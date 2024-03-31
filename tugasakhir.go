package createdbgo

import (
	"context"
	"fmt"
	"os"

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

func InsertTugasakhir(tugasBesar, kelompok, matkul string, biodata Mahasiswa) (interface{}, error) {
	// Membuat koneksi ke MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.TODO())

	// Mendapatkan koleksi tugas dari database
	tugasCollection := client.Database("week4").Collection("tugas")

	// Membuat tugas baru
	tugas := Tugasakhir{
		tugasBesar:  tugasBesar,
		kelomok:     kelompok,
		Mata_kuliah: matkul,
		Biodata:     biodata,
	}

	// Memasukkan tugas ke dalam database
	insertResult, err := tugasCollection.InsertOne(context.TODO(), tugas)
	if err != nil {
		return nil, err
	}

	return insertResult.InsertedID, nil
}
