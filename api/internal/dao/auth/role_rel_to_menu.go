// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
	"context"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
)

// internalRoleRelToMenuDao is internal type for wrapping internal DAO implements.
type internalRoleRelToMenuDao = *internal.RoleRelToMenuDao

// roleRelToMenuDao is the data access object for table auth_role_rel_to_menu.
// You can define custom methods on it to extend its functionality as you wish.
type roleRelToMenuDao struct {
	internalRoleRelToMenuDao
}

var (
	// RoleRelToMenu is globally public accessible object for table auth_role_rel_to_menu operations.
	RoleRelToMenu = roleRelToMenuDao{
		internal.NewRoleRelToMenuDao(),
	}
)

// 解析insert
func (dao *roleRelToMenuDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := []map[string]interface{}{}
		for index, item := range insert {
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][dao.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !dao.ColumnArrG().Contains(k) {
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
func (dao *roleRelToMenuDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[dao.Table()+"."+dao.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !dao.ColumnArrG().Contains(k) {
					continue
				}
				updateData[dao.Table()+"."+k] = v
			}
		}
		m = m.Data(updateData)
		return m
	}
}

// 解析field
func (dao *roleRelToMenuDao) ParseField(field []string, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(dao.Table() + "." + dao.PrimaryKey())
			default:
				if dao.ColumnArrG().Contains(v) {
					m = m.Fields(dao.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(dao.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (dao *roleRelToMenuDao) ParseFilter(filter map[string]interface{}, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id":
				m = m.Where(dao.Table()+"."+dao.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(dao.Table()+"."+dao.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(dao.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(dao.Table()+".createTime", v)
			default:
				kArr := strings.Split(k, " ")
				if dao.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(dao.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (dao *roleRelToMenuDao) ParseGroup(group []string, joinCodeArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(dao.Table() + "." + dao.PrimaryKey())
			default:
				if dao.ColumnArrG().Contains(v) {
					m = m.Group(dao.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (dao *roleRelToMenuDao) ParseOrder(order [][2]string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(dao.Table()+"."+dao.PrimaryKey(), v[1])
			default:
				if dao.ColumnArrG().Contains(v[0]) {
					m = m.Order(dao.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (dao *roleRelToMenuDao) ParseJoin(joinCode string, joinCodeArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		if garray.NewStrArrayFrom(*joinCodeArr).Contains(joinCode) {
			*joinCodeArr = append(*joinCodeArr, joinCode)
			switch joinCode {
			/* case "xxxx":
			m = m.LeftJoin("xxxx", "xxxx."+dao.PrimaryKey()+" = "+dao.Table()+"."+dao.PrimaryKey()) */
			}
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (dao *roleRelToMenuDao) AfterField(afterField []string) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for i, record := range result {
				for _, v := range afterField {
					switch v {
					/* case "xxxx":
					record[v] = gvar.New("") */
					}
				}
				result[i] = record
			}
			return
		},
	}
}

// 详情
func (dao *roleRelToMenuDao) Info(ctx context.Context, field []string, filter map[string]interface{}) (info gdb.Record, err error) {
	joinCodeArr := []string{}
	info, err = dao.Ctx(ctx).Handler(dao.ParseField(field, &joinCodeArr), dao.ParseFilter(filter, &joinCodeArr)).One()
	return
}

// 列表
func (dao *roleRelToMenuDao) List(ctx context.Context, field []string, filter map[string]interface{}) (list gdb.Result, err error) {
	joinCodeArr := []string{}
	list, err = dao.Ctx(ctx).Handler(dao.ParseField(field, &joinCodeArr), dao.ParseFilter(filter, &joinCodeArr)).All()
	return
}

// Fill with you ideas below.
