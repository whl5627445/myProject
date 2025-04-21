package api

//	type ModelInstance struct {
//		Name                string             `json:"name,omitempty"`
//		ClassDefinitionList []*ClassDefinition `json:"class_definition_list,omitempty"`
//	}
type ClassDefinition struct {
	Final          bool            `json:"final,omitempty"`
	Encapsulated   bool            `json:"encapsulated,omitempty"`
	ClassPrefixes  *ClassPrefixes  `json:"class_prefixes,omitempty"`
	ClassSpecifier *ClassSpecifier `json:"class_specifier,omitempty"`
	Start          int             `json:"start,omitempty"`
	Stop           int             `json:"stop,omitempty"`
	Path           string          `json:"path,omitempty"`
	PackageName    string          `json:"package_name,omitempty"`
	Hash           string          `json:"hash,omitempty"`
}

type ClassPrefixes struct {
	Partial  bool   `json:"partial,omitempty"`
	Prefixes string `json:"prefixes,omitempty"`
}

type ClassSpecifier struct {
	ClassSpecifierType string                `json:"class_specifier_type,omitempty"`
	Name               string                `json:"name,omitempty"`
	Parent             string                `json:"parent,omitempty"`
	Description        string                `json:"description,omitempty"`
	Extends            bool                  `json:"extends,omitempty"`
	ClassModification  []*Argument           `json:"class_modification,omitempty"` // *ClassModification
	Composition        *Composition          `json:"composition,omitempty"`
	AnnotationClause   []*Argument           `json:"annotation,omitempty"`
	BasePrefix         string                `json:"base_prefix,omitempty"`
	TypeSpecifier      string                `json:"specifier,omitempty"`
	ArraySubscripts    []string              `json:"array_subscripts,omitempty"`
	Enumeration        bool                  `json:"enumeration,omitempty"`
	EnumList           []*EnumerationLiteral `json:"enum_list,omitempty"`
	// FullName           string                `json:"full_name,omitempty"`
	// Hash               string                `json:"hash,omitempty"`
}

type LongClassSpecifier struct {
	Name              string       `json:"name,omitempty"`
	FullName          string       `json:"full_name,omitempty"`
	Hash              string       `json:"hash,omitempty"`
	Parent            string       `json:"parent,omitempty"`
	Description       string       `json:"description,omitempty"`
	Extends           bool         `json:"extends,omitempty"`
	ClassModification []*Argument  `json:"class_modification,omitempty"` // *ClassModification
	Composition       *Composition `json:"composition,omitempty"`
}

type ShortClassSpecifier struct {
	Name              string                `json:"name,omitempty"`
	FullName          string                `json:"full_name,omitempty"`
	Hash              string                `json:"hash,omitempty"`
	Parent            string                `json:"parent,omitempty"`
	Description       string                `json:"description,omitempty"`
	AnnotationClause  []*Argument           `json:"annotation,omitempty"`
	BasePrefix        string                `json:"base_prefix,omitempty"`
	TypeSpecifier     string                `json:"specifier,omitempty"`
	ArraySubscripts   []string              `json:"array_subscripts,omitempty"`
	Enumeration       bool                  `json:"enumeration,omitempty"`
	EnumList          []*EnumerationLiteral `json:"enum_list,omitempty"`
	ClassModification []*Argument           `json:"class_modification,omitempty"` // *ClassModification
}

type DerClassSpecifier struct {
	Name             string      `json:"name,omitempty"`
	FullName         string      `json:"full_name,omitempty"`
	Hash             string      `json:"hash,omitempty"`
	Parent           string      `json:"parent,omitempty"`
	TypeSpecifier    string      `json:"specifier,omitempty"`
	Description      string      `json:"description,omitempty"`
	AnnotationClause []*Argument `json:"annotation,omitempty"`
}

type EnumerationLiteral struct {
	Name        string       `json:"name,omitempty"`
	Description *Description `json:"description,omitempty"`
}

type Composition struct {
	ElementList           []*Element             `json:"elements,omitempty"`
	PublicElementList     []*Element             `json:"public_element_list,omitempty"`
	ProtectedElementList  []*Element             `json:"protected_element_list,omitempty"`
	EquationSection       []*EquationSection     `json:"equation_section,omitempty"`
	AlgorithmSection      []*AlgorithmSection    `json:"algorithm_section,omitempty"`    // *AlgorithmSection
	ExternalComposition   []*ExternalComposition `json:"external_composition,omitempty"` // *ExternalComposition
	LanguageSpecification string                 `json:"language_specification,omitempty"`
	ExternalFunctionCall  *ExternalFunctionCall  `json:"external_function_call,omitempty"` // *ExternalFunctionCall
	AnnotationClause      []*Argument            `json:"annotation,omitempty"`             // *AnnotationClause
}

type ExternalComposition struct {
	LanguageSpecification string                `json:"language_specification,omitempty"`
	ExternalFunctionCall  *ExternalFunctionCall `json:"external_function_call,omitempty"`
	AnnotationClause      []*Argument           `json:"annotation,omitempty"` // *AnnotationClause
}

type ExternalFunctionCall struct {
	ComponentReference string   `json:"component_reference,omitempty"`
	Name               string   `json:"name,omitempty"`
	ExpressionList     []string `json:"expression_list,omitempty"`
}

type Element struct {
	ImportClause                   bool                            `json:"import_clause,omitempty"`
	ImportClauseName               string                          `json:"import_clause_name,omitempty"`
	ExtendsClause                  bool                            `json:"extends,omitempty"`
	TypeSpecifier                  string                          `json:"specifier,omitempty"`
	InheritanceModification        []*InheritanceModification      `json:"inheritance_modification,omitempty"`
	AnnotationClause               []*Argument                     `json:"annotation,omitempty"`
	Binding                        string                          `json:"binding,omitempty"`
	Redeclare                      bool                            `json:"redeclare,omitempty"`
	Final                          bool                            `json:"final,omitempty"`
	Inner                          bool                            `json:"inner,omitempty"`
	Outer                          bool                            `json:"outer,omitempty"`
	ClassDefinition                *ClassDefinition                `json:"class_definition,omitempty"`
	TypePrefix                     string                          `json:"prefix,omitempty"`
	ArraySubscripts                []string                        `json:"array_subscripts,omitempty"`
	Name                           string                          `json:"name,omitempty"`
	ConditionAttribute             string                          `json:"condition_attribute,omitempty"`
	Replaceable                    bool                            `json:"replaceable,omitempty"`
	ConstrainingClause             *ConstrainingClause             `json:"constraining_clause,omitempty"`
	Constrainedby                  bool                            `json:"constrainedby,omitempty"`
	ClassModification              []*Argument                     `json:"class_modification,omitempty"`
	DescriptionString              string                          `json:"description,omitempty"`
	Modification                   *Modification                   `json:"modification,omitempty"` // Modification
	FullName                       string                          `json:"full_name,omitempty"`
	Hash                           string                          `json:"hash,omitempty"`
	Parent                         string                          `json:"parent,omitempty"`
	ClassOrInheritanceModification *ClassOrInheritanceModification `json:"class_or_inheritance_modification,omitempty"`
	Instance                       *ClassDefinition                `json:"instance,omitempty"`
}

type ImportClause struct {
	Import      bool         `json:"import,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description *Description `json:"description,omitempty"`
}

type ExtendsClause struct {
	Extends                        bool                            `json:"extends,omitempty"`
	TypeSpecifier                  string                          `json:"specifier,omitempty"`
	ClassOrInheritanceModification *ClassOrInheritanceModification `json:"class_or_inheritance_modification,omitempty"`
	AnnotationClause               *AnnotationClause               `json:"annotation,omitempty"`
}

type ConstrainingClause struct {
	Constrainedby     bool        `json:"constrainedby,omitempty"`
	TypeSpecifier     string      `json:"specifier,omitempty"`
	ClassModification []*Argument `json:"class_modification,omitempty"`
}

type ClassOrInheritanceModification struct {
	ArgumentList            []*Argument                `json:"argument,omitempty"` // Argument
	InheritanceModification []*InheritanceModification `json:"inheritance_modification,omitempty"`
}

type InheritanceModification struct {
	Break           bool        `json:"break,omitempty"`
	ConnectEquation *Connection `json:"connect_equation,omitempty"`
	Name            string      `json:"name,omitempty"`
}
type Connection struct {
	Left  string `json:"left,omitempty"`
	Right string `json:"right,omitempty"`
}

type ConnectionAnnotation struct {
	Line map[string]any `json:"line,omitempty"`
}

type ComponentClause struct {
	TypePrefix      string                  `json:"prefix,omitempty"`
	TypeSpecifier   string                  `json:"specifier,omitempty"`
	ArraySubscripts []string                `json:"array_subscripts,omitempty"`
	ComponentList   []*ComponentDeclaration `json:"component_list,omitempty"`
}

type ComponentList struct {
	ComponentList []*ComponentDeclaration `json:"component_list,omitempty"`
}

type ComponentDeclaration struct {
	ArraySubscripts    []string      `json:"array_subscripts,omitempty"`
	Modification       *Modification `json:"modification,omitempty"` // Modification
	Name               string        `json:"name,omitempty"`
	DescriptionString  string        `json:"description,omitempty"`
	AnnotationClause   []*Argument   `json:"annotation,omitempty"`
	ConditionAttribute string        `json:"condition_attribute,omitempty"`
}

type Declaration struct {
	ArraySubscripts []string      `json:"array_subscripts,omitempty"`
	Modification    *Modification `json:"modification,omitempty"` // Modification
	Name            string        `json:"name,omitempty"`
}

type Modification struct {
	ClassModification      []*Argument             `json:"class_modification,omitempty"`
	ModificationExpression *ModificationExpression `json:"modification_expression,omitempty"`
	Expression             string                  `json:"expression,omitempty"`
}

type ModificationExpression struct {
	Expression string `json:"expression,omitempty"`
	Break      bool   `json:"break,omitempty"`
}
type ClassModification struct {
	ArgumentList []*Argument `json:"argument_list,omitempty"`
}
type Argument struct {
	Each                             bool                              `json:"each,omitempty"`
	Final                            bool                              `json:"final,omitempty"`
	Name                             string                            `json:"name,omitempty"`
	Data                             any                               `json:"data,omitempty"`
	Binding                          string                            `json:"binding,omitempty"`
	Children                         []*Argument                       `json:"children,omitempty"`
	ElementReplaceable               *ElementReplaceable               `json:"element_replaceable,omitempty"`
	ElementModification              *ElementModification              `json:"element_modification,omitempty"`
	ElementRedeclaration             *ElementRedeclaration             `json:"element_redeclaration,omitempty"`
	ElementModificationOrReplaceable *ElementModificationOrReplaceable `json:"element_modification_or_replaceable,omitempty"`
}

type ElementModificationOrReplaceable struct {
	Each                bool                 `json:"each,omitempty"`
	Final               bool                 `json:"final,omitempty"`
	Name                string               `json:"name,omitempty"`
	ElementModification *ElementModification `json:"element_modification,omitempty"`
	ElementReplaceable  *ElementReplaceable  `json:"element_replaceable,omitempty"`
}

type ElementModification struct {
	Name              string        `json:"name,omitempty"`
	Modification      *Modification `json:"modification,omitempty"`
	DescriptionString string        `json:"description,omitempty"`
}
type ElementRedeclaration struct {
	Redeclare           bool                 `json:"redeclare,omitempty"`
	Each                bool                 `json:"each,omitempty"`
	Final               bool                 `json:"final,omitempty"`
	ClassPrefixes       *ClassPrefixes       `json:"class_prefixes,omitempty"`
	ShortClassSpecifier *ShortClassSpecifier `json:"short_class_specifier,omitempty"` // ShortClassSpecifier
	ComponentClause1    *ComponentClause1    `json:"component_clause_1,omitempty"`
	ElementReplaceable  *ElementReplaceable  `json:"element_replaceable,omitempty"`
}

type ElementReplaceable struct {
	Replaceable         bool                 `json:"replaceable,omitempty"`
	ClassPrefixes       *ClassPrefixes       `json:"class_prefixes,omitempty"`
	ShortClassSpecifier *ShortClassSpecifier `json:"short_class_specifier,omitempty"` // ShortClassSpecifier
	TypePrefix          string               `json:"prefix,omitempty"`
	TypeSpecifier       string               `json:"specifier,omitempty"`
	Declaration         *Declaration         `json:"declaration,omitempty"`
	Description         *Description         `json:"description,omitempty"`
	Constrainedby       bool                 `json:"constrainedby,omitempty"`
	ClassModification   []*Argument          `json:"class_modification,omitempty"`
}

type ComponentClause1 struct {
	TypePrefix    string       `json:"prefix,omitempty"`
	TypeSpecifier string       `json:"specifier,omitempty"`
	Declaration   *Declaration `json:"declaration,omitempty"`
	Description   *Description `json:"description,omitempty"`
}

type ComponentDeclaration1 struct {
	Declaration *Declaration `json:"declaration,omitempty"`
	Description *Description `json:"description,omitempty"`
}

type ShortClassDefinition struct {
	ClassPrefixes       *ClassPrefixes `json:"class_prefixes,omitempty"`
	ShortClassSpecifier string         `json:"short_class_specifier,omitempty"` // ShortClassSpecifier
}

type EquationSection struct {
	Initial      bool        `json:"initial,omitempty"`
	EquationList []*Equation `json:"equation_list,omitempty"`
}

type AlgorithmSection struct {
	Initial       bool     `json:"initial,omitempty"`
	StatementList []string `json:"statement_list,omitempty"` // []*Statement
}

type Equation struct {
	EquationName      string      `json:"equation_name,omitempty"`
	Connection        *Connection `json:"connection,omitempty"`
	EquationString    string      `json:"equation_string,omitempty"`
	DescriptionString string      `json:"description,omitempty"`
	Annotation        []*Argument `json:"annotation,omitempty"`
}

type Statement struct {
	ComponentReference   string                `json:"component_reference,omitempty"`
	Expression           Expression            `json:"expression,omitempty"`
	FunctionCallArgs     *FunctionCallArgs     `json:"function_call_args,omitempty"`
	OutputExpressionList *OutputExpressionList `json:"output_expression_list,omitempty"`
	Break                bool                  `json:"break,omitempty"`
	Return               bool                  `json:"return,omitempty"`
	IfStatement          string                `json:"if_statement,omitempty"`
	ForStatement         string                `json:"for_statement,omitempty"`
	WhileStatement       string                `json:"while_statement,omitempty"`
	WhenStatement        string                `json:"when_statement,omitempty"`
	Description          *Description          `json:"description,omitempty"`
}

type ComponentReference struct {
	ArraySubscripts *ArraySubscripts `json:"array_subscripts,omitempty"`
}

type FunctionCallArgs struct {
}

type ArraySubscripts struct {
	Subscripts []any `json:"subscripts,omitempty"`
}

type Description struct {
	DescriptionString string      `json:"description,omitempty"`
	AnnotationClause  []*Argument `json:"annotation,omitempty"`
}

type AnnotationClause struct {
	ClassModification []*Argument `json:"class_modification,omitempty"`
}

type Expression struct {
	SimpleExpression *SimpleExpression `json:"simple_expression,omitempty"`
	IfElseExpression *IfElseExpression `json:"if_else_expression,omitempty"`
}

type SimpleExpression struct {
	LogicalExpression []*LogicalExpression `json:"logical_expression,omitempty"`
}

type IfElseExpression struct {
	If         string                  `json:"if_expression,omitempty"`
	Then       string                  `json:"then_expression,omitempty"`
	ElseIfThen []*ElseIfThenExpression `json:"else_if_then_expression,omitempty"`
	Else       string                  `json:"else_expression,omitempty"`
}

type ElseIfThenExpression struct {
	ElseIf string `json:"else_if_expression,omitempty"`
	Then   string `json:"then_expression,omitempty"`
}

type LogicalExpression struct {
	LogicalTerm []*LogicalTerm `json:"logical_term,omitempty"`
}

type LogicalTerm struct {
	LogicalFactor []*LogicalFactor `json:"logical_factor,omitempty"`
}

type LogicalFactor struct {
	Relation *Relation `json:"relation,omitempty"`
	Not      bool      `json:"not,omitempty"`
}

type Relation struct {
	ArithmeticExpression           *ArithmeticExpression           `json:"arithmetic_expression,omitempty"`
	RelationalArithmeticExpression *RelationalArithmeticExpression `json:"relational_arithmetic_expression,omitempty"`
}

type RelationalArithmeticExpression struct {
	RelationalOperator   string                `json:"relational_operator,omitempty"`
	ArithmeticExpression *ArithmeticExpression `json:"arithmetic_expression,omitempty"`
}

type ArithmeticExpression struct {
	AddOperator string     `json:"add_operator,omitempty"`
	Term        *Term      `json:"term,omitempty"`
	AddTerm     []*AddTerm `json:"add_term,omitempty"`
}

type AddTerm struct {
	AddOperator string `json:"add_operator,omitempty"`
	Term        *Term  `json:"term,omitempty"`
}

type Term struct {
	Factor    *Factor      `json:"factor,omitempty"`
	MulFactor []*MulFactor `json:"mul_factor,omitempty"`
}
type MulFactor struct {
	Factor      *Factor `json:"factor,omitempty"`
	MulOperator string  `json:"mul_operator,omitempty"`
}

type Factor struct {
	Primary            *Primary            `json:"Primary,omitempty"`
	ExponentialPrimary *ExponentialPrimary `json:"exponential_primary,omitempty"`
}

type ExponentialPrimary struct {
	ExponentialOperator string   `json:"exponential_operator,omitempty"`
	Primary             *Primary `json:"Primary,omitempty"`
}

type Primary struct {
	Text                   string                  `json:"text,omitempty"`
	UnsignedNumber         string                  `json:"UNSIGNED_NUMBER,omitempty"`
	STRING                 string                  `json:"STRING,omitempty"`
	False                  bool                    `json:"false,omitempty"`
	True                   bool                    `json:"true,omitempty"`
	CdipFunctionCallArgs   *CdipFunctionCallArgs   `json:"cdip_function_call_args,omitempty"`
	ComponentReference     string                  `json:"component_reference,omitempty"`
	OutputExpressionList   *OutputExpressionList   `json:"output_expression_list,omitempty"`
	ExpressionListMultiple *ExpressionListMultiple `json:"expression_list__multiple,omitempty"`
	ArrayArguments         *ArrayArguments         `json:"array_arguments,omitempty"`
}

type CdipFunctionCallArgs struct {
	ComponentReference string             `json:"component_reference,omitempty"`
	Der                bool               `json:"der,omitempty"`
	Initial            bool               `json:"initial,omitempty"`
	Pure               bool               `json:"pure,omitempty"`
	FunctionCallArgs   *FunctionArguments `json:"function_arguments,omitempty"`
}

type ExpressionListMultiple struct {
	ExpressionList [][]string `json:"expression_list,omitempty"`
}

type FunctionArguments struct {
	ExpressionFunctionArgumentsNonFirstOrForIndices     *ExpressionFunctionArgumentsNonFirstOrForIndices     `json:"expression__function_arguments_non_first_or_for_indices,omitempty"`
	FunctionPartialApplicationFunctionArgumentsNonFirst *FunctionPartialApplicationFunctionArgumentsNonFirst `json:"function_partial_application__function_arguments_non_first,omitempty"`
	NamedArguments                                      *NamedArguments                                      `json:"named_arguments,omitempty"`
}

type ExpressionFunctionArgumentsNonFirstOrForIndices struct {
	Expression                string                     `json:"expression,omitempty"`
	FunctionArgumentsNonFirst *FunctionArgumentsNonFirst `json:"function_arguments_non_first,omitempty"`
	ForIndices                *ForIndices                `json:"for_indices,omitempty"`
}
type FunctionPartialApplicationFunctionArgumentsNonFirst struct {
	FunctionPartialApplication *FunctionPartialApplication `json:"function_partial_application,omitempty"`
	FunctionArgumentsNonFirst  *FunctionArgumentsNonFirst  `json:"function_arguments_non_first,omitempty"`
}

type FunctionArgumentsNonFirst struct {
	FunctionArgumentsOrNonFirst *FunctionArgumentsOrNonFirst `json:"function_arguments_or_non_first,omitempty"`
	NamedArguments              *NamedArguments              `json:"named_arguments,omitempty"`
}
type FunctionArgumentsOrNonFirst struct {
	FunctionArgument          *FunctionArgument          `json:"function_argument,omitempty"`
	FunctionArgumentsNonFirst *FunctionArgumentsNonFirst `json:"function_arguments_non_first,omitempty"`
}

type ArrayArguments struct {
	Expression             string                  `json:"expression,omitempty"`
	ArrayArgumentsNonFirst *ArrayArgumentsNonFirst `json:"array_arguments_non_first,omitempty"`
	ForIndices             *ForIndices             `json:"for_indices,omitempty"`
}

type ArrayArgumentsNonFirst struct {
	Expression             string                  `json:"expression,omitempty"`
	ArrayArgumentsNonFirst *ArrayArgumentsNonFirst `json:"array_arguments_non_first,omitempty"`
}

type NamedArguments struct {
	IDENT            []string            `json:"IDENT,omitempty"`
	FunctionArgument []*FunctionArgument `json:"function_argument,omitempty"`
}

type FunctionArgument struct {
	FunctionPartialApplication *FunctionPartialApplication `json:"function_partial_application,omitempty"`
	Expression                 string                      `json:"expression,omitempty"`
}

type FunctionPartialApplication struct {
	TypeSpecifier  string          `json:"type_specifier,omitempty"`
	NamedArguments *NamedArguments `json:"named_arguments,omitempty"`
}

type ForIndices struct {
	IDENTList      []string `json:"IDENT,omitempty"`
	ExpressionList []string `json:"expression,omitempty"`
}

type OutputExpressionList struct {
	Expression []string `json:"expression,omitempty"`
}
