package main

import (
	"fmt"

	"github.com/gitchander/asn1/der/params_design/v2"
	"github.com/gitchander/asn1/der/params_design/v3"
)

func main() {
	testV2()
	testV3()
}

func testV2() {
	v2.TestParams(v2.Tag(-9), v2.Name("hello"), v2.Optional{}, v2.Explicit{})
}

func testV3() {

	pb := v3.Tag(-6).Explicit()
	fmt.Println(pb.Params())

	v3.TestParamsBuilder(nil)
	v3.TestParamsBuilder(v3.Name("122334").Tag(-90))
	v3.TestParamsBuilder(v3.Tag(7))
	v3.TestParamsBuilder(v3.PB().Tag(12).Name("121212").Explicit().Tag(-3))
}
