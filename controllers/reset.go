package controllers

import (
	"github.com/lfq7413/tomato/errs"
	"github.com/lfq7413/tomato/rest"
	"github.com/lfq7413/tomato/types"
)

// ResetController 处理 /requestPasswordReset 接口的请求
type ResetController struct {
	ObjectsController
}

// HandleResetRequest 处理通过 email 重置迷密码的请求
// @router / [post]
func (r *ResetController) HandleResetRequest() {
	if r.JSONBody == nil && r.JSONBody["email"] == nil {
		r.Data["json"] = errs.ErrorMessageToMap(errs.EmailMissing, "you must provide an email")
		r.ServeJSON()
		return
	}
	email := r.JSONBody["email"].(string)
	err := rest.SendPasswordResetEmail(email)
	if err != nil {
		r.Data["json"] = errs.ErrorToMap(err)
		r.ServeJSON()
		return
	}

	r.Data["json"] = types.M{}
	r.ServeJSON()
}

// Get ...
// @router / [get]
func (r *ResetController) Get() {
	r.ObjectsController.Get()
}

// Delete ...
// @router / [delete]
func (r *ResetController) Delete() {
	r.ObjectsController.Delete()
}

// Put ...
// @router / [put]
func (r *ResetController) Put() {
	r.ObjectsController.Put()
}
