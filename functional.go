package main
import "fmt"

func (cs Cars) process (action func( *Car)){
	for _,c := range cs{
		action(c)
	}
}
func (cs Cars) findAll (selector func(* Car) bool) Cars{
	var selectedCars Cars
	cs.process(func(c *Car){
		if  selector(c){
			selectedCars= append(selectedCars,c)
		}
	})
	return selectedCars
}
func functional(){
	selectedCars := AllCars.findAll(func(c* Car) bool{return c.Manufacturer == targetManufacturer})
	fmt.Printf("select %v cars of %v using functional coding函数式编程 \n ", len(selectedCars), targetManufacturer)
}