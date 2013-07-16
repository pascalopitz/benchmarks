<?php

$name = substr($_SERVER['REQUEST_URI'], 1);
$time = date('h:ma');
	
?>
Hello <?= $name ?>, it is <?= $time ?>