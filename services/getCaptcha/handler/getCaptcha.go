package handler

import (
	"context"
	"encoding/json"
	"getCaptcha/model"
	pb "getCaptcha/proto"
	"github.com/afocus/captcha"
)

type GetCaptcha struct{}

func (e *GetCaptcha) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	cap := captcha.New()
	cap.SetFont("./conf/comic.ttf")

	img, str := cap.Create(5, captcha.ALL)
	err := model.SaveImgCode(str, req.Uuid)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(img)
	if err != nil {
		return err
	}
	rsp.Img = bytes

	return nil
}
