package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func Albums(config mysql.Config) ([]Album, error) {
	var albums []Album

	db := fetchDB(config)

	rows, err := db.Query("SELECT * FROM albums")
	if err != nil {
		return nil, fmt.Errorf("Albums: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var item Album

		if err := rows.Scan(&item.ID, &item.Title, &item.Artist, &item.Price); err != nil {
			return nil, fmt.Errorf("Album %v", err)
		}

		albums = append(albums, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Album %v", err)
	}

	return albums, nil
}

func AlbumByArists(config mysql.Config, name string) ([]Album, error) {
	var albums []Album

	db := fetchDB(config)

	rows, err := db.Query("SELECT * FROM albums WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumByArists %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var item Album

		if err := rows.Scan(&item.ID, &item.Title, &item.Artist, &item.Price); err != nil {
			return nil, fmt.Errorf("albumByArists %q: %v", name, err)
		}

		albums = append(albums, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumByArists %q: %v", name, err)
	}

	return albums, nil
}

func AlbumByID(config mysql.Config, id string) (Album, error) {
	var item Album

	db := fetchDB(config)

	row := db.QueryRow("SELECT * FROM albums WHERE ID = ?", id)
	if err := row.Scan(&item.ID, &item.Title, &item.Artist, &item.Price); err != nil {
		if err == sql.ErrNoRows {
			return item, fmt.Errorf("albumByID %d: no such album", id)
		}

		return item, fmt.Errorf("albumByID %d: %v", id, err)
	}

	return item, nil
}

// func AddAlbum(db *sql.DB, album Album) (int64, error) {
// 	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
// 	if err != nil {
// 		return 0, fmt.Errorf("addAlbum %v", err)
// 	}
//
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, fmt.Errorf("addAlbum %v", err)
// 	}
//
// 	return id, nil
// }

func fetchDB(config mysql.Config) *sql.DB {
	var err error
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")

	return db
}
