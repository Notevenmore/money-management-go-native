package repositories

import (
	"database/sql"
	"money-management-be/models"
)

type AssetRepository struct {
	DB *sql.DB
}

func NewAssetRepositories(db *sql.DB) *AssetRepository {
	return &AssetRepository{DB: db}
}

func (r *AssetRepository) GetAll() ([]models.Assets, error) {
	rows, err := r.DB.Query("SELECT id, name, is_reusable FROM assets")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []models.Assets
	for rows.Next() {
		var asset models.Assets
		if err := rows.Scan(&asset.ID, &asset.Name, &asset.IsReusable); err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

func (r *AssetRepository) Create(asset models.Assets) error {
	_, err := r.DB.Exec("INSERT INTO assets (name, is_reusable) VALUES ($1, $2)", asset.Name, asset.IsReusable)
	return err
}
