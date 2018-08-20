<?php


require dirname(__FILE__) . '/../vendor/autoload.php';

$appId = "123";
try {
    $client = new Meta\AgentClient('localhost:50051', ['credentials' => Grpc\ChannelCredentials::createInsecure(),]);
    $request = new Meta\MetaRequest();
    $request->setAppId($appId);
    list($reply, $status) = $client->GetInfo($request)->wait();
    echo $reply->serializeToJsonString() . PHP_EOL;
} catch (\Exception $e) {
    echo "error : " . $e->getMessage();
}

function greet($name)
{
    $client = new Helloworld\GreeterClient('127.0.0.1:50051', [
        'credentials' => Grpc\ChannelCredentials::createInsecure(),
    ]);
    $request = new Helloworld\HelloRequest();
    $request->setName($name);
    list($reply, $status) = $client->SayHello($request)->wait();
    $message = $reply->getMessage();
    return $message;
}

$name = !empty($argv[1]) ? $argv[1] : 'world';
echo "the greet: " . greet($name) . "\n";
