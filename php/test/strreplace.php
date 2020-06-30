<?php
/**
 * @date 2020/5/13
 */

function replace($target, $strs)
{
	if (empty($strs)) {
		return $strs;
	}
	$target = preg_replace("/(\r\n|\n|\r|\t)/i", '', $target);

	foreach ($strs as $str) {
		$target = str_replace($str, '', $target);
	}

	return $target;
}

