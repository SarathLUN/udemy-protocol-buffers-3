# Protobuf Data Evolution
>Evolve your .proto files in a safe way.

1. Why do we need to update the protocols?
- when you first declare a message in your protocol, you have a defined sets of requirements
- but as time go on, your business will evolve, and you have a difference set of requirements
- some fields may change, some fields may be added and other removed.
- now, also with time, many applications may read your messages using Protocol Buffers and you may not have the time to upgrade them
- for example, today we're asking for the First_Name and Last_Name of our customer, and that's our v1 schema, but tomorrow we may ask for their phone number. that would be our v2 of our schema
- so we need to be able to evolve the source data without breaking the other applications reading it
- thankfully, Protocol Buffers helps us tremendously with that as we'll see in th this section
- scenario 1: write data with new `.proto` -> read data with old `.proto`, this called **Forward** compatible change
- scenario 2: write data with old `.proto` -> read data with new `.proto`, this called **Backward** compatible change
- while protobuf provided us out of the box the both of forward and backward compatible which is called **FULL** compatible change

---

# Updating Protocol Rules (from the documentation)
1. Don't change the numeric tags for any existing fields.
2. You can add new fields, and old code will just ignore them.
3. likewise, if the old / new code reads unknown data, the default will take place
4. fields can be removed, as long as the tag number is not used again in your updated message type. you may want to rename the field instead, perhaps adding the prefix "OBSOLETE_", or make the tag reserved, so that future users of your `.proto` can't accidentally reuse the number.
5. for data type changes (int32 to int64 for example, please refer to the documentation)

---

# Adding fields
- let add a field to our schema (new tag number)
```protobuf
syntax = "proto3";
message MyMessage{
    int32 id = 1;
}
```
change to:
```protobuf
syntax = "proto3";
message MyMessage{
    int32 id = 1;
    string first_name = 2;
}
```
- if that field sends to old code, the old code will not know what that tag number corresponds to, and the field will be ignored or dropped.
- oppositely, if we read the old data with new code, the new field will not be found, and the default value will be assumed (empty string)
- ## The default values should always be interpreted with care.

---

# Renaming field
- let's rename a field in our schema
from:
```protobuf
syntax = "proto3";
message MyMessage{
    int32 id = 1;
    string first_name = 2;
}
```
to:
```protobuf
syntax = "proto3";
message MyMessage{
    int32 id = 1;
    string person_first_name = 2;
}
```
- in this case, nothing changes! field name can be change freely
- ## Only the tag number is important for protobuf

---

# Removing Fields
- let's remove a field in our schema
from:
```protobuf
syntax = "proto3";
message MyMessage{
    int32 id = 1;
    string first_name = 2;
}
```
to:
```protobuf
syntax = "proto3";
message MyMessage{
    int32 id = 1;
}
```
- if old code doesn't find the field anymore, the default value will be used
- oppositely, if we read old data with the new code, the deleted field will just drop
- ## Default values should always be interpreted with care
- ## when removing a field, you should always reserve the tage and the name
from:
```protobuf
syntax = "proto3";
message MyMessage{
    int32 id = 1;
    string first_name = 2;
}
```
to:
```protobuf
syntax = "proto3";
message MyMessage{
    reserved 2;
    reserved first_name;
    int32 id = 1;
}
```
- this prevents the tag to be re-used and this prevents the name to be re-used
- this is necessary to prevent conflicts in the codebase
- the alternative is that instead of removing a field, you rename it to `OBSOLETE_field_name`.
- the downside is that you may have to populate that field while your client get upgraded to use the newer field that replaces it (which has a new tag)

---

# Reserved keywords
- you can reserve TAGS and FIELD NAMES
- you can't mix TAGS and FIELD NAMES in the same `reserved` statement
```protobuf
syntax = "proto3";
message Foo{
    reserved 2,15, 9 to 11;
    reserved foo, bar;
}
```
- we reserve TAGS to prevent new fields from re-using tags (which would break old code at runtime)
- we reserve FIELD NAMES to prevent code bugs
- ## Do not <u>_ever_</u> remove any reserved tags.

---

# Beware of Default!
- defaults are great, but they are tricky to deal with
- default allow us to easily evolve protobuf files without breaking any existing or new codes
- they also ensure we know that a field will always have a non-null values
- but they are dangerous, because...
- you cannot differentiate from a missing field or if a value equal to the default was set
- now, what can we do about it?
  - make sure the default value doesn't have meaning for your business
  - deal with default value in your code if needed (with if statements)

---

# Evolving Enumerations
- Enumerations can evolve:
  - you can add
  - you can remove
  - you can reserve
- if the code doesn't know what the received Enum value corresponds to, the default value will be used
- therefore, I recommend you make the first value `UNKNOWN = 0`