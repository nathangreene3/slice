# Slice

```go
go get github.com/nathangreene3/slice
```

A library for manipulating slices.

## Subpackages

The following subpackages contain the same functionality specific to each type of slice they support.

* bts: extends `[]byte`
* flts: extends `[]float64`
* ints: extends `[]int`
* strs: extends `[]string`
* uints: extends `[]uint`

## Function purity

If a function maps a type to the same type, the given value will not be mutated and the returned value will be a new value.
