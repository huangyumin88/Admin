<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $actionId 权限操作ID
 * @property string $actionName 名称
 * @property string $actionCode 标识（代码中用于判断权限）
 * @property string $remark 备注
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Action extends AbstractDao
{
    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return boolean
     */
    protected function fieldOfAlone(string $key): bool
    {
        switch ($key) {
            case 'sceneIdArr':
                //需要id字段
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey());

                $this->afterField[] = $key;
                return true;
        }
        return false;
    }

    /**
     * 解析where（独有的）
     *
     * @param string $key
     * @param string|null $operator
     * @param [type] $value
     * @param string|null $boolean
     * @return boolean
     */
    protected function whereOfAlone(string $key, string $operator = null, $value, string $boolean = null): bool
    {
        switch ($key) {
            case 'sceneId':
                if (is_array($value)) {
                    if (count($value) === 1) {
                        $this->builder->where(getDao(ActionRelToScene::class)->getTable() . '.' . $key, $operator ?? '=', array_shift($value), $boolean ?? 'and');
                    } else {
                        $this->builder->whereIn(getDao(ActionRelToScene::class)->getTable() . '.' . $key, $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where(getDao(ActionRelToScene::class)->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and');
                }

                $this->joinOfAlone('actionRelToScene');
                return true;
            case 'selfAction': //获取当前登录身份可用的操作。参数：['sceneCode'=>场景标识, 'loginId'=>登录身份id]
                $sceneInfo = getContainer()->get(\App\Module\Logic\Auth\Scene::class)->getCurrentSceneInfo();    //当开启切面\App\Aspect\Scene时有值
                $sceneId = $sceneInfo === null ? getDao(Scene::class)->where(['sceneCode' => $value['sceneCode']])->getBuilder()->value('sceneId') : $sceneInfo->sceneId;
                $this->builder->where($this->getTable() . '.isStop', '=', 0, 'and');
                $this->builder->where(getDao(ActionRelToScene::class)->getTable() . '.sceneId', '=', $sceneId, 'and');
                $this->joinOfAlone('actionRelToScene');
                switch ($value['sceneCode']) {
                    case 'platformAdmin':
                        if ($value['loginId'] === getConfig('app.superPlatformAdminId')) { //平台超级管理员，不再需要其他条件
                            return true;
                        }
                        $this->builder->where(getDao(Role::class)->getTable() . '.isStop', '=', 0, 'and');
                        $this->builder->where(getDao(RoleRelOfPlatformAdmin::class)->getTable() . '.adminId', '=', $value['loginId'], 'and');

                        $this->joinOfAlone('roleRelToAction');
                        $this->joinOfAlone('role');
                        $this->joinOfAlone('roleRelOfPlatformAdmin');
                        break;
                }

                $this->group(['id']);
                return true;
        }
        return false;
    }

    /**
     * 解析join（独有的）
     *
     * @param string $key   键，用于确定关联表
     * @param [type] $value 值，用于确定关联表
     * @return boolean
     */
    protected function joinOfAlone(string $key, $value = null): bool
    {
        switch ($key) {
            case 'actionRelToScene':
                $actionRelToSceneDao = getDao(ActionRelToScene::class);
                $actionRelToSceneDaoTable = $actionRelToSceneDao->getTable();
                if (!in_array($actionRelToSceneDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $actionRelToSceneDaoTable;
                    $this->builder->leftJoin($actionRelToSceneDaoTable, $actionRelToSceneDaoTable . '.actionId', '=', $this->getTable() . '.actionId');
                }
                return true;
            case 'roleRelToAction':
                $roleRelToActionDao = getDao(RoleRelToAction::class);
                $roleRelToActionDaoTable = $roleRelToActionDao->getTable();
                if (!in_array($roleRelToActionDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $roleRelToActionDaoTable;
                    $this->builder->leftJoin($roleRelToActionDaoTable, $roleRelToActionDaoTable . '.actionId', '=', $this->getTable() . '.actionId');
                }
                return true;
            case 'role':
                $roleRelToActionDao = getDao(RoleRelToAction::class);
                $roleRelToActionDaoTable = $roleRelToActionDao->getTable();

                $roleDao = getDao(Role::class);
                $roleDaoTable = $roleDao->getTable();
                if (!in_array($roleDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $roleDaoTable;
                    $this->builder->leftJoin($roleDaoTable, $roleDaoTable . '.roleId', '=', $roleRelToActionDaoTable . '.roleId');
                }
                return true;
            case 'roleRelOfPlatformAdmin':
                $roleRelToActionDao = getDao(RoleRelToAction::class);
                $roleRelToActionDaoTable = $roleRelToActionDao->getTable();

                $roleRelOfPlatformAdminDao = getDao(RoleRelOfPlatformAdmin::class);
                $roleRelOfPlatformAdminDaoTable = $roleRelOfPlatformAdminDao->getTable();
                if (!in_array($roleRelOfPlatformAdminDaoTable, $this->joinCode)) {
                    $this->joinCode[] = $roleRelOfPlatformAdminDaoTable;
                    $this->builder->leftJoin($roleRelOfPlatformAdminDaoTable, $roleRelOfPlatformAdminDaoTable . '.roleId', '=', $roleRelToActionDaoTable . '.roleId');
                }
                return true;
        }
        return false;
    }

    /**
     * 获取数据后，再处理的字段（独有的）
     *
     * @param string $key
     * @param object $info
     * @return boolean
     */
    protected function afterFieldOfAlone(string $key, object &$info): bool
    {
        switch ($key) {
            case 'sceneIdArr':
                $info->{$key} = getDao(ActionRelToScene::class)->where(['actionId' => $info->{$this->getKey()}])->getBuilder()->pluck('sceneId')->toArray();
                return true;
        }
        return false;
    }
}