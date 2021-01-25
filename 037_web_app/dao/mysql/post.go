package mysql

import (
	"strings"
	"web_app/models"

	"github.com/jmoiron/sqlx"
)

// CreatePost 创建贴子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `
		insert into post(post_id, title, content, author_id, community_id)
		values(?,?,?,?,?)
	`
	_, err = DB.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 根据ID查询单个贴子数据
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time 
				from post where post_id = ?`
	err = DB.Get(post, sqlStr, pid)
	return
}

// GetPostList 查询贴子列表函数
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time 
	from post 
	ORDER BY create_time
	DESC
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2) //不要写成：make([]*models.Post, 2)
	err = DB.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// 根据给定的ID列表查询贴子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id,community_id,create_time from post
		where post_id in (?)
		order by FIND_IN_SET(post_id, ?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	query = DB.Rebind(query)
	DB.Select(postList, query, args...)
	return
}
