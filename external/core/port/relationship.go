package port

import "github.com/khanhtranrk/justice-alpha/external/core/domain"

type RelationshipRepository interface {
  ListRelationships() ([]*domain.Relationship, error)
  CreateRelationship(relationship *domain.Relationship) (*domain.Relationship, error)
  UpdateRelationship(relationship *domain.Relationship) (*domain.Relationship, error)
  DeleteRelationship(relationship *domain.Relationship) (*domain.Relationship, error)
}

type RelationshipService interface {
  ListRelationships() ([]*domain.Relationship, error)
  CreateRelationship(relationship *domain.Relationship) (*domain.Relationship, error)
  UpdateRelationship(relationship *domain.Relationship) (*domain.Relationship, error)
  DeleteRelationship(relationship *domain.Relationship) (*domain.Relationship, error)
}
