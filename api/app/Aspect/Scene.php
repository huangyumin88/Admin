<?php

declare(strict_types=1);

namespace App\Aspect;

use App\Module\Db\Dao\Auth\Scene as AuthScene;
use App\Module\Logic\Auth\Scene as LogicAuthScene;
use Hyperf\Context\Context;
use Hyperf\Di\Annotation\Aspect;
use Hyperf\Di\Aop\ProceedingJoinPoint;
use Hyperf\HttpServer\Contract\RequestInterface;
use Psr\Http\Message\ServerRequestInterface;

#[Aspect]
class Scene extends AbstractAspect
{
    //执行优先级（大值优先）
    public ?int $priority = 20;

    //要切入的类，可以多个，亦可通过 :: 标识到具体的某个方法，通过 * 可以模糊匹配
    public array $classes = [
        \App\Controller\Index::class,
        \App\Controller\Login::class,
        \App\Controller\Auth\AuthScene::class
    ];

    //要切入的注解，具体切入的还是使用了这些注解的类，仅可切入类注解和类方法注解
    public array $annotations = [];

    /**
     * @param ProceedingJoinPoint $proceedingJoinPoint
     * @return void
     */
    public function process(ProceedingJoinPoint $proceedingJoinPoint)
    {
        $sceneCode = $this->container->get(RequestInterface::class)->getHeaderLine('Scene');
        if (empty($sceneCode)) {
            throwFailJson('001001');
        }
        $sceneInfo = make(AuthScene::class)->where(['sceneCode' => $sceneCode])->getInfo();
        if (empty($sceneInfo)) {
            throwFailJson('001001');
        }
        $sceneInfo->sceneConfig = json_decode($sceneInfo->sceneConfig, true);
        //$this->container->get(LogicAuthScene::class)->setRequestSceneInfo($sceneInfo);
        $request = Context::get(ServerRequestInterface::class);
        $request->sceneInfo = $sceneInfo;
        Context::set(ServerRequestInterface::class, $request); //重新设置请求对象，改变协程上下文内的请求对象
        var_dump(Context::get(ServerRequestInterface::class)->sceneInfo);
        var_dump($this->container->get(RequestInterface::class)->sceneInfo);
        try {
            $response = $proceedingJoinPoint->process();
            return $response;
        } catch (\Throwable $th) {
            throw $th;
        }
    }
}
