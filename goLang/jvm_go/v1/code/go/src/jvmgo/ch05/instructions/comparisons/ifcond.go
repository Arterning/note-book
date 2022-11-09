package comparisons

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Branch if int comparison with zero succeeds
// if<cond>指令将操作数栈顶的int变量弹出，然后和0进行比较，满足条件则跳转

//因为是条件指令，所以传入的是base.BranchInstruction 其中包括了跳转偏移量
//ifeq : if equal 0, then jump
type IFEQ struct{ base.BranchInstruction }

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

//ifne : if not equal 0, then jump
type IFNE struct{ base.BranchInstruction }

func (self *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

//iflt : if left than 0, then jump
type IFLT struct{ base.BranchInstruction }

func (self *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

//ifle : if left or equal to 0, then jump
type IFLE struct{ base.BranchInstruction }

func (self *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

//ifgt : if greater to 0, then jump
type IFGT struct{ base.BranchInstruction }

func (self *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

//ifge : if greater or equal to 0, them jump
type IFGE struct{ base.BranchInstruction }

func (self *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}
