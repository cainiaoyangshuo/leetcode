<?php
/**
 * @date 2020/8/3
 */
//require_once './ListNode.php';
class LinkedList {

	public $head;
	public $size;

	public function __construct()
	{
		$this->head = new ListNode();
		$this->size  = 0;
	}

	public function addNodeFromHead($val)
	{
		$this->head = new ListNode($val, $this->head);
		$this->size++;
	}

	public function addNodeFromTail($val)
	{

	}


}