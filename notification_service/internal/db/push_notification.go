package db

import (
	_ "embed"
)

//go:embed queries/getPostLikeCount.sql
var getPostLikeCountQuery string

func (r *Repository) GetPostLikeCount(postId string) (int, error) {
	var likeCount int
	err := r.db.Get(&likeCount, getPostLikeCountQuery, postId)
	return likeCount, err
}
