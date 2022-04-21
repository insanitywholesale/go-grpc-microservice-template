package postgres_pgx

import (
	"context"

	"github.com/jackc/pgx/v4"
	models "gitlab.com/insanitywholesale/go-grpc-microservice-template/models/v1"
	pb "gitlab.com/insanitywholesale/go-grpc-microservice-template/proto/v1"
)

type postgresRepo struct {
	client *pgx.Conn
	pgURL  string
}

var ctx = context.Background()

func newPostgresClient(url string) (*pgx.Conn, error) {
	client, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(ctx, createHellosTableQuery)
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
	err := r.client.QueryRow(ctx, helloInsertQuery, hr.HelloWord).Scan(&id)
	if err != nil {
		return err
	}
	hr.Id = id
	return nil
}
