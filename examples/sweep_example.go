package main

import (
	"fmt"
	blockIO "github.com/BlockIo/block_io-go"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	toAddr := os.Getenv("TO_ADDRESS")
	privKey := os.Getenv("PRIVATE_KEY_FROM_ADDRESS")
	fromAddr := os.Getenv("FROM_ADDRESS")

	restClient := resty.New()
	rawSweepResponse, err := restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(map[string]interface{}{
			"to_address":   toAddr,
			"public_key":   blockIO.PubKeyFromWIF(privKey),
			"from_address": fromAddr,
		}).Post("https://block.io/api/v2/sweep_from_address?api_key=" + apiKey)

	if err != nil {
		panic(err)
	}

	fmt.Println("Raw sweep response: ")
	fmt.Println(rawSweepResponse)

	sweepData, sweepDataErr := blockIO.ParseResponseData(rawSweepResponse)

	if sweepDataErr != nil {
		panic(sweepDataErr)
	}

	signatureReq, signSweepReqErr := blockIO.SignSweepRequest(privKey, sweepData)

	if signSweepReqErr != nil {
		panic(signSweepReqErr)
	}

	signAndFinalizeRes, err := restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(map[string]interface{}{
			"signature_data": signatureReq,
		}).Post("https://block.io/api/v2/sign_and_finalize_sweep?api_key=" + apiKey)

	fmt.Println(signAndFinalizeRes)
}