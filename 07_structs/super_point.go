package super_point

// type for export should start with Uppercase
type SuperPoint struct { 
  X, Y int;    // only those fields export (Uppercase)
  name string; // this doesn't
}

// thats a trick to avoid  troubles with linker
// http://code.google.com/p/go/issues/detail?id=87&colspec=ID%20Type%20Status%20Owner%20Reporter%20Summary
func FixerOfIssue87() {
  
}


