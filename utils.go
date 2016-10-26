package mux

import (
	"fmt"
	// "encoding/json"
)

func prettyJson(data interface{}) {
	fmt.Printf("%+#v\n", data)
	/*
		d, _ := json.MarshalIndent(data, "", "	")
		fmt.Println(string(d))
	*/
}
