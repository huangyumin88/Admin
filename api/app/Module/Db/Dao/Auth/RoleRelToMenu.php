<?php

declare(strict_types=1);

namespace App\Module\Db\Dao\Auth;

use App\Module\Db\Dao\AbstractDao;

/**
 * @property int $roleId 权限角色ID
 * @property int $menuId 权限菜单ID
 * @property string $updateTime 更新时间
 * @property string $createTime 创建时间
 */
class RoleRelToMenu extends AbstractDao
{
}
