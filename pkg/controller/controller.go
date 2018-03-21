package controller

var (
	indexHandler index
)
func Startup(){
	indexHandler.registerHandlers()
}