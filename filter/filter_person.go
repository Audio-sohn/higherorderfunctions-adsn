package filter

import "strings"

type Person struct {
	Name string
	Age  int
}

// FilterMinAge erwartet eine Liste von Personen und ein Mindestalter.
// FilterMinAge liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// die mindestens das angegebene Alter haben.
func FilterMinAge(list []Person, minAge int) []Person {

	predicate := func(x Person) bool {
		return x.Age >= minAge
	}

	// wenn alter ausreichend, element verwenden und fortschreiten
	return FilterList(list, predicate)

}

// FilterLongNames erwartet eine Liste von Personen und eine Mindestlänge.
// FilterLongNames liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// deren Name mindestens die angegebene Länge hat.
func FilterLongNames(list []Person, minLength int) []Person {
	// rekursionsanker

	predicate := func(x Person) bool {

		return len(x.Name) >= minLength

	}

	return FilterList(list, predicate)

}

// FilterNamePrefix erwartet eine Liste von Personen und einen Namenspräfix.
// FilterNamePrefix liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// deren Name mit dem angegebenen Präfix beginnt.
func FilterNamePrefix(list []Person, prefix string) []Person {
	// rekursionsanker

	predicate := func(x Person) bool {

		return strings.HasPrefix(x.Name, prefix)

	}

	return FilterList(list, predicate)

}

// FilterChildren erwartet eine Liste von Personen.
// FilterChildren liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// die höchstens 13 Jahre alt sind.
func FilterChildren(list []Person) []Person {

	minAge := 13

	predicate := func(x Person) bool {

		return x.Age <= minAge

	}

	return FilterList(list, predicate)
}

// FilterChildrenWithLongNames erwartet eine Liste von Personen und eine Mindestlänge.
// FilterChildrenWithLongNames liefert eine neue Liste, die Personen aus der ursprünglichen Liste enthält,
// die höchstens 13 Jahre alt sind und deren Name mindestens die angegebene Länge hat.
func FilterChildrenWithLongNames(list []Person, minLength int) []Person {

	// wenn nicht, element überspringen
	return FilterLongNames(FilterChildren(list), minLength)
}
