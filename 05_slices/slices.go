package main

import fmt "fmt"

func main(  ) {
  fmt.Println( "Slices Lesson" );

  // slice factory
  var slice_100 = make([]int, 100);
  fmt.Println( slice_100 );  
  fmt.Printf( "Capacity of the slice is %d\n", cap(slice_100) );

  // slice of an array
  ar1 := []int{ 1, 2, 3, 4, 5 };
  fmt.Println( ar1[0:len(ar1)] );
  
  
}


