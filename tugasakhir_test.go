package createdbgo

import (
	"fmt"
	"testing"
)

func TestInsertTugasakhir(t *testing.T) {
	tugas2 := "Membuat chatbot web menggunakan API ChatGpt"
	tugas3 := "1. Kirito 2. Genos"
	phonenumber := "6284455621"
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
	hasil := InsertTugasakhir(tugas2, tugas3, "Pemrograman", phonenumber, biodata)
	fmt.Println(hasil)
}

func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
	phonenumber := "9999111999"
	biodata := GetMahasiswaFromPhone(phonenumber)
	fmt.Println(biodata)
}

func TestGetTugasFromNama(t *testing.T) {
	mhs, err := GetTugasFromNama("Kirito")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Nilai Mahasiswa:", mhs.Nama_tugasBesar)
	}
}

func TestGetNamaFromAlamat(t *testing.T) {
	mhs, err := GetNamaFromAlamat("Jalan Parapat")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Nama Mahasiswa:", mhs.Biodata.Nama)
	}
}

func TestGetAllTugas(t *testing.T) {
	data := GetAllTugas()
	fmt.Println(data)
}
