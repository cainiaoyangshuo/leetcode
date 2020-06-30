<?php
/**
 * 给定字符串haystack和一个要替换的字符串needle，在haystack中找到needle,替换为needle+自然数。在找到的第一个位置替换为“$needle1”,第二个位置替换为“$needle2”，依次类推，最终输出被替换后的haystack。
例如：
haystack："Hello everybody, I know English, I know Chinese"
needle："know"
replace："$needle$n"
替换后的结果：
"Hello everybody, I know1 English, I know2 Chinese"
 * @date 2020/4/29
 */


function replaceStr($haystack,$needle){
	//从这里开始写代码
	if (empty($haystack)) {
		return false;
	}
	$haystacks = explode(' ', $haystack);

	$counter = 1;
	foreach ($haystacks as $k => $h) {
		if ($h == $needle) {
			$haystacks[$k] = $h.$counter;
			++$counter;
		} else
			$haystacks[$k] = $h;
	}

	$haystack = implode(' ', $haystacks);
	return $haystack;

}

var_dump(replaceStr("Hello everybody, I know English, I know Chinese", 'know'));