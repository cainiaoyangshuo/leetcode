<?php

$json = [
	'pre' => 'https://zyb-jiazhang.cdnjtzy.com/parentywremind/pre.mp3',
	'mid' => 'https://zyb-jiazhang.cdnjtzy.com/parentywremind/mid.mp3',
	'end' => 'https://zyb-jiazhang.cdnjtzy.com/parentyuwen/parentyuwen_2acbfef21e0ae2abe86b9bbbefa5b27c.mp3',
	];

//for ($i = 1; $i <= 50; $i++) {
//	$json[$i] = 'https://zyb-jiazhang.cdnjtzy.com/parentywremind/' . $i . '.mp3';
//}
const md5_prefix = 'parentyuwen';
const prefix = 'parentyuwen_';
//$path = '/Users/yangshuo/Downloads/口算听写10.9/parentyuwenbook1114的副本/';
$path = '/Users/yangshuo/Downloads/1124/';

$dir = dir($path);
$timer = 0;
$fileArr = [];
$charArr = [];

// 读取文件名
while (false !== ($entry = $dir->read())) {
	if ($entry == '.' || $entry == '..' || $entry == '.DS_Store') {
		continue;
	}
	$fileArr[] = $entry;
	$charId = substr($entry, 0, strpos($entry, '.'));
	$charArr[] = $charId;
}
var_dump($fileArr);
// 重命名文件
for ($i = 0; $i<count($fileArr); $i++) {
	$entry = $fileArr[$i];

	$char = substr($entry, 0, strpos($entry, '.'));
	$newName = prefix . md5(md5_prefix . $char) . '.mp3';
	echo $timer . '--' . $entry . '  ' . $newName . PHP_EOL;
	$res = rename($path.$entry, $path.$newName);
	if ($res == false) {
		exit('rename failed');
	}
	++$timer;
}

$dir->close();

//$url = [
//	[
//		'id' => 3,
//		'version' => 0,
//	],
//	[
//		'id' => 4,
//		'version' => 0,
//	],
//	[
//		'id' => 5,
//		'version' => 1,
//	],
//];
//
//$arr = array_column($url, 'version');
//$uniq = array_unique($arr);
sort($charArr);
print_r($charArr);
//$yinpinPrefix = 'https://zyb-jiazhang.cdnjtzy.com/parentyuwenbook1114/';
$yinpinPrefix = 'https://zyb-jiazhang.cdnjtzy.com/parentywpinyin/1124/';
$handle = fopen("char1124.sql", 'a');
// 生成sql
foreach ($charArr as $char) {
	$url = $yinpinPrefix . prefix . md5(md5_prefix . $char) . '.mp3';
	$sql = sprintf("UPDATE `tblYwChar` SET yuyin_version=1124,new_yuyin_url='%s' WHERE id=%d;\n", $url, $char);


	if (false === $handle) {
		exit('fopen error');
	}

	if (fwrite($handle, $sql) === false) {
		exit("写入文件char.sql失败");
	}
	++$timer;
}
fclose($handle);
echo $timer . PHP_EOL;
echo PHP_EOL . 'ok' . PHP_EOL;