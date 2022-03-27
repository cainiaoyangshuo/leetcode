<?php
/**
 * @date 2020/7/14
 */

class Solution
{
    /**
     * 1 遇到左括号入栈
     * 2 非左括号则判断是否是栈顶左括号对应的右括号，是的话出栈，否则返回false
     * 3 遍历完字符串后判断栈是否为空，空则对
     * @param String $s
     * @return Boolean
     */
    function isValid($s)
    {
        $len = strlen($s);

        // 奇数直接返回
        if ($len % 2 == 1) {
            return false;
        }

        $map = [
            '(' => ')',
            '[' => ']',
            '{' => '}',
        ];

        $stack = new SplStack();

        for ($i = 0; $i < $len; $i++) {
            if (isset($map[$s[$i]])) {
                // 左括号入栈
                $stack->push($s[$i]);
            } else {
                // 非左括号则判断当前字符是否是栈顶左括号对应的右括号，是的话出栈，否返回false。因为只要不是左括号，如果栈空了那肯定错；如果栈不空说明有左括号，那如果不是栈顶左括号对应的右括号也不对。
                if (!$stack->isEmpty() && $map[$stack->top()] == $s[$i]) {
                    $stack->pop();
                } else {
                    return false;
                }
            }
        }

        return $stack->isEmpty();
    }
}

$re = new Solution();
$s = '([)]';
$res = $re->isValid($s);
var_dump($res);
