package createdbgo

import (
	"testing"
)

func TestInsertTugasakhir(t *testing.T) {
	biodata := Mahasiswa{
		Name:  "Venti",
		Kelas: "2B",
		Alamat: Alamat{
			Nama_jalan: "Mondstadt",
			No_rumah:   "Stormterror",
			RT:         "06",
			RW:         "66",
		},
	}

	insertedID, err := InsertTugasakhir("Membuat chatbot menggunakan API GPT-3.5-turbo", "1.Venti 2.Lisa", "Pemrograman", biodata)
	if err != nil {
		t.Errorf("InsertTugasakhir failed: %v", err)
		return
	}
	t.Logf("Inserted ID: %v", insertedID)
}
