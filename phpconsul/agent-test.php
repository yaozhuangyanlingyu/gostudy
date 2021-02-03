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
$agent = $sf->get(\SensioLabs\Consul\Services\AgentInterface::class);

// consul集群成员
// p( $agent->members() );

// 服务地址
$rsp = $agent->services();
$services = json_decode($rsp->getBody(), true);
$data = [];
foreach($services as $s){
    $data[$s['Service']][] = [
        'port'      => $s['Port'],
        'address'   => $s['Address'],
    ];
}
p( $data );
