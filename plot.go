package ouwsdk


type Plot struct {
	Area								int					`redis:"area"`
	Owner								string			`redis:"owner"`
	Farmable						int					`redis:"farmable"`
	Water								int					`redis:"water"`
	Elevation						int					`redis:"elevation"`
	Resources						map[string] Resource			`redis:"resources"`
	Inhabitants					map[string] User					`redis:"inhabitants"`
}


func addPlot() {

	
} // addPlot
