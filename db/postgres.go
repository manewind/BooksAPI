package db

import (
	"BooksAPI/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConnectionManager struct {
	conn *sql.DB
}

func (cm *ConnectionManager) Connect(config config.DBConfig) (*sql.DB, error) {
	if cm.conn != nil {
		return cm.conn, nil
	}
	connStr := fmt.Sprintf("host=%s port = %s user = %s password = %s dbname = %s sslmode = disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db,err := sql.Open("postgres",connStr)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to databse: %w",err)
	}

	if err := db.Ping();err !=nil{
		return nil, fmt.Errorf("database iss unreachable: %w", err)
	}

	cm.conn = db
	return cm.conn,nil
}


func (cm *ConnectionManager)Close()error{
	if cm.conn != nil{
		err := cm.conn.Close()
		cm.conn = nil
		return err
	}
	return nil
}
