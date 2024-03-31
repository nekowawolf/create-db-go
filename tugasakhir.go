package createdbgo

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertTugasakhir(tugas2 string, tugas3 string, matkul string, phonenumber string, Biodata Mahasiswa) (InsertedID interface{}) {
	var tugas Tugasakhir
	tugas.Nama_tugasBesar = tugas2
	tugas.Nama_kelompok = tugas3
	tugas.Matakuliah = matkul
	tugas.Biodata = Biodata
	return InsertOneDoc("week4", "tugas", tugas)
}

func GetMahasiswaFromPhone(phone_number string) (mhs Tugasakhir) {
	tugas := MongoConnect("week4").Collection("tugas")
	filter := bson.M{"phone_number": phone_number}
	err := tugas.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetMahasiswaFromPhone: %v\n", err)
	}
	return mhs
}

func GetTugasFromNama(nama string) (mhs Tugasakhir, err error) {
	tugas := MongoConnect("week4").Collection("tugas")
	filter := bson.M{"biodata.nama": nama}
	err = tugas.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetTugasFromNama: %v\n", err)
	}
	return mhs, err
}

func GetNamaFromAlamat(alamat string) (mhs Tugasakhir, err error) {
	tugas := MongoConnect("week4").Collection("tugas")
	filter := bson.M{"biodata.alamat.nama_jalan": alamat}
	err = tugas.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetNamaFromAlamat: %v\n", err)
	}
	return mhs, err
}

func GetAllTugas() (data []Tugasakhir) {
	tugas := MongoConnect("week4").Collection("tugas")
	filter := bson.M{}
	cursor, err := tugas.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllTugas :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
