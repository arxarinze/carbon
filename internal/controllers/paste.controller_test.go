package controllers

import (
	"carbon/internal/repo"
	"context"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func Test_pasteController_Create(t *testing.T) {
	type fields struct {
		ctx       context.Context
		pasteRepo repo.PasteRepo
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pasteController{
				ctx:       tt.fields.ctx,
				pasteRepo: tt.fields.pasteRepo,
			}
			if err := p.Create(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("pasteController.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
