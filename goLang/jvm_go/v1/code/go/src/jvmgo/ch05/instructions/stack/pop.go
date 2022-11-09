package stack

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/*
	栈指令 POP
	将操作数栈的顶部变量弹出
*/
// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
