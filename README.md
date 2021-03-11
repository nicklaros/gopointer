# gopointer

Minimalistic pointer utility for go.

# Why Gopointer?

Gopointer makes golang easier by taking the hassle out of defining intermediary variable for
composite literal just so you can take it's address.

![failed to take address](https://i.ibb.co/ZcrQLCz/why-gopointer.png)

# How to use

Import package as follows
```
import "github.com/nicklaros/gopointer"
```

For boolean pointer

```
value := true
result := gopointer.BoolPointer(value) // result type is *bool.
```

For float32 pointer
```
value := float32(99.99)
result := gopointer.Float32Pointer(value) // result type is *float32.
```

For float64 pointer
```
value := float64(99.99)
result := gopointer.Float64Pointer(value) // result type is *float64.
```

For int pointer
```
value := int(99)
result := gopointer.IntPointer(value) // result type is *int.
```

For int8 pointer
```
value := int8(99)
result := gopointer.Int8Pointer(value) // result type is *int8.
```

For int16 pointer
```
value := int16(99)
result := gopointer.Int16Pointer(value) // result type is *int16.
```

For int32 pointer
```
value := int32(99)
result := gopointer.Int32Pointer(value) // result type is *int32.
```

For int64 pointer
```
value := int64(99)
result := gopointer.Int64Pointer(value) // result type is *int64.
```

For string pointer
```
value := "Hi, you can store me in a pointer."
result := gopointer.StringPointer(value) // result type is *string.
```

For time pointer
```
import "time"

value := time.Now()
result := gopointer.TimePointer(value) // result type is *time
```

For uint pointer
```
value := uint(99)
result := gopointer.UintPointer(value) // result type is *uint.
```

For uint8 pointer
```
value := uint8(99)
result := gopointer.Uint8Pointer(value) // result type is *uint8.
```

For uint16 pointer
```
value := uint16(99)
result := gopointer.Uint16Pointer(value) // result type is *uint16.
```

For uint32 pointer
```
value := uint32(99)
result := gopointer.Uint32Pointer(value) // result type is *uint32.
```

For uint64 pointer
```
value := uint64(99)
result := gopointer.Uint64Pointer(value) // result type is *uint64.
```
