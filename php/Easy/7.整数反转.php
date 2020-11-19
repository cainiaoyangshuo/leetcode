<?php
/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
 */

class Solution
{
    /**
     * @param Integer $x
     * @return Integer
     */
    function reverse($x)
    {
        if (!is_int($x) || abs($x) < 10) {
            return $x;
        }
        if ($x < 0) {
            $x = abs($x);
            $res = 0;
            while ($x > 9) {
                $remainder = $x % 10;
                $x = ($x - $remainder) / 10;
                $res = $res * 10 + $remainder;
            }
            $result = -($res * 10 + $x);
            if ($result <= -pow(2, 31) - 1) {
                return 0;
            }
        } else {
            $res = 0;
            while ($x > 9) {
                //var_dump('x:  ' . $x);
                $remainder = $x % 10;
                //var_dump('x的余数： ' . $remainder);
                $res = $res * 10 + $remainder;
                //var_dump('??:  ' . $res);
                $x = ($x - $remainder) / 10;
                //var_dump('去掉最后一位后的x:  ' . $x);
            }
            $result = $res * 10 + $x;
            if ($result >= pow(2, 31) - 1) {
                return 0;
            }
        }

        return $result;
    }
}

$obj = new Solution();
$date = $obj->reverse(19);
var_dump(crc32($date));

var_dump(urlencode('www.baidu.com?key=3'));
