package models

type Customer struct {
	Name string `gorm:"not null"`
	Company string `gorm:"not null"`
	phone string `gorm:"not null"`
	Address string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Invoice []Invoice `gorm:"foreignkey:UserRefer"`//has many invoices
	ModelBase
}

/*

        'id','company', 'email', 'name', 'phone', 'address', 'created_at'

        'company', 'email', 'name', 'phone', 'address'
type Auth struct {
	User User `gorm:"foreignKey:UserID; not null"`
	UserID int `json:"userid"`
	Token string `gorm:"type:varchar(200); not null"`
	ModelBase
}*/