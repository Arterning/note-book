$(function () {
    var $li = $(".wrap>ul>li")
    $li.mouseenter(function () { $(this).children("ul").stop().slideDown(); })
    $li.mouseleave(function () { $(this).children("ul").stop().slideUp(); })
});


