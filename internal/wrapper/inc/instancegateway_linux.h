#ifndef INSTANCEGATEWAYS_H
#define INSTANCEGATEWAYS_H

#include <stdint.h>

uintptr_t init_ng(void *, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
uintptr_t command_ng(void *, uintptr_t);
uintptr_t running_ng(void *);
uintptr_t getVecInfo_ng(void *, uintptr_t);
uintptr_t circ_ng(void *, uintptr_t);
uintptr_t curPlot_ng(void *);
uintptr_t allPlots_ng(void *);

#endif