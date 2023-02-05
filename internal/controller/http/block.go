package http

import (
	"github.com/gofiber/fiber/v2"
	"pow-blockchain/internal/domain"
	"pow-blockchain/internal/domain/dto"
	"pow-blockchain/pkg/govalidator"
)

type BlockController struct {
	blockService domain.BlockService
	validator    govalidator.Validator
}

func (c *BlockController) GetBlockchain() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		result, err := c.blockService.GetBlockchain()
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.Status(fiber.StatusOK).JSON(result)
	}
}

func (c *BlockController) WriteBlock() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var p dto.WriteBlock
		if err := c.validator.ReadBody(ctx, &p); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		err := c.blockService.WriteBlock(p)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.SendStatus(fiber.StatusOK)
	}
}

func (c *BlockController) RegisterRoutes(blockGroup fiber.Router) {
	blockGroup.Get("", c.GetBlockchain())
	blockGroup.Post("", c.WriteBlock())
}

func NewBlockController(blockService domain.BlockService, validator govalidator.Validator) *BlockController {
	return &BlockController{blockService, validator}
}
