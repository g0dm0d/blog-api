package main

import (
	"blog-api/config"
	"blog-api/pkg/cron"
	"blog-api/rest"
	"blog-api/service"
	"blog-api/store/postgres"
	"blog-api/tools/tokenmanager"
	"log"
)

func check(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	// get cfg
	cfg, err := config.Load()
	check(err)

	// connect to postgres
	db, err := postgres.New(cfg.Database.DSN)
	check(err)

	// Stores
	userStore := postgres.NewUserStore(db)
	sessionStore := postgres.NewSessionStore(db)
	articleStore := postgres.NewArticleStore(db)

	// Cron
	c := cron.NewCron()

	c.AddTask(cron.Task{
		Name: "Delete exp session",
		Schedule: cron.Schedule{
			IsDate:  true,
			Day:     0,
			Hours:   0,
			Minutes: 1,
			Seconds: 0,
		},
		Action: sessionStore.ClearExpSession,
	})

	c.Start()

	// Tools
	tokenManager := tokenmanager.New(cfg.Server.Secret)

	// Services
	services := service.New(service.Opts{
		UserStore:    userStore,
		SessionStore: sessionStore,
		ArticleStore: articleStore,
		Token:        *tokenManager,
		// need for store img
		AssetsDir: cfg.Server.AssetsDir,
	})

	// generate params for http server
	server := rest.NewServer(&rest.Config{
		Addr:         cfg.Server.Addr,
		Port:         cfg.Server.Port,
		Service:      services,
		TokenManager: tokenManager,
	})
	// setup routs
	server.SetupRouter()

	// run server
	err = server.RunServer()
	check(err)
}
