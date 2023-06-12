package sorular

import "fmt"

type Ulke struct {
	Isim    string
	Baskent string
}

func (u Ulke) UlkeBilgileriYazdir() {
	fmt.Println("Ülke İsmi:", u.Isim)
	fmt.Println("Başkent:", u.Baskent)
}

func Demo8() {
	fmt.Println(Ulke{"Türkiye", "Ankara"})
}
