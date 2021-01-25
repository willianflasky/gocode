package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

// GetCommunityList xx
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetail xx
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
