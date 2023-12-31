package web

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nebula.xyz/global"
	"nebula.xyz/model/request"
	"nebula.xyz/model/system"
	"nebula.xyz/utils"
)

type UserService struct{}

// 用户登录
func (u *UserService) Login(user *system.SysUser) (userHandler *system.SysUser, err error) {
	if global.DB == nil {
		global.Logger.Error("数据库未配置")
		return nil, fmt.Errorf("数据库未配置")
	}
	var userTmp system.SysUser
	err = global.DB.Where("email = ?", user.Email).First(&userTmp).Error
	if err == nil {
		// 判断密码是否相等
		if ok := utils.BcryptCheck(user.PassWord, userTmp.PassWord); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &userTmp, err
}

func (u *UserService) GetUserInfoPagination(pagination request.Pagination) (users []system.SysUser, total int64, err error) {
	db := global.DB.Model(&system.SysUser{})

	err = db.Count(&total).Error
	if err != nil {
		global.Logger.Error("用户分页查询失败", zap.Error(err))
		return
	}
	if total < 0 {
		return
	}
	offset := (pagination.Page - 1) * pagination.Limit
	err = db.Offset(offset).Limit(pagination.Limit).Preload("Roles").Preload("Role").Find(&users).Error
	if err != nil {
		global.Logger.Error("设备分页查询失败", zap.Error(err))
	}
	return
}

func (u *UserService) CreateUser(create request.UserCreate) error {
	// 1. 判断邮箱是否存在
	user := system.SysUser{}
	if !errors.Is(global.DB.Model(&system.SysUser{}).Where("email = ?", create.Email).Find(&user).Error, gorm.ErrRecordNotFound) {
		global.Logger.Error("该邮箱已被注册")
		return errors.New("该邮箱已被注册")
	}

	// 2. 对密码进行加密
	passwordHash := utils.BcryptHash(create.Password)

	user.UserName = create.Name
	user.PassWord = passwordHash
	user.Email = create.Email
	user.Enable = 1
	user.RoleID = create.Roles[0]
	var sysRoles []system.SysRole
	for _, v := range create.Roles {
		sysRoles = append(sysRoles, system.SysRole{
			ID: v,
		})
	}
	user.Roles = sysRoles
	err := global.DB.Model(&system.SysUser{}).Create(&user).Error
	if err != nil {
		global.Logger.Error("创建用户失败", zap.Error(err))
	}
	return err
}

// EnableUser 更新用户状态
func (u *UserService) EnableUser(user request.EnableUser) error {
	db := global.DB.Model(&system.SysUser{})
	var sysUser system.SysUser
	if errors.Is(db.Where("id = ?", user.ID).Find(&sysUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("该用户不能存在")
	}
	err := db.Where("id = ?", user.ID).Find(&system.SysUser{}).Update("enable", user.Enable).Error
	if err != nil {
		global.Logger.Error("更新用户状态失败", zap.Error(err))
		return errors.New("更新用户状态失败")
	}
	return nil
}

func (u *UserService) UpdateUser(create request.UserUpdate) error {

	// 1. 判断邮箱是否存在
	user := system.SysUser{}
	err := global.DB.Model(&system.SysUser{}).Where("email = ? and id != ?", create.Email, create.ID).Find(&user).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			global.Logger.Error("该邮箱已被注册")
			return errors.New("该邮箱已被注册")
		}
		return err
	}

	user.UserName = create.Name
	user.Email = create.Email
	user.RoleID = create.Roles[0]

	return global.DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUserRole{}, "sys_user_id = ?", create.ID).Error
		if TxErr != nil {
			return TxErr
		}
		var sysRoles []system.SysUserRole
		for _, v := range create.Roles {
			sysRoles = append(sysRoles, system.SysUserRole{
				SysUserId: create.ID,
				SysRoleId: v,
			})
		}
		TxErr = tx.Model(&system.SysUserRole{}).Create(sysRoles).Error
		if TxErr != nil {
			global.Logger.Error("添加用户角色失败", zap.Error(TxErr))
			return errors.New("更新用户失败")
		}
		TxErr = tx.Where("id = ? ", create.ID).Updates(&user).Error
		if TxErr != nil {
			global.Logger.Error("更新用户失败", zap.Error(TxErr))
			return errors.New("更新用户失败")
		}
		// 返回 nil 提交事务
		return nil
	})
}

func (u *UserService) DeleteUser(id string) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&system.SysUser{}).Where("id = ?", id).Find(&system.SysUser{}).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("用户不存在")
			}
			return errors.New("删除用户错误")
		}
		err = tx.Model(&system.SysUserRole{}).Where("sys_user_id = ?", id).Delete(&system.SysUserRole{}).Error
		if err != nil {
			global.Logger.Error("删除用户失败", zap.Error(err))
			return errors.New("删除用户错误")
		}
		err = tx.Model(&system.SysUser{}).Where("id = ?", id).Delete(&system.SysUser{}).Error
		if err != nil {
			global.Logger.Error("删除用户失败", zap.Error(err))
			return errors.New("删除用户错误")
		}
		return nil
	})
}

func (u *UserService) GetUserInfo(id uint) (system.SysUser, error) {
	user := system.SysUser{}
	err := global.DB.Model(&system.SysUser{}).Where("id = ?", id).Preload("Role").Find(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Logger.Error("不存在对应用户", zap.Uint("id", id))
			return user, errors.New("用户不存在")
		}
		global.Logger.Error("查询错误", zap.Error(err))
		return user, errors.New("查询错误")
	}
	return user, nil
}

func (u *UserService) EditUser(user request.EditUser) error {
	db := global.DB.Model(&system.SysUser{})
	if user.UserName != "" {
		err := db.Where("id = ?", user.ID).Update("user_name", user.UserName).Error
		if err != nil {
			global.Logger.Error("修改错误", zap.Error(err))
			return errors.New("修改错误")
		}
	}
	if user.Password != "" {
		hash := utils.BcryptHash(user.Password)
		err := db.Where("id = ?", user.ID).Update("pass_word", hash).Error
		if err != nil {
			global.Logger.Error("修改错误", zap.Error(err))
			return errors.New("修改错误")
		}
	}
	return nil
}
