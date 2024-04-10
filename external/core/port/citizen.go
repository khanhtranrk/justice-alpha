package port

import "github.com/khanhtranrk/justice-alpha/external/core/domain"

type CitizenRepository interface {
  ListAllCitizens() ([]*domain.Citizen, error)
  GetCitizenById(id uint64) (*domain.Citizen, error)
}

type CitizenService interface {
  ListAllCitizens() ([]*domain.Citizen, error)
  GetCitizenById(id uint64) (*domain.Citizen, error)
}
