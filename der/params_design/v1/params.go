package v1

import (
	"log"

	opt "github.com/gitchander/asn1/der/params_design/optional"
)

type Node struct {
	b byte
}

type Serializer interface {
	SerializeDER(options ...Option) (*Node, error)
}

type Deserializer interface {
	DeserializeDER(n *Node, options ...Option) error
}

type Option func(*Options) error

type Options struct {
	Name     opt.OptString
	Optional bool
	Explicit bool
	Tag      opt.OptInt
	Min      opt.OptInt
	Max      opt.OptInt
	Size     opt.OptInt
}

var DefaultOptions = Options{
	//Tag: -1,
	//Tag: optInt{},
}

func Name(name string) Option {
	return func(o *Options) error {
		o.Name = opt.String(name)
		return nil
	}
}

func Tag(tag int) Option {
	return func(o *Options) error {
		o.Tag = opt.Int(tag)
		return nil
	}
}

func makeOptions(options ...Option) (*Options, error) {
	var opts Options
	for _, opt := range options {
		if err := opt(&opts); err != nil {
			return nil, err
		}
	}
	return &opts, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
