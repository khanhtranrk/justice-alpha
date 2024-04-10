package port

import "github.com/khanhtranrk/justice-alpha/external/core/domain"

type LetterRepository interface {
  CreateLetter(letter *domain.Letter, tableName string) (*domain.Letter, error)
  UpdateLetter(letter *domain.Letter, tableName string) (*domain.Letter, error)
  DeleteLetter(letter *domain.Letter, tableName string) (*domain.Letter, error)
  GetLetterById(id uint64, tableName string) (*domain.Letter, error)
}

type LetterService interface {
  CreateHandleLaterLetter(letter *domain.Letter) (*domain.Letter, error)
  CreateResponseLaterLetter(letter *domain.Letter) (*domain.Letter, error)
  CreateSentLetter(letter *domain.Letter) (*domain.Letter, error)
  CreateReceivedLetter(letter *domain.Letter) (*domain.Letter, error)

  DeleteHandleLaterLetter(letter *domain.Letter) (*domain.Letter, error)
  DeleteResponseLaterLetter(letter *domain.Letter) (*domain.Letter, error)
  DeleteSentLetter(letter *domain.Letter) (*domain.Letter, error)
  DeleteReceivedLetter(letter *domain.Letter) (*domain.Letter, error)

  UpdateHandleLaterLetter(letter *domain.Letter) (*domain.Letter, error)
  UpdateResponseLaterLetter(letter *domain.Letter) (*domain.Letter, error)
  UpdateSentLetter(letter *domain.Letter) (*domain.Letter, error)
  UpdateReceivedLetter(letter *domain.Letter) (*domain.Letter, error)

  GetHandleLaterLetterById(id string) (*domain.Letter, error)
  GetResponseLaterLetterById(id string) (*domain.Letter, error)
  GetSentLetterById(id string) (*domain.Letter, error)
  GetReceivedLetterById(id string) (*domain.Letter, error)
}

