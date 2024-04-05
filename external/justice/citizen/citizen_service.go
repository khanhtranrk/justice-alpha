package justice

type CitizenService struct {
  rp ICitizenRepository
}

func NewCitizenService(rp ICitizenRepository) *CitizenService {
  return &CitizenService{rp}
}

func (cs *CitizenService) ListAll() ([]*Citizen, error) {
  return cs.rp.ListAll()
}

func (cs *CitizenService) GetCitizenById(id uint64) (*Citizen, error) {
  return cs.rp.GetCitizenById(id)
}
