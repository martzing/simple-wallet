package auth

import "fmt"

func register(data *RegisterData) {
	fmt.Printf("%v", data)

	fmt.Printf(data.Username)
	fmt.Printf(data.Password)
}
