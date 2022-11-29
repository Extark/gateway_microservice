package main

import (
	"encoding/json"
	"github.com/extark/gateway_microservice/models"
	"github.com/extark/gateway_microservice/utils"
	"github.com/extark/go_jwt_auth"
	stargate "github.com/realbucksavage/stargate"
	"github.com/realbucksavage/stargate/listers"
	"github.com/samber/lo"
	"log"
	"net/http"
)

var confs []*models.ConfigJsonFormat

func main() {
	//Init the settings values inside a global var
	err := utils.InitSettings()
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	//read the conf of the routes
	confs, err = utils.ReadConf("config.json")
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	//creates the routes map
	routes := make(map[string][]string)
	for _, conf := range confs {
		routes[conf.Route] = conf.Nodes
	}

	//instance the listers
	listeners := listers.Static{Routes: routes}

	//instances a jwtCasbinAuthMiddleware
	gateway, err := stargate.NewRouter(listeners, stargate.WithMiddleware(jwtCasbinAuthMiddleware))
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatal(http.ListenAndServe(":8080", gateway))
}

func jwtCasbinAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj := r.URL.Path
		conf, ok := lo.Find[*models.ConfigJsonFormat](confs, func(i *models.ConfigJsonFormat) bool {
			return i.Route == obj
		})

		if !ok {
			log.Print("ERROR: request not found")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.StandardError{Error: "request not found"})
			return
		}
		if conf.Auth {
			go_jwt_auth.JwtCasbinAuthMiddleware(next, utils.Cfg.CASBINADAPTER, utils.Cfg.SECRET)
			return
		}

		next.ServeHTTP(w, r)
	})
}
