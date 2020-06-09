package models
import (
)

type InvoiceItem struct {
	InvoiceID uint64 `gorm:"not null"`
	Invoice Invoice `gorm:"foreignKey:InvoiceID; not null"`
	Description string `gorm:"not null"`
	Qty uint64 `gorm:"not null"`
	Unit_price float64 `gorm:"not null"`
	ModelBase
}

/*

        'invoice_id', 'description', 'qty', 'unit_price'
        'customer_id', 'title', 'date', 'due_date',
        'discount', 'sub_total', 'total'


        'company', 'email', 'name', 'phone', 'address'
type Auth struct {
	User User `gorm:"foreignKey:UserID; not null"`
	UserID int `json:"userid"`
	Token string `gorm:"type:varchar(200); not null"`
	ModelBase
}*/