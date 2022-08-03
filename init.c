#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <libguile.h>

void* GoScmWithGuileFunc(uint64_t ctxid);

void* scm_with_guile_func (void* data) {
  uint64_t ctxid = *((uint64_t*) data);
  return GoScmWithGuileFunc(ctxid);
}

void* goile_scm_with_guile (void* data) {
  return scm_with_guile(scm_with_guile_func, data);
}
