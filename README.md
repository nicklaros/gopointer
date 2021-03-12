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
result := gopointer.BoolPointer(true) // result type is *bool.
```

For float32 pointer
```
result := gopointer.Float32Pointer(99.99) // result type is *float32.
```

For float64 pointer
```
result := gopointer.Float64Pointer(99.99) // result type is *float64.
```

For int pointer
```
result := gopointer.IntPointer(99) // result type is *int.
```

For int8 pointer
```
result := gopointer.Int8Pointer(99) // result type is *int8.
```

For int16 pointer
```
result := gopointer.Int16Pointer(99) // result type is *int16.
```

For int32 pointer
```
result := gopointer.Int32Pointer(99) // result type is *int32.
```

For int64 pointer
```
result := gopointer.Int64Pointer(99) // result type is *int64.
```

For string pointer
```
result := gopointer.StringPointer("Hi, you can store me in a pointer.") // result type is *string.
```

For time pointer
```
import "time"

result := gopointer.TimePointer(time.Now()) // result type is *time
```

For uint pointer
```
result := gopointer.UintPointer(99) // result type is *uint.
```

For uint8 pointer
```
result := gopointer.Uint8Pointer(99) // result type is *uint8.
```

For uint16 pointer
```
result := gopointer.Uint16Pointer(99) // result type is *uint16.
```

For uint32 pointer
```
result := gopointer.Uint32Pointer(99) // result type is *uint32.
```

For uint64 pointer
```
result := gopointer.Uint64Pointer(99) // result type is *uint64.
```
