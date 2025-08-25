package models

import (
	"time"
)

type Tenant struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Domain    string    `json:"domain" db:"domain"`
	Config    JSONB     `json:"config" db:"config"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type JSONB map[string]interface{}

type WhitelabelConfig struct {
	Branding struct {
		PrimaryColor   string `json:"primary_color"`
		SecondaryColor string `json:"secondary_color"`
		LogoURL        string `json:"logo_url"`
		CompanyName    string `json:"company_name"`
	} `json:"branding"`
	Features struct {
		EnableFeatureX bool `json:"enable_feature_x"`
		EnableFeatureY bool `json:"enable_feature_y"`
		MaxUsers       int  `json:"max_users"`
	} `json:"features"`
	API struct {
		RateLimit      int      `json:"rate_limit"`
		AllowedOrigins []string `json:"allowed_origins"`
	} `json:"api"`
}
