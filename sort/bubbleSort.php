<?php
/**
 * @date 2019/9/7
 */

function bubbleSort($arr)
{
    $len = count($arr);

    for ($i = 0; $i < $len; $i++) {
        for ($j = 0; $j < $len - 1 - $i; $j++) {
            if ($arr[$j] > $arr[$j + 1]) {
                $temp = $arr[$j];
                $arr[$j] = $arr[$j + 1];
                $arr[$j + 1] = $temp;
            }
        }
    }

    return $arr;
}

for ($i = 0; $i < 10; $i++) {
    $arr[] = rand(0, 100);
}

var_dump(bubbleSort($arr));
