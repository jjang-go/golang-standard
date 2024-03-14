package network

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jjang-go/golang-standard/types"
)

var (
	userRouterInit     sync.Once // 해당 라우터 한번만 호출 시키기 위함
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		router.registerPOST("/", userRouterInstance.create)
		router.registerGET("/", userRouterInstance.get)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
	})

	return userRouterInstance
}

// CRUD
func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create.")
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get.")

	//response
	u.router.okResponse(c, &types.UserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "성공입니다.",
		},
		User: nil,
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update.")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete.")
}
