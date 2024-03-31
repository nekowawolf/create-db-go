package createdbgo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama          string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number  string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Alamat        Alamat             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Kelas         string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
	Jenis_kelamin string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty"`
}

type Alamat struct {
	Nama_jalan string `bson:"nama_jalan,omitempty" json:"nama_jalan,omitempty"`
	No_rumah   string `bson:"no_rumah,omitempty" json:"no_rumah,omitempty"`
	RT         string `bson:"rt,omitempty" json:"rt,omitempty"`
	RW         string `bson:"rw,omitempty" json:"rw,omitempty"`
	Kode_pos   int    `bson:"kode_pos,omitempty" json:"kode_pos,omitempty"`
}

type Tugasakhir struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_tugasBesar string             `bson:"nama_tugasbesar,omitempty" json:"nama_tugasbesar,omitempty"`
	Nama_kelompok   string             `bson:"nama_kelompok,omitempty" json:"nama_kelompok,omitempty"`
	Matakuliah      string             `bson:"matakuliah,omitempty" json:"matakuliah,omitempty"`
	Biodata         Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
