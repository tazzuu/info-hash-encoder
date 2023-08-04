package main

import  (
	"fmt"
	"encoding/hex"
	"net/url"
	"log"
	"os"
	// "strings"
)

func ConvertHash(input string) string {
	// https://github.com/anacrolix/torrent/blob/v1.52.4/types/infohash/infohash.go
	// 20-byte SHA1 hash ; also try hex.DecodeString(input)
	decodedArr := [20]byte{}
	_, err := hex.Decode(decodedArr[:], []byte(input)) // [20]uint8
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf(">>> %x\n", decodedArr[:])
	// fmt.Printf(">>> %#v\n", string(decodedArr[:]))

	// var encoded []string
	// for _, b := range decodedArr {
	// 	// fmt.Printf(">>> %%%02X\n", b)
	// 	encoded = append(encoded, fmt.Sprintf("%%%02X", b))
	// }
	// return strings.Join(encoded, "")

	// NOTE: special characters are not getting escaped correctly https://github.com/golang/go/issues/47379
	encoded := url.QueryEscape(string(decodedArr[:])) // also try PathEscape
	return encoded
}

func main() {
	inputStr := os.Args[1]
	fmt.Println(ConvertHash(inputStr))
}
