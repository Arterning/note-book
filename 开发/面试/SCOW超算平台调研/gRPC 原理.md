实际上就是定义一个统一的格式，和语言无关


然后把他翻译成语言相关的类型


比如

```protobuf
syntax = "proto3";
package product;
​
service ProductInfo {
  //添加商品
  rpc addProduct(Product) returns (ProductId);
  //获取商品
  rpc getProduct(ProductId) returns (Product);
}
​
message Product {
  string id = 1;
  string name = 2;
  string description = 3;
}
​
message ProductId {
  string value = 1;
}
```

在go中生成的结果，也就是存根，包括Product类型的定义，ProductInfo类型的getter setter方法