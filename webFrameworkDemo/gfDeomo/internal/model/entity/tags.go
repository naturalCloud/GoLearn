// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tags is the golang structure for table tags.
type Tags struct {
	Id        uint        `json:"id"        description:""`     //
	TagName   string      `json:"tagName"   description:"标签名字"` // 标签名字
	CreatedAt *gtime.Time `json:"createdAt" description:""`     //
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`     //
}