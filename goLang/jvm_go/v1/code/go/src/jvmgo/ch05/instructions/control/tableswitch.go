package control

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/*
tableswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4
jump offsets...
*/
// Access jump table by index and jump
/**
java switch-case有2种实现方式
1. 如果case的值可以编码成一个索引表，则实现tableswitch指令
2. 否则则实现成lookupswitch指令
*/
type TABLE_SWITCH struct {
	defaultOffset int32   //switch - case 的默认跳转偏移量
	low           int32   //case的最小取值范围
	high          int32   //case的最大取值范围
	jumpOffsets   []int32 //索引表，对应各种case情况下，执行跳转所需的字节码偏移量
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	base.Branch(frame, offset)
}
