<?php
/**
 * author: yangshuo5@ucfgroup.com
 */

/**
 * $arr = [1,3,5,7]; $target = 4;
 * return index1 = 1, index2 = 2;
 * Class sulotion
 */
class TwoSum
{
    public $name = 'a';
    /**
     *
     *
     */
    public function index($arr, $target)
    {
        $sum = [];
        foreach ($arr as $k => $v) {
            foreach ($arr as $item => $value) {
                if ($v + $value == $target) {
                    $tmp[] = $k;
                    $tmp[] = $item;
                    $sum[] = $tmp;
                }
            }
        }
        return $sum;
    }

    /**
     * 哈希 存target-当前值的key-value 如果在哈希表中找到则说明有 时间复杂度O(n) 空间复杂度O(n)
     * @param $nums
     * @param $target
     *
     * @return array
     */
    public function twoSum1($nums, $target)
    {
        $m = [];
        $result = [];
        foreach ($nums as $k => $v) {
            if (array_key_exists($v, $m)) {
                $result[] = $m[$v];
                $result[] = $k + 1;

                return $result;
            }
            $m[$target - $v] = $k + 1;
            //var_dump($m);
        }
        //var_dump($m);
        return $result;
    }

    public function two_sum($nums, $target)
    {
        $map = []; //key为(target-当前数组的value)，value为key
        if (empty($nums) || count($nums) < 2) {
            return $nums;
        }
        $result = [];
        foreach ($nums as $key => $value) {
            if (array_key_exists($value, $map)) {
                array_push($result, $map[$value]);
                array_push($result, $key);
                return $result;
            }
            $map[$target - $value] = $key;
            /*
             * 这行不能放在if前面，[3,2,4]，6   这个测试用例会输出[0,0]
             */
        }

        return $result;
    }

    function twoSum2($arr, $target)
    {
        if (count($arr) < 2) {
            return $arr;
        }

        $map = $result = [];

        foreach ($arr as $item => $value) {
            if (array_key_exists($value, $map)) {
                array_push($result, $map[$value]);
                array_push($result, $item);
                return $result;
            }

            $map[$target - $value] = $item;
        }

        return $result;
    }

    function twoSum3($arr, $target)
    {
        if (count($arr) < 2) {
            return $arr;
        }

        $map = $res = [];

        foreach ($arr as $key => $value) {
            if (array_key_exists($value, $map)) {
                array_push($res, $map[$value]);
                array_push($res, $key);
                return $res;
            }

            $map[$target - $value] = $key;
        }

        return $res;
    }
    /*
     * 字符在字符串中的个数
     */
    public function check($str, $target)
    {
        return substr_count($str, $target);
    }
}

$sulo = new TwoSum();
$arr = [1, 3, 4, 2, 8];
$target = 7;
print_r($sulo->two_sum($arr, $target));
//var_dump($sulo->twoSum($arr,$target));
// var_dump($sulo->twoSum1($arr,$target));
//$monthlyMealFee = '32+36+30+35+30+25+35+38+43+24+22+33+34+27+24+20+35+31+24+31+30+41+17+13+22+18';
//$didi = '61+72+75+67+81+80+86+30+57+76+72+85+62+73+62+67+56+78+71+77';
//var_dump($sulo->check($didi, '+'));

$b = $sulo;
$b->name = 'b';
echo $sulo->name;
