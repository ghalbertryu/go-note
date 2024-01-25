package concurrent

var (
	slice = make([]bool, 1)
)

func SetSlice(boolValue bool) {
	slice[0] = boolValue
}

func GetSlice() bool {
	return slice[0]
}

func AppendSlice(boolValue bool) {
	slice = append(slice, boolValue)
}
