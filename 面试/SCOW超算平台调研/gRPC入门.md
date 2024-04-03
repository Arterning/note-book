gRPC 是 Google 开源的一个高性能的 RPC(Remote Procedure Call) 框架，它具有如下的优点：

- 提供高效的进程间通信。gRPC 没有使用 XML 或者 JSON 这种文本格式，而是采用了基于 protocol buffers 的二进制协议；同时，gRPC 采用了 HTTP/2 做为通信协议，从而能够快速的处理进程间通信。
- 简单且良好的服务接口和模式。gRPC 为程序开发提供了一种契约优先的方式，必须首先定义服务接口，才能处理实现细节。
- 支持多语言。gRPC 是语言中立的，我们可以选择任意一种编程语言，都能够与 gRPC 客户端或者服务端进行交互。
- 成熟并且已被广泛使用。通过在 Google 的大量实战测试，gRPC 已经发展成熟。

下面通过一个简单的 demo 来初步了解 gRPC 的使用。

我们构建一个商品服务，命名为 ProductInfo，客户端和服务端的交互模式如下：


![[Pasted image 20240402090805.png]]


首先我们需要定义 protobuf：

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


我们定义了一个 ProductInfo 服务，其中有两个方法，分别是添加商品和获取商品，然后在 proto 文件所在的目录下执行命令 `protoc --go_out=plugins=grpc:../product ProductInfo.proto`。

如果没有安装 protoc，执行命令 `go get -u github.com/golang/protobuf/protoc-gen-go` 进行安装。虽然 gRPC 支持多种语言，但是为了统一，我文章中的代码都使用 Go。

执行完之后，在 proto 文件同级目录下会出现一个 `ProductInfo.pb.go` 文件：

这就是生成的服务端骨架，在其中会自动生成一个注册服务的方法，我们后面会用到

```go

func RegisterProductInfoServer(s *grpc.Server, srv ProductInfoServer) {
	s.RegisterService(&_ProductInfo_serviceDesc, srv)
}
...
```



然后在 product 文件夹下新建一个 server 文件夹，然后新建一个 `main.go` 文件，首先在文件中实现 ProductInfo 服务的两个方法的业务逻辑：

```go
package main
​
import (
   "context"
   "github.com/gofrs/uuid"
   "google.golang.org/grpc/codes"
   "google.golang.org/grpc/status"
   "grpc-demo/product"
)
​
type server struct {
   productMap map[string]*product.Product
}
​
//添加商品
func (s *server) AddProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
   resp = &product.ProductId{}
   out, err := uuid.NewV4()
   if err != nil {
      return resp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
   }
   
   req.Id = out.String()
   if s.productMap == nil {
      s.productMap = make(map[string]*product.Product) 
   }
   
   s.productMap[req.Id] = req
   resp.Value = req.Id
   return
}
​
//获取商品
func (s *server) GetProduct(ctx context.Context, req *product.ProductId) (resp *product.Product, err error) {
   if s.productMap == nil {
      s.productMap = make(map[string]*product.Product)
   }
​
   resp = s.productMap[req.Value]
   return
}
```

然后继续在 main.go 文件中添加一个 main 方法，建立一个 gRPC 服务器：

```go
func main() {
   listener, err := net.Listen("tcp", port)
   if err != nil {
      log.Println("net listen err ", err)
      return
   }
​
   s := grpc.NewServer()
   product.RegisterProductInfoServer(s, &server{})
   log.Println("start gRPC listen on port " + port)
   if err := s.Serve(listener); err != nil {
      log.Println("failed to serve...", err)
      return
   }
}
```




服务端的逻辑就到这里了，接下来再写一下客户端的逻辑，建立一个 client 文件夹，然后新建一个 main.go 文件，内容如下：

```go
package main
​
import (
   "context"
   "google.golang.org/grpc"
   "grpc-demo/product"
   "log"
)
​
const (
   address = "localhost:50051"
)
​
func main() {
   conn, err := grpc.Dial(address, grpc.WithInsecure())
   if err != nil {
      log.Println("did not connect.", err)
      return
   }
   defer conn.Close()
​
   client := product.NewProductInfoClient(conn)
   ctx := context.Background()
​
   id := AddProduct(ctx, client)
   GetProduct(ctx, client, id)
}
​
// 添加一个测试的商品
func AddProduct(ctx context.Context, client product.ProductInfoClient) (id string) {
   aMac := &product.Product{Name: "Mac Book Pro 2019", Description: "From Apple Inc."}
   productId, err := client.AddProduct(ctx, aMac)
   if err != nil {
      log.Println("add product fail.", err)
      return
   }
   log.Println("add product success, id = ", productId.Value)
   return productId.Value
}
​
// 获取一个商品
func GetProduct(ctx context.Context, client product.ProductInfoClient, id string) {
   p, err := client.GetProduct(ctx, &product.ProductId{Value: id})
   if err != nil {
      log.Println("get product err.", err)
      return
   }
   log.Printf("get prodcut success : %+v\n", p)
}
```





代码地址：
https://github.com/roseduan/grpc-demo/blob/main/product/ProductInfo.pb.go


