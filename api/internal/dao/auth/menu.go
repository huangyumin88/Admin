// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalMenuDao is internal type for wrapping internal DAO implements.
type internalMenuDao = *internal.MenuDao

// menuDao is the data access object for table auth_menu.
// You can define custom methods on it to extend its functionality as you wish.
type menuDao struct {
	internalMenuDao
}

var (
	// Menu is globally public accessible object for table auth_menu operations.
	Menu = menuDao{
		internal.NewMenuDao(),
	}
)

// 解析分库
func (daoThis *menuDao) ParseDbGroup(ctx context.Context, dbGroupSelData ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupSelData) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *menuDao) ParseDbTable(ctx context.Context, dbTableSelData ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableSelData) > 0 {
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *menuDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
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
func (daoThis *menuDao) ParseInsert(insert map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		hookData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `id`:
				insertData[daoThis.PrimaryKey()] = v
			case daoThis.Columns().Pid:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), v).Fields(daoThis.Columns().IdPath, daoThis.Columns().Level).One()
					hookData[`pIdPath`] = pInfo[daoThis.Columns().IdPath].String()
					hookData[`pLevel`] = pInfo[daoThis.Columns().Level].Uint()
				} else {
					hookData[`pIdPath`] = `0`
					hookData[`pLevel`] = 0
				}
			case daoThis.Columns().ExtraData:
				insertData[k] = v
				if gconv.String(v) == `` {
					insertData[k] = nil
				}
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
func (daoThis *menuDao) HookInsert(data map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			updateSelfData := map[string]interface{}{}
			for k, v := range data {
				switch k {
				case `pIdPath`:
					updateSelfData[daoThis.Columns().IdPath] = gconv.String(v) + `-` + gconv.String(id)
				case `pLevel`:
					updateSelfData[daoThis.Columns().Level] = gconv.Uint(v) + 1
				}
			}
			if len(updateSelfData) > 0 {
				daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).Data(updateSelfData).Update()
			}
			return
		},
	}
}

// 解析update
func (daoThis *menuDao) ParseUpdate(update map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[tableThis+`.`+daoThis.PrimaryKey()] = v
			case daoThis.Columns().Pid:
				updateData[tableThis+`.`+k] = v
				pIdPath := `0`
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), v).Fields(daoThis.Columns().IdPath, daoThis.Columns().Level).One()
					pIdPath = pInfo[daoThis.Columns().IdPath].String()
					pLevel = pInfo[daoThis.Columns().Level].Uint()
				}
				updateData[tableThis+`.`+daoThis.Columns().IdPath] = gdb.Raw(`CONCAT('` + pIdPath + `-', ` + daoThis.PrimaryKey() + `)`)
				updateData[tableThis+`.`+daoThis.Columns().Level] = pLevel + 1
			case daoThis.Columns().ExtraData:
				updateData[tableThis+`.`+k] = gvar.New(v)
				if gconv.String(v) == `` {
					updateData[tableThis+`.`+k] = nil
				}
			default:
				if daoThis.ColumnArrG().Contains(k) {
					updateData[tableThis+`.`+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
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
func (daoThis *menuDao) HookUpdate(data map[string]interface{}, idArr ...uint) gdb.HookHandler {
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

			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			for k, v := range data {
				switch k {
				case `updateChildIdPathAndLevelList`: //修改pid时，更新所有子孙级的idPath和level。参数：[]map[string]interface{}{newIdPath: 父级新idPath, oldIdPath: 父级旧idPath, newLevel: 父级新level, oldLevel: 父级旧level}
					val := v.([]map[string]interface{})
					for _, v1 := range val {
						daoThis.UpdateChildIdPathAndLevel(ctx, gconv.String(v1[`newIdPath`]), gconv.String(v1[`oldIdPath`]), gconv.Uint(v1[`newLevel`]), gconv.Uint(v1[`oldLevel`]))
					}
				}
			}
			return
		},
	}
}

// hook delete
func (daoThis *menuDao) HookDelete(idArr ...uint) gdb.HookHandler {
	return gdb.HookHandler{
		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */
			return
		},
	}
}

// 解析field
func (daoThis *menuDao) ParseField(field []string, fieldWithParam map[string]interface{}, afterField *[]string, afterFieldWithParam map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = m.Handler(daoThis.ParseJoin(Xxxx.ParseDbTable(ctx), joinTableArr))
			*afterField = append(*afterField, v) */
			case `id`:
				m = m.Fields(tableThis + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(tableThis + `.` + daoThis.Columns().MenuName + ` AS ` + v)
			case Scene.Columns().SceneName: //因前端页面已用该字段名显示，故不存在时改成`sceneName`（控制器也要改）。同时下面Fields方法改成m = m.Fields(tableScene + `.` + Scene.Columns().Xxxx + ` AS ` + v)
				tableScene := Scene.ParseDbTable(ctx)
				m = m.Fields(tableScene + `.` + v)
				m = m.Handler(daoThis.ParseJoin(tableScene, joinTableArr))
			case `pMenuName`:
				tableP := `p_` + tableThis
				m = m.Fields(tableP + `.` + daoThis.Columns().MenuName + ` AS ` + v)
				m = m.Handler(daoThis.ParseJoin(tableP, joinTableArr))
			case `tree`:
				m = m.Fields(tableThis + `.` + daoThis.PrimaryKey())
				m = m.Fields(tableThis + `.` + daoThis.Columns().Pid)
				m = m.Handler(daoThis.ParseOrder([]string{`tree`}, joinTableArr))
			case `showMenu`: //前端显示菜单需要以下字段，且title需要转换
				m = m.Fields(tableThis + `.` + daoThis.Columns().MenuName)
				m = m.Fields(tableThis + `.` + daoThis.Columns().MenuIcon)
				m = m.Fields(tableThis + `.` + daoThis.Columns().MenuUrl)
				m = m.Fields(tableThis + `.` + daoThis.Columns().ExtraData)
				// m = m.Fields(tableThis + `.` + daoThis.Columns().ExtraData + `->'$.i18n' AS i18n`)	//mysql5.6版本不支持
				// m = m.Fields(gdb.Raw(`JSON_UNQUOTE(JSON_EXTRACT(` + daoThis.Columns().ExtraData + `, \`$.i18n\`)) AS i18n`))	//mysql不能直接转成对象返回
				*afterField = append(*afterField, v)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(tableThis + `.` + v)
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
func (daoThis *menuDao) HookSelect(afterField *[]string, afterFieldWithParam map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range *afterField {
					switch v {
					case `showMenu`:
						extraDataJson := gjson.New(record[daoThis.Columns().ExtraData])
						record[`i18n`] = extraDataJson.Get(`i18n`)
						if record[`i18n`] == nil {
							record[`i18n`] = gvar.New(map[string]interface{}{`title`: map[string]interface{}{`zh-cn`: record[`menuName`]}})
						}
					default:
						record[v] = gvar.New(nil)
					}
				}
				/* for k, v := range afterFieldWithParam {
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
func (daoThis *menuDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for k, v := range filter {
			switch k {
			case `excId`, `excIdArr`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(tableThis+`.`+daoThis.PrimaryKey(), v)
				} else {
					m = m.WhereNot(tableThis+`.`+daoThis.PrimaryKey(), v)
				}
			case `id`, `idArr`:
				m = m.Where(tableThis+`.`+daoThis.PrimaryKey(), v)
			case `label`:
				m = m.WhereLike(tableThis+`.`+daoThis.Columns().MenuName, `%`+gconv.String(v)+`%`)
			case daoThis.Columns().MenuName:
				m = m.WhereLike(tableThis+`.`+k, `%`+gconv.String(v)+`%`)
			case `timeRangeStart`:
				m = m.WhereGTE(tableThis+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(tableThis+`.`+daoThis.Columns().CreatedAt, v)
			case `selfMenu`: //获取当前登录身份可用的菜单。参数：map[string]interface{}{`sceneCode`: `场景标识`, `sceneId`: 场景id, `loginId`: 登录身份id}
				val := gconv.Map(v)
				m = m.Where(tableThis+`.`+daoThis.Columns().SceneId, val[`sceneId`])
				m = m.Where(tableThis+`.`+daoThis.Columns().IsStop, 0)
				switch gconv.String(val[`sceneCode`]) {
				case `platform`:
					if gconv.Uint(val[`loginId`]) == g.Cfg().MustGet(ctx, `superPlatformAdminId`).Uint() { //平台超级管理员，不再需要其它条件
						continue
					}
					tableRole := Role.ParseDbTable(ctx)
					tableRoleRelToMenu := RoleRelToMenu.ParseDbTable(ctx)
					m = m.Where(tableRole+`.`+Role.Columns().IsStop, 0)
					m = m.Handler(daoThis.ParseJoin(tableRoleRelToMenu, joinTableArr))
					m = m.Handler(daoThis.ParseJoin(tableRole, joinTableArr))

					tableRoleRelOfPlatformAdmin := RoleRelOfPlatformAdmin.ParseDbTable(ctx)
					m = m.Where(tableRoleRelOfPlatformAdmin+`.`+RoleRelOfPlatformAdmin.Columns().AdminId, val[`loginId`])
					m = m.Handler(daoThis.ParseJoin(tableRoleRelOfPlatformAdmin, joinTableArr))
				default:
					m = m.Where(`1 = 0`)
				}
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Where(tableThis+`.`+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoThis *menuDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for _, v := range group {
			switch v {
			case `id`:
				m = m.Group(tableThis + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Group(tableThis + `.` + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoThis *menuDao) ParseOrder(order []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for _, v := range order {
			v = gstr.Trim(v)
			k := gstr.Split(v, ` `)[0]
			switch k {
			case `id`:
				m = m.Order(tableThis + `.` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))
			case `tree`:
				m = m.OrderAsc(tableThis + `.` + daoThis.Columns().Pid)
				m = m.OrderAsc(tableThis + `.` + daoThis.Columns().Sort)
				m = m.OrderAsc(tableThis + `.` + daoThis.PrimaryKey())
			case daoThis.Columns().Level:
				m = m.Order(tableThis + `.` + v)
				m = m.OrderDesc(tableThis + `.` + daoThis.PrimaryKey())
			case daoThis.Columns().Sort:
				m = m.Order(tableThis + `.` + v)
				m = m.OrderDesc(tableThis + `.` + daoThis.PrimaryKey())
			default:
				if daoThis.ColumnArrG().Contains(k) {
					m = m.Order(tableThis + `.` + v)
				} else {
					m = m.Order(v)
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *menuDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if garray.NewStrArrayFrom(*joinTableArr).Contains(joinCode) {
			return m
		}
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		*joinTableArr = append(*joinTableArr, joinCode)
		switch joinCode {
		/* case Xxxx.ParseDbTable(ctx):
		m = m.LeftJoin(joinCode, joinCode+`.`+Xxxx.Columns().XxxxId+` = `+tableThis+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(ctx)+` AS `+joinCode, joinCode+`.`+Xxxx.Columns().XxxxId+` = `+tableThis+`.`+daoThis.PrimaryKey()) */
		case Scene.ParseDbTable(ctx):
			m = m.LeftJoin(joinCode, joinCode+`.`+Scene.PrimaryKey()+` = `+tableThis+`.`+daoThis.Columns().SceneId)
		case `p_` + tableThis:
			m = m.LeftJoin(tableThis+` AS `+joinCode, joinCode+`.`+daoThis.PrimaryKey()+` = `+tableThis+`.`+daoThis.Columns().Pid)
		case Role.ParseDbTable(ctx):
			m = m.LeftJoin(joinCode, joinCode+`.`+Role.PrimaryKey()+` = `+RoleRelToMenu.ParseDbTable(ctx)+`.`+RoleRelToMenu.Columns().RoleId)
		case RoleRelOfPlatformAdmin.ParseDbTable(ctx):
			m = m.LeftJoin(joinCode, joinCode+`.`+RoleRelOfPlatformAdmin.Columns().RoleId+` = `+RoleRelToMenu.ParseDbTable(ctx)+`.`+RoleRelToMenu.Columns().RoleId)
		/* case RoleRelToMenu.ParseDbTable(ctx):
		m = m.LeftJoin(joinCode+` AS `+joinCode, joinCode+`.`+RoleRelToMenu.Columns().MenuId+` = `+tableThis+`.`+daoThis.PrimaryKey()) */
		default:
			m = m.LeftJoin(joinCode, joinCode+`.`+daoThis.PrimaryKey()+` = `+tableThis+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.

// 修改pid时，更新所有子孙级的idPath和level
func (daoThis *menuDao) UpdateChildIdPathAndLevel(ctx context.Context, newIdPath string, oldIdPath string, newLevel uint, oldLevel uint) {
	data := g.Map{
		daoThis.Columns().IdPath: gdb.Raw(`REPLACE(` + daoThis.Columns().IdPath + `, '` + oldIdPath + `', '` + newIdPath + `')`),
		daoThis.Columns().Level:  gdb.Raw(daoThis.Columns().Level + ` + ` + gconv.String(newLevel-oldLevel)),
	}
	if newLevel < oldLevel {
		data[daoThis.Columns().Level] = gdb.Raw(daoThis.Columns().Level + ` - ` + gconv.String(oldLevel-newLevel))
	}
	daoThis.ParseDbCtx(ctx).WhereLike(daoThis.Columns().IdPath, oldIdPath+`-%`).Data(data).Update()
}
