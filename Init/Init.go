package Init

func init() {
	modelLibraryInit()
	LogInit()
	go SimulationService()
	go NacosRegister()
	go OMCMessagesInit()
}
