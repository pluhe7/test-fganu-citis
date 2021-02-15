package controllers

import (
	"github.com/go-pg/pg"
	"github.com/pluhe7/test-fganu-citis/model"
)

func GetRecords(db *pg.DB) []model.Record {
	var recs []model.Record
	err := db.Model(&recs).Select()
	if err != nil {
		panic(err)
	}
	return recs
}

func GetRecordByStatus(status int, db *pg.DB) (*model.Record, error) {
	rec := &model.Record{Status: status}
	err := db.Model(rec).Where("status = ?", rec.Status).First()
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func UpdateRecord(rec *model.Record, db *pg.DB) error {
	_, err := db.Model(rec).WherePK().Update()
	if err != nil {
		return err
	}
	return nil
}
