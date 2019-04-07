package slice

import (
	"reflect"
	"testing"
)

func Test_SliceAtoi(t *testing.T) {
	type args struct {
		stringSlice []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "When supplied slice of string, should convert to slice of ints",
			args: args{
				stringSlice: []string{"1", "2", "3", "4"},
			},
			want:    []int{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name: "When supplied slice of non-convertable strings, should return nil and error",
			args: args{
				stringSlice: []string{"1", "2", "cabe"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "When supplied nil, should return nil and error",
			args: args{
				stringSlice: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SliceAtoi(tt.args.stringSlice)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceAtoi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SliceItoa(t *testing.T) {
	type args struct {
		intSlice []int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "When supplied slice of integer, should return slice of strings",
			args: args{
				intSlice: []int{1, 2, 3, 4},
			},
			want:    []string{"1", "2", "3", "4"},
			wantErr: false,
		},
		{
			name: "When supplied nil, should return nil and error",
			args: args{
				intSlice: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SliceItoa(tt.args.intSlice)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceItoa() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceItoa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ConvertIntSliceToCommaDelimitedString(t *testing.T) {
	type args struct {
		intSlice []int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "When supplied slice of ints, should convert to comma delimited string correctly",
			args: args{
				intSlice: []int{1, 2, 3, 4},
			},
			want:    "1,2,3,4",
			wantErr: false,
		},
		{
			name: "When supplied nil, should return empty string and error",
			args: args{
				intSlice: nil,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertIntSliceToCommaDelimitedString(tt.args.intSlice)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertIntSliceToCommaDelimitedString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertIntSliceToCommaDelimitedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ConvertStringSliceToCommaDelimitedString(t *testing.T) {
	type args struct {
		stringSlice []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "When supplied slice of strings, should convert to comma delimited string correctly",
			args: args{
				stringSlice: []string{"aku", "cinta", "tokopedia"},
			},
			want:    "aku,cinta,tokopedia",
			wantErr: false,
		},
		{
			name: "When supplied nil, should return empty string and error",
			args: args{
				stringSlice: nil,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertStringSliceToCommaDelimitedString(tt.args.stringSlice)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertStringSliceToCommaDelimitedString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertStringSliceToCommaDelimitedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
