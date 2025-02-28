package filter

type Person struct {
	Name string
	Age  int
}

// FilterMinAge erwartet eine Liste von Personen und ein Mindestalter.
// FilterMinAge liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// die mindestens das angegebene Alter haben.
func FilterMinAge(list []Person, minAge int) []Person {

	// rekursionsanker
	if len(list) == 0 {
		return nil
	}

	// wenn alter ausreichend, element verwenden und fortschreiten
	if list[0].Age >= minAge {

		return append([]Person{list[0]}, FilterMinAge(list[1:], minAge)...)

	}

	// wenn nicht, element überspringen
	return append([]Person{}, FilterMinAge(list[1:], minAge)...)

}

// FilterLongNames erwartet eine Liste von Personen und eine Mindestlänge.
// FilterLongNames liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// deren Name mindestens die angegebene Länge hat.
func FilterLongNames(list []Person, minLength int) []Person {
	// TODO
	return []Person{}
}

// FilterNamePrefix erwartet eine Liste von Personen und einen Namenspräfix.
// FilterNamePrefix liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// deren Name mit dem angegebenen Präfix beginnt.
func FilterNamePrefix(list []Person, prefix string) []Person {
	// TODO
	return []Person{}
}

// FilterChildren erwartet eine Liste von Personen.
// FilterChildren liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// die höchstens 13 Jahre alt sind.
func FilterChildren(list []Person) []Person {
	// TODO
	return []Person{}
}

// FilterChildrenWithLongNames erwartet eine Liste von Personen und eine Mindestlänge.
// FilterChildrenWithLongNames liefert eine neue Liste, die Personen aus der ursprünglichen Liste enthält,
// die höchstens 13 Jahre alt sind und deren Name mindestens die angegebene Länge hat.
func FilterChildrenWithLongNames(list []Person, minLength int) []Person {
	// TODO
	return []Person{}
}
