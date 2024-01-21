// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID          int64     `json:"id"`
	AuthorName  string    `json:"author_name"`
	Website     string    `json:"website"`
	Instagram   string    `json:"instagram"`
	Youtube     string    `json:"youtube"`
	UserCreated string    `json:"user_created"`
	CreatedAt   time.Time `json:"created_at"`
}

type Recipe struct {
	ID           int64     `json:"id"`
	RecipeName   string    `json:"recipe_name"`
	Link         string    `json:"link"`
	AuthorID     int64     `json:"author_id"`
	PrepTime     float64   `json:"prep_time"`
	PrepTimeUnit string    `json:"prep_time_unit"`
	UserCreated  string    `json:"user_created"`
	CreatedAt    time.Time `json:"created_at"`
}

type RecipeIngredient struct {
	ID             int64  `json:"id"`
	Rank           int32  `json:"rank"`
	IngredientName string `json:"ingredient_name"`
	Unit           string `json:"unit"`
	// cannot be negative
	Amount    float64   `json:"amount"`
	RecipeID  int64     `json:"recipe_id"`
	CreatedAt time.Time `json:"created_at"`
}

type RecipeStep struct {
	ID              int64     `json:"id"`
	Rank            int32     `json:"rank"`
	StepDescription string    `json:"step_description"`
	RecipeID        int64     `json:"recipe_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	Email             string    `json:"email"`
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
