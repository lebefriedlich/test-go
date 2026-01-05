package repository

import (
	"context"

	"gorm.io/gorm"
)

// MasterCategoryMerchant represents a row from master_category_merchant.
type MasterCategoryMerchant struct {
	ID          string  `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	Name        string  `gorm:"column:name" json:"name"`
	Description *string `gorm:"column:description" json:"description,omitempty"`
}

// TableName sets the table name for GORM.
func (MasterCategoryMerchant) TableName() string {
	return "master_category_merchant"
}

// MasterCategoryMerchantRepository handles queries for master_category_merchant data.
type MasterCategoryMerchantRepository struct {
	db *gorm.DB
}

func NewMasterCategoryMerchantRepository(db *gorm.DB) *MasterCategoryMerchantRepository {
	return &MasterCategoryMerchantRepository{db: db}
}

func (r *MasterCategoryMerchantRepository) List(ctx context.Context) ([]MasterCategoryMerchant, error) {
	var records []MasterCategoryMerchant
	if err := r.db.WithContext(ctx).Order("name").Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}