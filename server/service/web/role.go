package web

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nebula.xyz/global"
	"nebula.xyz/model/request"
	"nebula.xyz/model/response"
	"nebula.xyz/model/system"
	"nebula.xyz/service"
	"strconv"
	"strings"
)

type RoleService struct{}

func (s *RoleService) CreateRole(role request.CreateRole) error {
	sysRole := system.SysRole{}
	sysRole.Slug = role.Slug
	// 判断标识是否唯一
	err := sysRole.GetRoleBySlug()
	if err != nil {
		return err
	}
	if sysRole.ID != 0 {
		return errors.New(fmt.Sprintf("%s标识已经存在", role.Slug))
	}
	sysRole.Name = role.Name
	sysRole.ParentId = role.ParentId
	sysRole.Desc = role.Desc
	err = sysRole.SaveRole()
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) GetRolePagination(pagination request.Pagination) (roles []system.SysRole, total int64, err error) {
	db := global.DB.Model(&system.SysRole{})

	err = db.Where("parent_id = ?", 0).Count(&total).Error
	if err != nil {
		global.Logger.Error("角色分页查询失败", zap.Error(err))
		return
	}
	if total < 0 {
		return
	}

	offset := (pagination.Page - 1) * pagination.Limit
	err = db.Where("parent_id = ?", 0).Offset(offset).Limit(pagination.Limit).Find(&roles).Error
	if err != nil {
		global.Logger.Error("角色分页查询失败", zap.Error(err))
		return
	}
	for k := range roles {
		err = s.findChildrenRole(&roles[k])
	}
	return
}

func (s *RoleService) findChildrenRole(role *system.SysRole) error {
	err := global.DB.Where("parent_id = ?", role.ID).Find(&role.Children).Error
	if len(role.Children) > 0 {
		for k := range role.Children {
			err = s.findChildrenRole(&role.Children[k])
		}
	}
	return err
}

func (s *RoleService) UpdateRole(r request.UpdateRole) error {
	role := system.SysRole{}
	role.ID = r.ID
	role.Name = r.Name
	role.Desc = r.Desc
	global.Logger.Error("abcccccc")
	err := global.DB.Where("id = ?", role.ID).Find(&system.SysRole{}).Updates(&role).Error
	return err
}

func (s *RoleService) DeleteRole(id uint) error {

	role := system.SysRole{}
	role.ID = id
	result := global.DB.Where("id = ?", id).Find(&role)
	if result.RowsAffected == 0 {
		return errors.New("不存在当前角色")
	}
	var total int64 = 0
	err := global.DB.Model(&system.SysRole{}).Where("parent_id = ?", role.ID).Count(&total).Error
	if err != nil {
		return errors.New("查询子角色错误")
	}
	if total > 0 {
		return errors.New("存在子角色，请删除子角色")
	}
	err = global.DB.Model(&system.SysRole{}).Where("id = ?", id).Delete(role).Error
	if err != nil {
		global.Logger.Error("删除角色失败", zap.Error(err))
		return errors.New("删除角色失败")
	}
	return nil
}

func (s *RoleService) GetAllPermission(roleId uint64) (*response.PermissionResponse, error) {
	db := global.DB
	var role system.SysRole
	err := db.Model(&system.SysRole{}).
		Preload("SysMenus").
		Preload("SysButton").
		Where("id = ?", roleId).
		First(&role).Error
	if err != nil {
		return nil, errors.New("获取权限错误")
	}
	// 系统中的所有权限
	var sysMenus []system.SysMenu
	var sysButton []system.SysButton
	_ = db.Model(&system.SysMenu{}).Find(&sysMenus).Error
	_ = db.Model(&system.SysButton{}).Find(&sysButton).Error
	menuSet := make(map[uint]bool)
	buttonSet := make(map[uint]bool)
	for _, item := range role.SysMenus {
		menuSet[item.ID] = true
	}

	for _, item := range role.SysButton {
		buttonSet[item.ID] = true
	}

	// 当前用户的上级权限
	menuParentSet := make(map[uint]bool)
	buttonParentSet := make(map[uint]bool)

	if *role.ParentId != 0 {
		var sysRoleMenu []system.SysRoleMenus
		var sysRoleButton []system.SysRoleButton
		_ = db.Model(&system.SysRoleMenus{}).Where("sys_role_id = ?", role.ParentId).Find(&sysRoleMenu).Error
		_ = db.Model(&system.SysRoleButton{}).Where("sys_role_id = ?", role.ParentId).Find(&sysRoleButton).Error
		for _, item := range sysRoleMenu {
			menuParentSet[item.SysMenuId] = true
		}
		for _, item := range sysRoleButton {
			buttonParentSet[item.SysButtonId] = true
		}
	} else {
		for _, item := range sysMenus {
			menuParentSet[item.ID] = true
		}
		for _, item := range sysButton {
			buttonParentSet[item.ID] = true
		}
	}

	menus := make([]response.PermissionDetails, 0)
	for _, menu := range sysMenus {
		menus = append(menus, response.PermissionDetails{
			Label:   menu.Name,
			Value:   menu.ID,
			Disable: menuParentSet[menu.ID],
			Hold:    menuSet[menu.ID],
		})
	}

	buttons := make([]response.PermissionDetails, 0)
	for _, button := range sysButton {
		buttons = append(buttons, response.PermissionDetails{
			Label:   button.Name,
			Value:   button.ID,
			Disable: buttonParentSet[button.ID],
			Hold:    buttonSet[button.ID],
		})
	}
	return &response.PermissionResponse{
		Menus:   menus,
		Buttons: buttons,
	}, nil
}

func (s *RoleService) UpdatePermission(permission request.UpdatePermission) error {
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		var role system.SysRole
		err := tx.Model(&system.SysRole{}).Where("id = ?", permission.RoleId).First(&role).Error
		if err != nil {
			return errors.New("角色不存在")
		}
		// 获取上级权限
		menuParentMap := make(map[uint]system.SysMenu)
		buttonParentMap := make(map[uint]system.SysButton)
		var sysMenus []system.SysMenu
		var sysButton []system.SysButton
		if *role.ParentId != 0 {
			var sysRoleMenu []system.SysRoleMenus
			var sysRoleButton []system.SysRoleButton
			_ = tx.Model(&system.SysRoleMenus{}).Where("sys_role_id = ?", role.ParentId).Find(&sysRoleMenu).Error
			_ = tx.Model(&system.SysRoleButton{}).Where("sys_role_id = ?", role.ParentId).Find(&sysRoleButton).Error
			menusId := make([]uint, 0)
			for _, item := range sysRoleMenu {
				menusId = append(menusId, item.SysMenuId)
			}
			buttonId := make([]uint, 0)
			for _, item := range sysRoleButton {
				buttonId = append(buttonId, item.SysButtonId)
			}
			_ = tx.Model(&system.SysMenu{}).Where("id in ?", menusId).Find(&sysMenus).Error
			_ = tx.Model(&system.SysButton{}).Where("id in ?", buttonId).Find(&sysButton).Error

		} else {
			_ = tx.Model(&system.SysMenu{}).Find(&sysMenus).Error
			_ = tx.Model(&system.SysButton{}).Find(&sysButton).Error
		}
		for _, item := range sysMenus {
			menuParentMap[item.ID] = item
		}
		for _, item := range sysButton {
			buttonParentMap[item.ID] = item
		}
		// 判断权限是否在范围内
		menus := make([]system.SysRoleMenus, 0)
		for _, menu := range permission.Menu {
			if _, exists := menuParentMap[menu]; !exists {
				return errors.New("权限范围错误")
			}
			menus = append(menus, system.SysRoleMenus{
				SysMenuId: menu,
				SysRoleId: role.ID,
			})
		}
		buttons := make([]system.SysRoleButton, 0)
		for _, button := range permission.Button {
			if _, exists := buttonParentMap[button]; !exists {
				return errors.New("权限范围错误")
			}
			buttons = append(buttons, system.SysRoleButton{
				SysRoleId:   role.ID,
				SysButtonId: button,
			})
		}
		var sysRoleMenu []system.SysRoleMenus
		var sysRoleButton []system.SysRoleButton
		err = tx.Model(&system.SysRoleMenus{}).Where("sys_role_id = ?", role.ID).Delete(&sysRoleMenu).Error
		if err != nil {
			global.Logger.Error("删除角色菜单表失败", zap.Error(err))
			defer tx.Rollback()
			return errors.New("更新权限错误")
		}
		err = tx.Model(&system.SysRoleButton{}).Where("sys_role_id = ?", role.ID).Delete(&sysRoleButton).Error
		if err != nil {
			global.Logger.Error("删除角色按钮表失败", zap.Error(err))
			defer tx.Rollback()
			return errors.New("更新权限错误")
		}
		if len(menus) > 0 {
			err = tx.Model(&system.SysRoleMenus{}).Save(&menus).Error
			if err != nil {
				global.Logger.Error("保存角色菜单权限错误", zap.Error(err))
				defer tx.Rollback()
				return errors.New("更新权限错误")
			}
		}
		if len(buttons) > 0 {
			err = tx.Model(&system.SysRoleButton{}).Save(&buttons).Error
			if err != nil {
				global.Logger.Error("保存角色按钮权限错误", zap.Error(err))
				defer tx.Rollback()
				return errors.New("更新权限错误")
			}
		}

		type CasbinRuleTmp struct {
			url    string
			method string
		}
		rules := make([]CasbinRuleTmp, 0)
		for _, menu := range permission.Menu {
			sysMenu := menuParentMap[menu]
			methods := strings.Split(sysMenu.Method, "|")
			for _, method := range methods {
				rules = append(rules, CasbinRuleTmp{
					url:    sysMenu.Path,
					method: method,
				})
			}
		}

		for _, button := range permission.Button {
			sysButton := buttonParentMap[button]
			methods := strings.Split(sysButton.Method, "|")
			for _, method := range methods {
				rules = append(rules, CasbinRuleTmp{
					url:    sysButton.Path,
					method: method,
				})
			}
		}
		// 更新Casbin_Rule政策
		casbinService := service.SysServiceAll.CasbinService
		casbin := casbinService.Casbin()
		_, err = casbin.RemoveFilteredPolicy(0, strconv.FormatUint(uint64(role.ID), 10))
		//err = casbin.SavePolicy()
		if err != nil {
			global.Logger.Error("删除政策失败", zap.Error(err))
			defer tx.Rollback()
			return errors.New("更新权限错误")
		}
		global.Logger.Debug("删除政策成功")

		for _, rule := range rules {
			_, err := casbin.AddPolicy(strconv.FormatUint(uint64(role.ID), 10), rule.url, rule.method)
			if err != nil {
				global.Logger.Error("保存政策失败", zap.Error(err))
				tx.Rollback()
				return errors.New("更新权限错误")
			}
		}
		err = casbin.SavePolicy()
		if err != nil {
			global.Logger.Error("保存政策失败", zap.Error(err))
			defer tx.Rollback()
			return errors.New("更新权限错误")
		}
		return nil
	})
	return err
}
