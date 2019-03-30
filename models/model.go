package models

type Model struct {
	Base

	Email       string `gorm:"unique_index"`
	Password    string
	IsActivated bool `gorm:"default:0"`
	IsVerified  bool `gorm:"default:0"`

	Name     *string `gorm:"type:varchar(50);unique_index"`
	Bio      *string `gorm:"type:text"`
	Avatar   *string `gorm:"type:text"`
	Video    *string `gorm:"type:text"`
	Pictures *string `gorm:"type:json"`
}
