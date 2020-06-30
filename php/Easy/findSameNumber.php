<?php
/**
 * 找出两个有序数组相同元素
 * 思路，两个数组同时前进，并比较大小，小了就往后走，直到任意一个数组走完。
 * @date 2019/10/22
 */


function one($arr1, $arr2)
{
	$i = $j = 0;

	$len1 = count($arr1);
	$len2 = count($arr2);

	while ($i < $len1 && $j < $len2) {
		if ($arr1[$i] == $arr2[$j]) {
			return $arr1[$i];
		}

		if ($arr1[$i] > $arr2[$j]) {
			$j++;
		} else {
			$i++;
		}
	}

	return -1;
}


