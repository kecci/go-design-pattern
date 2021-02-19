package main

import "github.com/kecci/design-pattern-go/adapter/computer"

func main() {

	client := &computer.Client{}
	mac := &computer.Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &computer.Windows{}
	windowsMachineAdapter := &computer.WindowsAdapter{WindowMachine: windowsMachine}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
