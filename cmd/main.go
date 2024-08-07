package main

import (
	"github.com/dayroromero/storiChallenge/pkg/csv"
	"github.com/dayroromero/storiChallenge/pkg/db"
)

func main() {
	dbUrl := "postgres://postgres:postgres@localhost:5432/stori"

	db.Init(dbUrl)

	csv.File_processor()
}
