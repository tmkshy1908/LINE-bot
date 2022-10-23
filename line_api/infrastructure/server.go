package infrastructure

import (
	// "github.com/gorilla/mux"

	"github.com/gorilla/mux"

	db "github.com/tmkshy1908/LINE-bot/infrastructure/db"
)

type ControllHandler struct {
	// Common *interfaces.CommonController
}

func NewServer(h db.SqlConnHandler) (handler *mux.Router) {
	ch := &ControllHandler{
		//  Common: interfaces.NewController(h),
	}

	handler = NewRouter(ch)
	return
}
