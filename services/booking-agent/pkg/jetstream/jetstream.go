package jetstream

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/samarec1812/hotel-booking-mas/services/booking-agent/models"
	"github.com/sirupsen/logrus"
	"log"
)

const (
	streamName     = "PAYMENT"
	streamSubjects = "PAYMENT.*"
	pubSubjectName = "PAYMENT.created"
	subSubjectName = "PAYMENT.response"
)

var (
	JS nats.JetStreamContext
)

func init() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Fatalf("error jetstream connect: %s", err.Error())
	}

	JS, err = nc.JetStream()
	if err != nil {
		logrus.Fatalf("error jetstream create: %s", err.Error())
	}

	err = createStream(JS)
	if err != nil {
		logrus.Fatalf("error stream create: %s", err.Error())
	}
}

// createStream creates a stream by using JetStreamContext
func createStream(js nats.JetStreamContext) error {
	// Check if the ORDERS stream already exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
			//	Retention: nats.WorkQueuePolicy,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateOrder(input models.Payment, userId int, js nats.JetStreamContext) error {
	orderJSON, _ := json.Marshal(input)

	_, err := js.Publish(pubSubjectName, orderJSON)
	if err != nil {
		return err
	}
	logrus.Printf("Data for user with user_id:%d has been published\n", userId)

	return nil
}

type Response struct {
	Status string `json:"status"`
}

func ResponseOrder(js nats.JetStreamContext) Response {
	var resp Response
	// Create durable consumer monitor
	_, err := js.Subscribe(subSubjectName, func(msg *nats.Msg) {
		msg.Ack()

		err := json.Unmarshal(msg.Data, &resp)
		if err != nil {
			logrus.Errorf("Error response format from bank: %s", err.Error())
		}

		logrus.Printf("subscribes from subject:%s\n", msg.Subject)
		logrus.Printf("Status response:%s\n", resp.Status)
	}, nats.ManualAck())
	if err != nil {
		logrus.Printf("error with subscribes")
		return Response{
			Status: err.Error(),
		}
	}

	//	runtime.Goexit()
	return resp
}
