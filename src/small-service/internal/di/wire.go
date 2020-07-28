// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"small-service/internal/dao"
	"small-service/internal/service"
	"small-service/internal/server/grpc"
	"small-service/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
