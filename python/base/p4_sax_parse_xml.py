import xml.sax
 
class StudentHandler(xml.sax.ContentHandler):
    def __init__(self):
        self.id = ""
        self.name = ""
        self.age = ""
        self.gender = ""
 
    # 元素开始调用
    def startElement(self, tag, attributes):
        self.CurrentData = tag
        if tag == "student":
            stu_name = attributes["name"]
            print("stu_name:", stu_name)
 

def endElement(self, tag):
        if self.CurrentData == "id":
            print("id:", self.id)
        elif self.CurrentData == "name":
            print("name:", self.name)
        elif self.CurrentData == "age":
            print("age:", self.age)
        elif self.CurrentData == "gender":
            print("gender:", self.gender)
        self.CurrentData = ""
 
    # 读取字符时调用
def characters(self, content):
        if self.CurrentData == "id":
            self.id = content
        elif self.CurrentData == "name":
            self.name = content
        elif self.CurrentData == "age":
            self.age = content
        elif self.CurrentData == "gender":
            self.gender = content



if (__name__ == "__main__"):
    # 创建 XMLReader
    parser = xml.sax.make_parser()
    # 关闭命名空间
    parser.setFeature(xml.sax.handler.feature_namespaces, 0)
    # 重写 ContextHandler
    Handler = StudentHandler()
    parser.setContentHandler(Handler)
    parser.parse("test.xml")