package service

import (
	"github.com/khanhtranrk/justice-alpha/external/core/domain"
	"github.com/khanhtranrk/justice-alpha/external/core/port"
)

type RelationshipService struct {
  rr port.RelationshipRepository
}

func NewRelationshipService(rr port.RelationshipRepository) *RelationshipService {
  return &RelationshipService{rr}
}

func (rs *RelationshipService) ListRelationships() ([]*domain.Relationship, error) {
  return rs.rr.ListRelationships()
}

func (rs *RelationshipService) GetRelationshipByGuestKey(guestKey string) (*domain.Relationship, error) {
  return rs.rr.GetRelationshipByGuestKey(guestKey)
}
