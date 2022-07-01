package biz

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/mock/gomock"
)

func TestGreeterUsecase_CreateGreeter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockGreeterRepo(ctrl)

	arg := &Greeter{
		Hello: "wiken",
	}
	repo.EXPECT().Save(context.Background(), arg).Return(arg, nil)

	l := log.GetLogger()
	uc := &GreeterUsecase{
		repo: repo,
		log:  log.NewHelper(l),
	}

	got, err := uc.CreateGreeter(context.Background(), arg)
	if err != nil {
		t.Errorf("CreateGreeter failed: %s", err.Error())
	}
	want := &Greeter{
		Hello: "wiken",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CreateGreeter got: %v, want: %v", got, want)
	}
}

func TestNewGreeterUsecase(t *testing.T) {
	type args struct {
		repo   GreeterRepo
		logger log.Logger
	}
	tests := []struct {
		name string
		args args
		want *GreeterUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGreeterUsecase(tt.args.repo, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGreeterUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
