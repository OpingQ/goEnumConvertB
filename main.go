package main

import (
	"encoding/json"
	"fmt"
	"goEnumConvertB/tool"
	"io/ioutil"
)

func main() {

	// fmt.Println(bet.Right.Int())
	// fmt.Println(bet.Left.Int())
	//fmt.Println(bet.Jiod.FullString())
	// fmt.Println(Left)
	// fmt.Println(bet.Down)
	// fmt.Println(bet.Left)
	// fmt.Println(bet.Tuiosss)
	// fmt.Println(bet.Iodjvcsddf)

	if bs, er := ioutil.ReadFile("enumArgs.txt"); er == nil {
		v := new(tool.GoEnum)
		json.Unmarshal(bs, v)
		if err := tool.ConvGo(v.FileName, v.PackStr, v.MyEnumType, v.EnumSortNum, v.EnumSort); err != nil {
			fmt.Println(err)
		}
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
