<?php
require __DIR__.'/vendor/autoload.php';

use \Grpc\ChannelCredentials;
use Pb\GreeterClient;
use Pb\HelloRequest;
use Pb\HelloReply;

// 创建客户端实例
$client = new \Pb\GreeterClient('127.0.0.1:8972',[
    'credentials'   => Grpc\ChannelCredentials::createInsecure()
]);

// 实例化请求参数
$helloRequest = new HelloRequest();
$helloRequest->setName("zhangfei");

// 调用远程服务
$get = $client->sayHello($helloRequest)->wait();

//返回数据
list($reply,$status) = $get;
var_dump($status);
var_dump($reply->getMessage());
