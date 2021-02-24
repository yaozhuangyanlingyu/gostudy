<?php
require('./vendor/autoload.php');
function p($arr) {
    echo '<pre>';
    print_r($arr);
    die;
}

$options = [
    'base_uri'  => 'http://dev03.aplum-inc.com:8500',
];

// The simple way to use this SDK, is to instantiate the service factory:
$sf = new SensioLabs\Consul\ServiceFactory($options);

// Then, a service could be retrieve from this factory:
$agent = $sf->get(\SensioLabs\Consul\Services\HealthInterface::class);

// 服务地址
$name = 'product-srv-yaoxf';
$res = $agent->service($name, ['passing' => true]);
p( json_decode($res->getBody(), true) );
