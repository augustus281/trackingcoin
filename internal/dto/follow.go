package dto

import "time"

type UserFollowedAsset struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	AssetID    int       `json:"asset_id"`
	FollowedAt time.Time `json:"followed_at"`
}
