package main

import (
	"github.com/extark/gateway_microservice/utils"
	"github.com/extark/go_jwt_auth"
	stargate "github.com/realbucksavage/stargate"
	"github.com/realbucksavage/stargate/listers"
	"log"
	"net/http"
)

func main() {
	//Init the settings values inside a global var
	err := utils.InitSettings()
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	listeners := listers.Static{
		Routes: map[string][]string{
			"/": {},
		},
	}

	gateway, err := stargate.NewRouter(listeners, stargate.WithMiddleware(jwtCasbinAuthMiddleware))
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatal(http.ListenAndServe(":8080", gateway))
}

func jwtCasbinAuthMiddleware(next http.Handler) http.Handler {
	return go_jwt_auth.JwtCasbinAuthMiddleware(next, utils.Cfg.CASBINADAPTER, utils.Cfg.SECRET)
}
