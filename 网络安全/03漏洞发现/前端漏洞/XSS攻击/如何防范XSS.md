XSS漏洞修复建议
因为XSS漏洞涉及输入和输出两部分，所以其修复也分为两种。

过滤输入的数据，包括 '、"、<、>、on* 等非法字符。
对输出到页面的数据进行相应的编码转换，包括HTML实体编码、JavaScript编码等。



比如在React中，如果我们需要展示用户的评论，我们需要这样做

```js
const App = () => {
  const data = `lorem <b onmouseover="alert('mouseover');">ipsum</b>`;

  return <div dangerouslySetInnerHTML={{ __html: data }} />;
};
```