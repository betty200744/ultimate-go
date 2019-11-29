package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const hashedPassword = "$2b$04$B1gNQ1zQ3GJsI3aAQtnRROcXJCpyFVPrTF0M9GRTPtC8fTqr/PNGq"
const password = "12345678"

func main() {
	hashedPassword1, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	fmt.Println(string(hashedPassword1))
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)))
	fmt.Println(bcrypt.Cost([]byte(hashedPassword)))
}
