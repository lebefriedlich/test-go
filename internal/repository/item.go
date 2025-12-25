package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// MasterCategoryMerchant represents a row from master_category_merchant.
type MasterCategoryMerchant struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
}

// MasterCategoryMerchantRepository handles queries for master_category_merchant data.
type MasterCategoryMerchantRepository struct {
	pool *pgxpool.Pool
}

func NewMasterCategoryMerchantRepository(pool *pgxpool.Pool) *MasterCategoryMerchantRepository {
	return &MasterCategoryMerchantRepository{pool: pool}
}

func (r *MasterCategoryMerchantRepository) List(ctx context.Context) ([]MasterCategoryMerchant, error) {
	const query = `
		SELECT
			id::text AS id,
			name,
			description
		FROM master_category_merchant
		ORDER BY name
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (MasterCategoryMerchant, error) {
		var rec MasterCategoryMerchant
		if err := row.Scan(
			&rec.ID,
			&rec.Name,
			&rec.Description,
		); err != nil {
			return MasterCategoryMerchant{}, err
		}
		return rec, nil
	})
	if err != nil {
		return nil, err
	}

	return records, nil
}