package port

import "github.com/khanhtranrk/justice-alpha/external/core/domain"

type RelationshipRepository interface {
  CreateChain(relationship *domain.Relationship) error
}

type RelationshipService interface {
  CreateChain(relationship *domain.Relationship) error
}
