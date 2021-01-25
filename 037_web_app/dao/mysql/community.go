package mysql

import (
	"database/sql"
	"web_app/models"

	"go.uber.org/zap"
)

// GetCommunityList xx
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := DB.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

// GetCommunityDetailByID xx
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select 
		community_id,
		community_name,
		introduction,
		create_time
		from community where community_id = ?`
	if err := DB.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
