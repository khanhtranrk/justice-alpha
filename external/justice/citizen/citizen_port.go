package justice

type ICitizenRepository interface {
  ListAll() ([]*Citizen, error)
  GetCitizenById(id uint64) (*Citizen, error)
}

type ICitizenService interface {
  ListAll() ([]*Citizen, error)
  GetCitizenById(id uint64) (*Citizen, error)
}
