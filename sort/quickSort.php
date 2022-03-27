<?php
/**
 * @date 2019/8/22
 */

class QuickSort
{
    public static function sort($arr)
    {
        $count = count($arr);

        if ($count < 2) {
            return $arr;
        }

        $pivot = $arr[0];
        $left = [];
        $right = [];

        for ($i = 1; $i < $count; $i++) {
            if ($pivot > $arr[$i]) {
                $left[] = $arr[$i];
            } else {
                $right[] = $arr[$i];
            }
        }

        $left = self::sort($left);
        $right = self::sort($right);

        return array_merge($left, [$pivot], $right);
    }
}

$arr = [];
for ($i = 0; $i < 10; $i++) {
    $num = rand(0, 100);
    echo $num . ' ';
    $arr[] = $num;
}
echo "\n";

$res = QuickSort::sort($arr);

foreach ($res as $re) {
    echo $re . ' ';
}
