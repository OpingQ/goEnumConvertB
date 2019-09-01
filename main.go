package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testlorca/ioTest/tool"
)

func main() {

	if bs, er := ioutil.ReadFile("enumArgs.txt"); er == nil {
		v := new(tool.GoEnum)
		json.Unmarshal(bs, v)
		tool.ConvGo(v.FileName, v.PackStr, v.MyEnumType, v.EnumSort...)
	} else {
		fmt.Println(er)
	}

	// g := new(tool.GoEnum)
	// g.PackStr = "bet"
	// g.MyEnumType = "BetZone"
	// g.EnumSort = []string{"top", "down", "right", "left"}
	// bs, _ := json.Marshal(g)
	// fmt.Println(string(bs))

	//tool.ConvGo("bet", "BetZone", "top", "down", "right", "left")
	// fmt.Println("--------------------------")
	// tool.ConvJson()
	//fmt.Println(tool.F3.ToString())
}
