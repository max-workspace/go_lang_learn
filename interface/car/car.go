package car

type Any interface{}
type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}
type Cars []*Car

func (this Cars) Process(f func(car *Car)) {
	for _, carItem := range this {
		f(carItem)
	}
}

func (this Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)

	this.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func (this Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, len(this))
	ix := 0
	this.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)

	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}
