<?php

declare(strict_types=1);

namespace App\Module\Db\Dao;

use Hyperf\DbConnection\Db;

/**
 * 需使用框架命令快速生成Dao时
 *      继承extends \Hyperf\DbConnection\Model\Model
 *      设置$connection属性为default
 *      注释掉getConnection方法
 *      注释掉update方法
 */
abstract class AbstractDao/*  extends \Hyperf\DbConnection\Model\Model */
{
    protected ?string $connection/*  = 'default' */;  //分库情况下，解析后所确定的连接
    protected ?string $table;   //分表情况下，解析后所确定的表

    protected array $insert = [];   //解析后的insert。格式：[['字段' => '值',...],...]。无顺序要求
    protected array $update = [];   //解析后的update。格式：['字段' => '值',...]。无顺序要求

    protected array $joinCode = [];    //已联表标识。格式：['joinCode',...]

    protected array $afterField = [];    //获取数据后，再处理的字段

    protected array $jsonField = [];    //json类型字段。这些字段创建|更新时，需要特殊处理

    //#[Inject(value: \App\Module\Db\Model\Platform\Admin::class)]
    protected \App\Module\Db\Model\AbstractModel $model;   //模型
    protected \Hyperf\Database\Query\Builder $builder; //构造器

    public function __construct()
    {
        //子类未定义$model时会自动设置。注意：Dao类目录和Model目录的对应关系
        if (empty($this->model)) {
            $modelClassName = str_replace('\\Dao\\', '\\Model\\', get_class($this));
            $this->model = getModel($modelClassName);
        }
        $this->initBuilder();
    }

    /**
     * 获取连接
     *
     * @return string
     */
    final public function getConnection(): string
    {
        return empty($this->connection) ? $this->model->getConnectionName() : $this->connection;
    }

    /**
     * 获取表名
     *
     * @return string
     */
    final public function getTable(): string
    {
        return empty($this->table) ? $this->model->getTable() : $this->table;
    }

    /**
     * 获取主键名
     *
     * @return string
     */
    final public function getKey(): string
    {
        return $this->model->getKeyName();
    }

    /**
     * 获取全部列
     *
     * @return array
     */
    final public function getAllColumn(): array
    {
        return $this->model->getAllColumn();
    }

    /*----------------解析 开始----------------*/
    /**
     * 解析连接
     *
     * @param array $connectionSelectData   分库情况下，用于确定使用哪个连接
     * @return self
     */
    public function connection(array $connectionSelectData = []): self
    {
        //选择逻辑
        //$this->connection = ''; //设置当前使用的连接
        return $this;
    }

    /**
     * 解析表名
     *
     * @param array $tableSelectData    分表情况下，用于确定使用哪个表
     * @return self
     */
    public function table(array $tableSelectData = []): self
    {
        //选择逻辑
        //$this->table = ''; //设置当前使用的表名
        return $this;
    }

    /**
     * 初始化Db构造器
     *
     * @param string $tableRaw  表的原生表达式。当需要强制索引等特殊情况时使用。示例：'__TABLE__ FORCE INDEX (索引)')。
     * @return self
     */
    final public function initBuilder(string $tableRaw = ''): self
    {
        if (empty($tableRaw)) {
            $this->builder = Db::connection($this->getConnection())->table($this->getTable());
        } else {
            $tableRaw = str_replace('__TABLE__', $this->getTable(), $tableRaw);
            $this->builder = Db::connection($this->getConnection())->table(Db::raw($tableRaw));
        }
        return $this;
    }

    /**
     * 获取Db构造器
     *
     * @return \Hyperf\Database\Query\Builder
     */
    final public function getBuilder(): \Hyperf\Database\Query\Builder
    {
        return $this->builder;
    }

    /**
     * 解析insert（入口）
     *
     * @param array $insert 格式：['字段' => '值',...] 或 [['字段' => '值',...],...]
     * @return self
     */
    final public function insert(array $insert): self
    {
        if (isset($insert[0]) && is_array($insert[0])) {
            foreach ($insert as $k => $v) {
                foreach ($v as $k1 => $v1) {
                    if (!$this->insertOfAlone($k1, $v1, $k)) {
                        $this->insertOfCommon($k1, $v1, $k);
                    }
                }
            }
        } else {
            foreach ($insert as $k => $v) {
                if (!$this->insertOfAlone($k, $v)) {
                    $this->insertOfCommon($k, $v);
                }
            }
        }
        return $this;
    }

    /**
     * 解析update（入口）
     *
     * @param array $update 格式：['字段' => '值',...]
     * @return self
     */
    final public function update(array $update): self
    {
        foreach ($update as $k => $v) {
            if (!$this->updateOfAlone($k, $v)) {
                $this->updateOfCommon($k, $v);
            }
        }
        return $this;
    }

    /**
     * 解析field（入口）
     *
     * @param array $field  格式：['字段',...]
     * @return self
     */
    final public function field(array $field): self
    {
        foreach ($field as $v) {
            if (!$this->fieldOfAlone($v)) {
                $this->fieldOfCommon($v);
            }
        }
        return $this;
    }

    /**
     * 解析where（入口）
     *
     * @param array $where  格式：['字段' => '值', ['字段'，'运算符', '值', 'and|or'],...]
     * @return self
     */
    final public function where(array $where): self
    {
        foreach ($where as $k => $v) {
            if (is_numeric($k) && is_array($v)) {
                if (!$this->whereOfAlone(...$v)) {
                    $this->whereOfCommon(...$v);
                }
            } else {
                if (!$this->whereOfAlone($k, null, $v)) {
                    $this->whereOfCommon($k, null, $v);
                }
            }
        }
        return $this;
    }

    /**
     * 解析group（入口）
     *
     * @param array $group  格式：['字段',...]
     * @return self
     */
    final public function group(array $group): self
    {
        foreach ($group as $v) {
            if (!$this->groupOfAlone($v)) {
                $this->groupOfCommon($v);
            }
        }
        return $this;
    }

    /**
     * 解析order（入口）
     *
     * @param array $order  格式：['字段' => 'asc或desc',...]
     * @return self
     */
    final public function order(array $order): self
    {
        foreach ($order as $k => $v) {
            if (!$this->orderOfAlone($k, $v)) {
                $this->orderOfCommon($k, $v);
            }
        }
        return $this;
    }

    /**
     * 解析insert（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @param integer $index
     * @return boolean
     */
    final protected function insertOfCommon(string $key, $value, int $index = 0): bool
    {
        switch ($key) {
            case 'id':
                $this->insert[$index][$this->getKey()] = $value;
                return true;
            default:
                //数据库不存在的字段过滤掉
                if (in_array($key, $this->getAllColumn())) {
                    if (in_array($key, $this->jsonField)) {
                        if ($value === '' || $value === null) {
                            $this->insert[$index][$key] =  null;
                        } else {
                            $this->insert[$index][$key] =  is_array($value) ? json_encode($value, JSON_UNESCAPED_UNICODE) : $value;
                        }
                    } else {
                        $this->insert[$index][$key] = $value;
                    }
                    return true;
                }
        }
        return false;
    }

    /**
     * 解析update（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    final protected function updateOfCommon(string $key, $value): bool
    {
        switch ($key) {
            case 'id':
                $this->update[$this->getTable() . '.' . $this->getKey()] = $value;
                return true;
            default:
                /* //暂时不考虑其他复杂字段。复杂字段建议直接写入updateOfAlone方法
                list($realKey) = explode('->', $key);   //json情况
                list($realKey) = explode(' AS ', $realKey); //别名情况
                list($realKey) = explode(' as ', $realKey); //别名情况
                $realKey = trim($realKey);  //去除两边空白
                //数据库不存在的字段过滤掉
                if (in_array($realKey, $this->getAllColumn())) { */
                //数据库不存在的字段过滤掉
                if (in_array($key, $this->getAllColumn())) {
                    if (in_array($key, $this->jsonField)) {
                        if ($value === '' || $value === null) {
                            $this->update[$this->getTable() . '.' . $key] = null;
                        } else {
                            $this->update[$this->getTable() . '.' . $key] = is_array($value) ? json_encode($value, JSON_UNESCAPED_UNICODE) : $value;
                        }
                    } else {
                        $this->update[$this->getTable() . '.' . $key] = $value;
                    }
                    return true;
                }
        }
        return false;
    }

    /**
     * 解析field（公共的）
     *
     * @param string $key
     * @return boolean
     */
    final protected function fieldOfCommon(string $key): bool
    {
        switch ($key) {
            case '*':
                $this->builder->addSelect($key);
                return true;
            case 'id':
                $this->builder->addSelect($this->getTable() . '.' . $this->getKey() . ' AS ' . $key);
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->builder->addSelect($this->getTable() . '.' . $key);
                } else {
                    $this->builder->addSelect($key);
                }
                return true;
        }
        return false;
    }

    /**
     * 解析where（公共的）
     *
     * @param string $key
     * @param string|null $operator
     * @param [type] $value
     * @param string|null $boolean
     * @return boolean
     */
    final protected function whereOfCommon(string $key, string $operator = null, $value, string $boolean = null): bool
    {
        switch ($key) {
            case 'id':
                if (is_array($value)) {
                    if (count($value) === 1) {
                        $this->builder->where($this->getTable() . '.' . $this->getKey(), $operator ?? '=', array_shift($value), $boolean ?? 'and');
                    } else {
                        $this->builder->whereIn($this->getTable() . '.' . $this->getKey(), $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where($this->getTable() . '.' . $this->getKey(), $operator ?? '=', $value, $boolean ?? 'and');
                }
                return true;
            case 'excId':
                if (is_array($value)) {
                    if (count($value) === 1) {
                        $this->builder->where($this->getTable() . '.' . $this->getKey(), $operator ?? '<>', array_shift($value), $boolean ?? 'and');
                    } else {
                        $this->builder->whereNotIn($this->getTable() . '.' . $this->getKey(), $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where($this->getTable() . '.' . $this->getKey(), $operator ?? '<>', $value, $boolean ?? 'and');
                }
                return true;
            case 'startTime':
                $this->builder->where($this->getTable() . '.createTime', $operator ?? '>=', date('Y-m-d H:i:s', strtotime($value)), $boolean ?? 'and');
                return true;
            case 'endTime':
                $this->builder->where($this->getTable() . '.createTime', $operator ?? '<=', date('Y-m-d H:i:s', strtotime($value)), $boolean ?? 'and');
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    //id类型字段和部分字段，可通过传递数组做查询
                    if (strtolower(substr($key, -2)) === 'id' || in_array($key, ['configKey'])) {
                        if (is_array($value)) {
                            if (count($value) === 1) {
                                $this->builder->where($this->getTable() . '.' . $key, $operator ?? '=', array_shift($value), $boolean ?? 'and');
                            } else {
                                $this->builder->whereIn($this->getTable() . '.' . $key, $value, $boolean ?? 'and');
                            }
                        } else {
                            $this->builder->where($this->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and');
                        }
                    } else {
                        $this->builder->where($this->getTable() . '.' . $key, $operator ?? '=', $value, $boolean ?? 'and');
                    }
                } else {
                    $this->builder->where($key, $operator ?? '=', $value, $boolean ?? 'and');
                }
                return true;
        }
        return false;
    }

    /**
     * 解析group（公共的）
     *
     * @param string $key
     * @return boolean
     */
    final protected function groupOfCommon(string $key): bool
    {
        switch ($key) {
            case 'id':
                $this->builder->groupBy($this->getTable() . '.' . $this->getKey());
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->builder->groupBy($this->getTable() . '.' . $key);
                } else {
                    $this->builder->groupBy($key);
                }
                return true;
        }
        return false;
    }

    /**
     * 解析order（公共的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    final protected function orderOfCommon(string $key, $value): bool
    {
        switch ($key) {
            case 'id':
                $this->builder->orderBy($this->getTable() . '.' . $this->getKey(), $value);
                return true;
            default:
                if (in_array($key, $this->getAllColumn())) {
                    $this->builder->orderBy($this->getTable() . '.' . $key, $value);
                } else {
                    $this->builder->orderBy($key, $value);
                }
                return true;
        }
        return false;
    }

    /**
     * 解析insert（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @param integer $index
     * @return boolean
     */
    protected function insertOfAlone(string $key, $value, int $index = 0): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->insert[$index][$key] = $value;
                return true;
        } */
        return false;
    }

    /**
     * 解析update（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    protected function updateOfAlone(string $key, $value = null): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->update[$key] = $value;
                return true;
        } */
        return false;
    }

    /**
     * 解析field（独有的）
     *
     * @param string $key
     * @return self
     */
    protected function fieldOfAlone(string $key): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->builder->addSelect($key);
                //$this->builder->addSelect(Db::raw('JSON_UNQUOTE(JSON_EXTRACT(extraData, "$.' . $key . '")) AS ' . $key));   //不能防sql注入
                //$this->builder->selectRaw('IFNULL(字段名, \'\') AS ?', [$key]); //能防sql注入
                return true;
        } */
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
        /* switch ($key) {
            case 'xxxx':
                $this->builder->where($key, $operator ?? '=', $value, $boolean ?? 'and');
                //$this->builder->whereRaw(':key > :value', ['key' => $key, 'value' => $value], $boolean ?? 'and');
                return true;
        } */
        return false;
    }

    /**
     * 解析group（独有的）
     *
     * @param string $key
     * @return boolean
     */
    protected function groupOfAlone(string $key): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->builder->groupBy($key);
                //$this->builder->groupByRaw(':key', ['key' => $key]);
                return true;
        } */
        return false;
    }

    /**
     * 解析order（独有的）
     *
     * @param string $key
     * @param [type] $value
     * @return boolean
     */
    protected function orderOfAlone(string $key, $value = null): bool
    {
        /* switch ($key) {
            case 'xxxx':
                $this->builder->orderBy($key, $value);
                //$this->builder->orderByRaw(':key ' . $value, ['key' => $key]);
                return true;
        } */
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
        return false;
    }

    /**
     * 判断是否联表
     *
     * @return boolean
     */
    final public function isJoin(): bool
    {
        return !empty($this->joinCode);
    }
    /*----------------解析 结束----------------*/

    /*----------------封装部分方法方便使用 开始----------------*/
    /**
     * 保存插入
     *
     * @param boolean $isGetId
     * @return boolean|integer
     */
    final public function saveInsert(bool $isGetId = true): bool|int
    {
        $this->getBuilder();
        if (count($this->insert) > 1 || !$isGetId) {
            return $this->builder->insert($this->insert);
        }
        return $this->builder->insertGetId($this->insert[0]);
    }

    /**
     * 保存更新
     *
     * @param integer $offset
     * @param integer $limit
     * @return integer
     */
    final public function saveUpdate(int $offset = 0, int $limit = 0): int
    {
        $this->getBuilder();
        $this->handleLimit($offset, $limit);
        return $this->builder->update($this->update);
    }

    /**
     * 删除
     *
     * @return integer
     */
    final public function delete(int $offset = 0, int $limit = 0): int
    {
        $this->getBuilder();
        $this->handleLimit($offset, $limit);
        return $this->builder->delete();
    }

    /**
     * 获取信息
     *
     * @param boolean $isUseWriter
     * @return object|null
     */
    final public function getInfo(bool $isUseWriter = false): object|null
    {
        $this->getBuilder();
        if ($isUseWriter) {
            $this->builder->useWritePdo();
        }
        $info = $this->builder->first();
        if (empty($info)) {
            return $info;
        }
        $this->afterField($info);
        return $info;
    }

    /**
     * 获取列表
     *
     * @param integer $offset
     * @param integer $limit
     * @param boolean $isUseWriter  读写分离时，是否使用写库读（因读写分离有延迟，有些时候需要使用写库读取）
     * @return array
     */
    final public function getList(int $offset = 0, int $limit = 0, bool $isUseWriter = false): array
    {
        $this->getBuilder();
        if ($isUseWriter) {
            $this->builder->useWritePdo();
        }
        $this->handleLimit($offset, $limit);
        $list = $this->builder->get()->toArray();
        if (!empty($this->afterField)) {
            foreach ($list as &$v) {
                $this->afterField($v);
            }
        }
        return $list;
    }

    /**
     * 处理limit（limit不算常用，故不做解析，也不做解析后处理，即不在getBuilder方法中处理）
     *
     * @return self
     */
    final protected function handleLimit(int $offset, int $limit): self
    {
        if ($limit > 0) {
            $this->builder->offset($offset)->limit($limit);
        } elseif ($offset > 0) {    //当offset>0，limit==0时表示取剩下全部数据。需要limit足够大，这里写99999999，这样还不够的话，服务器也抗不住了
            $this->builder->offset($offset)->limit(99999999);
        }
        return $this;
    }

    /**
     * 获取数据库数据后，再处理的字段（入口）
     *
     * @param object $info
     * @return void
     */
    final protected function afterField(object &$info)
    {
        //isset($info->{$this->getKey()}) ? $info->id = $info->{$this->getKey()} : null;  //设置id字段
        foreach ($this->afterField as $field) {
            if (!$this->afterFieldOfAlone($field, $info)) {
                $this->afterFieldOfCommon($field, $info);
            }
        }
    }

    /**
     * 获取数据后，再处理的字段（公共的）
     *
     * @param string $key
     * @param object $info
     * @return boolean
     */
    final protected function afterFieldOfCommon(string $key, object &$info): bool
    {
        /* switch ($key) {
            case 'id':
                $info->{$key} = $info->{$this->getKey()};
                return true;
            default:
                getLogger('daoAfterField')->warning('未处理字段：' . $key);
        } */
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
        /* switch ($key) {
            case 'xxxx':
                $info->xxxx = 'xxxx';
                return true;
        } */
        return false;
    }
    /*----------------封装部分方法方便使用 结束----------------*/
}