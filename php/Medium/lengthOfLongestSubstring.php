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
    public function theLengthOfLongestSubstring($str)
    {
        $maxlen = 0;
        $lastRepeatPos = -1;
        $map = [];
        for ($i = 0; $i < strlen($str); $i++) {
            $lastRepeatPos = max($lastRepeatPos, $map[$str[$i]]);
            print_r($lastRepeatPos);
            $map[$str[$i]] = $i;
            $maxlen = max($maxlen, $i - $lastRepeatPos);
            print_r($map);
        }
        var_dump($map);
        return $maxlen;
    }

}

$str = 'pwwkew';
$ob = new LengthOfLongestSubstring();
var_dump($ob->theLengthOfLongestSubstring($str));
