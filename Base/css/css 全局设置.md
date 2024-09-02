```css
  box-shadow: 12px 12px 18px rgba(155, 196, 255, 0.42),
    inset 10px 10px 11px rgba(250, 252, 255, 0.48),
    inset -10px -10px 15px rgba(46, 129, 255, 0.22);
```

## 一些全局设置
```
* {
  box-sizing: border-box;
  padding: 0;
  margin: 0;
  font-family: Verdana, "PingFang SC", "Microsoft Yahei", sans-serif;
}

```

### 背景颜色渐变
```css
 background: linear-gradient(
          to bottom,
          hsl(0deg, 100%, 50%, 0.4) 0%,
          hsl(40deg, 100%, 50%) 20%,
          hsl(80deg, 100%, 50%) 30%,
          hsl(120deg, 100%, 50%) 40%,
          hsl(180deg, 100%, 50%) 50%,
          hsl(250deg, 100%, 50%) 80%,
          hsl(320deg, 100%, 50%, 0.4) 100%
        );
```