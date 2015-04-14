package main

import (
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
)

type Hivemind struct {
	procs  []*Process
	procWg sync.WaitGroup
	done   chan bool
}

func NewHivemind() (h *Hivemind) {
	h = &Hivemind{}
	h.createProcesses()
	return
}

func (h *Hivemind) createProcesses() {
	color := 32
	port := config.PortBase

	for _, entry := range parseProcfile("Procfile") {
		h.procs = append(
			h.procs,
			NewProcess(
				entry.Name,
				strings.Replace(entry.Command, "$PORT", strconv.Itoa(port), -1),
				color,
			),
		)

		color++
		port += config.PortStep
	}

	return
}

func (h *Hivemind) interrupted() chan os.Signal {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	return c
}

func (h *Hivemind) detectExit() {
	for {
		var exit bool

		select {
		case <-h.done:
			exit = config.ExitTogether
		case <-h.interrupted():
			exit = true
		}

		if exit {
			break
		}
	}
}

func (h *Hivemind) Run() {
	h.done = make(chan bool, len(h.procs))

	for _, proc := range h.procs {
		proc.Run(&h.procWg, h.done)
	}

	go func() {
		h.detectExit()

		for _, proc := range h.procs {
			proc.Term()
		}
	}()

	h.procWg.Wait()
}
