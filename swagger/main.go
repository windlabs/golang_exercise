package main

import (
	"flag"
	"fmt"
	"windlabs.com/base/swagger/controller"
)

// @title 这里写标题`
// @version 1.0`
// @description xxxx述信息`
// @termsOfService [http://swagger.io/terms/](http://swagger.io/terms/)`
// @contact.name xxxx这里写联系人信息`
// @contact.url [http://www.swagger.io/support](http://www.swagger.io/support)`
// @contact.email support@swagger.io`
// @license.name Apache 2.0`
// @license.url [http://www.apache.org/licenses/LICENSE-2.0.html](http://www.apache.org/licenses/LICENSE-2.0.html)`
// @host 这里写接口服务的host`
// @BasePath 这里写base path`

func main(){
	var configPath string
	flag.StringVar(&configPath, "config", "./configs", "config path")
	flag.Parse()
	fmt.Println(configPath)
	controller.PayList()
}
