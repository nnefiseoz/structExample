package sorular

import "fmt"

type Film struct {
	Baslik string
	Yil    int
	Tur    string
}

func (f Film) FilmBilgileriYazdir() {
	fmt.Println("Başlık:", f.Baslik)
	fmt.Println("Yıl:", f.Yil)
	fmt.Println("Tür:", f.Tur)
}

func Demo5() {
	fmt.Println(Film{"Delibal", 2015, "Romantk/Dram"})
}
