<?php
/**
 * @file     aa.php
 * @brief    #功能说明
 * @version  1.0
 * @author   yangshuo
 * @since
 * @date     2022/10/11
 */

$params = [
    [
        "id" => 1,
        "parent_id" => 100,
        "name" => "dir1",
    ],
    [
        "id" => 2,
        "parent_id" => 100,
        "name"=> "dir",
    ],
    [
        "id" => 100,
        "parent_id" => 10,
        "name" => "dir100",
    ]
];

class aa {
    public $res = [];
    
    public function api($id, $params, $deep = 0)
    {
        $deep += 2;
        foreach ($params as $k => $v) {
            if ($v["parent_id"] == $id) {
                $item = [
                    "id" => $v["id"],
                    "name" => $v["name"],
                    "parent_id" => $v["parent_id"],
                    "children_id" => [],
                ];
                
                $this->res[] = "|".str_repeat("--", $deep). $v["name"];
                return $this->api($v["id"], $params, $deep);
            }
            
        }
        return $this->res;
    }
    
    function getList( $data, $pid=0, &$result=array(), $deep = 0 ) {
        $deep+=2;
        foreach ( $data as $key => $val ) {
            if ( $pid == $val['pid'] ) {
                $result[] = "|".str_repeat("--", $deep).$val['name'];
                $this->getList( $data, $val['uid'], $result, $deep );
            }
        }
        return $result;
    }
}

$a = new aa();
$res = $a->api(10, $params);
var_dump($res);