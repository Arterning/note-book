# maven的pom文件<relativePath/>作用


maven项目，父子模块，父工程pom中的relativePath。

```text
<parent>
	<groupId>***</groupId>
	<artifactId>***</artifactId>
	<version>***</version>
	<relativePath/>
</parent>
```

这个<parent>下面的<relativePath>属性：**parent的pom文件的路径**

**1.默认值：**

```text
<parent>
	<groupId>***</groupId>
	<artifactId>***</artifactId>
	<version>***</version>

</parent>
```

默认我们不用写<relativePath>，那默认值就是 ../pom.xml，会从本地路径中获取parent的pom

**2.<relativePath/>:**

设定一个空值将始终从仓库中获取，不从本地路径获取.

```text
<parent>
	<groupId>***</groupId>
	<artifactId>***</artifactId>
	<version>***</version>
	<relativePath/>
</parent>
```

**3.<relativePath>某个pom的路径<relativePath/>:**

指定本地的路径，从本地路径获取parent的pom。

```text
<parent>
	<groupId>***</groupId>
	<artifactId>***</artifactId>
	<version>***</version>
        <relativePath>***<relativePath/>
</parent>
```

  

_为天地立心，为生民立命，为往圣继绝学，为万世开太平。_