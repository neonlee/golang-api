package models

import (
	"time"
)

type Tenant struct {
	Id        int              `gorm:"column:id;primaryKey" json:"id"`
	Name      string           `gorm:"column:name" json:"name"`
	Domain    string           `gorm:"column:domain" json:"domain"`
	Config    WhitelabelConfig `gorm:"column:config" json:"config"`
	IsActive  bool             `gorm:"column:is_active" json:"is_active"`
	CreatedAt time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time        `gorm:"column:updated_at" json:"updated_at"`
}

type WhitelabelConfig struct {
	Branding struct {
		PrimaryColor   string `gorm:"column:primary_color" json:"primary_color"`
		SecondaryColor string `gorm:"column:secondary_color" json:"secondary_color"`
		LogoURL        string `gorm:"column:logo_url" json:"logo_url"`
		CompanyName    string `gorm:"column:company_name" json:"company_name"`
	} `json:"branding"`
	Features struct {
		EnableFeatureX bool `gorm:"column:enable_feature_x" json:"enable_feature_x"`
		EnableFeatureY bool `gorm:"column:enable_feature_y" json:"enable_feature_y"`
		MaxUsers       int  `gorm:"column:max_users" json:"max_users"`
	} `json:"features"`
	API struct {
		RateLimit      int      `gorm:"column:rate_limit" json:"rate_limit"`
		AllowedOrigins []string `gorm:"column:allowed_origins" json:"allowed_origins"`
	} `json:"api"`
}
