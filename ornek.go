package sorular

import "fmt"

type Kisi struct {
	adi    string
	soyAdi string
	yasi   int
}

func (k Kisi) BilgileriGoster() {
	fmt.Printf("Kişi Bilgileri:\nAdı: %s\nSoyadı: %s\nYaşı: %d\n", k.adi, k.soyAdi, k.yasi)
}

func Demo1() {
	fmt.Println(Kisi{"Nefise", "Öztürk", 18})
}
