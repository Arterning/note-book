
# vscode-snippets：快速初始化一个React函数组件


Ctrl + Shift + P
snippets入口：首选项->Configure User Snippets


type this 

```js
{
    "React Typescript function component": {
            "prefix": "tsfc",
            "body": [
                    "import React from 'react';",
                    "", // 空白行
                    "interface I${1}Props {", // ${1}变量，相同编号表示同样的数值
                    "",
                    "}",
                    "",
                    "export const ${1}: React.FC<I${1}Props> = (props) => {",
                    "",
                    "    return (",
                    "",
                    "    );",
                    "};"
            ],
            "description": "react function component"
    },
    "React JavaScript function component": {
        "prefix": "jsfc",
        "body": [
                "import React from 'react';",
                
                "export const ${1} = (props) => {",
                "",
                "    return (",
                "      <div>",
                "          Here is ${1} Page",
                "      </div>",
                "    );",
                "};"
        ],
        "description": "react function component",
	"ES function": {
			"prefix": "df",
			"body": [
					"export const ${1} = () => {",
					"",
					"",
					"};"
			],
			"description": "react function component"
			},
}
}


```

[一个前端开发](https://juejin.cn/user/2400989128436141/posts)