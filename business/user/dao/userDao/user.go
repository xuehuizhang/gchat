package userDao

import (
	"com.gchat.business/user/dao"
	"com.gchat.business/user/model"
	"fmt"
	"gorm.io/gorm"
)

func Add(user *model.User) error {
	return dao.UDB.Create(user).Error
}

func Delete(id int64) error {
	return dao.UDB.Delete(&model.User{}, id).Error
}

func UpdateMap(id int64, m map[string]interface{}) error {
	return dao.UDB.Model(&model.User{}).Where("id=?", id).Updates(m).Error
}

func SaveOrUpdate(user *model.User) error {
	return dao.UDB.Save(user).Error
}

func GetOne(id int64) (*model.User, error) {
	var user *model.User
	err := dao.UDB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetByField(filedName string, value interface{}) (*model.User, error) {
	var user *model.User
	err := dao.UDB.First(&user, fmt.Sprintf("%v= ?", filedName), value).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return user, nil
}

func Count(sql string, params []interface{}) (int64, error) {
	var count int64
	err := dao.UDB.Model(&model.User{}).Where(sql, params).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func Page(selectFileds []string, sql string, params []interface{}, order string, offset, pageSize int) ([]*model.User, error) {
	var users []*model.User
	err := dao.UDB.Select(selectFileds).Where(sql, params).Order(order).Offset(offset).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Raw(sql string, params []interface{}) ([]*model.User, error) {
	var users []*model.User
	err := dao.UDB.Raw(sql, params...).Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
