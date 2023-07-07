#include <stdint.h>

typedef uintptr_t (*init)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
uintptr_t init_ng(void *f, uintptr_t a0, uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6)
{
    return ((init)f)(a0, a1, a2, a3, a4, a5, a6);
}
typedef uintptr_t (*command)(uintptr_t);
uintptr_t command_ng(void *f, uintptr_t a0)
{
    return ((command)f)(a0);
}

typedef uintptr_t (*running)();
uintptr_t running_ng(void *f)
{
    return ((running)f)();
}

typedef uintptr_t (*getVecInfo)(uintptr_t);
uintptr_t getVecInfo_ng(void *f, uintptr_t a0)
{
    return ((getVecInfo)f)(a0);
}

typedef uintptr_t (*circ)(uintptr_t);
uintptr_t circ_ng(void *f, uintptr_t a0)
{
    return ((circ)f)(a0);
}

typedef uintptr_t (*curPlot)();
uintptr_t curPlot_ng(void *f)
{
    return ((curPlot)f)();
}

typedef uintptr_t (*allPlots)();
uintptr_t allPlots_ng(void *f)
{
    return ((allPlots)f)();
}