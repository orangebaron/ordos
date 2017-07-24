package main

func serverFunc(quitChan chan struct{}) {
	for {
		select {
		case <-quitChan:
			return
		default:
			//do things
		}
	}
}
