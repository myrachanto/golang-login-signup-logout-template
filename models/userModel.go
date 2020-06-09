package models

type User struct {
	Name string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	ModelBase
}
type Auth struct {
	User User `gorm:"foreignKey:UserID; not null"`
	UserID uint64 `json:"userid"`
	Token string `gorm:"type:varchar(200); not null"`
	ModelBase
}