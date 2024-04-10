package repository

import (
	"database/sql"

	"github.com/khanhtranrk/justice-alpha/external/core/domain"
)

type LetterRepository struct {
  db *sql.DB
}

func NewLetterRepository(db *sql.DB) *LetterRepository {
  return &LetterRepository{db}
}

func (lt *LetterRepository) CreateLetter(letter *domain.Letter, tableName string) (*domain.Letter, error) {
  exec := "INSERT INTO " + tableName + " (code, foreign_id, commit_time, message, status) VALUES (?, ?, ?, ?, ?)"
  result, err := lt.db.Exec(
    exec,
    letter.Code,
    letter.ForeignId,
    letter.CommitTime,
    letter.Message,
    letter.Status,
  )
  if err != nil {
    return nil, err
  }

  newId, err := result.LastInsertId()
  if (err != nil) {
    newId = 0
  }

  letter.Id = uint64(newId)

  return letter, nil
}

func (lt *LetterRepository) UpdateLetter(letter *domain.Letter, tableName string) (*domain.Letter, error) {
  exec := "UPDATE " + tableName + " SET code = ?, foreign_id = ?, commit_time = ?, message = ?, status = ? WHERE id = ?"
  _, err := lt.db.Exec(
    exec,
    letter.Code,
    letter.ForeignId,
    letter.CommitTime,
    letter.Message,
    letter.Status,
    letter.Id,
  )
  if err != nil {
    return nil, err
  }

  return letter, nil
}


func (lt *LetterRepository) DeleteLetter(letter *domain.Letter, tableName string) (*domain.Letter, error) {
  exec := "DELETE FROM " + tableName + " WHERE id = ?"
  _, err := lt.db.Exec(exec, letter.Id)
  if err != nil {
    return nil, err
  }

  return letter, nil
}

func (lt *LetterRepository) GetLetterById(id uint64, tableName string) (*domain.Letter, error) {
  query := "SELECT id, code, foreign_id, commit_time, message, status From " + tableName + " WHERE id = ?"
  rows, err := lt.db.Query(query, id)
  if err != nil {
    return nil, err
  }

  var letter domain.Letter
  if rows.Next() {
    rows.Scan(
      &letter.Id,
      &letter.Code,
      &letter.ForeignId,
      &letter.CommitTime,
      &letter.Message,
      &letter.Status,
    )
  } else {
    return nil, nil
  }
  defer rows.Close()

  return &letter, nil
}
