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
$kv = $sf->get(\SensioLabs\Consul\Services\KVInterface::class);

$kv->put('test/foo/bar', 'yaoxf');
$kv->get('test/foo/bar', ['raw' => true]);
$kv->delete('test/foo/bar');

