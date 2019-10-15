package main

import (
	"fmt"

	"./Assembler"
)

func main() {
	ins := `
	lw r1,r0,0
	Lw r2,r0,4
	Addu r3, r1,r2
	Sw r3,r0,8
	halt
	`
	code, err := Assembler.Assemble(ins)
	if err != 0 {
		fmt.Println("Error line: ", err)
	} else {
		fmt.Print(*code)
	}
}
