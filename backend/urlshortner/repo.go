package urlshortner

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type URLShortnerRepo struct {
	Db *sql.DB
}

func (r *URLShortnerRepo) Insert(s ShortenedURL) error {
	fmt.Println(s.Key)
	stmt, err := r.Db.Prepare("INSERT INTO urls(key, url, banned, created_at) values(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(s.Key, s.URL, s.Banned, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *URLShortnerRepo) Find(s string) (string, error) {
	q := fmt.Sprintf("SELECT url FROM urls where key='%s';", s)
	rows, err := r.Db.Query(q)
	if err != nil {
		return "", err
	}
	var url string
	for rows.Next() {
		err = rows.Scan(&url)
	}
	rows.Close()
	return url, nil
}

func (r *URLShortnerRepo) Exists(s string) (string, error) {
	q := fmt.Sprintf("SELECT key FROM urls where url='%s';", s)
	rows, err := r.Db.Query(q)
	if err != nil {
		return "", err
	}
	var key string
	for rows.Next() {
		err = rows.Scan(&key)
	}
	rows.Close()
	return key, nil
}
