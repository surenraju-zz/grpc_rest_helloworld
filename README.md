# grpc_rest_helloworld

This project aims to teach a ```gRPC``` service with ```HTTP+JSON``` interface. A small amount of configuration in your service to attach HTTP semantics is all that's needed to generate a reverse-proxy with this library.

Installation
The grpc-gateway requires a local installation of the Google protocol buffers compiler protoc v3.0.0 or above. Please install this via your local package manager or by downloading one of the releases from the official repository:

```https://github.com/protocolbuffers/protobuf/releases```

Then, ```go get -u``` as usual the following packages:
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```
This will place three binaries in your ```$GOBIN```

- ```protoc-gen-grpc-gateway```
- ```protoc-gen-grpc-swagger```
- ```protoc-gen-go```

Create ```third_party``` folder in the ```grpc_rent_helloworld``` project and copy content of ```include``` folder of ```protocolbuffers/protobuf``` binary to ```third_party``` folder in the project

Copy content of ```%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google``` folder to ```third_party/google``` folder in the project

Create ```protoc-gen-swagger/options``` folder in the ```third_party``` project folder 

```
mkdir -p third_party\protoc-gen-swagger\options
```

then copy ```annotations.proto``` and ```openapiv2.proto``` files from ```%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options``` folder to ```third_party\protoc-gen-swagger/options``` folder in the project


Create ```protoc-gen.cmd``` (protoc-gen.sh for MacOS/Linux) file in the third_party folder
```
protoc --proto_path=greetingservice --proto_path=third_party --go_out=plugins=grpc:greetingservice greetingservice.proto
protoc --proto_path=greetingservice --proto_path=third_party --grpc-gateway_out=logtostderr=true:greetingservice greetingservice.proto
protoc --proto_path=greetingservice --proto_path=third_party --swagger_out=logtostderr=true:greetingservice greetingservice.proto
```

From ```grpc_rest_helloworld``` folder run compilation
```
.\third_party\protoc-gen.cmd
```

for MacOS/Linux:
```
./third_party/protoc-gen.sh
```
It creates 

- ```greetingservice.pb.go``` - gRPC generated stub
- ```greetingservice.pb.gw.go``` - REST/HTTP generated stub
- ```greetingservice.swagger.json``` - generated Swagger documentation 

files inside ```greetingservice``` folder


To compile and run the server, assuming you are in the folder $GOPATH/src/github.com/surenraju/grpc_rest_helloworld, simply:
```
go run server/main.go
```

If you see
```
2019/05/14 17:28:25 starting HTTP/REST gateway...
2019/05/14 17:28:25 starting gRPC server...
```

It means server is started. 
