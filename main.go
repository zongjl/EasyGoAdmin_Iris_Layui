// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package main

import (
	_ "easygoadmin/boot"
	"easygoadmin/conf"
	"easygoadmin/router"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	// 创建app结构体对象
	app := iris.New()
	// 设置调试模式
	app.Logger().SetLevel("debug")
	// 可选项添加两个内置的句柄（handlers）
	// 捕获相对于http产生的异常行为
	app.Use(recover.New())
	//记录请求日志
	app.Use(logger.New())
	// 初始化配置
	config := iris.WithConfiguration(iris.YAML("./conf/config.yml"))
	// 路由注册
	router.RegisterRouter(app)

	//// 谓词:   GET
	//// 资源: http://localhost:8080
	//app.Handle("GET", "/", func(ctx iris.Context) {
	//	ctx.HTML("<h1>Welcome</h1>")
	//})

	//app.Get("/message", func(ctx iris.Context) {
	//	// 绑定kv
	//	ctx.ViewData("message", "Hello world!")
	//	// 渲染模板文件./views.hello.html
	//	ctx.View("hello.html")
	//})

	//// 职级管理
	//level := app.Party("/message2")
	//{
	//	level.Get("/index", func(ctx iris.Context) {
	//		// 绑定kv
	//		ctx.ViewData("message", "Hello world!")
	//		// 渲染模板文件./views.hello.html
	//		ctx.View("level/edit.html")
	//	})
	//}

	//
	//// 等价于 app.Handle("GET", "/ping", [...])
	//// 谓词:   GET
	//// 资源: http://localhost:8080/ping
	//app.Get("/ping", func(ctx iris.Context) {
	//	ctx.WriteString("ping")
	//})
	//
	//// 谓词:   GET
	//// 资源: http://localhost:8080/hello
	//app.Get("/hello", func(ctx iris.Context) {
	//	ctx.JSON(iris.Map{"message": "Hello Iris!"})
	//})
	//
	//// http://localhost:8080
	//// http://localhost:8080/ping
	//// http://localhost:8080/hello
	//// Run 方法第二个参数为应用的配置参数
	//app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

	// Listens and serves incoming http requests
	// on http://localhost:8081.
	fmt.Println()
	app.Run(iris.Addr(conf.CONFIG.EGAdmin.Addr+":"+conf.CONFIG.EGAdmin.Port), config)
}
