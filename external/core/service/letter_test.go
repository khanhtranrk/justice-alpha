package service

import (
	"testing"
	"time"

	"github.com/khanhtranrk/justice-alpha/external/adapter/config"
	"github.com/khanhtranrk/justice-alpha/external/adapter/storage/sqlite"
	"github.com/khanhtranrk/justice-alpha/external/adapter/storage/sqlite/repository"
	"github.com/khanhtranrk/justice-alpha/external/core/domain"
)

func TestGetSentLetterById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  _, err = lttServ.GetSentLetterById(1)
  if err != nil {
    t.Fatalf("ListAll failed: %v", err)
  }
}

func TestCreateSentLetterById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 932749234,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateSentLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateSentLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateSentLetter failed: %v", err)
  }
}

func TestUpdateSentLetter(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 740580432,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateSentLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateSentLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateSentLetter failed: %v", err)
  }

  ltt.Status = 2
  ltt, err = lttServ.UpdateSentLetter(ltt)
  if err != nil {
    t.Fatalf("UpdateSentLetter failed: %v", err)
  }
}

func TestDeleteSentLetter(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 7534897094,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateSentLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateSentLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateSentLetter failed: %v", err)
  }

  ltt, err = lttServ.DeleteSentLetter(ltt)
  if err != nil {
    t.Fatalf("DeleteSentLetter failed: %v", err)
  }
}

func TestGetReceivedLetterById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  _, err = lttServ.GetReceivedLetterById(1)
  if err != nil {
    t.Fatalf("ListAll failed: %v", err)
  }
}

func TestCreateReceivedLetterById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 932749234,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateReceivedLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateReceivedLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateReceivedLetter failed: %v", err)
  }
}

func TestUpdateReceivedLetter(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 740580432,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateReceivedLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateReceivedLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateReceivedLetter failed: %v", err)
  }

  ltt.Status = 2
  ltt, err = lttServ.UpdateReceivedLetter(ltt)
  if err != nil {
    t.Fatalf("UpdateReceivedLetter failed: %v", err)
  }
}

func TestDeleteReceivedLetter(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 7534897094,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateReceivedLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateReceivedLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateReceivedLetter failed: %v", err)
  }

  ltt, err = lttServ.DeleteReceivedLetter(ltt)
  if err != nil {
    t.Fatalf("DeleteReceivedLetter failed: %v", err)
  }
}

func TestGetHandleLaterLetterById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  _, err = lttServ.GetHandleLaterLetterById(1)
  if err != nil {
    t.Fatalf("ListAll failed: %v", err)
  }
}

func TestCreateHandleLaterLetterById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 932749234,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateHandleLaterLetter(ltt)

  if err != nil {
    t.Fatalf("CreateHandleLaterLetter failed: %v", err)
  }

  if ltt.Id == 0 {
    t.Fatalf("CreateHandleLaterLetter failed: Id is zero")
  }

}

func TestUpdateHandleLaterLetter(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 740580432,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateHandleLaterLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateHandleLaterLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateHandleLaterLetter failed: %v", err)
  }

  ltt.Status = 2
  ltt, err = lttServ.UpdateHandleLaterLetter(ltt)
  if err != nil {
    t.Fatalf("UpdateHandleLaterLetter failed: %v", err)
  }
}

func TestDeleteHandleLaterLetter(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 7534897094,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateHandleLaterLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateHandleLaterLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateHandleLaterLetter failed: %v", err)
  }

  ltt, err = lttServ.DeleteHandleLaterLetter(ltt)
  if err != nil {
    t.Fatalf("DeleteHandleLaterLetter failed: %v", err)
  }
}

func TestGetResponseLaterLetterById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  _, err = lttServ.GetResponseLaterLetterById(1)
  if err != nil {
    t.Fatalf("ListAll failed: %v", err)
  }
}

func TestCreateResponseLaterLetterById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 932749234,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateResponseLaterLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateResponseLaterLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateResponseLaterLetter failed: %v", err)
  }
}

func TestUpdateResponseLaterLetter(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 740580432,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateResponseLaterLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateResponseLaterLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateResponseLaterLetter failed: %v", err)
  }

  ltt.Status = 2
  ltt, err = lttServ.UpdateResponseLaterLetter(ltt)
  if err != nil {
    t.Fatalf("UpdateResponseLaterLetter failed: %v", err)
  }
}

func TestDeleteResponseLaterLetter(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ltt := &domain.Letter{
    Id: 0,
    Code: 7534897094,
    ForeignId: 4,
    CommitTime: uint64(time.Now().UTC().UnixMilli()),
    Message: []byte{12,  13, 45, 53},
    Status: 1,
  }

  lttRepo := repository.NewLetterRepository(db)
  lttServ := NewLetterService(lttRepo)
  ltt, err = lttServ.CreateResponseLaterLetter(ltt)

  if ltt.Id == 0 {
    t.Fatalf("CreateResponseLaterLetter failed: Id is zero")
  }

  if err != nil {
    t.Fatalf("CreateResponseLaterLetter failed: %v", err)
  }

  ltt, err = lttServ.DeleteResponseLaterLetter(ltt)
  if err != nil {
    t.Fatalf("DeleteResponseLaterLetter failed: %v", err)
  }
}
