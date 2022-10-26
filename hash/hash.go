package gutil

import "github.com/spaolacci/murmur3"

func Hash64Byte(data []byte) uint64 {
	return murmur3.Sum64(data)
}

func Hash64(data string) uint64 {
	return murmur3.Sum64([]byte(data))
}

func Hash32Byte(data []byte) uint32 {
	return murmur3.Sum32(data)
}

func Hash32(data string) uint32 {
	return murmur3.Sum32([]byte(data))
}
