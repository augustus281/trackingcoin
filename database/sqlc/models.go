// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type NotificationType string

const (
	NotificationTypeEmail    NotificationType = "email"
	NotificationTypeDiscord  NotificationType = "discord"
	NotificationTypeTelegram NotificationType = "telegram"
)

func (e *NotificationType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NotificationType(s)
	case string:
		*e = NotificationType(s)
	default:
		return fmt.Errorf("unsupported scan type for NotificationType: %T", src)
	}
	return nil
}

type NullNotificationType struct {
	NotificationType NotificationType
	Valid            bool // Valid is true if NotificationType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullNotificationType) Scan(value interface{}) error {
	if value == nil {
		ns.NotificationType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.NotificationType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullNotificationType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.NotificationType), nil
}

type Asset struct {
	ID                int32
	CmcID             int64
	Name              string
	Slug              string
	Price             float64
	PercentChange1h   float32
	PercentChange24h  float32
	PercentChange7d   float32
	MarketCap         float64
	Volume24h         float64
	CirculatingSupply float32
	AllTimeHigh       float32
	AllTimeLow        float32
	Turnover          float32
	TotalSupply       float32
	MaxSupply         float32
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
}

type Notification struct {
	ID        int32
	UserID    sql.NullInt32
	NotiType  NotificationType
	Message   sql.NullString
	IsRead    sql.NullBool
	CreatedAt sql.NullTime
}

type User struct {
	ID             int32
	Email          string
	Fullname       string
	HashedPassword string
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
}

type UserFollowedAsset struct {
	ID         int32
	UserID     sql.NullInt32
	AssetID    sql.NullInt32
	FollowedAt sql.NullTime
}
