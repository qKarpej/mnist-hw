package main

import (
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "t10k-images-idx3-ubyte.gz"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// var tmp[i] - magicNumber, nImages, nRows, nCols
	tmp := make([]uint32, 4)
	for i := range tmp {
		err = binary.Read(r, binary.BigEndian, &tmp[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	image := make([]byte, tmp[2]*tmp[3]*10)
	err = binary.Read(r, binary.BigEndian, &image)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DebugPrint(uint32(tmp[3]), image))

}
func DebugPrint(columns uint32, pixels []byte) string {
	var b []byte
	for i, p := range pixels {
		if uint32(i)%columns == 0 {
			b = append(b, '\n')
		}
		if p == 0 {
			b = append(b, ' ', ' ')
		} else if p < 128 {
			b = append(b, '.', '.')
		} else {
			b = append(b, '#', '#')
		}
	}
	return string(b)
}
