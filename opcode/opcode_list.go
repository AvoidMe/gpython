package opcode

const (
	POP_TOP             = 1
	NOP                 = 9
	BINARY_ADD          = 23
	BINARY_SUBSTRACT    = 24
	BINARY_SUBSCR       = 25
	STORE_SUBSCR        = 60
	RETURN_VALUE        = 83
	STORE_NAME          = 90
	LOAD_CONST          = 100
	LOAD_NAME           = 101
	BUILD_LIST          = 103
	BUILD_MAP           = 105
	COMPARE_OP          = 107
	IS_OP               = 117
	CALL_FUNCTION       = 131
	BUILD_CONST_KEY_MAP = 156
	LIST_EXTEND         = 162
)
