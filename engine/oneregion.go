package engine

import (
	"fmt"
	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
)

func sInRegion(q OutputQuantity, r int) ScalarOutput {
	return AsScalarOutput(inRegion(q, r))
}

func vInRegion(q OutputQuantity, r int) VectorOutput {
	return AsVectorOutput(inRegion(q, r))
}

func sOneRegion(q OutputQuantity, r int) *sOneReg {
	util.Argument(q.NComp() == 1)
	return &sOneReg{oneReg{q, r}}
}

func vOneRegion(q OutputQuantity, r int) *vOneReg {
	util.Argument(q.NComp() == 3)
	return &vOneReg{oneReg{q, r}}
}

type sOneReg struct{ oneReg }

func (q *sOneReg) Average() float64 { return q.average()[0] }

type vOneReg struct{ oneReg }

func (q *vOneReg) Average() data.Vector { return unslice(q.average()) }

// represents a new quantity equal to q in the given region, 0 outside.
type oneReg struct {
	parent OutputQuantity
	region int
}

func inRegion(q OutputQuantity, region int) OutputQuantity {
	return &oneReg{q, region}
}

func (q *oneReg) NComp() int       { return q.parent.NComp() }
func (q *oneReg) Name() string     { return fmt.Sprint(q.parent.Name(), ".region", q.region) }
func (q *oneReg) Unit() string     { return q.parent.Unit() }
func (q *oneReg) Mesh() *data.Mesh { return q.parent.Mesh() }

// returns a new slice equal to q in the given region, 0 outside.
func (q *oneReg) Slice() (*data.Slice, bool) {
	src, r := q.parent.Slice()
	if r {
		defer cuda.Recycle(src)
	}
	out := cuda.Buffer(q.NComp(), q.Mesh().Size())
	cuda.RegionSelect(out, src, regions.Gpu(), byte(q.region))
	return out, true
}

func (q *oneReg) average() []float64 {
	slice, r := q.Slice()
	if r {
		defer cuda.Recycle(slice)
	}
	avg := sAverageUniverse(slice)
	sDiv(avg, regions.volume(q.region))
	return avg
}

func (q *oneReg) Average() []float64 { return q.average() }

// slice division
func sDiv(v []float64, x float64) {
	for i := range v {
		v[i] /= x
	}
}
