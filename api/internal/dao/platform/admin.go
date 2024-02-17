// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	daoIndex "api/internal/dao"
	daoAuth "api/internal/dao/auth"
	"api/internal/dao/platform/internal"
	"context"
	"database/sql"

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

// 获取daoHandler
func (daoThis *adminDao) HandlerCtx(ctx context.Context, dbOpt ...map[string]interface{}) *daoIndex.DaoHandler {
	return daoIndex.NewDaoHandler(ctx, daoThis, dbOpt...)
}

// 解析分库
func (daoThis *adminDao) ParseDbGroup(ctx context.Context, dbGroupOpt ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupOpt) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *adminDao) ParseDbTable(ctx context.Context, dbTableOpt ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableOpt) > 0 {
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *adminDao) ParseDbCtx(ctx context.Context, dbOpt ...map[string]interface{}) *gdb.Model {
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
func (daoThis *adminDao) ParseInsert(insert map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
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
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}
				salt := grand.S(8)
				insertData[daoThis.Columns().Salt] = salt
				password = gmd5.MustEncrypt(password + salt)
				insertData[k] = password
			case `roleIdArr`:
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
func (daoThis *adminDao) HookInsert(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			id, _ := result.LastInsertId()

			for k, v := range daoHandler.AfterInsert {
				switch k {
				case `roleIdArr`:
					daoThis.SaveRelRole(ctx, gconv.SliceUint(v), uint(id))
				}
			}
			return
		},
	}
}

// 解析update
func (daoThis *adminDao) ParseUpdate(update map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[daoHandler.DbTable+`.`+daoThis.PrimaryKey()] = v
			case daoThis.Columns().Phone:
				updateData[daoHandler.DbTable+`.`+k] = v
				if gconv.String(v) == `` {
					updateData[daoHandler.DbTable+`.`+k] = nil
				}
			case daoThis.Columns().Account:
				updateData[daoHandler.DbTable+`.`+k] = v
				if gconv.String(v) == `` {
					updateData[daoHandler.DbTable+`.`+k] = nil
				}
			case daoThis.Columns().Password:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}
				salt := grand.S(8)
				updateData[daoHandler.DbTable+`.`+daoThis.Columns().Salt] = salt
				password = gmd5.MustEncrypt(password + salt)
				updateData[daoHandler.DbTable+`.`+k] = password
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
func (daoThis *adminDao) HookUpdate(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}

			for k, v := range daoHandler.AfterUpdate {
				switch k {
				case `roleIdArr`:
					relIdArr := gconv.SliceUint(v)
					for _, id := range daoHandler.IdArr {
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
func (daoThis *adminDao) HookDelete(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
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

			daoAuth.RoleRelOfPlatformAdmin.ParseDbCtx(ctx).Where(daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId, daoHandler.IdArr).Delete()
			return
		},
	}
}

// 解析field
func (daoThis *adminDao) ParseField(field []string, fieldWithParam map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range field {
			switch v {
			/* case `xxxx`:
			m = m.Handler(daoThis.ParseJoin(Xxxx.ParseDbTable(m.GetCtx()), daoHandler))
			daoHandler.AfterField = append(daoHandler.AfterField, v) */
			case `id`:
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.PrimaryKey() + ` AS ` + v)
			case `label`:
				m = m.Fields(`IFNULL(` + daoHandler.DbTable + `.` + daoThis.Columns().Account + `, ` + daoHandler.DbTable + `.` + daoThis.Columns().Phone + `) AS ` + v)
			case `roleIdArr`:
				m = m.Fields(daoHandler.DbTable + `.` + daoThis.PrimaryKey())
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
func (daoThis *adminDao) HookSelect(daoHandler *daoIndex.DaoHandler) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range daoHandler.AfterField {
					switch v {
					case `roleIdArr`:
						idArr, _ := daoAuth.RoleRelOfPlatformAdmin.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), record[daoThis.PrimaryKey()]).Array(daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId)
						record[v] = gvar.New(idArr)
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
func (daoThis *adminDao) ParseFilter(filter map[string]interface{}, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
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
				m = m.Where(m.Builder().WhereLike(daoHandler.DbTable+`.`+daoThis.Columns().Account, `%`+gconv.String(v)+`%`).WhereOrLike(daoHandler.DbTable+`.`+daoThis.Columns().Phone, `%`+gconv.String(v)+`%`))
			case `timeRangeStart`:
				m = m.WhereGTE(daoHandler.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `timeRangeEnd`:
				m = m.WhereLTE(daoHandler.DbTable+`.`+daoThis.Columns().CreatedAt, v)
			case `loginName`:
				if g.Validator().Rules(`required|phone`).Data(v).Run(m.GetCtx()) == nil {
					m = m.Where(daoHandler.DbTable+`.`+daoThis.Columns().Phone, v)
				} else {
					m = m.Where(daoHandler.DbTable+`.`+daoThis.Columns().Account, v)
				}
			case daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId:
				tableRoleRelOfPlatformAdmin := daoAuth.RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx())
				m = m.Where(tableRoleRelOfPlatformAdmin+`.`+k, v)
				m = m.Handler(daoThis.ParseJoin(tableRoleRelOfPlatformAdmin, daoHandler))
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
func (daoThis *adminDao) ParseGroup(group []string, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
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
func (daoThis *adminDao) ParseOrder(order []string, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
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
func (daoThis *adminDao) ParseJoin(joinTable string, daoHandler *daoIndex.DaoHandler) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		if garray.NewStrArrayFrom(daoHandler.JoinTableArr).Contains(joinTable) {
			return m
		}
		daoHandler.JoinTableArr = append(daoHandler.JoinTableArr, joinTable)
		switch joinTable {
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` AS `+joinTable, joinTable+`.`+Xxxx.Columns().XxxxId+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey()) */
		/* case daoAuth.RoleRelOfPlatformAdmin.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+`.`+daoAuth.RoleRelOfPlatformAdmin.Columns().AdminId+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey()) */
		default:
			m = m.LeftJoin(joinTable, joinTable+`.`+daoThis.PrimaryKey()+` = `+daoHandler.DbTable+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.

// 保存关联角色
func (daoThis *adminDao) SaveRelRole(ctx context.Context, relIdArr []uint, id uint) {
	relDao := daoAuth.RoleRelOfPlatformAdmin
	priKey := relDao.Columns().AdminId
	relKey := relDao.Columns().RoleId
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
