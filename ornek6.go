package sorular

import "fmt"

type Telefon struct {
	Marka string
	Model string
	RAM   int
}

func (t Telefon) TelefonBilgileriYazdir() {
	fmt.Println("Marka:", t.Marka)
	fmt.Println("Model:", t.Model)
	fmt.Println("RAM:", t.RAM, "GB")
}

func Demo6() {
	fmt.Println(Telefon{"Apple", "14 Pro", 12})
}
