package usecase

// import (
// 	"context"
// 	"fmt"

// 	"github.com/tmkshy1908/LINE-bot/infrastructure/pdb"
// )

// type CommonRepository struct {
// 	DB pdb.SqlHandler
// }

// const (
// 	SELECT_SUCHEDULE string = "select id, day, contents from schedule;"
// 	INSERT_SUCHEDULE string = "insert into schedule (id,day,contents) values ($1,$2,$3)"
// 	UPDATE_SUCHEDULE string = "update schedule set day = $2, contents = $3 where id = $1"
// 	DELETE_SUCHEDULE string = "delete from schedule where id = $1"
// )

// func (r *CommonRepository) Get(ctx context.Context) (err error) {
// 	// rows, err := r.DB.Query(ctx, SELECT_SUCHEDULE)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }
