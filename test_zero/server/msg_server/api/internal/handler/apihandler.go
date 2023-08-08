package handler

import (
	"net/http"
	"station/common"

	"test_zero/server/msg_server/api/internal/logic"
	"test_zero/server/msg_server/api/internal/svc"
	"test_zero/server/msg_server/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ApiHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			common.RespErr(w, r, err)
			return
		}
		httpx.GetRemoteAddr(r)
		l := logic.NewApiLogic(r.Context(), ctx)
		resp, err := l.Api(req)
		if err != nil {
			common.RespErr(w, r, err)
		} else {
			common.RespOk(w, r, resp)
		}
	}
}
