package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Unknwon/macaron"
)

func GetUsersV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "Get V1 users,successfully"})
	return http.StatusOK, result
}

func PostUsersV1Handler(ctx *macaron.Context) (int, []byte) {
	result, _ := json.Marshal(map[string]string{"message": "Post V1 users,unauthorization"})
	return http.StatusUnauthorized, result
}