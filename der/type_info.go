package der

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type fieldInfo struct {
	name     string
	optional bool
	explicit bool
	tag      *int
	min      *int
	max      *int
	size     *int
}

func fieldInfoToString(fi *fieldInfo) string {
	const separator = ','
	var b strings.Builder
	b.WriteByte('(')
	fmt.Fprintf(&b, "name:%q", fi.name)
	if fi.tag != nil {
		b.WriteByte(separator)
		fmt.Fprintf(&b, "tag:%d", *(fi.tag))
	}
	if fi.optional {
		b.WriteByte(separator)
		b.WriteString("optional")
	}
	if fi.explicit {
		b.WriteByte(separator)
		b.WriteString("explicit")
	}
	b.WriteByte(')')
	return b.String()
}

func parseFieldInfo(s string) (*fieldInfo, error) {

	const (
		prefTag  = "tag:"
		prefMin  = "min:"
		prefMax  = "max:"
		prefSize = "size:"
	)

	var fi fieldInfo

	for _, part := range strings.Split(s, ",") {

		switch {

		case part == "optional":
			fi.optional = true

		case part == "explicit":
			fi.explicit = true

		case strings.HasPrefix(part, prefTag):
			{
				x, err := parseFIParamInt(prefTag, part[len(prefTag):])
				if err != nil {
					return nil, err
				}
				fi.tag = newInt(x)
			}

		case strings.HasPrefix(part, prefMin):
			{
				x, err := parseFIParamInt(prefMin, part[len(prefMin):])
				if err != nil {
					return nil, err
				}
				fi.min = newInt(x)
			}

		case strings.HasPrefix(part, prefMax):
			{
				i, err := parseFIParamInt(prefMax, part[len(prefMax):])
				if err != nil {
					return nil, err
				}
				fi.max = newInt(i)
			}

		case strings.HasPrefix(part, prefSize):
			{
				i, err := parseFIParamInt(prefSize, part[len(prefSize):])
				if err != nil {
					return nil, err
				}
				fi.size = newInt(i)
			}
		}
	}

	return &fi, nil
}

// It parses <field info param> with type int.
func parseFIParamInt(paramName string, s string) (int, error) {
	x, err := parseInt(s)
	if err != nil {
		return 0, fmt.Errorf("param <%s> has invalid value %q", paramName, s)
	}
	return x, nil
}

type typeInfo struct {
	fields []*fieldInfo
}

type typeInfoMap struct {
	guard sync.RWMutex
	tis   map[reflect.Type]*typeInfo
}

func newTypeInfoMap() *typeInfoMap {
	return &typeInfoMap{
		tis: make(map[reflect.Type]*typeInfo),
	}
}

func (p *typeInfoMap) getTypeInfo(t reflect.Type) (*typeInfo, error) {

	p.guard.RLock()
	ti, ok := p.tis[t]
	p.guard.RUnlock()

	if ok {
		return ti, nil
	}

	ti = new(typeInfo)

	n := t.NumField()
	for i := 0; i < n; i++ {
		f := t.Field(i)
		tag := f.Tag.Get("der")
		fi, err := parseFieldInfo(tag)
		if err != nil {
			return nil, fmt.Errorf("type <%v>, field <%s>, %s", t, f.Name, err)
		}
		fi.name = f.Name
		ti.fields = append(ti.fields, fi)
	}

	p.guard.Lock()
	p.tis[t] = ti
	p.guard.Unlock()

	return ti, nil
}

var tinfoMap = newTypeInfoMap()
