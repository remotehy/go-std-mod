package business

import (
	"github.com/google/wire"

	"github.com/remotehy/go-std-mod/app/tom/data"
)

var Provider = wire.NewSet(
	NewBusiness,
	wire.Bind(new(data.DataAccess), new(*data.DataAccessImpl)),
	data.Provider)

type Business interface {
	Greet() string
}

func NewBusiness(dao data.DataAccess) *BusinessImpl {
	return &BusinessImpl{dao: dao}
}

type BusinessImpl struct {
	dao data.DataAccess
}

func (b *BusinessImpl) Greet() string {
	return b.dao.GetData()
}
