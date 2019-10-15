package Assembler

import (
	"fmt"
	"strconv"
	"strings"
)

/*
AssembleLine assemble one line ARM instruction to machine code
*/
func AssembleLine(instrction string) *string {
	// to upper
	ins := strings.ToUpper(instrction)

	// remove space and \n
	ins = strings.TrimSpace(ins)

	// HALT
	if ins == "HALT" {
		code := "11111111\n11111111\n11111111\n11111111\n"
		return &code
	}

	// get instruction
	firstSpace := strings.Index(ins, " ")
	i := ins[:firstSpace]

	// split parameters
	ins = ins[firstSpace:]
	paras, n := splitParameter(ins)

	// code
	c := ""

	// add opcode
	switch i {
	case "ADDU", "SUBU":
		c += iToA(0, 6)
	case "BEQ":
		c += iToA(0x4, 6)
	case "LW":
		c += iToA(0x23, 6)
	case "SW":
		c += iToA(0x2B, 6)
	default:
		return nil
	}

	// add other code
	switch i {
	case "ADDU", "SUBU":
		if n != 3 {
			return nil
		}

		if v, err := decodeR(paras[1], 5); err {
			c += *v
		} else {
			return nil
		}
		if v, err := decodeR(paras[2], 5); err {
			c += *v
		} else {
			return nil
		}
		if v, err := decodeR(paras[0], 5); err {
			c += *v
		} else {
			return nil
		}
	case "BEQ", "LW", "SW":
		if n != 3 {
			return nil
		}

		if v, err := decodeR(paras[1], 5); err {
			c += *v
		} else {
			return nil
		}
		if v, err := decodeR(paras[0], 5); err {
			c += *v
		} else {
			return nil
		}
		if i, err := strconv.Atoi(paras[2]); err == nil {
			c += iToA(i, 16)
		} else {
			return nil
		}
	}

	// add funct
	switch i {
	case "ADDU":
		c += iToA(0x21, 11)
	case "SUBU":
		c += iToA(0x23, 11)
	}

	var b strings.Builder
	for i := 0; i < 4; i++ {
		b.WriteString(c[i*8 : (i+1)*8])
		b.WriteByte('\n')
	}
	c = b.String()

	return &c
}

/*
splitParameter split instrution to parameters
*/
func splitParameter(ins string) (paras []string, n int) {
	// remove space and \n
	ins = strings.TrimSpace(ins)

	p := strings.Split(ins, ",")
	for _, v := range p {
		v = strings.TrimSpace(v)
	}

	return p, len(p)
}

/*
iToa returns the string representation of i, the length of the string is len
*/
func iToA(i int, length int) string {
	s := strconv.FormatInt(int64(i), 2)
	if len(s) > length {
		return s[:length]
	}

	var b strings.Builder
	for b.Len() < length-len(s) {
		fmt.Fprintf(&b, "0")
	}
	b.WriteString(s)

	return b.String()
}

func decodeR(c string, length int) (*string, bool) {
	c = strings.TrimSpace(c)
	if strings.HasPrefix(c, "R") {
		s := strings.TrimLeft(c, "R")
		if i, err := strconv.Atoi(s); err == nil {
			code := iToA(i, length)
			return &code, true
		}
		return nil, false
	}
	return nil, false
}