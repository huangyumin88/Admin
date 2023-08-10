<?php

declare(strict_types=1);

/**
 *  接口返回码各位置说明
 *      1位：错误类型。 1框架错误；2数据库错误；3业务错误；4三方接口或回调错误；8参数错误；9公共错误；
 *      2-4位：功能模块标识。   999公共错误；
 *      5-8位：错误码。
 */
return [
    0 => '成功',

    99999999 => '失败',

    19990404 => '找不到路由',    //Not Found
    19990500 => '内部服务器错误',   //Internal Server Error.

    29991062 => '数据已存在(:errField)',   // #当数据库报1062重复索引时使用
    29999994 => '不能删除含有子级的数据',
    29999995 => '父级不能是自身的子孙级',
    29999996 => '父级不能是自身',
    29999997 => '父级不存在',
    29999998 => '数据不存在',

    30000000 => '不能修改平台超级管理员',
    30000001 => '不能删除平台超级管理员',

    39990000 => '账号不存在',
    39990001 => '密码错误',
    39990002 => '账号被停用',
    39990003 => '校验密码错误',
    /*--------前端收到以下39994xxx错误码，需强制退出登录 开始--------*/
    39994000 => '您还未登录，请先登录',
    39994001 => 'token无效',
    39994002 => '账号在其他地方登录',
    39994003 => '账号不存在',
    39994004 => '您的账号已被停用',
    /*--------前端收到以下39994xxx错误码，需强制退出登录 结束--------*/
    39999002 => '签名失败',
    39999003 => '签名方式不存在',
    39999997 => '无权限',
    39999998 => '场景停用',
    39999999 => '场景非法',

    79990000 => '阿里云OSS回调Header签名不能为空',
    79990001 => '阿里云OSS回调Header公钥Url不能为空',
    79990002 => '阿里云OSS回调Header公钥Url错误',
    79990003 => '阿里云OSS验证签名错误',

    89999996 => "不支持批量修改(:errField)",
    89999997 => '验证场景不存在',
    89999998 => '参数非法',   //正规操作不会报这个错
    89999999 => '参数错误',
];
