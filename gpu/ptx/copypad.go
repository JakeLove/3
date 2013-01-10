package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["copypad"] = COPYPAD }

const COPYPAD = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_000018e7_00000000-9_copypad.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/mx3/gpu/ptx/copypad.cu"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"

.visible .entry copypad(
	.param .u64 copypad_param_0,
	.param .u32 copypad_param_1,
	.param .u32 copypad_param_2,
	.param .u32 copypad_param_3,
	.param .u64 copypad_param_4,
	.param .u32 copypad_param_5,
	.param .u32 copypad_param_6,
	.param .u32 copypad_param_7,
	.param .u32 copypad_param_8,
	.param .u32 copypad_param_9,
	.param .u32 copypad_param_10
)
{
	.reg .pred 	%p<10>;
	.reg .s32 	%r<42>;
	.reg .f32 	%f<2>;
	.reg .s64 	%rd<9>;


	ld.param.u64 	%rd3, [copypad_param_0];
	ld.param.u32 	%r18, [copypad_param_1];
	ld.param.u32 	%r19, [copypad_param_2];
	ld.param.u32 	%r20, [copypad_param_3];
	ld.param.u64 	%rd4, [copypad_param_4];
	ld.param.u32 	%r21, [copypad_param_5];
	ld.param.u32 	%r22, [copypad_param_6];
	ld.param.u32 	%r23, [copypad_param_7];
	ld.param.u32 	%r24, [copypad_param_8];
	ld.param.u32 	%r25, [copypad_param_9];
	ld.param.u32 	%r26, [copypad_param_10];
	cvta.to.global.u64 	%rd1, %rd3;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 13 1
	mov.u32 	%r1, %ntid.y;
	mov.u32 	%r2, %ctaid.y;
	mov.u32 	%r3, %tid.y;
	mad.lo.s32 	%r27, %r1, %r2, %r3;
	.loc 2 14 1
	mov.u32 	%r4, %ntid.x;
	mov.u32 	%r5, %ctaid.x;
	mov.u32 	%r6, %tid.x;
	mad.lo.s32 	%r28, %r4, %r5, %r6;
	.loc 2 16 1
	setp.ge.s32 	%p1, %r28, %r23;
	setp.ge.s32 	%p2, %r27, %r22;
	or.pred  	%p3, %p1, %p2;
	setp.ge.s32 	%p4, %r27, %r19;
	or.pred  	%p5, %p3, %p4;
	setp.ge.s32 	%p6, %r28, %r20;
	or.pred  	%p7, %p5, %p6;
	@%p7 bra 	BB0_4;

	.loc 3 210 5
	min.s32 	%r7, %r21, %r18;
	.loc 2 24 1
	setp.lt.s32 	%p8, %r7, 1;
	@%p8 bra 	BB0_4;

	add.s32 	%r30, %r6, %r26;
	mad.lo.s32 	%r31, %r4, %r5, %r30;
	add.s32 	%r32, %r3, %r25;
	mad.lo.s32 	%r33, %r1, %r2, %r32;
	mad.lo.s32 	%r34, %r24, %r19, %r33;
	mad.lo.s32 	%r40, %r20, %r34, %r31;
	mul.lo.s32 	%r9, %r20, %r19;
	mad.lo.s32 	%r39, %r23, %r27, %r28;
	mul.lo.s32 	%r11, %r23, %r22;
	mov.u32 	%r41, 0;

BB0_3:
	.loc 2 26 1
	mul.wide.s32 	%rd5, %r39, 4;
	add.s64 	%rd6, %rd2, %rd5;
	mul.wide.s32 	%rd7, %r40, 4;
	add.s64 	%rd8, %rd1, %rd7;
	ld.global.f32 	%f1, [%rd6];
	st.global.f32 	[%rd8], %f1;
	.loc 2 24 1
	add.s32 	%r40, %r40, %r9;
	add.s32 	%r39, %r39, %r11;
	.loc 2 24 52
	add.s32 	%r41, %r41, 1;
	.loc 2 24 1
	setp.lt.s32 	%p9, %r41, %r7;
	@%p9 bra 	BB0_3;

BB0_4:
	.loc 2 28 2
	ret;
}


`
