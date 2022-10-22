<?php
/**
 * @date 2020/6/30
 */

class Solution
{
    /**
     * @desc  数组第k大的值 使用最小堆
     * @param $nums
     * @param $k
     *
     * @return mixed
     */
    function findKthLargest($nums, $k) {
        $heap = new SplMinHeap();
        $len = count($nums);
        
        for ($i = 0; $i < $len; $i++) {
            // 堆里最多有k个元素
            if ($heap->count() < $k) {
                $heap->insert($nums[$i]);
            } elseif ($nums[$i] > $heap->top()) {
                // 注意 只有最小堆元素到k个时才进行比较
                $heap->extract();
                $heap->insert($nums[$i]);
            }
        }
        
        return $heap->top();
    }
    
    function findMid($nums) {
        // 初始化最小堆
        $heap = new SplMinHeap();
        $len = count($nums);
        
        // 计算中位数数值
        if ($len % 2 == 0) {
            $k = $len/2;
        } else {
            $k = intval($len/2) + 1;
        }
    
        for ($i = 0; $i < $len; $i++) {
            // 堆里最多有k个元素
            if ($heap->count() < $k) {
                $heap->insert($nums[$i]);
            } elseif ($nums[$i] > $heap->top()) {
                // 注意 只有最小堆元素到k个时才进行比较
                $heap->extract();
                $heap->insert($nums[$i]);
            }
        }
    
        return $heap->top();
    }
}
