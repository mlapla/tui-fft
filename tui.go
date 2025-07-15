package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"time"
	"log"
)

type DataHandle struct {
	signal_real []float64
	signal_imag []float64
	frequency_real []float64
	frequency_imag []float64
}

type AppRoot struct {
	SignalPlot *widgets.Plot
	FrequencyPlot *widgets.Plot
	Grid *ui.Grid
}

func newApp(data DataHandle) AppRoot {

	sig_plt := widgets.NewPlot()
	sig_plt.Title = "Signal"
	sig_plt.Data = append(sig_plt.Data, data.signal_real)
	sig_plt.Data = append(sig_plt.Data, data.signal_imag)
	sig_plt.AxesColor = ui.ColorWhite
	sig_plt.LineColors[0] = ui.ColorYellow
	sig_plt.MaxVal = 1.9

	freq_plt := widgets.NewPlot()
	freq_plt.Title = "Frequency"
	freq_plt.Data = append(freq_plt.Data, data.frequency_real)
	freq_plt.Data = append(freq_plt.Data, data.frequency_imag)
	freq_plt.AxesColor = ui.ColorGreen
	freq_plt.LineColors[0] = ui.ColorYellow
	freq_plt.MaxVal = 150

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0, sig_plt),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0, freq_plt),
		),
	)

	app := &AppRoot {
		sig_plt, freq_plt, grid}

	return *app
}

var tickerCount int = 1

func StartApp(data DataHandle) AppRoot {
	
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	app := newApp(data)
	ui.Render(app.Grid)

	log.Println("New app created...")
	return app
}

func Loop(updateData func(tickerCount int)) {
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		tickerCount++
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			updateData(tickerCount)
			log.Println("tick")
			tickerCount++
		}
	}
}
