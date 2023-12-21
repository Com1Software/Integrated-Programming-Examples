
#include "hbapi.h"
HB_FUNC(UDF) {
   unsigned int a,b,c;
   a=16;
   b=5;
   c=a+b;
   hb_retni(c);
}