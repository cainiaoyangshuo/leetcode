<?php
/**
 * @date 2020/6/25
 */

class Solution
{
    /**
     * 异或运算 n^n=0 n^0=n
     * @param Integer[] $nums
     * @return Integer
     */
    function singleNumber($nums)
    {
        $len = count($nums);

        $res = 0;
        for ($i = 0; $i < $len; ++$i) {
            $res ^= $nums[$i];
        }

        return $res;
    }
}

$num = [];
var_dump((new Solution())->singleNumber($num));
