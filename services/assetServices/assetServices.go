package assetservices

import (
	"money-management-be/models"
	"money-management-be/repositories"
)

type AssetService struct {
	Repo *repositories.AssetRepository
}

func NewAssetService(repo *repositories.AssetRepository) *AssetService {
	return &AssetService{Repo: repo}
}

func (s *AssetService) GetAllAssets() ([]models.Assets, error) {
	assets, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (s *AssetService) CreateAsset(asset models.Assets) error {
	if asset.Name == "" {

	}
	return s.Repo.Create(asset)
}
