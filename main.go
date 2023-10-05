package main

import "fmt"

type Geng interface {
	Memukul() (string, int)
	Bertahan() string
}

type Anggota struct {
	Nama string
	Pukulan int
	Status string
}

func (a Anggota) Memukul() (string, int) {
	msg := "Nama Anggota: " + a.Nama + " Dengan status: " + a.Status
	return msg, a.Pukulan
}

func (a Anggota) Bertahan() string {
	switch a.Status {
		case "Pemimpin":
			return "Kuat"
		case "Anggota":
			return "Sedang"
		default:
			return "Lemah"
	}
}

func sendPointer(data *string) string {
	d := *data
	return d
}

func sendPointerReturnPointer(data *string) *string {
	return data
}

func returnPointer(data string) *string {
	return &data
}

func main(){

	data := "halo"
	data2 := &data
	*data2 = "Baru"

	fmt.Println(data)

	fmt.Println(sendPointer(&data))

	fmt.Println(*sendPointerReturnPointer(&data))

	fmt.Println(*returnPointer(data))

	Anggota1 := Anggota{
		Nama: "Jinbe",
        Pukulan: 20,
        Status: "Anggota",
	}
	message, damage := Anggota1.Memukul()
	fmt.Println(message)
	fmt.Println(damage)

}
