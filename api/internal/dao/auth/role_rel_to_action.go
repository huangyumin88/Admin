// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
)

// internalRoleRelToActionDao is internal type for wrapping internal DAO implements.
type internalRoleRelToActionDao = *internal.RoleRelToActionDao

// roleRelToActionDao is the data access object for table auth_role_rel_to_action.
// You can define custom methods on it to extend its functionality as you wish.
type roleRelToActionDao struct {
	internalRoleRelToActionDao
}

var (
	// RoleRelToAction is globally public accessible object for table auth_role_rel_to_action operations.
	RoleRelToAction = roleRelToActionDao{
		internal.NewRoleRelToActionDao(),
	}
)

// Fill with you ideas below.