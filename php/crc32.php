<?php
/**
 * @file     crc32.php
 * @brief    #功能说明
 * @version  1.0
 * @author   yangshuo
 * @since
 * @date     2022/11/1
 */
class LeetCode_php_Crc32
{
    public function str($str = "")
    {
        return $str % 20;
    }

    public function crc32($str = "")
    {
        return crc32($str) % 20;
    }

    public function number($num = 0)
    {
        return $num % 20;
    }
}

$str = " ";
$num = 2100001220;
$a = new LeetCode_php_Crc32();
echo sprintf("%s %% 20 = %d\n", $str, $a->str($str));

echo sprintf("crc32(%s) %% 20 = %d\n", $str, $a->crc32($str));

echo sprintf("%d %% 20 = %d\n", $num, $a->number($str));
