package service

import (
	"testing"

	"github.com/khanhtranrk/justice-alpha/external/adapter/config"
	"github.com/khanhtranrk/justice-alpha/external/adapter/storage/sqlite"
	"github.com/khanhtranrk/justice-alpha/external/adapter/storage/sqlite/repository"
)

func TestListRelationships(t *testing.T) {
  cfg, err := config.New()
  if err != nil {
    t.Fatalf("%v", err)
  }

  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("%v", err)
  }

  rp := repository.NewRelationshipRepository(db)
  sv := NewRelationshipService(rp)
  _, err = sv.ListRelationships()
  if err != nil {
    t.Fatalf("%v", err)
  }
}

func TestGetRelationshipByGuestKey(t *testing.T) {
  cfg, err := config.New()
  if err != nil {
    t.Fatalf("%v", err)
  }

  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("%v", err)
  }

  rp := repository.NewRelationshipRepository(db)
  sv := NewRelationshipService(rp)
  _, err = sv.GetRelationshipByGuestKey("test")
  if err != nil {
    t.Fatalf("%v", err)
  }
}

func TestIsTrustGuest(t *testing.T) {
  cfg, err := config.New()
  if err != nil {
    t.Fatalf("%v", err)
  }

  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("%v", err)
  }

  rp := repository.NewRelationshipRepository(db)
  sv := NewRelationshipService(rp)
  _, err = sv.IsTrustGuest("test", "test")
  if err != nil {
    t.Fatalf("%v", err)
  }
}

func TestUpdateSafeKeyForGuest(t *testing.T) {
  cfg, err := config.New()
  if err != nil {
    t.Fatalf("%v", err)
  }

  db, err := sqlite.New(cfg)
  if err != nil {
    t.Fatalf("%v", err)
  }

  rp := repository.NewRelationshipRepository(db)
  sv := NewRelationshipService(rp)
  err = sv.UpdateSafeKeyForGuest("test", "test")
  if err != nil {
    t.Fatalf("%v", err)
  }
}
