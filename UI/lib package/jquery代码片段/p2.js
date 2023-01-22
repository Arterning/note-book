$("button").click(function(){
    $("li").each(function(){
      alert($(this).text())
    });
  });



// element - 当前的元素（也可使用 "this" 选择器
$(selector).each(function(index,element));


append() - 在被选元素的结尾插入内容
prepend() - 在被选元素的开头插入内容
after() - 在被选元素之后插入内容
before() - 在被选元素之前插入内容