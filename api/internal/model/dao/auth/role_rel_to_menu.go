// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/model/dao/auth/internal"
	"context"
	"strings"

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
func (daoRoleRelToMenu *roleRelToMenuDao) ParseInsert(insert []map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := make([]map[string]interface{}, len(insert))
		for index, item := range insert {
			insertData[index] = map[string]interface{}{}
			for k, v := range item {
				switch k {
				case "id":
					insertData[index][daoRoleRelToMenu.PrimaryKey()] = v
				default:
					//数据库不存在的字段过滤掉，未传值默认true
					if (len(fill) == 0 || fill[0]) && !daoRoleRelToMenu.ColumnArrG().Contains(k) {
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
func (daoRoleRelToMenu *roleRelToMenuDao) ParseUpdate(update map[string]interface{}, fill ...bool) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case "id":
				updateData[daoRoleRelToMenu.Table()+"."+daoRoleRelToMenu.PrimaryKey()] = v
			default:
				//数据库不存在的字段过滤掉，未传值默认true
				if (len(fill) == 0 || fill[0]) && !daoRoleRelToMenu.ColumnArrG().Contains(k) {
					continue
				}
				updateData[daoRoleRelToMenu.Table()+"."+k] = v
			}
		}
		//m = m.Data(updateData) //字段被解析成`table.xxxx`，正确的应该是`table`.`xxxx`
		//解决字段被解析成`table.xxxx`的BUG
		fieldArr := []string{}
		valueArr := []interface{}{}
		for k, v := range updateData {
			fieldArr = append(fieldArr, k+" = ?")
			valueArr = append(valueArr, v)
		}
		data := []interface{}{strings.Join(fieldArr, ",")}
		data = append(data, valueArr...)
		m = m.Data(data...)
		return m
	}
}

// 解析field
func (daoRoleRelToMenu *roleRelToMenuDao) ParseField(field []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		afterField := []string{}
		for _, v := range field {
			switch v {
			/* case "xxxx":
			m = daoRoleRelToMenu.ParseJoin("xxxx", joinTableArr)(m)
			afterField = append(afterField, v) */
			case "id":
				m = m.Fields(daoRoleRelToMenu.Table() + "." + daoRoleRelToMenu.PrimaryKey() + " AS " + v)
			default:
				if daoRoleRelToMenu.ColumnArrG().Contains(v) {
					m = m.Fields(daoRoleRelToMenu.Table() + "." + v)
				} else {
					m = m.Fields(v)
				}
			}
		}
		if len(afterField) > 0 {
			m = m.Hook(daoRoleRelToMenu.AfterField(afterField))
		}
		return m
	}
}

// 解析filter
func (daoRoleRelToMenu *roleRelToMenuDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for k, v := range filter {
			switch k {
			case "id", "idArr":
				m = m.Where(daoRoleRelToMenu.Table()+"."+daoRoleRelToMenu.PrimaryKey(), v)
			case "excId":
				m = m.WhereNot(daoRoleRelToMenu.Table()+"."+daoRoleRelToMenu.PrimaryKey(), v)
			case "excIdArr":
				m = m.WhereNotIn(daoRoleRelToMenu.Table()+"."+daoRoleRelToMenu.PrimaryKey(), v)
			case "startTime":
				m = m.WhereGTE(daoRoleRelToMenu.Table()+".createTime", v)
			case "endTime":
				m = m.WhereLTE(daoRoleRelToMenu.Table()+".createTime", v)
			case "keyword":
				keywordField := strings.ReplaceAll(daoRoleRelToMenu.PrimaryKey(), "Id", "Name")
				switch v := v.(type) {
				case *string:
					m = m.WhereLike(daoRoleRelToMenu.Table()+"."+keywordField, *v)
				case string:
					m = m.WhereLike(daoRoleRelToMenu.Table()+"."+keywordField, v)
				default:
					m = m.Where(daoRoleRelToMenu.Table()+"."+keywordField, v)
				}
			default:
				kArr := strings.Split(k, " ")
				if daoRoleRelToMenu.ColumnArrG().Contains(kArr[0]) {
					m = m.Where(daoRoleRelToMenu.Table()+"."+k, v)
				} else {
					m = m.Where(k, v)
				}
			}
		}
		return m
	}
}

// 解析group
func (daoRoleRelToMenu *roleRelToMenuDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range group {
			switch v {
			case "id":
				m = m.Group(daoRoleRelToMenu.Table() + "." + daoRoleRelToMenu.PrimaryKey())
			default:
				if daoRoleRelToMenu.ColumnArrG().Contains(v) {
					m = m.Group(daoRoleRelToMenu.Table() + "." + v)
				} else {
					m = m.Group(v)
				}
			}
		}
		return m
	}
}

// 解析order
func (daoRoleRelToMenu *roleRelToMenuDao) ParseOrder(order [][2]string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		for _, v := range order {
			switch v[0] {
			case "id":
				m = m.Order(daoRoleRelToMenu.Table()+"."+daoRoleRelToMenu.PrimaryKey(), v[1])
			default:
				if daoRoleRelToMenu.ColumnArrG().Contains(v[0]) {
					m = m.Order(daoRoleRelToMenu.Table()+"."+v[0], v[1])
				} else {
					m = m.Order(v[0], v[1])
				}
			}
		}
		return m
	}
}

// 解析join
func (daoRoleRelToMenu *roleRelToMenuDao) ParseJoin(joinCode string, joinTableArr *[]string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		switch joinCode {
		/* case "xxxx":
		xxxxTable := xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(xxxxTable) {
			*joinTableArr = append(*joinTableArr, xxxxTable)
			m = m.LeftJoin(xxxxTable, xxxxTable+"."+daoRoleRelToMenu.PrimaryKey()+" = "+daoRoleRelToMenu.Table()+"."+daoRoleRelToMenu.PrimaryKey())
		} */
		}
		return m
	}
}

// 获取数据后，再处理的字段
func (daoRoleRelToMenu *roleRelToMenuDao) AfterField(afterField []string) gdb.HookHandler {
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
func (daoRoleRelToMenu *roleRelToMenuDao) Info(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (info gdb.Record, err error) {
	joinTableArr := []string{}
	model := daoRoleRelToMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRoleRelToMenu.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRoleRelToMenu.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRoleRelToMenu.ParseOrder(order, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 列表
func (daoRoleRelToMenu *roleRelToMenuDao) List(ctx context.Context, filter map[string]interface{}, field []string, order ...[2]string) (list gdb.Result, err error) {
	joinTableArr := []string{}
	model := daoRoleRelToMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRoleRelToMenu.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRoleRelToMenu.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRoleRelToMenu.ParseOrder(order, &joinTableArr))
	}
	list, err = model.All()
	return
}

// Fill with you ideas below.