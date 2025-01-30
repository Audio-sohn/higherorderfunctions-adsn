package filter

type Person struct {
	Name string
	Age  int
}

// FilterMinAge erwartet eine Liste von Personen und ein Mindestalter.
// FilterMinAge liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// die mindestens das angegebene Alter haben.
func FilterMinAge(list []Person, minAge int) []Person {
	// HINWEIS: Verwenden Sie die Funktion FilterList.
	// Übergeben Sie ihr eine Funktion, die für eine Person p
	// überprüft, ob sie mindestens das Alter minAge hat.
	// TODO
	return []Person{}
}

// FilterLongNames erwartet eine Liste von Personen und eine Mindestlänge.
// FilterLongNames liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// deren Name mindestens die angegebene Länge hat.
func FilterLongNames(list []Person, minLength int) []Person {
	// HINWEIS: Verwenden Sie die Funktion FilterList.
	// Übergeben Sie ihr eine Funktion, die für eine Person p
	// überprüft, ob ihr Name wenigstens minLength Zeichen hat.
	// TODO
	return []Person{}
}

// FilterNamePrefix erwartet eine Liste von Personen und einen Namenspräfix.
// FilterNamePrefix liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// deren Name mit dem angegebenen Präfix beginnt.
func FilterNamePrefix(list []Person, prefix string) []Person {
	// HINWEIS: Verwenden Sie die Funktion FilterList.
	// Übergeben Sie ihr eine Funktion, die für eine Person p
	// überprüft, ob ihr Name mit prefix beginnt.
	// TODO
	return []Person{}
}

// FilterChildren erwartet eine Liste von Personen.
// FilterChildren liefert eine neue Liste, die nur die Personen der ursprünglichen Liste enthält,
// die höchstens 13 Jahre alt sind.
func FilterChildren(list []Person) []Person {
	// HINWEIS: Verwenden Sie die Funktion FilterList.
	// Übergeben Sie ihr eine Funktion, die für eine Person p
	// überprüft, ob ihr Alter höchstens 13 ist.
	// TODO
	return []Person{}
}

// FilterChildrenWithLongNames erwartet eine Liste von Personen und eine Mindestlänge.
// FilterChildrenWithLongNames liefert eine neue Liste, die Personen aus der ursprünglichen Liste enthält,
// die höchstens 13 Jahre alt sind und deren Name mindestens die angegebene Länge hat.
func FilterChildrenWithLongNames(list []Person, minLength int) []Person {
	// HINWEIS: Kombinieren Sie geeignete Funktionen, die Sie bereits geschrieben haben.
	// TODO
	return []Person{}
}
