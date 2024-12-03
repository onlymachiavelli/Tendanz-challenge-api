package utils

import "math/rand"

func GenerateCode() string {


	code := ""
	for i := 0; i < 6; i++ {
		code += string(rune(48 + rand.Intn(10)))
	}
	return code


}