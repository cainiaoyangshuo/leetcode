<?php
/**
 * @date 2020/5/29
 */

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */

function preorder($root)
{
    if (!empty($root)) {
        $func = __FUNCTION__;
        echo $root->val . '';
        $func($root->left);
        $func($root->right);
    }
}

/**
获取url扩展名
redis
ipc方式
数组交集
查找字符串最长公共前缀
给一组非负整数，重排最大整数
502 504
mysql调优
分表，
主从同步读写分离，
索引设计，根据查询条件设计合理索引
SQL尽量使用覆盖索引，索引不进行计算
根据慢查询日志解决慢查询问题。

项目架构
审核系统，前后端分离，分别部署在不同机器
首先通过消息队列接收上游提供的数据，消费入库
然后根据创建的任务不断地将记录分发给指定的审核组
mysql 一主一从部署，读写分离
数据源表按月分表


高并发架构
可以从几个层面入手，负载均衡、分布式缓存、灾备、熔断、降级
首先通过硬件或者软件进行负载均衡(F5或nginx)，分担机器压力
其次对于QPS过高的接口尽量走内存，而不要大量的io操作
而且还要有完善的监控机制，宕机或异常应该第一时间通知到负责人
机房放在不同的地方，两个以上
对于非核心业务要有降级机制，防止影响核心业务。

 **/
