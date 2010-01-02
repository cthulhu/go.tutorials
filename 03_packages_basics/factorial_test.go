package factorial_test

import (
  "testing";
  "./factorial";
)

func TestCalc(t *testing.T){
  if factorial.Calc( 5 ) != 120 {
    t.Errorf("Error calculation factorial\nExpected 120 was %d\n", factorial.Calc( 5 ) )
  }
  if factorial.Calc( 0 ) != 1 {
    t.Errorf("Error calculation factorial\nExpected 1 was %d\n", factorial.Calc( 0 ) )
  }
  if factorial.Calc( 1 ) != 1 {
    t.Errorf("Error calculation factorial\nExpected 1 was %d\n", factorial.Calc( 1 ) )
  }
}

