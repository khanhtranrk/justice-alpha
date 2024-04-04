package domain

import "time"

type Relationship struct {
  RelationshipKey string
  OwnerKey string
  GuestKey string
  GuestRequestQueue string
  GuestResponseQueue string
  RequestKey string
  ResponseKey string
  CreatedAt time.Time
  UpdatedAt time.Time
}

func NewRelationship(
  RelationshipKey string,
  ownerKey string,
  guestKey string,
  guestRequestQueue string,
  GuestResponseQueue string,
  RequestKey string,
  ResponseKey string,
  CreatedAt time.Time,
  UpdatedAt time.Time,
) *Relationship {
  return &Relationship{
    RelationshipKey,
    ownerKey,
    guestKey,
    guestRequestQueue,
    GuestResponseQueue,
    RequestKey,
    ResponseKey,
    CreatedAt,
    UpdatedAt,
  }
}

