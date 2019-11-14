package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/crc32"
)

// сюда писать код
func worker(index int, data string) {
	// Md5 from data
	Md5 := DataSignerMd5(data)
	Crc32 := DataSignerCrc32(data)
	//Md5 := md5.Sum([]byte(data))
	// CRC32 from MD5
	tempHash := md5.Sum([]byte(data))
	Crc32FromMd5 := crc32.ChecksumIEEE([]byte(hex.EncodeToString(tempHash[:])))
	// CRC32 from data
	//Crc32 := crc32.ChecksumIEEE([]byte(data))

	//Crc32Md5 := crc32.ChecksumIEEE([]byte(hex.EncodeToString(Md5)))

	fmt.Printf("%v SingleHash data %v\n", index, data)
	fmt.Printf("%v SingleHash md5(data) %v\n", index, Md5)
	fmt.Printf("%v SingleHash crc32(md5(data)) %v\n", index, Crc32FromMd5)
	fmt.Printf("%v SingleHash crc32(data) %v\n", index, Crc32)
	fmt.Printf("%v SingleHash result %v~%v\n", index, Crc32, Crc32FromMd5)
}

func main() {
	data := []string{"0", "1"}
	for i, v := range data {
		worker(i, v)
	}
	_, _ = fmt.Scanln()
}
