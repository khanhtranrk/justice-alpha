package port

import "github.com/khanhtranrk/justice-alpha/external/core/domain"

type ChainRepository interface {
  CreateChain(chain *domain.Chain) error
}

type ChainService interface {
  CreateChain(chain *domain.Chain) error
}
