package service

import (
	"github.com/khanhtranrk/justice-alpha/external/core/domain"
	"github.com/khanhtranrk/justice-alpha/external/core/port"
)

type RequestMessageService struct {
  rr port.RelationshipRepository
}

func (rms *RequestMessageService) IsTrust(rm *domain.RequestMessage) (bool, error) {
  relationship, err := rms.rr.GetRelationshipByGuestKey(rm.RequesterId)

  if err != nil {
    return false, err
  }

  if relationship.RequestKey != rm.SafeKey {
    return false, nil
  }

  return true, nil
}
