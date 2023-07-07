#ifndef GATEWAYS_H
#define GATEWAYS_H

#include "sharedspice.h"

int sendChar_cgo(char *, int, void *);
int sendStat_cgo(char *, int, void *);
int controlledExit_cgo(int, bool, bool, int, void *);
int sendData_cgo(pvecvaluesall, int, int, void *);
int sendInitData_cgo(pvecinfoall, int, void *);
int bgThreadRunning_cgo(bool, int, void *);

#endif