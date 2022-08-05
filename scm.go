package libgoile

// #cgo pkg-config: guile-3.0
// #include<libguile.h>
import "C"

// CSCM is wrapper of SCM type of libguile. Guile represents all
// Scheme values with the single C type SCM.
type CSCM C.SCM
