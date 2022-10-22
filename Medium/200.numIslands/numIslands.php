<?php
/**
 * @file     numIslands.php
 * @brief    #功能说明
 * @version  1.0
 * @author   yangshuo
 * @since
 * @date     2022/3/28
 */
class Solution
{
    private $grid;
    /**
     * @param String[][] $grid
     * @return Integer
     */
    function numIslands($grid)
    {
        $this->grid = $grid;
        $num = 0;

        foreach ($this->grid as $key => $value) {
            foreach ($value as $k => $v) {
                if ($this->grid[$key][$k] == 1) {
                    $num++;
                    $this->dfs($key, $k);
                }
            }
        }

        return $num;
    }

    function dfs($key, $k)
    {
        if (!isset($this->grid[$key][$k]) || $this->grid[$key][$k] != 1) {
            return;
        }
        $this->grid[$key][$k] = 2;
        $this->dfs($key + 1, $k);
        $this->dfs($key - 1, $k);
        $this->dfs($key, $k + 1);
        $this->dfs($key, $k - 1);
    }
}
