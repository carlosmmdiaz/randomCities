// Package that implementa utility to read a list of words from a file
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
  list map[uint8][]string
}

// Struct constructor:
func (cities *Cities) New() {
  cities.list = make(map[uint8][]string)
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
          cities.list[word[0]] = append(cities.list[word[0]], word)
      }
    } 
}

// Picks up a random city and return it.
func (cities *Cities) pickUpRandomCity(k uint8) string{

	rand.Seed(time.Now().UTC().UnixNano())
	randomNumber := rand.Intn(len(cities.list[k]))

	return cities.list[k][randomNumber]
}

// Returns the random cities.
func (cities *Cities) RandomCities(inputFileName string) map[string]string{

    // Get cities from file sorted like this: { A: Abell, Avila...
    //                                          B: Barcelona, Bilbao... }
    cities.getCitiesFromFile(inputFileName)

    randomCities := make(map[string]string)

    // Get random cities and save in the file:
   	for k, _ := range cities.list {
       randomCities[string(k)] = cities.pickUpRandomCity(k)
   	}  

    return randomCities  
}