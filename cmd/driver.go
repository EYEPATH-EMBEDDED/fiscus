package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbFile = "./usage_log.db"

func InitDB() (*sql.DB, error) {
	// 파일이 없으면 새로 생성
	_, err := os.Stat(dbFile)
	firstRun := os.IsNotExist(err)

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, fmt.Errorf("DB 연결 실패: %w", err)
	}

	if firstRun {
		if err := createTables(db); err != nil {
			return nil, fmt.Errorf("테이블 생성 실패: %w", err)
		}
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS usage_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT NOT NULL,
		start_time DATETIME NOT NULL,
		end_time DATETIME NOT NULL,
		photos INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(schema)
	return err
}
