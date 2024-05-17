package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func Albums(pgpool *pgxpool.Pool) ([]Album, error) {
	var albums []Album

	rows, err := pgpool.Query(context.Background(), "SELECT * FROM albums")
	if err != nil {
		return nil, fmt.Errorf("albums %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var album Album

		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albums %v", err)
		}

		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albums %v", err)
	}

	return albums, nil
}

func AlbumByID(id string) (Album, error) {
	var album Album

	dbpool := fetchDB()
	defer dbpool.Close()

	row := dbpool.QueryRow(context.Background(), "SELECT * FROM albums WHERE id = $1", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumByID %s: no such album", id)
		}

		return album, fmt.Errorf("albumByID %s: %v", id, err)
	}

	return album, nil
}

func AlbumByArists(name string) ([]Album, error) {
	var albums []Album

	dbpool := fetchDB()
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "SELECT * FROM albums WHERE artist = ?", name)
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

func fetchDB() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("error create pgxpool %v", err)
	}

	connection, err := dbpool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("error acquiring a connection %v", err)
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatalf("error pinging connaction %v", err)
	}

	fmt.Println("Connected")

	return dbpool
}
