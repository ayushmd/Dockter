package sqlite

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var dbDir string = ""

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CreateConn() (*sql.DB, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dbDirPath := filepath.Join(pwd, "db")
	if !exists(dbDirPath) {
		err = os.Mkdir("db", os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	db, err := sql.Open("sqlite3", filepath.Join(dbDirPath, "dns.db"))
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS dns (Subdomain TEXT PRIMARY KEY, HostIp TEXT, HostPort TEXT, RunningPort TEXT, ImageName TEXT, ContainerID TEXT)")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Subdomain   string //a unique id
// 	URL         url.URL
// 	Hostport    string
// 	Runningport string
// 	ImageName   string
