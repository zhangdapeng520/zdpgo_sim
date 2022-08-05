<?php
// 定义变量并默认设置为空值
$name111 = $email = $gender = $comment = $website = '';

if ($_SERVER['REQUEST_METHOD'] == 'POST') {
    // 这里的变量名不一样了
    $name111 = test_input($_POST['name']);
    $email = test_input($_POST['email']);
    $website = test_input($_POST['website']);
    $comment = test_input($_POST['comment']);
    $gender = test_input($_POST['gender']);
}

function test_input($data)
{
    $data = trim($data);
    $data = stripslashes($data);
    $data = htmlspecialchars($data);
    return $data;
}
?>
