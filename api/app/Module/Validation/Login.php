<?php

declare(strict_types=1);

namespace App\Module\Validation;

class Login extends AbstractValidation
{
    protected array $rule = [
        'account' => 'required|alpha_dash|between:4,30',
        'password' => 'required|alpha_num|size:32',
    ];

    protected array $scene = [
        'encryptStr' => [
            'account'
        ],
    ];
}
