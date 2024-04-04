package port

import "github.com/khanhtranrk/justice-alpha/external/core/domain"

type RelationshipRepository interface {
  GetRelationshipByGuestKey(guestKey string) (*domain.Relationship, error)
  ListRelationships() ([]*domain.Relationship, error)
  UpdateRelationship(rl *domain.Relationship) (*domain.Relationship, error)
}

type RelationshipService interface {
  GetRelationshipByGuestKey(guestKey string) (*domain.Relationship, error)
  ListRelationships() ([]*domain.Relationship, error)
  UpdateSafeKeyForGuest(guestKey string, safeKey string) error
}
