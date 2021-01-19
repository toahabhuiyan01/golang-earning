package main 
  
import (
	"fmt"
	"encoding/json"
)
  
// Main function 
func main() { 
  
    // Creating and initializing 
    // the anonymous structure 
    Element := struct { 
        name      string 
        branch    string 
        language  string 
        Particles int
    }{ 
        name:      "Pikachu", 
        branch:    "ECE", 
        language:  "C++", 
        Particles: 498, 
    } 
  
    // Display the anonymous structure 
	fmt.Println(json.Marshal(Element))
	
	mp := map[string]int{"python": 5, "java": 4}
	_a, out := json.unmarshal(mp)
	fmt.Println(out)
}