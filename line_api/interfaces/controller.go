package interfaces

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/tmkshy1908/LINE-bot/infrastructure/pdb"
// )

// type CommonController struct {
// 	Interactor CommonInteractor
// 	Converter  CommonConverter
// }

// func NewController(SqlHandler pdb.SqlHandler) (cc *CommonController) {
// 	cc = &CommonController{
// 		Interactor: &usecase.CommonInteractor{
// 			CommonRepository: &CommonRepository{
// 				DB: SqlHandler,
// 			},
// 		},
// 	}
// 	cc.Converter = NewConvertController()
// 	return
// }

// func (cc *CommonController) SampleHandler(w http.ResponseWriter, r *http.Request) (err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
// 	defer cancel()

// 	resp, err := cc.Interactor.UseCaseSampleRepository(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }
