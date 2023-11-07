package concurrent

var (
	slice = make([]bool, 1)
)

func SetSlice() {
	slice[0] = true
}

func GetSlice() bool {
	return slice[0]
}
