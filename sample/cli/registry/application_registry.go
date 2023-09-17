package registry

import "github.com/onnga-wasabi/go-ddd/sample/application"

type ApplicationRegistry interface {
	NewUserApplicationService() application.UserApplicationService
}
