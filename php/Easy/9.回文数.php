<?php
/**
 * 是否是回文
 * @date 2019/7/29
 */


function isPalindrome($x)
{

    if ($x < 0) {
        //echo __LINE__ . "\n";
        return false;
    }

    $len = strlen($x);

    if ($len%2 == 0) {
        // 偶数右边起始下标应该是长度
        $start = intval($len/2);
    } else {
        $start = intval($len/2) + 1;
    }

    $left = substr($x, 0, intval($len/2));

    $right = strrev(substr($x, $start, intval($len/2)));

    if ($left == $right) {
        return true;
    }

    return false;
}


var_dump(isPalindrome(10101));


/*
 * 没想到的点（用例）
 * 1.偶数也是回文的情况
 */