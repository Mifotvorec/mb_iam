package models

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// User structure
type Logger struct {
	Date time.Time `json:"date"`
	Sid  string    `json:"sid"`
	Text string    `json:"text"`
}

// Function save one User to database
func Logsave(Sid string, Text string) error {

	stmt, err := database.Prepare(`
		INSERT INTO logs (timestamp, uuid, text)
		VALUES ($1, $2, $3)
	`)
	if err != nil {
		fmt.Println(time.Now(), "DB1:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(time.Now(), Sid, Text)

	if err != nil {
		fmt.Println(time.Now(), "DB2:", err)
		return err
	}
	return nil
}
