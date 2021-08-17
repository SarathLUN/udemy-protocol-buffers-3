# Protocol Buffers Advanced
## 1. Integer Type

- there exist many ways to represent an integer in protocol buffers
- int32, int64, uint32, uint64, sint32, sint64, fixed32, fixed64, sfixed32, sfixed64
- each type is basically constructed to handle:
  1. range of allowed value: 64 bits has more values than 32 bits
  2. whether negative value are allowed?
  3. size efficiency of serialization
- this is advanced and meant for performance and space optimization

---

## 2. Advanced Types: 
# `OneOf`

- you can use `oneof` to tell the protocol buffers that only one field can have a value:

```protobuf
syntax = "proto3";

message MyMessage {
    int32 id = 1;
    oneof my_one_of {
        string my_string = 2;
        bool my_bool = 3;
    }
}
```
- the `oneof` field cannot be repeated
- evolving schemas using `oneof` is completed (see documentation)
- on read, all fields will be null except the last one that was set at write

# `Maps`

- maps can be used to map scalars (except float/double) to value of any type
```protobuf
syntax = "proto3";

message MyMessage {
  int32 id = 1;
  map<string, Result> results = 2;
}

message Result { 
  string result = 1; 
} 
```
- map field cannot be repeated
- there's no ordering in map (it's a key=>value store)

# `Timestamp` (Well Known Types)
- protocol buffers contain a set of well known types - e.g. advanced types known by all programming languages
- the list is [here](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf)
- one of the types is `timestamp` - fields are _second_ and _nanosecond_ (UTC)
- don't forget to import to be able to use it
```protobuf
syntax = "proto3";
import "google/protobuf/timestamp.proto";
message MyMessage{
  google.protobuf.Timestamp my_field = 1;
}
```
# `Duration`
- duration is another well known type
- it represents the time span between two timestamps
- it contains, just like Timestamp, seconds and nanoseconds
```protobuf
syntax = "proto3";

import "google/protobuf/timestamp.proto";
import "google/protobuf/duration.proto";

message MyMessage {
  google.protobuf.Timestamp msg_date = 1;
  google.protobuf.Duration validaty = 2;
}
```

---

# Protocol Buffers Options
- options allow to alter the behavior of the protoc compiler when generating code for specific languages
- there are few bundled options (read the docs), here are a few:
```protobuf
syntax = "proto3";
option csharp_namespace = "Google.Protobuf.WellKnownTypes";
option cc_enable_arenas = true;
option go_package = "github.com/golang/protobuf/ptypes/duration";
option java_package = "com.google.protobuf";
option java_outer_classname = "DurationProto";
option java_multiple_files = true;
option objc_class_prefix = "GPB";
```

---

# Naming Convention From the docs
- refer to https://developers.google.com/protocol-buffers/docs/style
- use CamelCase for message name
- use underscore_separated_names for field name
```protobuf
syntax = "proto3";
message MyLongMessage{
  string my_long_field = 1;
}
```

---

# Introduction to Protocol Buffer Services
- protocol buffers can define services on top of messages
- a service is a set of endpoints your application can be assessable from
```protobuf
syntax = "proto3";

message SearchRequest { 
  int32 person_id = 1; 
}

message SearchResponse {
  int32 person_id = 1;
  string person_name = 2;
}

service SearchService { 
  rpc Search(SearchRequest) returns (SearchResponse); 
}
```
- service need to be interpreted by a framework to generate associated code
- the main framework is gRPC, but you may find others in the web
- the advantage of services and RPC is that you can call Server API from any client seamlessly
- gRPC for example is used by Google, Netflix, CoreOS (etcd), Google Cloud API, and is gaining popularity fast
