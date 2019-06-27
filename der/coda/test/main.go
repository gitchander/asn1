package main

import (
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gitchander/asn1/der/coda"
)

// http://luca.ntop.org/Teaching/Appunti/asn1.html

func main() {

	// const base = 128
	// fmt.Printf("%X\n", 255%base)
	// fmt.Printf("%X\n", 1<<7)
	// return

	asn1Test()
	//testHeader()
	testDecode()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testHeader() {

	h := &coda.Header{
		Tag:        1,
		Class:      coda.CLASS_CONTEXT_SPECIFIC,
		IsCompound: false,
	}
	data, err := coda.EncodeHeader(nil, h)
	checkError(err)
	fmt.Printf("%X\n", data)

	var m coda.Header

	_, err = coda.DecodeHeader(data, &m)
	checkError(err)

	fmt.Printf("%+v\n", m)
}

func testDecode() {
	s := "9F818080808000"

	data, err := hex.DecodeString(s)
	checkError(err)

	var h coda.Header
	_, err = coda.DecodeHeader(data, &h)
	checkError(err)

	fmt.Printf("%+v\n", h)
}

func asn1Test() {

	a := make([]byte, 15)

	//data, err := asn1.Marshal(Persone{Age: 0})

	params := "tag:1"
	data, err := asn1.MarshalWithParams(a, params)
	checkError(err)

	fmt.Printf("%X\n", data)

	var b []byte
	_, err = asn1.UnmarshalWithParams(data, &b, params)
	checkError(err)

	//fmt.Println(b)
}

type Persone struct {
	//Age int `asn1:"tag:5,explicit"`
	Age int `asn1:"tag:128"`
}
