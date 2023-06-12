package sorular

import "fmt"

type Kitap struct {
    Baslik string
    Yazar  string
    Yil    int
}

func (k Kitap) BilgileriGoster() {
    fmt.Printf("Kitap Bilgileri:\nBaşlık: %s\nYazar: %s\nYıl: %d\n", k.Baslik, k.Yazar, k.Yil)
}

func Demo3() {
    fmt.Println(Kitap{"Kürk Mantolu Madonna","Sabahattin Ali",1943})
}