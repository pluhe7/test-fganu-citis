package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"

	"github.com/pluhe7/test-fganu-citis/controllers"
	"github.com/pluhe7/test-fganu-citis/db"
	"github.com/pluhe7/test-fganu-citis/model"
)

func main() {
	dbConn := db.Connect()
	defer db.Disconnect(dbConn)
	recs := make(chan *model.Record)
	for {
		go findRecord(recs, dbConn)
		rec, ok := <-recs
		go updateRecord(rec, dbConn)
		if !ok {
			break
		}
	}
}

func findRecord(recs chan *model.Record, db *pg.DB) {
	rec, err := controllers.GetRecordByStatus(1, db)
	if err != nil {
		fmt.Println("Couldn't find record")
		close(recs)
		return
	}
	rec.Status = 2
	if err != nil {
		fmt.Println("Couldn't update record")
		return
	}
	recs <- rec
}

func updateRecord(rec *model.Record, db *pg.DB) {
	time.Sleep(1000 * time.Millisecond)
	rec.Status = 3
	rec.Comment = "Работа выполнена"
	err := controllers.UpdateRecord(rec, db)
	if err != nil {
		fmt.Println("Couldn't update record")
		return
	}
}
