package domain

type Relationship struct {
  OwnerKey string
  GuestKey string
  SafeKey string
  Gate string
}

func NewRelationship(
  ownerKey string,
  guestKey string,
  safeKey string,
  gate string,
) *Relationship {
  return &Relationship{
    ownerKey,
    guestKey,
    safeKey,
    gate,
  }
}

