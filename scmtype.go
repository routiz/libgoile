package libgoile

// #include<libguile.h>
import "C"

type Scm struct {
	p C.SCM
}
