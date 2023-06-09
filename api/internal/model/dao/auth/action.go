// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/auth/internal"
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalActionDao is internal type for wrapping internal DAO implements.
type internalActionDao = *internal.ActionDao

// actionDao is the data access object for table auth_action.
// You can define custom methods on it to extend its functionality as you wish.
type actionDao struct {
	internalActionDao
}

var (
	// Action is globally public accessible object for table auth_action operations.
	Action = actionDao{
		internal.NewActionDao(),
	}
)

// 解析分库
func (daoThis *actionDao) ParseDbGroup(dbGroupSeldata map[string]interface{}) string {
	group := daoThis.Group()
	if len(dbGroupSeldata) > 0 { //分库逻辑
	}
	return group
}

// 解析分表
func (daoThis *actionDao) ParseDbTable(dbTableSelData map[string]interface{}) string {
	table := daoThis.Table()
	if len(dbTableSelData) > 0 { //分表逻辑
	}
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *actionDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
	switch len(dbSelDataList) {
	case 1:
		return g.DB(daoThis.ParseDbGroup(dbSelDataList[0])).Model(daoThis.Table()).Safe().Ctx(ctx)
	case 2:
		return g.DB(daoThis.ParseDbGroup(dbSelDataList[0])).Model(daoThis.ParseDbTable(dbSelDataList[1])).Safe().Ctx(ctx)
	default:
		return daoThis.Ctx(ctx)
	}
}

// 解析insert
func (daoThis *actionDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case `id`:
					insertData[index][daoThis.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
						continue
					}
					insertData[index][k] = v
				}
			}
		}
		if len(insertData) == 1 {
			m = m.Data(insertData[0])
		} else {
			m = m.Data(insertData)
		}
		return m
	}
}

// 解析update
func (daoThis *actionDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoThis.Table()+`.`+daoThis.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoThis.Table()+`.`+k] = v
			}
		}
		//m = m.Data(updateData) //字段被解析成`table.xxxx`，正确的应该是`table`.`xxxx`
		//解决字段被解析成`table.xxxx`的BUG
		fieldArr := []string{}
		valueArr := []interface{}{}
		for k, v := range updateData {
			_, ok := v.(gdb.Raw)
			if ok {
				fieldArr = append(fieldArr, k+` = `+gconv.String(v))
			} else {
				fieldArr = append(fieldArr, k+` = ?`)
				valueArr = append(valueArr, v)
			}
		}
		data := []interface{}{strings.Join(fieldArr, `,`)}
		data = append(data, valueArr...)
		m = m.Data(data...)
		return m
	}
}

// 解析field
func (daoThis *actionDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = daoThis.ParseJoin(`xxxx`, joinTableArr)(m)
			afterField = append(afterField, v) */
			case `id`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `sceneIdArr`:
				//需要id字段
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
				afterField = append(afterField, v)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(daoThis.Table() + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoThis.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoThis *actionDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case `id`, `idArr`:
				val := gconv.SliceInt(v)
				if len(val) == 1 {
					m = m.Where(daoThis.Table()+`.`+daoThis.PrimaryKey(), val[0])
				} else {
					m = m.Where(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
				}
			case `excId`, `excIdArr`:
				val := gconv.SliceInt(v)
				switch len(val) {
				case 0: //gconv.SliceInt会把0转换成[]int{}，故不能用转换后的val。必须用原始数据v
					m = m.WhereNot(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
				case 1:
					m = m.WhereNot(daoThis.Table()+`.`+daoThis.PrimaryKey(), val[0])
				default:
					m = m.WhereNotIn(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
				}
			case `startTime`:
				m = m.WhereGTE(daoThis.Table()+`.createdAt`, v)
			case `endTime`:
				m = m.WhereLTE(daoThis.Table()+`.createdAt`, v)
			case `keyword`:
				keywordField := strings.ReplaceAll(daoThis.PrimaryKey(), `Id`, `Name`)
				m = m.WhereLike(daoThis.Table()+`.`+keywordField, `%`+gconv.String(v)+`%`)
			case `sceneId`:
				m = m.Where(ActionRelToScene.Table()+`.`+k, v)

				m = daoThis.ParseJoin(`actionRelToScene`, joinTableArr)(m)
			case `selfAction`: //获取当前登录身份可用的操作。参数：map[string]interface{}{`sceneCode`: `场景标识`, `sceneId`=>场景id, `loginId`: 登录身份id}
				val := v.(map[string]interface{})

				m = m.Where(daoThis.Table()+`.isStop`, 0)
				m = m.Where(ActionRelToScene.Table()+`.sceneId`, val[`sceneId`])
				m = daoThis.ParseJoin(`actionRelToScene`, joinTableArr)(m)
				switch val[`sceneCode`].(string) {
				case `platformAdmin`:
					if gconv.Int(val[`loginId`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Int() { //平台超级管理员，不再需要其他条件
						return m
					}
					m = m.Where(Role.Table()+`.isStop`, 0)
					m = m.Where(RoleRelOfPlatformAdmin.Table()+`.adminId`, val[`loginId`])

					m = daoThis.ParseJoin(`roleRelToAction`, joinTableArr)(m)
					m = daoThis.ParseJoin(`role`, joinTableArr)(m)
					m = daoThis.ParseJoin(`roleRelOfPlatformAdmin`, joinTableArr)(m)
				}
				m = daoThis.ParseGroup([]string{`id`}, joinTableArr)(m)
			default:
				kArr := strings.Split(k, ` `) //支持`id > ?`等k
				if daoThis.ColumnArrG().Contains(kArr[0]) {
					if len(kArr) == 1 {
						if gstr.ToLower(gstr.SubStr(kArr[0], -2)) == `id` {
							val := gconv.SliceInt(v)
							if len(val) == 1 {
								m = m.Where(daoThis.Table()+`.`+k, val[0])
							} else {
								m = m.Where(daoThis.Table()+`.`+k, v)
							}
						} else if gstr.ToLower(gstr.SubStr(kArr[0], -4)) == `name` {
							m = m.WhereLike(daoThis.Table()+`.`+k, `%`+gconv.String(v)+`%`)
						} else {
							m = m.Where(daoThis.Table()+`.`+k, v)
						}
					} else {
						m = m.Where(daoThis.Table()+`.`+k, v)
					}
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoThis *actionDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoThis.Table() + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Group(daoThis.Table() + `.` + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *actionDao) ParseOrder(order [][2]string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case `id`:
				m = m.Order(daoThis.Table()+`.`+daoThis.PrimaryKey(), v[1])
			default:
				if daoThis.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoThis.Table()+`.`+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *actionDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case `xxxx`:
		xxxxTable := xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(xxxxTable) {
			*joinTableArr = append(*joinTableArr, xxxxTable)
			m = m.LeftJoin(xxxxTable, xxxxTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
		} */
		case `actionRelToScene`:
			actionRelToSceneTable := ActionRelToScene.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(actionRelToSceneTable) {
				*joinTableArr = append(*joinTableArr, actionRelToSceneTable)
				m = m.LeftJoin(actionRelToSceneTable, actionRelToSceneTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
			}
		case `roleRelToAction`:
			roleRelToActionTable := RoleRelToAction.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleRelToActionTable) {
				*joinTableArr = append(*joinTableArr, roleRelToActionTable)
				m = m.LeftJoin(roleRelToActionTable, roleRelToActionTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
			}
		case `role`:
			roleTable := Role.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleTable) {
				*joinTableArr = append(*joinTableArr, roleTable)
				roleRelToActionTable := RoleRelToAction.Table()
				m = m.LeftJoin(roleTable, roleTable+`.`+Role.PrimaryKey()+` = `+roleRelToActionTable+`.`+Role.PrimaryKey())
			}
		case `roleRelOfPlatformAdmin`:
			roleRelOfPlatformAdminTable := RoleRelOfPlatformAdmin.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleRelOfPlatformAdminTable) {
				*joinTableArr = append(*joinTableArr, roleRelOfPlatformAdminTable)
				roleRelToActionTable := RoleRelToAction.Table()
				m = m.LeftJoin(roleRelOfPlatformAdminTable, roleRelOfPlatformAdminTable+`.`+Role.PrimaryKey()+` = `+roleRelToActionTable+`.`+Role.PrimaryKey())
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoThis *actionDao) AfterField(afterField []string) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for i, record := range result {
				for _, v := range afterField {
					switch v {
					/* case `xxxx`:
					record[v] = gvar.New(``) */
					case `sceneIdArr`:
						sceneIdArr, _ := ActionRelToScene.ParseDbCtx(ctx).Where(`actionId`, record[daoThis.PrimaryKey()]).Array(`sceneId`)
						record[v] = gvar.New(sceneIdArr)
					}
				}
				result[i] = record
			}
			return
		},
	}
}

// Fill with you ideas below.

// 保存关联场景
func (daoThis *actionDao) SaveRelScene(ctx context.Context, sceneIdArr []int, id int) {
	sceneIdArrOfOldTmp, _ := ActionRelToScene.ParseDbCtx(ctx).Where(`actionId`, id).Array(`sceneId`)
	sceneIdArrOfOld := gconv.SliceInt(sceneIdArrOfOldTmp)

	/**----新增关联场景 开始----**/
	inserttSceneIdArr := gset.NewIntSetFrom(sceneIdArr).Diff(gset.NewIntSetFrom(sceneIdArrOfOld)).Slice()
	if len(inserttSceneIdArr) > 0 {
		insertList := []map[string]interface{}{}
		for _, v := range inserttSceneIdArr {
			insertList = append(insertList, map[string]interface{}{
				`actionId`: id,
				`sceneId`:  v,
			})
		}
		ActionRelToScene.ParseDbCtx(ctx).Data(insertList).Insert()
	}
	/**----新增关联场景 结束----**/

	/**----删除关联场景 开始----**/
	deleteSceneIdArr := gset.NewIntSetFrom(sceneIdArrOfOld).Diff(gset.NewIntSetFrom(sceneIdArr)).Slice()
	if len(deleteSceneIdArr) > 0 {
		ActionRelToScene.ParseDbCtx(ctx).Where(`actionId`, id).Where(`sceneId`, deleteSceneIdArr).Delete()
	}
	/**----删除关联场景 结束----**/
}
