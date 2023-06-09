<?php

declare(strict_types=1);

namespace App\Module\Db\Model\Auth;

use App\Module\Db\Model\AbstractModel;

/**
 * @property int $roleId 权限角色ID
 * @property int $actionId 权限操作ID
 * @property string $updatedAt 更新时间
 * @property string $createdAt 创建时间
 */
class RoleRelToAction extends AbstractModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'auth_role_rel_to_action';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = ['roleId', 'actionId', 'updatedAt', 'createdAt'];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['roleId' => 'integer', 'actionId' => 'integer'];
}
