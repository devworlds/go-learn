package interfaces

import (
	"testing"
)

func TestDataConsumerSuccess(t *testing.T){
	data := Data{
		Msg: "data",
	}
	var consumer Consumer = &data
	res, err := consumer.DataConsumer()
	if res != "data"{
		panic(err)
	}
}

func TestDataConsumerFail(t *testing.T){
	data := Data{
		Msg: "",
	}
	var consumer Consumer = &data
	res, err := consumer.DataConsumer()
	if res != ""{
		panic(err)
	}
}