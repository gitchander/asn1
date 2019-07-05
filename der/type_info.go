package der

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type typeInfo struct {
	fields []fieldInfo
}

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
	var sb strings.Builder
	const separator = ','
	sb.WriteByte('(')
	fmt.Fprintf(&sb, "name:%q", fi.name)
	if fi.tag != nil {
		sb.WriteByte(separator)
		fmt.Fprintf(&sb, "tag:%d", *(fi.tag))
	}
	if fi.optional {
		sb.WriteByte(separator)
		sb.WriteString("optional")
	}
	if fi.explicit {
		sb.WriteByte(separator)
		sb.WriteString("explicit")
	}
	sb.WriteByte(')')
	return sb.String()
}

var (
	tinfoMap   = make(map[reflect.Type]*typeInfo)
	tinfoMutex sync.RWMutex
)

func getTypeInfo(t reflect.Type) (*typeInfo, error) {

	tinfoMutex.RLock()
	tinfo, ok := tinfoMap[t]
	tinfoMutex.RUnlock()

	if ok {
		return tinfo, nil
	}

	tinfo = new(typeInfo)

	n := t.NumField()
	for i := 0; i < n; i++ {
		f := t.Field(i)
		tag := f.Tag.Get("der")
		params := parseFieldInfo(tag)
		params.name = f.Name
		tinfo.fields = append(tinfo.fields, params)
	}

	tinfoMutex.Lock()
	tinfoMap[t] = tinfo
	tinfoMutex.Unlock()

	return tinfo, nil
}

func parseFieldInfo(s string) (fp fieldInfo) {

	const (
		prefTag  = "tag:"
		prefMin  = "min:"
		prefMax  = "max:"
		prefSize = "size:"
	)

	for _, part := range strings.Split(s, ",") {

		switch {

		case part == "optional":
			fp.optional = true

		case part == "explicit":
			fp.explicit = true

		case strings.HasPrefix(part, prefTag):
			{
				i, err := strconv.ParseInt(part[len(prefTag):], 10, 32)
				if err == nil {
					fp.tag = new(int)
					*fp.tag = int(i)
				}
			}

		case strings.HasPrefix(part, prefMin):
			{
				i, err := strconv.ParseInt(part[len(prefMin):], 10, 32)
				if err == nil {
					fp.min = new(int)
					*fp.min = int(i)
				}
			}

		case strings.HasPrefix(part, prefMax):
			{
				i, err := strconv.ParseInt(part[len(prefMax):], 10, 32)
				if err == nil {
					fp.max = new(int)
					*fp.max = int(i)
				}
			}

		case strings.HasPrefix(part, prefSize):
			{
				i, err := strconv.ParseInt(part[len(prefSize):], 10, 32)
				if err == nil {
					fp.size = new(int)
					*fp.size = int(i)
				}
			}
		}
	}

	return
}
