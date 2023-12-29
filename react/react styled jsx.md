```css
<style jsx>{`
          .flip-card {
            width: 100%;
            height: 100%;
            display: inline-block;
            position: relative;
            transform-style: preserve-3d;
            transition: transform 0.2s;
          }
          
          .flip-card-front,
          .flip-card-back {
            position: absolute;
            width: 100%;
            height: 100%;
            backface-visibility: hidden;
          }
          
          .flip-card-front {
            z-index: 2;
            transform: rotateY(0);
          }
          
          .flip-card-back {
            transform: rotateY(180deg);
          }
          
          .flip-card.flipped {
            transform: rotateY(180deg);
          }
        `}</style>
```

style 表示后面是样式 花括号表示JSX表达式 反引号是模板字符串