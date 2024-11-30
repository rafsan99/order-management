package models

import (
	"time"
)

type Order struct {
	ID                 uint      `gorm:"primaryKey"`
	OrderConsignmentID string    `gorm:"unique;not null" json:"order_consignment_id"`
	OrderDescription   string    `gorm:"" json:"order_description"`
	StoreID            uint      `gorm:"" json:"store_id"`
	MerchantOrderID    string    `gorm:"" json:"merchant_order_id"`
	RecipientName      string    `gorm:"not null" json:"recipient_name"`
	RecipientPhone     string    `gorm:"not null" json:"recipient_phone"`
	RecipientAddress   string    `gorm:"not null" json:"recipient_address"`
	RecipientCity      uint      `gorm:"not null" json:"recipient_city"`
	RecipientZone      uint      `gorm:"not null" json:"recipient_zone"`
	RecipientArea      uint      `gorm:"not null" json:"recipient_area"`
	DeliveryType       uint      `gorm:"not null" json:"delivery_type"`
	ItemType           uint      `gorm:"not null" json:"item_type"`
	ItemQuantity       uint      `gorm:"not null" json:"item_quantity"`
	ItemWeight         float64   `gorm:"not null" json:"item_weight"`
	SpecialInstruction string    `gorm:"" json:"instruction"`
	ItemDescription    string    `gorm:"" json:"item_description"`
	AmountToCollect    float64   `gorm:"not null" json:"order_amount"`
	CODFee             float64   `gorm:"not null" json:"cod_fee"`
	DeliveryFee        float64   `gorm:"not null" json:"delivery_fee"`
	PromoDiscount      float64   `gorm:"default:0" json:"promo_discount"`
	Discount           float64   `gorm:"default:0" json:"discount"`
	TotalFee           float64   `gorm:"not null" json:"total_fee"`
	OrderStatus        string    `gorm:"not null;default:'Pending'" json:"order_status"`
	OrderType          string    `gorm:"not null" json:"order_type"`
	UserID             uint      `gorm:"not null" json:"user_id"`
	OrderCreatedAt     time.Time `gorm:"autoCreateTime" json:"order_created_at"`
}
