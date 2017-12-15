package main

import "fmt"

type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	var personDB map[string] PersonInfo
	    personDB = make(map[string] PersonInfo)

	    personDB["12345"] = PersonInfo{"12345", "tom", "room203"}
	    personDB["1"] = PersonInfo{"1", "jack", "room101"}
	    personDB["1234"] = PersonInfo{"1", "simith", "room101"}

        delete(personDB, "234")  
	    person, ok := personDB["1234"]

	    if ok {
	    	fmt.Println("Found person", person.Name, " with ID 1234.")
	    } else {

            fmt.Println("Did not found person with ID 1234.")
	    }
}