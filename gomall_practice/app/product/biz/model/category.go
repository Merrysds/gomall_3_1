package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `json:"product" gorm:"many2many:product_category"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) CategoryQuery {
	return CategoryQuery{ctx: ctx, db: db}
}

func (q CategoryQuery) GetProductsByCategoryName(name string) (category []Category, err error) {
	err = q.db.WithContext(q.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&category).Error
	return category, err
}
