package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-pantheon/janus/app/gate/internal/conf"
	"github.com/google/wire"
	"github.com/pkg/errors"
	etcdclient "go.etcd.io/etcd/client/v3"
)

var ProviderSet = wire.NewSet(NewTCPServer, NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) (registry.Registrar, error) {
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: conf.Etcd.Endpoints,
		Username:  conf.Etcd.Username,
		Password:  conf.Etcd.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "create etcd client failed")
	}

	return etcd.New(client), nil
}
