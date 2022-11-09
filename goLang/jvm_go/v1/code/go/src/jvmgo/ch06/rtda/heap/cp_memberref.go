package heap

import "jvmgo/ch06/classfile"

//字段符号引用 用于存放字段和方法符号引用共有的信息
type MemberRef struct {
	SymRef
	name       string //名称
	descriptor string //类型描述符
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
