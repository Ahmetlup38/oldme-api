// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"oldme-api/internal/dao/internal"
)

// internalArticleGroupDao is internal type for wrapping internal DAO implements.
type internalArticleGroupDao = *internal.ArticleGroupDao

// articleGroupDao is the data access object for table article_group.
// You can define custom methods on it to extend its functionality as you wish.
type articleGroupDao struct {
	internalArticleGroupDao
}

var (
	// ArticleGroup is globally public accessible object for table article_group operations.
	ArticleGroup = articleGroupDao{
		internal.NewArticleGroupDao(),
	}
)

// Fill with you ideas below.
