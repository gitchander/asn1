package v4

import (
	opt "github.com/gitchander/asn1/der/params_design/optional"
)

type Params struct {
	name     opt.OptString
	optional bool
	explicit bool
	tag      opt.OptInt
	min      opt.OptInt
	max      opt.OptInt
	size     opt.OptInt
}

func (p *Params) Name() (name string, ok bool) {
	if p == nil {
		return "", false
	}
	if !(p.name.Present) {
		return "", false
	}
	return p.name.Value, true
}

func (p *Params) Optional() bool {
	if p == nil {
		return false
	}
	return p.optional
}

func (p *Params) Explicit() bool {
	if p == nil {
		return false
	}
	return p.explicit
}

func (p *Params) Tag() (tag int, ok bool) {
	if p == nil {
		return 0, false
	}
	if !(p.tag.Present) {
		return 0, false
	}
	return p.tag.Value, true
}
