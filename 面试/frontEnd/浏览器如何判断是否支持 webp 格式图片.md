
### 通过canvas来判断（这个比较常用）



```js
 /**
   * 判断浏览器是否支持webp格式
   * @return {[Boolean]}
   */
  function isSupportWebp () {
    try {
      return document.createElement('canvas').toDataURL('image/webp').indexOf('data:image/webp') === 0
    } catch (err) {
      return false
    }
  }

```

创建一个canvas元素，然后把它转成`image/webp`格式的data_url,如果这个data_url里面包含`webp`,则代表当前浏览器支持`webp`格式， 反之则不支持。

  

