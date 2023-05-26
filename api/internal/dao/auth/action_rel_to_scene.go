// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"api/internal/dao/auth/internal"
)

// internalActionRelToSceneDao is internal type for wrapping internal DAO implements.
type internalActionRelToSceneDao = *internal.ActionRelToSceneDao

// actionRelToSceneDao is the data access object for table auth_action_rel_to_scene.
// You can define custom methods on it to extend its functionality as you wish.
type actionRelToSceneDao struct {
	internalActionRelToSceneDao
}

var (
	// ActionRelToScene is globally public accessible object for table auth_action_rel_to_scene operations.
	ActionRelToScene = actionRelToSceneDao{
		internal.NewActionRelToSceneDao(),
	}
)

// Fill with you ideas below.