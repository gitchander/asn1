package der

type Parameter interface {
	isParameter()
}

type Name string
type Tag int
type Optional struct{}
type Explicit struct{}

func (Name) isParameter()     {}
func (Tag) isParameter()      {}
func (Optional) isParameter() {}
func (Explicit) isParameter() {}

func GetTagByParams(params []Parameter) (tag int, ok bool) {
	for _, p := range params {
		if tag, ok := p.(Tag); ok {
			return int(tag), true
		}
	}
	return 0, false
}
