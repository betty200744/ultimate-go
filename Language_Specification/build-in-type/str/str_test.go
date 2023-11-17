package str

import (
	"reflect"
	"testing"
)

func TestIdsString2Slice(t *testing.T) {
	type args struct {
		ids string
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "IdsString2Slice",
			args: args{ids: "1,2,3"},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdsString2Slice(tt.args.ids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String2SliceInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice2IdsString(t *testing.T) {
	type args struct {
		ids []int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "IdsString2Slice",
			args: args{
				ids: []int64{1, 2, 3},
			},
			want: "1,2,3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice2IdsString(tt.args.ids); got != tt.want {
				t.Errorf("Slice2IdsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64Encode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Base64Encode",
			args: args{s: "sss"},
			want: "c3Nz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Encode(tt.args.s); got != tt.want {
				t.Errorf("Base64Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64Decode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Base64Decode",
			args: args{
				s: "c3Nz",
			},
			want: "sss",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Decode(tt.args.s); got != tt.want {
				t.Errorf("Base64Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
