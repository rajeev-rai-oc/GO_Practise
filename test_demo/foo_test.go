package main
import "testing"
func TestFooer(t *testing.T) {
    result := Fooer(4)
    if result != "Three" {
    t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Three")
    }
}

func TestFooerTableDriven(t *testing.T) {
	// Defining the columns of the table
	  var tests = []struct {
	  name string
		  input int
		  want  string
	  }{
		  // the table itself
		  {"9 should be Three", 9, "Three"},
		  {"3 should be Three", 3, "Three"},
		  {"1 is not Three", 1, "1"},
		  {"0 should be Three", 0, "Three"},
	  }
	// The execution loop
	  for _, tt := range tests {
		  t.Run(tt.name, func(t *testing.T) {
			  ans := Fooer(tt.input)
			  if ans != tt.want {
				  t.Errorf("got %s, want %s", ans, tt.want)
			  }
		  })
	  }
  }