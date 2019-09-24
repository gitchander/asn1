package optional

type OptString struct {
	Optional
	Value string
}

func (v OptString) String() string {
	if v.Present {
		return v.Value
	}
	return absentValue
}

func String(v string) OptString {
	return OptString{
		Optional: Optional{
			Present: true,
		},
		Value: v,
	}
}
