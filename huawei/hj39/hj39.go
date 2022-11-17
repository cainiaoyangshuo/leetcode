/**
 * @desc 判断两个ip是否属于同一个子网
描述
IP地址是由4个0-255之间的整数构成的，用"."符号相连。
二进制的IP地址格式有32位，例如：10000011，01101011，00000011，00011000;每八位用十进制表示就是131.107.3.24
子网掩码是用来判断任意两台计算机的IP地址是否属于同一子网络的根据。
子网掩码与IP地址结构相同，是32位二进制数，由1和0组成，且1和0分别连续，其中网络号部分全为“1”和主机号部分全为“0”。
你可以简单的认为子网掩码是一串连续的1和一串连续的0拼接而成的32位二进制数，左边部分都是1，右边部分都是0。
利用子网掩码可以判断两台主机是否在同一子网中。
若两台主机的IP地址分别与它们的子网掩码进行逻辑“与”运算（按位与/AND）后的结果相同，则说明这两台主机在同一子网中。

示例：
I P 地址　 192.168.0.1
子网掩码　 255.255.255.0

转化为二进制进行运算：

I P 地址　  11000000.10101000.00000000.00000001
子网掩码　11111111.11111111.11111111.00000000

AND运算   11000000.10101000.00000000.00000000

转化为十进制后为：
192.168.0.0


I P 地址　 192.168.0.254
子网掩码　 255.255.255.0


转化为二进制进行运算：

I P 地址　11000000.10101000.00000000.11111110
子网掩码  11111111.11111111.11111111.00000000

AND运算  11000000.10101000.00000000.00000000

转化为十进制后为：
192.168.0.0

通过以上对两台计算机IP地址与子网掩码的AND运算后，我们可以看到它运算结果是一样的。均为192.168.0.0，所以这二台计算机可视为是同一子网络。

输入一个子网掩码以及两个ip地址，判断这两个ip地址是否是一个子网络。
若IP地址或子网掩码格式非法则输出1，若IP1与IP2属于同一子网络输出0，若IP1与IP2不属于同一子网络输出2。

注:
有效掩码与IP的性质为：
1. 掩码与IP每一段在 0 - 255 之间
2. 掩码的二进制字符串前缀为网络号，都由‘1’组成；后缀为主机号，都由'0'组成

输入描述：
3行输入，第1行是输入子网掩码、第2，3行是输入两个ip地址
题目的示例中给出了三组数据，但是在实际提交时，你的程序可以只处理一组数据（3行）。

输出描述：
若IP地址或子网掩码格式非法则输出1，若IP1与IP2属于同一子网络输出0，若IP1与IP2不属于同一子网络输出2

示例1
输入：
255.255.255.0
192.168.224.256
192.168.10.4
255.0.0.0
193.194.202.15
232.43.7.59
255.255.255.0
192.168.0.254
192.168.0.1
复制
输出：
1
2
0
复制
说明：
对于第一个例子:
255.255.255.0
192.168.224.256
192.168.10.4
其中IP:192.168.224.256不合法，输出1

对于第二个例子:
255.0.0.0
193.194.202.15
232.43.7.59
2个与运算之后，不在同一个子网，输出2

对于第三个例子，2个与运算之后，如题目描述所示，在同一个子网，输出0
 * @date 2022/11/17
 * @user yangshuo
 */


/*
思路：

判断ip和子网掩码的合法性。
然后通过&来获得子网掩码。
比较两个子网号是否一样。
由于Golang不像C和C++一样，能很好的操作底层，不能很好的将4个字节总体进行按位操作，因此把IP抽象一下，直接操作四个数，这样是比较简单。 另外，IP地址违法，这个题目也比较笼统，我随便输入一个字符串，我也可以说是违法的IP地址。考虑太多，这个题就不好做了。根据例子，非法的IP地址，一般就是字节不在0~255之间，子网掩码的合法性稍微复杂点。抽象成四个数后，每个数都是有限制的，当后面的数有值时，前面的数必须是255。最后面的一个数，必须是0xff,0xfe,0xfc,0xf8,0xf0,0xe0,0xc0,0x80,0x00中的某一个，这些数，是根据子网掩码的位限制得出来的。
 */
package main

import (
"fmt"
)

type ipaddr struct {
	byte1 int
	byte2 int
	byte3 int
	byte4 int
}

type info struct {
	mark    string
	ip1     string
	ip2     string
	netmark ipaddr
	ipaddr1 ipaddr
	ipaddr2 ipaddr
}

func main() {
	var in info

	for _, err := fmt.Scanf("%s\n", &in.mark); err == nil; _, err = fmt.Scanf("%s\n", &in.mark) {
		fmt.Scanf("%s\n", &in.ip1)
		fmt.Scanf("%s\n", &in.ip2)
		fmt.Sscanf(in.mark, "%d.%d.%d.%d", &in.netmark.byte1, &in.netmark.byte2, &in.netmark.byte3, &in.netmark.byte4)
		fmt.Sscanf(in.ip1, "%d.%d.%d.%d", &in.ipaddr1.byte1, &in.ipaddr1.byte2, &in.ipaddr1.byte3, &in.ipaddr1.byte4)
		fmt.Sscanf(in.ip2, "%d.%d.%d.%d", &in.ipaddr2.byte1, &in.ipaddr2.byte2, &in.ipaddr2.byte3, &in.ipaddr2.byte4)

		if !in.netmark.subnet_valid() || !in.ipaddr1.ip_valid() || !in.ipaddr2.ip_valid() {
			fmt.Println("1")
			continue
		}

		if in.same_network() {
			fmt.Println("0")
			continue
		}

		fmt.Println("2")
	}
}

func (ip *ipaddr) ip_valid() bool {
	r1 := (ip.byte1 <= 0 || ip.byte1 > 255)
	r2 := (ip.byte2 < 0 || ip.byte2 > 255)
	r3 := (ip.byte3 < 0 || ip.byte3 > 255)
	r4 := (ip.byte4 < 0 || ip.byte4 > 255)
	if r1 || r2 || r3 || r4 {
		return false
	}
	return true
}

func (ip *ipaddr) subnet_valid() bool {
	if ip.byte1 < ip.byte2 || ip.byte2 < ip.byte3 || ip.byte3 < ip.byte4 {
		return false
	}

	if ip.byte4 != 0 {
		if ip.byte3 != 0xff {
			return false
		}
	}

	if ip.byte3 != 0 {
		if ip.byte2 != 0xff {
			return false
		}
	}

	if ip.byte2 != 0 {
		if ip.byte1 != 0xff {
			return false
		}
	}

	if !ip_byte_check(ip.byte1) || !ip_byte_check(ip.byte2) || !ip_byte_check(ip.byte3) || !ip_byte_check(ip.byte4) {
		return false
	}

	return true
}

func ip_byte_check(b int) bool {
	if b == 0xff || b == 0xfe || b == 0xfc || b == 0xf8 || b == 0xf0 || b == 0xe0 || b == 0xc0 || b == 0x80 || b == 0x00 {
		return true
	}

	return false
}

func (ip *ipaddr) get_subnet(ip1, ip2 ipaddr) (ipaddr, ipaddr) {
	var net1 ipaddr
	var net2 ipaddr

	net1.byte1 = ip.byte1 & ip1.byte1
	net1.byte2 = ip.byte2 & ip1.byte2
	net1.byte3 = ip.byte3 & ip1.byte3
	net1.byte4 = ip.byte4 & ip1.byte4

	net2.byte1 = ip.byte1 & ip2.byte1
	net2.byte2 = ip.byte2 & ip2.byte2
	net2.byte3 = ip.byte3 & ip2.byte3
	net2.byte4 = ip.byte4 & ip2.byte4

	return net1, net2
}

func (i *info) same_network() bool {
	net1, net2 := i.netmark.get_subnet(i.ipaddr1, i.ipaddr2)

	if net1.byte1 == net2.byte1 && net1.byte2 == net2.byte2 && net1.byte3 == net2.byte3 && net1.byte4 == net2.byte4 {
		return true
	}

	return false
}

