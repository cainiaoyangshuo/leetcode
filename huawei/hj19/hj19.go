/**
 * @desc
 * @date 2022/11/9
 * @user yangshuo
 */

package main

import (
"bufio"
"fmt"
"log"
"os"
"strconv"
"strings"
)

func main() {
	simpleErrorRecord()
}

type errRecordPro struct {
	name        string
	lineCount   int
	recordCount int
}

const errRecordProArrLen int = 8

type errRecordProArr [errRecordProArrLen]errRecordPro

func simpleErrorRecord() {
	inLine := bufio.NewScanner(os.Stdin)
	var in []string
	for inLine.Scan() {
		//         if inLine.Text() == "stop" {
		//             break
		//         }
		in = append(in, inLine.Text())
	}
	var (
		errRecord errRecordProArr
		ma        map[string][]int
		position  = -1
	)
	ma = make(map[string][]int, 0)
	for _, v := range in {
		name, count := dealFileNameAndLineCount(v)
		mva, ok := ma[name]
		if !ok {
			ma[name] = append(ma[name], count)
		}
		if errRecord.isInRecord(name, count) {
			continue
		}
		if ok {
			var flag bool
			for _, v := range mva {
				if v == count {
					flag = true
					break
				}
			}
			if flag {
				continue
			}
			ma[name] = append(ma[name], count)
		}
		position = (position + 1) % errRecordProArrLen
		errRecord.insertRecord(name, count, position)
	}
	if position == -1 {
		return
	}
	errRecord.printResult(position)
}

// 处理文件名和行号
// 是否与错误记录中的某一条数据为同一个文件
// 记录错误记录
// 打印错误记录

func dealFileNameAndLineCount(in string) (string, int) {
	arr := strings.Split(in, " ")
	if len(arr) != 2 {
		log.Panic()
	}
	names, count := arr[0], arr[1]
	name := strings.Split(names, "\\")
	na := name[len(name)-1]
	if len(na) > 16 {
		na = na[len(na)-16:]
	}
	cou, _ := strconv.Atoi(count)
	return na, cou
}

func (er *errRecordProArr) isInRecord(name string, count int) bool {
	for i, _ := range er {
		if er[i].name == name && er[i].lineCount == count {
			er[i].recordCount++
			return true
		}
	}
	return false
}

func (er *errRecordProArr) insertRecord(name string, count, pos int) {
	er[pos].name = name
	er[pos].lineCount = count
	er[pos].recordCount = 1

}

func (er *errRecordProArr) printResult(start int) {
	var (
		end = (start + 1) % errRecordProArrLen
		pos = end
		res string
	)
	for {
		if er[pos].name == "" {
			pos = (pos + 1) % errRecordProArrLen
			continue
		}
		res += er[pos].name + " " + strconv.Itoa(er[pos].lineCount) + " " + strconv.Itoa(er[pos].recordCount) + "\n"
		pos = (pos + 1) % errRecordProArrLen
		if pos == end {
			break
		}
	}
	fmt.Print(res)
}