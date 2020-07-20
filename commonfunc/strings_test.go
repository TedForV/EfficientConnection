package commonfunc

import (
	"reflect"
	"testing"
)

func Test_containsKey(t *testing.T) {
	type args struct {
		source          []string
		caseInsensitive bool
		key             string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "oneInsensityTest_True",
			args: args{
				source:          []string{"a", "b", "c", "D", "e"},
				caseInsensitive: true,
				key:             "d",
			},
			want: true,
		},
		{
			name: "oneInsensityTest_False",
			args: args{
				source:          []string{"a", "b", "c", "D", "e"},
				caseInsensitive: true,
				key:             "f",
			},
			want: false,
		},
		{
			name: "oneSensitiveTest_True",
			args: args{
				source:          []string{"a", "b", "c", "D", "e"},
				caseInsensitive: true,
				key:             "e",
			},
			want: true,
		},
		{
			name: "oneSensitiveTest_False",
			args: args{
				source:          []string{"a", "b", "c", "D", "e"},
				caseInsensitive: false,
				key:             "d",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsKey(tt.args.source, tt.args.caseInsensitive, tt.args.key); got != tt.want {
				t.Errorf("containsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsKeys(t *testing.T) {
	type args struct {
		source          []string
		caseInsensitive bool
		keys            []string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []string
	}{
		{
			name: "InsensityTest_True",
			args: args{
				source:          []string{"a", "b", "c", "D", "e"},
				caseInsensitive: true,
				keys:            []string{"d", "A"},
			},
			want:  true,
			want1: []string{},
		},
		{
			name: "InsensityTest_False",
			args: args{
				source:          []string{"a", "b", "c", "D", "e"},
				caseInsensitive: true,
				keys:            []string{"f", "a"},
			},
			want:  false,
			want1: []string{"f"},
		},
		{
			name: "SensitiveTest_True",
			args: args{
				source:          []string{"a", "b", "c", "D", "e"},
				caseInsensitive: false,
				keys:            []string{"e", "D"},
			},
			want:  true,
			want1: []string{},
		},
		{
			name: "SensitiveTest_False",
			args: args{
				source:          []string{"a", "b", "c", "D", "e"},
				caseInsensitive: false,
				keys:            []string{"d", "E"},
			},
			want:  false,
			want1: []string{"d", "E"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := containsKeys(tt.args.source, tt.args.caseInsensitive, tt.args.keys)
			if got != tt.want {
				t.Errorf("containsKeys() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("containsKeys() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
