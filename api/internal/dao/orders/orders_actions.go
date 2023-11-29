// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/orders/internal"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// internalOrdersActionsDao is internal type for wrapping internal DAO implements.
type internalOrdersActionsDao = *internal.OrdersActionsDao

// ordersActionsDao is the data access object for table app_card_orders_actions.
// You can define custom methods on it to extend its functionality as you wish.
type ordersActionsDao struct {
	internalOrdersActionsDao
}

var (
	// OrdersActions is globally public accessible object for table app_card_orders_actions operations.
	OrdersActions = ordersActionsDao{
		internal.NewOrdersActionsDao(),
	}
)

// 解析分库
func (daoThis *ordersActionsDao) ParseDbGroup(ctx context.Context, dbGroupSelData ...map[string]interface{}) string {
	group := daoThis.Group()
	// 分库逻辑
	/* if len(dbGroupSelData) > 0 {
	} */
	return group
}

// 解析分表
func (daoThis *ordersActionsDao) ParseDbTable(ctx context.Context, dbTableSelData ...map[string]interface{}) string {
	table := daoThis.Table()
	// 分表逻辑
	/* if len(dbTableSelData) > 0 {
	} */
	return table
}

// 解析分库分表（对外暴露使用）
func (daoThis *ordersActionsDao) ParseDbCtx(ctx context.Context, dbSelDataList ...map[string]interface{}) *gdb.Model {
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
func (daoThis *ordersActionsDao) ParseInsert(insert map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		insertData := map[string]interface{}{}
		hookData := map[string]interface{}{}
		for k, v := range insert {
			switch k {
			case `id`:
				insertData[daoThis.PrimaryKey()] = v
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
func (daoThis *ordersActionsDao) HookInsert(data map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			// id, _ := result.LastInsertId()
			return
		},
	}
}

// 解析update
func (daoThis *ordersActionsDao) ParseUpdate(update map[string]interface{}) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		updateData := map[string]interface{}{}
		for k, v := range update {
			switch k {
			case `id`:
				updateData[tableThis+`.`+daoThis.PrimaryKey()] = v
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
func (daoThis *ordersActionsDao) HookUpdate(data map[string]interface{}, idArr ...uint) gdb.HookHandler {
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

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */
			return
		},
	}
}

// hook delete
func (daoThis *ordersActionsDao) HookDelete(idArr ...uint) gdb.HookHandler {
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
func (daoThis *ordersActionsDao) ParseField(field []string, fieldWithParam map[string]interface{}, afterField *[]string, afterFieldWithParam map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
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
func (daoThis *ordersActionsDao) HookSelect(afterField *[]string, afterFieldWithParam map[string]interface{}) gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			for _, record := range result {
				for _, v := range *afterField {
					switch v {
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
func (daoThis *ordersActionsDao) ParseFilter(filter map[string]interface{}, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for k, v := range filter {
			switch k {
			case `excId`, `excIdArr`:
				val := gconv.SliceUint(v)
				switch len(val) {
				case 0: //gconv.SliceUint会把0转换成[]uint{}，故不能用转换后的val。必须用原始数据v
					m = m.WhereNot(tableThis+`.`+daoThis.PrimaryKey(), v)
				case 1:
					m = m.WhereNot(tableThis+`.`+daoThis.PrimaryKey(), val[0])
				default:
					m = m.WhereNotIn(tableThis+`.`+daoThis.PrimaryKey(), v)
				}
			case `id`, `idArr`:
				m = m.Where(tableThis+`.`+daoThis.PrimaryKey(), v)
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
func (daoThis *ordersActionsDao) ParseGroup(group []string, joinTableArr *[]string) gdb.ModelHandler {
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
func (daoThis *ordersActionsDao) ParseOrder(order []string, joinTableArr *[]string) gdb.ModelHandler {
	return func(m *gdb.Model) *gdb.Model {
		ctx := m.GetCtx()
		tableThis := daoThis.ParseDbTable(ctx)
		for _, v := range order {
			v = gstr.Trim(v)
			k := gstr.Split(v, ` `)[0]
			switch k {
			case `id`:
				m = m.Order(tableThis + `.` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))
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
func (daoThis *ordersActionsDao) ParseJoin(joinCode string, joinTableArr *[]string) gdb.ModelHandler {
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
		default:
			m = m.LeftJoin(joinCode, joinCode+`.`+daoThis.PrimaryKey()+` = `+tableThis+`.`+daoThis.PrimaryKey())
		}
		return m
	}
}

// Fill with you ideas below.
