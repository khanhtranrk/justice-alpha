package repository

import (
  "database/sql"
  "time"

  jc_domain "github.com/khanhtranrk/justice-alpha/justice_chain/core/domain"
)

type ChainRepository struct {
  db *sql.DB
}

func NewChainRepository(db *sql.DB) *ChainRepository {
  return &ChainRepository{db}
}

type Chain struct {
  PrimarySafeKey string
  PayloadKey string
  SecondarySafeKey string
  ChainKey string
  RelationshipKey string
  CreatedAt time.Time
}

func (r *ChainRepository) CreateChain(chain *jc_domain.Chain) (*jc_domain.Chain, error) {
  _, err := r.db.Exec(
    `INSERT INTO justice_chains (primary_safe_key, secondary_safe_key, payload_key, chain_key, relationship_key, created_at)
     VALUES (?, ?, ?, ?, ?, ?)`,
     chain.PrimarySafeKey,
     chain.SecondarySafeKey,
     chain.PayloadKey,
     chain.ChainKey,
     chain.RelationshipKey,
     chain.CreatedAt,
  )

  if err != nil {
    return chain, err
  }

  return chain, nil
}

