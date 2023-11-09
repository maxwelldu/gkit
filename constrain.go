package gkit

// RealNumber in most cases, you should use this to express the meaning of numbers
type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Number interface {
	RealNumber | ~complex64 | ~complex128
}
