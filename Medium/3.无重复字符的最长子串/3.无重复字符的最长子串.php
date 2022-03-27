<?php
/**
 * @date 2018/9/16
 * @author: yangshuo5@ucfgroup.com
 */

/**
 * Class lengthOfLongestSubstring
 *
 *
 *  给定一个字符串，找出不含有重复字符的最长子串的长度。
    示例 1:

    输入: "abcabcbb"
    输出: 3
    解释: 无重复字符的最长子串是 "abc"，其长度为 3。
    示例 2:

    输入: "bbbbb"
    输出: 1
    解释: 无重复字符的最长子串是 "b"，其长度为 1。
    示例 3:

    输入: "pwwkew"
    输出: 3
    解释: 无重复字符的最长子串是 "wke"，其长度为 3。
    请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
 */
class LengthOfLongestSubstring
{
    /**
     * 思路
     * 1 定义1个临时数组保存字符，1个记录最大长度的变量
     * 2 遍历目标字符串，将字符放入临时数组，并判断是否已经在临时数组
     * 如果不在则继续，如果存在则找到第一个所在位置，并将该字符后面所有字符作为新的临时数组
     * 3 判断临时数组长度，将大的保存在记录最大长度的变量中，以便和后面的比较
     * 时间复杂度：O(n) 字符串长度
     * 空间复杂度：字符集大小
     * @param $str
     *
     * @return int|mixed
     */
    public function theLengthOfLongestSubstring($str)
    {
        $maxlen = 0;
        $subStr = '';
        $len = strlen($str);
        for ($i = 0; $i < $len; $i++) {
            $inStr = strpos($subStr, $str[$i]); // 找到字符第一次出现的位置

            $subStr .= $str[$i]; // 不管之前有没有都放进去，有的话就从第一次出现的位置开始往后截取
            if ($inStr !== false) {
                $subStr = substr($subStr, $inStr + 1); // 如果有重复的，截取重复的字符后面的，并变为新的临时数组
            }

            $maxlen = max(strlen($subStr), $maxlen);
        }

        return $maxlen;
    }
}

$str = 'pwwkew';
$ob = new LengthOfLongestSubstring();
var_dump($ob->theLengthOfLongestSubstring($str));
