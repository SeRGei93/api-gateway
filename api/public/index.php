<?php

require __DIR__ . '/../vendor/autoload.php';

use Slim\Factory\AppFactory;

$app = AppFactory::create();

$app->get('/', function ($request, $response) {
    $response->getBody()->write(json_encode(['status' => 'API работает!']));
    return $response->withHeader('Content-Type', 'application/json');
});

$app->post('/users', function ($request, $response) {
    $response->getBody()->write(json_encode(['users' => ['User1', 'User2']]));
    return $response->withHeader('Content-Type', 'application/json');
});

$app->post('/cards', function ($request, $response) {
    $response->getBody()->write(json_encode(['cards' => ['Card1', 'Card2']]));
    return $response->withHeader('Content-Type', 'application/json');
});

$app->run();