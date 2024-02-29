#include "usage.h"
#include <stdio.h>

struct rusage ReadUsage(void) {
    struct rusage r;
    getrusage(RUSAGE_SELF, &r);
    return r;
}