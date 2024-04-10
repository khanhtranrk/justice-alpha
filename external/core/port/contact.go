package port

import "github.com/khanhtranrk/justice-alpha/external/core/domain"

type ContactRepository interface {
  ListAllContacts() ([]*domain.Contact, error)
  GetContactByCitizenId(id uint64) (*domain.Contact, error)
}

type ContactService interface {
  ListAllContacts() ([]*domain.Contact, error)
  GetContactByCitizenId(id uint64) (*domain.Contact, error)
}
