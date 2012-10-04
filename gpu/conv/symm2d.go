package conv

import (
	"github.com/barnex/cuda4/cu"
	"github.com/barnex/cuda4/safe"
	"nimble-cube/core"
	"nimble-cube/gpu"
)

type Symm2D struct {
	size        [3]int             // 3D size of the input/output data
	kernSize    [3]int             // Size of kernel and logical FFT size.
	fftKernSize [3]int             // Size of real, FFTed kernel
	n           int                // product of size
	input       [3]gpu.RChan       // TODO: fuse with input
	output      [3]gpu.Chan        // TODO: fuse with output
	fftRBuf     [2]safe.Float32s   // FFT input buf for FFT, shares storage with fftCBuf. 
	fftCBuf     [2]safe.Complex64s // FFT output buf, shares storage with fftRBuf
	//                                element [0] used for X; elements [0],[1] for Y, Z
	gpuFFTKern [3][3]safe.Float32s // FFT kernel on device: TODO: xfer if needed
	fwPlan     safe.FFT3DR2CPlan   // Forward FFT (1 component)
	bwPlan     safe.FFT3DC2RPlan   // Backward FFT (1 component)
	stream     cu.Stream           // 
	kern       [3][3][]float32     // Real-space kernel
	kernArr    [3][3][][][]float32 // Real-space kernel
	fftKern    [3][3][]float32     // FFT kernel on host
}

func (c *Symm2D) init() {
	core.Log("initializing 2D symmetric convolution")
	padded := c.kernSize

	{ // init FFT plans
		c.stream = cu.StreamCreate()
		c.fwPlan = safe.FFT3DR2C(padded[0], padded[1], padded[2])
		c.fwPlan.SetStream(c.stream)
		c.bwPlan = safe.FFT3DC2R(padded[0], padded[1], padded[2])
		c.bwPlan.SetStream(c.stream)
	}

	{ // init FFT kernel
		ffted := fftR2COutputSizeFloats(padded)
		realsize := ffted
		realsize[2] /= 2
		c.fftKernSize = realsize
		halfkern := realsize
		halfkern[1] = halfkern[1]/2 + 1
		fwPlan := c.fwPlan
		output := safe.MakeComplex64s(fwPlan.OutputLen())
		input := output.Float().Slice(0, fwPlan.InputLen())

		// upper triangular part
		for i := 0; i < 3; i++ {
			for j := i; j < 3; j++ {
				if c.kern[i][j] != nil { // ignore 0's
					input.CopyHtoD(c.kern[i][j])
					fwPlan.Exec(input, output)
					fwPlan.Stream().Synchronize() // !!
					c.fftKern[i][j] = make([]float32, prod(halfkern))
					scaleRealParts(c.fftKern[i][j], output.Float().Slice(0, prod(halfkern)*2), 1/float32(fwPlan.InputLen()))
					c.gpuFFTKern[i][j] = safe.MakeFloat32s(len(c.fftKern[i][j]))
					c.gpuFFTKern[i][j].CopyHtoD(c.fftKern[i][j])
				}
			}
		}
		output.Free()
	}

	{ // init device buffers
		for i := 0; i < 2; i++ {
			c.fftCBuf[i] = safe.MakeComplex64s(prod(fftR2COutputSizeFloats(c.kernSize)) / 2)
			c.fftRBuf[i] = c.fftCBuf[i].Float().Slice(0, prod(c.kernSize))
		}
	}
}

func (c *Symm2D) Run() {
	core.Log("running symmetric 2D convolution")
	gpu.LockCudaThread()
	c.init()

	padded := c.kernSize
	offset := [3]int{0, 0, 0}
	N1, N2 := c.fftKernSize[1], c.fftKernSize[2]

	for {

		// Convolution is separated into 
		// a 1D convolution for x
		// and a 2D convolution for yz.
		// so only 2 FFT buffers are then needed at the same time.

		// FFT x
		c.input[0].ReadNext(c.n)
		c.fftRBuf[0].MemsetAsync(0, c.stream) // copypad does NOT zero remainder.
		copyPad(c.fftRBuf[0], c.input[0].UnsafeData(), padded, c.size, offset, c.stream)
		c.fwPlan.Exec(c.fftRBuf[0], c.fftCBuf[0])
		//c.stream.Synchronize()
		c.input[0].ReadDone()

		// kern mul X	
		kernMulRSymm2Dx(c.fftCBuf[0], c.gpuFFTKern[0][0], N1, N2, c.stream)
		//c.stream.Synchronize()

		// bw FFT x
		c.output[0].WriteNext(c.n)
		c.bwPlan.Exec(c.fftCBuf[0], c.fftRBuf[0])
		copyPad(c.output[0].UnsafeData(), c.fftRBuf[0], c.size, padded, offset, c.stream)
		c.stream.Synchronize()
		c.output[0].WriteDone()

		// FW FFT yz
		// use FFT buffers 0 and 1 (not 1, 2) -> fftBuf[i-1]
		for i := 1; i < 3; i++ {
			c.input[i].ReadNext(c.n)
			c.fftRBuf[i-1].MemsetAsync(0, c.stream)
			copyPad(c.fftRBuf[i-1], c.input[i].UnsafeData(), padded, c.size, offset, c.stream)
			c.fwPlan.Exec(c.fftRBuf[i-1], c.fftCBuf[i-1])
			c.stream.Synchronize()
			c.input[i].ReadDone()
		}

		// kern mul yz
		kernMulRSymm2Dyz(c.fftCBuf[0], c.fftCBuf[1],
			c.gpuFFTKern[1][1], c.gpuFFTKern[2][2], c.gpuFFTKern[1][2],
			N1, N2, c.stream)
		c.stream.Synchronize()

		// BW FFT yz
		for i := 1; i < 3; i++ {
			c.output[i].WriteNext(c.n)
			c.bwPlan.Exec(c.fftCBuf[i-1], c.fftRBuf[i-1])
			copyPad(c.output[i].UnsafeData(), c.fftRBuf[i-1], c.size, padded, offset, c.stream)
			c.stream.Synchronize()
			c.output[i].WriteDone()
		}
	}
}

func(c*Symm2D)is2D()bool{
	return c.size[0] == 1
}

func(c*Symm2D)is3D()bool{
	return !c.is2D()
}

func NewSymm2D(size [3]int, kernel [3][3][][][]float32, input [3]gpu.RChan, output [3]gpu.Chan) *Symm2D {
	core.Assert(size[0] == 1) // 3D not supported
	c := new(Symm2D)
	c.size = size
	c.kernArr = kernel
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if kernel[i][j] != nil {
				c.kern[i][j] = core.Contiguous(kernel[i][j])
			}
		}
	}
	c.n = prod(size)
	c.kernSize = core.SizeOf(kernel[0][0])
	c.input = input
	c.output = output

	return c
	// TODO: self-test
}
