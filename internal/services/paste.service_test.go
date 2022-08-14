package services

import (
	"carbon/internal/helpers"
	"carbon/internal/interfaces"
	"carbon/internal/models"
	"carbon/internal/repo"
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_pasteService_Create(t *testing.T) {
	type fields struct {
		ctx       context.Context
		pasteRepo repo.PasteRepo
	}
	type args struct {
		payload interfaces.PasteRequest
		time    time.Time
	}
	ctrl := gomock.NewController(t)
	mPasteRepo := repo.NewMockPasteRepo(ctrl)
	now := time.Now()
	etime, _ := helpers.GenerateIfTime("2022-08-14")
	url := helpers.GenerateUrl("2022-08-14", "hello", now)
	mPasteRepo.EXPECT().CreatePaste(&models.Paste{
		ID:     0,
		Text:   "hello",
		Expiry: etime,
		Url:    url,
	}).Return(errors.New("err")).AnyTimes()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Create Post",
			fields: fields{
				ctx:       context.Background(),
				pasteRepo: mPasteRepo,
			},
			args: args{
				payload: interfaces.PasteRequest{
					Text:   "hello",
					Expiry: "2022-08-14",
				},
				time: now,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &pasteService{
				ctx:       tt.fields.ctx,
				pasteRepo: tt.fields.pasteRepo,
			}
			got, err := r.Create(tt.args.payload, tt.args.time)
			if (err != nil) != tt.wantErr {
				t.Errorf("pasteService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pasteService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pasteService_View(t *testing.T) {
	type fields struct {
		ctx       context.Context
		pasteRepo repo.PasteRepo
	}
	type args struct {
		url string
	}
	ctrl := gomock.NewController(t)
	mPasteRepo := repo.NewMockPasteRepo(ctrl)
	now := time.Now()
	url := helpers.GenerateUrl("2022-08-14", "hello", now)
	mPasteRepo.EXPECT().ViewPasteByUrl(url).Return(&models.Paste{
		ID:   0,
		Text: "asdasdsa",
	}, errors.New("err")).AnyTimes()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "View Post",
			fields: fields{
				ctx:       context.Background(),
				pasteRepo: mPasteRepo,
			},
			args: args{
				url: url,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pasteService{
				ctx:       tt.fields.ctx,
				pasteRepo: tt.fields.pasteRepo,
			}
			got, err := p.View(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("pasteService.View() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got, tt.want)
				t.Errorf("pasteService.View() = %v, want %v", got, tt.want)
			}
		})
	}
}
