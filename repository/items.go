package repository

import (
	"database/sql"
	"time"

	"goTestApi/models"
)

type ItemStore struct {
	db *sql.DB
}

func NewItemStore(db *sql.DB) *ItemStore {
	return &ItemStore{
		db: db,
	}
}

func (t *ItemStore) GetAll() ([]models.Items, error) {
	var m []models.Items

	// Get a Tx for making transaction requests.
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// query
	rows, err := t.db.Query("select id, name, price, description, created_at from items")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		var i models.Items

		// unmarshal data
		err = rows.Scan(&i.ID, &i.Name, &i.Price, &i.Description, &i.CreatedAt)
		if err != nil {
			// log.Fatalf("Unable to scan the row. %v", err)
			return nil, err
		}

		m = append(m, i)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return m, nil
}

func (t *ItemStore) Create(i *models.Items) (*models.Items, error) {
	// Get a Tx for making transaction requests.
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// define and submit query
	created := time.Now()
	stmt, err := t.db.Prepare("insert into items (name, price, description, created_at) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(i.Name, i.Price, i.Description, created)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// assign values back to struct
	i.ID = id
	i.CreatedAt = created

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return i, nil
}

func (t *ItemStore) GetById(id int) (*models.Items, error) {
	// Get a Tx for making transaction requests.
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	var i models.Items

	// submit query
	if err := t.db.QueryRow(
		"select id, name, price, description, created_at from items where id = ?",
		id,
	).Scan(&i.ID, &i.Name, &i.Price, &i.Description, &i.CreatedAt); err != nil {
		return nil, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &i, nil
}