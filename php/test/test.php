<?php
/**
 * @date 2019/10/8
 */



function genN($n)
{
	function genT($len)
	{
		for ($i = 1; $i < $len; $i++) {
			yield $i;

		}
	}
	foreach (genT($n) as $out) {
		printf("%d", $out);
	}

	printf("%d", ++$out);
}


//genN(15);

// 上个月1号
$b = 5;
$a = &$b;
$b = null;
sleep(6);
var_dump(empty($b));
