package controllers

import (
	"carbon/internal/interfaces"
	"carbon/internal/repo"
	"carbon/internal/services"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type PasteController interface {
	Create(c *fiber.Ctx) error
	View(c *fiber.Ctx) error
}

type pasteController struct {
	ctx          context.Context
	pasteRepo    repo.PasteRepo
	pasteService services.PasteService
}

// Create implements PasteController
func (p *pasteController) Create(c *fiber.Ctx) error {
	payload := new(interfaces.PasteRequest)
	c.BodyParser(&payload)

	url, err := p.pasteService.Create(*payload, time.Now())
	if err != nil {
		return c.Status(500).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(interfaces.PasteResponse{
		Url: c.BaseURL() + "/" + url,
	})
}

// View implements PasteController
func (p *pasteController) View(c *fiber.Ctx) error {
	url := c.Params("url")
	data, err := p.pasteService.View(url)
	if err != nil {
		return c.Status(500).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(data)
}

func NewPasteController(ctx context.Context, repo repo.PasteRepo) PasteController {
	service := services.NewPasteService(ctx, repo)
	return &pasteController{
		ctx,
		repo,
		service,
	}
}
