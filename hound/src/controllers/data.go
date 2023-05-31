package controllers

func HandleData(data []byte, originIp string) <-chan string {
	c := make(chan string)

	go func() {
		c <- handleData(data, originIp)
	}()

	return c
}

func handleData(data []byte, originIp string) string {
	return "WebSocket OK"
}
