- file: simple.proto
- we need to add `option go_package` to define output
- syntax `option go_package = "<output_path>;<output_package_name>";`
- this example will output the generated file into `./example/simple/` and the package name will be `simple`

```protobuf
syntax = "proto3";

option go_package = "./example/simple;simple"; //ADDED THIS LINE HERE FOR GOLANG GRPC

package example.simple;

message SimpleMessage {
  int32 id = 1;
  bool is_simple = 2;
  string name = 3;
  repeated int32 sample_list = 4;
}
```
- generate via below command
```shell
> protoc -I . --go_out=. simple.proto
> tree example
```
```shell
example
└── simple
    └── simple.pb.go

1 directory, 1 file
```