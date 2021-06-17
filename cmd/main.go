package main

import "github.com/nathangreene3/slice/strs"

func main() {
	a, b := []string{"a", "b", "c"}, []string{"a", "b", "c"}
	strs.Equal(a, b)
	strs.Equal0(a, b)
}
