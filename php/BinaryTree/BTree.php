<?php
/**
 * @date 2020/5/5
 */

class BTree
{
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($value)
    {
        $this->val = $value;
    }
}

class Creator
{
    public $arr = [];

    public static function preOrder($root)
    {
        if (!empty($root)) {
            $func = __FUNCTION__;
            echo $root->val . '';
            $func($root->left);
            $func($root->right);
        }
    }

    public function inorder($root)
    {
        if (!empty($root)) {
            $func = __FUNCTION__;
            $func($root->left);
            echo $root->val . '';
            $func($root->right);
        }
    }

    public static function postOrder($root)
    {
        if (!empty($root)) {
            $func = __FUNCTION__;

            $func($root->left);
            $func($root->right);
            echo $root->val . '';
        }
    }

    public static function levelOrder($root)
    {
        self::order($root, 0);
        return self::$arr;
    }

    public function order($root, $level)
    {
        if (empty($root)) {
            return $root;
        }

        $this->arr[$level][] = $root->val;
        $level++;

        if (!empty($root->left)) {
            $this->order($root->left, $level);
        }

        if (!empty($root->right)) {
            $this->order($root->right, $level);
        }
    }

    /**
     * 层序遍历非递归
     * @param $root
     */
    public static function levelOrder1($root)
    {
        if (empty($root)) {
            return [];
        }

        $res = [];
        $queue = [];

        array_push($queue, $root);
        $level = 0;
        while ($count = count($queue)) {
            for ($i = 0; $i < $count; $i++) {
                $node = array_shift($queue);
                $res[$level][] = $node->val;

                if ($node->left) {
                    $queue[] = $node->left;
                }

                if ($node->right) {
                    $queue[] = $node->right;
                }
            }
            $level++;
        }

        return $res;
    }

    /** 先序遍历非递归 */
    public static function preOrder1($root)
    {
        if (empty($root)) {
            return [];
        }

        $res = [];
        $stack = [];
        array_push($stack, $root);

        while (count($stack)) {
            $node = array_pop($stack);
            $res[] = $node->val;
            if (!empty($node->right)) {
                array_push($stack, $node->right);
            }
            if (!empty($node->left)) {
                array_push($stack, $node->left);
            }
        }

        return $res;
    }

    /**
     * 中序遍历 迭代
     * @param $root
     *
     * @return array
     */
    public static function inorderTraversal($root)
    {
        if (empty($root)) {
            return [];
        }

        $res = [];
        $stack = [];
        while (!empty($root) || !empty($stack)) {
            if (!empty($root)) {
                array_push($stack, $root);
                $root = $root->left;
            } else {
                $root = array_pop($stack);
                $res[] = $root->val;
                $root = $root->right;
            }
        }

        return $res;
    }
}
