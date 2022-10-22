<?php
/**
 * @file     quicksort.php
 * @brief    #功能说明
 * @version  1.0
 * @author   yangshuo
 * @since
 * @date     2022/10/21
 */

class Sort
{
    public function quickSort($nums)
    {
        return $this->dosort($nums, 0, count($nums)-1);
    }
    
    public function dosort($nums, $left, $right)
    {
        echo sprintf("left : %d, right : %d", $left, $right);
        
        if ($left > $right) {
            return false;
        }
        
        $i = $left;
        $j = $right;
        $pivot = $nums[$left];
        
        echo sprintf("i : %d, j : %d", $i, $j);
        echo sprintf("nums : %s", json_encode($nums));
        
        while ($i < $j) {
            //寻找小于主元的右边元素
            while ($i < $j && $nums[$j] >= $pivot) {
                $j--;
            }
            
            while($i < $j && $nums[$i] <= $pivot) {
                $i++;
            }
            //交换左边第一个小于主元的元素和右边第一个大于主元的元素
            $tmp = $nums[$i];
            $nums[$i] = $nums[$j];
            $nums[$j] = $tmp;
        }
    
        $tmp = $nums[$i];
        $nums[$i] = $nums[$left];
        $nums[$left] = $tmp;
        
        $this->dosort($nums, $left, $i-1);
        $this->dosort($nums, $i+1, $right);
        
        return $nums;
    }
    
    function swap(&$a, &$b)
    {
        $tmp = $a;
        $a = $b;
        $b = $tmp;
    }
}


$arr = [5, 2, 3, 9];
$obj = new Sort();

$a = 1;
$b = 3;

$obj->swap($a, $b);

$res = $obj->quickSort($arr);
print_r($res);