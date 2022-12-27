<?php

declare(strict_types=1);

/**
 *  接口返回码各位置说明
 *      1位：应用标识。 0公共标识
 *      2-3位：功能模块标识。   00公共模块；01权限模块
 *      4-6位：错误码。
 */
return [
    '00000000' => 'success',
    '99999999' => 'fail',
];
