package justice

import (
	"database/sql"
	"fmt"
)

type ContactRepository struct {
  db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
  return &ContactRepository{db}
}

func (cr *ContactRepository) ListAll() ([]*Contact, error) {
  query := `SELECT citizen_id, safe_key FROM contacts`
  rows, err := cr.db.Query(query)
  if err != nil {
    return nil, err
  }

  var contacts []*Contact
  for rows.Next() {
    var contact Contact

    err = rows.Scan(&contact.CitizenId, &contact.SafeKey)
    if err != nil {
      return nil, err
    }

    contacts = append(contacts, &contact)
  }

  return contacts, nil
}

func (cr *ContactRepository) GetContactByCitizenId(id uint64) (*Contact, error) {
  query := "SELECT citizen_id, safe_key FROM contacts WHERE id = ?"
  rows, err := cr.db.Query(query, id)
  if err != nil {
    return nil, err
  }

  if !rows.Next() {
    return nil, fmt.Errorf("Contact not found: %v", id)
  }

  var contact Contact
  err = rows.Scan(&contact.CitizenId, &contact.SafeKey)
  if err != nil {
    return nil, err
  }

  return &contact, nil
}
