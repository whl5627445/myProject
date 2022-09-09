package Init

func init() {
	ModelLibraryInit()
	LogInit()
	go SimulationService()
	go NacosRegister()
}
