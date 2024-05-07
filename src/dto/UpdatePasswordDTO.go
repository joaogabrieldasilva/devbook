package dto

type UpdatePasswordDTO struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassowrd string `json:"newPassword"`
}