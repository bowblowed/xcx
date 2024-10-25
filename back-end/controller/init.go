package controller

import (
	"back-end/config"
	"fmt"
)

var baseUrl = fmt.Sprintf("http://%s:%d", config.Server.Ip, config.Server.Port)
