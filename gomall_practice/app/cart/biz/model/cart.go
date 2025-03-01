package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null;"`
	Qty       uint32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, c *Cart) error {
	var row Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: c.UserId, ProductId: c.ProductId}).
			UpdateColumn("qty", gorm.Expr("qty+?", c.Qty)).Error
	}
	return db.WithContext(ctx).Create(c).Error
}

func EmptyCart(db *gorm.DB, ctx context.Context, userId uint32) error {
	if userId == 0 {
		return errors.New("user_is is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var rows []*Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).Where(&Cart{UserId: userId}).
		Find(&rows).Error
	return rows, err
}
