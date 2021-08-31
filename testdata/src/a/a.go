package a

// func decls
func _(_, _, _, _, _ int)    {}         // OK
func _(_, _, _, _, _, _ int) {}         // want "too many params"
func _() (_, _, _ int)       { return } // OK
func _() (_, _, _, _ int)    { return } // want "too many results"

// func literals
func _() {
	_ = func(_, _, _, _, _ int) {}         // OK
	_ = func(_, _, _, _, _, _ int) {}      // want "too many params"
	_ = func() (_, _, _ int) { return }    // OK
	_ = func() (_, _, _, _ int) { return } // want "too many results"
}
