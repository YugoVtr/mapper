# Package mapper 

`import "github.com/yugovtr/mapper"`

## Description

**Mapper** is intended to map values from one structure to another of different types using reflexion.

This mapping can be better implemented using [generic types](https://go.dev/blog/generics-proposal) that should be released in version 1.18

## Index 

```go
func Mapper(source interface{}, target interface{}) (err error)
```

## Examples


### Package files `mapper.go`
### **func Mapper**

> func Mapper(source interface{}, target interface{}) (err error)


Mapper traces fields in a structure source to fields of the same name in target structure.

- Source is a Struct.
- Target is a reference to Struct.

If the target has an attribute of type pointer, the pointer needs to store a struct to be mapped. Fields with same name but different types are ignored.

### Example Code:

```golang

target := struct {
    A string
    B int
    C bool
}{}

source := struct {
    A string
    B int
    C bool
}{
  A: "A",
  B: 1,
  C: true
}

err := Mapper(source, &target)

fmt.Printf("Target: %vnError: %v", target, err)
```

```
Output:
Target: {A 1 true}
Error: <nil>
```