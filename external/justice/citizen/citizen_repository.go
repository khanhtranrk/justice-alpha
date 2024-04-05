package justice

import (
	"database/sql"
	"fmt"
)

type CitizenRepository struct {
  db *sql.DB
}

func NewCitizenRepository(db *sql.DB) *CitizenRepository {
  return &CitizenRepository{db}
}

func (cr *CitizenRepository) ListAll() ([]*Citizen, error) {
  query := `SELECT id, gate_id FROM citizens`
  rows, err := cr.db.Query(query)
  if err != nil {
    return nil, err
  }

  var citizens []*Citizen
  for rows.Next() {
    var citizen Citizen

    err = rows.Scan(&citizen.Id, &citizen.GateId)
    if err != nil {
      return nil, err
    }

    citizens = append(citizens, &citizen)
  }

  return citizens, nil
}

func (cr *CitizenRepository) GetCitizenById(id uint64) (*Citizen, error) {
  query := "SELECT id, gate_id FROM citizens WHERE id = ?"
  rows, err := cr.db.Query(query, id)
  if err != nil {
    return nil, err
  }

  if !rows.Next() {
    return nil, fmt.Errorf("Citizen not found: %v", id)
  }

  var citizen Citizen
  err = rows.Scan(&citizen.Id, &citizen.GateId)
  if err != nil {
    return nil, err
  }

  return &citizen, nil
}
