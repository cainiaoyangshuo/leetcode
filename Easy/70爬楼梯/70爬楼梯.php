<?php
/**
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 * @date 2020/5/1
 */

/**
 * 斐波那契公式
 *
 * @param Integer $n
 * @return Integer
 */
function climbStairs($n)
{
    $sqrt5 = sqrt(5);
    $fibn = pow((1 + $sqrt5) / 2, $n + 1) - pow((1 - $sqrt5) / 2, $n + 1);
    return $fibn / $sqrt5;
}

function climbStairs1($n)
{
    if ($n == 1 || $n == 2) {
        return 1;
    }

    echo $n . ' ';
    return climbStairs1($n - 1) + climbStairs1($n - 2);
}

function climbStairs2($n)
{
    if ($n < 0) {
        return -1;
    }
    if ($n == 1 || $n == 2) {
        return 1;
    }

    $f1 = 1;
    $f2 = 1;
    for ($i = 3; $i <= $n; $i++) {
        $f3 = $f1 + $f2;
        $f1 = $f2;
        $f2 = $f3;
    }
    return $f3;
}

var_dump(climbStairs2(6));
