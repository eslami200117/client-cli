package entities
type AuthEntity struct {
	Username  string `gorm:"primaryKey" json:"username"`
	AuthToken string `json:"authToken"`
}
