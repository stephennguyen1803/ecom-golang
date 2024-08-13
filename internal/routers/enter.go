package routers

import (
	"ecom-project/internal/routers/manage"
	"ecom-project/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
