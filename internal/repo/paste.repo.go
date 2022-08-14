package repo

import (
	"carbon/internal/models"
	"context"

	"gorm.io/gorm"
)

type PasteRepo interface {
	CreatePaste(paste *models.Paste) error
	ViewPasteByUrl(url string) (*models.Paste, error)
}

type pasteRepo struct {
	ctx context.Context
	db  gorm.DB
}

// ViewPaste implements PasteRepo
func (p *pasteRepo) ViewPasteByUrl(url string) (*models.Paste, error) {
	result := new(models.Paste)
	tX := p.db.First(&result, &models.Paste{
		Url: url,
	})

	if tX.Error != nil {
		return &models.Paste{}, tX.Error
	}
	return result, nil
}

func NewPasteRepo(ctx context.Context, db gorm.DB) PasteRepo {
	return &pasteRepo{
		ctx,
		db,
	}
}

func (a *pasteRepo) CreatePaste(paste *models.Paste) error {
	tx := a.db.Create(&paste)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
