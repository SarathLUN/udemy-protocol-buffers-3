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
---
# Go working on data model
- file: main.go

```go
package main

import (
	"github.com/SarathLUN/udemy-protocol-buffers-3/my-code/06-golang/01/example/simple"
	"log"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simple.SimpleMessage{
		Id:         123,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 23, 45, 6},
	}
	log.Println(sm)
}

```
- initialize Go module and run
```shell
> go mod init
> go mod tidy
> go run main.go
```
- output:
```shell
{{{} [] [] <nil>} 0 [] 123 true My Simple Message [1 23 45 6]}
```