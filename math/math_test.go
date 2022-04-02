package math

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "t01",
			args: args{num1: 2, num2: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.num1, tt.args.num2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		num1 int64
		num2 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "t01",
			args: args{num1: 2, num2: 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.num1, tt.args.num2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
