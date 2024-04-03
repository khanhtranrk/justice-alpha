package domain

type Relationship struct {
  RelationshipKey string
  OwnerKey string
  GuestKey string
  GuestRequestQueue string
  GuestResponseQueue string
}

func NewRelationship(RelationshipKey string, ownerKey string, guestKey string, guestRequestQueue string, GuestResponseQueue string) *Relationship {
  return &Relationship{RelationshipKey, ownerKey, guestKey, guestRequestQueue, GuestResponseQueue}
}

