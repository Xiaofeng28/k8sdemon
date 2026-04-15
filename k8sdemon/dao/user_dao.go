package dao

import (
	"fmt"
	"go_code/k8sdemon/config"
	"go_code/k8sdemon/models"
)

type UserDao struct {
}

// FindFistUser 查询用户 by id
func FindFistUser() models.User {
	var user models.User
	config.DB.Where("user_id = ?", 1).First(&user)
	return user
}

// GetAllUserInfo 获取所有用户信息
func GetAllUserInfo() ([]models.User, error) {
	var userSlice []models.User
	// 查询所有用户记录
	result := config.DB.Find(&userSlice)
	// 检查是否有错误
	if result.Error != nil {
		return nil, result.Error
	}
	// 返回查询结果
	return userSlice, nil
}

// AddOneUser 添加一个用户
func AddOneUser(user models.User) (int, error) {
	// 添加
	result := config.DB.Create(&user)
	// 检查是否成功
	if result.Error != nil {
		fmt.Printf("创建失败: %v\n", result.Error)
		return 0, result.Error
	}
	// 返回插入的条数
	return int(result.RowsAffected), nil
}

// DeleteOneUserByUserID 删除用户根据用户id
func DeleteOneUserByUserID(userid int) (int, error) {
	// 删除用户
	result := config.DB.Delete(&models.User{}, userid)
	if result.Error != nil {
		fmt.Println("DeleteOneUserByUserID() 函数删除用户失败！")
		return 0, result.Error
	}
	// 删除成功
	return int(result.RowsAffected), nil
}

// UpdateOneUserInfoById 根据用户id更新用户信息
func UpdateOneUserInfoById(user models.User) (int, error) {
	// 更新用户
	result := config.DB.Model(&user).Updates(user)
	if result.Error != nil {
		fmt.Println("UpdateOneUserInfoById() 函数更新用户失败！")
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}
