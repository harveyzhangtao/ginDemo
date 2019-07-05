package models

type User struct {
	Model
	UserName     string `gorm:"type:varchar(50);not null"`
	AvatarUrl    string `gorm:"type:varchar(255);not null"`
	OpenId       string `gorm:"type:varchar(64);not null"`
	Gender       int `gorm:"type:TINYINT(1);not null"`
	UnionId      string `gorm:"type:varchar(64)"`
	Mobile       string `gorm:"type:varchar(11)"`
	City         string `gorm:"type:varchar(50)"`
	Province     string `gorm:"type:varchar(20)"`
	Country      string `gorm:"type:varchar(20)"`
	IsRobot      int `gorm:"type:TINYINT(1);not null;default(0)"`
	IsNew       int `gorm:"type:TINYINT(1);not null;default(0)"`
}
