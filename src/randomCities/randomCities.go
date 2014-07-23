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

//	Gets the words from the file and return them in a map[uint8][]string.
func getCitiesFromFile(filename string)  map[uint8][]string{
	
    // Create a map to save the cities:
    cities := make(map[uint8][]string)

    // Open the file to read:
    bs, err := ioutil.ReadFile(filename)
    
    if err != nil {
        return cities
    }

    str := string(bs)

    // Creates a scanner to read the file:
    scanner := bufio.NewScanner(strings.NewReader(str))

    scanner.Split(bufio.ScanLines)

    // Read file line by line:
  	for scanner.Scan() {
  	    word := scanner.Text()
  	   	cities[word[0]] = append(cities[word[0]], word)
  	}

    return cities
}

// Picks up a random city and return it.
func pickUpRandomCity(cities map[uint8][]string, k uint8) string{

	rand.Seed(time.Now().UTC().UnixNano())
	randomNumber := rand.Intn(len(cities[k]))

	return cities[k][randomNumber]
}

// Returns the random cities.
func RandomCities(inputFileName string) map[string]string{

    // Get cities from file sorted like this: { A: Abell, Avila...
    //                                          B: Barcelona, Bilbao... }
    cities := getCitiesFromFile(inputFileName)

    randomCities := make(map[string]string)

    // Get random cities and save in the file:
   	for k, _ := range cities {
       randomCities[string(k)] = pickUpRandomCity(cities, k)
   	}  

    return randomCities  
}