package Init

func init() {
	ModelLibraryInit()
	LogInit()
	go simulationService()
	go Register()
	go OMCMessagesInit()
	go checkOMC()
	go ModelCodeAutoSave()
	go GcCollectAndUnmap()
}
