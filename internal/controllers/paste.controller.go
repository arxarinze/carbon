package controllers

import (
	"carbon/internal/interfaces"
	"carbon/internal/models"
	"context"
	"time"

	"carbon/internal/helpers"
	"carbon/internal/repo"

	"github.com/gofiber/fiber/v2"
)

type PasteController interface {
	Create(c *fiber.Ctx) error
	View(c *fiber.Ctx) error
}

type pasteController struct {
	ctx       context.Context
	pasteRepo repo.PasteRepo
}

// Create implements PasteController
func (p *pasteController) Create(c *fiber.Ctx) error {
	payload := new(interfaces.PasteRequest)
	c.BodyParser(&payload)

	url := helpers.GenerateUrl(payload.Expiry, payload.Text)
	etime, err := helpers.GenerateIfTime(payload.Expiry)
	if err != nil {
		return c.Status(500).JSON(map[string]string{
			"error": "wrong time format use yyyy-mm-dd",
		})
	}
	err = p.pasteRepo.CreatePaste(&models.Paste{
		Url:    url,
		Text:   payload.Text,
		Expiry: etime,
	})
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
	view, err := p.pasteRepo.ViewPasteByUrl(url)
	if err != nil {
		return c.Status(404).JSON(map[string]string{
			"error": "Not Found",
		})
	}
	stat := view.Expiry
	if !stat.Valid {
		return c.JSON(map[string]string{
			"text": view.Text,
		})
	}
	checkTime := view.Expiry.Time
	if checkTime.Before(time.Now()) {
		return c.Status(404).JSON(map[string]string{
			"error": "Expired",
		})
	}
	return c.JSON(map[string]string{
		"text": view.Text,
	})
}

func NewPasteController(ctx context.Context, repo repo.PasteRepo) PasteController {
	return &pasteController{
		ctx,
		repo,
	}
}
