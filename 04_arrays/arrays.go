
package main

import fmt "fmt"
import sort "sort"  

// U need to remember that arrays are values not pointers like in C
func try_to_transform_array( a [5] int ){
  a[0] = 1;
}

// So u need to put pointer to specify C like behaviour 
func transform_array( a *[5] int ){
  a[0] = 1;
}

func array_reversing( ar *[ 5 ] int ) (*[5] int){
  var tmp int;
  for i := 0 ; i < len( ar ) / 2 ; i++ {
    tmp = ar[ i ]; 
    ar[ i ] = ar[ len( ar ) - 1 - i ];
    ar[ len( ar ) - 1 - i ] = tmp;
  }
  return ar;
}

func main() {
  var ar [5] int;
  fmt.Println( ar ); // nice trick to output all elements

  // arrays and pointers
  try_to_transform_array( ar );
  fmt.Println( ar ); 
  transform_array( &ar );
  fmt.Println( ar );

  // Arrays with initializing
  // first 3 elements
  ar2 := [5] int { 1, 2, 3 };
  fmt.Println( ar2 );
  // random elements
  ar3 := [5] int { 0:1, 2:3, 4:5 };
  fmt.Println( ar3 );
  
  // without size specifying
  ar4 := [...] int { 1, 2, 3, 4, 5 };
  fmt.Println( ar4 );
  
  // without size specifying and just 1 ellement
  ar5 := [...] int { 4:5 };
  fmt.Println( ar5 );
 
  // passing array litteral as argument to function
  ar6 := array_reversing( &[ 5 ]int{ 1, 2, 3, 4, 5 } );
  fmt.Println( *ar6 );

  // array functions
  // len - size
  ar7 := []int{ 1, 3, 2, 5, 1 };
  fmt.Printf( "Array size is %d\n", len( ar7 ) );

  // sorting
  sort.Ints( ar7 );
  fmt.Println( ar7 );
}


