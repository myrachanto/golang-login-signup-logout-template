package models
import (
	"time"
)

type Invoice struct {
	CustomerID uint64 `gorm:"not null"`
	Customer Customer `gorm:"foreignKey:CustomerID; not null"`
	Title string `gorm:"not null"`
	Dated time.Time `gorm:"not null"`
	Due_date time.Time `gorm:"not null"`
	Discount float64 `gorm:"not null"`
	Sub_total float64 `gorm:"not null"`
	Total float64 `gorm:"not null"`
	InvoiceItem []InvoiceItem `gorm:"foreignkey:UserRefer"`//has many invoiceitems
	ModelBase
}

/*

        'customer_id', 'title', 'date', 'due_date',
        'discount', 'sub_total', 'total'


        'company', 'email', 'name', 'phone', 'address'
type Auth struct {
	User User `gorm:"foreignKey:UserID; not null"`
	UserID int `json:"userid"`
	Token string `gorm:"type:varchar(200); not null"`
	ModelBase
}*/