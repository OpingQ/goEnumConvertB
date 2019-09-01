package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type GoEnum struct {
	FileName   string
	PackStr    string
	MyEnumType string
	EnumSort   []string
}

const (
	ipackStr    = "_pkg_"
	imyEnumType = "_MyEnumType_"
	imead       = "_MyEnumArrDown_"
	imearr      = "_MyEnumArr_"
)

func ConvGo(fileName, packStr, myEnumType string, enumsort ...string) {
	bs, err := ioutil.ReadFile("enumFormat.txt")
	var str string
	if err == nil {
		str = string(bs)
		str = strings.ReplaceAll(str, ipackStr, packStr)
		str = strings.ReplaceAll(str, imyEnumType, myEnumType)

		for k, v := range enumsort {
			enumsort[k] = strings.Title(v)
		}
		newd := strings.Join(enumsort, "\n\t")
		na := make([]string, 0, 5)
		for _, v := range enumsort {
			na = append(na, fmt.Sprintf("%q", v))
		}
		newa := strings.Join(na, ",")

		str = strings.ReplaceAll(str, imead, newd)
		str = strings.ReplaceAll(str, imearr, newa)
		ioutil.WriteFile(fileName+".go", []byte(str), os.ModeIrregular)
		//fmt.Println(str)
	} else {

	}
}

type Product struct {
	Name  string
	Price int
}

func ConvJson() {
	ss := Product{
		Name:  "您好",
		Price: 456,
	}
	ssa := make(map[string]Product)
	ssa["1"] = ss
	ssa["2"] = ss
	ssa["3"] = ss
	ssa["4"] = ss

	b, _ := json.Marshal(ssa)
	fmt.Println(string(b))
}
