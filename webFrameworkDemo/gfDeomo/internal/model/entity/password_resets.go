// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PasswordResets is the golang structure for table password_resets.
type PasswordResets struct {
	Email     string      `json:"email"     description:""` //
	Token     string      `json:"token"     description:""` //
	CreatedAt *gtime.Time `json:"createdAt" description:""` //
}
