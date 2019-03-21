package routers

import (
	"LogView/controllers"
	"github.com/astaxie/beego"
)

func init() {
	loginController :=&controllers.LoginController{}
    beego.Router("/login", loginController,"get:Login")
	beego.Router("/login/post", loginController,"post:Post")
	beego.Router("/getusername", loginController,"post:GetUserName")
	beego.Router("/logout", loginController,"post:Logout")
	mainController :=&controllers.MainController{}
	beego.Router("/", mainController,"get:Get")
	beego.Router("/getDomains", mainController,"post:GetDomains")
	beego.Router("/getMachines", mainController,"post:GetMachines")
	beego.Router("/getLogsName", mainController,"post:GetLogsName")
	beego.Router("/getLogs", mainController,"post:GetLogContents")
	websocketController :=&controllers.WebSocketController{}
	beego.Router("/ws", websocketController, "get:WsHandler")
	serverController :=&controllers.ServerController{}
	beego.Router("/list", serverController, "get:Get")
	beego.Router("/getEnv", serverController, "post:GetEnv")
	beego.Router("/getApplications", serverController, "post:GetApplications")
	beego.Router("/getHosts", serverController, "post:GetHosts")
	beego.Router("/getExecLogs", serverController, "post:GetExecLogs")
	userController := &controllers.UserController{}
	beego.Router("/admin/user/edit", userController, "get:Edit")
	beego.Router("/admin/user/delete", userController, "post:Delete")
	beego.Router("/admin/user/list", userController, "get:List")
	beego.Router("/admin/user/add", userController, "get:Add")
	beego.Router("/admin/user/post", userController, "post:Post")
	beego.Router("/admin/user/update", userController, "post:Update")
	softwareController :=&controllers.SoftwareController{}
	beego.Router("/admin/software/edit", softwareController, "get:Edit")
	beego.Router("/admin/software/delete", softwareController, "post:Delete")
	beego.Router("/admin/software/list", softwareController, "get:List")
	beego.Router("/admin/software/add", softwareController, "get:Add")
	beego.Router("/admin/software/post", softwareController, "post:Post")
	beego.Router("/admin/software/update", softwareController, "post:Update")
	perjectController := &controllers.PerjectController{}
	beego.Router("/admin/perject/edit", perjectController, "get:Edit")
	beego.Router("/admin/perject/delete", perjectController, "post:Delete")
	beego.Router("/admin/perject/list", perjectController, "get:List")
	beego.Router("/admin/perject/add", perjectController, "get:Add")
	beego.Router("/admin/perject/post", perjectController, "post:Post")
	beego.Router("/admin/perject/update", perjectController, "post:Update")
	envController := &controllers.EnvController{}
	beego.Router("/admin/env/edit", envController, "get:Edit")
	beego.Router("/admin/env/delete", envController, "post:Delete")
	beego.Router("/admin/env/list", envController, "get:List")
	beego.Router("/admin/env/add", envController, "get:Add")
	beego.Router("/admin/env/post", envController, "post:Post")
	beego.Router("/admin/env/update", envController, "post:Update")
	applicationController := &controllers.ApplicationController{}
	beego.Router("/admin/app/edit",applicationController, "get:Edit")
	beego.Router("/admin/app/delete", applicationController, "post:Delete")
	beego.Router("/admin/app/list", applicationController, "get:List")
	beego.Router("/admin/app/add", applicationController, "get:Add")
	beego.Router("/admin/app/post", applicationController, "post:Post")
	beego.Router("/admin/app/update", applicationController, "post:Update")
	healthController :=&controllers.InspectionController{}
	beego.Router("/admin/health/edit",healthController, "get:Edit")
	beego.Router("/admin/health/delete", healthController, "post:Delete")
	beego.Router("/admin/health/list", healthController, "get:List")
	beego.Router("/admin/health/add", healthController, "get:Add")
	beego.Router("/admin/health/post", healthController, "post:Post")
	beego.Router("/admin/health/update", healthController, "post:Update")
	beego.Router("/health", healthController, "get:Get")
	beego.Router("/getHealth", healthController, "post:GetHealth")
}
