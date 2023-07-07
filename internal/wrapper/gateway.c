#include <stdbool.h>

#include "sharedspice.h"

int sendChar_cgo(char *c, int i, void *ptr)
{
    int sendChar(char *, int, void *);
    return sendChar(c, i, ptr);
}

int sendStat_cgo(char *c, int b, void *ptr)
{
    int sendChar(char *, int, void *);
    return sendChar(c, b, ptr);
}

int controlledExit_cgo(int i1, bool b1, bool b2, int i2, void *ptr)
{
    int controlledExit(int, bool, bool, int, void *);
    return controlledExit(i1, b1, b2, i2, ptr);
}

int sendData_cgo(pvecvaluesall vecValAll, int i1, int i2, void *ptr)
{
    int sendData(pvecvaluesall, int, int, void *);
    return sendData(vecValAll, i1, i2, ptr);
}

int sendInitData_cgo(pvecinfoall vecInfoAll, int i, void *ptr)
{
    int sendInitData(pvecinfoall, int, void *);
    return sendInitData(vecInfoAll, i, ptr);
}

int bgThreadRunning_cgo(bool b, int i, void *ptr)
{
    int bgThreadRunning(bool, int, void *);
    return bgThreadRunning(b, i, ptr);
}