<?php

require_once __DIR__ . '/vendor/autoload.php';

$app = new Silex\Application();

$app->register(new Silex\Provider\TwigServiceProvider(), array(
    'twig.path' => __DIR__.'/views',
));

$app->get('/{name}', function ($name) use ($app) {

	$date = new DateTime();

    return $app['twig']->render('index.twig', array(
        'name' => $name,
        'time' => $date->format('H:ia')
    ));
});

$app->run();