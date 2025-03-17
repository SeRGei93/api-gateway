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

func NewMysqlClientStore(config *dbv4.Config) *ClientStore {
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	return &ClientStore{db: db, tableName: "oauth2_clients"}
}

func (cs *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=? LIMIT 1", cs.tableName)
	item := &models.Client{}
	err := cs.db.QueryRow(query, id).Scan(
		&item.ID,
		&item.Secret,
		&item.Domain,
		&item.Public,
		&item.UserID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return item, nil
}
