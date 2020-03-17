package store

import (
	"testing"
	"time"
)

func TestCartAdd(t *testing.T) {
	type args struct {
		id      int32
		product Product
		count   int
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	tests = append(tests, struct {
		name    string
		args    args
		wantErr bool
	}{
		name: "test1", args: args{
			id: 32519,
			product: Product{
				ProductId: "1231",
				Name:      "wtfww",
				ImageUri:  []string{"http://wtf.img"},
				Information: Type{
					Category: "wtf",
					Brand:    "wtf",
				},
				Price:    999,
				Off:      25,
				Owner:    "",
				CreateAt: time.Now(),
			},
			count: 3,
		}, wantErr: false})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CartAdd(int32(tt.args.id), tt.args.product, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("CartAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCartDel(t *testing.T) {
	type args struct {
		id      int32
		product Product
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	tests = append(tests, struct {
		name    string
		args    args
		wantErr bool
	}{
		name: "test1", args: args{
			id: 32519,
			product: Product{
				ProductId: "123123123",
				Name:      "wtfww",
				ImageUri:  []string{"http://wtf.img"},
				Information: Type{
					Category: "wtf",
					Brand:    "wtf",
				},
				Price:    999,
				Off:      25,
				Owner:    "",
				CreateAt: time.Now(),
			},
		}, wantErr: false})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CartDel(tt.args.id, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("CartAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
