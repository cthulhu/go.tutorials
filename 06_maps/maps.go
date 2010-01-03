
package main

import fmt "fmt"

func main() {
  fmt.Println( "Maps Lesson" );
  // creation with litterals
  m := map[string] int {"one":1, "two":2};
  fmt.Printf( "Map size %d\n", len(m) );
  // inserting element
  m["three"] = 3;
  fmt.Println( m );
  fmt.Printf( "Map size %d\n", len(m) );

  // creation with factory
  var m2 = make(map[string] int);
  fmt.Println( m2 );  
  m2["one"] = 1;
  m2["two"] = 2;
  fmt.Println( m2 );  
  
  // assignment is a coping
  
  var m3 = m2;
  fmt.Println( "m2=", m2 );  
  fmt.Println( "m3=", m3 );  
  m2["four"] = 4;
  fmt.Println( "m2=", m2 );  
  fmt.Println( "m3=", m3 );    
  m3["five"] = 5;
  fmt.Println( "m2=", m2 );  
  fmt.Println( "m3=", m3 );    
  
  // trick for testing presence
  var value int;
  var presence bool;
  m3["five"] = 5;  

  value, presence = m3["five"];
  fmt.Println( "Value =", value, " Presence =", presence );    

  value, presence = m3["six"];
  fmt.Println( "Value =", value, " Presence =", presence );    
  
  // deletion element
  m3["five"] = 5, false;    
  value, presence = m3["five"];
  fmt.Println( "Value =", value, " Presence =", presence );    
  
  // loop over hash with for
  // key value
  for key, value := range m3 {
    fmt.Printf("key %v, value %v\n", key, value)
  } 
  // keys only
  for key := range m3 {
    fmt.Printf( "key %v\n", key )
  } 

  
}


