<?php
class TreeNode
{
    public $left = null;
    public $right = null;
    public $value = null;
}

function reverseTree($root)
{
    if ($root === null) {
        return null;
    }

    $left = $root->left;
    $right = $root->right;
    $root->left = $this->reverseTree($right);
    $root->right = $this->reverseTree($left);
    return $root;
}

/**
 * @param TreeNode $root
 * @return TreeNode
 */
function invertTree($root)
{
    //左右数对调

    if ($root == null) {
        return null;
    }

    list($root->left, $root->right) = [
        $this->invertTree($root->right),
        $this->invertTree($root->left),
    ];

    return $root;
}

/**
 * 非递归算法
 * 这个方法的思路就是，我们需要交换树中所有节点的左孩子和右孩子。
 * 因此可以创一个队列来存储所有左孩子和右孩子还没有被交换过的节点。
 * 开始的时候，只有根节点在这个队列里面。
 * 只要这个队列不空，就一直从队列中出队节点，然后互换这个节点的左右孩子节点，接着再把孩子节点入队到队列，对于其中的空节点不需要加入队列。
 * 最终队列一定会空，这时候所有节点的孩子节点都被互换过了，直接返回最初的根节点就可以了。
 */
function invertTree1($root)
{
    if ($root === null) {
        return $root;
    }

    $queue = [];
    $queue[] = $root;

    while (count($queue) > 0) {
        $cur = array_shift($queue);
        $temp = $cur->left;
        $cur->left = $cur->right;
        $cur->right = $temp;

        if (!empty($cur->left)) {
            array_push($queue, $cur->left);
        }
        if (!empty($cur->right)) {
            array_push($queue, $cur->right);
        }
    }

    return $root;
}
