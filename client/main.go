package main

func main() {

	go doSyncFromLocalEventStore()
	go doPullFileFormServer()
	go doPushFileToServer()

}
