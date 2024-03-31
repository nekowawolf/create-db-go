package createdbgo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama          string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number  string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Alamat        Alamat             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Kelas         string             `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Jenis_kelamin string             `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty"`
}

type Alamat struct {
	Nama_jalan string `bson:"durasi,omitempty" json:"durasi,omitempty"`
	No_rumah   string `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	RT         string `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	RW         string `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Kode_pos   int    `bson:"hari,omitempty" json:"hari,omitempty"`
}

type Tugasakhir struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_tugasBesar string             `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Nama_kelompok   string             `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Matakuliah      string             `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number    string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Biodata         Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
