package factorial

// Functions those start from uppercase are pulic and would be imported 
// to access from other packages
func Calc( n int ) (rv int) {
  if n == 0 {
    rv = 1
  } else {
    rv = n * Calc( n - 1 )
  }
  return rv
}
