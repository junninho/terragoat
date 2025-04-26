<?php
// File inclusion vulnerability
$page = $_GET['page'];
include($page . '.php');
?> 