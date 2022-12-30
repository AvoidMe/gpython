package opcode

const (
	POP_TOP          = 1
	PUSH_NULL        = 2
	NOP              = 9
	BINARY_ADD       = 23
	BINARY_SUBSTRACT = 24
	BINARY_SUBSCR    = 25
	STORE_SUBSCR     = 60
	RETURN_VALUE     = 83
	STORE_NAME       = 90
	LOAD_CONST       = 100
	LOAD_NAME        = 101
	// BUILD_TUPLE is never productd opcode (I think)
	// BUILD_TUPLE         = 102
	BUILD_LIST          = 103
	BUILD_MAP           = 105
	COMPARE_OP          = 107
	IS_OP               = 117
	RESUME              = 151
	BUILD_CONST_KEY_MAP = 156
	LIST_EXTEND         = 162
	PRECALL             = 166
	CALL                = 171
)
