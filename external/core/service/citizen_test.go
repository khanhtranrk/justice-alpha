package service

import (
	"testing"

	"github.com/khanhtranrk/justice-alpha/external/adapter/config"
	"github.com/khanhtranrk/justice-alpha/external/adapter/storage/sqlite"
	"github.com/khanhtranrk/justice-alpha/external/adapter/storage/sqlite/repository"
)


func TestListAllCitizens(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ctzRepo := repository.NewCitizenRepository(db)
  ctzServ := NewCitizenService(ctzRepo)
  _, err = ctzServ.rp.ListAllCitizens()
  if err != nil {
    t.Fatalf("ListAll failed: %v", err)
  }
}

func TestGetCitizenById(t *testing.T) {
  cfg := &config.DB{DatabaseURL: "justice_test.db"}
  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("Could not connect to database: %v", err)
  }

  ctzRepo := repository.NewCitizenRepository(db)
  ctzServ := NewCitizenService(ctzRepo)
  _, err = ctzServ.GetCitizenById(4)
  if err != nil {
    t.Fatalf("ListAll failed: %v", err)
  }
}
