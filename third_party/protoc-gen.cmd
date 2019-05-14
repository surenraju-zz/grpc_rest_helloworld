protoc --proto_path=greetingservice --proto_path=third_party --go_out=plugins=grpc:greetingservice greetingservice.proto
protoc --proto_path=greetingservice --proto_path=third_party --grpc-gateway_out=logtostderr=true:greetingservice greetingservice.proto
protoc --proto_path=greetingservice --proto_path=third_party --swagger_out=logtostderr=true:greetingservice greetingservice.proto
