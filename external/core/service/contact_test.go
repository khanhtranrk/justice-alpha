package service

import (
	"testing"

	"github.com/khanhtranrk/justice-alpha/external/adapter/config"
	"github.com/khanhtranrk/justice-alpha/external/adapter/storage/sqlite"
	"github.com/khanhtranrk/justice-alpha/external/adapter/storage/sqlite/repository"
)


func TestListAllContacts(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  cntRepo := repository.NewContactRepository(db)
  cntServ := NewContactService(cntRepo)
  _, err = cntServ.ListAllContacts()
  if err != nil {
    t.Fatalf("ListAll failed: %v", err)
  }
}

func TestGetContactById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  cntRepo := repository.NewContactRepository(db)
  cntServ := NewContactService(cntRepo)
  _, err = cntServ.GetContactByCitizenId(4)
  if err != nil {
    t.Fatalf("GetContactByCitizenId failed: %v", err)
  }
}
