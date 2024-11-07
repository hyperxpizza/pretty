package pretty

import (
	"encoding/json"
	"fmt"
)

func Print(object any) {
	byteData, _ := json.MarshalIndent(object, "", " ")
	fmt.Println(string(byteData))
}
