package DataBaseModel

import (
	"time"

	"gorm.io/datatypes"

	"gorm.io/gorm"
)

type YssimModels struct {
	ID             string         `gorm:"index;primaryKey;type:varchar(128);comment:package唯一标识"`
	LibraryId      string         `gorm:"index;column:library_id;type:varchar(128);default:\"\";comment:library库唯一标识"`
	PackageName    string         `gorm:"index;column:package_name;type:varchar(128);comment:package名称，一般称为包名或库的名字"`
	Version        string         `gorm:"column:version;default:\"\";type:varchar(32);comment:package版本号"`
	SysUser        string         `gorm:"index;column:sys_or_user;default:\"\";type:varchar(32);comment:是用户模型的话则为用户名，系统模型则是sys"`
	FilePath       string         `gorm:"column:file_path;default:\"\";type:varchar(512);comment:用户模型所在路径，常驻系统模型则为空"`
	CreatedAt      *time.Time     `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdatedAt      *time.Time     `gorm:"column:update_time;comment:更新时间"`
	UserSpaceId    string         `gorm:"column:userspace_id;default:\"\";type:varchar(128);comment:package所在用户空间的唯一识别标识"`
	VersionControl bool           `gorm:"column:version_control;default:false;type:bool;comment:是否有版本控制"`
	VersionBranch  string         `gorm:"column:version_branch;default:\"\";type:varchar(128);comment:版本控制分支"`
	VersionTag     string         `gorm:"column:version_tag;default:\"\";type:varchar(128);comment:版本控制tag"`
	Deleted        gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
	Default        bool           `gorm:"column:default_version;default:0;type:bool;comment:是否默认加载"`
	Encryption     bool           `gorm:"column:encryption;default:0;type:bool;comment:是否是加密模型"`
	//PipeNet        bool           `gorm:"column:pipe_net;default:0;type:bool;comment:是否是管网模型"`
}

type SystemLibrary struct {
	ID             string         `gorm:"index;primaryKey;type:varchar(128);comment:package唯一标识"`
	UserName       string         `gorm:"index;column:username;type:varchar(128);comment:用户名"`
	PackageName    string         `gorm:"index;column:package_name;type:varchar(128);comment:package名称，一般称为包名或库的名字"`
	Version        string         `gorm:"column:version;default:\"\";type:varchar(32);comment:package版本号"`
	FilePath       string         `gorm:"column:file_path;default:\"\";type:varchar(256);comment:package所在路径"`
	VersionControl bool           `gorm:"column:version_control;default:false;type:bool;comment:是否有版本控制"`
	VersionBranch  string         `gorm:"column:version_branch;default:\"\";type:varchar(128);comment:版本控制分支"`
	VersionTag     string         `gorm:"column:version_tag;default:\"\";type:varchar(128);comment:版本控制tag"`
	CreatedAt      *time.Time     `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdatedAt      *time.Time     `gorm:"column:update_time;comment:更新时间"`
	Deleted        gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
	Encryption     bool           `gorm:"column:encryption;default:0;type:bool;comment:是否是加密模型"`
}

type UserLibrary struct {
	ID                string         `gorm:"index;primaryKey;type:varchar(128);comment:package唯一标识"`
	UserName          string         `gorm:"index;column:username;type:varchar(128);comment:用户名"`
	PackageName       string         `gorm:"index;column:package_name;type:varchar(128);comment:package名称，一般称为包名或库的名字"`
	Version           string         `gorm:"column:version;default:\"\";type:varchar(32);comment:package版本号"`
	Used              bool           `gorm:"column:used;type:bool;comment:该package是否已经被某空间使用"`
	FilePath          string         `gorm:"column:file_path;default:\"\";type:varchar(256);comment:package所在路径"`
	VersionControl    bool           `gorm:"column:version_control;default:false;type:bool;comment:是否有版本控制"`
	VersionBranch     string         `gorm:"column:version_branch;default:\"master\";type:varchar(128);comment:版本控制分支"`
	VersionTag        string         `gorm:"column:version_tag;default:\"\";type:varchar(128);comment:版本控制tag"`
	RepositoryAddress string         `gorm:"column:repository_address;default:\"\";type:varchar(256);comment:git地址"`
	AnotherName       string         `gorm:"column:another_name;default:\"\";type:varchar(128);comment:别名"`
	CreatedAt         *time.Time     `gorm:"column:create_time;autoCreateTime;comment:创建时间"`
	UpdatedAt         *time.Time     `gorm:"column:update_time;comment:更新时间"`
	Deleted           gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间"`
}

type ParameterCalibrationRecord struct {
	ID                      string         `gorm:"index;primaryKey;type:varchar(128);comment:参数标定记录的唯一标识"`
	PackageId               string         `gorm:"index;column:package_id;type:varchar(128);comment:package唯一标识"`
	Version                 string         `gorm:"column:version;default:\"\";type:varchar(32);comment:package版本号"`
	UserName                string         `gorm:"index;column:username;type:varchar(128);comment:用户名"`
	ModelName               string         `gorm:"column:model_name;type:varchar(256);comment:模型名称"`
	PackagePath             string         `gorm:"column:package_path;type:varchar(256);comment:标定模型所在包的加载文件"`
	CompileDependencies     datatypes.JSON `gorm:"column:compile_Dependencies;type:json;comment:编译时所需的包环境"`
	CompilePath             string         `gorm:"column:compile_path;type:varchar(128);comment:编译好的文件存放路径"`
	CompileStatus           string         `gorm:"column:compile_status;default:\"0\";type:varchar(32);comment:编译结果状态码，0(初始状态)、3(失败)、4(成功)、6(编译中)"`
	CompileStartTime        int64          `gorm:"column:compile_start_time;type:int;comment:编译开始时间"`
	CompileStopTime         int64          `gorm:"column:compile_stop_time;type:int;comment:编译结束时间"`
	SimulateModelResultPath string         `gorm:"column:simulate_model_result_path;type:varchar(256);comment:仿真结果存储路径"`
	SimulateResultStr       string         `gorm:"column:simulate_result_str;comment:仿真结果输出字符串"`
	SimulateStatus          string         `gorm:"column:simulate_status;default:\"0\";type:varchar(32);comment:仿真状态"`
	StartTime               string         `gorm:"column:start_time;type:varchar(32);comment:仿真开始时间"`
	StopTime                string         `gorm:"column:stop_time;type:varchar(32);comment:仿真结束时间"`
	Tolerance               string         `gorm:"column:tolerance;type:varchar(32);comment:仿真积分误差"`
	NumberOfIntervals       string         `gorm:"column:number_of_intervals;type:varchar(32);comment:仿真间隔数"`
	Interval                string         `gorm:"column:interval;type:varchar(32);comment:仿真时间间隔"`
	Method                  string         `gorm:"column:method;type:varchar(32);comment:仿真积分方法"`
	Percentage              datatypes.JSON `gorm:"column:percentage;type:json;comment:仿真进度, 每次仿真有单独状态"`
	ActualData              datatypes.JSON `gorm:"column:actual_data;type:json;comment:实测数据字段与数据"`
	SimulateResult          datatypes.JSON `gorm:"column:simulate_result;type:json;comment:仿真结果数据"`
	RatedCondition          datatypes.JSON `gorm:"column:rated_condition;type:json;comment:额定工况参数信息"`
	ConditionParameters     datatypes.JSON `gorm:"column:condition_parameters;type:json;comment:条件参数信息"`
	ResultParameters        datatypes.JSON `gorm:"column:result_parameters;type:json;comment:结果参数信息"`
	FormulaString           string         `gorm:"column:formula_string;type:longtext;comment:公式的原始数据"`
	Formula                 datatypes.JSON `gorm:"column:formula;type:json;comment:解析出来的公式"`
	VariableList            datatypes.JSON `gorm:"column:variable_list;type:json;comment:解析出来的公式变量数值"`
	CoefficientName         datatypes.JSON `gorm:"column:coefficient_name;type:json;comment:公式系数的名字"`
	Coefficient             datatypes.JSON `gorm:"column:coefficient;type:json;comment:公式系数的具体值"`
	CoefficientScore        float64        `gorm:"column:coefficient_score;type:float;comment:公式拟合得分"`
	ComponentParameters     datatypes.JSON `gorm:"column:component_parameters;type:json;comment:拟合设置的组件参数与值"`
	AssociatedParameters    datatypes.JSON `gorm:"column:associated_parameters;type:json;comment:公式变量与实测数据参数名的映射"`
	CreatedAt               *time.Time     `gorm:"column:create_time;autoCreateTime;type:datetime;comment:创建时间" json:"-"`
	UpdatedAt               *time.Time     `gorm:"column:update_time;type:datetime;comment:更新时间" json:"-"`
	UserSpaceId             string         `gorm:"column:userspace_id;type:varchar(128);comment:package所在用户空间的唯一识别标识"`
	Deleted                 gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"-"`
}

type ParameterCalibrationTemplate struct {
	ID                      string         `gorm:"index;primaryKey;type:varchar(128);comment:参数标定模板的唯一标识"`
	TemplateName            string         `gorm:"column:template_name;type:varchar(128);comment:参数标定模板名称"`
	RecordID                string         `gorm:"index;column:record_id;type:varchar(128);comment:参数标定记录的唯一标识"`
	PackageId               string         `gorm:"index;column:package_id;type:varchar(128);comment:package唯一标识"`
	Version                 string         `gorm:"column:version;default:\"\";type:varchar(32);comment:package版本号"`
	UserName                string         `gorm:"index;column:username;type:varchar(128);comment:用户名"`
	ModelName               string         `gorm:"column:model_name;type:varchar(256);comment:模型名称"`
	PackagePath             string         `gorm:"column:package_path;type:varchar(256);comment:标定模型所在包的加载文件"`
	CompileDependencies     datatypes.JSON `gorm:"column:compile_Dependencies;type:json;comment:编译时所需的包环境"`
	CompilePath             string         `gorm:"column:compile_path;type:varchar(128);comment:编译好的文件存放路径"`
	CompileStatus           string         `gorm:"column:compile_status;default:\"0\";type:varchar(32);comment:编译结果状态码，0(初始状态)、3(失败)、4(成功)、6(编译中)"`
	CompileStartTime        int64          `gorm:"column:compile_start_time;type:int;comment:编译开始时间"`
	CompileStopTime         int64          `gorm:"column:compile_stop_time;type:int;comment:编译结束时间"`
	SimulateModelResultPath string         `gorm:"column:simulate_model_result_path;type:varchar(256);comment:仿真结果存储路径"`
	SimulateResultStr       string         `gorm:"column:simulate_result_str;comment:仿真结果输出字符串"`
	SimulateStatus          string         `gorm:"column:simulate_status;default:\"0\";type:varchar(32);comment:仿真状态"`
	StartTime               string         `gorm:"column:start_time;type:varchar(32);comment:仿真开始时间"`
	StopTime                string         `gorm:"column:stop_time;type:varchar(32);comment:仿真结束时间"`
	Tolerance               string         `gorm:"column:tolerance;type:varchar(32);comment:仿真积分误差"`
	NumberOfIntervals       string         `gorm:"column:number_of_intervals;type:varchar(32);comment:仿真间隔数"`
	Interval                string         `gorm:"column:interval;type:varchar(32);comment:仿真时间间隔"`
	Method                  string         `gorm:"column:method;type:varchar(32);comment:仿真积分方法"`
	Percentage              datatypes.JSON `gorm:"column:percentage;type:json;comment:仿真进度, 每次仿真有单独状态"`
	ActualData              datatypes.JSON `gorm:"column:actual_data;type:json;comment:实测数据字段与数据"`
	SimulateResult          datatypes.JSON `gorm:"column:simulate_result;type:json;comment:仿真结果数据"`
	RatedCondition          datatypes.JSON `gorm:"column:rated_condition;type:json;comment:额定工况参数信息"`
	ConditionParameters     datatypes.JSON `gorm:"column:condition_parameters;type:json;comment:条件参数信息"`
	ResultParameters        datatypes.JSON `gorm:"column:result_parameters;type:json;comment:结果参数信息"`
	FormulaString           string         `gorm:"column:formula_string;type:longtext;comment:公式的原始数据"`
	Formula                 datatypes.JSON `gorm:"column:formula;type:json;comment:解析出来的公式"`
	VariableList            datatypes.JSON `gorm:"column:variable_list;type:json;comment:解析出来的公式变量数值"`
	CoefficientName         datatypes.JSON `gorm:"column:coefficient_name;type:json;comment:公式系数的名字"`
	Coefficient             datatypes.JSON `gorm:"column:coefficient;type:json;comment:公式系数的具体值"`
	CoefficientScore        float64        `gorm:"column:coefficient_score;type:float;comment:公式拟合得分"`
	ComponentParameters     datatypes.JSON `gorm:"column:component_parameters;type:json;comment:拟合设置的组件参数与值"`
	AssociatedParameters    datatypes.JSON `gorm:"column:associated_parameters;type:json;comment:公式变量与实测数据参数名的映射"`
	CreatedAt               *time.Time     `gorm:"column:create_time;autoCreateTime;type:datetime;comment:创建时间" json:"-"`
	UpdatedAt               *time.Time     `gorm:"column:update_time;type:datetime;comment:更新时间" json:"-"`
	UserSpaceId             string         `gorm:"column:userspace_id;type:varchar(128);comment:package所在用户空间的唯一识别标识"`
	Deleted                 gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"-"`
}
