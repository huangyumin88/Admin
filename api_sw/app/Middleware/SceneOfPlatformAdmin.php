<?php

declare(strict_types=1);

namespace App\Middleware;

use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class SceneOfPlatformAdmin implements \Psr\Http\Server\MiddlewareInterface
{
    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        try {
            $container = getContainer();
            $sceneCode = $container->get(\App\Module\Logic\Auth\Scene::class)->getCurrentSceneCode();
            if ($sceneCode == 'platformAdmin') {
                $container->get(\App\Module\Service\Login::class)->verifyToken($sceneCode);
            }
            $response = $handler->handle($request);
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
