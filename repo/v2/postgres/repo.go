package postgres

import (
	"database/sql"

	_ "github.com/lib/pq" // Driver for postgres, can be swapped out
	models "gitlab.com/insanitywholesale/go-grpc-microservice-template/models/v2"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v2"
)

type postgresRepo struct {
	client *sql.DB
	pgURL  string
}

func newPostgresClient(url string) (*sql.DB, error) {
	client, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = client.Ping()
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createHellosTableQuery)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewPostgresRepo(url string) (models.HelloRepo, error) {
	pgclient, err := newPostgresClient(url)
	if err != nil {
		return nil, err
	}
	repo := &postgresRepo{
		pgURL:  url,
		client: pgclient,
	}
	return repo, nil
}

func (r *postgresRepo) StoreHello(hr *pb.HelloResponse) error {
	var id uint32
	err := r.client.QueryRow(helloInsertQuery, hr.HelloWord).Scan(&id)
	if err != nil {
		return err
	}
	hr.Id = id
	return nil
}