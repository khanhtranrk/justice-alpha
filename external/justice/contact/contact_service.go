package justice

type ContactService struct {
  rp IContactRepository
}

func NewContactService(rp IContactRepository) *ContactService {
  return &ContactService{rp}
}

func (cs *ContactService) ListAll() ([]*Contact, error) {
  return cs.rp.ListAll()
}

func (cs *ContactService) GetContactByCitizenId(citizenid uint64) (*Contact, error) {
  return cs.rp.GetContactByCitizenId(citizenid)
}
