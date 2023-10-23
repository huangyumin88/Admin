// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"context"
	"database/sql"
	"sync"

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
func (daoThis *actionDao) ParseDbGroup(ctx context.Context, dbGroupSelData ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupSelData) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *actionDao) ParseDbTable(ctx context.Context, dbTableSelData ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableSelData) > 0 {
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *actionDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
	switch len(dbSelDataList) {
	case 1:
		return g.DB(daoThis.ParseDbGroup(ctx, dbSelDataList[0])).Model(daoThis.ParseDbTable(ctx)).Ctx(ctx)
	case 2:
		return g.DB(daoThis.ParseDbGroup(ctx, dbSelDataList[0])).Model(daoThis.ParseDbTable(ctx, dbSelDataList[1])).Ctx(ctx)
	default:
		return g.DB(daoThis.ParseDbGroup(ctx)).Model(daoThis.ParseDbTable(ctx)).Ctx(ctx)
	}
}

// 解析insert
func (daoThis *actionDao) ParseInsert(insert map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		hookData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `id`:
				insertData[daoThis.PrimaryKey()] = v
			case `sceneIdArr`:
				hookData[k] = v
			default:
				if daoThis.ColumnArrG().Contains(k) {
					insertData[k] = v
				}
			}
		}
		m = m.Data(insertData)
		if len(hookData) > 0 {
			m = m.Hook(daoThis.HookInsert(hookData))
		}
		return m
	}
}

// hook insert
func (daoThis *actionDao) HookInsert(data map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range data {
				switch k {
				case `sceneIdArr`:
					daoThis.SaveRelScene(ctx, gconv.SliceInt(v), int(id))
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *actionDao) ParseUpdate(update map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoThis.Table()+`.`+daoThis.PrimaryKey()] = v
			default:
				if daoThis.ColumnArrG().Contains(k) {
					updateData[daoThis.Table()+`.`+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
				}
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
		data := []interface{}{gstr.Join(fieldArr, `,`)}
		data = append(data, valueArr...)
		m = m.Data(data...)
		return m
	}
}

// hook update
func (daoThis *actionDao) HookUpdate(data map[string]interface{}, idArr ...int) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			/* //不能这样拿idArr，联表时会有bug
			var idArr []*gvar.Var
			if len(data) > 0 {
				idArr, _ = daoThis.ParseDbCtx(ctx).Where(in.Condition, in.Args[len(in.Args)-gstr.Count(in.Condition, `?`):]...).Array(daoThis.PrimaryKey())
			} */
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			for k, v := range data {
				switch k {
				case `sceneIdArr`:
					relIdArr := gconv.SliceInt(v)
					for _, id := range idArr {
						daoThis.SaveRelScene(ctx, relIdArr, id)
					}
				}
			}

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */
			return
		},
	}
}

// hook delete
func (daoThis *actionDao) HookDelete(idArr ...int) gdb.HookHandler {
	return gdb.HookHandler{
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			ActionRelToScene.ParseDbCtx(ctx).Where(ActionRelToScene.Columns().ActionId, idArr).Delete()
			RoleRelToAction.ParseDbCtx(ctx).Where(RoleRelToAction.Columns().ActionId, idArr).Delete()
			return
		},
	}
}

// 解析field
func (daoThis *actionDao) ParseField(field []string, fieldWithParam map[string]interface{}, afterField *[]string, afterFieldWithParam map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = m.Handler(daoThis.ParseJoin(Xxxx.Table(), joinTableArr))
			*afterField = append(*afterField, v) */
			case `id`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.Columns().ActionName + ` AS ` + v)
			case `sceneIdArr`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
				*afterField = append(*afterField, v)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(daoThis.Table() + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		for k, v := range fieldWithParam {
			switch k {
			default:
				afterFieldWithParam[k] = v
			}
		}
		return m
	}
}

// hook select
func (daoThis *actionDao) HookSelect(afterField *[]string, afterFieldWithParam map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			var wg sync.WaitGroup
			for _, record := range result {
				wg.Add(1)
				go func(record gdb.Record) {
					defer wg.Done()
					for _, v := range *afterField {
						switch v {
						/* case `xxxx`:
						record[v] = gvar.New(``) */
						case `sceneIdArr`:
							idArr, _ := ActionRelToScene.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(ActionRelToScene.Columns().SceneId)
							record[v] = gvar.New(idArr)
						}
					}
					/* for k, v := range afterFieldWithParam {
						switch k {
						case `xxxx`:
							record[k] = gvar.New(v)
						}
					} */
				}(record)
			}
			wg.Wait()
			return
		},
	}
}

// 解析filter
func (daoThis *actionDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
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
			case `id`, `idArr`:
				m = m.Where(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
			case `label`:
				m = m.WhereLike(daoThis.Table()+`.`+daoThis.Columns().ActionName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().ActionName:
				m = m.WhereLike(daoThis.Table()+`.`+k, `%`+gconv.String(v)+`%`)
			case `timeRangeStart`:
				m = m.WhereGTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `sceneId`:
				m = m.Where(ActionRelToScene.Table()+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(ActionRelToScene.Table(), joinTableArr))
			case `selfAction`: //获取当前登录身份可用的操作。参数：map[string]interface{}{`sceneCode`: `场景标识`, `sceneId`=>场景id, `loginId`: 登录身份id}
				val := gconv.Map(v)
				m = m.Where(daoThis.Table()+`.`+daoThis.Columns().IsStop, 0)
				m = m.Where(ActionRelToScene.Table()+`.`+ActionRelToScene.Columns().SceneId, val[`sceneId`])
				m = m.Handler(daoThis.ParseJoin(ActionRelToScene.Table(), joinTableArr))
				switch gconv.String(val[`sceneCode`]) {
				case `platform`:
					if gconv.Int(val[`loginId`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Int() { //平台超级管理员，不再需要其它条件
						continue
					}
					m = m.Where(Role.Table()+`.`+Role.Columns().IsStop, 0)
					m = m.Handler(daoThis.ParseJoin(RoleRelToAction.Table(), joinTableArr))
					m = m.Handler(daoThis.ParseJoin(Role.Table(), joinTableArr))

					m = m.Where(RoleRelOfPlatformAdmin.Table()+`.`+RoleRelOfPlatformAdmin.Columns().AdminId, val[`loginId`])
					m = m.Handler(daoThis.ParseJoin(RoleRelOfPlatformAdmin.Table(), joinTableArr))
				default:
					m = m.Where(`1 = 0`)
				}
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Where(daoThis.Table()+`.`+k, v)
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
func (daoThis *actionDao) ParseOrder(order []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			k := gstr.Split(v, ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoThis.Table() + `.` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Order(daoThis.Table() + `.` + v)
				} else {
					m = m.Order(v)
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *actionDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if garray.NewStrArrayFrom(*joinTableArr).Contains(joinCode) {
			return m
		}
		*joinTableArr = append(*joinTableArr, joinCode)
		switch joinCode {
		/* case Xxxx.Table():
		m = m.LeftJoin(joinCode, joinCode+`.`+Xxxx.Columns().XxxxId+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.Table()+` AS `+joinCode, joinCode+`.`+Xxxx.Columns().XxxxId+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey()) */
		case ActionRelToScene.Table():
			m = m.LeftJoin(joinCode, joinCode+`.`+ActionRelToScene.Columns().ActionId+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
		case RoleRelToAction.Table():
			m = m.LeftJoin(joinCode, joinCode+`.`+RoleRelToAction.Columns().ActionId+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
		case Role.Table():
			m = m.LeftJoin(joinCode, joinCode+`.`+Role.PrimaryKey()+` = `+RoleRelToAction.Table()+`.`+RoleRelToAction.Columns().RoleId)
		case RoleRelOfPlatformAdmin.Table():
			m = m.LeftJoin(joinCode, joinCode+`.`+RoleRelOfPlatformAdmin.Columns().RoleId+` = `+RoleRelToAction.Table()+`.`+RoleRelToAction.Columns().RoleId)
		}
		return m
	}
}

// Fill with you ideas below.

// 保存关联场景
func (daoThis *actionDao) SaveRelScene(ctx context.Context, relIdArr []int, id int) {
	relDao := ActionRelToScene
	priKey := relDao.Columns().ActionId
	relKey := relDao.Columns().SceneId
	relIdArrOfOldTmp, _ := relDao.ParseDbCtx(ctx).Where(priKey, id).Array(relKey)
	relIdArrOfOld := gconv.SliceInt(relIdArrOfOldTmp)

	/**----新增关联 开始----**/
	insertRelIdArr := gset.NewIntSetFrom(relIdArr).Diff(gset.NewIntSetFrom(relIdArrOfOld)).Slice()
	if len(insertRelIdArr) > 0 {
		insertList := []map[string]interface{}{}
		for _, v := range insertRelIdArr {
			insertList = append(insertList, map[string]interface{}{
				priKey: id,
				relKey: v,
			})
		}
		relDao.ParseDbCtx(ctx).Data(insertList).Insert()
	}
	/**----新增关联 结束----**/

	/**----删除关联 开始----**/
	deleteRelIdArr := gset.NewIntSetFrom(relIdArrOfOld).Diff(gset.NewIntSetFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		relDao.ParseDbCtx(ctx).Where(priKey, id).Where(relKey, deleteRelIdArr).Delete()
	}
	/**----删除关联 结束----**/
}
