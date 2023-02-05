package govalidator

import (
	"context"
	gov "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator struct {
	val *gov.Validate
}

func (v *Validator) Validate(ctx context.Context, s any) error {
	return v.val.StructCtx(ctx, s)
}

func (v *Validator) ReadBody(c *fiber.Ctx, request interface{}) error {
	if err := c.BodyParser(request); err != nil {
		return err
	}

	return v.val.StructCtx(c.Context(), request)
}
func New() *Validator {
	return &Validator{gov.New()}
}
