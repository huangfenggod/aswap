package utils

import (
	"encoding/json"
	"fmt"
)

func HandleJson()  {
	jsonData :=`{"Name":"lilei","Age":30,"list":1}`
	var value interface{}
	json.Unmarshal([]byte(jsonData),&value)
	data :=value.(map[string]interface{})

	for k,v:=range data{
		if k=="Name" {
			fmt.Println(v.(string))
		}
	}
}
