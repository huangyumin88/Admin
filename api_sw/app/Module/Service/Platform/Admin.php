<?php

declare(strict_types=1);

namespace App\Module\Service\Platform;

use App\Module\Db\Dao\Auth\RoleRelOfPlatformAdmin;
use App\Module\Logic\Platform\Admin as PlatformAdmin;
use App\Module\Service\AbstractService;

class Admin extends AbstractService
{
    /**
     * 创建
     *
     * @param array $data
     * @return void
     */
    public function create(array $data)
    {
        $id = $this->getDao()->insert($data)->saveInsert();
        if (empty($id)) {
            throwFailJson();
        }
        if (isset($data['roleIdArr'])) {
            $this->container->get(PlatformAdmin::class)->saveRelRole($data['roleIdArr'], $id);
        }
        throwSuccessJson();
    }

    /**
     * 更新
     *
     * @param array $data
     * @param array $where
     * @return void
     */
    public function update(array $data, array $where)
    {
        if (isset($data['checkPassword']) && $data['checkPassword'] != $this->getDao()->where($where)->getBuilder()->value('password')) {
            throwFailJson(39990003);
        }

        if (isset($data['roleIdArr'])) {
            $idArr = $this->getIdArr($where);
            foreach ($idArr as $id) {
                $this->container->get(PlatformAdmin::class)->saveRelRole($data['roleIdArr'], $id);
            }
            $this->getDao()->where($where)->update($data)->saveUpdate();    //有可能只改roleIdArr
        } else {
            $result = $this->getDao()->where($where)->update($data)->saveUpdate();
            if (empty($result)) {
                throwFailJson();
            }
        }
        throwSuccessJson();
    }

    /**
     * 删除
     *
     * @param array $where
     * @return void
     */
    public function delete(array $where)
    {
        $idArr = $this->getIdArr($where);
        $result = $this->getDao()->where($where)->delete();
        if (empty($result)) {
            throwFailJson();
        }
        getDao(RoleRelOfPlatformAdmin::class)->where(['adminId' => $idArr])->delete();
        throwSuccessJson();
    }
}