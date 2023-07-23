package translationTextServcice

import "fmt"

type TranslationTextService struct {
}

func (us *TranslationTextService) health() {
	fmt.Print("health")
}
