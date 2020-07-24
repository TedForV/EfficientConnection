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

func TestConvertToMysqlColumnName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "CapitalTest",
			args: args{
				name: "UserInfo",
			},
			want:    "user_info",
			wantErr: false,
		},
		{
			name: "CapitalTest2",
			args: args{
				name: "userInfo",
			},
			want:    "user_info",
			wantErr: false,
		},
		{
			name: "UnderscoreTest",
			args: args{
				name: "user_info",
			},
			want:    "user_info",
			wantErr: false,
		},
		{
			name: "CapitalAndUnderscoreTest",
			args: args{
				name: "User_Info",
			},
			want:    "user_info",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToMysqlColumnName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToMysqlColumnName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToMysqlColumnName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToStructPropertyName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "FirstCapital",
			args: args{
				name: "user_info",
			},
			want:    "UserInfo",
			wantErr: false,
		},
		{
			name: "Capital",
			args: args{
				name: "user_Info",
			},
			want:    "UserInfo",
			wantErr: false,
		},
		{
			name: "None",
			args: args{
				name: "UserInfo",
			},
			want:    "UserInfo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToStructPropertyName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToStructPropertyName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToStructPropertyName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToJSONPropertyName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "FirstCapital",
			args: args{
				name: "user_info",
			},
			want:    "userInfo",
			wantErr: false,
		},
		{
			name: "Capital",
			args: args{
				name: "user_Info",
			},
			want:    "userInfo",
			wantErr: false,
		},
		{
			name: "None",
			args: args{
				name: "UserInfo",
			},
			want:    "userInfo",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToJSONPropertyName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToJSONPropertyName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToJSONPropertyName() = %v, want %v", got, tt.want)
			}
		})
	}
}
