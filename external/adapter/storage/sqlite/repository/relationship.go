package repository

import (
	"database/sql"

	"github.com/khanhtranrk/justice-alpha/external/core/domain"
)

type RelationshipRepository struct {
  db *sql.DB
}

func NewRelationshipRepository(db *sql.DB) *RelationshipRepository {
  return &RelationshipRepository{db}
}

func (rr *RelationshipRepository) GetRelationshipByGuestKey(guestKey string) (*domain.Relationship, error) {
  query := `SELECT owner_key, guest_key, safe_key, gate
            FROM relationships
            WHERE owner_key = ?`

  rows, err := rr.db.Query(query, guestKey)

  if err != nil {
    return nil, err
  }

  var relationship domain.Relationship

  if rows.Next() {
    err := rows.Scan(
      &relationship.OwnerKey,
      &relationship.GuestKey,
      &relationship.SafeKey,
      &relationship.Gate,
    )

    if err != nil {
      return nil, err
    }

    return &relationship, nil
  }

  return nil, nil
}

func (rr *RelationshipRepository) ListRelationships() ([]*domain.Relationship, error) {
  query := `SELECT owner_key, guest_key, safe_key, gate
            FROM relationships`

  rows, err := rr.db.Query(query)

  if err != nil {
    return nil, err
  }

  var relationships []*domain.Relationship
  for rows.Next() {
    var relationship domain.Relationship

    err := rows.Scan(
      &relationship.OwnerKey,
      &relationship.GuestKey,
      &relationship.SafeKey,
      &relationship.Gate,
    )

    if err != nil {
      return nil, err
    }

    relationships = append(relationships, &relationship)
  }

  return relationships, nil
}

func (rr *RelationshipRepository) UpdateRelationship(rl *domain.Relationship) (*domain.Relationship, error) {
  exec := `UPDATE relationships SET safe_key = ? WHERE guest_key = ?`
  _, err := rr.db.Exec(exec, rl.SafeKey, rl.GuestKey)

  if err != nil {
    return nil, err
  }

  return rl, nil
}
