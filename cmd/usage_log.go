package main

import (
	"database/sql"
	"time"
)

type MonthlyUsage struct {
	TotalDurationMinutes int
	TotalPhotos          int
}

func InsertUsageLog(db *sql.DB, log UsageLog) error {
	_, err := db.Exec(`
		INSERT INTO usage_log (user_id, start_time, end_time, photos)
		VALUES (?, ?, ?, ?)
	`, log.UserID, log.StartTime, log.EndTime, log.Photos)
	return err
}

func GetMonthlyUsage(db *sql.DB, userID string, year, month int) (MonthlyUsage, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	row := db.QueryRow(`
		SELECT 
			COALESCE(SUM(strftime('%s', end_time) - strftime('%s', start_time)) / 60, 0) AS total_minutes,
			COALESCE(SUM(photos), 0) AS total_photos
		FROM usage_log
		WHERE user_id = ?
		  AND start_time >= ?
		  AND start_time < ?
	`, userID, startDate, endDate)

	var result MonthlyUsage
	err := row.Scan(&result.TotalDurationMinutes, &result.TotalPhotos)
	if err != nil {
		return MonthlyUsage{}, err
	}

	return result, nil
}
