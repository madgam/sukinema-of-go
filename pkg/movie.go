package pkg

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

// DbInfo : データベース接続情報
type DbInfo struct {
	User         string `default:"root"`
	Password     string `default:"root"`
	Host         string `default:"localhost"`
	Databasename string `default:"mysql"`
}

// Movie : 映画情報の構造体
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Pref        string `json:"pref"`
	Theater     string `json:"theater"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Time        string `json:"time"`
	AllTime     string `json:"all_time"`
	Review      string `json:"review"`
	ReleaseDate string `json:"release_date"`
	DropPath    string `json:"drop_path"`
	PosterPath  string `json:"poster_path"`
}

// Movies : 映画情報の配列
type Movies *[]Movie

const (
	// SelAll 全件取得時のSQL
	SelAll = "SELECT DISTINCT ID, TITLE, PREF, THEATER, DESCRIPTION, LINK, LATITUDE, LONGITUDE, TIME, ALL_TIME, REVIEW, RELEASE_DATE, DROP_PATH, POSTER_PATH FROM MOVIES"
	// SelPref 県別の情報取得時のSQL
	SelPref = "SELECT DISTINCT ID, TITLE, PREF, THEATER, DESCRIPTION, LINK, LATITUDE, LONGITUDE, TIME, ALL_TIME, REVIEW, RELEASE_DATE, DROP_PATH, POSTER_PATH FROM MOVIES WHERE PREF = ?"
)

// GetAllMovie 映画情報をJSON形式で取得
func GetAllMovie(ctx context.Context) ([]*Movie, error) {

	con, err := GetSQLConnector()
	if err != nil {
		log.Fatal(err)
	}
	err = con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	rows, err := con.Query(SelAll)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	movies := make([]*Movie, 0)
	for rows.Next() {
		m := new(Movie)
		err := rows.Scan(&m.ID, &m.Title, &m.Pref, &m.Theater, &m.Description, &m.Link, &m.Latitude, &m.Longitude, &m.Time, &m.AllTime, &m.Review, &m.ReleaseDate, &m.DropPath, &m.PosterPath)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, m)
	}

	return movies, err
}

// GetPrefMovie 県別の映画情報をJSON形式で取得
func GetPrefMovie(ctx context.Context, prefID string) ([]*Movie, error) {

	con, err := GetSQLConnector()
	if err != nil {
		log.Fatal(err)
	}
	err = con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	stmt, err := con.Prepare(SelPref)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(prefID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	movies := make([]*Movie, 0)
	for rows.Next() {
		m := new(Movie)
		err := rows.Scan(&m.ID, &m.Title, &m.Pref, &m.Theater, &m.Description, &m.Link, &m.Latitude, &m.Longitude, &m.Time, &m.AllTime, &m.Review, &m.ReleaseDate, &m.DropPath, &m.PosterPath)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, m)
	}

	return movies, err
}

// GetSQLConnector SQLコネクターを取得
func GetSQLConnector() (*sql.DB, error) {
	info := DbInfo{}
	err := envconfig.Process("mysql", &info)
	if err != nil {
		return nil, err
	}
	con, err := sql.Open("mysql", info.User+":"+info.Password+"@tcp("+info.Host+":3306)/"+info.Databasename)

	return con, err
}
