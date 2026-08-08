package maids

import (
	"butler/internal/model/dto"
	maiddto "butler/internal/model/dto/maid"
	"butler/internal/service/maids"
	"common/public/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterMaid 注册女仆节点
// POST /butler/maids/register
func RegisterMaid(c *gin.Context) error {
	var req maiddto.RegisterMaidReq
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		return pkg.BadArgumentsError(err.Error())
	}

	maidId, err := maids.RegisterMaid(&req)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, pkg.DataResp[maiddto.RegisterMaidResp]{
		Data: &maiddto.RegisterMaidResp{
			MaidId: dto.Uint64ToString(maidId),
		},
		HttpResp: pkg.SuccessMessageResp(""),
	})
	return nil
}

// DismissMaid 解除注册
// POST /butler/maids/{maidId}/dismiss
func DismissMaid(c *gin.Context) error {
	maidIdStr := c.Request.PathValue("maidId")
	maidId, parseErr := strconv.ParseUint(maidIdStr, 10, 64)
	if parseErr != nil {
		return pkg.BadArgumentsError(parseErr.Error())
	}

	err := maids.DismissMaid(maidId)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, pkg.SuccessMessageResp(""))
	return nil
}

// UpdateMaid 更新女仆节点信息
// POST /butler/maids/{maidId}/update
func UpdateMaid(c *gin.Context) error {
	maidIdStr := c.Request.PathValue("maidId")
	maidId, parseErr := strconv.ParseUint(maidIdStr, 10, 64)
	if parseErr != nil {
		return pkg.BadArgumentsError(parseErr.Error())
	}

	var req maiddto.UpdateMaidReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return pkg.BadArgumentsError(err.Error())
	}

	err = maids.UpdateMaid(maidId, &req)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, pkg.SuccessMessageResp(""))
	return nil
}

// ListMaids 查看女仆节点列表
// GET /butler/maids/list
func ListMaids(c *gin.Context) error {
	var req maiddto.ListMaidsReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		return pkg.BadArgumentsError(err.Error())
	}

	resp, err := maids.ListMaids(&req)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, resp)
	return nil
}