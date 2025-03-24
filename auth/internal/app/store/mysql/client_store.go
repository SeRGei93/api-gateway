package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	dbv4 "github.com/go-oauth2/mysql/v4"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type ClientStore struct {
	db        *sql.DB
	tableName string
}

func NewMysqlClientStore(config *dbv4.Config) (*ClientStore, error) {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	store := ClientStore{db: db, tableName: "oauth2_clients"}
	err = store.createTable()
	if err != nil {
		return nil, err
	}

	return &store, nil
}

func (cs *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	var clientId, clientSecret, domain, userid string
	var public bool

	query := fmt.Sprintf("SELECT * FROM %s WHERE client_id=?", cs.tableName)
	err := cs.db.QueryRowContext(ctx, query, id).Scan(
		&clientId,
		&clientSecret,
		&domain,
		&public,
		&userid,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: domain,
		Public: public,
		UserID: userid,
	}, nil
}

func (cs *ClientStore) createTable() error {
	const op = "storage.client.createTable"

	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			client_id VARCHAR(255) PRIMARY KEY,
			client_secret VARCHAR(255) NOT NULL UNIQUE,
			domain VARCHAR(255) NOT NULL,
			public BOOLEAN DEFAULT FALSE,
			user_id VARCHAR(255)
		)`, cs.tableName)

	stmt, err := cs.db.Prepare(q)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
