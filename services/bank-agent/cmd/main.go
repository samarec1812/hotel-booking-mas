package main

import (
	"bytes"
	"context"
	"github.com/go-co-op/gocron"
	"github.com/goccy/go-json"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	bank_agent "github.com/samarec1812/hotel-booking-mas/services/bank-agent"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/models"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/pkg/handler"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/pkg/repository"
	"github.com/samarec1812/hotel-booking-mas/services/bank-agent/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	subSubjectName       = "PAYMENT.created"
	pubSubjectName       = "PAYMENT.response"
	urlCreateTransaction = "http://localhost:8082/bank/transaction/create"
)

func printMsg(m *nats.Msg, i int) {
	logrus.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}

func Transaction(m *nats.Msg) {

}

//func Subscriber() {
//	// Connect to NATS
//
//	nc, err := nats.Connect(nats.DefaultURL)
//	if err != nil {
//		log.Fatal(err)
//	}
//	js, err := nc.JetStream()
//	if err != nil {
//		log.Fatal(err)
//	}
//	i := 0
//	js.Subscribe(subSubjectName, func(msg *nats.Msg) {
//		i += 1
//		printMsg(msg, i)
//		var payment models.Payment
//		err := json.Unmarshal(msg.Data, &payment)
//		if err != nil {
//			logrus.Printf("[#%d] [%s]: Error in payment data'%s'", i, msg.Subject)
//			resp := Response{
//				Status: "error",
//			}
//			reviewOrder(js, resp)
//		}
//		printMsg(msg, i)
//	})
//	nc.Flush()
//
//	if err := nc.LastError(); err != nil {
//		log.Fatal(err)
//	}
//
//	logrus.Printf("Listening on [%s]", subSubjectName)
//
//	runtime.Goexit()
//
//}

func Subscriber() {
	// Connect to NATS
	nc, _ := nats.Connect(nats.DefaultURL)
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Create Pull based consumer with maximum 128 inflight.
	// PullMaxWaiting defines the max inflight pull requests.
	sub, _ := js.PullSubscribe(subSubjectName, "order-review", nats.PullMaxWaiting(128))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		msgs, _ := sub.Fetch(10, nats.MaxWait(5*time.Minute))
		for i, msg := range msgs {
			msg.Ack()
			var payment models.Payment
			err := json.Unmarshal(msg.Data, &payment)
			if err != nil {
				logrus.Printf("[#%d] [%s]: Error in payment data", i, msg.Subject)
				reviewOrder(js, Response{
					Status: "bad",
				})
			}
			resp, err := http.Post(urlCreateTransaction, "application/json", bytes.NewBuffer(msg.Data))
			if err != nil {
				logrus.Fatalf("bank server: send bad request: %s", err.Error())
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logrus.Fatalf("bank server: error read response body: %s", err.Error())
			}
			var respToService = Response{
				Status: "not have money on balance",
			}
			err = json.Unmarshal(body, &respToService)
			if err != nil {
				logrus.Fatalf("bank server: error unmarshal response body: %s", err.Error())
			}
			//log.Println("order-review service")

			printMsg(msg, i)
			//log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n", order.OrderID, order.CustomerID, order.Status)
			reviewOrder(js, respToService)
		}
	}
}

type Response struct {
	Status string `json:"status"`
}

func reviewOrder(js nats.JetStreamContext, resp Response) {
	// Changing the Order status

	respJSON, _ := json.Marshal(resp)
	_, err := js.Publish(pubSubjectName, respJSON)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Printf("Send reponse with status: %s\n", resp.Status)
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		//Password: os.Getenv("DB_PASSWORD"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(bank_agent.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running server: %s", err.Error())
		}
	}()
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Minute().Do(Subscriber)

	s.StartAsync()

	// wait signal to shutdown server with a timeout
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Println("Shutting down server. ")

}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
