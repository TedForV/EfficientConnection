package autoform

import (
	"os"
	"testing"
)

var rootPath = "F:\\modProjects\\EfficientConnection\\formmodel\\"

var formInfoStr1 = `
{
    "formName":"测试一",
    "formCode":"code_1",
    "displayText":"单元测试",
    "needFlow":false,
    "columns":[{
        "id":1,
        "name":"age",
        "DisplayText":"年龄",
        "type":{
            "id":1,
            "len":0
        }
    },{
        "id":2,
        "name":"salary",
        "DisplayText":"薪水",
        "type":{
            "id":2,
            "len":0
        }
    },{
        "id":3,
        "name":"point",
        "DisplayText":"得分",
        "type":{
            "id":3,
            "len":0
        }
    },{
        "id":4,
        "name":"name",
        "DisplayText":"姓名",
        "type":{
            "id":4,
            "len":100
        }
    },{
        "id":5,
        "name":"is_student",
        "DisplayText":"是否学生",
        "type":{
            "id":5,
            "len":0
        }
    },{
        "id":6,
        "name":"attachment_id",
        "DisplayText":"附件",
        "type":{
            "id":6,
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

func Test_modelCreator_GenerateScript(t *testing.T) {
	form1, err := analyse(formInfoStr1)
	if err != nil {
		t.Error("json is wrong")
	}
	type args struct {
		form           FormInfo
		hasDecimalType bool
	}
	tests := []struct {
		name    string
		mc      *modelCreator
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "common type test",
			mc:   &modelCreator{},
			args: args{
				form:           form1,
				hasDecimalType: true,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.mc.GenerateScript(tt.args.form, tt.args.hasDecimalType)
			filePath := rootPath + tt.args.form.FormCode + ".go"
			os.Remove(filePath)
			f, err := os.Create(filePath)
			defer f.Close()
			f.WriteString(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("modelCreator.GenerateScript() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("modelCreator.GenerateScript() = %v, want %v", got, tt.want)
			// }
		})
	}
}
