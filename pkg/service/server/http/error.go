package http

import (
	"net/http"

	"github.com/geoirb/event/pkg/service"
)

var errorList map[error]int = map[error]int{
	service.ErrNotFound: http.StatusNotFound,
}

func getStatusCode(err error) int{
	if err != nil {
		if code, isExist := errorList[err]; isExist{
			return code
		}
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
