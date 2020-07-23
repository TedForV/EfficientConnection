package autoform

import (
	"testing"
)

var normalJson string = `
{
    "formName":"测试一",
    "formCode":"table_code_1",
    "displayText":"单元测试",
    "needFlow":false,
    "columns":[{
        "id":1,
        "name":"name",
        "DisplayText":"姓名",
        "type":{
            "id":1,
            "len":500
        }
    },{
        "id":2,
        "name":"sex",
        "DisplayText":"性别",
        "type":{
            "id":3,
            "len":0
        }
    }],
    "layout":{
        "column":2,
        "details":[
            [1,3],
            [2,4],
            [5,-1],
            [6,0]
        ]
    }
}
`

func Test_analyse(t *testing.T) {
	type args struct {
		formJSONStr string
	}
	tests := []struct {
		name    string
		args    args
		want    FormInfo
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				formJSONStr: normalJson,
			},
			want: FormInfo{
				FormName:    "测试一",
				FormCode:    "table_code_1",
				DisplayText: "单元测试",
				NeedFlow:    false,
				Columns: []FormColumn{
					FormColumn{
						ID:          1,
						Name:        "name",
						DisplayText: "姓名",
						ColumnType: FormColumnType{
							ID:  1,
							Len: 500,
						},
					},
					FormColumn{
						ID:          2,
						Name:        "sex",
						DisplayText: "性别",
						ColumnType: FormColumnType{
							ID:  3,
							Len: 0,
						},
					},
				},
				Layout: FormLayout{
					Column: 2,
					Details: [][]int{
						[]int{1, 3},
						[]int{2, 4},
						[]int{5, -1},
						[]int{6, 0},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := analyse(tt.args.formJSONStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("analyse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("analyse() = %v, want %v", got, tt.want)
			// }
		})
	}
}
