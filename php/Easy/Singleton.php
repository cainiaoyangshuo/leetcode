<?php
/**
 * @date 2019/4/7
 * @author: yangshuo5@ucfgroup.com
 */

class Singleton
{
    private static $name;

    private function __construct()
    {
    }

    public static function getInstance()
    {
        //if (self::$name) {
        //    return self::$name;
        //} else {
        //    self::$name= new self();
        //    return self::$name;
        //}

        //self::$name = self::$name ? self::$name : new self();
        return self::$name ?: new self();
    }

    private function __clone()
    {
        // TODO: Implement __clone() method.
    }
}

$db = Singleton::getInstance();
$str = [1, 2, 3];
//foreach ($str as $key => $val) {
//    var_dump(pow(2,32));exit;
//    $str[$key] += $val % 2 ? $val++ : $val;
//}

$str[0] = $str[0] + ++$str[0];
var_dump($str[0]);
//$val = 0;

//var_dump(implode('', $str));
