// 改变 HTML 内容
document.getElementById(id).innerHTML= ""

// 改变 HTML 属性
document.getElementById(id).attribute=""

// 改变 HTML 样式
document.getElementById(id).style.property=新样式
document.getElementById('id1').style.color='red'

document.getElementsByTagName('body').className = 'snow-container'; //设置为新的
document.getElementsByTagName('body').className += ' snow-container'; //在原来的后面加这个
document.getElementsByTagName('body').classList.add("snow-container"); //与第一个等价

document.getElementById("myBtn").addEventListener("click", displayDate);

element.addEventListener("click", function(){ alert("Hello World!"); });


/**
 * <div>
 *  <p>kkkkkk</p>
 * </div>
 */
var para = document.createElement("p");
var node = document.createTextNode("这是一个新的段落。");
para.appendChild(node);





