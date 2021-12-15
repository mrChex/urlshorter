package storage

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mrChex/urlshorter/backend/model"
	"log"
)

type Postgresql struct {
	conn *pgx.Conn
}

func NewPostgresql(dburl string) (Storage, error) {
	conn, err := pgx.Connect(context.Background(), dburl)
	if err != nil {
		return nil, err
	}
	return &Postgresql{
		conn: conn,
	}, nil
}

// Storage interface implementation

func (p *Postgresql) PutLink(url string) (model.Link, error) {
	link := model.Link{
		URL: url,
	}
	err := p.conn.QueryRow(
		context.Background(),
		`INSERT INTO urls(url) VALUES ($1)
             ON CONFLICT(url) DO UPDATE SET url = urls.url
             RETURNING id`,
		url,
	).Scan(&link.ID)
	if err != nil {
		return link, err
	}
	return link, nil
}

func (p *Postgresql) GetLinkByID(id int64) (model.Link, error) {
	link := model.Link{
		ID: id,
	}
	err := p.conn.QueryRow(
		context.Background(),
		`SELECT url FROM urls WHERE id=$1`,
		id,
	).Scan(&link.URL)
	if err != nil {
		if err == pgx.ErrNoRows {
			return link, ErrNotFound
		}
		return link, err
	}
	return link, nil
}

// Postgres specific

func (p *Postgresql) Close() {
	if err := p.conn.Close(context.Background()); err != nil {
		log.Printf("Error while closing pg conenction: %v", err)
	}
}
