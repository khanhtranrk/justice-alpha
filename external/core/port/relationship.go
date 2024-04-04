package port

import "github.com/khanhtranrk/justice-alpha/external/core/domain"

type RelationshipRepository interface {
  GetRelationshipByGuestKey(guestKey string) (*domain.Relationship, error)
  ListRelationships() ([]*domain.Relationship, error)
  UpdateRelationship(relationship *domain.Relationship) (*domain.Relationship, error)
}

type RelationshipService interface {
  ListRelationships() ([]*domain.Relationship, error)
}
