package dto

type (
	WriteBlock struct {
		Data int64 `json:"data" validate:"required"`
	}
)
