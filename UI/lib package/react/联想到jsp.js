/**

由于工作的关系 用了很长时间的jsp
都说jsp是过时的东西 但是从其思想来看 
模板语法 指令 和嵌入代码 这种设计并没有过时

jsx无非就是在html中嵌入javascript代码

//模拟一些数据
const data = ['Angular','React','Vue']
//1.创建虚拟DOM
const VDOM = (
    <div>
        <h1>前端js框架列表</h1>
        <ul>
            {
                data.map((item,index)=>{
                    return <li key={index}>{item}</li>
                })
            }
        </ul>
    </div>
)
嵌入的代码必须使用{} 包裹起来

这一点其实就类似于JSP中的嵌入JAVA代码
<%
    ...
%>
可以想见 这种方式是非常灵活的

而JSP中的各种指令和模板 就类似于Vue中的指令和模板
比如
<c:if>
<c:forEach>
<c:choose>
<c:when>
等等
JSP也支持类似于Vue中的数据双向绑定
{varName}

所以我总结如下
JSP相当于React和Vue在后端的综合体
当初学习JSP还是很有价值的。。。
 * 
 */