package phoenix

// #include "go-clang.h"
import "C"

// A group of CXDiagnostics.
type DiagnosticSet struct {
	c C.CXDiagnosticSet
}

// Determine the number of diagnostics in a CXDiagnosticSet.
func (ds DiagnosticSet) NumDiagnosticsInSet() uint16 {
	return uint16(C.clang_getNumDiagnosticsInSet(ds.c))
}

// Retrieve a diagnostic associated with the given CXDiagnosticSet. \param Diags the CXDiagnosticSet to query. \param Index the zero-based diagnostic number to retrieve. \returns the requested diagnostic. This diagnostic must be freed via a call to \c clang_disposeDiagnostic().
func (ds DiagnosticSet) DiagnosticInSet(Index uint16) Diagnostic {
	return Diagnostic{C.clang_getDiagnosticInSet(ds.c, C.uint(Index))}
}

// Release a CXDiagnosticSet and all of its contained diagnostics.
func (ds DiagnosticSet) Dispose() {
	C.clang_disposeDiagnosticSet(ds.c)
}
