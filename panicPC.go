package commontools

// Pc three lines in to one
func Pc(err error) {
	if err != nil {
		panic(err)
	}
}

// A Alias
var A = Pc
