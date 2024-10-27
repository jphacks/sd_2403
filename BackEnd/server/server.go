package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"sushi/controllor"
	"sushi/service"
	"sushi/utils/DB"
	"sushi/utils/config"
)

type Server struct {
	config     *config.Config
	log        *logrus.Logger
	router     *gin.Engine
	controller *controllor.Controller
	service    *service.Service
	db         *DB.DB
}

func CreateServer() *http.Server {

	conf, err := config.NewConfig()
	if err != nil {
		panic(any("error reading config.yaml, " + err.Error()))
	}

	log := logrus.New()
	log.Out = os.Stdout
	log.Level = conf.LogLevel()

	if conf.LogFileLocation() == "" {
		log.Fatal("missing log_file_location config.yaml variable")
	}
	logfile, err := os.OpenFile(conf.LogFileLocation(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("failed to open file for logging")
	} else {
		log.Out = logfile
		log.Formatter = &logrus.JSONFormatter{}
	}

	/*
		Initialize Server
	*/
	svr := NewServer(conf, log)

	/*
		Initialize DB
	*/

	svr.db = DB.NewDB_MySQL(log, conf.DBConnectionPath())
	/*
		Initialize Services
	*/
	closers := svr.NewService()

	/*
		Initialize Controllers
	*/
	svr.controller = controllor.NewControllor(svr.service, log, conf)

	socketServer := newSocketServer()
	/*
		Initialize Router
	*/
	svr.router = NewRouter(svr, *conf, socketServer)

	/*
		Start HTTP Server
	*/
	// initialize server
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", conf.HTTPPort())
	httpServer := makeHttpServer(addr, svr.router)

	closers = append(closers, socketServer)
	// handle graceful shutdown
	go handleGracefulShutdown(httpServer, closers)

	return httpServer
}

func NewServer(conf *config.Config, log *logrus.Logger) *Server {
	return &Server{
		config: conf,
		log:    log,
	}
}

func Start() error {
	srv := CreateServer()

	// listen and serve
	err := srv.ListenAndServe()
	if err == http.ErrServerClosed {
		log.Println("server shutting down gracefully...")
	} else {
		log.Println("unexpected server shutdown...")
		log.Println("ERR: ", err)
	}
	return err
}
