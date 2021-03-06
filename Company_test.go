package ras_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	ras "github.com/rafaeltokyo/rasempresa-sdk-go"
)

func TestGetInfo(t *testing.T) {
	godotenv.Load(".env.test")
	client := ras.New(os.Getenv("KEY"), os.Getenv("SECRET"), os.Getenv("ENV"))
	response, errAPI, err := client.CompanyService().GetInfo()
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}
