// Package that implementa utility to read a List of words from a file
// and return a list of random words.
//
// input: cities_US.txt like (Abbeville Abbotsford Abbott Abbottsburg Abbottstown Abbyville Abell Abercrombie Aberdeen)
// output: map[string]string like { A: Abell
//                                  B: Bronwood }

package RandomCities

import (
        "io/ioutil"
        "strings"
        "bufio"
        "math/rand"
        "time"
)

// Struct to save and manage the cities from the file:
type Cities struct {
  List map[uint8][]string
  RandomCities map[uint8]string
}

// Struct constructor:
func (cities *Cities) New() {
  
  // Make the map:
  cities.List = make(map[uint8][]string)
  
  // Create the final map:
  cities.RandomCities = make(map[uint8]string)
}

//	Gets the words from the file and return them in a map[uint8][]string.
func (cities *Cities) getCitiesFromFile(filename string) {

    // Open the file to read:
    bs, err := ioutil.ReadFile(filename)
    
    if err == nil {
      
      str := string(bs)

      // Creates a scanner to read the file:
      scanner := bufio.NewScanner(strings.NewReader(str))

      scanner.Split(bufio.ScanLines)

      // Read file line by line:
      for scanner.Scan() {
          word := scanner.Text()
          cities.List[word[0]] = append(cities.List[word[0]], word)
      }
    } 
}

// Picks up a random city and return it.
func (cities *Cities) pickUpRandomCity(k uint8, c chan string) {

	rand.Seed(time.Now().UTC().UnixNano())
	randomNumber := rand.Intn(len(cities.List[k]))

	chosenCity := string(k) + " - " + cities.List[k][randomNumber]
  c <- chosenCity
}

// Returns the random cities:
func (cities *Cities) GetRandomCities(inputFileName string) {

    // Get cities from file sorted like this: { A: Abell, Avila...
    //                                          B: Barcelona, Bilbao... }
    cities.getCitiesFromFile(inputFileName)

    c := make(chan string)

    // Get random cities and save in the file:
   	for k, _ := range cities.List {
       go cities.pickUpRandomCity(k, c)
   	}

    for i := 0; i < len(cities.List); i++ {
      cities.RandomCities[uint8(i)] = <- c
    }


}

