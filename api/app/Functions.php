<?php

declare(strict_types=1);

use Hyperf\DbConnection\Db;

/*----------------基于业务逻辑封装的函数  开始----------------*/

if (!function_exists('getConfig')) {
    function getConfig(string $key, mixed $default = null, string $type = 'default'): mixed
    {
        switch ($type) {
            case 'platformConfig':
                $allPlatformConfig = getContainer()->get('allPlatformConfig');
                return $allPlatformConfig[$key] ?? $default;
                break;
            default:
                return config($key, $default);
        }
    }
}

if (!function_exists('dbTablePartition')) {
    /**
     * 数据库表按时间做分区（通用，默认以分区最大日期作为分区名）
     * 注意：如果表有唯一索引（含主键），则用于建立分区的字段必须和唯一索引字段组成联合索引
     * 建议：分区间隔时间，分区数量设置后，两者总时长要比定时器间隔多几天时间，方便分区失败时，有时间让技术人工处理
     *
     * @param string $daoClassName  数据库表对应的dao类。示例：App\Module\Db\Dao\Log\Request::class
     * @param string $partitionField    分区字段，即根据该字段做分区
     * @param integer $partitionTime    间隔多长时间创建一个分区，单位：秒
     * @param integer $partitionNumber  当前时间后面，需要新增的分区数量
     * @return void
     */
    function dbTablePartition(string $daoClassName, int $partitionTime = 24 * 60 * 60, int $partitionNumber = 1, string $partitionField = 'createTime')
    {
        /* //查看分区
        SELECT PARTITION_NAME, PARTITION_EXPRESSION, PARTITION_DESCRIPTION, TABLE_ROWS
        FROM INFORMATION_SCHEMA.PARTITIONS
        WHERE TABLE_SCHEMA = SCHEMA() AND TABLE_NAME = '表名';
        //修改表为分区表
        ALTER TABLE `表名` PARTITION BY RANGE (TO_DAYS( 分区字段 ))
        (PARTITION `20220115` VALUES LESS THAN (TO_DAYS('2022-01-16')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
        PARTITION `max` VALUES LESS THAN (MAXVALUE) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 );
        //新增分区
        ALTER TABLE `表名` ADD PARTITION
        (PARTITION `20220115` VALUES LESS THAN (TO_DAYS('2022-01-16')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
        PARTITION `20220116` VALUES LESS THAN (TO_DAYS('2022-01-17')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 );
        //删除分区
        ALTER TABLE `表名` DROP PARTITION 20220115,20220116; */
        $dao = getDao($daoClassName);
        $connection = $dao->getConnection();
        $table = $dao->getTable();
        $db = Db::connection($connection);

        /**--------查询分区（不是分区表或无分区，查询结果都会有一项，且第一项内PARTITION_NAME值为null） 开始--------**/
        $partitionSelSql = 'SELECT PARTITION_NAME FROM INFORMATION_SCHEMA.PARTITIONS WHERE TABLE_SCHEMA = SCHEMA() AND TABLE_NAME = \'' . $table . '\'';
        $partitionResult = $db->select($partitionSelSql, [], false);
        $partitionList = [];
        if ($partitionResult[0]->PARTITION_NAME !== null) {
            foreach ($partitionResult as $v) {
                $partitionList[] = $v->PARTITION_NAME;
            }
        }
        /**--------查询分区（不是分区表或无分区，查询结果都会有一项，且第一项内PARTITION_NAME值为null） 结束--------**/

        /**--------无分区则建立当前分区 开始--------**/
        if (empty($partitionList)) {
            $ltTime = time() + $partitionTime;
            $ltDate = date('Y-m-d', $ltTime);
            $partitionName = date('Ymd', $ltTime - 24 * 60 * 60); //该分区的最大日期作为分区名
            $partitionCreateSql = 'PARTITION `' . $partitionName . '` VALUES LESS THAN (TO_DAYS( \'' . $ltDate . '\' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0';
            $partitionCreateSql = 'ALTER TABLE `' . $table . '` PARTITION BY RANGE (TO_DAYS( ' . $partitionField . ' )) ( ' . $partitionCreateSql . ' )';
            $db->update($partitionCreateSql);
            $maxPartitionName = $partitionName;
        } else {
            $maxPartitionName = $partitionList[count($partitionList) - 1];
        }
        /**--------无分区则建立当前分区 结束--------**/

        /**--------检测需要创建的分区是否存在，没有则新增分区 开始--------**/
        $partitionAddSqlList = [];
        for ($i = 1; $i <= $partitionNumber; $i++) {
            //时间超过最大的分区后才开始新增分区，且以最大分区的时间开始往后计算
            if (date('Ymd', time() + ($i + 1) * $partitionTime - 24 * 60 * 60) > $maxPartitionName) {
                $ltTime = strtotime($maxPartitionName) + ($i + 1) * $partitionTime;
                $ltDate = date('Y-m-d', $ltTime);
                $partitionName = date('Ymd', $ltTime - 24 * 60 * 60); //该分区的最大日期作为分区名
                $partitionAddSqlList[] = 'PARTITION `' . $partitionName . '` VALUES LESS THAN (TO_DAYS( \'' . $ltDate . '\' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0';
            }
        }
        if (!empty($partitionAddSqlList)) {
            $partitionAddSql = implode(',', $partitionAddSqlList);
            $partitionAddSql = 'ALTER TABLE `' . $table . '` ADD PARTITION ( ' . $partitionAddSql . ' )';
            $db->update($partitionAddSql);
        }
        /**--------检测需要创建的分区是否存在，没有则新增分区 结束--------**/
    }
}
/*----------------基于业务逻辑封装的函数  结束----------------*/


/*----------------基于当前框架封装的函数  开始----------------*/
if (!function_exists('throwSuccessJson')) {
    /**
     * 抛出错误（利用错误处理返回结果json格式。好处：深层次调用无需反复return）
     *
     * @param array $data
     * @param string $code
     * @param string $msg
     * @throws \App\Exception\Json
     * @return void
     */
    function throwSuccessJson(array $data = [], string $code = '00000000', string $msg = '')
    {
        throw make(\App\Exception\Json::class, ['code' => $code, 'msg' => $msg, 'data' => $data]);
    }
}

if (!function_exists('throwFailJson')) {
    /**
     * 抛出错误（利用错误处理返回结果json格式。好处：深层次调用无需反复return）
     *
     * @param string $code
     * @param string $msg
     * @param array $data
     * @throws \App\Exception\Json
     * @return void
     */
    function throwFailJson(string $code = '99999999', string $msg = '', array $data = [])
    {
        throw make(\App\Exception\Json::class, ['code' => $code, 'msg' => $msg, 'data' => $data]);
    }
}

if (!function_exists('throwRaw')) {
    /**
     * 抛出错误（利用错误处理返回结果raw格式。好处：深层次调用无需反复return）
     *
     * @param string $raw
     * @throws \App\Exception\Raw
     * @return void
     */
    function throwRaw(string $raw)
    {
        throw make(\App\Exception\Raw::class, ['raw' => $raw]);
    }
}

if (!function_exists('getContainer')) {
    /**
     * 获取容器
     *
     * @return \Psr\Container\ContainerInterface
     */
    function getContainer(): \Psr\Container\ContainerInterface
    {
        return \Hyperf\Utils\ApplicationContext::getContainer();
    }
}

if (!function_exists('getDao')) {
    /**
     * 获取Dao对象
     *  注意：
     *      app\Module\Db\Dao文件夹内的类统一使用此方法实例化。防止误使用容器获取，容器获取的实例状态改变会污染框架环境
     *
     * @param string $className
     * @return object
     */
    function getDao(string $className): object
    {
        return make($className);
    }
}

if (!function_exists('getCache')) {
    /**
     * 获取Cache对象
     *  注意：
     *      当确定使用的缓存对象一定不会切换连接库时（即不改变app\Module\Cache\AbstractCache类的$cache变量），可使用容器缓存获取
     * @param string $className
     * @return object
     */
    function getCache(string $className): object
    {
        //return make($className);
        return getContainer()->get($className);
    }
}

if (!function_exists('getLogger')) {
    /**
     * 获取Logger对象
     * 
     * @param string $name  对应文件内开头的name
     * @param string $group 对应config/autoload/logger.php内的key
     * @return \Psr\Log\LoggerInterface
     */
    function getLogger(string $name = 'hyperf', string $group = 'default'): \Psr\Log\LoggerInterface
    {
        return getContainer()->get(\Hyperf\Logger\LoggerFactory::class)->get($name, $group);
    }
}

if (!function_exists('getRequest')) {
    /**
     * 获取Request对象
     * 
     * @return \Hyperf\HttpServer\Contract\RequestInterface
     */
    function getRequest(): \Hyperf\HttpServer\Contract\RequestInterface
    {
        return getContainer()->get(\Hyperf\HttpServer\Contract\RequestInterface::class);
    }
}

if (!function_exists('setCurrentRequestAttribute')) {
    /**
     * 在当前请求中，设置属性
     * 
     * @param string $attrName
     * @param mixed $value
     * @return void
     */
    function setCurrentRequestAttribute(string $attrName, mixed $value)
    {
        $request = \Hyperf\Context\Context::get(\Psr\Http\Message\ServerRequestInterface::class);
        $request = $request->withAttribute($attrName, $value);
        \Hyperf\Context\Context::set(\Psr\Http\Message\ServerRequestInterface::class, $request);
    }
}

if (!function_exists('getCurrentRequestAttribute')) {
    /**
     * 在当前请求中，获取属性
     * 
     * @param string $attrName
     * @return mixed
     */
    function getCurrentRequestAttribute(string $attrName): mixed
    {
        return getContainer()->get(\Hyperf\HttpServer\Contract\RequestInterface::class)->getAttribute($attrName);
    }
}

if (!function_exists('getRequestScheme')) {
    /**
     * 获取当前请求是http还是https
     *
     * @return string
     */
    function getRequestScheme(): string
    {
        //nginx转发过来的请求，hyperf框架无法识别是否是https，默认都是http。
        //如要识别，需要nginx域名配置文件中设置转发时，增加一个头部用于说明。下面是nginx中所需增加配置，X-Forwarded-Proto名称可自定义
        /* map $http_x_forwarded_proto $admin_scheme {
            default $scheme;
            https https;
        }
        proxy_set_header X-Forwarded-Proto $admin_scheme; */
        return getRequest()->header('x-forwarded-proto', 'http');
    }
}

if (!function_exists('getRequestUrl')) {
    /**
     * 获取当前请求Url
     *
     * @param integer $type  类型。以下返回示例
     *      0：http(s)://www.xxxx.com
     *      1：http(s)://www.xxxx.com/test
     *      2：http(s)://www.xxxx.com/test?a=1&b=2
     * @return string
     */
    function getRequestUrl(int $type = 0): string
    {
        switch ($type) {
            case 1:
                $url = getRequest()->url();
                $scheme = getRequestScheme();
                return $scheme == 'https' ? str_replace('http://', $scheme . '://', $url) : $url;
            case 2:
                $url = getRequest()->fullUrl();
                $scheme = getRequestScheme();
                return $scheme == 'https' ? str_replace('http://', $scheme . '://', $url) : $url;
            case 0:
            default:
                $url = getRequestUrl(1);
                $path = getRequest()->getPathInfo();
                return $path == '/' ? $url : str_replace($path, '', $url);
        }
    }
}
/*----------------基于当前框架封装的函数  结束----------------*/


/*----------------基于PHP封装的函数  开始----------------*/
if (!function_exists('randStr')) {
    /**
     * 随机字符串
     *
     * @param int $length
     * @param bool $number
     * @param bool $lower
     * @param bool $upper
     * @return string
     */
    function randStr(int $length, bool $number = true, bool $lower = true, bool $upper = true): string
    {
        mt_srand();
        $str = '';
        $number ? $str .= '0123456789' : null;
        $lower ? $str .= 'abcdefghijklmnopqrstuvwxyz' : null;
        $upper ? $str .= 'ABCDEFGHIJKLMNOPQRSTUVWXYZ' : null;
        $strLen = strlen($str);
        $randStr = '';
        for ($i = 0; $i < $length; $i++) {
            $index = mt_rand(0, $strLen - 1);
            $randStr .= $str[$index];
        }
        return $randStr;
    }
}
/*----------------基于PHP封装的函数  结束----------------*/
