// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/auth/internal"
	"context"
	"encoding/json"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
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
func (daoThis *menuDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
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
func (daoThis *menuDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoThis.Table()+`.`+daoThis.PrimaryKey()] = v
			case `pidPathOfChild`: //更新所有子孙级的pidPath。参数：map[string]interface{}{`newVal`: `父级新pidPath`, `oldVal`:`父级旧pidPath`}
				val := gconv.Map(v)
				updateData[daoThis.Table()+`.pidPath`] = gdb.Raw(`REPLACE(` + daoThis.Table() + `.pidPath, '` + gconv.String(val[`oldVal`]) + `', '` + gconv.String(val[`newVal`]) + `')`)
			case `levelOfChild`: //更新所有子孙级的level。参数：map[string]interface{}{`newVal`: 父级新level, `oldVal`:父级旧level}
				val := gconv.Map(v)
				updateData[daoThis.Table()+`.level`] = gdb.Raw(daoThis.Table() + `.level + ` + gconv.String(gconv.Int(val[`newVal`])-gconv.Int(val[`oldVal`])))
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
func (daoThis *menuDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = daoThis.ParseJoin(`xxxx`, joinTableArr)(m)
			afterField = append(afterField, v) */
			case `id`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `menuTree`: //树状需要以下字段和排序方式
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
				m = m.Fields(daoThis.Table() + `.pid`)

				m = daoThis.ParseOrder([][2]string{{`menuTree`, ``}}, joinTableArr)(m) //排序方式
			case `showMenu`: //前端显示菜单需要以下字段，且title需要转换
				m = m.Fields(daoThis.Table() + `.menuName`)
				m = m.Fields(daoThis.Table() + `.menuIcon`)
				m = m.Fields(daoThis.Table() + `.menuUrl`)
				m = m.Fields(daoThis.Table() + `.extraData->'$.i18n' AS i18n`)
				//m = m.Fields(gdb.Raw(`JSON_UNQUOTE(JSON_EXTRACT(extraData, \`$.i18n\`)) AS i18n`))//mysql不能直接转成对象返回
				afterField = append(afterField, v)
			case `sceneName`:
				m = m.Fields(Scene.Table() + `.` + v)
				m = daoThis.ParseJoin(`scene`, joinTableArr)(m)
			case `pMenuName`:
				m = m.Fields(`p_` + daoThis.Table() + `.menuName AS ` + v)
				m = daoThis.ParseJoin(`pMenu`, joinTableArr)(m)
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
func (daoThis *menuDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
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
			case `selfMenu`: //获取当前登录身份可用的菜单。参数：map[string]interface{}{`sceneCode`: `场景标识`, `sceneId`: 场景id, `loginId`: 登录身份id}
				val := v.(map[string]interface{})

				m = m.Where(daoThis.Table()+`.sceneId`, val[`sceneId`])
				m = m.Where(daoThis.Table()+`.isStop`, 0)
				switch val[`sceneCode`].(string) {
				case `platformAdmin`:
					if gconv.Int(val[`loginId`]) == g.Cfg().MustGet(m.GetCtx(), `superPlatformAdminId`).Int() { //平台超级管理员，不再需要其他条件
						return m
					}
					m = m.Where(Role.Table()+`.isStop`, 0)
					m = m.Where(RoleRelOfPlatformAdmin.Table()+`.adminId`, val[`loginId`])

					m = daoThis.ParseJoin(`roleRelToMenu`, joinTableArr)(m)
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
func (daoThis *menuDao) ParseOrder(order [][2]string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case `id`:
				m = m.Order(daoThis.Table()+`.`+daoThis.PrimaryKey(), v[1])
			case `menuTree`:
				m = m.Order(daoThis.Table()+`.pid`, `ASC`)
				m = m.Order(daoThis.Table()+`.sort`, `ASC`)
				m = m.Order(daoThis.Table()+`.`+daoThis.PrimaryKey(), `ASC`)
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
func (daoThis *menuDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case `xxxx`:
		xxxxTable := xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(xxxxTable) {
			*joinTableArr = append(*joinTableArr, xxxxTable)
			m = m.LeftJoin(xxxxTable, xxxxTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
		} */
		case `scene`:
			sceneTable := Scene.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(sceneTable) {
				*joinTableArr = append(*joinTableArr, sceneTable)
				m = m.LeftJoin(sceneTable, sceneTable+`.`+Scene.PrimaryKey()+` = `+daoThis.Table()+`.`+Scene.PrimaryKey())
			}
		case `pMenu`:
			pMenuTable := `p_` + daoThis.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(pMenuTable) {
				*joinTableArr = append(*joinTableArr, pMenuTable)
				m = m.LeftJoin(daoThis.Table()+` AS `+pMenuTable, pMenuTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.pid`)
			}
		case `roleRelToMenu`:
			roleRelToMenuTable := RoleRelToMenu.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleRelToMenuTable) {
				*joinTableArr = append(*joinTableArr, roleRelToMenuTable)
				m = m.LeftJoin(roleRelToMenuTable, roleRelToMenuTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
			}
		case `role`:
			roleTable := Role.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(roleTable) {
				*joinTableArr = append(*joinTableArr, roleTable)
				roleRelToMenuTable := RoleRelToMenu.Table()
				m = m.LeftJoin(roleTable, roleTable+`.`+Role.PrimaryKey()+` = `+roleRelToMenuTable+`.`+Role.PrimaryKey())
			}
		case `roleRelOfPlatformAdmin`:
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

// 获取数据后，再处理的字段
func (daoThis *menuDao) AfterField(afterField []string) gdb.HookHandler {
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
				result[i] = record
			}
			return
		},
	}
}

// Fill with you ideas below.
