<?php
/**
 * @date 2019/9/5
 */

function binarySearch($arr, $target)
{
    $len = count($arr);
    $low = 0;
    $high = $len - 1;

    while ($low <= $high) {
        //

        $mid = intval(($high + $low) / 2); //
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

//$arr = [1, 2, 3, 4, 5, 6, 7, 8, 9];
//
//$res = binarySearch($arr, 2);
//
//echo $res . "\n";


function findMidAbs($nums) {
    if (empty($nums)) {
        return 0;
    }
    
    // 正数序列直接返回首元素
    if ($nums[0] >= 0) {
        return $nums[0];
    }
    
    // 负数序列返回尾元素
    $len = count($nums);
    if ($nums[$len-1] <= 0) {
        return $nums[$len-1];
    }
    
    // 通过二分查找 mid > 0：向左找；< 0 ：向右找
    $low = 0;
    $high = $len - 1;
    
    while ($low <= $high) {
        $mid = intval(($high + $low) / 2); //
        
        if ($nums[$mid] > 0) {
            $high = $mid - 1;
        } elseif ($nums[$mid] < 0) {
            $low = $mid + 1;
        } else {
            // 0的绝对值最小，直接返回
            return $nums[$mid];
        }
    }
    
    // 如果数组没有0，返回low左右的最小值
    return min(abs($nums[$low-1]), abs($nums[$low + 1]));
}

$arr = [-2, -1, 2, 3, 4, 5, 6, 7, 8, 9];

$res = findMidAbs($arr);
//exit;

echo $res . "\n";