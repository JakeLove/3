package gpu

import (
	"code.google.com/p/mx3/core"
	"code.google.com/p/mx3/gpu/ptx"
	"code.google.com/p/mx3/nimble"
)

type Stencil struct {
	in     nimble.RChanN
	out    nimble.ChanN
	Weight [7]float32
}

func NewStencil(tag, unit string, in nimble.ChanN, Weight [7]float32) *Stencil {
	core.Assert(in.NComp() == 1)
	r := in.NewReader() // TODO: buffer
	w := nimble.MakeChanN(1, tag, unit, r.Mesh(), in.MemType(), -1)
	return &Stencil{r, w, Weight}
}

func (s *Stencil) Exec() {
	N := s.out.Mesh().NCell()
	dst := s.out.WriteNext(N)
	Memset(dst, 0)
	src := s.in.ReadNext(N)
	StencilAdd(dst, src, s.out.Mesh(), &s.Weight)
}

func (s *Stencil) Output() nimble.ChanN {
	return s.out
}

// StencilAdd adds to dst the stencil result.
func StencilAdd(dst, src nimble.Slice, mesh *nimble.Mesh, weight *[7]float32) {
	core.Assert(dst.Len() == src.Len() && src.Len() == mesh.NCell())
	core.Assert(dst.NComp() == 1 && src.NComp() == 1)

	size := mesh.Size()
	N0, N1, N2 := size[0], size[1], size[2]
	wrap := mesh.PBC()
	core.Assert(wrap == [3]int{0, 0, 0})
	gridDim, blockDim := Make2DConf(N2, N1)

	ptx.K_stencil3(dst.DevPtr(0), src.DevPtr(0),
		weight[0], weight[1], weight[2], weight[3], weight[4], weight[5], weight[6],
		wrap[0], wrap[1], wrap[2], N0, N1, N2, gridDim, blockDim)
}
