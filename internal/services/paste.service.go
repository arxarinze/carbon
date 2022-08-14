package services

import (
	"carbon/internal/helpers"
	"carbon/internal/interfaces"
	"carbon/internal/models"
	"carbon/internal/repo"
	"context"
	"time"
)

type PasteService interface {
	Create(payload interfaces.PasteRequest, time time.Time) (string, error)
	View(usrl string) (map[string]string, error)
}

type pasteService struct {
	ctx       context.Context
	pasteRepo repo.PasteRepo
}

// View implements PasteService
func (p *pasteService) View(url string) (map[string]string, error) {
	view, err := p.pasteRepo.ViewPasteByUrl(url)
	if err != nil {
		return nil, err
	}
	stat := view.Expiry
	if !stat.Valid {
		return map[string]string{
			"text": view.Text,
		}, nil
	}
	checkTime := view.Expiry.Time
	if checkTime.Before(time.Now()) {
		return map[string]string{
			"error": "Expired",
		}, nil
	}
	return map[string]string{
		"text": view.Text,
	}, nil
}

func NewPasteService(ctx context.Context, repo repo.PasteRepo) PasteService {
	return &pasteService{
		ctx:       ctx,
		pasteRepo: repo,
	}
}

func (r *pasteService) Create(payload interfaces.PasteRequest, time time.Time) (string, error) {
	url := helpers.GenerateUrl(payload.Expiry, payload.Text, time)
	etime, err := helpers.GenerateIfTime(payload.Expiry)
	if err != nil {
		return "", err
	}
	err = r.pasteRepo.CreatePaste(&models.Paste{
		Url:    url,
		Text:   payload.Text,
		Expiry: etime,
	})
	if err != nil {
		return "", err
	}
	return url, nil
}
