package db

import (
	"api/config"
	"api/model"
	"errors"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDataMahasiswa(requestData model.DataMahasiswa) error {
	db := mongo.MongoConnect(DBATS)
	insertedID := mongo.InsertOneDoc(db, config.MahasiswaColl, requestData)
	if insertedID == nil {
		return errors.New("couldn't insert data")
	}
	return nil
}

func CountDataMahasiswa(version string) error {
	db := mongo.MongoConnect(DBATS)
	filter := bson.M{"codename": version}
	checkData, err := mongo.CountDocuments(db, config.MahasiswaColl, filter)
	if err != nil {
		return err // Return the error if there's an issue with counting documents
	}
	if checkData > 0 {
		return errors.New("data already exists") // Return an error if data already exists
	}
	return nil
}

func GetDataMahasiswaFilter(filter bson.M) ([]model.DataMahasiswa, error) {
	db := mongo.MongoConnect(DBATS)
	var datas []model.DataMahasiswa
	err := mongo.GetAllDocByFilter[model.DataMahasiswa](db, config.MahasiswaColl, filter, &datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetOneDataMahasiswaFilter(filter bson.M) (model.DataMahasiswa, error) {
	db := mongo.MongoConnect(DBATS)
	var data model.DataMahasiswa
	err := mongo.GetOneDocByFilter[model.DataMahasiswa](db, config.MahasiswaColl, filter, &data)
	if err != nil {
		return model.DataMahasiswa{}, err
	}
	return data, nil
}
