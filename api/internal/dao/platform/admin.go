// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/dao/platform/internal"
	"context"
	"database/sql"
	"sync"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// internalAdminDao is internal type for wrapping internal DAO implements.
type internalAdminDao = *internal.AdminDao

// adminDao is the data access object for table platform_admin.
// You can define custom methods on it to extend its functionality as you wish.
type adminDao struct {
	internalAdminDao
}

var (
	// Admin is globally public accessible object for table platform_admin operations.
	Admin = adminDao{
		internal.NewAdminDao(),
	}
)

// 解析分库
func (daoThis *adminDao) ParseDbGroup(ctx context.Context, dbGroupSeldata ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupSeldata) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *adminDao) ParseDbTable(ctx context.Context, dbTableSelData ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableSelData) > 0 {
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *adminDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
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
func (daoThis *adminDao) ParseInsert(insert map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		hookData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `id`:
				insertData[daoThis.PrimaryKey()] = v
			case daoThis.Columns().Phone:
				insertData[k] = v
				if gconv.String(v) == `` {
					insertData[k] = nil
				}
			case daoThis.Columns().Account:
				insertData[k] = v
				if gconv.String(v) == `` {
					insertData[k] = nil
				}
			case daoThis.Columns().Password:
				salt := grand.S(8)
				insertData[daoThis.Columns().Salt] = salt
				insertData[daoThis.Columns().Password] = gmd5.MustEncrypt(gconv.String(v) + salt)
			case `roleIdArr`:
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
func (daoThis *adminDao) HookInsert(data map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range data {
				switch k {
				case `roleIdArr`:
					daoThis.SaveRelRole(ctx, gconv.SliceInt(v), int(id))
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *adminDao) ParseUpdate(update map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoThis.Table()+`.`+daoThis.PrimaryKey()] = v
			case daoThis.Columns().Phone:
				updateData[daoThis.Table()+`.`+k] = v
				if gconv.String(v) == `` {
					updateData[daoThis.Table()+`.`+k] = nil
				}
			case daoThis.Columns().Account:
				updateData[daoThis.Table()+`.`+k] = v
				if gconv.String(v) == `` {
					updateData[daoThis.Table()+`.`+k] = nil
				}
			case daoThis.Columns().Password:
				salt := grand.S(8)
				updateData[daoThis.Table()+`.`+daoThis.Columns().Salt] = salt
				updateData[daoThis.Table()+`.`+daoThis.Columns().Password] = gmd5.MustEncrypt(gconv.String(v) + salt)
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
func (daoThis *adminDao) HookUpdate(data map[string]interface{}, idArr ...int) gdb.HookHandler {
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
				case `roleIdArr`:
					relIdArr := gconv.SliceInt(v)
					for _, id := range idArr {
						daoThis.SaveRelRole(ctx, relIdArr, id)
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
func (daoThis *adminDao) HookDelete(idArr ...int) gdb.HookHandler {
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

			daoAuth.RoleRelOfPlatformAdmin.ParseDbCtx(ctx).Where(daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId, idArr).Delete()
			return
		},
	}
}

// 解析field
func (daoThis *adminDao) ParseField(field []string, fieldWithParam map[string]interface{}, afterField *[]string, afterFieldWithParam map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = daoThis.ParseJoin(Xxxx.Table(), joinTableArr)(m)
			*afterField = append(*afterField, v) */
			case `id`:
				m = m.Fields(daoThis.Table() + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(`IFNULL(` + daoThis.Table() + `.` + daoThis.Columns().Account + `, ` + daoThis.Table() + `.` + daoThis.Columns().Phone + `) AS ` + v)
			case `roleIdArr`:
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
func (daoThis *adminDao) HookSelect(afterField *[]string, afterFieldWithParam map[string]interface{}) gdb.HookHandler {
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
						case `roleIdArr`:
							idArr, _ := daoAuth.RoleRelOfPlatformAdmin.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId)
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
func (daoThis *adminDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
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
				m = m.Where(m.Builder().WhereLike(daoThis.Table()+`.`+daoThis.Columns().Account, `%`+gconv.String(v)+`%`).WhereOrLike(daoThis.Table()+`.`+daoThis.Columns().Phone, `%`+gconv.String(v)+`%`))
			case `timeRangeStart`:
				m = m.WhereGTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(daoThis.Table()+`.`+daoThis.Columns().CreatedAt, v)
			case `accountOrPhone`:
				if g.Validator().Rules(`required|integer`).Data(v).Run(m.GetCtx()) == nil {
					m = m.Where(daoThis.Table()+`.`+daoThis.Columns().Phone, v)
				} else {
					m = m.Where(daoThis.Table()+`.`+daoThis.Columns().Account, v)
				}
			case `roleId`:
				m = m.Where(daoAuth.RoleRelOfPlatformAdmin.Table()+`.`+k, v)
				m = daoThis.ParseJoin(daoAuth.RoleRelOfPlatformAdmin.Table(), joinTableArr)(m)
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
func (daoThis *adminDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
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
func (daoThis *adminDao) ParseOrder(order []string, joinTableArr *[]string) gdb.ModelHandler {
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
func (daoThis *adminDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case Xxxx.Table():
		relTable := Xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
			*joinTableArr = append(*joinTableArr, relTable)
			m = m.LeftJoin(relTable, relTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
		} */
		case daoAuth.RoleRelOfPlatformAdmin.Table():
			relTable := daoAuth.RoleRelOfPlatformAdmin.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(relTable, relTable+`.`+daoThis.PrimaryKey()+` = `+daoThis.Table()+`.`+daoThis.PrimaryKey())
			}
		}
		return m
	}
}

// Fill with you ideas below.

// 保存关联角色
func (daoThis *adminDao) SaveRelRole(ctx context.Context, relIdArr []int, id int) {
	relDao := daoAuth.RoleRelOfPlatformAdmin
	priKey := relDao.Columns().AdminId
	relKey := relDao.Columns().RoleId
	relIdArrOfOldTmp, _ := relDao.ParseDbCtx(ctx).Where(priKey, id).Array(relKey)
	relIdArrOfOld := gconv.SliceInt(relIdArrOfOldTmp)

	/**----新增关联角色 开始----**/
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
	/**----新增关联角色 结束----**/

	/**----删除关联角色 开始----**/
	deleteRelIdArr := gset.NewIntSetFrom(relIdArrOfOld).Diff(gset.NewIntSetFrom(relIdArr)).Slice()
	if len(deleteRelIdArr) > 0 {
		relDao.ParseDbCtx(ctx).Where(priKey, id).Where(relKey, deleteRelIdArr).Delete()
	}
	/**----删除关联角色 结束----**/
}
