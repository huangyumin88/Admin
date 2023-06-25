// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"api/internal/utils"
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
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
func (daoThis *menuDao) ParseDbGroup(dbGroupSeldata map[string]interface{}) string {
	group := daoThis.Group()
	if len(dbGroupSeldata) > 0 { //分库逻辑
	}
	return group
}

// 解析分表
func (daoThis *menuDao) ParseDbTable(dbTableSelData map[string]interface{}) string {
	table := daoThis.Table()
	if len(dbTableSelData) > 0 { //分表逻辑
	}
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *menuDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
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
func (daoThis *menuDao) ParseInsert(insert map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		hookData := map[string]interface{}{}
		insertData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `id`:
				insertData[daoThis.PrimaryKey()] = v
			case `pid`:
				insertData[k] = v
				if gconv.Int(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), v).Fields(daoThis.Columns().IdPath, daoThis.Columns().Level).One()
					hookData[`pIdPath`] = pInfo[daoThis.Columns().IdPath].String()
					hookData[`pLevel`] = pInfo[daoThis.Columns().Level].Int()
				} else {
					hookData[`pIdPath`] = `0`
					hookData[`pLevel`] = 0
				}
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
					continue
				}
				insertData[k] = v
			}
		}
		m = m.Data(insertData)
		m = m.Hook(daoThis.HookInsert(hookData))
		return m
	}
}

// hook insert
func (daoThis *menuDao) HookInsert(data map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
				if len(match) > 0 {
					err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
				}
				return
			}
			id, _ := result.LastInsertId()
			updateSelfData := map[string]interface{}{}
			for k, v := range data {
				switch k {
				case `pIdPath`:
					updateSelfData[daoThis.Columns().IdPath] = gconv.String(v) + `-` + gconv.String(id)
				case `pLevel`:
					updateSelfData[daoThis.Columns().Level] = gconv.Int(v) + 1
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
func (daoThis *menuDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		hookData := map[string]interface{}{}
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoThis.Table()+`.`+daoThis.PrimaryKey()] = v
			case `pid`:
				updateData[daoThis.Table()+`.`+k] = v
				pIdPath := `0`
				pLevel := 0
				if gconv.Int(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), v).Fields(daoThis.Columns().IdPath, daoThis.Columns().Level).One()
					pIdPath = pInfo[daoThis.Columns().IdPath].String()
					pLevel = pInfo[daoThis.Columns().Level].Int()
				}
				updateData[daoThis.Table()+`.`+daoThis.Columns().IdPath] = gdb.Raw(`CONCAT('` + pIdPath + `-', ` + daoThis.PrimaryKey() + `)`)
				updateData[daoThis.Table()+`.`+daoThis.Columns().Level] = pLevel + 1
			case `idPathOfChild`: //更新所有子孙级的idPath。参数：map[string]interface{}{`newVal`: `父级新idPath`, `oldVal`:`父级旧idPath`}
				val := gconv.Map(v)
				updateData[daoThis.Table()+`.`+daoThis.Columns().IdPath] = gdb.Raw(`REPLACE(` + daoThis.Table() + `.` + daoThis.Columns().IdPath + `, '` + gconv.String(val[`oldVal`]) + `', '` + gconv.String(val[`newVal`]) + `')`)
			case `levelOfChild`: //更新所有子孙级的level。参数：map[string]interface{}{`newVal`: 父级新level, `oldVal`:父级旧level}
				val := gconv.Map(v)
				updateData[daoThis.Table()+`.`+daoThis.Columns().Level] = gdb.Raw(daoThis.Table() + `.` + daoThis.Columns().Level + ` + ` + gconv.String(gconv.Int(val[`newVal`])-gconv.Int(val[`oldVal`])))
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoThis.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoThis.Table()+`.`+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
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
		m = m.Hook(daoThis.HookUpdate(hookData))
		return m
	}
}

// hook update
func (daoThis *menuDao) HookUpdate(data map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			/* var idArr []*gvar.Var
			if len(data) > 0 {
				idArr, _ = daoThis.ParseDbCtx(ctx).Where(in.Condition, in.Args[len(in.Args)-gstr.Count(in.Condition, `?`):]...).Array(daoThis.PrimaryKey())
			} */
			result, err = in.Next(ctx)
			if err != nil {
				match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
				if len(match) > 0 {
					err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
				}
				return
			}
			row, _ := result.RowsAffected()

			/* for k, v := range data {
				switch k {
				case `xxxx`:
					xxxx := gconv.SliceInt(v)
					for _, id := range idArr {
					}
					row = 1 //有可能只改menuIdArr
				}
			} */

			if row == 0 {
				err = utils.NewErrorCode(ctx, 99999999, ``)
			}
			return
		},
	}
}

// 解析field
func (daoThis *menuDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = daoThis.ParseJoin(Xxxx.Table(), joinTableArr)(m)
			afterField = append(afterField, v) */
			case `id`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `name`:
				nameField := gstr.SubStr(gstr.CaseCamel(daoThis.PrimaryKey()), 0, -2) + `Name`
				if daoThis.ColumnArrG().Contains(gstr.CaseCamelLower(nameField)) {
					m = m.Fields(daoThis.Table() + `.` + gstr.CaseCamelLower(nameField) + ` AS ` + v)
				} else if daoThis.ColumnArrG().Contains(gstr.CaseSnakeFirstUpper(nameField)) {
					m = m.Fields(daoThis.Table() + `.` + gstr.CaseSnakeFirstUpper(nameField) + ` AS ` + v)
				}
			case `menuTree`: //树状需要以下字段和排序方式
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
				m = m.Fields(daoThis.Table() + `.` + daoThis.Columns().Pid)

				m = daoThis.ParseOrder([]string{`menuTree`}, joinTableArr)(m) //排序方式
			case `showMenu`: //前端显示菜单需要以下字段，且title需要转换
				m = m.Fields(daoThis.Table() + `.` + daoThis.Columns().MenuName)
				m = m.Fields(daoThis.Table() + `.` + daoThis.Columns().MenuIcon)
				m = m.Fields(daoThis.Table() + `.` + daoThis.Columns().MenuUrl)
				m = m.Fields(daoThis.Table() + `.` + daoThis.Columns().ExtraData + `->'$.i18n' AS i18n`)
				//m = m.Fields(gdb.Raw(`JSON_UNQUOTE(JSON_EXTRACT(` + daoThis.Columns().ExtraData + `, \`$.i18n\`)) AS i18n`))//mysql不能直接转成对象返回
				afterField = append(afterField, v)
			case `sceneName`:
				m = m.Fields(Scene.Table() + `.` + v)
				m = daoThis.ParseJoin(Scene.Table(), joinTableArr)(m)
			case `pMenuName`:
				m = m.Fields(`p_` + daoThis.Table() + `.` + daoThis.Columns().MenuName + ` AS ` + v)
				m = daoThis.ParseJoin(`p_`+daoThis.Table(), joinTableArr)(m)
			default:
				if daoThis.ColumnArrG().Contains(v) {
					m = m.Fields(daoThis.Table() + `.` + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoThis.HookSelect(afterField))
		}
		return m
	}
}

// hook select
func (daoThis *menuDao) HookSelect(afterField []string) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for index, record := range result {
				for _, v := range afterField {
					switch v {
					/* case `xxxx`:
					record[v] = gvar.New(``) */
					case `showMenu`:
						if record[`i18n`] == nil {
							record[`i18n`] = gvar.New(map[string]interface{}{`title`: map[string]interface{}{`zh-cn`: record[`menuName`]}})
						} else {
							i18n := map[string]interface{}{}
							json.Unmarshal([]byte(record[`i18n`].String()), &i18n)
							record[`i18n`] = gvar.New(i18n)
						}
					}
				}
				result[index] = record
			}
			return
		},
	}
}

// 解析filter
func (daoThis *menuDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case `id`, `idArr`:
				m = m.Where(daoThis.Table()+`.`+daoThis.PrimaryKey(), v)
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
				m = m.WhereGTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `endTime`:
				m = m.WhereLTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `name`:
				nameField := gstr.SubStr(gstr.CaseCamel(daoThis.PrimaryKey()), 0, -2) + `Name`
				if daoThis.ColumnArrG().Contains(gstr.CaseCamelLower(nameField)) {
					m = m.WhereLike(daoThis.Table()+`.`+gstr.CaseCamelLower(nameField), `%`+gconv.String(v)+`%`)
				} else if daoThis.ColumnArrG().Contains(gstr.CaseSnakeFirstUpper(nameField)) {
					m = m.WhereLike(daoThis.Table()+`.`+gstr.CaseSnakeFirstUpper(nameField), `%`+gconv.String(v)+`%`)
				}
			case `selfMenu`: //获取当前登录身份可用的菜单。参数：map[string]interface{}{`sceneCode`: `场景标识`, `sceneId`: 场景id, `loginId`: 登录身份id}
				val := v.(map[string]interface{})
				m = m.Where(daoThis.Table()+`.`+daoThis.Columns().SceneId, val[`sceneId`])
				m = m.Where(daoThis.Table()+`.`+daoThis.Columns().IsStop, 0)
				switch val[`sceneCode`].(string) {
				case `platform`:
					if gconv.Int(val[`loginId`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Int() { //平台超级管理员，不再需要其他条件
						continue
					}
					m = m.Where(Role.Table()+`.`+Role.Columns().IsStop, 0)
					m = m.Where(RoleRelOfPlatformAdmin.Table()+`.`+RoleRelOfPlatformAdmin.Columns().AdminId, val[`loginId`])

					m = daoThis.ParseJoin(RoleRelToMenu.Table(), joinTableArr)(m)
					m = daoThis.ParseJoin(Role.Table(), joinTableArr)(m)
					m = daoThis.ParseJoin(RoleRelOfPlatformAdmin.Table(), joinTableArr)(m)
				}
				m = daoThis.ParseGroup([]string{`id`}, joinTableArr)(m)
			default:
				kArr := strings.Split(k, ` `) //支持`id > ?`等k
				if !daoThis.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(k, v)
					continue
				}
				if len(kArr) == 1 {
					if gstr.SubStr(gstr.CaseCamel(kArr[0]), -4) == `Name` {
						m = m.WhereLike(daoThis.Table()+`.`+k, `%`+gconv.String(v)+`%`)
						continue
					}
				}
				m = m.Where(daoThis.Table()+`.`+k, v)
			}
		}
		return m
	}
}

// 解析group
func (daoThis *menuDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
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
func (daoThis *menuDao) ParseOrder(order []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			kArr := strings.Split(v, ` `)
			if len(kArr) == 1 {
				kArr = append(kArr, `ASC`)
			}
			switch kArr[0] {
			case `id`:
				m = m.Order(daoThis.Table()+`.`+daoThis.PrimaryKey(), kArr[1])
			case `menuTree`:
				m = m.Order(daoThis.Table()+`.`+daoThis.Columns().Pid, `ASC`)
				m = m.Order(daoThis.Table()+`.`+daoThis.Columns().Sort, `ASC`)
				m = m.Order(daoThis.Table()+`.`+daoThis.PrimaryKey(), `ASC`)
			default:
				if daoThis.ColumnArrG().Contains(kArr[0]) {
					m = m.Order(daoThis.Table()+`.`+kArr[0], kArr[1])
				} else {
					m = m.Order(kArr[0], kArr[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoThis *menuDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case Xxxx.Table():
		relTable := Xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
			*joinTableArr = append(*joinTableArr, relTable)
			m = m.LeftJoin(relTable, relTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
		} */
		case Scene.Table():
			relTable := Scene.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(relTable, relTable+`.`+Scene.PrimaryKey()+` = `+daoThis.Table()+`.`+Scene.PrimaryKey())
			}
		case `p_` + daoThis.Table():
			relTable := `p_` + daoThis.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(daoThis.Table()+` AS `+relTable, relTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.Columns().Pid)
			}
		case RoleRelToMenu.Table():
			relTable := RoleRelToMenu.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(relTable, relTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
			}
		case Role.Table():
			relTable := Role.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				roleRelToMenuTable := RoleRelToMenu.Table()
				m = m.LeftJoin(relTable, relTable+`.`+Role.PrimaryKey()+` = `+roleRelToMenuTable+`.`+Role.PrimaryKey())
			}
		case RoleRelOfPlatformAdmin.Table():
			roleRelOfPlatformAdminTable := RoleRelOfPlatformAdmin.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleRelOfPlatformAdminTable) {
				*joinTableArr = append(*joinTableArr, roleRelOfPlatformAdminTable)
				roleRelToMenuTable := RoleRelToMenu.Table()
				m = m.LeftJoin(roleRelOfPlatformAdminTable, roleRelOfPlatformAdminTable+`.`+Role.PrimaryKey()+` = `+roleRelToMenuTable+`.`+Role.PrimaryKey())
			}
		}
		return m
	}
}

// Fill with you ideas below.
