package dto

import (
	"common/public/pkg"
	"time"
)

// RegisterMaidReq 注册女仆节点请求
type RegisterMaidReq struct {
	MaidName    string `json:"maidName" binding:"required"`
	HostPort    string `json:"hostPort" binding:"required"`
	AccessSecret string `json:"accessSecret" binding:"required"`
}

// RegisterMaidResp 注册女仆节点响应
type RegisterMaidResp struct {
	MaidId string `json:"maidId"`
}

// UpdateMaidReq 更新女仆节点信息请求
type UpdateMaidReq struct {
	MaidName    string `json:"maidName" binding:"required"`
	HostPort    string `json:"hostPort" binding:"required"`
	AccessSecret string `json:"accessSecret" binding:"required"`
}

// ListMaidsReq 查询女仆节点列表请求
type ListMaidsReq struct {
	MaidName string `json:"maidName" form:"maidName"`
	pkg.PageableReq
}

// ListMaidsResp 查询女仆节点列表响应项
type ListMaidsResp struct {
	MaidId      string     `json:"maidId"`
	MaidName    string     `json:"maidName"`
	HostPort    string     `json:"hostPort"`
	AccessSecret string     `json:"accessSecret"`
	Enabled     int8       `json:"enabled"`
	DeletedTime *time.Time `json:"deletedTime"`
	CreateTime  time.Time  `json:"createTime"`
	OwnerId     string     `json:"ownerId"`
	UpdateTime  time.Time  `json:"updateTime"`
}