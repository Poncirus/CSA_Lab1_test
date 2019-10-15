package main

import (
	"fmt"

	"./assembler"
)

func main() {
	ins := `
	lw r1,r0,0
	Lw r2,r0,4
	Addu r3, r1,r2
	ADDU R4, R3, R2
	ADDU R5, R3, R4
	LW R6, R0, 8
	ADDU R7, R6, R5
	SW R7, R0, 12
	LW R8, R0, 4
	Sw r8,r0,16
	halt
	`
	code, err := assembler.Assemble(ins)
	if err != 0 {
		fmt.Println("Error line: ", err)
	} else {
		fmt.Print(*code)
	}
}
