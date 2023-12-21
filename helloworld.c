/*
 * Harbour 3.2.0dev (r2312082217)
 * GNU C 12.2 (64-bit)
 * Generated C source from "helloworld.prg"
 */

#include "hbvmpub.h"
#include "hbinit.h"


HB_FUNC( HELLOWORLD );
HB_FUNC( MAIN );
HB_FUNC_EXTERN( QOUT );


HB_INIT_SYMBOLS_BEGIN( hb_vm_SymbolInit_HELLOWORLD )
{ "HELLOWORLD", {HB_FS_PUBLIC | HB_FS_FIRST | HB_FS_LOCAL}, {HB_FUNCNAME( HELLOWORLD )}, NULL },
{ "MAIN", {HB_FS_PUBLIC | HB_FS_LOCAL}, {HB_FUNCNAME( MAIN )}, NULL },
{ "QOUT", {HB_FS_PUBLIC}, {HB_FUNCNAME( QOUT )}, NULL }
HB_INIT_SYMBOLS_EX_END( hb_vm_SymbolInit_HELLOWORLD, "helloworld.prg", 0x0, 0x0003 )

#if defined( HB_PRAGMA_STARTUP )
   #pragma startup hb_vm_SymbolInit_HELLOWORLD
#elif defined( HB_DATASEG_STARTUP )
   #define HB_DATASEG_BODY    HB_DATASEG_FUNC( hb_vm_SymbolInit_HELLOWORLD )
   #include "hbiniseg.h"
#endif

HB_FUNC( HELLOWORLD )
{
	static const HB_BYTE pcode[] =
	{
		7
	};

	hb_vmExecute( pcode, symbols );
}

HB_FUNC( MAIN )
{
	static const HB_BYTE pcode[] =
	{
		36,5,0,176,2,0,106,14,72,101,108,108,111,44,
		32,119,111,114,108,100,33,0,20,1,36,7,0,7
	};

	hb_vmExecute( pcode, symbols );
}

