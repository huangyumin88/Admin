// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	daoIndex "api/internal/dao/handler"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalRoleDao is internal type for wrapping internal DAO implements.
type internalRoleDao = *internal.RoleDao

// roleDao is the data access object for table auth_role.
// You can define custom methods on it to extend its functionality as you wish.
type roleDao struct {
	internalRoleDao
}

var (
	// Role is globally public accessible object for table auth_role operations.
	Role = roleDao{
		internal.NewRoleDao(),
	}
)

// 获取daoHandler
func (daoThis *roleDao) HandlerCtx(ctx context.Context, dbOpt ...map[string]interface{}) *daoIndex.DaoHandler {
	return daoIndex.NewDaoHandler(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *roleDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *roleDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *roleDao) ParseDbCtx(ctx context.Context, dbOpt ...map[string]interface{}) *gdb.Model {
	switch len(dbOpt) {
	case 1:
		return g.DB(daoThis.ParseDbGroup(ctx, dbOpt[0])).Model(daoThis.ParseDbTable(ctx)).Ctx(ctx)
	case 2:
		return g.DB(daoThis.ParseDbGroup(ctx, dbOpt[0])).Model(daoThis.ParseDbTable(ctx, dbOpt[1])).Ctx(ctx)
	default:
		return g.DB(daoThis.ParseDbGroup(ctx)).Model(daoThis.ParseDbTable(ctx)).Ctx(ctx)
	}
}

// 解析insert
func (daoThis *roleDao) ParseInsert(insert map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `id`:
				insertData[daoThis.PrimaryKey()] = v
			case `menuIdArr`, `actionIdArr`:
				daoHandler.AfterInsert[k] = v
			default:
				if daoThis.ColumnArrG().Contains(k) {
					insertData[k] = v
				}
			}
		}
		m = m.Data(insertData)
		if len(daoHandler.AfterInsert) > 0 {
			m = m.Hook(daoThis.HookInsert(daoHandler))
		}
		return m
	}
}

// hook insert
func (daoThis *roleDao) HookInsert(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range daoHandler.AfterInsert {
				switch k {
				case `menuIdArr`:
					daoThis.SaveRelMenu(ctx, gconv.SliceUint(v), uint(id))
				case `actionIdArr`:
					daoThis.SaveRelAction(ctx, gconv.SliceUint(v), uint(id))
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *roleDao) ParseUpdate(update map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoHandler.DbTable+`.`+daoThis.PrimaryKey()] = v
			default:
				if daoThis.ColumnArrG().Contains(k) {
					updateData[daoHandler.DbTable+`.`+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
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
func (daoThis *roleDao) HookUpdate(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			for k, v := range daoHandler.AfterUpdate {
				switch k {
				case `menuIdArr`:
					relIdArr := gconv.SliceUint(v)
					for _, id := range daoHandler.IdArr {
						daoThis.SaveRelMenu(ctx, relIdArr, id)
					}
				case `actionIdArr`:
					relIdArr := gconv.SliceUint(v)
					for _, id := range daoHandler.IdArr {
						daoThis.SaveRelAction(ctx, relIdArr, id)
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
func (daoThis *roleDao) HookDelete(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
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

			RoleRelToMenu.ParseDbCtx(ctx).Where(RoleRelToMenu.Columns().RoleId, daoHandler.IdArr).Delete()
			RoleRelToAction.ParseDbCtx(ctx).Where(RoleRelToAction.Columns().RoleId, daoHandler.IdArr).Delete()
			RoleRelOfPlatformAdmin.ParseDbCtx(ctx).Where(RoleRelOfPlatformAdmin.Columns().RoleId, daoHandler.IdArr).Delete()
			return
		},
	}
}

// 解析field
func (daoThis *roleDao) ParseField(field []string, fieldWithParam map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = m.Handler(daoThis.ParseJoin(Xxxx.ParseDbTable(ctx), daoHandler))
			daoHandler.AfterField = append(daoHandler.AfterField, v) */
			case `id`:
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.Columns().RoleName + ` AS ` + v)
			case Scene.Columns().SceneName: //因前端页面已用该字段名显示，故不存在时改成`sceneName`（控制器也要改）。同时下面Fields方法改成m = m.Fields(tableScene + `.` + Scene.Columns().Xxxx + ` AS ` + v)
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Fields(tableScene + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoHandler))
			case `menuIdArr`, `actionIdArr`:
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.PrimaryKey())
				daoHandler.AfterField = append(daoHandler.AfterField, v)
			case `tableName`:
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.Columns().TableId)
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Fields(tableScene + `.` + Scene.Columns().SceneCode)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoHandler))
				daoHandler.AfterField = append(daoHandler.AfterField, v)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(daoHandler.DbTable + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		for k, v := range fieldWithParam {
			switch k {
			default:
				daoHandler.AfterFieldWithParam[k] = v
			}
		}
		return m
	}
}

// hook select
func (daoThis *roleDao) HookSelect(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range daoHandler.AfterField {
					switch v {
					case `menuIdArr`:
						idArr, _ := RoleRelToMenu.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(RoleRelToMenu.Columns().MenuId)
						record[v] = gvar.New(idArr)
					case `actionIdArr`:
						idArr, _ := RoleRelToAction.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(RoleRelToAction.Columns().ActionId)
						record[v] = gvar.New(idArr)
					case `tableName`:
						if record[daoThis.Columns().TableId].Uint() == 0 {
							record[v] = gvar.New(`平台`)
							continue
						}
						switch record[Scene.Columns().SceneCode].String() {
						case `platform`:
						}
					default:
						record[v] = gvar.New(nil)
					}
				}
				/* for k, v := range daoHandler.AfterFieldWithParam {
					switch k {
					case `xxxx`:
						record[k] = gvar.New(v)
					}
				} */
			}
			return
		},
	}
}

// 解析filter
func (daoThis *roleDao) ParseFilter(filter map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case `excId`, `excIdArr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoHandler.DbTable+`.`+daoThis.PrimaryKey(), v)
				} else {
					m = m.WhereNot(daoHandler.DbTable+`.`+daoThis.PrimaryKey(), v)
				}
			case `id`, `idArr`:
				m = m.Where(daoHandler.DbTable+`.`+daoThis.PrimaryKey(), v)
			case `label`:
				m = m.WhereLike(daoHandler.DbTable+`.`+daoThis.Columns().RoleName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().RoleName:
				m = m.WhereLike(daoHandler.DbTable+`.`+k, `%`+gconv.String(v)+`%`)
			case `timeRangeStart`:
				m = m.WhereGTE(daoHandler.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(daoHandler.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case Scene.Columns().SceneCode:
				tableScene := Scene.ParseDbTable(m.GetCtx())
				m = m.Where(tableScene+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableScene, daoHandler))
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Where(daoHandler.DbTable+`.`+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoThis *roleDao) ParseGroup(group []string, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(daoHandler.DbTable + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Group(daoHandler.DbTable + `.` + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *roleDao) ParseOrder(order []string, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			v = gstr.Trim(v)
			k := gstr.Split(v, ` `)[0]
			switch k {
			case `id`:
				m = m.Order(daoHandler.DbTable + `.` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Order(daoHandler.DbTable + `.` + v)
				} else {
					m = m.Order(v)
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *roleDao) ParseJoin(joinTable string, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if garray.NewStrArrayFrom(daoHandler.JoinTableArr).Contains(joinTable) {
			return m
		}
		daoHandler.JoinTableArr = append(daoHandler.JoinTableArr, joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey()) */
		case Scene.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`.`+Scene.PrimaryKey()+` = `+daoHandler.DbTable+`.`+daoThis.Columns().SceneId)
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.

// 保存关联菜单
func (daoThis *roleDao) SaveRelMenu(ctx context.Context, relIdArr []uint, id uint) {
	relDao := RoleRelToMenu
	priKey := relDao.Columns().RoleId
	relKey := relDao.Columns().MenuId
	relIdArrOfOldTmp, _ := relDao.ParseDbCtx(ctx).Where(priKey, id).Array(relKey)
	relIdArrOfOld := gconv.SliceUint(relIdArrOfOldTmp)

	/**----新增关联 开始----**/
	insertRelIdArr := gset.NewFrom(relIdArr).Diff(gset.NewFrom(relIdArrOfOld)).Slice()
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
	deleteRelIdArr := gset.NewFrom(relIdArrOfOld).Diff(gset.NewFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		relDao.ParseDbCtx(ctx).Where(priKey, id).Where(relKey, deleteRelIdArr).Delete()
	}
	/**----删除关联 结束----**/
}

// 保存关联操作
func (daoThis *roleDao) SaveRelAction(ctx context.Context, relIdArr []uint, id uint) {
	relDao := RoleRelToAction
	priKey := relDao.Columns().RoleId
	relKey := relDao.Columns().ActionId
	relIdArrOfOldTmp, _ := relDao.ParseDbCtx(ctx).Where(priKey, id).Array(relKey)
	relIdArrOfOld := gconv.SliceUint(relIdArrOfOldTmp)

	/**----新增关联 开始----**/
	insertRelIdArr := gset.NewFrom(relIdArr).Diff(gset.NewFrom(relIdArrOfOld)).Slice()
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
	/**----新增关联 开始----**/

	/**----删除关联 结束----**/
	deleteRelIdArr := gset.NewFrom(relIdArrOfOld).Diff(gset.NewFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		relDao.ParseDbCtx(ctx).Where(priKey, id).Where(relKey, deleteRelIdArr).Delete()
	}
	/**----删除关联 结束----**/
}
