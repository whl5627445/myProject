package Init

func init() {
	ModelLibraryInit()
	LogInit()
	go simulationService()
	go NacosRegister()
	go OMCMessagesInit()
	go checkOMC()
}
