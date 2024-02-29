package domains

import "time"

type Users struct {
	UserID           uint      `gorm:"column:user_id;primaryKey" json:"userID"`
	Username         string    `gorm:"column:user_name;not null" json:"username"`
	Email            string    `gorm:"column:email;unique;not null" json:"email"`
	Password         string    `gorm:"column:password;not null" json:"-"`
	RegistrationDate time.Time `gorm:"column:registration_date;default:CURRENT_TIMESTAMP" json:"registrationDate"`
	PremiumStatus    bool      `gorm:"column:premium_status" json:"premiumStatus"`
	Otp              int       `gorm:"column:otp" json:"otp"`
	IsActive         bool      `gorm:"column:is_active" json:"isActive"`
}

func (Users) TableName() string {
	return "users"
}
