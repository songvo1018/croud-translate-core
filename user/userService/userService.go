package userService

import "fmt"

type UserService struct {
}

func (us *UserService) health() {
	fmt.Print("health")
}
