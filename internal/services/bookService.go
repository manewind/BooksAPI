package services

import (
	//"context"
	"database/sql"
	//"fmt"
)


type BookDatabase struct{
	db *sql.DB
}


func NewBookDatabase(db *sql.DB) *BookDatabase{
	return &BookDatabase{db:db}
}

func AddBooks(bd *BookDatabase)
