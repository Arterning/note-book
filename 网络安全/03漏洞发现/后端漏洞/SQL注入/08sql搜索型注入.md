其实后端代码就是
select .. from .. like '%search%'
我们只需要根据这种语句来构造闭合就行了
所以payload就是

kobe%' or 1=1 #