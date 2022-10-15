关于Jibx编译并生成building.xml, pojo.xsd的问题：

命令参数是：
java -cp D:\mvnRepo\com\thoughtworks\qdox\qdox\1.12.1\qdox-1.12.1.jar;D:\mvnRepo\log4j\log4j\1.2.17\log4j-1.2.17.jar;D:\mvnRepo\org\jibx\jibx-schema\1.2.6\jibx-schema-1.2.6.jar;D:\mvnRepo\org\apache\bcel\bcel\5.2\bcel-5.2.jar;D:\mvnRepo\org\jibx\jibx-run\1.2.6\jibx-run-1.2.6.jar;D:\mvnRepo\org\jibx\jibx-bind\1.2.6\jibx-bind-1.2.6.jar;D:\mvnRepo\org\jibx\jibx-tools\1.2.6\jibx-tools-1.2.6.jar org.jibx.binding.generator.BindGen -b binding.xml com.cyh.xml.pojo.Order

没成功：
  估计是由于JDK原因，无法编译
  我自己本机是1.8，好像版本有点高，运行报错
  降为1.7或者可以，但是我没有尝试，不想弄


和使用Java代码 GenerateXmlXsd 直接运行报一样的错：

Exception in thread "main" java.lang.IllegalStateException: Internal error: instance signatures not found for class java.lang.String
	at org.jibx.binding.model.ClassWrapper.isImplements(ClassWrapper.java:136)
	at org.jibx.custom.classes.SharedValueBase.fillType(SharedValueBase.java:367)
	at org.jibx.custom.classes.ValueCustom.fillDetails(ValueCustom.java:316)
	at org.jibx.custom.classes.ClassCustom.apply(ClassCustom.java:800)
	at org.jibx.custom.classes.GlobalCustom.addClassCustomization(GlobalCustom.java:377)
	at org.jibx.binding.generator.BindGen.isValueClass(BindGen.java:134)
	at org.jibx.binding.generator.BindGen.expandReferences(BindGen.java:225)
	at org.jibx.binding.generator.BindGen.findReferences(BindGen.java:1010)
	at org.jibx.binding.generator.BindGen.generate(BindGen.java:1124)
	at org.jibx.binding.generator.BindGen.main(BindGen.java:1302)
	at GenerateXmlXsd.main(GenerateXmlXsd.java:13)


