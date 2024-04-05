package justice

type IContactRepository interface {
  ListAll() ([]*Contact, error)
  GetContactByCitizenId(id uint64) (*Contact, error)
}

type IContactService interface {
  ListAll() ([]*Contact, error)
  GetContactByCitizenId(id uint64) (*Contact, error)
}
