package createdbgo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"nama"`
	Kelas  string             `bson:"kelas"`
	Alamat Alamat             `bson:"alamt"`
}

type Alamat struct {
	Nama_jalan string `bson:"nama_jalan"`
	No_rumah   string `bson:"no_rumah"`
	RT         string `bson:"rt"`
	RW         string `bson:"rw"`
}

type Tugasakhir struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	tugasBesar  string             `bson:"nama_tugasbesar,"`
	kelomok     string             `bson:"nama_kelompok,"`
	Mata_kuliah string             `bson:"matakuliah,"`
	Biodata     Mahasiswa          `bson:"biodata,"`
}
