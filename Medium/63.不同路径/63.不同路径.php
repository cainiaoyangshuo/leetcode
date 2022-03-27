<?php
/**
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
	网格中的障碍物和空位置分别用 1 和 0 来表示。

	说明：m 和 n 的值均不超过 100。

	示例 1:
	输入:
	[
	  [0,0,0],
	  [0,1,0],
	  [0,0,0]
	]
	输出: 2
	解释:
	3x3 网格的正中间有一个障碍物。
	从左上角到右下角一共有 2 条不同的路径：
	1. 向右 -> 向右 -> 向下 -> 向下
	2. 向下 -> 向下 -> 向右 -> 向右
 * @date 2020/7/6
 */

class Solution
{
    /**
     *
     * @param Integer[][] $obstacleGrid
     * @return Integer
     */
    function uniquePathsWithObstacles($obstacleGrid)
    {
        $n = count($obstacleGrid); // 列
        $m = count($obstacleGrid[0]); //行

        $f = [];

        $f[0] = $obstacleGrid[0][0] == 0 ? 1 : 0; // 判断开始是不是障碍

        for ($i = 0; $i < $n; $i++) {
            for ($j = 0; $j < $m; $j++) {
                if ($obstacleGrid[$i][$j] == 1) {
                    $f[$j] = 0;
                    continue;
                }

                if ($j - 1 >= 0 && $obstacleGrid[$i][$j - 1] == 0) {
                    $f[$j] += $f[$j - 1];
                }
            }
        }

        return $f[$m - 1];
    }
}

$obj = new Solution();

$obs = [[0, 0, 0], [0, 0, 1], [0, 0, 0]];
$res = $obj->uniquePathsWithObstacles($obs);
print_r($res);
