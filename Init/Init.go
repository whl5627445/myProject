package Init

func init() {
	go Register()
	go OMCMessagesInit()
	go ListenTaskDispatcher()
	go initWorkSpace()
	go checkWorkSpace()
}
