package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	blockhttp "github.com/alexvelfr/tmp/delivery/http"
	"github.com/alexvelfr/tmp/models"
	"github.com/alexvelfr/tmp/testmon"
	"github.com/alexvelfr/tmp/testmon/blockrepo"
	"github.com/alexvelfr/tmp/testmon/notificator"
	"github.com/alexvelfr/tmp/testmon/recipientrepo"
	"github.com/alexvelfr/tmp/testmon/uc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	uc     testmon.Usecase
	cancle context.CancelFunc
	db     *gorm.DB
	ctx    context.Context
}

func NewApp(ctx context.Context, cancle context.CancelFunc) *App {
	res := &App{}
	db := res.getDBConnect()
	blockRepo := blockrepo.NewBlockRepo(db)
	recipientRepo := recipientrepo.NewBlockRepo(db)
	notify := notificator.NewNotificator()
	usecase := uc.NewUC(
		recipientRepo,
		blockRepo,
		notify,
	)
	res.uc = usecase
	res.db = db
	res.cancle = cancle
	res.ctx = ctx
	return res
}

func (a *App) Run() {
	go func() {
		defer a.cancle()
		a.db.AutoMigrate(
			&models.Block{},
			&models.Recipient{},
		)
		a.RunTasks()

		router := gin.Default()
		api := router.Group("api/")
		blockhttp.RegisterBlockEndpoints(api, a.uc)

		srv := &http.Server{
			Addr:    viper.GetString("app.port"),
			Handler: router,
		}
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-c
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()
}

func (a *App) RunTasks() {
	go func() {
		for {
			select {
			case <-a.ctx.Done():
				return
			default:
				a.uc.CheckReglamentBlocks()
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func (a *App) getDBConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(viper.GetString("app.db.name")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
