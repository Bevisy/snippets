package slice

import (
	"reflect"
	"testing"
)

func Test_unique(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "t01",
			args: args{s: []string{"abc", "cde", "efg", "efg", "abc", "cde"}},
			want: []string{"abc", "cde", "efg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
