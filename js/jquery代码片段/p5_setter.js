// text不能解析html
$("#test1").text("Hello world!");

// html可以解析html
$("#test2").html("<b>Hello world!</b>");

// val专门设置input等表单值的value
$("#test3").val("RUNOOB");


$("#runoob").attr("href","http://www.runoob.com/jquery");

$("#runoob").attr({
    "href" : "http://www.runoob.com/jquery",
    "title" : "jQuery 教程"
});

$('#demo').attr('class','class1 class2');


$("p:first").addClass("intro");
$('#demo').removeClass('class1');


$("p").css("background-color");
$("p").css("background-color","yellow");
$("p").css({"background-color":"yellow","font-size":"200%"});
