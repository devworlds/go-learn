package interfaces

import (
	"errors"
	"log"
)

type Data struct{
	Msg string
}

type Consumer interface{
	DataConsumer() (string, error)
}

func (d *Data) DataConsumer() (string, error) {
	if d.Msg == ""{
		log.Println("DataConsumer: Failed")
		return "", errors.New("Failed to consume")
	}
	log.Println("DataConsumer: Success")
	return  d.Msg, nil
}