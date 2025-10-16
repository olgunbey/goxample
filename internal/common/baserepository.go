package common

import (
	"database/sql"
	"fmt"
	"log"
)

type BaseRepository[T any] struct {
	Db        *sql.DB
	TableName string
}

func NewGenericRepository[T any](db *sql.DB, tableName string) *BaseRepository[T] {
	return &BaseRepository[T]{Db: db, TableName: tableName}
}

func (r *BaseRepository[T]) Add(item *T) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s DEFAULT VALUES RETURNING id", r.TableName)

	var id int
	err := r.Db.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert failed: %w", err)
	}

	return id, nil
}

func (r *BaseRepository[T]) RemoveGetById(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", r.TableName)

	var deletedId int
	err := r.Db.QueryRow(query, id).Scan(&deletedId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("record with id %d not found", id)
		}
		return 0, fmt.Errorf("delete failed: %w", err)
	}

	return deletedId, nil
}

func (r *BaseRepository[T]) GetAll() ([]T, error) {
	query := fmt.Sprintf("SELECT * FROM %s", r.TableName)

	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var result []T

	for rows.Next() {
		var entity T
		err := rows.Scan(&entity)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		result = append(result, entity)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return result, nil
}

func (r *BaseRepository[T]) GetById(id int) (*T, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", r.TableName)

	var entity T
	err := r.Db.QueryRow(query, id).Scan(&entity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &entity, nil
}
