package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	getCaptcha "ihome/web/proto"
	"image/png"

	"github.com/gin-gonic/gin"
	"ihome/web/utils"

	"net/http"
)

func GetSession(ctx *gin.Context) {
	resp := make(map[string]string)

	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	ctx.JSON(http.StatusOK, resp)

}

func GetImageCd(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	reg := consul.NewRegistry()
	service := micro.NewService(
		micro.Registry(reg),
	)
	microClient := getCaptcha.NewGetCaptchaService("getcaptcha", service.Client())
	response, err := microClient.Call(context.TODO(), &getCaptcha.CallRequest{Uuid: uuid})
	if err != nil {
		return
	}

	var img captcha.Image
	json.Unmarshal(response.Img, &img)
	png.Encode(ctx.Writer, img)
	fmt.Println("uuid = ", uuid)

}
