<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $roleId 权限角色ID
 * @property int $sceneId 权限场景ID
 * @property int $tableId 关联表ID（0表示平台创建，其他值根据authSceneId对应不同表，表示是哪个表内哪个机构或个人创建）
 * @property string $roleName 名称
 * @property int $isStop 是否停用：0否 1是
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class Role extends AbstractDao
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
            case 'menuName':
                $this->joinOfAlone($key);
                $this->field['select'][] = getDao(Menu::class)->getTable() . '.' . $key;
                return true;
            case 'menuIdArr':
            case 'actionIdArr':
                $this->afterField[] = $key;

                //需要id字段
                $this->field['select'][] = $this->getTable() . '.' . $this->getKey();
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
            case 'menuName':
                $menuDao = getDao(Menu::class);
                $menuDaoTable = $menuDao->getTable();
                if (!isset($this->join[$menuDaoTable])) {
                    $this->join[$menuDaoTable] = [
                        'method' => 'leftJoin',
                        'param' => [
                            $menuDaoTable,
                            $menuDaoTable . '.menuId',
                            '=',
                            $this->getTable() . '.menuId'
                        ]
                    ];
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
            case 'menuIdArr':
                $info->sceneIdArr = getDao(RoleRelToMenu::class)->where(['roleId' => $info->{$this->getKey()}])->getBuilder()->pluck('menuId')->toArray();
                return true;
            case 'actionIdArr':
                $info->sceneIdArr = getDao(RoleRelToAction::class)->where(['roleId' => $info->{$this->getKey()}])->getBuilder()->pluck('actionId')->toArray();
                return true;
        }
        return false;
    }
}
