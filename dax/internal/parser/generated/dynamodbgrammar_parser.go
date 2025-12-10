// Code generated from DynamoDbGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package generated // DynamoDbGrammar

import (
	"fmt"
	"strconv"
	"sync"

	antlr "github.com/antlr4-go/antlr/v4"
)

// import com.amazon.dynamodb.grammar.exceptions.RedundantParenthesesException;

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DynamoDbGrammarParser struct {
	*antlr.BaseParser
}

var DynamoDbGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dynamodbgrammarParserInit() {
	staticData := &DynamoDbGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'('", "')'", "'.'", "'['", "']'", "'*'", "", "'='", "'<>'",
		"'<'", "'<='", "'>'", "'>='", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "WS", "EQ", "NE", "LT", "LE", "GT",
		"GE", "PLUS", "MINUS", "IN", "BETWEEN", "NOT", "AND", "OR", "SET", "ADD",
		"DELETE", "REMOVE", "INSERT", "INTO", "VALUE", "VALUES", "UPDATE", "ON",
		"RETURNING", "ALL", "KEYS", "KEY", "ONLY", "NEW", "NONE", "OLD", "UPDATED",
		"DUPLICATE", "CREATE", "TABLE", "PRIMARY", "NUMBER", "BINARY", "STRING",
		"IF", "CAPACITY", "GLOBAL", "LOCAL", "INDEXKEYWORD", "PROJECTION", "SHOW",
		"TABLES", "SELECT", "FROM", "USE", "ENABLE", "SCAN", "WHERE", "DROP",
		"ALTER", "DESCRIBE", "OPTION", "INDEX", "ID", "ATTRIBUTE_NAME_SUB",
		"LITERAL_SUB", "STRING_LITERAL", "UNKNOWN",
	}
	staticData.RuleNames = []string{
		"projection_", "projection", "condition_", "condition", "comparator_symbol",
		"update_", "update", "set_section", "set_action", "add_section", "add_action",
		"delete_section", "delete_action", "remove_section", "remove_action",
		"set_value", "arithmetic", "operand", "function", "path", "id", "dereference",
		"literal", "expression_attr_names_sub", "expression_attr_values_sub",
		"statement_", "statement", "dropTableStatement", "describeStatement",
		"alterTableStatement", "alterTableStatementType", "setCapacity", "addIndex",
		"dropIndex", "alterIndex", "updateStatement", "deleteStatement", "insertStatement",
		"createTableStatement", "showTablesStatement", "selectStatement", "selectProjection",
		"optionBlock", "option", "singleOption", "keyValueOption", "optionKey",
		"optionValue", "optionValueString", "optionValueNumber", "star", "hint",
		"indexHint", "indexHintName", "scanInfo", "totalSegment", "segment",
		"where", "primaryKeyDecl", "secondaryIndexDecl", "secondaryIndexType",
		"indexName", "projectionIndex", "projectionIndexKeysOnly", "projectionIndexVector",
		"capacity", "readUnits", "writeUnits", "indexDecl", "attributeDecl",
		"hashKey", "rangeKey", "attributeName", "attributeType", "returning",
		"returningValue", "onDuplicateKeyUpdate", "ifClause", "tableName", "ddlName",
		"string", "unknown",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 655, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2, 73, 7,
		73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78, 7, 78,
		2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 5, 1, 171, 8, 1, 10, 1, 12, 1, 174, 9, 1, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 190,
		8, 3, 10, 3, 12, 3, 193, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 211, 8, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 219, 8, 3, 10, 3, 12, 3, 222,
		9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 4, 6, 233,
		8, 6, 11, 6, 12, 6, 234, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 241, 8, 7, 10, 7,
		12, 7, 244, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9,
		254, 8, 9, 10, 9, 12, 9, 257, 9, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 5, 11, 266, 8, 11, 10, 11, 12, 11, 269, 9, 11, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 278, 8, 13, 10, 13, 12, 13,
		281, 9, 13, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15, 287, 8, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 298, 8, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 308, 8, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 315, 8, 18, 10, 18, 12, 18, 318,
		9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 324, 8, 19, 10, 19, 12, 19, 327,
		9, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 336, 8,
		21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3,
		26, 358, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 376, 8,
		30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3,
		35, 398, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 405, 8, 36, 1,
		37, 1, 37, 3, 37, 409, 8, 37, 1, 37, 1, 37, 3, 37, 413, 8, 37, 1, 37, 1,
		37, 3, 37, 417, 8, 37, 1, 37, 3, 37, 420, 8, 37, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 5, 38, 429, 8, 38, 10, 38, 12, 38, 432, 9, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 440, 8, 38, 10, 38, 12,
		38, 443, 9, 38, 3, 38, 445, 8, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 457, 8, 40, 1, 40, 3, 40, 460, 8,
		40, 1, 40, 3, 40, 463, 8, 40, 1, 41, 1, 41, 3, 41, 467, 8, 41, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 5, 42, 474, 8, 42, 10, 42, 12, 42, 477, 9, 42,
		1, 42, 1, 42, 1, 43, 1, 43, 3, 43, 483, 8, 43, 1, 44, 1, 44, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 3, 47, 495, 8, 47, 1, 48,
		1, 48, 1, 48, 3, 48, 500, 8, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1,
		51, 3, 51, 508, 8, 51, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53,
		3, 53, 517, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 3, 54, 527, 8, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57,
		1, 58, 1, 58, 1, 58, 1, 58, 3, 58, 540, 8, 58, 1, 59, 3, 59, 543, 8, 59,
		1, 59, 1, 59, 1, 59, 1, 59, 5, 59, 549, 8, 59, 10, 59, 12, 59, 552, 9,
		59, 3, 59, 554, 8, 59, 1, 59, 1, 59, 1, 59, 3, 59, 559, 8, 59, 1, 59, 3,
		59, 562, 8, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 3, 62, 570, 8,
		62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 5, 64,
		581, 8, 64, 10, 64, 12, 64, 584, 9, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1,
		65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 1, 68,
		1, 68, 1, 68, 3, 68, 603, 8, 68, 1, 68, 1, 68, 1, 69, 1, 69, 1, 69, 1,
		70, 1, 70, 1, 71, 1, 71, 1, 72, 1, 72, 1, 73, 1, 73, 1, 74, 1, 74, 1, 74,
		1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 1, 75, 3, 75, 630,
		8, 75, 1, 76, 1, 76, 1, 76, 1, 76, 1, 76, 3, 76, 637, 8, 76, 1, 77, 1,
		77, 1, 77, 1, 78, 1, 78, 1, 79, 1, 79, 3, 79, 646, 8, 79, 1, 80, 1, 80,
		1, 81, 4, 81, 651, 8, 81, 11, 81, 12, 81, 652, 1, 81, 0, 1, 6, 82, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
		78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110,
		112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140,
		142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 0, 6, 1, 0, 9, 14,
		1, 0, 15, 16, 1, 0, 67, 68, 1, 0, 28, 29, 1, 0, 50, 51, 1, 0, 45, 47, 645,
		0, 164, 1, 0, 0, 0, 2, 167, 1, 0, 0, 0, 4, 175, 1, 0, 0, 0, 6, 210, 1,
		0, 0, 0, 8, 223, 1, 0, 0, 0, 10, 225, 1, 0, 0, 0, 12, 232, 1, 0, 0, 0,
		14, 236, 1, 0, 0, 0, 16, 245, 1, 0, 0, 0, 18, 249, 1, 0, 0, 0, 20, 258,
		1, 0, 0, 0, 22, 261, 1, 0, 0, 0, 24, 270, 1, 0, 0, 0, 26, 273, 1, 0, 0,
		0, 28, 282, 1, 0, 0, 0, 30, 286, 1, 0, 0, 0, 32, 297, 1, 0, 0, 0, 34, 307,
		1, 0, 0, 0, 36, 309, 1, 0, 0, 0, 38, 321, 1, 0, 0, 0, 40, 328, 1, 0, 0,
		0, 42, 335, 1, 0, 0, 0, 44, 337, 1, 0, 0, 0, 46, 339, 1, 0, 0, 0, 48, 342,
		1, 0, 0, 0, 50, 345, 1, 0, 0, 0, 52, 357, 1, 0, 0, 0, 54, 359, 1, 0, 0,
		0, 56, 363, 1, 0, 0, 0, 58, 366, 1, 0, 0, 0, 60, 375, 1, 0, 0, 0, 62, 377,
		1, 0, 0, 0, 64, 380, 1, 0, 0, 0, 66, 383, 1, 0, 0, 0, 68, 387, 1, 0, 0,
		0, 70, 392, 1, 0, 0, 0, 72, 399, 1, 0, 0, 0, 74, 406, 1, 0, 0, 0, 76, 421,
		1, 0, 0, 0, 78, 448, 1, 0, 0, 0, 80, 451, 1, 0, 0, 0, 82, 466, 1, 0, 0,
		0, 84, 468, 1, 0, 0, 0, 86, 482, 1, 0, 0, 0, 88, 484, 1, 0, 0, 0, 90, 486,
		1, 0, 0, 0, 92, 490, 1, 0, 0, 0, 94, 494, 1, 0, 0, 0, 96, 499, 1, 0, 0,
		0, 98, 501, 1, 0, 0, 0, 100, 503, 1, 0, 0, 0, 102, 505, 1, 0, 0, 0, 104,
		509, 1, 0, 0, 0, 106, 516, 1, 0, 0, 0, 108, 518, 1, 0, 0, 0, 110, 528,
		1, 0, 0, 0, 112, 530, 1, 0, 0, 0, 114, 532, 1, 0, 0, 0, 116, 535, 1, 0,
		0, 0, 118, 542, 1, 0, 0, 0, 120, 563, 1, 0, 0, 0, 122, 565, 1, 0, 0, 0,
		124, 569, 1, 0, 0, 0, 126, 571, 1, 0, 0, 0, 128, 575, 1, 0, 0, 0, 130,
		587, 1, 0, 0, 0, 132, 594, 1, 0, 0, 0, 134, 596, 1, 0, 0, 0, 136, 598,
		1, 0, 0, 0, 138, 606, 1, 0, 0, 0, 140, 609, 1, 0, 0, 0, 142, 611, 1, 0,
		0, 0, 144, 613, 1, 0, 0, 0, 146, 615, 1, 0, 0, 0, 148, 617, 1, 0, 0, 0,
		150, 629, 1, 0, 0, 0, 152, 631, 1, 0, 0, 0, 154, 638, 1, 0, 0, 0, 156,
		641, 1, 0, 0, 0, 158, 645, 1, 0, 0, 0, 160, 647, 1, 0, 0, 0, 162, 650,
		1, 0, 0, 0, 164, 165, 3, 2, 1, 0, 165, 166, 5, 0, 0, 1, 166, 1, 1, 0, 0,
		0, 167, 172, 3, 38, 19, 0, 168, 169, 5, 1, 0, 0, 169, 171, 3, 38, 19, 0,
		170, 168, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172,
		173, 1, 0, 0, 0, 173, 3, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 176, 3,
		6, 3, 0, 176, 177, 5, 0, 0, 1, 177, 5, 1, 0, 0, 0, 178, 179, 6, 3, -1,
		0, 179, 180, 3, 34, 17, 0, 180, 181, 3, 8, 4, 0, 181, 182, 3, 34, 17, 0,
		182, 211, 1, 0, 0, 0, 183, 184, 3, 34, 17, 0, 184, 185, 5, 17, 0, 0, 185,
		186, 5, 2, 0, 0, 186, 191, 3, 34, 17, 0, 187, 188, 5, 1, 0, 0, 188, 190,
		3, 34, 17, 0, 189, 187, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0, 191, 189, 1,
		0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 1, 0, 0, 0, 193, 191, 1, 0, 0,
		0, 194, 195, 5, 3, 0, 0, 195, 211, 1, 0, 0, 0, 196, 197, 3, 34, 17, 0,
		197, 198, 5, 18, 0, 0, 198, 199, 3, 34, 17, 0, 199, 200, 5, 20, 0, 0, 200,
		201, 3, 34, 17, 0, 201, 211, 1, 0, 0, 0, 202, 211, 3, 36, 18, 0, 203, 204,
		5, 2, 0, 0, 204, 205, 3, 6, 3, 0, 205, 206, 5, 3, 0, 0, 206, 207, 6, 3,
		-1, 0, 207, 211, 1, 0, 0, 0, 208, 209, 5, 19, 0, 0, 209, 211, 3, 6, 3,
		3, 210, 178, 1, 0, 0, 0, 210, 183, 1, 0, 0, 0, 210, 196, 1, 0, 0, 0, 210,
		202, 1, 0, 0, 0, 210, 203, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 220,
		1, 0, 0, 0, 212, 213, 10, 2, 0, 0, 213, 214, 5, 20, 0, 0, 214, 219, 3,
		6, 3, 3, 215, 216, 10, 1, 0, 0, 216, 217, 5, 21, 0, 0, 217, 219, 3, 6,
		3, 2, 218, 212, 1, 0, 0, 0, 218, 215, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0,
		220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 7, 1, 0, 0, 0, 222, 220,
		1, 0, 0, 0, 223, 224, 7, 0, 0, 0, 224, 9, 1, 0, 0, 0, 225, 226, 3, 12,
		6, 0, 226, 227, 5, 0, 0, 1, 227, 11, 1, 0, 0, 0, 228, 233, 3, 14, 7, 0,
		229, 233, 3, 18, 9, 0, 230, 233, 3, 22, 11, 0, 231, 233, 3, 26, 13, 0,
		232, 228, 1, 0, 0, 0, 232, 229, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232,
		231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235,
		1, 0, 0, 0, 235, 13, 1, 0, 0, 0, 236, 237, 5, 22, 0, 0, 237, 242, 3, 16,
		8, 0, 238, 239, 5, 1, 0, 0, 239, 241, 3, 16, 8, 0, 240, 238, 1, 0, 0, 0,
		241, 244, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243,
		15, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 245, 246, 3, 38, 19, 0, 246, 247,
		5, 9, 0, 0, 247, 248, 3, 30, 15, 0, 248, 17, 1, 0, 0, 0, 249, 250, 5, 23,
		0, 0, 250, 255, 3, 20, 10, 0, 251, 252, 5, 1, 0, 0, 252, 254, 3, 20, 10,
		0, 253, 251, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255,
		256, 1, 0, 0, 0, 256, 19, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 259, 3,
		38, 19, 0, 259, 260, 3, 44, 22, 0, 260, 21, 1, 0, 0, 0, 261, 262, 5, 24,
		0, 0, 262, 267, 3, 24, 12, 0, 263, 264, 5, 1, 0, 0, 264, 266, 3, 24, 12,
		0, 265, 263, 1, 0, 0, 0, 266, 269, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267,
		268, 1, 0, 0, 0, 268, 23, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 270, 271, 3,
		38, 19, 0, 271, 272, 3, 44, 22, 0, 272, 25, 1, 0, 0, 0, 273, 274, 5, 25,
		0, 0, 274, 279, 3, 28, 14, 0, 275, 276, 5, 1, 0, 0, 276, 278, 3, 28, 14,
		0, 277, 275, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279,
		280, 1, 0, 0, 0, 280, 27, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 283, 3,
		38, 19, 0, 283, 29, 1, 0, 0, 0, 284, 287, 3, 34, 17, 0, 285, 287, 3, 32,
		16, 0, 286, 284, 1, 0, 0, 0, 286, 285, 1, 0, 0, 0, 287, 31, 1, 0, 0, 0,
		288, 289, 3, 34, 17, 0, 289, 290, 7, 1, 0, 0, 290, 291, 3, 34, 17, 0, 291,
		298, 1, 0, 0, 0, 292, 293, 5, 2, 0, 0, 293, 294, 3, 32, 16, 0, 294, 295,
		5, 3, 0, 0, 295, 296, 6, 16, -1, 0, 296, 298, 1, 0, 0, 0, 297, 288, 1,
		0, 0, 0, 297, 292, 1, 0, 0, 0, 298, 33, 1, 0, 0, 0, 299, 308, 3, 38, 19,
		0, 300, 308, 3, 44, 22, 0, 301, 308, 3, 36, 18, 0, 302, 303, 5, 2, 0, 0,
		303, 304, 3, 34, 17, 0, 304, 305, 5, 3, 0, 0, 305, 306, 6, 17, -1, 0, 306,
		308, 1, 0, 0, 0, 307, 299, 1, 0, 0, 0, 307, 300, 1, 0, 0, 0, 307, 301,
		1, 0, 0, 0, 307, 302, 1, 0, 0, 0, 308, 35, 1, 0, 0, 0, 309, 310, 5, 67,
		0, 0, 310, 311, 5, 2, 0, 0, 311, 316, 3, 34, 17, 0, 312, 313, 5, 1, 0,
		0, 313, 315, 3, 34, 17, 0, 314, 312, 1, 0, 0, 0, 315, 318, 1, 0, 0, 0,
		316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 319, 1, 0, 0, 0, 318,
		316, 1, 0, 0, 0, 319, 320, 5, 3, 0, 0, 320, 37, 1, 0, 0, 0, 321, 325, 3,
		40, 20, 0, 322, 324, 3, 42, 21, 0, 323, 322, 1, 0, 0, 0, 324, 327, 1, 0,
		0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 39, 1, 0, 0, 0,
		327, 325, 1, 0, 0, 0, 328, 329, 7, 2, 0, 0, 329, 41, 1, 0, 0, 0, 330, 331,
		5, 4, 0, 0, 331, 336, 3, 40, 20, 0, 332, 333, 5, 5, 0, 0, 333, 334, 5,
		66, 0, 0, 334, 336, 5, 6, 0, 0, 335, 330, 1, 0, 0, 0, 335, 332, 1, 0, 0,
		0, 336, 43, 1, 0, 0, 0, 337, 338, 5, 69, 0, 0, 338, 45, 1, 0, 0, 0, 339,
		340, 5, 68, 0, 0, 340, 341, 5, 0, 0, 1, 341, 47, 1, 0, 0, 0, 342, 343,
		5, 69, 0, 0, 343, 344, 5, 0, 0, 1, 344, 49, 1, 0, 0, 0, 345, 346, 3, 52,
		26, 0, 346, 347, 5, 0, 0, 1, 347, 51, 1, 0, 0, 0, 348, 358, 3, 74, 37,
		0, 349, 358, 3, 76, 38, 0, 350, 358, 3, 78, 39, 0, 351, 358, 3, 70, 35,
		0, 352, 358, 3, 72, 36, 0, 353, 358, 3, 80, 40, 0, 354, 358, 3, 54, 27,
		0, 355, 358, 3, 58, 29, 0, 356, 358, 3, 56, 28, 0, 357, 348, 1, 0, 0, 0,
		357, 349, 1, 0, 0, 0, 357, 350, 1, 0, 0, 0, 357, 351, 1, 0, 0, 0, 357,
		352, 1, 0, 0, 0, 357, 353, 1, 0, 0, 0, 357, 354, 1, 0, 0, 0, 357, 355,
		1, 0, 0, 0, 357, 356, 1, 0, 0, 0, 358, 53, 1, 0, 0, 0, 359, 360, 5, 62,
		0, 0, 360, 361, 5, 43, 0, 0, 361, 362, 3, 156, 78, 0, 362, 55, 1, 0, 0,
		0, 363, 364, 5, 64, 0, 0, 364, 365, 3, 156, 78, 0, 365, 57, 1, 0, 0, 0,
		366, 367, 5, 63, 0, 0, 367, 368, 5, 43, 0, 0, 368, 369, 3, 156, 78, 0,
		369, 370, 3, 60, 30, 0, 370, 59, 1, 0, 0, 0, 371, 376, 3, 62, 31, 0, 372,
		376, 3, 64, 32, 0, 373, 376, 3, 66, 33, 0, 374, 376, 3, 68, 34, 0, 375,
		371, 1, 0, 0, 0, 375, 372, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 374,
		1, 0, 0, 0, 376, 61, 1, 0, 0, 0, 377, 378, 5, 22, 0, 0, 378, 379, 3, 130,
		65, 0, 379, 63, 1, 0, 0, 0, 380, 381, 5, 23, 0, 0, 381, 382, 3, 118, 59,
		0, 382, 65, 1, 0, 0, 0, 383, 384, 5, 62, 0, 0, 384, 385, 5, 52, 0, 0, 385,
		386, 3, 122, 61, 0, 386, 67, 1, 0, 0, 0, 387, 388, 5, 63, 0, 0, 388, 389,
		5, 52, 0, 0, 389, 390, 3, 122, 61, 0, 390, 391, 3, 62, 31, 0, 391, 69,
		1, 0, 0, 0, 392, 393, 5, 30, 0, 0, 393, 394, 3, 156, 78, 0, 394, 395, 3,
		12, 6, 0, 395, 397, 3, 114, 57, 0, 396, 398, 3, 148, 74, 0, 397, 396, 1,
		0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 71, 1, 0, 0, 0, 399, 400, 5, 24, 0,
		0, 400, 401, 5, 57, 0, 0, 401, 402, 3, 156, 78, 0, 402, 404, 3, 114, 57,
		0, 403, 405, 3, 148, 74, 0, 404, 403, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0,
		405, 73, 1, 0, 0, 0, 406, 408, 5, 26, 0, 0, 407, 409, 5, 27, 0, 0, 408,
		407, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 412,
		3, 156, 78, 0, 411, 413, 7, 3, 0, 0, 412, 411, 1, 0, 0, 0, 412, 413, 1,
		0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 416, 3, 44, 22, 0, 415, 417, 3, 152,
		76, 0, 416, 415, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 419, 1, 0, 0, 0,
		418, 420, 3, 148, 74, 0, 419, 418, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420,
		75, 1, 0, 0, 0, 421, 422, 5, 42, 0, 0, 422, 423, 5, 43, 0, 0, 423, 424,
		3, 156, 78, 0, 424, 425, 5, 2, 0, 0, 425, 430, 3, 138, 69, 0, 426, 427,
		5, 1, 0, 0, 427, 429, 3, 138, 69, 0, 428, 426, 1, 0, 0, 0, 429, 432, 1,
		0, 0, 0, 430, 428, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 433, 1, 0, 0,
		0, 432, 430, 1, 0, 0, 0, 433, 434, 5, 1, 0, 0, 434, 444, 3, 116, 58, 0,
		435, 436, 5, 1, 0, 0, 436, 441, 3, 118, 59, 0, 437, 438, 5, 1, 0, 0, 438,
		440, 3, 118, 59, 0, 439, 437, 1, 0, 0, 0, 440, 443, 1, 0, 0, 0, 441, 439,
		1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0,
		0, 0, 444, 435, 1, 0, 0, 0, 444, 445, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0,
		446, 447, 5, 3, 0, 0, 447, 77, 1, 0, 0, 0, 448, 449, 5, 54, 0, 0, 449,
		450, 5, 55, 0, 0, 450, 79, 1, 0, 0, 0, 451, 452, 5, 56, 0, 0, 452, 453,
		3, 82, 41, 0, 453, 454, 5, 57, 0, 0, 454, 456, 3, 156, 78, 0, 455, 457,
		3, 102, 51, 0, 456, 455, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 459, 1,
		0, 0, 0, 458, 460, 3, 114, 57, 0, 459, 458, 1, 0, 0, 0, 459, 460, 1, 0,
		0, 0, 460, 462, 1, 0, 0, 0, 461, 463, 3, 84, 42, 0, 462, 461, 1, 0, 0,
		0, 462, 463, 1, 0, 0, 0, 463, 81, 1, 0, 0, 0, 464, 467, 3, 2, 1, 0, 465,
		467, 3, 100, 50, 0, 466, 464, 1, 0, 0, 0, 466, 465, 1, 0, 0, 0, 467, 83,
		1, 0, 0, 0, 468, 469, 5, 65, 0, 0, 469, 470, 5, 2, 0, 0, 470, 475, 3, 86,
		43, 0, 471, 472, 5, 1, 0, 0, 472, 474, 3, 86, 43, 0, 473, 471, 1, 0, 0,
		0, 474, 477, 1, 0, 0, 0, 475, 473, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476,
		478, 1, 0, 0, 0, 477, 475, 1, 0, 0, 0, 478, 479, 5, 3, 0, 0, 479, 85, 1,
		0, 0, 0, 480, 483, 3, 88, 44, 0, 481, 483, 3, 90, 45, 0, 482, 480, 1, 0,
		0, 0, 482, 481, 1, 0, 0, 0, 483, 87, 1, 0, 0, 0, 484, 485, 3, 92, 46, 0,
		485, 89, 1, 0, 0, 0, 486, 487, 3, 92, 46, 0, 487, 488, 5, 9, 0, 0, 488,
		489, 3, 94, 47, 0, 489, 91, 1, 0, 0, 0, 490, 491, 3, 158, 79, 0, 491, 93,
		1, 0, 0, 0, 492, 495, 3, 96, 48, 0, 493, 495, 3, 98, 49, 0, 494, 492, 1,
		0, 0, 0, 494, 493, 1, 0, 0, 0, 495, 95, 1, 0, 0, 0, 496, 500, 3, 160, 80,
		0, 497, 500, 5, 38, 0, 0, 498, 500, 5, 67, 0, 0, 499, 496, 1, 0, 0, 0,
		499, 497, 1, 0, 0, 0, 499, 498, 1, 0, 0, 0, 500, 97, 1, 0, 0, 0, 501, 502,
		5, 66, 0, 0, 502, 99, 1, 0, 0, 0, 503, 504, 5, 7, 0, 0, 504, 101, 1, 0,
		0, 0, 505, 507, 3, 104, 52, 0, 506, 508, 3, 108, 54, 0, 507, 506, 1, 0,
		0, 0, 507, 508, 1, 0, 0, 0, 508, 103, 1, 0, 0, 0, 509, 510, 5, 58, 0, 0,
		510, 511, 3, 106, 53, 0, 511, 105, 1, 0, 0, 0, 512, 513, 5, 52, 0, 0, 513,
		517, 3, 122, 61, 0, 514, 515, 5, 44, 0, 0, 515, 517, 5, 52, 0, 0, 516,
		512, 1, 0, 0, 0, 516, 514, 1, 0, 0, 0, 517, 107, 1, 0, 0, 0, 518, 519,
		5, 59, 0, 0, 519, 526, 5, 60, 0, 0, 520, 521, 5, 2, 0, 0, 521, 522, 3,
		110, 55, 0, 522, 523, 5, 1, 0, 0, 523, 524, 3, 112, 56, 0, 524, 525, 5,
		3, 0, 0, 525, 527, 1, 0, 0, 0, 526, 520, 1, 0, 0, 0, 526, 527, 1, 0, 0,
		0, 527, 109, 1, 0, 0, 0, 528, 529, 5, 66, 0, 0, 529, 111, 1, 0, 0, 0, 530,
		531, 5, 66, 0, 0, 531, 113, 1, 0, 0, 0, 532, 533, 5, 61, 0, 0, 533, 534,
		3, 6, 3, 0, 534, 115, 1, 0, 0, 0, 535, 536, 5, 44, 0, 0, 536, 537, 5, 35,
		0, 0, 537, 539, 3, 136, 68, 0, 538, 540, 3, 130, 65, 0, 539, 538, 1, 0,
		0, 0, 539, 540, 1, 0, 0, 0, 540, 117, 1, 0, 0, 0, 541, 543, 3, 120, 60,
		0, 542, 541, 1, 0, 0, 0, 542, 543, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544,
		553, 5, 52, 0, 0, 545, 550, 3, 138, 69, 0, 546, 547, 5, 1, 0, 0, 547, 549,
		3, 138, 69, 0, 548, 546, 1, 0, 0, 0, 549, 552, 1, 0, 0, 0, 550, 548, 1,
		0, 0, 0, 550, 551, 1, 0, 0, 0, 551, 554, 1, 0, 0, 0, 552, 550, 1, 0, 0,
		0, 553, 545, 1, 0, 0, 0, 553, 554, 1, 0, 0, 0, 554, 555, 1, 0, 0, 0, 555,
		556, 3, 122, 61, 0, 556, 558, 3, 136, 68, 0, 557, 559, 3, 124, 62, 0, 558,
		557, 1, 0, 0, 0, 558, 559, 1, 0, 0, 0, 559, 561, 1, 0, 0, 0, 560, 562,
		3, 130, 65, 0, 561, 560, 1, 0, 0, 0, 561, 562, 1, 0, 0, 0, 562, 119, 1,
		0, 0, 0, 563, 564, 7, 4, 0, 0, 564, 121, 1, 0, 0, 0, 565, 566, 3, 158,
		79, 0, 566, 123, 1, 0, 0, 0, 567, 570, 3, 126, 63, 0, 568, 570, 3, 128,
		64, 0, 569, 567, 1, 0, 0, 0, 569, 568, 1, 0, 0, 0, 570, 125, 1, 0, 0, 0,
		571, 572, 5, 53, 0, 0, 572, 573, 5, 34, 0, 0, 573, 574, 5, 36, 0, 0, 574,
		127, 1, 0, 0, 0, 575, 576, 5, 53, 0, 0, 576, 577, 5, 2, 0, 0, 577, 582,
		3, 144, 72, 0, 578, 579, 5, 1, 0, 0, 579, 581, 3, 144, 72, 0, 580, 578,
		1, 0, 0, 0, 581, 584, 1, 0, 0, 0, 582, 580, 1, 0, 0, 0, 582, 583, 1, 0,
		0, 0, 583, 585, 1, 0, 0, 0, 584, 582, 1, 0, 0, 0, 585, 586, 5, 3, 0, 0,
		586, 129, 1, 0, 0, 0, 587, 588, 5, 49, 0, 0, 588, 589, 5, 2, 0, 0, 589,
		590, 3, 132, 66, 0, 590, 591, 5, 1, 0, 0, 591, 592, 3, 134, 67, 0, 592,
		593, 5, 3, 0, 0, 593, 131, 1, 0, 0, 0, 594, 595, 5, 66, 0, 0, 595, 133,
		1, 0, 0, 0, 596, 597, 5, 66, 0, 0, 597, 135, 1, 0, 0, 0, 598, 599, 5, 2,
		0, 0, 599, 602, 3, 140, 70, 0, 600, 601, 5, 1, 0, 0, 601, 603, 3, 142,
		71, 0, 602, 600, 1, 0, 0, 0, 602, 603, 1, 0, 0, 0, 603, 604, 1, 0, 0, 0,
		604, 605, 5, 3, 0, 0, 605, 137, 1, 0, 0, 0, 606, 607, 3, 144, 72, 0, 607,
		608, 3, 146, 73, 0, 608, 139, 1, 0, 0, 0, 609, 610, 3, 158, 79, 0, 610,
		141, 1, 0, 0, 0, 611, 612, 3, 158, 79, 0, 612, 143, 1, 0, 0, 0, 613, 614,
		3, 158, 79, 0, 614, 145, 1, 0, 0, 0, 615, 616, 7, 5, 0, 0, 616, 147, 1,
		0, 0, 0, 617, 618, 5, 32, 0, 0, 618, 619, 3, 150, 75, 0, 619, 149, 1, 0,
		0, 0, 620, 630, 5, 38, 0, 0, 621, 622, 5, 33, 0, 0, 622, 630, 5, 39, 0,
		0, 623, 624, 5, 40, 0, 0, 624, 630, 5, 39, 0, 0, 625, 626, 5, 33, 0, 0,
		626, 630, 5, 37, 0, 0, 627, 628, 5, 40, 0, 0, 628, 630, 5, 37, 0, 0, 629,
		620, 1, 0, 0, 0, 629, 621, 1, 0, 0, 0, 629, 623, 1, 0, 0, 0, 629, 625,
		1, 0, 0, 0, 629, 627, 1, 0, 0, 0, 630, 151, 1, 0, 0, 0, 631, 632, 5, 31,
		0, 0, 632, 633, 5, 41, 0, 0, 633, 634, 5, 35, 0, 0, 634, 636, 5, 30, 0,
		0, 635, 637, 3, 154, 77, 0, 636, 635, 1, 0, 0, 0, 636, 637, 1, 0, 0, 0,
		637, 153, 1, 0, 0, 0, 638, 639, 5, 48, 0, 0, 639, 640, 3, 6, 3, 0, 640,
		155, 1, 0, 0, 0, 641, 642, 3, 158, 79, 0, 642, 157, 1, 0, 0, 0, 643, 646,
		5, 67, 0, 0, 644, 646, 3, 160, 80, 0, 645, 643, 1, 0, 0, 0, 645, 644, 1,
		0, 0, 0, 646, 159, 1, 0, 0, 0, 647, 648, 5, 70, 0, 0, 648, 161, 1, 0, 0,
		0, 649, 651, 5, 71, 0, 0, 650, 649, 1, 0, 0, 0, 651, 652, 1, 0, 0, 0, 652,
		650, 1, 0, 0, 0, 652, 653, 1, 0, 0, 0, 653, 163, 1, 0, 0, 0, 52, 172, 191,
		210, 218, 220, 232, 234, 242, 255, 267, 279, 286, 297, 307, 316, 325, 335,
		357, 375, 397, 404, 408, 412, 416, 419, 430, 441, 444, 456, 459, 462, 466,
		475, 482, 494, 499, 507, 516, 526, 539, 542, 550, 553, 558, 561, 569, 582,
		602, 629, 636, 645, 652,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// DynamoDbGrammarParserInit initializes any static state used to implement DynamoDbGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDynamoDbGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DynamoDbGrammarParserInit() {
	staticData := &DynamoDbGrammarParserStaticData
	staticData.once.Do(dynamodbgrammarParserInit)
}

// NewDynamoDbGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDynamoDbGrammarParser(input antlr.TokenStream) *DynamoDbGrammarParser {
	DynamoDbGrammarParserInit()
	this := new(DynamoDbGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DynamoDbGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "DynamoDbGrammar.g4"

	return this
}

// Note that '@members' cannot be changed now, but this should have been 'globals'
// If you are looking to have variables for each instance, use '@structmembers'

// private static void validateRedundantParentheses(bool redundantParens) {
//     if (redundantParens) {
//         throw new RedundantParenthesesException();
//     }
// }

// DynamoDbGrammarParser tokens.
const (
	DynamoDbGrammarParserEOF                = antlr.TokenEOF
	DynamoDbGrammarParserT__0               = 1
	DynamoDbGrammarParserT__1               = 2
	DynamoDbGrammarParserT__2               = 3
	DynamoDbGrammarParserT__3               = 4
	DynamoDbGrammarParserT__4               = 5
	DynamoDbGrammarParserT__5               = 6
	DynamoDbGrammarParserT__6               = 7
	DynamoDbGrammarParserWS                 = 8
	DynamoDbGrammarParserEQ                 = 9
	DynamoDbGrammarParserNE                 = 10
	DynamoDbGrammarParserLT                 = 11
	DynamoDbGrammarParserLE                 = 12
	DynamoDbGrammarParserGT                 = 13
	DynamoDbGrammarParserGE                 = 14
	DynamoDbGrammarParserPLUS               = 15
	DynamoDbGrammarParserMINUS              = 16
	DynamoDbGrammarParserIN                 = 17
	DynamoDbGrammarParserBETWEEN            = 18
	DynamoDbGrammarParserNOT                = 19
	DynamoDbGrammarParserAND                = 20
	DynamoDbGrammarParserOR                 = 21
	DynamoDbGrammarParserSET                = 22
	DynamoDbGrammarParserADD                = 23
	DynamoDbGrammarParserDELETE             = 24
	DynamoDbGrammarParserREMOVE             = 25
	DynamoDbGrammarParserINSERT             = 26
	DynamoDbGrammarParserINTO               = 27
	DynamoDbGrammarParserVALUE              = 28
	DynamoDbGrammarParserVALUES             = 29
	DynamoDbGrammarParserUPDATE             = 30
	DynamoDbGrammarParserON                 = 31
	DynamoDbGrammarParserRETURNING          = 32
	DynamoDbGrammarParserALL                = 33
	DynamoDbGrammarParserKEYS               = 34
	DynamoDbGrammarParserKEY                = 35
	DynamoDbGrammarParserONLY               = 36
	DynamoDbGrammarParserNEW                = 37
	DynamoDbGrammarParserNONE               = 38
	DynamoDbGrammarParserOLD                = 39
	DynamoDbGrammarParserUPDATED            = 40
	DynamoDbGrammarParserDUPLICATE          = 41
	DynamoDbGrammarParserCREATE             = 42
	DynamoDbGrammarParserTABLE              = 43
	DynamoDbGrammarParserPRIMARY            = 44
	DynamoDbGrammarParserNUMBER             = 45
	DynamoDbGrammarParserBINARY             = 46
	DynamoDbGrammarParserSTRING             = 47
	DynamoDbGrammarParserIF                 = 48
	DynamoDbGrammarParserCAPACITY           = 49
	DynamoDbGrammarParserGLOBAL             = 50
	DynamoDbGrammarParserLOCAL              = 51
	DynamoDbGrammarParserINDEXKEYWORD       = 52
	DynamoDbGrammarParserPROJECTION         = 53
	DynamoDbGrammarParserSHOW               = 54
	DynamoDbGrammarParserTABLES             = 55
	DynamoDbGrammarParserSELECT             = 56
	DynamoDbGrammarParserFROM               = 57
	DynamoDbGrammarParserUSE                = 58
	DynamoDbGrammarParserENABLE             = 59
	DynamoDbGrammarParserSCAN               = 60
	DynamoDbGrammarParserWHERE              = 61
	DynamoDbGrammarParserDROP               = 62
	DynamoDbGrammarParserALTER              = 63
	DynamoDbGrammarParserDESCRIBE           = 64
	DynamoDbGrammarParserOPTION             = 65
	DynamoDbGrammarParserINDEX              = 66
	DynamoDbGrammarParserID                 = 67
	DynamoDbGrammarParserATTRIBUTE_NAME_SUB = 68
	DynamoDbGrammarParserLITERAL_SUB        = 69
	DynamoDbGrammarParserSTRING_LITERAL     = 70
	DynamoDbGrammarParserUNKNOWN            = 71
)

// DynamoDbGrammarParser rules.
const (
	DynamoDbGrammarParserRULE_projection_                = 0
	DynamoDbGrammarParserRULE_projection                 = 1
	DynamoDbGrammarParserRULE_condition_                 = 2
	DynamoDbGrammarParserRULE_condition                  = 3
	DynamoDbGrammarParserRULE_comparator_symbol          = 4
	DynamoDbGrammarParserRULE_update_                    = 5
	DynamoDbGrammarParserRULE_update                     = 6
	DynamoDbGrammarParserRULE_set_section                = 7
	DynamoDbGrammarParserRULE_set_action                 = 8
	DynamoDbGrammarParserRULE_add_section                = 9
	DynamoDbGrammarParserRULE_add_action                 = 10
	DynamoDbGrammarParserRULE_delete_section             = 11
	DynamoDbGrammarParserRULE_delete_action              = 12
	DynamoDbGrammarParserRULE_remove_section             = 13
	DynamoDbGrammarParserRULE_remove_action              = 14
	DynamoDbGrammarParserRULE_set_value                  = 15
	DynamoDbGrammarParserRULE_arithmetic                 = 16
	DynamoDbGrammarParserRULE_operand                    = 17
	DynamoDbGrammarParserRULE_function                   = 18
	DynamoDbGrammarParserRULE_path                       = 19
	DynamoDbGrammarParserRULE_id                         = 20
	DynamoDbGrammarParserRULE_dereference                = 21
	DynamoDbGrammarParserRULE_literal                    = 22
	DynamoDbGrammarParserRULE_expression_attr_names_sub  = 23
	DynamoDbGrammarParserRULE_expression_attr_values_sub = 24
	DynamoDbGrammarParserRULE_statement_                 = 25
	DynamoDbGrammarParserRULE_statement                  = 26
	DynamoDbGrammarParserRULE_dropTableStatement         = 27
	DynamoDbGrammarParserRULE_describeStatement          = 28
	DynamoDbGrammarParserRULE_alterTableStatement        = 29
	DynamoDbGrammarParserRULE_alterTableStatementType    = 30
	DynamoDbGrammarParserRULE_setCapacity                = 31
	DynamoDbGrammarParserRULE_addIndex                   = 32
	DynamoDbGrammarParserRULE_dropIndex                  = 33
	DynamoDbGrammarParserRULE_alterIndex                 = 34
	DynamoDbGrammarParserRULE_updateStatement            = 35
	DynamoDbGrammarParserRULE_deleteStatement            = 36
	DynamoDbGrammarParserRULE_insertStatement            = 37
	DynamoDbGrammarParserRULE_createTableStatement       = 38
	DynamoDbGrammarParserRULE_showTablesStatement        = 39
	DynamoDbGrammarParserRULE_selectStatement            = 40
	DynamoDbGrammarParserRULE_selectProjection           = 41
	DynamoDbGrammarParserRULE_optionBlock                = 42
	DynamoDbGrammarParserRULE_option                     = 43
	DynamoDbGrammarParserRULE_singleOption               = 44
	DynamoDbGrammarParserRULE_keyValueOption             = 45
	DynamoDbGrammarParserRULE_optionKey                  = 46
	DynamoDbGrammarParserRULE_optionValue                = 47
	DynamoDbGrammarParserRULE_optionValueString          = 48
	DynamoDbGrammarParserRULE_optionValueNumber          = 49
	DynamoDbGrammarParserRULE_star                       = 50
	DynamoDbGrammarParserRULE_hint                       = 51
	DynamoDbGrammarParserRULE_indexHint                  = 52
	DynamoDbGrammarParserRULE_indexHintName              = 53
	DynamoDbGrammarParserRULE_scanInfo                   = 54
	DynamoDbGrammarParserRULE_totalSegment               = 55
	DynamoDbGrammarParserRULE_segment                    = 56
	DynamoDbGrammarParserRULE_where                      = 57
	DynamoDbGrammarParserRULE_primaryKeyDecl             = 58
	DynamoDbGrammarParserRULE_secondaryIndexDecl         = 59
	DynamoDbGrammarParserRULE_secondaryIndexType         = 60
	DynamoDbGrammarParserRULE_indexName                  = 61
	DynamoDbGrammarParserRULE_projectionIndex            = 62
	DynamoDbGrammarParserRULE_projectionIndexKeysOnly    = 63
	DynamoDbGrammarParserRULE_projectionIndexVector      = 64
	DynamoDbGrammarParserRULE_capacity                   = 65
	DynamoDbGrammarParserRULE_readUnits                  = 66
	DynamoDbGrammarParserRULE_writeUnits                 = 67
	DynamoDbGrammarParserRULE_indexDecl                  = 68
	DynamoDbGrammarParserRULE_attributeDecl              = 69
	DynamoDbGrammarParserRULE_hashKey                    = 70
	DynamoDbGrammarParserRULE_rangeKey                   = 71
	DynamoDbGrammarParserRULE_attributeName              = 72
	DynamoDbGrammarParserRULE_attributeType              = 73
	DynamoDbGrammarParserRULE_returning                  = 74
	DynamoDbGrammarParserRULE_returningValue             = 75
	DynamoDbGrammarParserRULE_onDuplicateKeyUpdate       = 76
	DynamoDbGrammarParserRULE_ifClause                   = 77
	DynamoDbGrammarParserRULE_tableName                  = 78
	DynamoDbGrammarParserRULE_ddlName                    = 79
	DynamoDbGrammarParserRULE_string                     = 80
	DynamoDbGrammarParserRULE_unknown                    = 81
)

// IProjection_Context is an interface to support dynamic dispatch.
type IProjection_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Projection() IProjectionContext
	EOF() antlr.TerminalNode

	// IsProjection_Context differentiates from other interfaces.
	IsProjection_Context()
}

type Projection_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjection_Context() *Projection_Context {
	var p = new(Projection_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projection_
	return p
}

func InitEmptyProjection_Context(p *Projection_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projection_
}

func (*Projection_Context) IsProjection_Context() {}

func NewProjection_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Projection_Context {
	var p = new(Projection_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projection_

	return p
}

func (s *Projection_Context) GetParser() antlr.Parser { return s.parser }

func (s *Projection_Context) Projection() IProjectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionContext)
}

func (s *Projection_Context) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Projection_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Projection_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Projection_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjection_(s)
	}
}

func (s *Projection_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjection_(s)
	}
}

func (p *DynamoDbGrammarParser) Projection_() (localctx IProjection_Context) {
	localctx = NewProjection_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DynamoDbGrammarParserRULE_projection_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Projection()
	}
	{
		p.SetState(165)
		p.Match(DynamoDbGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProjectionContext is an interface to support dynamic dispatch.
type IProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPath() []IPathContext
	Path(i int) IPathContext

	// IsProjectionContext differentiates from other interfaces.
	IsProjectionContext()
}

type ProjectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionContext() *ProjectionContext {
	var p = new(ProjectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projection
	return p
}

func InitEmptyProjectionContext(p *ProjectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projection
}

func (*ProjectionContext) IsProjectionContext() {}

func NewProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionContext {
	var p = new(ProjectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projection

	return p
}

func (s *ProjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionContext) AllPath() []IPathContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPathContext); ok {
			len++
		}
	}

	tst := make([]IPathContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPathContext); ok {
			tst[i] = t.(IPathContext)
			i++
		}
	}

	return tst
}

func (s *ProjectionContext) Path(i int) IPathContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *ProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjection(s)
	}
}

func (s *ProjectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjection(s)
	}
}

func (p *DynamoDbGrammarParser) Projection() (localctx IProjectionContext) {
	localctx = NewProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DynamoDbGrammarParserRULE_projection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Path()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(168)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Path()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICondition_Context is an interface to support dynamic dispatch.
type ICondition_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Condition() IConditionContext
	EOF() antlr.TerminalNode

	// IsCondition_Context differentiates from other interfaces.
	IsCondition_Context()
}

type Condition_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondition_Context() *Condition_Context {
	var p = new(Condition_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_condition_
	return p
}

func InitEmptyCondition_Context(p *Condition_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_condition_
}

func (*Condition_Context) IsCondition_Context() {}

func NewCondition_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_Context {
	var p = new(Condition_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_condition_

	return p
}

func (s *Condition_Context) GetParser() antlr.Parser { return s.parser }

func (s *Condition_Context) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Condition_Context) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Condition_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterCondition_(s)
	}
}

func (s *Condition_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitCondition_(s)
	}
}

func (p *DynamoDbGrammarParser) Condition_() (localctx ICondition_Context) {
	localctx = NewCondition_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DynamoDbGrammarParserRULE_condition_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.condition(0)
	}
	{
		p.SetState(176)
		p.Match(DynamoDbGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHasOuterParens returns the hasOuterParens attribute.
	GetHasOuterParens() bool

	// SetHasOuterParens sets the hasOuterParens attribute.
	SetHasOuterParens(bool)

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	hasOuterParens bool // TODO = false
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) GetHasOuterParens() bool { return s.hasOuterParens }

func (s *ConditionContext) SetHasOuterParens(v bool) { s.hasOuterParens = v }

func (s *ConditionContext) CopyAll(ctx *ConditionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
	s.hasOuterParens = ctx.hasOuterParens
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OrContext struct {
	ConditionContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserOR, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOr(s)
	}
}

type NegationContext struct {
	ConditionContext
}

func NewNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationContext {
	var p = new(NegationContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) NOT() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNOT, 0)
}

func (s *NegationContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *NegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterNegation(s)
	}
}

func (s *NegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitNegation(s)
	}
}

type InContext struct {
	ConditionContext
}

func NewInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InContext {
	var p = new(InContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *InContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *InContext) IN() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserIN, 0)
}

func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIn(s)
	}
}

type AndContext struct {
	ConditionContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserAND, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAnd(s)
	}
}

type BetweenContext struct {
	ConditionContext
}

func NewBetweenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BetweenContext {
	var p = new(BetweenContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *BetweenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *BetweenContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *BetweenContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserBETWEEN, 0)
}

func (s *BetweenContext) AND() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserAND, 0)
}

func (s *BetweenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterBetween(s)
	}
}

func (s *BetweenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitBetween(s)
	}
}

type FunctionConditionContext struct {
	ConditionContext
}

func NewFunctionConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionConditionContext {
	var p = new(FunctionConditionContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *FunctionConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionConditionContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterFunctionCondition(s)
	}
}

func (s *FunctionConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitFunctionCondition(s)
	}
}

type ComparatorContext struct {
	ConditionContext
}

func NewComparatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparatorContext {
	var p = new(ComparatorContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *ComparatorContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *ComparatorContext) Comparator_symbol() IComparator_symbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparator_symbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparator_symbolContext)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitComparator(s)
	}
}

type ConditionGroupingContext struct {
	ConditionContext
	c IConditionContext
}

func NewConditionGroupingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionGroupingContext {
	var p = new(ConditionGroupingContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *ConditionGroupingContext) GetC() IConditionContext { return s.c }

func (s *ConditionGroupingContext) SetC(v IConditionContext) { s.c = v }

func (s *ConditionGroupingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionGroupingContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionGroupingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterConditionGrouping(s)
	}
}

func (s *ConditionGroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitConditionGrouping(s)
	}
}

func (p *DynamoDbGrammarParser) Condition() (localctx IConditionContext) {
	return p.condition(0)
}

func (p *DynamoDbGrammarParser) condition(_p int) (localctx IConditionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewConditionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, DynamoDbGrammarParserRULE_condition, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewComparatorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(179)
			p.Operand()
		}
		{
			p.SetState(180)
			p.Comparator_symbol()
		}
		{
			p.SetState(181)
			p.Operand()
		}

	case 2:
		localctx = NewInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(183)
			p.Operand()
		}
		{
			p.SetState(184)
			p.Match(DynamoDbGrammarParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(DynamoDbGrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Operand()
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DynamoDbGrammarParserT__0 {
			{
				p.SetState(187)
				p.Match(DynamoDbGrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(188)
				p.Operand()
			}

			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(194)
			p.Match(DynamoDbGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewBetweenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(196)
			p.Operand()
		}
		{
			p.SetState(197)
			p.Match(DynamoDbGrammarParserBETWEEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Operand()
		}
		{
			p.SetState(199)
			p.Match(DynamoDbGrammarParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Operand()
		}

	case 4:
		localctx = NewFunctionConditionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(202)
			p.Function()
		}

	case 5:
		localctx = NewConditionGroupingContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(203)
			p.Match(DynamoDbGrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)

			var _x = p.condition(0)

			localctx.(*ConditionGroupingContext).c = _x
		}
		{
			p.SetState(205)
			p.Match(DynamoDbGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		// validateRedundantParentheses(localctx.(*ConditionGroupingContext).GetC().GetHasOuterParens());
		localctx.(*ConditionGroupingContext).SetHasOuterParens(true)

	case 6:
		localctx = NewNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(208)
			p.Match(DynamoDbGrammarParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.condition(3)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndContext(p, NewConditionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DynamoDbGrammarParserRULE_condition)
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(213)
					p.Match(DynamoDbGrammarParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(214)
					p.condition(3)
				}

			case 2:
				localctx = NewOrContext(p, NewConditionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DynamoDbGrammarParserRULE_condition)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(216)
					p.Match(DynamoDbGrammarParserOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(217)
					p.condition(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparator_symbolContext is an interface to support dynamic dispatch.
type IComparator_symbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode

	// IsComparator_symbolContext differentiates from other interfaces.
	IsComparator_symbolContext()
}

type Comparator_symbolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparator_symbolContext() *Comparator_symbolContext {
	var p = new(Comparator_symbolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_comparator_symbol
	return p
}

func InitEmptyComparator_symbolContext(p *Comparator_symbolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_comparator_symbol
}

func (*Comparator_symbolContext) IsComparator_symbolContext() {}

func NewComparator_symbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparator_symbolContext {
	var p = new(Comparator_symbolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_comparator_symbol

	return p
}

func (s *Comparator_symbolContext) GetParser() antlr.Parser { return s.parser }

func (s *Comparator_symbolContext) EQ() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEQ, 0)
}

func (s *Comparator_symbolContext) NE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNE, 0)
}

func (s *Comparator_symbolContext) LT() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserLT, 0)
}

func (s *Comparator_symbolContext) LE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserLE, 0)
}

func (s *Comparator_symbolContext) GT() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserGT, 0)
}

func (s *Comparator_symbolContext) GE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserGE, 0)
}

func (s *Comparator_symbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparator_symbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparator_symbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterComparator_symbol(s)
	}
}

func (s *Comparator_symbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitComparator_symbol(s)
	}
}

func (p *DynamoDbGrammarParser) Comparator_symbol() (localctx IComparator_symbolContext) {
	localctx = NewComparator_symbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DynamoDbGrammarParserRULE_comparator_symbol)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32256) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdate_Context is an interface to support dynamic dispatch.
type IUpdate_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Update() IUpdateContext
	EOF() antlr.TerminalNode

	// IsUpdate_Context differentiates from other interfaces.
	IsUpdate_Context()
}

type Update_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_Context() *Update_Context {
	var p = new(Update_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_update_
	return p
}

func InitEmptyUpdate_Context(p *Update_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_update_
}

func (*Update_Context) IsUpdate_Context() {}

func NewUpdate_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_Context {
	var p = new(Update_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_update_

	return p
}

func (s *Update_Context) GetParser() antlr.Parser { return s.parser }

func (s *Update_Context) Update() IUpdateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateContext)
}

func (s *Update_Context) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Update_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterUpdate_(s)
	}
}

func (s *Update_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitUpdate_(s)
	}
}

func (p *DynamoDbGrammarParser) Update_() (localctx IUpdate_Context) {
	localctx = NewUpdate_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DynamoDbGrammarParserRULE_update_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Update()
	}
	{
		p.SetState(226)
		p.Match(DynamoDbGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdateContext is an interface to support dynamic dispatch.
type IUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSet_section() []ISet_sectionContext
	Set_section(i int) ISet_sectionContext
	AllAdd_section() []IAdd_sectionContext
	Add_section(i int) IAdd_sectionContext
	AllDelete_section() []IDelete_sectionContext
	Delete_section(i int) IDelete_sectionContext
	AllRemove_section() []IRemove_sectionContext
	Remove_section(i int) IRemove_sectionContext

	// IsUpdateContext differentiates from other interfaces.
	IsUpdateContext()
}

type UpdateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateContext() *UpdateContext {
	var p = new(UpdateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_update
	return p
}

func InitEmptyUpdateContext(p *UpdateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_update
}

func (*UpdateContext) IsUpdateContext() {}

func NewUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateContext {
	var p = new(UpdateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_update

	return p
}

func (s *UpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateContext) AllSet_section() []ISet_sectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISet_sectionContext); ok {
			len++
		}
	}

	tst := make([]ISet_sectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISet_sectionContext); ok {
			tst[i] = t.(ISet_sectionContext)
			i++
		}
	}

	return tst
}

func (s *UpdateContext) Set_section(i int) ISet_sectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_sectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_sectionContext)
}

func (s *UpdateContext) AllAdd_section() []IAdd_sectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdd_sectionContext); ok {
			len++
		}
	}

	tst := make([]IAdd_sectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdd_sectionContext); ok {
			tst[i] = t.(IAdd_sectionContext)
			i++
		}
	}

	return tst
}

func (s *UpdateContext) Add_section(i int) IAdd_sectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdd_sectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdd_sectionContext)
}

func (s *UpdateContext) AllDelete_section() []IDelete_sectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDelete_sectionContext); ok {
			len++
		}
	}

	tst := make([]IDelete_sectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDelete_sectionContext); ok {
			tst[i] = t.(IDelete_sectionContext)
			i++
		}
	}

	return tst
}

func (s *UpdateContext) Delete_section(i int) IDelete_sectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelete_sectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelete_sectionContext)
}

func (s *UpdateContext) AllRemove_section() []IRemove_sectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRemove_sectionContext); ok {
			len++
		}
	}

	tst := make([]IRemove_sectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRemove_sectionContext); ok {
			tst[i] = t.(IRemove_sectionContext)
			i++
		}
	}

	return tst
}

func (s *UpdateContext) Remove_section(i int) IRemove_sectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemove_sectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemove_sectionContext)
}

func (s *UpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterUpdate(s)
	}
}

func (s *UpdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitUpdate(s)
	}
}

func (p *DynamoDbGrammarParser) Update() (localctx IUpdateContext) {
	localctx = NewUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DynamoDbGrammarParserRULE_update)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62914560) != 0) {
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case DynamoDbGrammarParserSET:
			{
				p.SetState(228)
				p.Set_section()
			}

		case DynamoDbGrammarParserADD:
			{
				p.SetState(229)
				p.Add_section()
			}

		case DynamoDbGrammarParserDELETE:
			{
				p.SetState(230)
				p.Delete_section()
			}

		case DynamoDbGrammarParserREMOVE:
			{
				p.SetState(231)
				p.Remove_section()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISet_sectionContext is an interface to support dynamic dispatch.
type ISet_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET() antlr.TerminalNode
	AllSet_action() []ISet_actionContext
	Set_action(i int) ISet_actionContext

	// IsSet_sectionContext differentiates from other interfaces.
	IsSet_sectionContext()
}

type Set_sectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_sectionContext() *Set_sectionContext {
	var p = new(Set_sectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_set_section
	return p
}

func InitEmptySet_sectionContext(p *Set_sectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_set_section
}

func (*Set_sectionContext) IsSet_sectionContext() {}

func NewSet_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_sectionContext {
	var p = new(Set_sectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_set_section

	return p
}

func (s *Set_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_sectionContext) SET() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSET, 0)
}

func (s *Set_sectionContext) AllSet_action() []ISet_actionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISet_actionContext); ok {
			len++
		}
	}

	tst := make([]ISet_actionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISet_actionContext); ok {
			tst[i] = t.(ISet_actionContext)
			i++
		}
	}

	return tst
}

func (s *Set_sectionContext) Set_action(i int) ISet_actionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_actionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_actionContext)
}

func (s *Set_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSet_section(s)
	}
}

func (s *Set_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSet_section(s)
	}
}

func (p *DynamoDbGrammarParser) Set_section() (localctx ISet_sectionContext) {
	localctx = NewSet_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DynamoDbGrammarParserRULE_set_section)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(DynamoDbGrammarParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Set_action()
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(238)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Set_action()
		}

		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISet_actionContext is an interface to support dynamic dispatch.
type ISet_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	EQ() antlr.TerminalNode
	Set_value() ISet_valueContext

	// IsSet_actionContext differentiates from other interfaces.
	IsSet_actionContext()
}

type Set_actionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_actionContext() *Set_actionContext {
	var p = new(Set_actionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_set_action
	return p
}

func InitEmptySet_actionContext(p *Set_actionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_set_action
}

func (*Set_actionContext) IsSet_actionContext() {}

func NewSet_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_actionContext {
	var p = new(Set_actionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_set_action

	return p
}

func (s *Set_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_actionContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Set_actionContext) EQ() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEQ, 0)
}

func (s *Set_actionContext) Set_value() ISet_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_valueContext)
}

func (s *Set_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSet_action(s)
	}
}

func (s *Set_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSet_action(s)
	}
}

func (p *DynamoDbGrammarParser) Set_action() (localctx ISet_actionContext) {
	localctx = NewSet_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DynamoDbGrammarParserRULE_set_action)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Path()
	}
	{
		p.SetState(246)
		p.Match(DynamoDbGrammarParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Set_value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdd_sectionContext is an interface to support dynamic dispatch.
type IAdd_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	AllAdd_action() []IAdd_actionContext
	Add_action(i int) IAdd_actionContext

	// IsAdd_sectionContext differentiates from other interfaces.
	IsAdd_sectionContext()
}

type Add_sectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdd_sectionContext() *Add_sectionContext {
	var p = new(Add_sectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_add_section
	return p
}

func InitEmptyAdd_sectionContext(p *Add_sectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_add_section
}

func (*Add_sectionContext) IsAdd_sectionContext() {}

func NewAdd_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Add_sectionContext {
	var p = new(Add_sectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_add_section

	return p
}

func (s *Add_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Add_sectionContext) ADD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserADD, 0)
}

func (s *Add_sectionContext) AllAdd_action() []IAdd_actionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdd_actionContext); ok {
			len++
		}
	}

	tst := make([]IAdd_actionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdd_actionContext); ok {
			tst[i] = t.(IAdd_actionContext)
			i++
		}
	}

	return tst
}

func (s *Add_sectionContext) Add_action(i int) IAdd_actionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdd_actionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdd_actionContext)
}

func (s *Add_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Add_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Add_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAdd_section(s)
	}
}

func (s *Add_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAdd_section(s)
	}
}

func (p *DynamoDbGrammarParser) Add_section() (localctx IAdd_sectionContext) {
	localctx = NewAdd_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DynamoDbGrammarParserRULE_add_section)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(DynamoDbGrammarParserADD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Add_action()
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(251)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.Add_action()
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdd_actionContext is an interface to support dynamic dispatch.
type IAdd_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Literal() ILiteralContext

	// IsAdd_actionContext differentiates from other interfaces.
	IsAdd_actionContext()
}

type Add_actionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdd_actionContext() *Add_actionContext {
	var p = new(Add_actionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_add_action
	return p
}

func InitEmptyAdd_actionContext(p *Add_actionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_add_action
}

func (*Add_actionContext) IsAdd_actionContext() {}

func NewAdd_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Add_actionContext {
	var p = new(Add_actionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_add_action

	return p
}

func (s *Add_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Add_actionContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Add_actionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Add_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Add_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Add_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAdd_action(s)
	}
}

func (s *Add_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAdd_action(s)
	}
}

func (p *DynamoDbGrammarParser) Add_action() (localctx IAdd_actionContext) {
	localctx = NewAdd_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DynamoDbGrammarParserRULE_add_action)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Path()
	}
	{
		p.SetState(259)
		p.Literal()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDelete_sectionContext is an interface to support dynamic dispatch.
type IDelete_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE() antlr.TerminalNode
	AllDelete_action() []IDelete_actionContext
	Delete_action(i int) IDelete_actionContext

	// IsDelete_sectionContext differentiates from other interfaces.
	IsDelete_sectionContext()
}

type Delete_sectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_sectionContext() *Delete_sectionContext {
	var p = new(Delete_sectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_section
	return p
}

func InitEmptyDelete_sectionContext(p *Delete_sectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_section
}

func (*Delete_sectionContext) IsDelete_sectionContext() {}

func NewDelete_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_sectionContext {
	var p = new(Delete_sectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_section

	return p
}

func (s *Delete_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_sectionContext) DELETE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDELETE, 0)
}

func (s *Delete_sectionContext) AllDelete_action() []IDelete_actionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDelete_actionContext); ok {
			len++
		}
	}

	tst := make([]IDelete_actionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDelete_actionContext); ok {
			tst[i] = t.(IDelete_actionContext)
			i++
		}
	}

	return tst
}

func (s *Delete_sectionContext) Delete_action(i int) IDelete_actionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelete_actionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelete_actionContext)
}

func (s *Delete_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDelete_section(s)
	}
}

func (s *Delete_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDelete_section(s)
	}
}

func (p *DynamoDbGrammarParser) Delete_section() (localctx IDelete_sectionContext) {
	localctx = NewDelete_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DynamoDbGrammarParserRULE_delete_section)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(DynamoDbGrammarParserDELETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Delete_action()
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(263)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Delete_action()
		}

		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDelete_actionContext is an interface to support dynamic dispatch.
type IDelete_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Literal() ILiteralContext

	// IsDelete_actionContext differentiates from other interfaces.
	IsDelete_actionContext()
}

type Delete_actionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelete_actionContext() *Delete_actionContext {
	var p = new(Delete_actionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_action
	return p
}

func InitEmptyDelete_actionContext(p *Delete_actionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_action
}

func (*Delete_actionContext) IsDelete_actionContext() {}

func NewDelete_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delete_actionContext {
	var p = new(Delete_actionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_delete_action

	return p
}

func (s *Delete_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Delete_actionContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Delete_actionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Delete_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delete_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delete_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDelete_action(s)
	}
}

func (s *Delete_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDelete_action(s)
	}
}

func (p *DynamoDbGrammarParser) Delete_action() (localctx IDelete_actionContext) {
	localctx = NewDelete_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DynamoDbGrammarParserRULE_delete_action)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Path()
	}
	{
		p.SetState(271)
		p.Literal()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRemove_sectionContext is an interface to support dynamic dispatch.
type IRemove_sectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REMOVE() antlr.TerminalNode
	AllRemove_action() []IRemove_actionContext
	Remove_action(i int) IRemove_actionContext

	// IsRemove_sectionContext differentiates from other interfaces.
	IsRemove_sectionContext()
}

type Remove_sectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemove_sectionContext() *Remove_sectionContext {
	var p = new(Remove_sectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_section
	return p
}

func InitEmptyRemove_sectionContext(p *Remove_sectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_section
}

func (*Remove_sectionContext) IsRemove_sectionContext() {}

func NewRemove_sectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Remove_sectionContext {
	var p = new(Remove_sectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_section

	return p
}

func (s *Remove_sectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Remove_sectionContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserREMOVE, 0)
}

func (s *Remove_sectionContext) AllRemove_action() []IRemove_actionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRemove_actionContext); ok {
			len++
		}
	}

	tst := make([]IRemove_actionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRemove_actionContext); ok {
			tst[i] = t.(IRemove_actionContext)
			i++
		}
	}

	return tst
}

func (s *Remove_sectionContext) Remove_action(i int) IRemove_actionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemove_actionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemove_actionContext)
}

func (s *Remove_sectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_sectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Remove_sectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterRemove_section(s)
	}
}

func (s *Remove_sectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitRemove_section(s)
	}
}

func (p *DynamoDbGrammarParser) Remove_section() (localctx IRemove_sectionContext) {
	localctx = NewRemove_sectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DynamoDbGrammarParserRULE_remove_section)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(DynamoDbGrammarParserREMOVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Remove_action()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(275)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Remove_action()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRemove_actionContext is an interface to support dynamic dispatch.
type IRemove_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext

	// IsRemove_actionContext differentiates from other interfaces.
	IsRemove_actionContext()
}

type Remove_actionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemove_actionContext() *Remove_actionContext {
	var p = new(Remove_actionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_action
	return p
}

func InitEmptyRemove_actionContext(p *Remove_actionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_action
}

func (*Remove_actionContext) IsRemove_actionContext() {}

func NewRemove_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Remove_actionContext {
	var p = new(Remove_actionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_remove_action

	return p
}

func (s *Remove_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Remove_actionContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Remove_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Remove_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterRemove_action(s)
	}
}

func (s *Remove_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitRemove_action(s)
	}
}

func (p *DynamoDbGrammarParser) Remove_action() (localctx IRemove_actionContext) {
	localctx = NewRemove_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DynamoDbGrammarParserRULE_remove_action)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Path()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISet_valueContext is an interface to support dynamic dispatch.
type ISet_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSet_valueContext differentiates from other interfaces.
	IsSet_valueContext()
}

type Set_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_valueContext() *Set_valueContext {
	var p = new(Set_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_set_value
	return p
}

func InitEmptySet_valueContext(p *Set_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_set_value
}

func (*Set_valueContext) IsSet_valueContext() {}

func NewSet_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_valueContext {
	var p = new(Set_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_set_value

	return p
}

func (s *Set_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_valueContext) CopyAll(ctx *Set_valueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Set_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArithmeticValueContext struct {
	Set_valueContext
}

func NewArithmeticValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticValueContext {
	var p = new(ArithmeticValueContext)

	InitEmptySet_valueContext(&p.Set_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Set_valueContext))

	return p
}

func (s *ArithmeticValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticValueContext) Arithmetic() IArithmeticContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}

func (s *ArithmeticValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterArithmeticValue(s)
	}
}

func (s *ArithmeticValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitArithmeticValue(s)
	}
}

type OperandValueContext struct {
	Set_valueContext
}

func NewOperandValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperandValueContext {
	var p = new(OperandValueContext)

	InitEmptySet_valueContext(&p.Set_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Set_valueContext))

	return p
}

func (s *OperandValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandValueContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *OperandValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOperandValue(s)
	}
}

func (s *OperandValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOperandValue(s)
	}
}

func (p *DynamoDbGrammarParser) Set_value() (localctx ISet_valueContext) {
	localctx = NewSet_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DynamoDbGrammarParserRULE_set_value)
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOperandValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Operand()
		}

	case 2:
		localctx = NewArithmeticValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Arithmetic()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArithmeticContext is an interface to support dynamic dispatch.
type IArithmeticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHasOuterParens returns the hasOuterParens attribute.
	GetHasOuterParens() bool

	// SetHasOuterParens sets the hasOuterParens attribute.
	SetHasOuterParens(bool)

	// IsArithmeticContext differentiates from other interfaces.
	IsArithmeticContext()
}

type ArithmeticContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	hasOuterParens bool // TODO = false
}

func NewEmptyArithmeticContext() *ArithmeticContext {
	var p = new(ArithmeticContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_arithmetic
	return p
}

func InitEmptyArithmeticContext(p *ArithmeticContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_arithmetic
}

func (*ArithmeticContext) IsArithmeticContext() {}

func NewArithmeticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticContext {
	var p = new(ArithmeticContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_arithmetic

	return p
}

func (s *ArithmeticContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticContext) GetHasOuterParens() bool { return s.hasOuterParens }

func (s *ArithmeticContext) SetHasOuterParens(v bool) { s.hasOuterParens = v }

func (s *ArithmeticContext) CopyAll(ctx *ArithmeticContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
	s.hasOuterParens = ctx.hasOuterParens
}

func (s *ArithmeticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PlusMinusContext struct {
	ArithmeticContext
}

func NewPlusMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusMinusContext {
	var p = new(PlusMinusContext)

	InitEmptyArithmeticContext(&p.ArithmeticContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArithmeticContext))

	return p
}

func (s *PlusMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusMinusContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *PlusMinusContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PlusMinusContext) PLUS() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserPLUS, 0)
}

func (s *PlusMinusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserMINUS, 0)
}

func (s *PlusMinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterPlusMinus(s)
	}
}

func (s *PlusMinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitPlusMinus(s)
	}
}

type ArithmeticParensContext struct {
	ArithmeticContext
	a IArithmeticContext
}

func NewArithmeticParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArithmeticParensContext {
	var p = new(ArithmeticParensContext)

	InitEmptyArithmeticContext(&p.ArithmeticContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArithmeticContext))

	return p
}

func (s *ArithmeticParensContext) GetA() IArithmeticContext { return s.a }

func (s *ArithmeticParensContext) SetA(v IArithmeticContext) { s.a = v }

func (s *ArithmeticParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticParensContext) Arithmetic() IArithmeticContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticContext)
}

func (s *ArithmeticParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterArithmeticParens(s)
	}
}

func (s *ArithmeticParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitArithmeticParens(s)
	}
}

func (p *DynamoDbGrammarParser) Arithmetic() (localctx IArithmeticContext) {
	localctx = NewArithmeticContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DynamoDbGrammarParserRULE_arithmetic)
	var _la int

	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPlusMinusContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.Operand()
		}
		{
			p.SetState(289)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DynamoDbGrammarParserPLUS || _la == DynamoDbGrammarParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(290)
			p.Operand()
		}

	case 2:
		localctx = NewArithmeticParensContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(DynamoDbGrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)

			var _x = p.Arithmetic()

			localctx.(*ArithmeticParensContext).a = _x
		}
		{
			p.SetState(294)
			p.Match(DynamoDbGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		// validateRedundantParentheses(localctx.(*ArithmeticParensContext).GetA().GetHasOuterParens());
		localctx.(*ArithmeticParensContext).SetHasOuterParens(true)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHasOuterParens returns the hasOuterParens attribute.
	GetHasOuterParens() bool

	// SetHasOuterParens sets the hasOuterParens attribute.
	SetHasOuterParens(bool)

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	hasOuterParens bool // TODO = false
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) GetHasOuterParens() bool { return s.hasOuterParens }

func (s *OperandContext) SetHasOuterParens(v bool) { s.hasOuterParens = v }

func (s *OperandContext) CopyAll(ctx *OperandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
	s.hasOuterParens = ctx.hasOuterParens
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PathOperandContext struct {
	OperandContext
}

func NewPathOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PathOperandContext {
	var p = new(PathOperandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *PathOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathOperandContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *PathOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterPathOperand(s)
	}
}

func (s *PathOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitPathOperand(s)
	}
}

type LiteralOperandContext struct {
	OperandContext
}

func NewLiteralOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralOperandContext {
	var p = new(LiteralOperandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *LiteralOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralOperandContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterLiteralOperand(s)
	}
}

func (s *LiteralOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitLiteralOperand(s)
	}
}

type FunctionOperandContext struct {
	OperandContext
}

func NewFunctionOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionOperandContext {
	var p = new(FunctionOperandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *FunctionOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionOperandContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterFunctionOperand(s)
	}
}

func (s *FunctionOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitFunctionOperand(s)
	}
}

type ParenOperandContext struct {
	OperandContext
	o IOperandContext
}

func NewParenOperandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenOperandContext {
	var p = new(ParenOperandContext)

	InitEmptyOperandContext(&p.OperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*OperandContext))

	return p
}

func (s *ParenOperandContext) GetO() IOperandContext { return s.o }

func (s *ParenOperandContext) SetO(v IOperandContext) { s.o = v }

func (s *ParenOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenOperandContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *ParenOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterParenOperand(s)
	}
}

func (s *ParenOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitParenOperand(s)
	}
}

func (p *DynamoDbGrammarParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DynamoDbGrammarParserRULE_operand)
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPathOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(299)
			p.Path()
		}

	case 2:
		localctx = NewLiteralOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.Literal()
		}

	case 3:
		localctx = NewFunctionOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(301)
			p.Function()
		}

	case 4:
		localctx = NewParenOperandContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(302)
			p.Match(DynamoDbGrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)

			var _x = p.Operand()

			localctx.(*ParenOperandContext).o = _x
		}
		{
			p.SetState(304)
			p.Match(DynamoDbGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		// validateRedundantParentheses(localctx.(*ParenOperandContext).GetO().GetHasOuterParens());
		localctx.(*ParenOperandContext).SetHasOuterParens(true)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) CopyAll(ctx *FunctionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionCallContext struct {
	FunctionContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyFunctionContext(&p.FunctionContext)
	p.parser = parser
	p.CopyAll(ctx.(*FunctionContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserID, 0)
}

func (s *FunctionCallContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *FunctionCallContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *DynamoDbGrammarParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DynamoDbGrammarParserRULE_function)
	var _la int

	localctx = NewFunctionCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(DynamoDbGrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Match(DynamoDbGrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Operand()
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(312)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Operand()
		}

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(319)
		p.Match(DynamoDbGrammarParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id() IIdContext
	AllDereference() []IDereferenceContext
	Dereference(i int) IDereferenceContext

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_path
	return p
}

func InitEmptyPathContext(p *PathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_path
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *PathContext) AllDereference() []IDereferenceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDereferenceContext); ok {
			len++
		}
	}

	tst := make([]IDereferenceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDereferenceContext); ok {
			tst[i] = t.(IDereferenceContext)
			i++
		}
	}

	return tst
}

func (s *PathContext) Dereference(i int) IDereferenceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDereferenceContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDereferenceContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitPath(s)
	}
}

func (p *DynamoDbGrammarParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DynamoDbGrammarParserRULE_path)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Id()
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(322)
				p.Dereference()
			}

		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdContext is an interface to support dynamic dispatch.
type IIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ATTRIBUTE_NAME_SUB() antlr.TerminalNode

	// IsIdContext differentiates from other interfaces.
	IsIdContext()
}

type IdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdContext() *IdContext {
	var p = new(IdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_id
	return p
}

func InitEmptyIdContext(p *IdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_id
}

func (*IdContext) IsIdContext() {}

func NewIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdContext {
	var p = new(IdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_id

	return p
}

func (s *IdContext) GetParser() antlr.Parser { return s.parser }

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserID, 0)
}

func (s *IdContext) ATTRIBUTE_NAME_SUB() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserATTRIBUTE_NAME_SUB, 0)
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitId(s)
	}
}

func (p *DynamoDbGrammarParser) Id() (localctx IIdContext) {
	localctx = NewIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DynamoDbGrammarParserRULE_id)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DynamoDbGrammarParserID || _la == DynamoDbGrammarParserATTRIBUTE_NAME_SUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDereferenceContext is an interface to support dynamic dispatch.
type IDereferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDereferenceContext differentiates from other interfaces.
	IsDereferenceContext()
}

type DereferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDereferenceContext() *DereferenceContext {
	var p = new(DereferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_dereference
	return p
}

func InitEmptyDereferenceContext(p *DereferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_dereference
}

func (*DereferenceContext) IsDereferenceContext() {}

func NewDereferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DereferenceContext {
	var p = new(DereferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_dereference

	return p
}

func (s *DereferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *DereferenceContext) CopyAll(ctx *DereferenceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DereferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DereferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListAccessContext struct {
	DereferenceContext
}

func NewListAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListAccessContext {
	var p = new(ListAccessContext)

	InitEmptyDereferenceContext(&p.DereferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*DereferenceContext))

	return p
}

func (s *ListAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAccessContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *ListAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterListAccess(s)
	}
}

func (s *ListAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitListAccess(s)
	}
}

type MapAccessContext struct {
	DereferenceContext
}

func NewMapAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapAccessContext {
	var p = new(MapAccessContext)

	InitEmptyDereferenceContext(&p.DereferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*DereferenceContext))

	return p
}

func (s *MapAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapAccessContext) Id() IIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdContext)
}

func (s *MapAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterMapAccess(s)
	}
}

func (s *MapAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitMapAccess(s)
	}
}

func (p *DynamoDbGrammarParser) Dereference() (localctx IDereferenceContext) {
	localctx = NewDereferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DynamoDbGrammarParserRULE_dereference)
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserT__3:
		localctx = NewMapAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Match(DynamoDbGrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Id()
		}

	case DynamoDbGrammarParserT__4:
		localctx = NewListAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(332)
			p.Match(DynamoDbGrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(DynamoDbGrammarParserINDEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(DynamoDbGrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralSubContext struct {
	LiteralContext
}

func NewLiteralSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralSubContext {
	var p = new(LiteralSubContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LiteralSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralSubContext) LITERAL_SUB() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserLITERAL_SUB, 0)
}

func (s *LiteralSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterLiteralSub(s)
	}
}

func (s *LiteralSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitLiteralSub(s)
	}
}

func (p *DynamoDbGrammarParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DynamoDbGrammarParserRULE_literal)
	localctx = NewLiteralSubContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(DynamoDbGrammarParserLITERAL_SUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpression_attr_names_subContext is an interface to support dynamic dispatch.
type IExpression_attr_names_subContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRIBUTE_NAME_SUB() antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsExpression_attr_names_subContext differentiates from other interfaces.
	IsExpression_attr_names_subContext()
}

type Expression_attr_names_subContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_attr_names_subContext() *Expression_attr_names_subContext {
	var p = new(Expression_attr_names_subContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_names_sub
	return p
}

func InitEmptyExpression_attr_names_subContext(p *Expression_attr_names_subContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_names_sub
}

func (*Expression_attr_names_subContext) IsExpression_attr_names_subContext() {}

func NewExpression_attr_names_subContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_attr_names_subContext {
	var p = new(Expression_attr_names_subContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_names_sub

	return p
}

func (s *Expression_attr_names_subContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_attr_names_subContext) ATTRIBUTE_NAME_SUB() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserATTRIBUTE_NAME_SUB, 0)
}

func (s *Expression_attr_names_subContext) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Expression_attr_names_subContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_attr_names_subContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_attr_names_subContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterExpression_attr_names_sub(s)
	}
}

func (s *Expression_attr_names_subContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitExpression_attr_names_sub(s)
	}
}

func (p *DynamoDbGrammarParser) Expression_attr_names_sub() (localctx IExpression_attr_names_subContext) {
	localctx = NewExpression_attr_names_subContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DynamoDbGrammarParserRULE_expression_attr_names_sub)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(DynamoDbGrammarParserATTRIBUTE_NAME_SUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Match(DynamoDbGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpression_attr_values_subContext is an interface to support dynamic dispatch.
type IExpression_attr_values_subContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LITERAL_SUB() antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsExpression_attr_values_subContext differentiates from other interfaces.
	IsExpression_attr_values_subContext()
}

type Expression_attr_values_subContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_attr_values_subContext() *Expression_attr_values_subContext {
	var p = new(Expression_attr_values_subContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_values_sub
	return p
}

func InitEmptyExpression_attr_values_subContext(p *Expression_attr_values_subContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_values_sub
}

func (*Expression_attr_values_subContext) IsExpression_attr_values_subContext() {}

func NewExpression_attr_values_subContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_attr_values_subContext {
	var p = new(Expression_attr_values_subContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_expression_attr_values_sub

	return p
}

func (s *Expression_attr_values_subContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_attr_values_subContext) LITERAL_SUB() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserLITERAL_SUB, 0)
}

func (s *Expression_attr_values_subContext) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Expression_attr_values_subContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_attr_values_subContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_attr_values_subContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterExpression_attr_values_sub(s)
	}
}

func (s *Expression_attr_values_subContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitExpression_attr_values_sub(s)
	}
}

func (p *DynamoDbGrammarParser) Expression_attr_values_sub() (localctx IExpression_attr_values_subContext) {
	localctx = NewExpression_attr_values_subContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DynamoDbGrammarParserRULE_expression_attr_values_sub)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(DynamoDbGrammarParserLITERAL_SUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(DynamoDbGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatement_Context is an interface to support dynamic dispatch.
type IStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement() IStatementContext
	EOF() antlr.TerminalNode

	// IsStatement_Context differentiates from other interfaces.
	IsStatement_Context()
}

type Statement_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatement_Context() *Statement_Context {
	var p = new(Statement_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_statement_
	return p
}

func InitEmptyStatement_Context(p *Statement_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_statement_
}

func (*Statement_Context) IsStatement_Context() {}

func NewStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_Context {
	var p = new(Statement_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_statement_

	return p
}

func (s *Statement_Context) GetParser() antlr.Parser { return s.parser }

func (s *Statement_Context) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Statement_Context) EOF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEOF, 0)
}

func (s *Statement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statement_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterStatement_(s)
	}
}

func (s *Statement_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitStatement_(s)
	}
}

func (p *DynamoDbGrammarParser) Statement_() (localctx IStatement_Context) {
	localctx = NewStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DynamoDbGrammarParserRULE_statement_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Statement()
	}
	{
		p.SetState(346)
		p.Match(DynamoDbGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InsertStatement() IInsertStatementContext
	CreateTableStatement() ICreateTableStatementContext
	ShowTablesStatement() IShowTablesStatementContext
	UpdateStatement() IUpdateStatementContext
	DeleteStatement() IDeleteStatementContext
	SelectStatement() ISelectStatementContext
	DropTableStatement() IDropTableStatementContext
	AlterTableStatement() IAlterTableStatementContext
	DescribeStatement() IDescribeStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) InsertStatement() IInsertStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertStatementContext)
}

func (s *StatementContext) CreateTableStatement() ICreateTableStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateTableStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateTableStatementContext)
}

func (s *StatementContext) ShowTablesStatement() IShowTablesStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShowTablesStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShowTablesStatementContext)
}

func (s *StatementContext) UpdateStatement() IUpdateStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateStatementContext)
}

func (s *StatementContext) DeleteStatement() IDeleteStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteStatementContext)
}

func (s *StatementContext) SelectStatement() ISelectStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *StatementContext) DropTableStatement() IDropTableStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropTableStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropTableStatementContext)
}

func (s *StatementContext) AlterTableStatement() IAlterTableStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlterTableStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlterTableStatementContext)
}

func (s *StatementContext) DescribeStatement() IDescribeStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescribeStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescribeStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *DynamoDbGrammarParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DynamoDbGrammarParserRULE_statement)
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserINSERT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.InsertStatement()
		}

	case DynamoDbGrammarParserCREATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.CreateTableStatement()
		}

	case DynamoDbGrammarParserSHOW:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(350)
			p.ShowTablesStatement()
		}

	case DynamoDbGrammarParserUPDATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(351)
			p.UpdateStatement()
		}

	case DynamoDbGrammarParserDELETE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(352)
			p.DeleteStatement()
		}

	case DynamoDbGrammarParserSELECT:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(353)
			p.SelectStatement()
		}

	case DynamoDbGrammarParserDROP:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(354)
			p.DropTableStatement()
		}

	case DynamoDbGrammarParserALTER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(355)
			p.AlterTableStatement()
		}

	case DynamoDbGrammarParserDESCRIBE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(356)
			p.DescribeStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropTableStatementContext is an interface to support dynamic dispatch.
type IDropTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DROP() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	TableName() ITableNameContext

	// IsDropTableStatementContext differentiates from other interfaces.
	IsDropTableStatementContext()
}

type DropTableStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropTableStatementContext() *DropTableStatementContext {
	var p = new(DropTableStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_dropTableStatement
	return p
}

func InitEmptyDropTableStatementContext(p *DropTableStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_dropTableStatement
}

func (*DropTableStatementContext) IsDropTableStatementContext() {}

func NewDropTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropTableStatementContext {
	var p = new(DropTableStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_dropTableStatement

	return p
}

func (s *DropTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DropTableStatementContext) DROP() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDROP, 0)
}

func (s *DropTableStatementContext) TABLE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserTABLE, 0)
}

func (s *DropTableStatementContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *DropTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropTableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDropTableStatement(s)
	}
}

func (s *DropTableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDropTableStatement(s)
	}
}

func (p *DynamoDbGrammarParser) DropTableStatement() (localctx IDropTableStatementContext) {
	localctx = NewDropTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DynamoDbGrammarParserRULE_dropTableStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(DynamoDbGrammarParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(360)
		p.Match(DynamoDbGrammarParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.TableName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDescribeStatementContext is an interface to support dynamic dispatch.
type IDescribeStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DESCRIBE() antlr.TerminalNode
	TableName() ITableNameContext

	// IsDescribeStatementContext differentiates from other interfaces.
	IsDescribeStatementContext()
}

type DescribeStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescribeStatementContext() *DescribeStatementContext {
	var p = new(DescribeStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_describeStatement
	return p
}

func InitEmptyDescribeStatementContext(p *DescribeStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_describeStatement
}

func (*DescribeStatementContext) IsDescribeStatementContext() {}

func NewDescribeStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescribeStatementContext {
	var p = new(DescribeStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_describeStatement

	return p
}

func (s *DescribeStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DescribeStatementContext) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDESCRIBE, 0)
}

func (s *DescribeStatementContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *DescribeStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescribeStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescribeStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDescribeStatement(s)
	}
}

func (s *DescribeStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDescribeStatement(s)
	}
}

func (p *DynamoDbGrammarParser) DescribeStatement() (localctx IDescribeStatementContext) {
	localctx = NewDescribeStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DynamoDbGrammarParserRULE_describeStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(DynamoDbGrammarParserDESCRIBE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.TableName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAlterTableStatementContext is an interface to support dynamic dispatch.
type IAlterTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALTER() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	TableName() ITableNameContext
	AlterTableStatementType() IAlterTableStatementTypeContext

	// IsAlterTableStatementContext differentiates from other interfaces.
	IsAlterTableStatementContext()
}

type AlterTableStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterTableStatementContext() *AlterTableStatementContext {
	var p = new(AlterTableStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatement
	return p
}

func InitEmptyAlterTableStatementContext(p *AlterTableStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatement
}

func (*AlterTableStatementContext) IsAlterTableStatementContext() {}

func NewAlterTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterTableStatementContext {
	var p = new(AlterTableStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatement

	return p
}

func (s *AlterTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterTableStatementContext) ALTER() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserALTER, 0)
}

func (s *AlterTableStatementContext) TABLE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserTABLE, 0)
}

func (s *AlterTableStatementContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *AlterTableStatementContext) AlterTableStatementType() IAlterTableStatementTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlterTableStatementTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlterTableStatementTypeContext)
}

func (s *AlterTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterTableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAlterTableStatement(s)
	}
}

func (s *AlterTableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAlterTableStatement(s)
	}
}

func (p *DynamoDbGrammarParser) AlterTableStatement() (localctx IAlterTableStatementContext) {
	localctx = NewAlterTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DynamoDbGrammarParserRULE_alterTableStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(DynamoDbGrammarParserALTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Match(DynamoDbGrammarParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(368)
		p.TableName()
	}
	{
		p.SetState(369)
		p.AlterTableStatementType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAlterTableStatementTypeContext is an interface to support dynamic dispatch.
type IAlterTableStatementTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SetCapacity() ISetCapacityContext
	AddIndex() IAddIndexContext
	DropIndex() IDropIndexContext
	AlterIndex() IAlterIndexContext

	// IsAlterTableStatementTypeContext differentiates from other interfaces.
	IsAlterTableStatementTypeContext()
}

type AlterTableStatementTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterTableStatementTypeContext() *AlterTableStatementTypeContext {
	var p = new(AlterTableStatementTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatementType
	return p
}

func InitEmptyAlterTableStatementTypeContext(p *AlterTableStatementTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatementType
}

func (*AlterTableStatementTypeContext) IsAlterTableStatementTypeContext() {}

func NewAlterTableStatementTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterTableStatementTypeContext {
	var p = new(AlterTableStatementTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_alterTableStatementType

	return p
}

func (s *AlterTableStatementTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterTableStatementTypeContext) SetCapacity() ISetCapacityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetCapacityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetCapacityContext)
}

func (s *AlterTableStatementTypeContext) AddIndex() IAddIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddIndexContext)
}

func (s *AlterTableStatementTypeContext) DropIndex() IDropIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropIndexContext)
}

func (s *AlterTableStatementTypeContext) AlterIndex() IAlterIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlterIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlterIndexContext)
}

func (s *AlterTableStatementTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterTableStatementTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterTableStatementTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAlterTableStatementType(s)
	}
}

func (s *AlterTableStatementTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAlterTableStatementType(s)
	}
}

func (p *DynamoDbGrammarParser) AlterTableStatementType() (localctx IAlterTableStatementTypeContext) {
	localctx = NewAlterTableStatementTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DynamoDbGrammarParserRULE_alterTableStatementType)
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserSET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.SetCapacity()
		}

	case DynamoDbGrammarParserADD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.AddIndex()
		}

	case DynamoDbGrammarParserDROP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.DropIndex()
		}

	case DynamoDbGrammarParserALTER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
			p.AlterIndex()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISetCapacityContext is an interface to support dynamic dispatch.
type ISetCapacityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET() antlr.TerminalNode
	Capacity() ICapacityContext

	// IsSetCapacityContext differentiates from other interfaces.
	IsSetCapacityContext()
}

type SetCapacityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetCapacityContext() *SetCapacityContext {
	var p = new(SetCapacityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_setCapacity
	return p
}

func InitEmptySetCapacityContext(p *SetCapacityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_setCapacity
}

func (*SetCapacityContext) IsSetCapacityContext() {}

func NewSetCapacityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetCapacityContext {
	var p = new(SetCapacityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_setCapacity

	return p
}

func (s *SetCapacityContext) GetParser() antlr.Parser { return s.parser }

func (s *SetCapacityContext) SET() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSET, 0)
}

func (s *SetCapacityContext) Capacity() ICapacityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapacityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapacityContext)
}

func (s *SetCapacityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetCapacityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetCapacityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSetCapacity(s)
	}
}

func (s *SetCapacityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSetCapacity(s)
	}
}

func (p *DynamoDbGrammarParser) SetCapacity() (localctx ISetCapacityContext) {
	localctx = NewSetCapacityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DynamoDbGrammarParserRULE_setCapacity)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Match(DynamoDbGrammarParserSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.Capacity()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddIndexContext is an interface to support dynamic dispatch.
type IAddIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SecondaryIndexDecl() ISecondaryIndexDeclContext

	// IsAddIndexContext differentiates from other interfaces.
	IsAddIndexContext()
}

type AddIndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddIndexContext() *AddIndexContext {
	var p = new(AddIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_addIndex
	return p
}

func InitEmptyAddIndexContext(p *AddIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_addIndex
}

func (*AddIndexContext) IsAddIndexContext() {}

func NewAddIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddIndexContext {
	var p = new(AddIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_addIndex

	return p
}

func (s *AddIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *AddIndexContext) ADD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserADD, 0)
}

func (s *AddIndexContext) SecondaryIndexDecl() ISecondaryIndexDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISecondaryIndexDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISecondaryIndexDeclContext)
}

func (s *AddIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAddIndex(s)
	}
}

func (s *AddIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAddIndex(s)
	}
}

func (p *DynamoDbGrammarParser) AddIndex() (localctx IAddIndexContext) {
	localctx = NewAddIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, DynamoDbGrammarParserRULE_addIndex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(DynamoDbGrammarParserADD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.SecondaryIndexDecl()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDropIndexContext is an interface to support dynamic dispatch.
type IDropIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DROP() antlr.TerminalNode
	INDEXKEYWORD() antlr.TerminalNode
	IndexName() IIndexNameContext

	// IsDropIndexContext differentiates from other interfaces.
	IsDropIndexContext()
}

type DropIndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropIndexContext() *DropIndexContext {
	var p = new(DropIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_dropIndex
	return p
}

func InitEmptyDropIndexContext(p *DropIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_dropIndex
}

func (*DropIndexContext) IsDropIndexContext() {}

func NewDropIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropIndexContext {
	var p = new(DropIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_dropIndex

	return p
}

func (s *DropIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *DropIndexContext) DROP() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDROP, 0)
}

func (s *DropIndexContext) INDEXKEYWORD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEXKEYWORD, 0)
}

func (s *DropIndexContext) IndexName() IIndexNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *DropIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDropIndex(s)
	}
}

func (s *DropIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDropIndex(s)
	}
}

func (p *DynamoDbGrammarParser) DropIndex() (localctx IDropIndexContext) {
	localctx = NewDropIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, DynamoDbGrammarParserRULE_dropIndex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(383)
		p.Match(DynamoDbGrammarParserDROP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.Match(DynamoDbGrammarParserINDEXKEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)
		p.IndexName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAlterIndexContext is an interface to support dynamic dispatch.
type IAlterIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALTER() antlr.TerminalNode
	INDEXKEYWORD() antlr.TerminalNode
	IndexName() IIndexNameContext
	SetCapacity() ISetCapacityContext

	// IsAlterIndexContext differentiates from other interfaces.
	IsAlterIndexContext()
}

type AlterIndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlterIndexContext() *AlterIndexContext {
	var p = new(AlterIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_alterIndex
	return p
}

func InitEmptyAlterIndexContext(p *AlterIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_alterIndex
}

func (*AlterIndexContext) IsAlterIndexContext() {}

func NewAlterIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlterIndexContext {
	var p = new(AlterIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_alterIndex

	return p
}

func (s *AlterIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *AlterIndexContext) ALTER() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserALTER, 0)
}

func (s *AlterIndexContext) INDEXKEYWORD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEXKEYWORD, 0)
}

func (s *AlterIndexContext) IndexName() IIndexNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *AlterIndexContext) SetCapacity() ISetCapacityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetCapacityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetCapacityContext)
}

func (s *AlterIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlterIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlterIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAlterIndex(s)
	}
}

func (s *AlterIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAlterIndex(s)
	}
}

func (p *DynamoDbGrammarParser) AlterIndex() (localctx IAlterIndexContext) {
	localctx = NewAlterIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, DynamoDbGrammarParserRULE_alterIndex)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(DynamoDbGrammarParserALTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(388)
		p.Match(DynamoDbGrammarParserINDEXKEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(389)
		p.IndexName()
	}
	{
		p.SetState(390)
		p.SetCapacity()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUpdateStatementContext is an interface to support dynamic dispatch.
type IUpdateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UPDATE() antlr.TerminalNode
	TableName() ITableNameContext
	Update() IUpdateContext
	Where() IWhereContext
	Returning() IReturningContext

	// IsUpdateStatementContext differentiates from other interfaces.
	IsUpdateStatementContext()
}

type UpdateStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateStatementContext() *UpdateStatementContext {
	var p = new(UpdateStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_updateStatement
	return p
}

func InitEmptyUpdateStatementContext(p *UpdateStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_updateStatement
}

func (*UpdateStatementContext) IsUpdateStatementContext() {}

func NewUpdateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateStatementContext {
	var p = new(UpdateStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_updateStatement

	return p
}

func (s *UpdateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateStatementContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUPDATE, 0)
}

func (s *UpdateStatementContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *UpdateStatementContext) Update() IUpdateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateContext)
}

func (s *UpdateStatementContext) Where() IWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *UpdateStatementContext) Returning() IReturningContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturningContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturningContext)
}

func (s *UpdateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterUpdateStatement(s)
	}
}

func (s *UpdateStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitUpdateStatement(s)
	}
}

func (p *DynamoDbGrammarParser) UpdateStatement() (localctx IUpdateStatementContext) {
	localctx = NewUpdateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, DynamoDbGrammarParserRULE_updateStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Match(DynamoDbGrammarParserUPDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.TableName()
	}
	{
		p.SetState(394)
		p.Update()
	}
	{
		p.SetState(395)
		p.Where()
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserRETURNING {
		{
			p.SetState(396)
			p.Returning()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeleteStatementContext is an interface to support dynamic dispatch.
type IDeleteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE() antlr.TerminalNode
	FROM() antlr.TerminalNode
	TableName() ITableNameContext
	Where() IWhereContext
	Returning() IReturningContext

	// IsDeleteStatementContext differentiates from other interfaces.
	IsDeleteStatementContext()
}

type DeleteStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStatementContext() *DeleteStatementContext {
	var p = new(DeleteStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_deleteStatement
	return p
}

func InitEmptyDeleteStatementContext(p *DeleteStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_deleteStatement
}

func (*DeleteStatementContext) IsDeleteStatementContext() {}

func NewDeleteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStatementContext {
	var p = new(DeleteStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_deleteStatement

	return p
}

func (s *DeleteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStatementContext) DELETE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDELETE, 0)
}

func (s *DeleteStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserFROM, 0)
}

func (s *DeleteStatementContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *DeleteStatementContext) Where() IWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *DeleteStatementContext) Returning() IReturningContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturningContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturningContext)
}

func (s *DeleteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDeleteStatement(s)
	}
}

func (p *DynamoDbGrammarParser) DeleteStatement() (localctx IDeleteStatementContext) {
	localctx = NewDeleteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, DynamoDbGrammarParserRULE_deleteStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(DynamoDbGrammarParserDELETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(400)
		p.Match(DynamoDbGrammarParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.TableName()
	}
	{
		p.SetState(402)
		p.Where()
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserRETURNING {
		{
			p.SetState(403)
			p.Returning()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInsertStatementContext is an interface to support dynamic dispatch.
type IInsertStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INSERT() antlr.TerminalNode
	TableName() ITableNameContext
	Literal() ILiteralContext
	INTO() antlr.TerminalNode
	OnDuplicateKeyUpdate() IOnDuplicateKeyUpdateContext
	Returning() IReturningContext
	VALUES() antlr.TerminalNode
	VALUE() antlr.TerminalNode

	// IsInsertStatementContext differentiates from other interfaces.
	IsInsertStatementContext()
}

type InsertStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertStatementContext() *InsertStatementContext {
	var p = new(InsertStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_insertStatement
	return p
}

func InitEmptyInsertStatementContext(p *InsertStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_insertStatement
}

func (*InsertStatementContext) IsInsertStatementContext() {}

func NewInsertStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertStatementContext {
	var p = new(InsertStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_insertStatement

	return p
}

func (s *InsertStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertStatementContext) INSERT() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINSERT, 0)
}

func (s *InsertStatementContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *InsertStatementContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *InsertStatementContext) INTO() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINTO, 0)
}

func (s *InsertStatementContext) OnDuplicateKeyUpdate() IOnDuplicateKeyUpdateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOnDuplicateKeyUpdateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOnDuplicateKeyUpdateContext)
}

func (s *InsertStatementContext) Returning() IReturningContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturningContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturningContext)
}

func (s *InsertStatementContext) VALUES() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserVALUES, 0)
}

func (s *InsertStatementContext) VALUE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserVALUE, 0)
}

func (s *InsertStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterInsertStatement(s)
	}
}

func (s *InsertStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitInsertStatement(s)
	}
}

func (p *DynamoDbGrammarParser) InsertStatement() (localctx IInsertStatementContext) {
	localctx = NewInsertStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, DynamoDbGrammarParserRULE_insertStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(DynamoDbGrammarParserINSERT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserINTO {
		{
			p.SetState(407)
			p.Match(DynamoDbGrammarParserINTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(410)
		p.TableName()
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserVALUE || _la == DynamoDbGrammarParserVALUES {
		{
			p.SetState(411)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DynamoDbGrammarParserVALUE || _la == DynamoDbGrammarParserVALUES) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(414)
		p.Literal()
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserON {
		{
			p.SetState(415)
			p.OnDuplicateKeyUpdate()
		}

	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserRETURNING {
		{
			p.SetState(418)
			p.Returning()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateTableStatementContext is an interface to support dynamic dispatch.
type ICreateTableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	TableName() ITableNameContext
	AllAttributeDecl() []IAttributeDeclContext
	AttributeDecl(i int) IAttributeDeclContext
	PrimaryKeyDecl() IPrimaryKeyDeclContext
	AllSecondaryIndexDecl() []ISecondaryIndexDeclContext
	SecondaryIndexDecl(i int) ISecondaryIndexDeclContext

	// IsCreateTableStatementContext differentiates from other interfaces.
	IsCreateTableStatementContext()
}

type CreateTableStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableStatementContext() *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_createTableStatement
	return p
}

func InitEmptyCreateTableStatementContext(p *CreateTableStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_createTableStatement
}

func (*CreateTableStatementContext) IsCreateTableStatementContext() {}

func NewCreateTableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableStatementContext {
	var p = new(CreateTableStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_createTableStatement

	return p
}

func (s *CreateTableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableStatementContext) CREATE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserCREATE, 0)
}

func (s *CreateTableStatementContext) TABLE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserTABLE, 0)
}

func (s *CreateTableStatementContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *CreateTableStatementContext) AllAttributeDecl() []IAttributeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeDeclContext); ok {
			len++
		}
	}

	tst := make([]IAttributeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeDeclContext); ok {
			tst[i] = t.(IAttributeDeclContext)
			i++
		}
	}

	return tst
}

func (s *CreateTableStatementContext) AttributeDecl(i int) IAttributeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeDeclContext)
}

func (s *CreateTableStatementContext) PrimaryKeyDecl() IPrimaryKeyDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryKeyDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyDeclContext)
}

func (s *CreateTableStatementContext) AllSecondaryIndexDecl() []ISecondaryIndexDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISecondaryIndexDeclContext); ok {
			len++
		}
	}

	tst := make([]ISecondaryIndexDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISecondaryIndexDeclContext); ok {
			tst[i] = t.(ISecondaryIndexDeclContext)
			i++
		}
	}

	return tst
}

func (s *CreateTableStatementContext) SecondaryIndexDecl(i int) ISecondaryIndexDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISecondaryIndexDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISecondaryIndexDeclContext)
}

func (s *CreateTableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterCreateTableStatement(s)
	}
}

func (s *CreateTableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitCreateTableStatement(s)
	}
}

func (p *DynamoDbGrammarParser) CreateTableStatement() (localctx ICreateTableStatementContext) {
	localctx = NewCreateTableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, DynamoDbGrammarParserRULE_createTableStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Match(DynamoDbGrammarParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(422)
		p.Match(DynamoDbGrammarParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(423)
		p.TableName()
	}
	{
		p.SetState(424)
		p.Match(DynamoDbGrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(425)
		p.AttributeDecl()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(426)
				p.Match(DynamoDbGrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(427)
				p.AttributeDecl()
			}

		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Match(DynamoDbGrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)
		p.PrimaryKeyDecl()
	}
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(435)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.SecondaryIndexDecl()
		}
		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DynamoDbGrammarParserT__0 {
			{
				p.SetState(437)
				p.Match(DynamoDbGrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(438)
				p.SecondaryIndexDecl()
			}

			p.SetState(443)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(446)
		p.Match(DynamoDbGrammarParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShowTablesStatementContext is an interface to support dynamic dispatch.
type IShowTablesStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SHOW() antlr.TerminalNode
	TABLES() antlr.TerminalNode

	// IsShowTablesStatementContext differentiates from other interfaces.
	IsShowTablesStatementContext()
}

type ShowTablesStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShowTablesStatementContext() *ShowTablesStatementContext {
	var p = new(ShowTablesStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_showTablesStatement
	return p
}

func InitEmptyShowTablesStatementContext(p *ShowTablesStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_showTablesStatement
}

func (*ShowTablesStatementContext) IsShowTablesStatementContext() {}

func NewShowTablesStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShowTablesStatementContext {
	var p = new(ShowTablesStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_showTablesStatement

	return p
}

func (s *ShowTablesStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ShowTablesStatementContext) SHOW() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSHOW, 0)
}

func (s *ShowTablesStatementContext) TABLES() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserTABLES, 0)
}

func (s *ShowTablesStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShowTablesStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShowTablesStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterShowTablesStatement(s)
	}
}

func (s *ShowTablesStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitShowTablesStatement(s)
	}
}

func (p *DynamoDbGrammarParser) ShowTablesStatement() (localctx IShowTablesStatementContext) {
	localctx = NewShowTablesStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, DynamoDbGrammarParserRULE_showTablesStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Match(DynamoDbGrammarParserSHOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(449)
		p.Match(DynamoDbGrammarParserTABLES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	SelectProjection() ISelectProjectionContext
	FROM() antlr.TerminalNode
	TableName() ITableNameContext
	Hint() IHintContext
	Where() IWhereContext
	OptionBlock() IOptionBlockContext

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_selectStatement
	return p
}

func InitEmptySelectStatementContext(p *SelectStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_selectStatement
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSELECT, 0)
}

func (s *SelectStatementContext) SelectProjection() ISelectProjectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectProjectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectProjectionContext)
}

func (s *SelectStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserFROM, 0)
}

func (s *SelectStatementContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *SelectStatementContext) Hint() IHintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHintContext)
}

func (s *SelectStatementContext) Where() IWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *SelectStatementContext) OptionBlock() IOptionBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionBlockContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSelectStatement(s)
	}
}

func (s *SelectStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSelectStatement(s)
	}
}

func (p *DynamoDbGrammarParser) SelectStatement() (localctx ISelectStatementContext) {
	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, DynamoDbGrammarParserRULE_selectStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(DynamoDbGrammarParserSELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(452)
		p.SelectProjection()
	}
	{
		p.SetState(453)
		p.Match(DynamoDbGrammarParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(454)
		p.TableName()
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserUSE {
		{
			p.SetState(455)
			p.Hint()
		}

	}
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserWHERE {
		{
			p.SetState(458)
			p.Where()
		}

	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserOPTION {
		{
			p.SetState(461)
			p.OptionBlock()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectProjectionContext is an interface to support dynamic dispatch.
type ISelectProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Projection() IProjectionContext
	Star() IStarContext

	// IsSelectProjectionContext differentiates from other interfaces.
	IsSelectProjectionContext()
}

type SelectProjectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectProjectionContext() *SelectProjectionContext {
	var p = new(SelectProjectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_selectProjection
	return p
}

func InitEmptySelectProjectionContext(p *SelectProjectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_selectProjection
}

func (*SelectProjectionContext) IsSelectProjectionContext() {}

func NewSelectProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectProjectionContext {
	var p = new(SelectProjectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_selectProjection

	return p
}

func (s *SelectProjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectProjectionContext) Projection() IProjectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionContext)
}

func (s *SelectProjectionContext) Star() IStarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *SelectProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectProjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectProjectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSelectProjection(s)
	}
}

func (s *SelectProjectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSelectProjection(s)
	}
}

func (p *DynamoDbGrammarParser) SelectProjection() (localctx ISelectProjectionContext) {
	localctx = NewSelectProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, DynamoDbGrammarParserRULE_selectProjection)
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserID, DynamoDbGrammarParserATTRIBUTE_NAME_SUB:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(464)
			p.Projection()
		}

	case DynamoDbGrammarParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
			p.Star()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionBlockContext is an interface to support dynamic dispatch.
type IOptionBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTION() antlr.TerminalNode
	AllOption() []IOptionContext
	Option(i int) IOptionContext

	// IsOptionBlockContext differentiates from other interfaces.
	IsOptionBlockContext()
}

type OptionBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionBlockContext() *OptionBlockContext {
	var p = new(OptionBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionBlock
	return p
}

func InitEmptyOptionBlockContext(p *OptionBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionBlock
}

func (*OptionBlockContext) IsOptionBlockContext() {}

func NewOptionBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionBlockContext {
	var p = new(OptionBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionBlock

	return p
}

func (s *OptionBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionBlockContext) OPTION() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserOPTION, 0)
}

func (s *OptionBlockContext) AllOption() []IOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionContext); ok {
			len++
		}
	}

	tst := make([]IOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionContext); ok {
			tst[i] = t.(IOptionContext)
			i++
		}
	}

	return tst
}

func (s *OptionBlockContext) Option(i int) IOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *OptionBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionBlock(s)
	}
}

func (s *OptionBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionBlock(s)
	}
}

func (p *DynamoDbGrammarParser) OptionBlock() (localctx IOptionBlockContext) {
	localctx = NewOptionBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, DynamoDbGrammarParserRULE_optionBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		p.Match(DynamoDbGrammarParserOPTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(469)
		p.Match(DynamoDbGrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(470)
		p.Option()
	}
	p.SetState(475)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(471)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)
			p.Option()
		}

		p.SetState(477)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(478)
		p.Match(DynamoDbGrammarParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SingleOption() ISingleOptionContext
	KeyValueOption() IKeyValueOptionContext

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_option
	return p
}

func InitEmptyOptionContext(p *OptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_option
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) SingleOption() ISingleOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleOptionContext)
}

func (s *OptionContext) KeyValueOption() IKeyValueOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValueOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyValueOptionContext)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *DynamoDbGrammarParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, DynamoDbGrammarParserRULE_option)
	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(480)
			p.SingleOption()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(481)
			p.KeyValueOption()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleOptionContext is an interface to support dynamic dispatch.
type ISingleOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionKey() IOptionKeyContext

	// IsSingleOptionContext differentiates from other interfaces.
	IsSingleOptionContext()
}

type SingleOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleOptionContext() *SingleOptionContext {
	var p = new(SingleOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_singleOption
	return p
}

func InitEmptySingleOptionContext(p *SingleOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_singleOption
}

func (*SingleOptionContext) IsSingleOptionContext() {}

func NewSingleOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleOptionContext {
	var p = new(SingleOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_singleOption

	return p
}

func (s *SingleOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleOptionContext) OptionKey() IOptionKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionKeyContext)
}

func (s *SingleOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSingleOption(s)
	}
}

func (s *SingleOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSingleOption(s)
	}
}

func (p *DynamoDbGrammarParser) SingleOption() (localctx ISingleOptionContext) {
	localctx = NewSingleOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, DynamoDbGrammarParserRULE_singleOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(484)
		p.OptionKey()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyValueOptionContext is an interface to support dynamic dispatch.
type IKeyValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionKey() IOptionKeyContext
	EQ() antlr.TerminalNode
	OptionValue() IOptionValueContext

	// IsKeyValueOptionContext differentiates from other interfaces.
	IsKeyValueOptionContext()
}

type KeyValueOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueOptionContext() *KeyValueOptionContext {
	var p = new(KeyValueOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_keyValueOption
	return p
}

func InitEmptyKeyValueOptionContext(p *KeyValueOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_keyValueOption
}

func (*KeyValueOptionContext) IsKeyValueOptionContext() {}

func NewKeyValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueOptionContext {
	var p = new(KeyValueOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_keyValueOption

	return p
}

func (s *KeyValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueOptionContext) OptionKey() IOptionKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionKeyContext)
}

func (s *KeyValueOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserEQ, 0)
}

func (s *KeyValueOptionContext) OptionValue() IOptionValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionValueContext)
}

func (s *KeyValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterKeyValueOption(s)
	}
}

func (s *KeyValueOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitKeyValueOption(s)
	}
}

func (p *DynamoDbGrammarParser) KeyValueOption() (localctx IKeyValueOptionContext) {
	localctx = NewKeyValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, DynamoDbGrammarParserRULE_keyValueOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(486)
		p.OptionKey()
	}
	{
		p.SetState(487)
		p.Match(DynamoDbGrammarParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(488)
		p.OptionValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionKeyContext is an interface to support dynamic dispatch.
type IOptionKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DdlName() IDdlNameContext

	// IsOptionKeyContext differentiates from other interfaces.
	IsOptionKeyContext()
}

type OptionKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionKeyContext() *OptionKeyContext {
	var p = new(OptionKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionKey
	return p
}

func InitEmptyOptionKeyContext(p *OptionKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionKey
}

func (*OptionKeyContext) IsOptionKeyContext() {}

func NewOptionKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionKeyContext {
	var p = new(OptionKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionKey

	return p
}

func (s *OptionKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionKeyContext) DdlName() IDdlNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDdlNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *OptionKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionKey(s)
	}
}

func (s *OptionKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionKey(s)
	}
}

func (p *DynamoDbGrammarParser) OptionKey() (localctx IOptionKeyContext) {
	localctx = NewOptionKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, DynamoDbGrammarParserRULE_optionKey)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.DdlName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionValueContext is an interface to support dynamic dispatch.
type IOptionValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionValueString() IOptionValueStringContext
	OptionValueNumber() IOptionValueNumberContext

	// IsOptionValueContext differentiates from other interfaces.
	IsOptionValueContext()
}

type OptionValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueContext() *OptionValueContext {
	var p = new(OptionValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValue
	return p
}

func InitEmptyOptionValueContext(p *OptionValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValue
}

func (*OptionValueContext) IsOptionValueContext() {}

func NewOptionValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueContext {
	var p = new(OptionValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValue

	return p
}

func (s *OptionValueContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueContext) OptionValueString() IOptionValueStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionValueStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionValueStringContext)
}

func (s *OptionValueContext) OptionValueNumber() IOptionValueNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionValueNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionValueNumberContext)
}

func (s *OptionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionValue(s)
	}
}

func (s *OptionValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionValue(s)
	}
}

func (p *DynamoDbGrammarParser) OptionValue() (localctx IOptionValueContext) {
	localctx = NewOptionValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, DynamoDbGrammarParserRULE_optionValue)
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserNONE, DynamoDbGrammarParserID, DynamoDbGrammarParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(492)
			p.OptionValueString()
		}

	case DynamoDbGrammarParserINDEX:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(493)
			p.OptionValueNumber()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionValueStringContext is an interface to support dynamic dispatch.
type IOptionValueStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() IStringContext
	NONE() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsOptionValueStringContext differentiates from other interfaces.
	IsOptionValueStringContext()
}

type OptionValueStringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueStringContext() *OptionValueStringContext {
	var p = new(OptionValueStringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueString
	return p
}

func InitEmptyOptionValueStringContext(p *OptionValueStringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueString
}

func (*OptionValueStringContext) IsOptionValueStringContext() {}

func NewOptionValueStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueStringContext {
	var p = new(OptionValueStringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueString

	return p
}

func (s *OptionValueStringContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueStringContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *OptionValueStringContext) NONE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNONE, 0)
}

func (s *OptionValueStringContext) ID() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserID, 0)
}

func (s *OptionValueStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionValueStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionValueString(s)
	}
}

func (s *OptionValueStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionValueString(s)
	}
}

func (p *DynamoDbGrammarParser) OptionValueString() (localctx IOptionValueStringContext) {
	localctx = NewOptionValueStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, DynamoDbGrammarParserRULE_optionValueString)
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(496)
			p.String_()
		}

	case DynamoDbGrammarParserNONE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(497)
			p.Match(DynamoDbGrammarParserNONE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DynamoDbGrammarParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(498)
			p.Match(DynamoDbGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionValueNumberContext is an interface to support dynamic dispatch.
type IOptionValueNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEX() antlr.TerminalNode

	// IsOptionValueNumberContext differentiates from other interfaces.
	IsOptionValueNumberContext()
}

type OptionValueNumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionValueNumberContext() *OptionValueNumberContext {
	var p = new(OptionValueNumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueNumber
	return p
}

func InitEmptyOptionValueNumberContext(p *OptionValueNumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueNumber
}

func (*OptionValueNumberContext) IsOptionValueNumberContext() {}

func NewOptionValueNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionValueNumberContext {
	var p = new(OptionValueNumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_optionValueNumber

	return p
}

func (s *OptionValueNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionValueNumberContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *OptionValueNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionValueNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionValueNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOptionValueNumber(s)
	}
}

func (s *OptionValueNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOptionValueNumber(s)
	}
}

func (p *DynamoDbGrammarParser) OptionValueNumber() (localctx IOptionValueNumberContext) {
	localctx = NewOptionValueNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, DynamoDbGrammarParserRULE_optionValueNumber)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(501)
		p.Match(DynamoDbGrammarParserINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStarContext is an interface to support dynamic dispatch.
type IStarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStarContext differentiates from other interfaces.
	IsStarContext()
}

type StarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStarContext() *StarContext {
	var p = new(StarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_star
	return p
}

func InitEmptyStarContext(p *StarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_star
}

func (*StarContext) IsStarContext() {}

func NewStarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StarContext {
	var p = new(StarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_star

	return p
}

func (s *StarContext) GetParser() antlr.Parser { return s.parser }
func (s *StarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterStar(s)
	}
}

func (s *StarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitStar(s)
	}
}

func (p *DynamoDbGrammarParser) Star() (localctx IStarContext) {
	localctx = NewStarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, DynamoDbGrammarParserRULE_star)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(503)
		p.Match(DynamoDbGrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHintContext is an interface to support dynamic dispatch.
type IHintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IndexHint() IIndexHintContext
	ScanInfo() IScanInfoContext

	// IsHintContext differentiates from other interfaces.
	IsHintContext()
}

type HintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHintContext() *HintContext {
	var p = new(HintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_hint
	return p
}

func InitEmptyHintContext(p *HintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_hint
}

func (*HintContext) IsHintContext() {}

func NewHintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HintContext {
	var p = new(HintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_hint

	return p
}

func (s *HintContext) GetParser() antlr.Parser { return s.parser }

func (s *HintContext) IndexHint() IIndexHintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexHintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexHintContext)
}

func (s *HintContext) ScanInfo() IScanInfoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScanInfoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScanInfoContext)
}

func (s *HintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterHint(s)
	}
}

func (s *HintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitHint(s)
	}
}

func (p *DynamoDbGrammarParser) Hint() (localctx IHintContext) {
	localctx = NewHintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, DynamoDbGrammarParserRULE_hint)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		p.IndexHint()
	}
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserENABLE {
		{
			p.SetState(506)
			p.ScanInfo()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexHintContext is an interface to support dynamic dispatch.
type IIndexHintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USE() antlr.TerminalNode
	IndexHintName() IIndexHintNameContext

	// IsIndexHintContext differentiates from other interfaces.
	IsIndexHintContext()
}

type IndexHintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexHintContext() *IndexHintContext {
	var p = new(IndexHintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHint
	return p
}

func InitEmptyIndexHintContext(p *IndexHintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHint
}

func (*IndexHintContext) IsIndexHintContext() {}

func NewIndexHintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexHintContext {
	var p = new(IndexHintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHint

	return p
}

func (s *IndexHintContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexHintContext) USE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUSE, 0)
}

func (s *IndexHintContext) IndexHintName() IIndexHintNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexHintNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexHintNameContext)
}

func (s *IndexHintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexHintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexHintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIndexHint(s)
	}
}

func (s *IndexHintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIndexHint(s)
	}
}

func (p *DynamoDbGrammarParser) IndexHint() (localctx IIndexHintContext) {
	localctx = NewIndexHintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, DynamoDbGrammarParserRULE_indexHint)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(509)
		p.Match(DynamoDbGrammarParserUSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(510)
		p.IndexHintName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexHintNameContext is an interface to support dynamic dispatch.
type IIndexHintNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEXKEYWORD() antlr.TerminalNode
	IndexName() IIndexNameContext
	PRIMARY() antlr.TerminalNode

	// IsIndexHintNameContext differentiates from other interfaces.
	IsIndexHintNameContext()
}

type IndexHintNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexHintNameContext() *IndexHintNameContext {
	var p = new(IndexHintNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHintName
	return p
}

func InitEmptyIndexHintNameContext(p *IndexHintNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHintName
}

func (*IndexHintNameContext) IsIndexHintNameContext() {}

func NewIndexHintNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexHintNameContext {
	var p = new(IndexHintNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_indexHintName

	return p
}

func (s *IndexHintNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexHintNameContext) INDEXKEYWORD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEXKEYWORD, 0)
}

func (s *IndexHintNameContext) IndexName() IIndexNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *IndexHintNameContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserPRIMARY, 0)
}

func (s *IndexHintNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexHintNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexHintNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIndexHintName(s)
	}
}

func (s *IndexHintNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIndexHintName(s)
	}
}

func (p *DynamoDbGrammarParser) IndexHintName() (localctx IIndexHintNameContext) {
	localctx = NewIndexHintNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, DynamoDbGrammarParserRULE_indexHintName)
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserINDEXKEYWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(512)
			p.Match(DynamoDbGrammarParserINDEXKEYWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(513)
			p.IndexName()
		}

	case DynamoDbGrammarParserPRIMARY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(514)
			p.Match(DynamoDbGrammarParserPRIMARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(515)
			p.Match(DynamoDbGrammarParserINDEXKEYWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScanInfoContext is an interface to support dynamic dispatch.
type IScanInfoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENABLE() antlr.TerminalNode
	SCAN() antlr.TerminalNode
	TotalSegment() ITotalSegmentContext
	Segment() ISegmentContext

	// IsScanInfoContext differentiates from other interfaces.
	IsScanInfoContext()
}

type ScanInfoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScanInfoContext() *ScanInfoContext {
	var p = new(ScanInfoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_scanInfo
	return p
}

func InitEmptyScanInfoContext(p *ScanInfoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_scanInfo
}

func (*ScanInfoContext) IsScanInfoContext() {}

func NewScanInfoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScanInfoContext {
	var p = new(ScanInfoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_scanInfo

	return p
}

func (s *ScanInfoContext) GetParser() antlr.Parser { return s.parser }

func (s *ScanInfoContext) ENABLE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserENABLE, 0)
}

func (s *ScanInfoContext) SCAN() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSCAN, 0)
}

func (s *ScanInfoContext) TotalSegment() ITotalSegmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITotalSegmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITotalSegmentContext)
}

func (s *ScanInfoContext) Segment() ISegmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISegmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISegmentContext)
}

func (s *ScanInfoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScanInfoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScanInfoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterScanInfo(s)
	}
}

func (s *ScanInfoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitScanInfo(s)
	}
}

func (p *DynamoDbGrammarParser) ScanInfo() (localctx IScanInfoContext) {
	localctx = NewScanInfoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, DynamoDbGrammarParserRULE_scanInfo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(518)
		p.Match(DynamoDbGrammarParserENABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(519)
		p.Match(DynamoDbGrammarParserSCAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserT__1 {
		{
			p.SetState(520)
			p.Match(DynamoDbGrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(521)
			p.TotalSegment()
		}
		{
			p.SetState(522)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(523)
			p.Segment()
		}
		{
			p.SetState(524)
			p.Match(DynamoDbGrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITotalSegmentContext is an interface to support dynamic dispatch.
type ITotalSegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEX() antlr.TerminalNode

	// IsTotalSegmentContext differentiates from other interfaces.
	IsTotalSegmentContext()
}

type TotalSegmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTotalSegmentContext() *TotalSegmentContext {
	var p = new(TotalSegmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_totalSegment
	return p
}

func InitEmptyTotalSegmentContext(p *TotalSegmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_totalSegment
}

func (*TotalSegmentContext) IsTotalSegmentContext() {}

func NewTotalSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TotalSegmentContext {
	var p = new(TotalSegmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_totalSegment

	return p
}

func (s *TotalSegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *TotalSegmentContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *TotalSegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TotalSegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TotalSegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterTotalSegment(s)
	}
}

func (s *TotalSegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitTotalSegment(s)
	}
}

func (p *DynamoDbGrammarParser) TotalSegment() (localctx ITotalSegmentContext) {
	localctx = NewTotalSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, DynamoDbGrammarParserRULE_totalSegment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		p.Match(DynamoDbGrammarParserINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISegmentContext is an interface to support dynamic dispatch.
type ISegmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEX() antlr.TerminalNode

	// IsSegmentContext differentiates from other interfaces.
	IsSegmentContext()
}

type SegmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySegmentContext() *SegmentContext {
	var p = new(SegmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_segment
	return p
}

func InitEmptySegmentContext(p *SegmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_segment
}

func (*SegmentContext) IsSegmentContext() {}

func NewSegmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SegmentContext {
	var p = new(SegmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_segment

	return p
}

func (s *SegmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SegmentContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *SegmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SegmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SegmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSegment(s)
	}
}

func (s *SegmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSegment(s)
	}
}

func (p *DynamoDbGrammarParser) Segment() (localctx ISegmentContext) {
	localctx = NewSegmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, DynamoDbGrammarParserRULE_segment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
		p.Match(DynamoDbGrammarParserINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhereContext is an interface to support dynamic dispatch.
type IWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	Condition() IConditionContext

	// IsWhereContext differentiates from other interfaces.
	IsWhereContext()
}

type WhereContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereContext() *WhereContext {
	var p = new(WhereContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_where
	return p
}

func InitEmptyWhereContext(p *WhereContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_where
}

func (*WhereContext) IsWhereContext() {}

func NewWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereContext {
	var p = new(WhereContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_where

	return p
}

func (s *WhereContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereContext) WHERE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserWHERE, 0)
}

func (s *WhereContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterWhere(s)
	}
}

func (s *WhereContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitWhere(s)
	}
}

func (p *DynamoDbGrammarParser) Where() (localctx IWhereContext) {
	localctx = NewWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, DynamoDbGrammarParserRULE_where)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
		p.Match(DynamoDbGrammarParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(533)
		p.condition(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryKeyDeclContext is an interface to support dynamic dispatch.
type IPrimaryKeyDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIMARY() antlr.TerminalNode
	KEY() antlr.TerminalNode
	IndexDecl() IIndexDeclContext
	Capacity() ICapacityContext

	// IsPrimaryKeyDeclContext differentiates from other interfaces.
	IsPrimaryKeyDeclContext()
}

type PrimaryKeyDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryKeyDeclContext() *PrimaryKeyDeclContext {
	var p = new(PrimaryKeyDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_primaryKeyDecl
	return p
}

func InitEmptyPrimaryKeyDeclContext(p *PrimaryKeyDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_primaryKeyDecl
}

func (*PrimaryKeyDeclContext) IsPrimaryKeyDeclContext() {}

func NewPrimaryKeyDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryKeyDeclContext {
	var p = new(PrimaryKeyDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_primaryKeyDecl

	return p
}

func (s *PrimaryKeyDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryKeyDeclContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserPRIMARY, 0)
}

func (s *PrimaryKeyDeclContext) KEY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserKEY, 0)
}

func (s *PrimaryKeyDeclContext) IndexDecl() IIndexDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexDeclContext)
}

func (s *PrimaryKeyDeclContext) Capacity() ICapacityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapacityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapacityContext)
}

func (s *PrimaryKeyDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryKeyDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryKeyDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterPrimaryKeyDecl(s)
	}
}

func (s *PrimaryKeyDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitPrimaryKeyDecl(s)
	}
}

func (p *DynamoDbGrammarParser) PrimaryKeyDecl() (localctx IPrimaryKeyDeclContext) {
	localctx = NewPrimaryKeyDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, DynamoDbGrammarParserRULE_primaryKeyDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(535)
		p.Match(DynamoDbGrammarParserPRIMARY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(536)
		p.Match(DynamoDbGrammarParserKEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(537)
		p.IndexDecl()
	}
	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserCAPACITY {
		{
			p.SetState(538)
			p.Capacity()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISecondaryIndexDeclContext is an interface to support dynamic dispatch.
type ISecondaryIndexDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEXKEYWORD() antlr.TerminalNode
	IndexName() IIndexNameContext
	IndexDecl() IIndexDeclContext
	SecondaryIndexType() ISecondaryIndexTypeContext
	AllAttributeDecl() []IAttributeDeclContext
	AttributeDecl(i int) IAttributeDeclContext
	ProjectionIndex() IProjectionIndexContext
	Capacity() ICapacityContext

	// IsSecondaryIndexDeclContext differentiates from other interfaces.
	IsSecondaryIndexDeclContext()
}

type SecondaryIndexDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecondaryIndexDeclContext() *SecondaryIndexDeclContext {
	var p = new(SecondaryIndexDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexDecl
	return p
}

func InitEmptySecondaryIndexDeclContext(p *SecondaryIndexDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexDecl
}

func (*SecondaryIndexDeclContext) IsSecondaryIndexDeclContext() {}

func NewSecondaryIndexDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecondaryIndexDeclContext {
	var p = new(SecondaryIndexDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexDecl

	return p
}

func (s *SecondaryIndexDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SecondaryIndexDeclContext) INDEXKEYWORD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEXKEYWORD, 0)
}

func (s *SecondaryIndexDeclContext) IndexName() IIndexNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *SecondaryIndexDeclContext) IndexDecl() IIndexDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexDeclContext)
}

func (s *SecondaryIndexDeclContext) SecondaryIndexType() ISecondaryIndexTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISecondaryIndexTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISecondaryIndexTypeContext)
}

func (s *SecondaryIndexDeclContext) AllAttributeDecl() []IAttributeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeDeclContext); ok {
			len++
		}
	}

	tst := make([]IAttributeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeDeclContext); ok {
			tst[i] = t.(IAttributeDeclContext)
			i++
		}
	}

	return tst
}

func (s *SecondaryIndexDeclContext) AttributeDecl(i int) IAttributeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeDeclContext)
}

func (s *SecondaryIndexDeclContext) ProjectionIndex() IProjectionIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionIndexContext)
}

func (s *SecondaryIndexDeclContext) Capacity() ICapacityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapacityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapacityContext)
}

func (s *SecondaryIndexDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondaryIndexDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SecondaryIndexDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSecondaryIndexDecl(s)
	}
}

func (s *SecondaryIndexDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSecondaryIndexDecl(s)
	}
}

func (p *DynamoDbGrammarParser) SecondaryIndexDecl() (localctx ISecondaryIndexDeclContext) {
	localctx = NewSecondaryIndexDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, DynamoDbGrammarParserRULE_secondaryIndexDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserGLOBAL || _la == DynamoDbGrammarParserLOCAL {
		{
			p.SetState(541)
			p.SecondaryIndexType()
		}

	}
	{
		p.SetState(544)
		p.Match(DynamoDbGrammarParserINDEXKEYWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(545)
			p.AttributeDecl()
		}
		p.SetState(550)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DynamoDbGrammarParserT__0 {
			{
				p.SetState(546)
				p.Match(DynamoDbGrammarParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(547)
				p.AttributeDecl()
			}

			p.SetState(552)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(555)
		p.IndexName()
	}
	{
		p.SetState(556)
		p.IndexDecl()
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserPROJECTION {
		{
			p.SetState(557)
			p.ProjectionIndex()
		}

	}
	p.SetState(561)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserCAPACITY {
		{
			p.SetState(560)
			p.Capacity()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISecondaryIndexTypeContext is an interface to support dynamic dispatch.
type ISecondaryIndexTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GLOBAL() antlr.TerminalNode
	LOCAL() antlr.TerminalNode

	// IsSecondaryIndexTypeContext differentiates from other interfaces.
	IsSecondaryIndexTypeContext()
}

type SecondaryIndexTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecondaryIndexTypeContext() *SecondaryIndexTypeContext {
	var p = new(SecondaryIndexTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexType
	return p
}

func InitEmptySecondaryIndexTypeContext(p *SecondaryIndexTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexType
}

func (*SecondaryIndexTypeContext) IsSecondaryIndexTypeContext() {}

func NewSecondaryIndexTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecondaryIndexTypeContext {
	var p = new(SecondaryIndexTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_secondaryIndexType

	return p
}

func (s *SecondaryIndexTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SecondaryIndexTypeContext) GLOBAL() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserGLOBAL, 0)
}

func (s *SecondaryIndexTypeContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserLOCAL, 0)
}

func (s *SecondaryIndexTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondaryIndexTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SecondaryIndexTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterSecondaryIndexType(s)
	}
}

func (s *SecondaryIndexTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitSecondaryIndexType(s)
	}
}

func (p *DynamoDbGrammarParser) SecondaryIndexType() (localctx ISecondaryIndexTypeContext) {
	localctx = NewSecondaryIndexTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, DynamoDbGrammarParserRULE_secondaryIndexType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(563)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DynamoDbGrammarParserGLOBAL || _la == DynamoDbGrammarParserLOCAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexNameContext is an interface to support dynamic dispatch.
type IIndexNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DdlName() IDdlNameContext

	// IsIndexNameContext differentiates from other interfaces.
	IsIndexNameContext()
}

type IndexNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexNameContext() *IndexNameContext {
	var p = new(IndexNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexName
	return p
}

func InitEmptyIndexNameContext(p *IndexNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexName
}

func (*IndexNameContext) IsIndexNameContext() {}

func NewIndexNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexNameContext {
	var p = new(IndexNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_indexName

	return p
}

func (s *IndexNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexNameContext) DdlName() IDdlNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDdlNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *IndexNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIndexName(s)
	}
}

func (s *IndexNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIndexName(s)
	}
}

func (p *DynamoDbGrammarParser) IndexName() (localctx IIndexNameContext) {
	localctx = NewIndexNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, DynamoDbGrammarParserRULE_indexName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(565)
		p.DdlName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProjectionIndexContext is an interface to support dynamic dispatch.
type IProjectionIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ProjectionIndexKeysOnly() IProjectionIndexKeysOnlyContext
	ProjectionIndexVector() IProjectionIndexVectorContext

	// IsProjectionIndexContext differentiates from other interfaces.
	IsProjectionIndexContext()
}

type ProjectionIndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionIndexContext() *ProjectionIndexContext {
	var p = new(ProjectionIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndex
	return p
}

func InitEmptyProjectionIndexContext(p *ProjectionIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndex
}

func (*ProjectionIndexContext) IsProjectionIndexContext() {}

func NewProjectionIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionIndexContext {
	var p = new(ProjectionIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndex

	return p
}

func (s *ProjectionIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionIndexContext) ProjectionIndexKeysOnly() IProjectionIndexKeysOnlyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionIndexKeysOnlyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionIndexKeysOnlyContext)
}

func (s *ProjectionIndexContext) ProjectionIndexVector() IProjectionIndexVectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectionIndexVectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectionIndexVectorContext)
}

func (s *ProjectionIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjectionIndex(s)
	}
}

func (s *ProjectionIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjectionIndex(s)
	}
}

func (p *DynamoDbGrammarParser) ProjectionIndex() (localctx IProjectionIndexContext) {
	localctx = NewProjectionIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, DynamoDbGrammarParserRULE_projectionIndex)
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(567)
			p.ProjectionIndexKeysOnly()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(568)
			p.ProjectionIndexVector()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProjectionIndexKeysOnlyContext is an interface to support dynamic dispatch.
type IProjectionIndexKeysOnlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROJECTION() antlr.TerminalNode
	KEYS() antlr.TerminalNode
	ONLY() antlr.TerminalNode

	// IsProjectionIndexKeysOnlyContext differentiates from other interfaces.
	IsProjectionIndexKeysOnlyContext()
}

type ProjectionIndexKeysOnlyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionIndexKeysOnlyContext() *ProjectionIndexKeysOnlyContext {
	var p = new(ProjectionIndexKeysOnlyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexKeysOnly
	return p
}

func InitEmptyProjectionIndexKeysOnlyContext(p *ProjectionIndexKeysOnlyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexKeysOnly
}

func (*ProjectionIndexKeysOnlyContext) IsProjectionIndexKeysOnlyContext() {}

func NewProjectionIndexKeysOnlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionIndexKeysOnlyContext {
	var p = new(ProjectionIndexKeysOnlyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexKeysOnly

	return p
}

func (s *ProjectionIndexKeysOnlyContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionIndexKeysOnlyContext) PROJECTION() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserPROJECTION, 0)
}

func (s *ProjectionIndexKeysOnlyContext) KEYS() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserKEYS, 0)
}

func (s *ProjectionIndexKeysOnlyContext) ONLY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserONLY, 0)
}

func (s *ProjectionIndexKeysOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionIndexKeysOnlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionIndexKeysOnlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjectionIndexKeysOnly(s)
	}
}

func (s *ProjectionIndexKeysOnlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjectionIndexKeysOnly(s)
	}
}

func (p *DynamoDbGrammarParser) ProjectionIndexKeysOnly() (localctx IProjectionIndexKeysOnlyContext) {
	localctx = NewProjectionIndexKeysOnlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, DynamoDbGrammarParserRULE_projectionIndexKeysOnly)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(571)
		p.Match(DynamoDbGrammarParserPROJECTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(572)
		p.Match(DynamoDbGrammarParserKEYS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(573)
		p.Match(DynamoDbGrammarParserONLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProjectionIndexVectorContext is an interface to support dynamic dispatch.
type IProjectionIndexVectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROJECTION() antlr.TerminalNode
	AllAttributeName() []IAttributeNameContext
	AttributeName(i int) IAttributeNameContext

	// IsProjectionIndexVectorContext differentiates from other interfaces.
	IsProjectionIndexVectorContext()
}

type ProjectionIndexVectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectionIndexVectorContext() *ProjectionIndexVectorContext {
	var p = new(ProjectionIndexVectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexVector
	return p
}

func InitEmptyProjectionIndexVectorContext(p *ProjectionIndexVectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexVector
}

func (*ProjectionIndexVectorContext) IsProjectionIndexVectorContext() {}

func NewProjectionIndexVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectionIndexVectorContext {
	var p = new(ProjectionIndexVectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_projectionIndexVector

	return p
}

func (s *ProjectionIndexVectorContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectionIndexVectorContext) PROJECTION() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserPROJECTION, 0)
}

func (s *ProjectionIndexVectorContext) AllAttributeName() []IAttributeNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeNameContext); ok {
			len++
		}
	}

	tst := make([]IAttributeNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeNameContext); ok {
			tst[i] = t.(IAttributeNameContext)
			i++
		}
	}

	return tst
}

func (s *ProjectionIndexVectorContext) AttributeName(i int) IAttributeNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeNameContext)
}

func (s *ProjectionIndexVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectionIndexVectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectionIndexVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterProjectionIndexVector(s)
	}
}

func (s *ProjectionIndexVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitProjectionIndexVector(s)
	}
}

func (p *DynamoDbGrammarParser) ProjectionIndexVector() (localctx IProjectionIndexVectorContext) {
	localctx = NewProjectionIndexVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, DynamoDbGrammarParserRULE_projectionIndexVector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(575)
		p.Match(DynamoDbGrammarParserPROJECTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(576)
		p.Match(DynamoDbGrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(577)
		p.AttributeName()
	}
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(578)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(579)
			p.AttributeName()
		}

		p.SetState(584)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(585)
		p.Match(DynamoDbGrammarParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICapacityContext is an interface to support dynamic dispatch.
type ICapacityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CAPACITY() antlr.TerminalNode
	ReadUnits() IReadUnitsContext
	WriteUnits() IWriteUnitsContext

	// IsCapacityContext differentiates from other interfaces.
	IsCapacityContext()
}

type CapacityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapacityContext() *CapacityContext {
	var p = new(CapacityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_capacity
	return p
}

func InitEmptyCapacityContext(p *CapacityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_capacity
}

func (*CapacityContext) IsCapacityContext() {}

func NewCapacityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapacityContext {
	var p = new(CapacityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_capacity

	return p
}

func (s *CapacityContext) GetParser() antlr.Parser { return s.parser }

func (s *CapacityContext) CAPACITY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserCAPACITY, 0)
}

func (s *CapacityContext) ReadUnits() IReadUnitsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadUnitsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadUnitsContext)
}

func (s *CapacityContext) WriteUnits() IWriteUnitsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWriteUnitsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWriteUnitsContext)
}

func (s *CapacityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapacityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CapacityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterCapacity(s)
	}
}

func (s *CapacityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitCapacity(s)
	}
}

func (p *DynamoDbGrammarParser) Capacity() (localctx ICapacityContext) {
	localctx = NewCapacityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, DynamoDbGrammarParserRULE_capacity)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(587)
		p.Match(DynamoDbGrammarParserCAPACITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(588)
		p.Match(DynamoDbGrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(589)
		p.ReadUnits()
	}
	{
		p.SetState(590)
		p.Match(DynamoDbGrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(591)
		p.WriteUnits()
	}
	{
		p.SetState(592)
		p.Match(DynamoDbGrammarParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReadUnitsContext is an interface to support dynamic dispatch.
type IReadUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEX() antlr.TerminalNode

	// IsReadUnitsContext differentiates from other interfaces.
	IsReadUnitsContext()
}

type ReadUnitsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadUnitsContext() *ReadUnitsContext {
	var p = new(ReadUnitsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_readUnits
	return p
}

func InitEmptyReadUnitsContext(p *ReadUnitsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_readUnits
}

func (*ReadUnitsContext) IsReadUnitsContext() {}

func NewReadUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadUnitsContext {
	var p = new(ReadUnitsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_readUnits

	return p
}

func (s *ReadUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadUnitsContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *ReadUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterReadUnits(s)
	}
}

func (s *ReadUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitReadUnits(s)
	}
}

func (p *DynamoDbGrammarParser) ReadUnits() (localctx IReadUnitsContext) {
	localctx = NewReadUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, DynamoDbGrammarParserRULE_readUnits)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(594)
		p.Match(DynamoDbGrammarParserINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWriteUnitsContext is an interface to support dynamic dispatch.
type IWriteUnitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEX() antlr.TerminalNode

	// IsWriteUnitsContext differentiates from other interfaces.
	IsWriteUnitsContext()
}

type WriteUnitsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWriteUnitsContext() *WriteUnitsContext {
	var p = new(WriteUnitsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_writeUnits
	return p
}

func InitEmptyWriteUnitsContext(p *WriteUnitsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_writeUnits
}

func (*WriteUnitsContext) IsWriteUnitsContext() {}

func NewWriteUnitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WriteUnitsContext {
	var p = new(WriteUnitsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_writeUnits

	return p
}

func (s *WriteUnitsContext) GetParser() antlr.Parser { return s.parser }

func (s *WriteUnitsContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserINDEX, 0)
}

func (s *WriteUnitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteUnitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WriteUnitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterWriteUnits(s)
	}
}

func (s *WriteUnitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitWriteUnits(s)
	}
}

func (p *DynamoDbGrammarParser) WriteUnits() (localctx IWriteUnitsContext) {
	localctx = NewWriteUnitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, DynamoDbGrammarParserRULE_writeUnits)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(596)
		p.Match(DynamoDbGrammarParserINDEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexDeclContext is an interface to support dynamic dispatch.
type IIndexDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HashKey() IHashKeyContext
	RangeKey() IRangeKeyContext

	// IsIndexDeclContext differentiates from other interfaces.
	IsIndexDeclContext()
}

type IndexDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexDeclContext() *IndexDeclContext {
	var p = new(IndexDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexDecl
	return p
}

func InitEmptyIndexDeclContext(p *IndexDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_indexDecl
}

func (*IndexDeclContext) IsIndexDeclContext() {}

func NewIndexDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexDeclContext {
	var p = new(IndexDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_indexDecl

	return p
}

func (s *IndexDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexDeclContext) HashKey() IHashKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHashKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHashKeyContext)
}

func (s *IndexDeclContext) RangeKey() IRangeKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeKeyContext)
}

func (s *IndexDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIndexDecl(s)
	}
}

func (s *IndexDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIndexDecl(s)
	}
}

func (p *DynamoDbGrammarParser) IndexDecl() (localctx IIndexDeclContext) {
	localctx = NewIndexDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, DynamoDbGrammarParserRULE_indexDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(598)
		p.Match(DynamoDbGrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(599)
		p.HashKey()
	}
	p.SetState(602)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserT__0 {
		{
			p.SetState(600)
			p.Match(DynamoDbGrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(601)
			p.RangeKey()
		}

	}
	{
		p.SetState(604)
		p.Match(DynamoDbGrammarParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeDeclContext is an interface to support dynamic dispatch.
type IAttributeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttributeName() IAttributeNameContext
	AttributeType() IAttributeTypeContext

	// IsAttributeDeclContext differentiates from other interfaces.
	IsAttributeDeclContext()
}

type AttributeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeDeclContext() *AttributeDeclContext {
	var p = new(AttributeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeDecl
	return p
}

func InitEmptyAttributeDeclContext(p *AttributeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeDecl
}

func (*AttributeDeclContext) IsAttributeDeclContext() {}

func NewAttributeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeDeclContext {
	var p = new(AttributeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeDecl

	return p
}

func (s *AttributeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeDeclContext) AttributeName() IAttributeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeNameContext)
}

func (s *AttributeDeclContext) AttributeType() IAttributeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeContext)
}

func (s *AttributeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAttributeDecl(s)
	}
}

func (s *AttributeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAttributeDecl(s)
	}
}

func (p *DynamoDbGrammarParser) AttributeDecl() (localctx IAttributeDeclContext) {
	localctx = NewAttributeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, DynamoDbGrammarParserRULE_attributeDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(606)
		p.AttributeName()
	}
	{
		p.SetState(607)
		p.AttributeType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHashKeyContext is an interface to support dynamic dispatch.
type IHashKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DdlName() IDdlNameContext

	// IsHashKeyContext differentiates from other interfaces.
	IsHashKeyContext()
}

type HashKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHashKeyContext() *HashKeyContext {
	var p = new(HashKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_hashKey
	return p
}

func InitEmptyHashKeyContext(p *HashKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_hashKey
}

func (*HashKeyContext) IsHashKeyContext() {}

func NewHashKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HashKeyContext {
	var p = new(HashKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_hashKey

	return p
}

func (s *HashKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *HashKeyContext) DdlName() IDdlNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDdlNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *HashKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HashKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterHashKey(s)
	}
}

func (s *HashKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitHashKey(s)
	}
}

func (p *DynamoDbGrammarParser) HashKey() (localctx IHashKeyContext) {
	localctx = NewHashKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, DynamoDbGrammarParserRULE_hashKey)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(609)
		p.DdlName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeKeyContext is an interface to support dynamic dispatch.
type IRangeKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DdlName() IDdlNameContext

	// IsRangeKeyContext differentiates from other interfaces.
	IsRangeKeyContext()
}

type RangeKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeKeyContext() *RangeKeyContext {
	var p = new(RangeKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_rangeKey
	return p
}

func InitEmptyRangeKeyContext(p *RangeKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_rangeKey
}

func (*RangeKeyContext) IsRangeKeyContext() {}

func NewRangeKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeKeyContext {
	var p = new(RangeKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_rangeKey

	return p
}

func (s *RangeKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeKeyContext) DdlName() IDdlNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDdlNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *RangeKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterRangeKey(s)
	}
}

func (s *RangeKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitRangeKey(s)
	}
}

func (p *DynamoDbGrammarParser) RangeKey() (localctx IRangeKeyContext) {
	localctx = NewRangeKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, DynamoDbGrammarParserRULE_rangeKey)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(611)
		p.DdlName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeNameContext is an interface to support dynamic dispatch.
type IAttributeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DdlName() IDdlNameContext

	// IsAttributeNameContext differentiates from other interfaces.
	IsAttributeNameContext()
}

type AttributeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeNameContext() *AttributeNameContext {
	var p = new(AttributeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeName
	return p
}

func InitEmptyAttributeNameContext(p *AttributeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeName
}

func (*AttributeNameContext) IsAttributeNameContext() {}

func NewAttributeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeNameContext {
	var p = new(AttributeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeName

	return p
}

func (s *AttributeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeNameContext) DdlName() IDdlNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDdlNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *AttributeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAttributeName(s)
	}
}

func (s *AttributeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAttributeName(s)
	}
}

func (p *DynamoDbGrammarParser) AttributeName() (localctx IAttributeNameContext) {
	localctx = NewAttributeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, DynamoDbGrammarParserRULE_attributeName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(613)
		p.DdlName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeTypeContext is an interface to support dynamic dispatch.
type IAttributeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BINARY() antlr.TerminalNode

	// IsAttributeTypeContext differentiates from other interfaces.
	IsAttributeTypeContext()
}

type AttributeTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeTypeContext() *AttributeTypeContext {
	var p = new(AttributeTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeType
	return p
}

func InitEmptyAttributeTypeContext(p *AttributeTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeType
}

func (*AttributeTypeContext) IsAttributeTypeContext() {}

func NewAttributeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeTypeContext {
	var p = new(AttributeTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_attributeType

	return p
}

func (s *AttributeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeTypeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNUMBER, 0)
}

func (s *AttributeTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSTRING, 0)
}

func (s *AttributeTypeContext) BINARY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserBINARY, 0)
}

func (s *AttributeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterAttributeType(s)
	}
}

func (s *AttributeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitAttributeType(s)
	}
}

func (p *DynamoDbGrammarParser) AttributeType() (localctx IAttributeTypeContext) {
	localctx = NewAttributeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, DynamoDbGrammarParserRULE_attributeType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(615)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&246290604621824) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturningContext is an interface to support dynamic dispatch.
type IReturningContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURNING() antlr.TerminalNode
	ReturningValue() IReturningValueContext

	// IsReturningContext differentiates from other interfaces.
	IsReturningContext()
}

type ReturningContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturningContext() *ReturningContext {
	var p = new(ReturningContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_returning
	return p
}

func InitEmptyReturningContext(p *ReturningContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_returning
}

func (*ReturningContext) IsReturningContext() {}

func NewReturningContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturningContext {
	var p = new(ReturningContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_returning

	return p
}

func (s *ReturningContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturningContext) RETURNING() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserRETURNING, 0)
}

func (s *ReturningContext) ReturningValue() IReturningValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturningValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturningValueContext)
}

func (s *ReturningContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturningContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturningContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterReturning(s)
	}
}

func (s *ReturningContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitReturning(s)
	}
}

func (p *DynamoDbGrammarParser) Returning() (localctx IReturningContext) {
	localctx = NewReturningContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, DynamoDbGrammarParserRULE_returning)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(617)
		p.Match(DynamoDbGrammarParserRETURNING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(618)
		p.ReturningValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturningValueContext is an interface to support dynamic dispatch.
type IReturningValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NONE() antlr.TerminalNode
	ALL() antlr.TerminalNode
	OLD() antlr.TerminalNode
	UPDATED() antlr.TerminalNode
	NEW() antlr.TerminalNode

	// IsReturningValueContext differentiates from other interfaces.
	IsReturningValueContext()
}

type ReturningValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturningValueContext() *ReturningValueContext {
	var p = new(ReturningValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_returningValue
	return p
}

func InitEmptyReturningValueContext(p *ReturningValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_returningValue
}

func (*ReturningValueContext) IsReturningValueContext() {}

func NewReturningValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturningValueContext {
	var p = new(ReturningValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_returningValue

	return p
}

func (s *ReturningValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturningValueContext) NONE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNONE, 0)
}

func (s *ReturningValueContext) ALL() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserALL, 0)
}

func (s *ReturningValueContext) OLD() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserOLD, 0)
}

func (s *ReturningValueContext) UPDATED() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUPDATED, 0)
}

func (s *ReturningValueContext) NEW() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserNEW, 0)
}

func (s *ReturningValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturningValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturningValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterReturningValue(s)
	}
}

func (s *ReturningValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitReturningValue(s)
	}
}

func (p *DynamoDbGrammarParser) ReturningValue() (localctx IReturningValueContext) {
	localctx = NewReturningValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, DynamoDbGrammarParserRULE_returningValue)
	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(620)
			p.Match(DynamoDbGrammarParserNONE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(621)
			p.Match(DynamoDbGrammarParserALL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)
			p.Match(DynamoDbGrammarParserOLD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(623)
			p.Match(DynamoDbGrammarParserUPDATED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(624)
			p.Match(DynamoDbGrammarParserOLD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(625)
			p.Match(DynamoDbGrammarParserALL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(626)
			p.Match(DynamoDbGrammarParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(627)
			p.Match(DynamoDbGrammarParserUPDATED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(628)
			p.Match(DynamoDbGrammarParserNEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOnDuplicateKeyUpdateContext is an interface to support dynamic dispatch.
type IOnDuplicateKeyUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ON() antlr.TerminalNode
	DUPLICATE() antlr.TerminalNode
	KEY() antlr.TerminalNode
	UPDATE() antlr.TerminalNode
	IfClause() IIfClauseContext

	// IsOnDuplicateKeyUpdateContext differentiates from other interfaces.
	IsOnDuplicateKeyUpdateContext()
}

type OnDuplicateKeyUpdateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnDuplicateKeyUpdateContext() *OnDuplicateKeyUpdateContext {
	var p = new(OnDuplicateKeyUpdateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_onDuplicateKeyUpdate
	return p
}

func InitEmptyOnDuplicateKeyUpdateContext(p *OnDuplicateKeyUpdateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_onDuplicateKeyUpdate
}

func (*OnDuplicateKeyUpdateContext) IsOnDuplicateKeyUpdateContext() {}

func NewOnDuplicateKeyUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OnDuplicateKeyUpdateContext {
	var p = new(OnDuplicateKeyUpdateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_onDuplicateKeyUpdate

	return p
}

func (s *OnDuplicateKeyUpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *OnDuplicateKeyUpdateContext) ON() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserON, 0)
}

func (s *OnDuplicateKeyUpdateContext) DUPLICATE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserDUPLICATE, 0)
}

func (s *OnDuplicateKeyUpdateContext) KEY() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserKEY, 0)
}

func (s *OnDuplicateKeyUpdateContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUPDATE, 0)
}

func (s *OnDuplicateKeyUpdateContext) IfClause() IIfClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfClauseContext)
}

func (s *OnDuplicateKeyUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OnDuplicateKeyUpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OnDuplicateKeyUpdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterOnDuplicateKeyUpdate(s)
	}
}

func (s *OnDuplicateKeyUpdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitOnDuplicateKeyUpdate(s)
	}
}

func (p *DynamoDbGrammarParser) OnDuplicateKeyUpdate() (localctx IOnDuplicateKeyUpdateContext) {
	localctx = NewOnDuplicateKeyUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, DynamoDbGrammarParserRULE_onDuplicateKeyUpdate)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(631)
		p.Match(DynamoDbGrammarParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(632)
		p.Match(DynamoDbGrammarParserDUPLICATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(633)
		p.Match(DynamoDbGrammarParserKEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(634)
		p.Match(DynamoDbGrammarParserUPDATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(636)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DynamoDbGrammarParserIF {
		{
			p.SetState(635)
			p.IfClause()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfClauseContext is an interface to support dynamic dispatch.
type IIfClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Condition() IConditionContext

	// IsIfClauseContext differentiates from other interfaces.
	IsIfClauseContext()
}

type IfClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfClauseContext() *IfClauseContext {
	var p = new(IfClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_ifClause
	return p
}

func InitEmptyIfClauseContext(p *IfClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_ifClause
}

func (*IfClauseContext) IsIfClauseContext() {}

func NewIfClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfClauseContext {
	var p = new(IfClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_ifClause

	return p
}

func (s *IfClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *IfClauseContext) IF() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserIF, 0)
}

func (s *IfClauseContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterIfClause(s)
	}
}

func (s *IfClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitIfClause(s)
	}
}

func (p *DynamoDbGrammarParser) IfClause() (localctx IIfClauseContext) {
	localctx = NewIfClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, DynamoDbGrammarParserRULE_ifClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(638)
		p.Match(DynamoDbGrammarParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(639)
		p.condition(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DdlName() IDdlNameContext

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_tableName
	return p
}

func InitEmptyTableNameContext(p *TableNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_tableName
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) DdlName() IDdlNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDdlNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDdlNameContext)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *DynamoDbGrammarParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, DynamoDbGrammarParserRULE_tableName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(641)
		p.DdlName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDdlNameContext is an interface to support dynamic dispatch.
type IDdlNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	String_() IStringContext

	// IsDdlNameContext differentiates from other interfaces.
	IsDdlNameContext()
}

type DdlNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDdlNameContext() *DdlNameContext {
	var p = new(DdlNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_ddlName
	return p
}

func InitEmptyDdlNameContext(p *DdlNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_ddlName
}

func (*DdlNameContext) IsDdlNameContext() {}

func NewDdlNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DdlNameContext {
	var p = new(DdlNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_ddlName

	return p
}

func (s *DdlNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DdlNameContext) ID() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserID, 0)
}

func (s *DdlNameContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *DdlNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DdlNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DdlNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterDdlName(s)
	}
}

func (s *DdlNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitDdlName(s)
	}
}

func (p *DynamoDbGrammarParser) DdlName() (localctx IDdlNameContext) {
	localctx = NewDdlNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, DynamoDbGrammarParserRULE_ddlName)
	p.SetState(645)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DynamoDbGrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(643)
			p.Match(DynamoDbGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DynamoDbGrammarParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(644)
			p.String_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserSTRING_LITERAL, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *DynamoDbGrammarParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, DynamoDbGrammarParserRULE_string)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(647)
		p.Match(DynamoDbGrammarParserSTRING_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnknownContext is an interface to support dynamic dispatch.
type IUnknownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUNKNOWN() []antlr.TerminalNode
	UNKNOWN(i int) antlr.TerminalNode

	// IsUnknownContext differentiates from other interfaces.
	IsUnknownContext()
}

type UnknownContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnknownContext() *UnknownContext {
	var p = new(UnknownContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_unknown
	return p
}

func InitEmptyUnknownContext(p *UnknownContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DynamoDbGrammarParserRULE_unknown
}

func (*UnknownContext) IsUnknownContext() {}

func NewUnknownContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnknownContext {
	var p = new(UnknownContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DynamoDbGrammarParserRULE_unknown

	return p
}

func (s *UnknownContext) GetParser() antlr.Parser { return s.parser }

func (s *UnknownContext) AllUNKNOWN() []antlr.TerminalNode {
	return s.GetTokens(DynamoDbGrammarParserUNKNOWN)
}

func (s *UnknownContext) UNKNOWN(i int) antlr.TerminalNode {
	return s.GetToken(DynamoDbGrammarParserUNKNOWN, i)
}

func (s *UnknownContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnknownContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnknownContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.EnterUnknown(s)
	}
}

func (s *UnknownContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DynamoDbGrammarListener); ok {
		listenerT.ExitUnknown(s)
	}
}

func (p *DynamoDbGrammarParser) Unknown() (localctx IUnknownContext) {
	localctx = NewUnknownContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, DynamoDbGrammarParserRULE_unknown)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(650)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DynamoDbGrammarParserUNKNOWN {
		{
			p.SetState(649)
			p.Match(DynamoDbGrammarParserUNKNOWN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(652)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *DynamoDbGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ConditionContext = nil
		if localctx != nil {
			t = localctx.(*ConditionContext)
		}
		return p.Condition_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DynamoDbGrammarParser) Condition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
