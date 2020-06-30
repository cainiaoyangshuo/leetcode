<?php
/**
 * @date 2019/10/24
 */


function my_scandir($dir)
{
	$files = array();
	if($handle = opendir($dir))
	{
		while (($file = readdir($handle))!== false)
		{
			if($file != '..' && $file != '.')
			{
				if(is_dir($dir."/".$file))
				{
					$files[$file]=my_scandir($dir. '/' .$file);
				}
				else
				{
					$files[] = $file;
				}
			}
		}

		closedir($handle);
		return $files;
	}

	return false;
}

var_dump(__FILE__);
var_dump(dirname(__FILE__));

