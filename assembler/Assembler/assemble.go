package Assembler

import "strings"

/*
Assemble ARM instructions to machine code
*/
func Assemble(instrctions string) (code *string, errorLine int) {
	ins := strings.Split(strings.TrimSpace(instrctions), "\n")
	var b strings.Builder
	for i, v := range ins {
		if c := AssembleLine(v); c != nil {
			b.WriteString(*c)
		} else {
			return nil, i + 1
		}
	}
	c := b.String()
	return &c, 0
}
