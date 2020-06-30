<?php
/**
 * @date 2019/9/5
 */


function binarySearch($arr, $target)
{
    $len = count($arr);
    $low = 0;
    $high = $len - 1;


    while ($low <= $high) {   //

        $mid = intval(($high + $low) / 2);   //
        echo $mid . ' ';

        if ($arr[$mid] > $target) {
            $high = $mid - 1;
        } elseif ($arr[$mid] < $target) {
            $low = $mid + 1;
        } else {
            return $mid;
        }
    }

    return -1;
}

$arr = [1, 2, 3, 4, 5, 6,7,8,9];

$res = binarySearch($arr, 2);

echo $res . "\n";