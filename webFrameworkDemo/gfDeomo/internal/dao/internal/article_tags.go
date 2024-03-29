// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleTagsDao is the data access object for table article_tags.
type ArticleTagsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ArticleTagsColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleTagsColumns defines and stores column names for table article_tags.
type ArticleTagsColumns struct {
	ArticleId string //
	TagId     string //
}

//  articleTagsColumns holds the columns for table article_tags.
var articleTagsColumns = ArticleTagsColumns{
	ArticleId: "article_id",
	TagId:     "tag_id",
}

// NewArticleTagsDao creates and returns a new DAO object for table data access.
func NewArticleTagsDao() *ArticleTagsDao {
	return &ArticleTagsDao{
		group:   "default",
		table:   "article_tags",
		columns: articleTagsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleTagsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleTagsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleTagsDao) Columns() ArticleTagsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleTagsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleTagsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleTagsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
