append() - 在被选元素的结尾插入内容
prepend() - 在被选元素的开头插入内容


after() - 在被选元素之后插入内容
before() - 在被选元素之前插入内容



// append/prepend 是在选择元素内部嵌入。
// after/before 是在元素外面追加。

/**
 * 
 * 
<p>
  <span class="s1">s1</span>
</p>
<script>
$("p").append('<span class="s2">s2</span>');
</script>

<p>
  <span class="s1">s1</span>
  <span class="s2">s2</span>
</p>
 */

/**
 * 参数可以是个 list:
 */
 var txt1="<b>I </b>";                    // 使用 HTML 创建元素
 var txt2=$("<i></i>").text("love ");     // 使用 jQuery 创建元素
 var txt3=document.createElement("big");  // 使用 DOM 创建元素
 txt3.innerHTML="jQuery!";
 $("img").after([txt1,txt2,txt3]);          // 在图片后添加文本