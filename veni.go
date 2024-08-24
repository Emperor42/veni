package veni

import "fmt"

type VeniContext struct {
	systemTest string
}

func (v *VeniContext) ProcessHeader() {
	fmt.Println("temp")
}
