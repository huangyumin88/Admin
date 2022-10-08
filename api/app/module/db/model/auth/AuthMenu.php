<?php

declare(strict_types=1);

namespace app\module\db\model\auth;

use app\module\db\model\AbstractModel;

class AuthMenu extends AbstractModel
{
    protected string $table = 'auth_menu';
    protected string $primaryKey = 'menuId';
}
