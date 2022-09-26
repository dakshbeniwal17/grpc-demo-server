package daos

import (
	"context"
	"log"

	"grpc-demo-server/beans"
	models "grpc-demo-server/repository"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Create(ctx context.Context, exec boil.ContextExecutor, bean *beans.CreateHostParams) (*models.Host, error) {
	host := &models.Host{
		Name:    bean.Name,
		RAM:     null.Float64From(bean.Ram),
		Cores:   null.IntFrom(bean.Cores),
		Threads: null.IntFrom(bean.Threads),
		Ready:   null.IntFrom(bean.Ready),
		JobURL:  null.StringFrom(bean.JobUrl),
		Token:   bean.Token,
		Gpu:     null.StringFrom(bean.Gpu),
	}
	// fmt.Printf("host to be inserted: %v\n", *host)
	if err := host.Insert(ctx, exec, boil.Infer()); err != nil {
		log.Printf("error while creating host: %v", err)
		return nil, err
	}
	return host, nil
}
