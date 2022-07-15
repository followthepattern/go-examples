## Built-in Types
Golang has similar built-in types to other languages, with couple of exceptions, here is a table of listed built-in types.
| Type | Short description | Default value | Short example
| - | - | - | - |
| `string` | string type | `""` | text := "hello"
| `bool` | represents the logical type | `false` | isVisible := true
| `int` | signed integer type that is at least 32 bits in size | `0` | index := 1
| `int8` |  set of all signed 8-bit integers | `0` | var index int8 = 127
| `int16` | set of all signed 16-bit integers | `0` | var index int16 = 32767
| `int32` | set of all signed 32-bit integers | `0` | var index int32 = 2147483647
| `int64` | set of all signed 64-bit integers. | `0` | var index int64 = 9223372036854775807
| `uint` | unsigned integer type that is at least 32 bits in size | `0` | index := 1
| `uint8` | set of all unsigned 8-bit integers | `0` | var index uint8 = 255
| `uint16` | set of all unsigned 16-bit integers | `0` | var index uint16 = 65535
| `uint32` | set of all unsigned 32-bit integers| `0` | var index uint32 = 4294967295
| `uint64` | set of all unsigned 64-bit integers | `0` | var index uint64 = 18446744073709551615
| `byte` | byte is an alias for uint8 and is equivalent to uint8 in all ways | `0` | var index byte = 255
| `rune` | rune is an alias for int32 and is equivalent to int32 in all ways | `0` | var index rune = 4294967295
| `float32` | set of all 32-bit floating-point numbers | `0` | var index float32 = -342.3439
| `float64` | set of all 64-bit floating-point numbers | `0` var index float64 = -342.3439
| `complex64` | set of all complex numbers with float32 real and imaginary parts | `(0+0i)` | index := complex(32.3, 18.9)
| `complex128` | set of all complex numbers with float64 real and imaginary parts | `(0+0i)` | index := complex(128.2323823, 1238.2829932)

In this folder, you can also find examples for how to create and use them.