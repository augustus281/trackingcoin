package asset

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	db "github.com/augustus281/trackingcoin/database/sqlc"
	"github.com/augustus281/trackingcoin/global"
)

type IAssetRepository interface {
	FollowAsset(ctx *gin.Context, assetID, userID int) (db.UserFollowedAsset, error)
	UnfollowAsset(ctx *gin.Context, assetID, userID int) error
}

type assetRepo struct{}

func NewAssetRepo() IAssetRepository {
	return &assetRepo{}
}

func (r *assetRepo) FollowAsset(ctx *gin.Context, assetID, userID int) (db.UserFollowedAsset, error) {
	record, err := global.Db.CreateUserFollowedAsset(ctx, db.CreateUserFollowedAssetParams{
		UserID:  sql.NullInt32{Int32: int32(userID), Valid: userID != 0},
		AssetID: sql.NullInt32{Int32: int32(assetID), Valid: assetID != 0},
	})
	if err != nil {
		return db.UserFollowedAsset{}, err
	}
	return record, nil
}

func (r *assetRepo) UnfollowAsset(ctx *gin.Context, assetID, userID int) error {
	err := global.Db.DeleteUserFollowedAsset(ctx, db.DeleteUserFollowedAssetParams{
		UserID:  sql.NullInt32{Int32: int32(userID), Valid: userID != 0},
		AssetID: sql.NullInt32{Int32: int32(assetID), Valid: assetID != 0},
	})
	return err
}
