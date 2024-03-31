package createdbgo

import (
	"fmt"
	"testing"
)

func TestInsertTugasakhir(t *testing.T) {
	tugas2 := "Membuat chatbot web menggunakan API GPT-3.5-turbo"
	tugas3 := "1. Kirito 2. Genos"
	phonenumber := "9999111999"
	biodata := Mahasiswa{
		Nama:         "Kirito",
		Phone_number: "9999111999",
		Kelas:        "2B",
		Alamat: Alamat{
			Nama_jalan: "alfheim",
			No_rumah:   "21",
			RT:         "22",
			RW:         "23",
			Kode_pos:   8888,
		},
	}
	hasil := InsertTugasakhir(tugas2, tugas3, "Pemrograman 10", phonenumber, biodata)
	fmt.Println(hasil)
}

func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
	phonenumber := "9999111999"
	biodata := GetMahasiswaFromPhone(phonenumber)
	fmt.Println(biodata)
}

func TestGetAllTugas(t *testing.T) {
	data := GetAllTugas()
	fmt.Println(data)
}
