package sorular

import "fmt"

type Sarki struct {
	Adi     string
	Sanatci string
	Uzunluk int
	Tarih   int
}

func (s Sarki) SarkiBilgileriYazdir() {
	fmt.Println("Adı:", s.Adi)
	fmt.Println("Sanatçı:", s.Sanatci)
	fmt.Println("Uzunluk:", s.Uzunluk, "dakika")
	fmt.Println("Tarih:", s.Tarih)
}

func Demo7() {
	fmt.Println(Sarki{"Aşk Kırıntıları", "Teoman", 4, 2006})
}
