package yices2

const (
	NO_ERROR ErrorCodeT = iota
	/*
	 * Errors in type or term construction
	 */
	INVALID_TYPE
	INVALID_TERM
	INVALID_CONSTANT_INDEX
	INVALID_VAR_INDEX // Not used anymore
	INVALID_TUPLE_INDEX
	INVALID_RATIONAL_FORMAT
	INVALID_FLOAT_FORMAT
	INVALID_BVBIN_FORMAT
	INVALID_BVHEX_FORMAT
	INVALID_BITSHIFT
	INVALID_BVEXTRACT
	INVALID_BITEXTRACT // added 2014/02/17
	TOO_MANY_ARGUMENTS
	TOO_MANY_VARS
	MAX_BVSIZE_EXCEEDED
	DEGREE_OVERFLOW
	DIVISION_BY_ZERO
	POS_INT_REQUIRED
	NONNEG_INT_REQUIRED
	SCALAR_OR_UTYPE_REQUIRED
	FUNCTION_REQUIRED
	TUPLE_REQUIRED
	VARIABLE_REQUIRED
	ARITHTERM_REQUIRED
	BITVECTOR_REQUIRED
	SCALAR_TERM_REQUIRED
	WRONG_NUMBER_OF_ARGUMENTS
	TYPE_MISMATCH
	INCOMPATIBLE_TYPES
	DUPLICATE_VARIABLE
	INCOMPATIBLE_BVSIZES
	EMPTY_BITVECTOR
	ARITHCONSTANT_REQUIRED // added 2013/01/23
	INVALID_MACRO          // added 2013/03/31
	TOO_MANY_MACRO_PARAMS  // added 2013/03/31
	TYPE_VAR_REQUIRED      // added 2013/03/31
	DUPLICATE_TYPE_VAR     // added 2013/03/31
	BVTYPE_REQUIRED        // added 2013/05/27
	BAD_TERM_DECREF        // added 2013/10/03
	BAD_TYPE_DECREF        // added 2013/10/03
	INVALID_TYPE_OP        // added 2014/12/03
	INVALID_TERM_OP        // added 2014/12/04
)

const (
	/*
	 * Parser errors
	 */
	INVALID_TOKEN ErrorCodeT = 100 + iota
	SYNTAX_ERROR
	UNDEFINED_TYPE_NAME
	UNDEFINED_TERM_NAME
	REDEFINED_TYPE_NAME
	REDEFINED_TERM_NAME
	DUPLICATE_NAME_IN_SCALAR
	DUPLICATE_VAR_NAME
	INTEGER_OVERFLOW
	INTEGER_REQUIRED
	RATIONAL_REQUIRED
	SYMBOL_REQUIRED
	TYPE_REQUIRED
	NON_CONSTANT_DIVISOR
	NEGATIVE_BVSIZE
	INVALID_BVCONSTANT
	TYPE_MISMATCH_IN_DEF
	ARITH_ERROR
	BVARITH_ERROR
)

const (
	/*
	 * Errors in assertion processing.
	 * These codes mean that the context as configured
	 * cannot process the assertions.
	 */
	CTX_FREE_VAR_IN_FORMULA ErrorCodeT = 300 + iota
	CTX_LOGIC_NOT_SUPPORTED
	CTX_UF_NOT_SUPPORTED
	CTX_ARITH_NOT_SUPPORTED
	CTX_BV_NOT_SUPPORTED
	CTX_ARRAYS_NOT_SUPPORTED
	CTX_QUANTIFIERS_NOT_SUPPORTED
	CTX_LAMBDAS_NOT_SUPPORTED
	CTX_NONLINEAR_ARITH_NOT_SUPPORTED
	CTX_FORMULA_NOT_IDL
	CTX_FORMULA_NOT_RDL
	CTX_TOO_MANY_ARITH_VARS
	CTX_TOO_MANY_ARITH_ATOMS
	CTX_TOO_MANY_BV_VARS
	CTX_TOO_MANY_BV_ATOMS
	CTX_ARITH_SOLVER_EXCEPTION
	CTX_BV_SOLVER_EXCEPTION
	CTX_ARRAY_SOLVER_EXCEPTION
	CTX_SCALAR_NOT_SUPPORTED // added 2015/03/26
	CTX_TUPLE_NOT_SUPPORTED  // added 2015/03/26
	CTX_UTYPE_NOT_SUPPORTED  // added 2015/03/26
)

const (
	/*
	 * Error codes for other operations
	 */
	CTX_INVALID_OPERATION ErrorCodeT = 400 + iota
	CTX_OPERATION_NOT_SUPPORTED
)

const (
	/*
	 * Errors in context configurations and search parameter settings
	 */
	CTX_INVALID_CONFIG ErrorCodeT = 500 + iota
	CTX_UNKNOWN_PARAMETER
	CTX_INVALID_PARAMETER_VALUE
	CTX_UNKNOWN_LOGIC
)

const (
	/*
	 * Error codes for model queries
	 */
	EVAL_UNKNOWN_TERM ErrorCodeT = 600 + iota
	EVAL_FREEVAR_IN_TERM
	EVAL_QUANTIFIER
	EVAL_LAMBDA
	EVAL_OVERFLOW
	EVAL_FAILED
	EVAL_CONVERSION_FAILED
	EVAL_NO_IMPLICANT
	EVAL_NOT_SUPPORTED
)

const (
	/*
	 * Error codes for model construction
	 */
	MDL_UNINT_REQUIRED ErrorCodeT = 700 + iota
	MDL_CONSTANT_REQUIRED
	MDL_DUPLICATE_VAR
	MDL_FTYPE_NOT_ALLOWED
	MDL_CONSTRUCTION_FAILED
)

const (
	/*
	 * Error codes in DAG/node queries
	 */
	YVAL_INVALID_OP ErrorCodeT = 800 + iota
	YVAL_OVERFLOW
	YVAL_NOT_SUPPORTED
)

const (
	/*
	 * Error codes for model generalization
	 */
	MDL_GEN_TYPE_NOT_SUPPORTED ErrorCodeT = 900 + iota
	MDL_GEN_NONLINEAR
	MDL_GEN_FAILED
)

const (
	/*
	 * MCSAT error codes
	 */
	MCSAT_ERROR_UNSUPPORTED_THEORY ErrorCodeT = 1000 + iota
)

const (
	/*
	 * Input/output and system errors
	 */
	OUTPUT_ERROR ErrorCodeT = 9000 + iota
)

const (
	/*
	 * Catch-all code for anything else.
	 * This is a symptom that a bug has been found.
	 */
	INTERNAL_EXCEPTION ErrorCodeT = 9999 + iota
)
