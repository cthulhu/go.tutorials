
package main

import fmt "fmt" 

// implicit return value 
func concat_strings( message1 string, message2 string ) ( result string ) {
  result = message1 + message2;
  return
}

// explicit return value 
func string_length( message string ) ( result int ) {
  result = len(message);
  return result
}

// few return values
func concat_strings_and_get_lengths( message1 string, message2 string ) ( result_string string, result_length int ) {
  result_string = message1 + message2;
  result_length = len( result_string );
  return result_string, result_length
}

// defer
func string_length_with_output( message string ) int {
  defer fmt.Printf( "String %s is %d symbols\n", message, len( message ) );
  fmt.Printf( "Just before defer\n" );
  return len( message );
}

// closures

func string_length_calculator() ( func (string) int ) {
  return func ( message string ) int {
    return len( message );
  }
}

func main() {
    fmt.Printf( concat_strings("Hello ", "world!\n") );
    fmt.Printf( "%d\n", string_length("Hello") );
    var result_string string;
    var result_length int;
    result_string, result_length = concat_strings_and_get_lengths("Hello ", "world!");
    fmt.Printf( "String %s is %d symbols length \n", result_string, result_length );
    fmt.Printf( "%d\n", string_length_with_output( "Hello and defer should be" ) );
    var length_calculator = string_length_calculator();
    fmt.Printf( "%d\n", length_calculator("Hello") );        
}


