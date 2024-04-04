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
  query := `SELECT relationship_key, owner_key, guest_key,
                   guest_request_queue, guest_response_queue,
                   request_key, response_key, created_at, updated_at
            FROM relationships
            WHERE owner_key = ?`

  rows, err := rr.db.Query(query)

  if err != nil {
    return nil, err
  }

  var relationship *domain.Relationship

  if rows.Next() {
    err := rows.Scan(
      relationship.RelationshipKey,
      relationship.OwnerKey,
      relationship.GuestKey,
      relationship.GuestRequestQueue,
      relationship.GuestResponseQueue,
      relationship.RequestKey,
      relationship.ResponseKey,
      relationship.CreatedAt,
      relationship.UpdatedAt,
    )

    if err != nil {
      return nil, err
    }

    return relationship, nil
  }

  return nil, nil
}

func (rr *RelationshipRepository) ListRelationships() ([]*domain.Relationship, error) {
  query := `SELECT relationship_key, owner_key, guest_key,
                   guest_request_queue, guest_response_queue,
                   request_key, response_key, created_at, updated_at
            FROM relationships`

  rows, err := rr.db.Query(query)

  if err != nil {
    return nil, err
  }

  var relationships []*domain.Relationship
  for rows.Next() {
    var relationship *domain.Relationship

    err := rows.Scan(
      relationship.RelationshipKey,
      relationship.OwnerKey,
      relationship.GuestKey,
      relationship.GuestRequestQueue,
      relationship.GuestResponseQueue,
      relationship.RequestKey,
      relationship.ResponseKey,
      relationship.CreatedAt,
      relationship.UpdatedAt,
    )

    if err != nil {
      return nil, err
    }

    relationships = append(relationships, relationship)
  }

  return relationships, nil
}

