package main
import "fmt"

type Car struct {
	Model string
	Manufacturer string
	BuildYear int
}
type Cars []*Car
var AllCars Cars = []*Car{
&Car{ "Explorer","Ford", 2008},
&Car{"3i", "BMW", 2011},
&Car{"C300", "Mercedes", 2009},
&Car{"x4", "BMW", 2008}}

var targetManufacturer = "BMW"

func main(){
	var selectedCars Cars

	for _,item := range AllCars{
		if item.Manufacturer == targetManufacturer {
			selectedCars = append(selectedCars,item)
		}
	}
	fmt.Printf("select %v cars of %v using imperative coding命令式编程 \n", len(selectedCars), targetManufacturer)
	functional()
}