package main

import (
	"os"
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"
	ui "github.com/gizak/termui/v3"
)

var logger *log.Logger

func setupLogs() {
	fileName := "app.log"

	// open log file
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func process(y []complex128) DataHandle {
	padded_size := 512
	padded := Pad(y, padded_size)

	w := Fft(padded,padded_size)

	sr, sim := Flatten(padded)
	fr, fim := Flatten(w)

	handle := &DataHandle {
		sr, sim, fr, fim }
	
	return *handle
}

type signal struct {
	Re []float64 `json:"Re"`
	Im []float64 `json:"Im"`
}

func fetchSignal() []complex128 {
	requestURL := fmt.Sprintf("http://localhost:%d/signals/1", SERVER_PORT)

	res, err := http.Get(requestURL)
	defer res.Body.Close()

	if err != nil {
		log.Println("Error from api", err)
	}

	log.Println("Received signal, ", res.StatusCode)

	var result signal

	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &result) ; err != nil {
		log.Println("Could not parse json from api.")
	}

	return Zip(result.Re, result.Im)
}

func main() {

	setupLogs()
			
	log.Println("Booting...")
	go RunApi(log.Default())
	
	sig := fetchSignal()

	handle := process(sig)

	app := StartApp(handle)

	Loop(func(tickerCount int) {

		y := fetchSignal()
		handle := process(y)

		app.SignalPlot.Data[0] = handle.signal_real
		app.SignalPlot.Data[1] = handle.signal_imag
		app.FrequencyPlot.Data[0] = handle.frequency_real
		app.FrequencyPlot.Data[1] = handle.frequency_imag
		ui.Render(app.Grid)
	})

}

