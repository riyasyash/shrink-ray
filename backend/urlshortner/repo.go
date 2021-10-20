package urlshortner

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type URLShortnerRepo struct {
	Db *sql.DB
}

func (r *URLShortnerRepo) Insert(s ShortenedURL) error {
	stmt:="INSERT INTO urls(key, url, banned, created_at) values($1,$2,$3,$4) RETURNING ID;"
	// if err != nil {
	// 	fmt.Println(stmt)
	// 	fmt.Println("error", err.Error())
	// 	return err
	// }
	_, err := r.Db.Query(stmt, s.Key, s.URL, s.Banned, time.Now())
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
		fmt.Println("error",err.Error())
		return "", err
	}
	var key string
	for rows.Next() {
		err = rows.Scan(&key)
	}
	rows.Close()
	return key, nil
}
