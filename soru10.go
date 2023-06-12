package sorular

import "fmt"

type Motor struct {
	Marka string
	Model string
	Tarih int
}

func (m Motor) MotorunBilgileriniYazdir() {
	fmt.Println("Markası:", m.Marka)
	fmt.Println("Model:", m.Model)
	fmt.Println("Tarih:", m.Tarih)
}

func Demo10() {
	fmt.Println(Motor{"BMW", "R 1250 GS", 2021})
}
