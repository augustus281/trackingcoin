package asset

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/augustus281/trackingcoin/global"
	"github.com/augustus281/trackingcoin/internal/dto"
	repository "github.com/augustus281/trackingcoin/internal/repository/asset"
)

type IAssetService interface {
	FollowAsset(ctx *gin.Context, userID, assetID int) (dto.UserFollowedAsset, error)
	UnfollowAsset(ctx *gin.Context, userID, assetID int) error
}

type assetService struct {
	repo repository.IAssetRepository
}

func NewAssetService(repo repository.IAssetRepository) IAssetService {
	return &assetService{
		repo: repo,
	}
}

func (s *assetService) FollowAsset(ctx *gin.Context, userID, assetID int) (dto.UserFollowedAsset, error) {
	record, err := s.repo.FollowAsset(ctx, assetID, userID)
	if err != nil {
		global.Logger.Error("follow asset failed ",
			zap.Error(err),
			zap.Int("userID", userID),
			zap.Int("assetID", assetID),
		)
		return dto.UserFollowedAsset{}, err
	}
	return dto.UserFollowedAsset{
		ID:         int(record.ID),
		UserID:     int(record.UserID.Int32),
		AssetID:    int(record.AssetID.Int32),
		FollowedAt: record.FollowedAt.Time,
	}, nil
}

func (s *assetService) UnfollowAsset(ctx *gin.Context, userID, assetID int) error {
	err := s.repo.UnfollowAsset(ctx, assetID, userID)
	if err != nil {
		global.Logger.Error("unfollow asset ", zap.Error(err))
		return err
	}
	return nil
}
