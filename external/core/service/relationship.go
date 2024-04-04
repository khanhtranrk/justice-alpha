package service

import (
	"fmt"

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

func (rs *RelationshipService) IsTrustGuest(guestKey string, safeKey string) (bool, error) {
  r, err := rs.rr.GetRelationshipByGuestKey(guestKey)
  if err != nil {
    return false, err
  }

  if r == nil {
    return false, nil
  }

  if r.SafeKey != safeKey {
    return false, nil
  }

  return true, nil
}

func (rs *RelationshipService) UpdateSafeKeyForGuest(guestKey string, safeKey string) error {
  r, err := rs.rr.GetRelationshipByGuestKey(guestKey)
  if err != nil {
    return err
  }


  if r == nil {
    return fmt.Errorf("guestKey not found")
  }

  r.SafeKey = safeKey
  rs.rr.UpdateRelationship(r)

  return nil
}

