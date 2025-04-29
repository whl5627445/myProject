package api

// TODO：该文件数据结构暂时没有用
// TODO：该文件数据结构暂时没有用到

type ClassDefinition struct {
	Final          bool            `json:"final,omitempty"`           // final关键字标记
	Encapsulated   bool            `json:"encapsulated,omitempty"`    // encapsulated关键字标记
	ClassPrefixes  *ClassPrefixes  `json:"class_prefixes,omitempty"`  // 模型前缀
	ClassSpecifier *ClassSpecifier `json:"class_specifier,omitempty"` // 模型主要部分的数据
	WithinName     string          `json:"within_name,omitempty"`     // within关键字后标记的所属包名
	PackageName    string          `json:"package_name,omitempty"`    // 模型所属哪个顶层包名
	Hash           string          `json:"hash,omitempty"`
}

type ClassPrefixes struct {
	Partial  bool   `json:"partial,omitempty"`  // partial关键字标记
	Prefixes string `json:"prefixes,omitempty"` // 模型前缀。类似model、package、record
}

type ClassSpecifier struct {
	ClassSpecifierType string                `json:"class_specifier_type,omitempty"` // 模型说明符类型、 有LongClassSpecifier、ShortClassSpecifier两类
	Name               string                `json:"name,omitempty"`                 // 模型短类名，类似于Modelica.Blocks.Examples.PID_Controller中的PID_Controller名称
	Parent             string                `json:"parent,omitempty"`               // 模型的父节点
	Description        string                `json:"description,omitempty"`          // 模型的描述
	Extends            bool                  `json:"extends,omitempty"`              // 标记是否是继承项
	ClassModification  []*Argument           `json:"class_modification,omitempty"`   // 模型与组件修饰符列表
	Composition        *Composition          `json:"composition,omitempty"`          // 模型组成，一般是组件、方程、算法、annotation等
	AnnotationClause   []*Argument           `json:"annotation,omitempty"`           // 模型注解，包含图形信息、图标信息、文档以及实验等信息
	BasePrefix         string                `json:"base_prefix,omitempty"`          // 表示output、input等
	TypeSpecifier      string                `json:"specifier,omitempty"`            // 表示模型的基本类型
	ArraySubscripts    []string              `json:"array_subscripts,omitempty"`     // 数组表示
	Enumeration        bool                  `json:"enumeration,omitempty"`          // 是否是枚举类型
	EnumList           []*EnumerationLiteral `json:"enum_list,omitempty"`            // 枚举类型的选项列表
	Hash               string                `json:"hash,omitempty"`
}

type ShortClassSpecifier struct {
	Name              string                `json:"name,omitempty"`               // 模型短类名，类似于Modelica.Blocks.Examples.PID_Controller中的PID_Controller名称
	Parent            string                `json:"parent,omitempty"`             // 模型的父节点
	Description       string                `json:"description,omitempty"`        // 模型的描述
	AnnotationClause  []*Argument           `json:"annotation,omitempty"`         // 模型注解，包含图形信息、图标信息、文档以及实验等信息
	BasePrefix        string                `json:"base_prefix,omitempty"`        // 表示output、input等
	TypeSpecifier     string                `json:"specifier,omitempty"`          // 表示模型的基本类型
	ArraySubscripts   []string              `json:"array_subscripts,omitempty"`   // 数组表示
	Enumeration       bool                  `json:"enumeration,omitempty"`        // 是否是枚举类型
	EnumList          []*EnumerationLiteral `json:"enum_list,omitempty"`          // 枚举类型的选项列表
	ClassModification []*Argument           `json:"class_modification,omitempty"` // 模型与组件修饰符列表
	Hash              string                `json:"hash,omitempty"`
	FullName          string                `json:"-"`
}

type EnumerationLiteral struct {
	Name        string       `json:"name,omitempty"`        // 枚举选项名称
	Description *Description `json:"description,omitempty"` // 枚举选项的描述
}

type Composition struct {
	ElementList          []*Element         `json:"elements,omitempty"`               // 模型组件列表
	PublicElementList    []*Element         `json:"public_element_list,omitempty"`    // 公开的组件
	ProtectedElementList []*Element         `json:"protected_element_list,omitempty"` // 受保护的组件
	EquationSection      []*EquationSection `json:"equation_section,omitempty"`       // 方程列表
	AnnotationClause     []*Argument        `json:"annotation,omitempty"`             // 模型注解，包含图形信息、图标信息、文档以及实验等信息
}

type Element struct {
	ImportClause                   bool                            `json:"import_clause,omitempty"`                     // 标记是否是导入的包
	ImportClauseName               string                          `json:"import_clause_name,omitempty"`                // 导入包的名称
	ExtendsClause                  bool                            `json:"extends,omitempty"`                           // 标记是否是继承项
	TypeSpecifier                  string                          `json:"specifier,omitempty"`                         // 表示模型的基本类型
	InheritanceModification        []*InheritanceModification      `json:"inheritance_modification,omitempty"`          // 继承的修改项
	AnnotationClause               []*Argument                     `json:"annotation,omitempty"`                        // 注解，包含图形信息、图标信息、坐标以及一些分组等信息
	Binding                        string                          `json:"binding,omitempty"`                           // 绑定的值
	Redeclare                      bool                            `json:"redeclare,omitempty"`                         // redeclare关键字标记
	Final                          bool                            `json:"final,omitempty"`                             // final关键字标记
	Inner                          bool                            `json:"inner,omitempty"`                             // inner关键字标记
	Outer                          bool                            `json:"outer,omitempty"`                             // outer关键字标记
	ClassDefinition                *ClassDefinition                `json:"class_definition,omitempty"`                  // 模型中子模型项
	TypePrefix                     string                          `json:"prefix,omitempty"`                            // 前缀
	ArraySubscripts                []string                        `json:"array_subscripts,omitempty"`                  // 数组表示
	Name                           string                          `json:"name,omitempty"`                              // 名称
	ConditionAttribute             string                          `json:"condition_attribute,omitempty"`               // 条件表达式，一般用于表达是否启用
	Replaceable                    bool                            `json:"replaceable,omitempty"`                       // replaceable关键字标记
	ConstrainingClause             *ConstrainingClause             `json:"constraining,omitempty"`                      // 约束条件
	ClassModification              []*Argument                     `json:"class_modification,omitempty"`                // 修饰符列表
	DescriptionString              string                          `json:"description,omitempty"`                       // 描述
	Modification                   *Modification                   `json:"modification,omitempty"`                      // 表示修改项
	Parent                         string                          `json:"parent,omitempty"`                            // 模型的父节点
	Instance                       *ClassDefinition                `json:"instance,omitempty"`                          // 模型实例数据
	ClassOrInheritanceModification *ClassOrInheritanceModification `json:"class_or_inheritance_modification,omitempty"` // 类或继承项的修改项
	Hash                           string                          `json:"hash,omitempty"`
	FullName                       string                          `json:"-"`
}

type ConstrainingClause struct {
	Constrainedby     bool        `json:"constrainedby,omitempty"`      // constrainedby关键字标记
	TypeSpecifier     string      `json:"specifier,omitempty"`          // 表示基本类型
	ClassModification []*Argument `json:"class_modification,omitempty"` // 类修改项
}

type ClassOrInheritanceModification struct {
	ArgumentList            []*Argument                `json:"argument_list,omitempty"`
	InheritanceModification []*InheritanceModification `json:"inheritance_modification,omitempty"`
}

type InheritanceModification struct {
	Break           bool        `json:"break,omitempty"`
	ConnectEquation *Connection `json:"connect_equation,omitempty"`
	Name            string      `json:"name,omitempty"`
}
type Connection struct {
	Left  string `json:"left,omitempty"`  // 起始节点
	Right string `json:"right,omitempty"` // 终止节点
}

type Declaration struct {
	ArraySubscripts []string      `json:"array_subscripts,omitempty"` // 数组表示
	Modification    *Modification `json:"modification,omitempty"`     // 修改项
	Name            string        `json:"name,omitempty"`             // 名称
}

type Modification struct {
	ClassModification      []*Argument             `json:"class_modification,omitempty"`      // 类修改项
	ModificationExpression *ModificationExpression `json:"modification_expression,omitempty"` // 修改项表达式
	Expression             string                  `json:"expression,omitempty"`              // 表达式
}

type ModificationExpression struct {
	Expression string `json:"expression,omitempty"` // 表达式
	Break      bool   `json:"break,omitempty"`      // 标记break关键字
}

type Argument struct {
	Each                 bool                  `json:"each,omitempty"`                  // 标记each关键字
	Final                bool                  `json:"final,omitempty"`                 // 标记final关键字
	Name                 string                `json:"name,omitempty"`                  // 名称
	Data                 any                   `json:"data,omitempty"`                  // 绑定的数据对象
	Binding              string                `json:"binding,omitempty"`               // 绑定的数据
	Children             []*Argument           `json:"children,omitempty"`              // 子节点，结构完全一致
	ElementReplaceable   *ElementReplaceable   `json:"element_replaceable,omitempty"`   // 可更换组件
	ElementModification  *ElementModification  `json:"element_modification,omitempty"`  // 元素修改项
	ElementRedeclaration *ElementRedeclaration `json:"element_redeclaration,omitempty"` // 元素重新声明
}

type ElementModification struct {
	Name              string        `json:"name,omitempty"`         // 名称
	Modification      *Modification `json:"modification,omitempty"` // 表示修改项
	DescriptionString string        `json:"description,omitempty"`  // 描述
}
type ElementRedeclaration struct {
	Redeclare           bool                 `json:"redeclare,omitempty"`             // redeclare关键字标记
	Each                bool                 `json:"each,omitempty"`                  // each关键字标记
	Final               bool                 `json:"final,omitempty"`                 // final关键字标记
	ClassPrefixes       *ClassPrefixes       `json:"class_prefixes,omitempty"`        // 前缀
	ShortClassSpecifier *ShortClassSpecifier `json:"short_class_specifier,omitempty"` // 短类数据
	ComponentClause1    *ComponentClause1    `json:"component_clause_1,omitempty"`    // 组件子语句，类似表达式
	ElementReplaceable  *ElementReplaceable  `json:"element_replaceable,omitempty"`   // 组件替换
}

type ElementReplaceable struct {
	Replaceable         bool                 `json:"replaceable,omitempty"`           // replaceable关键字标记
	ClassPrefixes       *ClassPrefixes       `json:"class_prefixes,omitempty"`        // 前缀
	ShortClassSpecifier *ShortClassSpecifier `json:"short_class_specifier,omitempty"` // 短类数据
	TypePrefix          string               `json:"prefix,omitempty"`                // 类型前缀
	TypeSpecifier       string               `json:"specifier,omitempty"`             // 说明符
	Declaration         *Declaration         `json:"declaration,omitempty"`           // 声明
	Description         *Description         `json:"description,omitempty"`           // 描述
	Constrainedby       bool                 `json:"constrainedby,omitempty"`         // constrainedby关键字标记
	ClassModification   []*Argument          `json:"class_modification,omitempty"`    // 类修改项
}

type ComponentClause1 struct {
	TypePrefix    string       `json:"prefix,omitempty"`      // 类型前缀
	TypeSpecifier string       `json:"specifier,omitempty"`   // 说明符
	Declaration   *Declaration `json:"declaration,omitempty"` // 声明
	Description   *Description `json:"description,omitempty"` // 描述
}

type EquationSection struct {
	Initial      bool        `json:"initial,omitempty"`       // initial关键字标记
	EquationList []*Equation `json:"equation_list,omitempty"` // 方程列表
}

type Equation struct {
	EquationName      string      `json:"equation_name,omitempty"` // 方程名称
	Connection        *Connection `json:"connection,omitempty"`    // 方程连接
	DescriptionString string      `json:"description,omitempty"`   // 描述
	Annotation        []*Argument `json:"annotation,omitempty"`    // 注解
}

type Description struct {
	DescriptionString string      `json:"description,omitempty"` // 描述
	AnnotationClause  []*Argument `json:"annotation,omitempty"`  // 注解
}
