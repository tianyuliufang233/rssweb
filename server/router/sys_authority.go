package router

import (
	"github.com/gin-gonic/gin"
	"server/api/v1"
)

func InitAuthorityRouter(Router *gin.RouterGroup){
	AuthorityRouter :=Router.Group("authority")
	{
		AuthorityRouter.POST("createAuthority",v1.CreateAuthority) // 创建角色
		AuthorityRouter.POST("getAuthorityList", v1.GetAuthorityList) // 获取角色列表
	}
}