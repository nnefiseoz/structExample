package sorular

import "fmt"

type Hayvan struct {
	Ad  string
	Tur string
}

func (h Hayvan) HayvanBilgileriYazdir() {
	fmt.Println("Adı:", h.Ad)
	fmt.Println("Türü:", h.Tur)
}
