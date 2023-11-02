package service

import (
	"fmt"

	"github.com/google/wire"

	"github.com/remotehy/go-std-mod/app/tom/business"
)

type Service interface {
	HandleRequest(name string)
}

var Provider = wire.NewSet(
	wire.Bind(new(Service), new(*ServiceImpl)),
	NewService,
	wire.Bind(new(business.Business), new(*business.BusinessImpl)),
	business.Provider,
	wire.Bind(new(Thinger), new(MyThinger)),
	wire.Value(Foo(1)),
	wire.Value(Bar("hello")),
	wire.Struct(new(MyThinger), "MyFoo", "MyBar"))

// NewService 一般就是直接返回一个 concrete type 类型, 而不是一个接口类型
// 在 Go 语言体系中, 一般不是提前定义接口, 而是在必要的时候才会去定义接口的
func NewService(buz business.Business, th Thinger) *ServiceImpl {
	return &ServiceImpl{buz: buz, th: th}
}

type ServiceImpl struct {
	buz business.Business
	th  Thinger
}

func (s *ServiceImpl) HandleRequest(name string) {
	fmt.Printf("hello %s, response: %s\n", name, s.buz.Greet())
	s.th.Hello()
}

type Foo int
type Bar string
type Thinger interface {
	Hello()
}

type MyThinger struct {
	MyFoo Foo
	MyBar Bar
}

func (MyThinger) Hello() {
	fmt.Println("my-thinger")
}
