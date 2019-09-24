package v3

import (
	"fmt"

	opt "github.com/gitchander/asn1/der/params_design/optional"
)

type Node struct{}

type Serializer interface {
	SerializeDER(*ParamsBuilder) (*Node, error)
}

type Deserializer interface {
	DeserializeDER(*Node, *ParamsBuilder) error
}

// Parameters
type Params struct {
	Name     opt.OptString
	Optional bool
	Explicit bool
	Tag      opt.OptInt
	Min      opt.OptInt
	Max      opt.OptInt
	Size     opt.OptInt
}

type ParamsBuilder struct {
	params Params
}

func PB() *ParamsBuilder {
	return new(ParamsBuilder)
}

func NewParamsBuilder() *ParamsBuilder {
	return new(ParamsBuilder)
}

func (pb *ParamsBuilder) Params() Params {
	if pb == nil {
		return Params{}
	}
	return pb.params
}

func Name(name string) *ParamsBuilder {
	return &ParamsBuilder{
		params: Params{
			Name: opt.String(name),
		},
	}
}

func Optional() *ParamsBuilder {
	return &ParamsBuilder{
		params: Params{
			Optional: true,
		},
	}
}

func Explicit() *ParamsBuilder {
	return &ParamsBuilder{
		params: Params{
			Explicit: true,
		},
	}
}

func Tag(tag int) *ParamsBuilder {
	return &ParamsBuilder{
		params: Params{
			Tag: opt.Int(tag),
		},
	}
}

//----------------------------------------------------------------

func (pb *ParamsBuilder) Name(name string) *ParamsBuilder {
	pb.params.Name = opt.String(name)
	return pb
}

func (pb *ParamsBuilder) Optional() *ParamsBuilder {
	pb.params.Optional = true
	return pb
}

func (pb *ParamsBuilder) Explicit() *ParamsBuilder {
	pb.params.Explicit = true
	return pb
}

func (pb *ParamsBuilder) Tag(tag int) *ParamsBuilder {
	pb.params.Tag = opt.Int(tag)
	return pb
}

// func (p *Params) SetTag(tag int) *Params {
// 	p.tag = opt.Int(tag)
// 	return p
// }

// func ParamsBuilder struct {
// 	p *Params
// }

// func PB()

func TestParamsBuilder(pb *ParamsBuilder) {
	fmt.Printf("%+v\n", pb.Params())
}

// func main() {
// 	//testParams(New().Tag(-2).Name("ABCDEF"))

// 	testParams(NewParams().SetTag(-2).SetName("ABCDEF"))
// 	testParams(
// 		&Params{
// 			Name: opt.String("ABCDEF"),
// 			tag:  opt.Int(-2),
// 		})
// }
