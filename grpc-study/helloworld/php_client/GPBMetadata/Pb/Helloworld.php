<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pb/helloworld.proto

namespace GPBMetadata\Pb;

class Helloworld
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(hex2bin(
            "0a99010a1370622f68656c6c6f776f726c642e70726f746f12027062221c" .
            "0a0c48656c6c6f52657175657374120c0a046e616d65180120012809221d" .
            "0a0a48656c6c6f5265706c79120f0a076d65737361676518012001280932" .
            "390a0747726565746572122e0a0853617948656c6c6f12102e70622e4865" .
            "6c6c6f526571756573741a0e2e70622e48656c6c6f5265706c7922006206" .
            "70726f746f33"
        ), true);

        static::$is_initialized = true;
    }
}
