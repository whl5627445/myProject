package Init

func init() {
	go Register()
	go OMCMessagesInit()
	go checkOMC()
	go ModelCodeAutoSave()
	go GcCollectAndUnmap()
}
