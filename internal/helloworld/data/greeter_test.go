package data

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"web_template/internal/helloworld/biz"
)

func Test_greeterRepo_FindByID(t *testing.T) {
	type fields struct {
		data *Data
		log  *log.Helper
	}
	type args struct {
		in0 context.Context
		in1 int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *biz.Greeter
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			fields:  fields{},
			args:    args{},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &greeterRepo{
				data: tt.fields.data,
				log:  tt.fields.log,
			}
			got, err := r.FindByID(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greeterRepo_ListAll(t *testing.T) {
	type fields struct {
		data *Data
		log  *log.Helper
	}
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*biz.Greeter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &greeterRepo{
				data: tt.fields.data,
				log:  tt.fields.log,
			}
			got, err := r.ListAll(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greeterRepo_ListByHello(t *testing.T) {
	type fields struct {
		data *Data
		log  *log.Helper
	}
	type args struct {
		in0 context.Context
		in1 string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*biz.Greeter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &greeterRepo{
				data: tt.fields.data,
				log:  tt.fields.log,
			}
			got, err := r.ListByHello(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListByHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListByHello() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greeterRepo_Save(t *testing.T) {
	type fields struct {
		data *Data
		log  *log.Helper
	}
	type args struct {
		ctx context.Context
		g   *biz.Greeter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *biz.Greeter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &greeterRepo{
				data: tt.fields.data,
				log:  tt.fields.log,
			}
			got, err := r.Save(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greeterRepo_Update(t *testing.T) {
	type fields struct {
		data *Data
		log  *log.Helper
	}
	type args struct {
		ctx context.Context
		g   *biz.Greeter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *biz.Greeter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &greeterRepo{
				data: tt.fields.data,
				log:  tt.fields.log,
			}
			got, err := r.Update(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
