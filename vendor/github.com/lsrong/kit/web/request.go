package web

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func GetIntParam(ctx context.Context, name string) int {
	v := GetParam(ctx, name)
	intV, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}

	return intV
}

func GetParam(ctx context.Context, name string) string {
	params := httprouter.ParamsFromContext(ctx)

	return params.ByName(name)
}

func Decode(r *http.Request, dest interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dest); err != nil {
		return err
	}

	return nil
}
