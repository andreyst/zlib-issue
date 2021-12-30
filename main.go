package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"

	otherzlib "github.com/4kills/go-zlib"
)

func main() {
	all()
	stream()

	// Alas it hangs :(
	// otherZlib()
}

func all() {
	fmt.Println("==== RUNNING DECOMPRESSION OF all.in")
	fmt.Println("")

	f, openErr := os.Open("./chunks/all.in")
	if openErr != nil {
		log.Fatalf("openErr=%v\n", openErr)
	}

	buf := make([]byte, 100000)
	n, readErr := f.Read(buf)
	if readErr != nil {
		log.Fatalf("readErr=%v\n", readErr)
	}
	fmt.Printf("Read bytes, n=%v\n", n)

	buffer := bytes.NewBuffer(buf)
	zlibReader, zlibErr := zlib.NewReader(buffer)
	if zlibErr != nil {
		log.Fatalf("zlibErr=%v\n", zlibErr)
	}

	out := new(bytes.Buffer)
	written, copyErr := io.Copy(out, zlibReader)
	if copyErr != nil {
		log.Printf("copyErr=%v\n", copyErr)
	}
	fmt.Printf("Written bytes, n=%v, out:\n%v\n", written, out.String())
	fmt.Println("")
}

func stream() {
	fmt.Println("==== RUNNING DECOMPRESSION OF SEPARATE CHUNKS")
	fmt.Println("")

	buffer := new(bytes.Buffer)

	f0, open0Err := os.Open("./chunks/0.in")
	if open0Err != nil {
		log.Fatalf("openErr=%v\n", open0Err)
	}

	buf0 := make([]byte, 100000)
	n0, read0Err := f0.Read(buf0)
	if read0Err != nil {
		log.Fatalf("read0Err=%v\n", read0Err)
	}
	fmt.Printf("Read bytes, n=%v\n", n0)
	buffer.Write(buf0)

	zlibReader, zlibErr := zlib.NewReader(buffer)
	if zlibErr != nil {
		log.Fatalf("zlibErr=%v\n", zlibErr)
	}

	out0 := new(bytes.Buffer)
	written0, copy0Err := io.Copy(out0, zlibReader)
	if copy0Err != nil {
		log.Printf("copy0Err=%v\n", copy0Err)
	}
	fmt.Printf("Written decompressed bytes, n0=%v, out:\n%v\n", written0, out0.String())

	f1, open1Err := os.Open("./chunks/1.in")
	if open1Err != nil {
		log.Fatalf("open1Err=%v\n", open1Err)
	}

	buf1 := make([]byte, 100000)
	n1, read1Err := f1.Read(buf1)
	if read1Err != nil {
		log.Fatalf("read1Err=%v\n", read1Err)
	}
	fmt.Printf("Read bytes, n1=%v\n", n1)
	buffer.Write(buf1)

	out1 := new(bytes.Buffer)
	written1, copy1Err := io.Copy(out1, zlibReader)
	if copy1Err != nil {
		log.Printf("copy1Err=%v\n", copy1Err)
	}
	fmt.Printf("Written decompressed bytes, n1=%v, out:\n%v\n", written1, out1.String())
	fmt.Println("")
}

func otherZlib() {
	fmt.Println("==== RUNNING DECOMPRESSION OF SEPARATE CHUNKS WITH OTHER ZLIB")
	fmt.Println("")

	// buffer := new(bytes.Buffer)

	f0, open0Err := os.Open("./chunks/0.in")
	if open0Err != nil {
		log.Fatalf("openErr=%v\n", open0Err)
	}

	buf0 := make([]byte, 100000)
	n0, read0Err := f0.Read(buf0)
	if read0Err != nil {
		log.Fatalf("read0Err=%v\n", read0Err)
	}
	fmt.Printf("Read bytes, n=%v\n", n0)
	// buffer.Write(buf0)

	zlibReader, zlibErr := otherzlib.NewReader(nil)
	if zlibErr != nil {
		log.Fatalf("zlibErr=%v\n", zlibErr)
	}

	nz0, decompressed0, readBuffer0err := zlibReader.ReadBuffer(buf0[:n0], nil)
	if readBuffer0err != nil {
		log.Printf("readBuffer0err=%v\n", readBuffer0err)
	}
	fmt.Printf("Decompressed bytes, nz0=%v, out:\n%v\n", nz0, string(decompressed0))

	// f1, open1Err := os.Open("./chunks/1.in")
	// if open1Err != nil {
	// 	log.Fatalf("open1Err=%v\n", open1Err)
	// }

	// buf1 := make([]byte, 100000)
	// n1, read1Err := f1.Read(buf1)
	// if read1Err != nil {
	// 	log.Fatalf("read1Err=%v\n", read1Err)
	// }
	// fmt.Printf("Read bytes, n1=%v\n", n1)
	// buffer.Write(buf1)

	// out1 := new(bytes.Buffer)
	// written1, copy1Err := io.Copy(out1, zlibReader)
	// if copy1Err != nil {
	// 	log.Printf("copy1Err=%v\n", copy1Err)
	// }
	// fmt.Printf("Written decompressed bytes, n1=%v, out:\n%v\n", written1, out1.String())

}
