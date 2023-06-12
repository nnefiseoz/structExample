package sorular

import "fmt"

type Araba struct {
	Marka string
	Model string
	Yil   int
}

func (a Araba) BilgileriGoster() {
	fmt.Printf("Araba Bilgileri:\nMarka: %s\nModel: %s\nYıl: %d\n", a.Marka, a.Model, a.Yil)
}

func Demo2() {
	fmt.Println(Araba{"Nissan", "Qashqai", 2015})
}
