package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const hashedPassword = "$2b$04$B1gNQ1zQ3GJsI3aAQtnRROcXJCpyFVPrTF0M9GRTPtC8fTqr/PNGq"
const password = "whaletao8986"

func main() {
	hashedPassword1, _ := bcrypt.GenerateFromPassword([]byte(password), 1)
	fmt.Println("whaletao8986: ", string(hashedPassword1))
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)))
	fmt.Println(bcrypt.Cost([]byte(hashedPassword)))
}
