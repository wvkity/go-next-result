package core

// DRKey Data Result Key类型
type DRKey interface {
	~string | ~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64
}
