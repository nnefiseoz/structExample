package sorular

import "fmt"

type Urun struct {
	Adi   string
	Marka string
	Fiyat int
}

func (u Urun) BilgileriGoster() {
	fmt.Printf("Ürün Adı: %s\nMarka: %s\nFiyat: %d\n", u.Adi, u.Marka, u.Fiyat)
}

func Demo4() {
	fmt.Println("Karam", "Ülker", 7)
}
