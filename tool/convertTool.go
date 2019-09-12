package tool

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type GoEnum struct {
	FileName    string
	PackStr     string
	MyEnumType  string
	EnumSortNum []int
	EnumSort    []string
}

const (
	ipackStr    = "_pkg_"
	imyEnumType = "_MyEnumType_"
	imead       = "_MyEnumArrDown_"
	imearr      = "_MyEnumArr_"
)

// ConvGo 產生go檔案
func ConvGo(fileName, packStr, myEnumType string, enumSortNum []int, enumsort []string) error {
	bs, err := ioutil.ReadFile("enumFormat.txt")
	var str string
	if err == nil {
		fmt.Println(enumSortNum)
		convertMinusOne(enumSortNum)
		fmt.Println(enumSortNum)
		if err2 := checkOk(enumSortNum, enumsort); err2 != nil {
			return err2
		}
		fmt.Println(enumsort)
		fmt.Println("-------------------------")
		sortOrder(enumSortNum, enumsort)
		fmt.Println(enumSortNum)
		fmt.Println(enumsort)
		str = string(bs)
		str = strings.ReplaceAll(str, ipackStr, packStr)
		str = strings.ReplaceAll(str, imyEnumType, myEnumType)

		newd := ""
		astr := ""
		offset := -1
		for k, v := range enumsort {
			astr = ""
			if enumSortNum[k] != -1 && enumSortNum[k] != k+1 {
				curoffset := enumSortNum[k] - k - 1
				if offset != curoffset {
					offset = curoffset
					astr = fmt.Sprintf(" %s = %d + iota", myEnumType, offset)
				}
			}
			newd += strings.Title(v) + astr + "\n\t"
		}

		na := make([]string, 0, 5)
		for _, v := range enumsort {
			na = append(na, fmt.Sprintf("%q", v))
		}
		newa := strings.Join(na, ",")

		str = strings.ReplaceAll(str, imead, newd)
		str = strings.ReplaceAll(str, imearr, newa)
		createFileAndFolder(packStr, fileName, str)
		return nil
	}
	return err
}

func createFileAndFolder(packStr string, fileName string, data string) {
	folder := "./" + packStr
	if _, err := os.Stat(folder); err != nil {
		os.Mkdir(folder, os.ModePerm)
	}

	ioutil.WriteFile(folder+"/"+fileName+".go", []byte(data), os.ModeIrregular)
}

func checkOk(enumSortNum []int, enumsort []string) error {
	count := len(enumsort)
	ncount := len(enumSortNum)

	if ncount != count {
		return fmt.Errorf("num與enum數需相等 num: %d  enum: %d", ncount, count)
	}

	enumSortNumMap := make(map[int]int, 0)
	enumSortMap := make(map[string]int, 0)
	curSort := -1
	for _, v := range enumSortNum {
		if v == -1 {
			curSort--
			enumSortNumMap[curSort] = 0
			continue
		}
		enumSortNumMap[v] = 0
	}
	for _, v := range enumsort {
		enumSortMap[v] = 0
	}
	if len(enumSortNumMap) != count {
		return errors.New("Enum數字有重複")
	}
	if len(enumSortMap) != count {
		return errors.New("Enum文字有重複")
	}

	return nil
}

func convertMinusOne(enumSortNum []int) {
	begin := 0
	for k := range enumSortNum {
		if enumSortNum[k] == -1 {
			enumSortNum[k] = begin + 1
		}
		begin = enumSortNum[k]
	}
}

func sortOrder(enumSortNum []int, enumsort []string) {
	sort.SliceStable(enumSortNum, func(i, j int) bool {
		b := enumSortNum[i] < enumSortNum[j]
		if b {
			tstr := enumsort[i]
			enumsort[i] = enumsort[j]
			enumsort[j] = tstr
		}
		return b
	})
}
