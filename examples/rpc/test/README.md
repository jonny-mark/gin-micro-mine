


//額外import套件
import "google/api/annotations.proto";

//產生proto檔
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis greeter.proto --go_out=plugins=grpc:.
//產生gw檔
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis greeter.proto --grpc-gateway_out=logtostderr=true:.

参考资料：[gRPC Gateway]https://ithelp.ithome.com.tw/articles/10243864?sc=hot