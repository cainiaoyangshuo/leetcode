<?php
/**
 * @date 2019/8/16
 */

class Food
{
    protected static $num = 0;

    public function __construct()
    {
        echo self::$num;
        ++self::$num;
    }
}


$a = $b = $c = new Food();

$a = [1,2,3,4];
//$b = [1,2,3];
//
//$res = array_diff($b, $a);
//var_dump($res);



$c = '2019-8-9';
$d = '2019-8-23';
$e = false;
$s = strpos($c, $d);
//var_dump(binarySearch($a, 3));
//print_r(dateDiff($c, $d));
$c = '不行abc3';
//var_dump(rev($c));
//var_dump(8%(-3));
function binarySearch($arr, $target)
{
    $len = count($arr);
    $low = 0;
    $high = $len-1;

    while ($low <= $high) {

        $mid = intval(($high + $low) / 2);

        if ($target < $arr[$mid]) {
            $high = $mid - 1;
        } elseif ($target > $arr[$mid]) {
            $low = $mid + 1;
        } else {
            return $mid;
        }
    }

    return -1;

}

/*
 * 计算时间差集
 */
function dateDiff($date1, $date2)
{
    return date_diff(date_create($date1), date_create($date2));
}


/*
 * 反转字符串包括中文
 */

function rev($str)
{
    if (!is_string($str)){
        throw new \Exception('输入错误');
    }

    // 判断有没有汉字  没有就用strrev
    if (!preg_match('/[\u4e00-\u9fa5]{0,}/', $str)) {
        return strrev($str);
    }

    $arr = [];

    $len = mb_strlen($str);

    for ($i=0; $i<$len; $i++) {
        array_unshift($arr, mb_substr($str, $i, 1));
    }

    return implode($arr);
}

function removeRepeat($str)
{
    $len = strlen($str);
    if ($len < 3) {
        return $str;
    }

    $res = preg_replace('/([a-zA-Z])\1{2,}/', '', $str);
    //$res = preg_match('/[a-zA-Z]{3,}/', $str);
//return $res;
    if (empty($res) || strcmp($res, $str) == 0) {
        return $str;
    }
    return removeRepeat($res);
}

var_dump(removeRepeat('abcccbbdd'));