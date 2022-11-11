/**
 * @desc https://www.nowcoder.com/practice/de538edd6f7e4bc3a5689723a7435682?tpId=37&tqId=21241&rp=1&ru=/exam/oj/ta&qru=/exam/oj/ta&sourceUrl=%2Fexam%2Foj%2Fta%3Fdifficulty%3D4%26page%3D1%26pageSize%3D50%26search%3D%26tpId%3D37%26type%3D37&difficulty=4&judgeStatus=undefined&tags=&title=
 * @date 2022/11/9
 * @user yangshuo
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// IP就是一个32位的数字

var (
	typeAStart = 1 << 24
	typeAEnd = 126 << 24 + 255 << 16 + 255 << 8 + 255
	typeBStart = 128 << 24
	typeBEnd = 191 << 24 + 255 << 16 + 255 << 8 + 255
	typeCStart = 192 << 24
	typeCEnd = 223 << 24 + 255 << 16 + 255 << 8 + 255
	typeDStart = 224 << 24
	typeDEnd = 239 << 24 + 255 << 16 + 255 << 8 + 255
	typeEStart = 240 << 24
	typeEEnd = 255 << 24 + 255 << 16 + 255 << 8 + 255

	typePrivate1Start = 10 << 24
	typePrivate1End = 10 << 24 + 255 << 16 + 255 << 8 + 255
	typePrivate2Start = 172 << 24 + 16 << 16
	typePrivate2End = 172 << 24 + 31 << 16 + 255 << 8 + 255
	typePrivate3Start = 192 << 24 + 168 << 16
	typePrivate3End = 192 << 24 + 168 << 16 + 255 << 8 + 255
)

func main() {
	var ip, mask string
	var lineParts []string
	var a, b, c, d, e, errCount, private int
	var ipNum, maskNum int
	var binMask string
	var max1Idx, min0Idx int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lineParts = strings.Split(scanner.Text(), "~")
		ip = lineParts[0]
		mask = lineParts[1]

		ipNum = ipStringToByteSlice(ip)
		maskNum = ipStringToByteSlice(mask)

		if ipNum >> 24 == 0 || ipNum >> 24 == 127 {
			continue // 类似于【0.*.*.*】和【127.*.*.*】的IP地址不属于上述输入的任意一类，也不属于不合法ip地址，计数时请忽略
		}

		if ipNum == -1 || maskNum == -1 {
			errCount++ // ip 或者mask 转数字失败
			continue
		}

		max1Idx = 0
		min0Idx = 0
		binMask = strconv.FormatInt(int64(maskNum), 2)
		for i:=0; i<len(binMask); i++ { // 找1最后位置 0 最早位置 0后面还有1的话，错误的mask
			if binMask[i] == '1' {
				max1Idx = i
			}

			if binMask[i] == '0' && min0Idx == 0 {
				min0Idx = i
			}
		}

		if max1Idx == len(binMask) - 1 || min0Idx == 0 || min0Idx <= max1Idx { // 全是1 或者 0出现在开头 或者 0后面还有1
			errCount++
			continue
		}

		// 以下开始 IP类型计数

		if (typePrivate1Start <= ipNum && ipNum <= typePrivate1End) ||
			(typePrivate2Start <= ipNum && ipNum <= typePrivate2End) ||
			(typePrivate3Start <= ipNum && ipNum <= typePrivate3End) {
			private++
		}

		if typeAStart <= ipNum && ipNum <= typeAEnd {
			a++
		} else if typeBStart <= ipNum && ipNum <= typeBEnd {
			b++
		} else if typeCStart <= ipNum && ipNum <= typeCEnd {
			c++
		} else if typeDStart <= ipNum && ipNum <= typeDEnd {
			d++
		} else if typeEStart <= ipNum && ipNum <= typeEEnd {
			e++
		}
	}

	fmt.Println(a, b, c, d, e, errCount, private)
}

func ipStringToByteSlice(ip string) int {
	s := strings.Split(ip, ".")
	if len(s) != 4 {
		return -1
	}

	r := make([]int, 4)
	var num int
	var err error
	for i:=0; i<4; i++ {
		num, err = strconv.Atoi(s[i])
		if err != nil {
			return -1
		}
		r[i] = num
	}

	return r[0] << 24 + r[1] << 16 + r[2] << 8 + r[3]
}

