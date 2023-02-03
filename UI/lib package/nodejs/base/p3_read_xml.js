/**
 * 我讨厌xml 都用node javascript了 还扯什么xml 给老子滚！
 */
var xmlreader = require("xmlreader");
var fs = require("fs");
var xml_string = '<response id="1" shop="aldi">'			+ 		'This is some other content'			+		'<who name="james">James May</who>'			+ 		'<who name="sam">'			+			'Sam Decrock'			+			'<location>Belgium</location>'			+		'</who>'			+ 		'<who name="jack">Jack Johnsen</who>'			+		'<games age="6">'			+			'<game>Some great game</game>'			+			'<game>Some other great game</game>'			+		'</games>'			+		'<note>These are some notes</note>'			+	'</response>';

xmlreader.read(xml_string, function(errors, response
){	if(null !== errors )
  {
    console.log(errors)
    	return;
    	}
console.log( response.response );
console.log( response.response.text() );});


// npm install xml2js
var parseString = require('xml2js').parseString;
var xml = '<?xml version="1.0" encoding="UTF-8" ?><business><company>Code Blog</company><owner>Nic Raboy</owner><employee><firstname>Nic</firstname><lastname>Raboy</lastname></employee><employee><firstname>Maria</firstname><lastname>Campos</lastname></employee></business>';
parseString(xml, function (err, result) {
    console.dir(JSON.stringify(result));
});


//另外还有一个使用C语言写的xml解析器 node-expat，性能更好，
//不过使用也很“底层“, 对性能有一定要求的应用可以尝试一下：
