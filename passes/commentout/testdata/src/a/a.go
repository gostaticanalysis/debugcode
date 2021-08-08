package a

//func f() { // want "do not leave a comment outed debug code with out reason"
//	panic("not implement")
//}

// with reason // OK
//func f() {
//	panic("not implement")
//}

// with reason
//
//func f() {
//	panic("not implement")
//}

var _ = struct{}{}

// no reason

//func f() { // want "do not leave a comment outed debug code with out reason"
//	panic("not implement")
//}

func f() {
	//var _ = struct{}{} // OK - in func

	//func() { // OK - in func
	// panic("not implement")
	//}()

	var _ = struct{}{}

	// with reason
	//var _ = struct{}{}
}
