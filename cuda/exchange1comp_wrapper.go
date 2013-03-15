package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var addexchange1comp_code cu.Function

type addexchange1comp_args struct {
	arg_h  unsafe.Pointer
	arg_m  unsafe.Pointer
	arg_wx float32
	arg_wy float32
	arg_wz float32
	arg_N0 int
	arg_N1 int
	arg_N2 int
	argptr [8]unsafe.Pointer
}

// Wrapper for addexchange1comp CUDA kernel, asynchronous.
func k_addexchange1comp_async(h unsafe.Pointer, m unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config, str cu.Stream) {
	if addexchange1comp_code == 0 {
		addexchange1comp_code = fatbinLoad(addexchange1comp_map, "addexchange1comp")
	}

	var a addexchange1comp_args

	a.arg_h = h
	a.argptr[0] = unsafe.Pointer(&a.arg_h)
	a.arg_m = m
	a.argptr[1] = unsafe.Pointer(&a.arg_m)
	a.arg_wx = wx
	a.argptr[2] = unsafe.Pointer(&a.arg_wx)
	a.arg_wy = wy
	a.argptr[3] = unsafe.Pointer(&a.arg_wy)
	a.arg_wz = wz
	a.argptr[4] = unsafe.Pointer(&a.arg_wz)
	a.arg_N0 = N0
	a.argptr[5] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[6] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[7] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(addexchange1comp_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for addexchange1comp CUDA kernel, synchronized.
func k_addexchange1comp(h unsafe.Pointer, m unsafe.Pointer, wx float32, wy float32, wz float32, N0 int, N1 int, N2 int, cfg *config) {
	str := stream()
	k_addexchange1comp_async(h, m, wx, wy, wz, N0, N1, N2, cfg, str)
	syncAndRecycle(str)
}

var addexchange1comp_map = map[int]string{0: "",
	20: addexchange1comp_ptx_20,
	30: addexchange1comp_ptx_30,
	35: addexchange1comp_ptx_35}

const (
	addexchange1comp_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry addexchange1comp(
	.param .u64 addexchange1comp_param_0,
	.param .u64 addexchange1comp_param_1,
	.param .f32 addexchange1comp_param_2,
	.param .f32 addexchange1comp_param_3,
	.param .f32 addexchange1comp_param_4,
	.param .u32 addexchange1comp_param_5,
	.param .u32 addexchange1comp_param_6,
	.param .u32 addexchange1comp_param_7
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<71>;
	.reg .f32 	%f<26>;
	.reg .s64 	%rd<21>;


	ld.param.u64 	%rd2, [addexchange1comp_param_0];
	ld.param.u64 	%rd3, [addexchange1comp_param_1];
	ld.param.f32 	%f5, [addexchange1comp_param_2];
	ld.param.f32 	%f6, [addexchange1comp_param_3];
	ld.param.f32 	%f7, [addexchange1comp_param_4];
	ld.param.u32 	%r26, [addexchange1comp_param_5];
	ld.param.u32 	%r27, [addexchange1comp_param_6];
	ld.param.u32 	%r28, [addexchange1comp_param_7];
	.loc 2 10 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 2 11 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 13 1
	setp.lt.s32 	%p1, %r8, %r28;
	setp.lt.s32 	%p2, %r4, %r27;
	and.pred  	%p3, %p2, %p1;
	.loc 2 17 1
	setp.gt.s32 	%p4, %r26, 0;
	.loc 2 13 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_5;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 23 1
	add.s32 	%r30, %r8, -1;
	mov.u32 	%r70, 0;
	.loc 3 238 5
	max.s32 	%r31, %r30, %r70;
	.loc 2 24 1
	add.s32 	%r32, %r28, -1;
	add.s32 	%r33, %r8, 1;
	.loc 3 210 5
	min.s32 	%r34, %r33, %r32;
	.loc 2 27 1
	add.s32 	%r35, %r4, -1;
	.loc 3 238 5
	max.s32 	%r36, %r35, %r70;
	.loc 2 28 1
	add.s32 	%r37, %r27, -1;
	add.s32 	%r38, %r4, 1;
	.loc 3 210 5
	min.s32 	%r39, %r38, %r37;
	.loc 2 17 1
	mad.lo.s32 	%r69, %r39, %r28, %r8;
	mad.lo.s32 	%r68, %r36, %r28, %r8;
	mad.lo.s32 	%r67, %r28, %r4, %r34;
	mad.lo.s32 	%r66, %r28, %r4, %r31;
	mad.lo.s32 	%r65, %r28, %r4, %r8;
	cvta.to.global.u64 	%rd4, %rd2;

BB0_2:
	.loc 2 20 1
	mul.wide.s32 	%rd5, %r65, 4;
	add.s64 	%rd1, %rd4, %rd5;
	cvta.to.global.u64 	%rd6, %rd3;
	.loc 2 21 1
	add.s64 	%rd7, %rd6, %rd5;
	.loc 2 23 1
	mul.wide.s32 	%rd8, %r66, 4;
	add.s64 	%rd9, %rd6, %rd8;
	.loc 2 24 1
	mul.wide.s32 	%rd10, %r67, 4;
	add.s64 	%rd11, %rd6, %rd10;
	.loc 2 23 1
	ld.global.f32 	%f8, [%rd9];
	.loc 2 21 1
	ld.global.f32 	%f1, [%rd7];
	.loc 2 25 1
	sub.f32 	%f9, %f8, %f1;
	.loc 2 24 1
	ld.global.f32 	%f10, [%rd11];
	.loc 2 25 1
	sub.f32 	%f11, %f10, %f1;
	add.f32 	%f12, %f9, %f11;
	.loc 2 20 1
	ld.global.f32 	%f13, [%rd1];
	.loc 2 25 1
	fma.rn.f32 	%f14, %f12, %f7, %f13;
	.loc 2 27 1
	mul.wide.s32 	%rd12, %r68, 4;
	add.s64 	%rd13, %rd6, %rd12;
	.loc 2 28 1
	mul.wide.s32 	%rd14, %r69, 4;
	add.s64 	%rd15, %rd6, %rd14;
	.loc 2 27 1
	ld.global.f32 	%f15, [%rd13];
	.loc 2 29 1
	sub.f32 	%f16, %f15, %f1;
	.loc 2 28 1
	ld.global.f32 	%f17, [%rd15];
	.loc 2 29 1
	sub.f32 	%f18, %f17, %f1;
	add.f32 	%f19, %f16, %f18;
	fma.rn.f32 	%f25, %f19, %f6, %f14;
	setp.eq.s32 	%p6, %r26, 1;
	.loc 2 32 1
	@%p6 bra 	BB0_4;

	.loc 2 17 18
	add.s32 	%r48, %r70, 1;
	.loc 2 33 1
	add.s32 	%r49, %r26, -1;
	.loc 3 210 5
	min.s32 	%r50, %r48, %r49;
	.loc 2 33 1
	mad.lo.s32 	%r55, %r50, %r27, %r4;
	mad.lo.s32 	%r56, %r55, %r28, %r8;
	mul.wide.s32 	%rd17, %r56, 4;
	add.s64 	%rd18, %rd6, %rd17;
	.loc 2 34 1
	add.s32 	%r57, %r70, -1;
	mov.u32 	%r58, 0;
	.loc 3 238 5
	max.s32 	%r59, %r57, %r58;
	.loc 2 34 1
	mad.lo.s32 	%r60, %r59, %r27, %r4;
	mad.lo.s32 	%r61, %r60, %r28, %r8;
	mul.wide.s32 	%rd19, %r61, 4;
	add.s64 	%rd20, %rd6, %rd19;
	.loc 2 33 1
	ld.global.f32 	%f20, [%rd18];
	.loc 2 35 1
	sub.f32 	%f21, %f20, %f1;
	.loc 2 34 1
	ld.global.f32 	%f22, [%rd20];
	.loc 2 35 1
	sub.f32 	%f23, %f22, %f1;
	add.f32 	%f24, %f21, %f23;
	fma.rn.f32 	%f25, %f24, %f5, %f25;

BB0_4:
	.loc 2 38 1
	st.global.f32 	[%rd1], %f25;
	.loc 2 17 1
	mad.lo.s32 	%r69, %r28, %r27, %r69;
	mad.lo.s32 	%r68, %r28, %r27, %r68;
	mad.lo.s32 	%r67, %r28, %r27, %r67;
	mad.lo.s32 	%r66, %r28, %r27, %r66;
	mad.lo.s32 	%r65, %r28, %r27, %r65;
	.loc 2 17 18
	add.s32 	%r70, %r70, 1;
	.loc 2 17 1
	setp.lt.s32 	%p7, %r70, %r26;
	@%p7 bra 	BB0_2;

BB0_5:
	.loc 2 40 2
	ret;
}


`
	addexchange1comp_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry addexchange1comp(
	.param .u64 addexchange1comp_param_0,
	.param .u64 addexchange1comp_param_1,
	.param .f32 addexchange1comp_param_2,
	.param .f32 addexchange1comp_param_3,
	.param .f32 addexchange1comp_param_4,
	.param .u32 addexchange1comp_param_5,
	.param .u32 addexchange1comp_param_6,
	.param .u32 addexchange1comp_param_7
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<68>;
	.reg .f32 	%f<26>;
	.reg .s64 	%rd<21>;


	ld.param.u64 	%rd2, [addexchange1comp_param_0];
	ld.param.u64 	%rd3, [addexchange1comp_param_1];
	ld.param.f32 	%f5, [addexchange1comp_param_2];
	ld.param.f32 	%f6, [addexchange1comp_param_3];
	ld.param.f32 	%f7, [addexchange1comp_param_4];
	ld.param.u32 	%r27, [addexchange1comp_param_5];
	ld.param.u32 	%r28, [addexchange1comp_param_6];
	ld.param.u32 	%r29, [addexchange1comp_param_7];
	.loc 2 10 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 2 11 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 13 1
	setp.lt.s32 	%p1, %r8, %r29;
	setp.lt.s32 	%p2, %r4, %r28;
	and.pred  	%p3, %p2, %p1;
	.loc 2 17 1
	setp.gt.s32 	%p4, %r27, 0;
	.loc 2 13 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_5;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 23 1
	add.s32 	%r31, %r8, -1;
	mov.u32 	%r67, 0;
	.loc 3 238 5
	max.s32 	%r32, %r31, %r67;
	.loc 2 24 1
	add.s32 	%r33, %r29, -1;
	add.s32 	%r34, %r8, 1;
	.loc 3 210 5
	min.s32 	%r35, %r34, %r33;
	.loc 2 27 1
	add.s32 	%r36, %r4, -1;
	.loc 3 238 5
	max.s32 	%r37, %r36, %r67;
	.loc 2 28 1
	add.s32 	%r38, %r28, -1;
	add.s32 	%r39, %r4, 1;
	.loc 3 210 5
	min.s32 	%r40, %r39, %r38;
	.loc 2 17 1
	mad.lo.s32 	%r66, %r40, %r29, %r8;
	mul.lo.s32 	%r10, %r29, %r28;
	mad.lo.s32 	%r65, %r37, %r29, %r8;
	mad.lo.s32 	%r64, %r29, %r4, %r35;
	mad.lo.s32 	%r63, %r29, %r4, %r32;
	mad.lo.s32 	%r62, %r29, %r4, %r8;
	cvta.to.global.u64 	%rd4, %rd2;

BB0_2:
	.loc 2 20 1
	mul.wide.s32 	%rd5, %r62, 4;
	add.s64 	%rd1, %rd4, %rd5;
	cvta.to.global.u64 	%rd6, %rd3;
	.loc 2 21 1
	add.s64 	%rd7, %rd6, %rd5;
	.loc 2 23 1
	mul.wide.s32 	%rd8, %r63, 4;
	add.s64 	%rd9, %rd6, %rd8;
	.loc 2 24 1
	mul.wide.s32 	%rd10, %r64, 4;
	add.s64 	%rd11, %rd6, %rd10;
	.loc 2 23 1
	ld.global.f32 	%f8, [%rd9];
	.loc 2 21 1
	ld.global.f32 	%f1, [%rd7];
	.loc 2 25 1
	sub.f32 	%f9, %f8, %f1;
	.loc 2 24 1
	ld.global.f32 	%f10, [%rd11];
	.loc 2 25 1
	sub.f32 	%f11, %f10, %f1;
	add.f32 	%f12, %f9, %f11;
	.loc 2 20 1
	ld.global.f32 	%f13, [%rd1];
	.loc 2 25 1
	fma.rn.f32 	%f14, %f12, %f7, %f13;
	.loc 2 27 1
	mul.wide.s32 	%rd12, %r65, 4;
	add.s64 	%rd13, %rd6, %rd12;
	.loc 2 28 1
	mul.wide.s32 	%rd14, %r66, 4;
	add.s64 	%rd15, %rd6, %rd14;
	.loc 2 27 1
	ld.global.f32 	%f15, [%rd13];
	.loc 2 29 1
	sub.f32 	%f16, %f15, %f1;
	.loc 2 28 1
	ld.global.f32 	%f17, [%rd15];
	.loc 2 29 1
	sub.f32 	%f18, %f17, %f1;
	add.f32 	%f19, %f16, %f18;
	fma.rn.f32 	%f25, %f19, %f6, %f14;
	setp.eq.s32 	%p6, %r27, 1;
	.loc 2 32 1
	@%p6 bra 	BB0_4;

	.loc 2 17 18
	add.s32 	%r49, %r67, 1;
	.loc 2 33 1
	add.s32 	%r50, %r27, -1;
	.loc 3 210 5
	min.s32 	%r51, %r49, %r50;
	.loc 2 33 1
	mad.lo.s32 	%r52, %r51, %r28, %r4;
	mad.lo.s32 	%r53, %r52, %r29, %r8;
	mul.wide.s32 	%rd17, %r53, 4;
	add.s64 	%rd18, %rd6, %rd17;
	.loc 2 34 1
	add.s32 	%r54, %r67, -1;
	mov.u32 	%r55, 0;
	.loc 3 238 5
	max.s32 	%r56, %r54, %r55;
	.loc 2 34 1
	mad.lo.s32 	%r57, %r56, %r28, %r4;
	mad.lo.s32 	%r58, %r57, %r29, %r8;
	mul.wide.s32 	%rd19, %r58, 4;
	add.s64 	%rd20, %rd6, %rd19;
	.loc 2 33 1
	ld.global.f32 	%f20, [%rd18];
	.loc 2 35 1
	sub.f32 	%f21, %f20, %f1;
	.loc 2 34 1
	ld.global.f32 	%f22, [%rd20];
	.loc 2 35 1
	sub.f32 	%f23, %f22, %f1;
	add.f32 	%f24, %f21, %f23;
	fma.rn.f32 	%f25, %f24, %f5, %f25;

BB0_4:
	.loc 2 38 1
	st.global.f32 	[%rd1], %f25;
	.loc 2 17 1
	add.s32 	%r66, %r66, %r10;
	add.s32 	%r65, %r65, %r10;
	add.s32 	%r64, %r64, %r10;
	add.s32 	%r63, %r63, %r10;
	add.s32 	%r62, %r62, %r10;
	.loc 2 17 18
	add.s32 	%r67, %r67, 1;
	.loc 2 17 1
	setp.lt.s32 	%p7, %r67, %r27;
	@%p7 bra 	BB0_2;

BB0_5:
	.loc 2 40 2
	ret;
}


`
	addexchange1comp_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry addexchange1comp(
	.param .u64 addexchange1comp_param_0,
	.param .u64 addexchange1comp_param_1,
	.param .f32 addexchange1comp_param_2,
	.param .f32 addexchange1comp_param_3,
	.param .f32 addexchange1comp_param_4,
	.param .u32 addexchange1comp_param_5,
	.param .u32 addexchange1comp_param_6,
	.param .u32 addexchange1comp_param_7
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<61>;
	.reg .f32 	%f<26>;
	.reg .s64 	%rd<21>;


	ld.param.u64 	%rd2, [addexchange1comp_param_0];
	ld.param.u64 	%rd3, [addexchange1comp_param_1];
	ld.param.f32 	%f5, [addexchange1comp_param_2];
	ld.param.f32 	%f6, [addexchange1comp_param_3];
	ld.param.f32 	%f7, [addexchange1comp_param_4];
	ld.param.u32 	%r27, [addexchange1comp_param_5];
	ld.param.u32 	%r28, [addexchange1comp_param_6];
	ld.param.u32 	%r29, [addexchange1comp_param_7];
	.loc 3 10 1
	mov.u32 	%r1, %ctaid.x;
	mov.u32 	%r2, %ntid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r2, %r1, %r3;
	.loc 3 11 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 3 13 1
	setp.lt.s32 	%p1, %r8, %r29;
	setp.lt.s32 	%p2, %r4, %r28;
	and.pred  	%p3, %p2, %p1;
	.loc 3 17 1
	setp.gt.s32 	%p4, %r27, 0;
	.loc 3 13 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_5;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 23 1
	add.s32 	%r31, %r8, -1;
	mov.u32 	%r60, 0;
	.loc 4 238 5
	max.s32 	%r32, %r31, %r60;
	.loc 3 24 1
	add.s32 	%r33, %r29, -1;
	add.s32 	%r34, %r8, 1;
	.loc 4 210 5
	min.s32 	%r35, %r34, %r33;
	.loc 3 27 1
	add.s32 	%r36, %r4, -1;
	.loc 4 238 5
	max.s32 	%r37, %r36, %r60;
	.loc 3 28 1
	add.s32 	%r38, %r28, -1;
	add.s32 	%r39, %r4, 1;
	.loc 4 210 5
	min.s32 	%r40, %r39, %r38;
	.loc 3 17 1
	mad.lo.s32 	%r59, %r40, %r29, %r8;
	mul.lo.s32 	%r10, %r29, %r28;
	mad.lo.s32 	%r58, %r37, %r29, %r8;
	mad.lo.s32 	%r57, %r29, %r4, %r35;
	mad.lo.s32 	%r56, %r29, %r4, %r32;
	mad.lo.s32 	%r55, %r29, %r4, %r8;
	cvta.to.global.u64 	%rd4, %rd2;

BB2_2:
	.loc 3 20 1
	mul.wide.s32 	%rd5, %r55, 4;
	add.s64 	%rd1, %rd4, %rd5;
	ld.global.f32 	%f8, [%rd1];
	cvta.to.global.u64 	%rd6, %rd3;
	.loc 3 21 1
	add.s64 	%rd7, %rd6, %rd5;
	ld.global.nc.f32 	%f1, [%rd7];
	.loc 3 23 1
	mul.wide.s32 	%rd8, %r56, 4;
	add.s64 	%rd9, %rd6, %rd8;
	ld.global.nc.f32 	%f9, [%rd9];
	.loc 3 24 1
	mul.wide.s32 	%rd10, %r57, 4;
	add.s64 	%rd11, %rd6, %rd10;
	ld.global.nc.f32 	%f10, [%rd11];
	.loc 3 25 1
	sub.f32 	%f11, %f9, %f1;
	sub.f32 	%f12, %f10, %f1;
	add.f32 	%f13, %f11, %f12;
	fma.rn.f32 	%f14, %f13, %f7, %f8;
	.loc 3 27 1
	mul.wide.s32 	%rd12, %r58, 4;
	add.s64 	%rd13, %rd6, %rd12;
	ld.global.nc.f32 	%f15, [%rd13];
	.loc 3 28 1
	mul.wide.s32 	%rd14, %r59, 4;
	add.s64 	%rd15, %rd6, %rd14;
	ld.global.nc.f32 	%f16, [%rd15];
	.loc 3 29 1
	sub.f32 	%f17, %f15, %f1;
	sub.f32 	%f18, %f16, %f1;
	add.f32 	%f19, %f17, %f18;
	fma.rn.f32 	%f25, %f19, %f6, %f14;
	setp.eq.s32 	%p6, %r27, 1;
	.loc 3 32 1
	@%p6 bra 	BB2_4;

	.loc 3 17 18
	add.s32 	%r44, %r60, 1;
	.loc 3 33 1
	add.s32 	%r45, %r27, -1;
	.loc 4 210 5
	min.s32 	%r46, %r44, %r45;
	.loc 3 33 1
	mad.lo.s32 	%r47, %r46, %r28, %r4;
	mad.lo.s32 	%r48, %r47, %r29, %r8;
	mul.wide.s32 	%rd17, %r48, 4;
	add.s64 	%rd18, %rd6, %rd17;
	ld.global.nc.f32 	%f20, [%rd18];
	.loc 3 34 1
	add.s32 	%r49, %r60, -1;
	mov.u32 	%r50, 0;
	.loc 4 238 5
	max.s32 	%r51, %r49, %r50;
	.loc 3 34 1
	mad.lo.s32 	%r52, %r51, %r28, %r4;
	mad.lo.s32 	%r53, %r52, %r29, %r8;
	mul.wide.s32 	%rd19, %r53, 4;
	add.s64 	%rd20, %rd6, %rd19;
	ld.global.nc.f32 	%f21, [%rd20];
	.loc 3 35 1
	sub.f32 	%f22, %f20, %f1;
	sub.f32 	%f23, %f21, %f1;
	add.f32 	%f24, %f22, %f23;
	fma.rn.f32 	%f25, %f24, %f5, %f25;

BB2_4:
	.loc 3 38 1
	st.global.f32 	[%rd1], %f25;
	.loc 3 17 1
	add.s32 	%r59, %r59, %r10;
	add.s32 	%r58, %r58, %r10;
	add.s32 	%r57, %r57, %r10;
	add.s32 	%r56, %r56, %r10;
	add.s32 	%r55, %r55, %r10;
	.loc 3 17 18
	add.s32 	%r60, %r60, 1;
	.loc 3 17 1
	setp.lt.s32 	%p7, %r60, %r27;
	@%p7 bra 	BB2_2;

BB2_5:
	.loc 3 40 2
	ret;
}


`
)
