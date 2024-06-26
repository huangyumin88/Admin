package utils

import (
	"api/internal/consts"
	"api/internal/dao"
	"context"
	"os/exec"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/tools/imports"
)

// 生成错误码
func NewErrorCode(ctx context.Context, code int, msg string, data ...map[string]interface{}) error {
	detail := map[string]interface{}{}
	if len(data) > 0 && data[0] != nil {
		detail = data[0]
	}
	if msg == `` {
		switch code {
		case 89999996:
			msg = g.I18n().Tf(ctx, `code.`+gconv.String(code), gconv.String(detail[`errField`]))
			delete(detail, `errField`)
		default:
			msg = g.I18n().T(ctx, `code.`+gconv.String(code))
		}
	}
	return gerror.NewCode(gcode.New(code, ``, detail), msg)
}

// Http返回json
func HttpWriteJson(ctx context.Context, data map[string]interface{}, code int, msg string) {
	resData := map[string]interface{}{
		`code`: code,
		`msg`:  msg,
		`data`: data,
	}
	if msg == `` {
		resData[`msg`] = g.I18n().T(ctx, `code.`+gconv.String(code))
	}
	g.RequestFromCtx(ctx).Response.WriteJson(resData)
}

// 设置场景信息
func SetCtxSceneInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.ConstCtxSceneInfoName, info)
}

// 获取场景信息
func GetCtxSceneInfo(ctx context.Context) gdb.Record {
	tmp := ctx.Value(consts.ConstCtxSceneInfoName)
	if tmp == nil {
		return nil
	}
	return tmp.(gdb.Record)
}

// 设置登录身份信息
func SetCtxLoginInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.ConstCtxLoginInfoName, info)
}

// 获取登录身份信息
func GetCtxLoginInfo(ctx context.Context) gdb.Record {
	tmp := ctx.Value(consts.ConstCtxLoginInfoName)
	if tmp == nil {
		return nil
	}
	return tmp.(gdb.Record)
}

// 获取当前请求Url
func GetRequestUrl(ctx context.Context, flag int) (url string) {
	r := g.RequestFromCtx(ctx)
	switch flag {
	case 0: //http(s)://www.xxxx.com
		url = gstr.Replace(r.GetUrl(), r.URL.String(), ``)
	case 1: //http(s)://www.xxxx.com/test
		url = gstr.Replace(r.GetUrl(), r.URL.String(), ``) + r.URL.Path
	case 2: //http(s)://www.xxxx.com/test?a=1&b=2
		url = r.GetUrl()
	}
	return
}

// 获取服务器外网ip
func GetServerNetworkIp() string {
	cmd := exec.Command(`/bin/bash`, `-c`, `curl -s ifconfig.me`)
	output, _ := cmd.CombinedOutput()
	return string(output)
}

// 获取服务器内网ip
func GetServerLocalIp() string {
	cmd := exec.Command(`/bin/bash`, `-c`, `hostname -I`)
	output, _ := cmd.CombinedOutput()
	return gstr.Trim(string(output))
}

// go文件代码格式化
func GoFileFmt(filePath string) {
	fmtFuc := func(path, content string) string {
		res, err := imports.Process(path, []byte(content), nil)
		if err != nil {
			return content
		}
		return string(res)
	}
	gfile.ReplaceFileFunc(fmtFuc, filePath)
}

// 数据库表按时间做分区（通用，默认以分区最大日期作为分区名）
// 注意：如果表有唯一索引（含主键），则用于建立分区的字段必须和唯一索引字段组成联合索引
// 建议：分区间隔时间，分区数量设置后，两者总时长要比定时器间隔多几天时间，方便分区失败时，有时间让技术人工处理
// dbGroup			db分组
// dbTable			db表
// partitionNumber	当前时间后面，需要新增的分区数量
// partitionTime	间隔多长时间创建一个分区，单位：秒
// partitionField	分区字段，即根据该字段做分区
func DbTablePartition(ctx context.Context, dbGroup string, dbTable string, partitionNumber int64, partitionTime int64, partitionField string) (err error) {
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
	/**--------查询分区 开始--------**/
	partitionSelSql := `SELECT MAX(PARTITION_NAME) AS maxPartitionName FROM INFORMATION_SCHEMA.PARTITIONS WHERE TABLE_SCHEMA = SCHEMA() AND TABLE_NAME = '` + dbTable + `'`
	maxPartitionNameTmp, err := g.DB(dbGroup).GetValue(ctx, partitionSelSql) //不是分区表或无分区，查询结果都会有一项，且第一项maxPartitionName值为null
	if err != nil {
		return
	}
	maxPartitionName := maxPartitionNameTmp.String()
	/**--------查询分区 结束--------**/

	/**--------无分区则建立当前分区 开始--------**/
	if maxPartitionName == `` {
		ltTime := gtime.Now().Unix() + partitionTime
		ltDate := gtime.New(ltTime).Format(`Y-m-d`)
		partitionName := gtime.New(ltTime - 24*60*60).Format(`Ymd`) //该分区的最大日期作为分区名
		partitionCreateSql := "PARTITION `" + partitionName + "` VALUES LESS THAN (TO_DAYS( '" + ltDate + "' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0"
		partitionCreateSql = "ALTER TABLE `" + dbTable + "` PARTITION BY RANGE (TO_DAYS( " + partitionField + " )) ( " + partitionCreateSql + " )"
		_, err = g.DB(dbGroup).Exec(ctx, partitionCreateSql)
		if err != nil {
			return
		}
		maxPartitionName = partitionName
	}
	/**--------无分区则建立当前分区 结束--------**/

	/**--------检测需要创建的分区是否存在，没有则新增分区 开始--------**/
	partitionAddSqlList := []string{}
	nowTime := gtime.Now().Unix()
	maxPartitionTimeTmp, _ := gtime.StrToTimeFormat(maxPartitionName, `Ymd`)
	maxPartitionTime := maxPartitionTimeTmp.Unix()
	var i int64
	for i = 0; i < partitionNumber; i++ {
		//时间超过最大的分区后才开始新增分区，且以最大分区的时间开始往后计算
		if gtime.New(nowTime+(i+1)*partitionTime-24*60*60).Format(`Ymd`) > maxPartitionName {
			ltTime := maxPartitionTime + (i+1)*partitionTime
			ltDate := gtime.New(ltTime).Format(`Y-m-d`)
			partitionName := gtime.New(ltTime - 24*60*60).Format(`Ymd`) //该分区的最大日期作为分区名
			partitionAddSqlList = append(partitionAddSqlList, "PARTITION `"+partitionName+"` VALUES LESS THAN (TO_DAYS( '"+ltDate+"' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0")
		}

	}
	if len(partitionAddSqlList) > 0 {
		partitionAddSql := gstr.Join(partitionAddSqlList, `,`)
		partitionAddSql = "ALTER TABLE `" + dbTable + "` ADD PARTITION ( " + partitionAddSql + " )"
		_, err = g.DB(dbGroup).Exec(ctx, partitionAddSql)
		if err != nil {
			return
		}
	}
	/**--------检测需要创建的分区是否存在，没有则新增分区 结束--------**/
	return
}

// 列表转树状
func Tree(list g.List, id uint, priKey string, pidKey string) (tree g.List) {
	tree = g.List{}
	for k, v := range list {
		if gconv.Uint(v[pidKey]) == id {
			v[`children`] = Tree(list[(k+1):], gconv.Uint(v[priKey]), priKey, pidKey)
			tree = append(tree, v)
		}
	}
	return
}

// 保存关联表（一对多）。关联表除主表关联id外，只剩1个有用字段
func SaveArrRelMany(ctx context.Context, relDao dao.DaoInterface, idField string, valField string, id interface{}, valArr []interface{}) {
	valArrOfOldTmp, _ := relDao.CtxDaoModel(ctx).Filter(idField, id).Array(valField)
	valArrOfOld := gconv.SliceAny(valArrOfOldTmp)

	/**----新增关联 开始----**/
	insertValArr := gset.NewFrom(valArr).Diff(gset.NewFrom(valArrOfOld)).Slice()
	if len(insertValArr) > 0 {
		insertList := []map[string]interface{}{}
		for _, v := range insertValArr {
			insertList = append(insertList, map[string]interface{}{
				idField:  id,
				valField: v,
			})
		}
		relDao.CtxDaoModel(ctx).Data(insertList).Insert()
	}
	/**----新增关联 结束----**/

	/**----删除关联 开始----**/
	deleteValArr := gset.NewFrom(valArrOfOld).Diff(gset.NewFrom(valArr)).Slice()
	if len(deleteValArr) > 0 {
		relDao.CtxDaoModel(ctx).Filters(g.Map{
			idField:  id,
			valField: deleteValArr,
		}).Delete()
	}
	/**----删除关联 结束----**/
}

// 保存关联表（一对多），有顺序要求时使用。关联表除主表关联id外，只剩1个有用字段
func SaveArrRelManyWithSort(ctx context.Context, relDao dao.DaoInterface, idField string, valField string, idArr []interface{}, valArr []interface{}) {
	relDao.CtxDaoModel(ctx).Filter(idField, idArr).Delete()
	insertList := []map[string]interface{}{}
	for _, id := range idArr {
		for _, v := range valArr {
			insertList = append(insertList, map[string]interface{}{
				idField:  id,
				valField: v,
			})
		}
	}
	relDao.CtxDaoModel(ctx).Data(insertList).Insert()
}

// 保存关联表（一对多）。关联表除主表关联id外，至少还剩2个有用字段
func SaveListRelMany(ctx context.Context, relDao dao.DaoInterface, idField string, idSuffixFieldArr []string, id interface{}, valList []map[string]interface{}) {
	inStrArr := []string{}
	for _, v := range valList {
		saveItem := gjson.New(gjson.MustEncodeString(v)).Map()
		saveItem[idField] = id
		relDao.CtxDaoModel(ctx).Data(saveItem).Save()

		idArr := []string{gconv.String(id)}
		for _, idSuffixField := range idSuffixFieldArr {
			idArr = append(idArr, gconv.String(v[idSuffixField]))
		}
		inStrArr = append(inStrArr, `('`+gstr.Join(idArr, `', '`)+`')`)
	}
	relDao.CtxDaoModel(ctx).GetModel().Where(`(` + gstr.Join(append([]string{idField}, idSuffixFieldArr...), `, `) + `) NOT IN (` + gstr.Join(inStrArr, `, `) + `)`)
}

// 保存关联表（一对多），有顺序要求时使用。关联表除主表关联id外，至少还剩2个有用字段
func SaveListRelManyWithSort(ctx context.Context, relDao dao.DaoInterface, idField string, idArr []interface{}, valList []map[string]interface{}) {
	relDao.CtxDaoModel(ctx).Filter(idField, idArr).Delete()
	insertList := []map[string]interface{}{}
	for _, id := range idArr {
		for _, v := range valList {
			insertItem := gjson.New(gjson.MustEncodeString(v)).Map()
			insertItem[idField] = id
			insertList = append(insertList, insertItem)
		}
	}
	relDao.CtxDaoModel(ctx).Data(insertList).Insert()
}
