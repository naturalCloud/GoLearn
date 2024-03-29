// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Categories is the golang structure of table categories for DAO operations like Where/Data.
type Categories struct {
	g.Meta    `orm:"table:categories, do:true"`
	Id        interface{} //
	ParentId  interface{} //
	Lft       interface{} //
	Rgt       interface{} //
	Depth     interface{} //
	Name      interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
