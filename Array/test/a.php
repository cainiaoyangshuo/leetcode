<?php
/**
 * @file     a.php
 * @brief    #功能说明
 * @version  1.0
 * @author   yangshuo
 * @since
 * @date     2022/9/28
 */
//$a1 = ['red', 'green'];
//$a2 = ['blue', 'green'];
//print_r(array_merge($a1, $a2));
//
//$b1 = ['a' => 'red', 'b' => 'green'];
//$b2 = ['c' => 'blue', 'b' => 'green'];
//print_r(array_merge($b1, $b2));
//
//echo date('Y-m-d', strtotime('-2 day', time()));
//
//
//$str1 = null;
//$str2 = false;
//echo $str1==$str2 ? '相等 ' : '不相等 ';
//$str3 = '';
//$str4 = 0;
//echo $str3==$str4 ? '相等 ' : '不相等 ';
//$str5 = 0;
//$str6 = '0';
//echo $str5===$str6 ? '相等' : '不相等';
//
//
//$a[0] = 1;
//$a[3] = 3;
//$a[5] = 5;
//echo count($a);



$s = '2020-05-16 19:20:34|user.login|name=Charles&location=Beijing&device=iPhone';

function detectLog($str) {
    // "|" 分隔字符串
    $arr = explode("|", $str);
    
    // 取数组最后
    $params = $arr[count($arr)-1];
    
    // "&"分割字符串
    $params = explode('&', $params);
    $res = [];
    foreach ($params as $k => &$v) {
        $kv = explode('=', $v);
        $res[$kv[0]] = $kv[1];
    }
    
    $res = json_encode($res);
    return $res;
}

$len = strlen("{\"scanImageUrl\":\"https://zyb-kunpeng-mkinteral-1253445850.cos.ap-beijing.myqcloud.com/exscanner%2Fscanner%2F1648008504497.jpg?q-sign-algorithm=sha1\u0026q-ak=AKIDrTqxpSqxqUt25av20mIftteCUXJp45vv\u0026q-sign-time=1648008509%3B3224808509\u0026q-key-time=1648008509%3B3224808509\u0026q-header-list=host\u0026q-url-param-list=\u0026q-signature=7dc19f06a87c6391583aad22445ab40d27988fef\",\"alignImageUrl\":\"https://zyb-kunpeng-mkinteral-1253445850.cos.ap-beijing.myqcloud.com/exscanner%2Falgo%2F1506482973648620672_align.jpg?q-sign-algorithm=sha1\u0026q-ak=AKIDrTqxpSqxqUt25av20mIftteCUXJp45vv\u0026q-sign-time=1648008509%3B3224808509\u0026q-key-time=1648008509%3B3224808509\u0026q-header-list=host\u0026q-url-param-list=\u0026q-signature=ae545af79db10ce7c4dd36d97e5d645fc4051ae3\",\"alignImage\":\"exscanner/algo/1506482973648620672_align.jpg\",\"uploadAlignImageUrl\":\"https://zyb-kunpeng-mkinteral-1253445850.cos.ap-beijing.myqcloud.com/exscanner%2Falgo%2F1506482973648620672_align.jpg?q-sign-algorithm=sha1\u0026q-ak=AKIDrTqxpSqxqUt25av20mIftteCUXJp45vv\u0026q-sign-time=1648008509%3B1648015709\u0026q-key-time=1648008509%3B1648015709\u0026q-header-list=host\u0026q-url-param-list=\u0026q-signature=5fcb0489becbda3a45ac9a710f4588112b5ba649\",\"sid\":\"1506482973648620672\",\"pages\":[{\"answerSheetPageNo\":2,\"templateImageUrl\":\"https://zyb-kunpeng-mkinteral-1253445850.cos.ap-beijing.myqcloud.com/exanno%2Ftplimg%2F1648006872295_109618_1.png?q-sign-algorithm=sha1\u0026q-ak=AKIDrTqxpSqxqUt25av20mIftteCUXJp45vv\u0026q-sign-time=1648008509%3B3224808509\u0026q-key-time=1648008509%3B3224808509\u0026q-header-list=host\u0026q-url-param-list=\u0026q-signature=93f839453be798ab15f04fc9f4aead079353b55e\",\"content\":{\"boxes\":[{\"id\":\"d3a2ba8ad366628d_1_tag\",\"type\":18,\"pos\":[2118,380,2216,445],\"optionValue\":\"1200681705\",\"filling\":0},{\"id\":\"fd15cf27d83a8267_1_tag\",\"type\":18,\"pos\":[2215,378,2313,443],\"optionValue\":\"1200673707\",\"filling\":0},{\"id\":\"2bc6c64080cbccd7_1_tag\",\"type\":18,\"pos\":[2015,1223,2115,1288],\"optionValue\":\"1200670139\",\"filling\":0},{\"id\":\"2bc6c64080cbccd7_2_tag\",\"type\":18,\"pos\":[2013,1878,2111,1946],\"optionValue\":\"1200670139\",\"filling\":0},{\"id\":\"a9a66f898b03c154_1_tag\",\"type\":18,\"pos\":[2118,1220,2203,1288],\"optionValue\":\"1200673687\",\"filling\":0},{\"id\":\"a9a66f898b03c154_2_tag\",\"type\":18,\"pos\":[2111,1873,2204,1951],\"optionValue\":\"1200673687\",\"filling\":0},{\"id\":\"de35febca7b3ae88_1_tag\",\"type\":18,\"pos\":[2206,1220,2306,1285],\"optionValue\":\"1200671171\",\"filling\":0},{\"id\":\"de35febca7b3ae88_2_tag\",\"type\":18,\"pos\":[2203,1875,2296,1948],\"optionValue\":\"1200671171\",\"filling\":0}],\"locationArea\":[{\"id\":\"pos-widget1648006820846-2\",\"type\":19,\"pos\":[25,3305,140,3415],\"optionValue\":\"\",\"filling\":0},{\"id\":\"pos-widget1648006820846-3\",\"type\":19,\"pos\":[53,45,173,158],\"optionValue\":\"\",\"filling\":0}],\"needResArea\":{\"studentNum\":{\"recogRes\":\"\",\"boxes\":[]}}}},{\"answerSheetPageNo\":1,\"templateImageUrl\":\"https://zyb-kunpeng-mkinteral-1253445850.cos.ap-beijing.myqcloud.com/exanno%2Ftplimg%2F1648006857065_109618_0.png?q-sign-algorithm=sha1\u0026q-ak=AKIDrTqxpSqxqUt25av20mIftteCUXJp45vv\u0026q-sign-time=1648008509%3B3224808509\u0026q-key-time=1648008509%3B3224808509\u0026q-header-list=host\u0026q-url-param-list=\u0026q-signature=44dc63f98c5fa7f9168bcf8c9bc2df7cb1763f76\",\"content\":{\"boxes\":[{\"id\":\"2a0624df20388d22-0-0-1\",\"type\":1,\"pos\":[264,1329,315,1357],\"optionValue\":\"A\",\"filling\":0},{\"id\":\"2a0624df20388d22-0-1-1\",\"type\":1,\"pos\":[348,1329,399,1357],\"optionValue\":\"B\",\"filling\":0},{\"id\":\"2a0624df20388d22-0-2-1\",\"type\":1,\"pos\":[432,1329,483,1357],\"optionValue\":\"C\",\"filling\":0},{\"id\":\"2a0624df20388d22-0-3-1\",\"type\":1,\"pos\":[516,1329,567,1357],\"optionValue\":\"D\",\"filling\":0},{\"id\":\"ed567da778b770a3-0-0-1\",\"type\":1,\"pos\":[264,1388,315,1416],\"optionValue\":\"A\",\"filling\":0},{\"id\":\"ed567da778b770a3-0-1-1\",\"type\":1,\"pos\":[348,1388,399,1416],\"optionValue\":\"B\",\"filling\":0},{\"id\":\"ed567da778b770a3-0-2-1\",\"type\":1,\"pos\":[432,1388,483,1416],\"optionValue\":\"C\",\"filling\":0},{\"id\":\"ed567da778b770a3-0-3-1\",\"type\":1,\"pos\":[516,1388,567,1416],\"optionValue\":\"D\",\"filling\":0},{\"id\":\"88037620d59f9e44-0-0-1\",\"type\":1,\"pos\":[264,1447,315,1475],\"optionValue\":\"A\",\"filling\":0},{\"id\":\"88037620d59f9e44-0-1-1\",\"type\":1,\"pos\":[348,1447,399,1475],\"optionValue\":\"B\",\"filling\":0},{\"id\":\"88037620d59f9e44-0-2-1\",\"type\":1,\"pos\":[432,1447,483,1475],\"optionValue\":\"C\",\"filling\":0},{\"id\":\"88037620d59f9e44-0-3-1\",\"type\":1,\"pos\":[516,1447,567,1475],\"optionValue\":\"D\",\"filling\":0},{\"id\":\"22d93110c943946f-0-0-1\",\"type\":1,\"pos\":[264,1506,315,1534],\"optionValue\":\"A\",\"filling\":0},{\"id\":\"22d93110c943946f-0-1-1\",\"type\":1,\"pos\":[348,1506,399,1534],\"optionValue\":\"B\",\"filling\":0},{\"id\":\"22d93110c943946f-0-2-1\",\"type\":1,\"pos\":[432,1506,483,1534],\"optionValue\":\"C\",\"filling\":0},{\"id\":\"22d93110c943946f-0-3-1\",\"type\":1,\"pos\":[516,1506,567,1534],\"optionValue\":\"D\",\"filling\":0},{\"id\":\"0ce7b07925da18d9-0-0-1\",\"type\":1,\"pos\":[264,1565,315,1593],\"optionValue\":\"A\",\"filling\":0},{\"id\":\"0ce7b07925da18d9-0-1-1\",\"type\":1,\"pos\":[348,1565,399,1593],\"optionValue\":\"B\",\"filling\":0},{\"id\":\"0ce7b07925da18d9-0-2-1\",\"type\":1,\"pos\":[432,1565,483,1593],\"optionValue\":\"C\",\"filling\":0},{\"id\":\"0ce7b07925da18d9-0-3-1\",\"type\":1,\"pos\":[516,1565,567,1593],\"optionValue\":\"D\",\"filling\":0}],\"locationArea\":[{\"id\":\"pos-widget1648006820846-0\",\"type\":19,\"pos\":[28,43,143,148],\"optionValue\":\"\",\"filling\":0},{\"id\":\"pos-widget1648006820846-1\",\"type\":19,\"pos\":[15,3293,135,3420],\"optionValue\":\"\",\"filling\":0}],\"needResArea\":{\"studentNum\":{\"recogRes\":\"\",\"boxes\":[{\"id\":\"\",\"type\":14,\"pos\":[1758,725,1806,755],\"optionValue\":\"0\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1758,773,1806,803],\"optionValue\":\"1\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1758,821,1806,851],\"optionValue\":\"2\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1758,869,1806,899],\"optionValue\":\"3\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1758,917,1806,947],\"optionValue\":\"4\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1758,965,1806,995],\"optionValue\":\"5\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1758,1013,1806,1043],\"optionValue\":\"6\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1758,1061,1806,1091],\"optionValue\":\"7\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1758,1109,1806,1139],\"optionValue\":\"8\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1758,1157,1806,1187],\"optionValue\":\"9\",\"groupIndex\":0},{\"id\":\"\",\"type\":14,\"pos\":[1840,725,1888,755],\"optionValue\":\"0\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1840,773,1888,803],\"optionValue\":\"1\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1840,821,1888,851],\"optionValue\":\"2\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1840,869,1888,899],\"optionValue\":\"3\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1840,917,1888,947],\"optionValue\":\"4\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1840,965,1888,995],\"optionValue\":\"5\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1840,1013,1888,1043],\"optionValue\":\"6\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1840,1061,1888,1091],\"optionValue\":\"7\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1840,1109,1888,1139],\"optionValue\":\"8\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1840,1157,1888,1187],\"optionValue\":\"9\",\"groupIndex\":1},{\"id\":\"\",\"type\":14,\"pos\":[1922,725,1970,755],\"optionValue\":\"0\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[1922,773,1970,803],\"optionValue\":\"1\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[1922,821,1970,851],\"optionValue\":\"2\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[1922,869,1970,899],\"optionValue\":\"3\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[1922,917,1970,947],\"optionValue\":\"4\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[1922,965,1970,995],\"optionValue\":\"5\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[1922,1013,1970,1043],\"optionValue\":\"6\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[1922,1061,1970,1091],\"optionValue\":\"7\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[1922,1109,1970,1139],\"optionValue\":\"8\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[1922,1157,1970,1187],\"optionValue\":\"9\",\"groupIndex\":2},{\"id\":\"\",\"type\":14,\"pos\":[2004,725,2052,755],\"optionValue\":\"0\",\"groupIndex\":3},{\"id\":\"\",\"type\":14,\"pos\":[2004,773,2052,803],\"optionValue\":\"1\",\"groupIndex\":3},{\"id\":\"\",\"type\":14,\"pos\":[2004,821,2052,851],\"optionValue\":\"2\",\"groupIndex\":3},{\"id\":\"\",\"type\":14,\"pos\":[2004,869,2052,899],\"optionValue\":\"3\",\"groupIndex\":3},{\"id\":\"\",\"type\":14,\"pos\":[2004,917,2052,947],\"optionValue\":\"4\",\"groupIndex\":3},{\"id\":\"\",\"type\":14,\"pos\":[2004,965,2052,995],\"optionValue\":\"5\",\"groupIndex\":3},{\"id\":\"\",\"type\":14,\"pos\":[2004,1013,2052,1043],\"optionValue\":\"6\",\"groupIndex\":3},{\"id\":\"\",\"type\":14,\"pos\":[2004,1061,2052,1091],\"optionValue\":\"7\",\"groupIndex\":3},{\"id\":\"\",\"type\":14,\"pos\":[2004,1109,2052,1139],\"optionValue\":\"8\",\"groupIndex\":3},{\"id\":\"\",\"type\":14,\"pos\":[2004,1157,2052,1187],\"optionValue\":\"9\",\"groupIndex\":3}]}}}}]}");
echo $len;
exit;