package dto

type UpdatePassword struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassowrd string `json:"newPassword"`
}