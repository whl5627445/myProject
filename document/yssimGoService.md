
# yssimGoService

> v1.0.0

# Default

## POST 授权

POST /license/authorize

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|username|header|string| 否 |none|
|Authorization|header|string| 否 |none|
|space_id|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# yssim-go

## POST 测试命令

POST /test

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|cmd|query|string| 是 |-|
|username|header|string| 否 |none|
|Authorization|header|string| 否 |none|
|space_id|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 测试命令2

GET /

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|cmd|query|string| 否 |-|
|username|header|string| 否 |none|
|Authorization|header|string| 否 |none|
|space_id|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# yssim-go/模型

## GET 检查模型

GET /model/check

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|model_name|query|string| 是 |模型名称|
|username|header|string| 否 |none|
|Authorization|header|string| 否 |none|
|space_id|header|string| 否 |none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "模型检查完成",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 复制模型

POST /model/class/copy

```text
暂无描述
```

> Body 请求参数

```json
{
  "parent_name": "test5",
  "model_name": "Filter8",
  "copied_class_name": "Modelica.Blocks.Examples.Filter",
  "package_id": "41b464e1-d5a1-4e70-a81c-508e04950a54"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» parent_name|body|string| 是 | 父节点名称|none|
|» model_name|body|string| 是 | 模型全名|none|
|» copied_class_name|body|string| 是 | 被复制的模型全名|none|
|» package_id|body|string| 是 | 复制到包的id|模型被复制到哪个包了，id就是哪个包的|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 设置模型文档

POST /model/document/set

```text
暂无描述
```

> Body 请求参数

```json
"{\n  \"revisions\": \"1.0\",\n  \"document\": \"<html><head></head><body><p>\\nThis example123123 demonstrates various options of the\\n<a>Filter</a> block.\\nA step input starts at 0.1 s with an offset of 0.1,in order to demonstrate\\nthe initialization options. This step input drives 4 filter blocks that\\nhave identical parameters,with the only exception of the used analog filter type\\n(CriticalDamping,Bessel,Butterworth,Chebyshev of type I). All the main options\\ncan be set via parameters and are then applied to all the 4 filters.\\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\\n2 Hz resulting in the following outputs:\\n</p>\\n\\n<img src=\\\"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\\\" alt=\\\"Filter1.png\\\" href=\\\"模型缩略图\\\"/>\\n</body></html>\",\n  \"model_name\": \"Filter1\",\n  \"package_id\": \"b282a18a-de5a-469f-84c1-3f9acfd9f048\",\n}"
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» revisions|body|string| 是 | 版本|none|
|» document|body|string| 是 | 文档的html代码片段|none|
|» model_name|body|string| 是 | 模型全名|none|
|» package_id|body|string| 是 | 包id|none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "修改成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## GET 获取模型源码

GET /model/code

```text
# 获取模型的源码数据
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|package_id|query|string| 是 ||模型包id值|
|model_name|query|string| 是 ||模型名称|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": "model Filter \"Demonstrates the Continuous.Filter block with various options\"\n  extends Modelica.Icons.Example;\n  parameter Integer order = 3;\n  parameter Modelica.SIunits.Frequency f_cut = 2;\n  parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass \"Type of filter (LowPass/HighPass)\";\n  parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState \"Type of initialization (no init/steady state/initial state/initial output)\";\n  parameter Boolean normalized = true;\n  Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(\n    Placement(transformation(extent = {{-60, 40}, {-40, 60}})));\n  Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, 40}, {0, 60}})));\n  Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, 0}, {0, 20}})));\n  Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, -40}, {0, -20}})));\n  Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, -80}, {0, -60}})));\nequation\n  connect(step.y, CriticalDamping.u) annotation(\n    Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));\n  connect(step.y, Bessel.u) annotation(\n    Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));\n  connect(Butterworth.u, step.y) annotation(\n    Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n  connect(ChebyshevI.u, step.y) annotation(\n    Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n  annotation(\n    experiment(StopTime = 0.9),\n    Documentation(info = \"<html>\n\n<p>\nThis example demonstrates various options of the\n<a href=\\\"modelica://Modelica.Blocks.Continuous.Filter\\\">Filter</a> block.\nA step input starts at 0.1 s with an offset of 0.1, in order to demonstrate\nthe initialization options. This step input drives 4 filter blocks that\nhave identical parameters, with the only exception of the used analog filter type\n(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main options\ncan be set via parameters and are then applied to all the 4 filters.\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\n2 Hz resulting in the following outputs:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/Filter1.png\\\"\n     alt=\\\"Filter1.png\\\">\n</html>\"));\nend Filter;",
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|string|true|none|模型源码字符串|示例：model Filter "Demonstrates the Continuous.Filter block with various options"  extends Modelica.Icons.Example;  parameter Integer order = 3;  parameter Modelica.SIunits.Frequency f_cut = 2;  parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass "Type of filter (LowPass/HighPass)";  parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState "Type of initialization (no init/steady state/initial state/initial output)";  parameter Boolean normalized = true;  Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(    Placement(transformation(extent = {{-60, 40}, {-40, 60}})));  Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(    Placement(transformation(extent = {{-20, 40}, {0, 60}})));  Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(    Placement(transformation(extent = {{-20, 0}, {0, 20}})));  Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(    Placement(transformation(extent = {{-20, -40}, {0, -20}})));  Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(    Placement(transformation(extent = {{-20, -80}, {0, -60}})));equation  connect(step.y, CriticalDamping.u) annotation(    Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));  connect(step.y, Bessel.u) annotation(    Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));  connect(Butterworth.u, step.y) annotation(    Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));  connect(ChebyshevI.u, step.y) annotation(    Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));  annotation(    experiment(StopTime = 0.9),    Documentation(info = "<html><p>This example demonstrates various options of the<a href=\"modelica://Modelica.Blocks.Continuous.Filter\">Filter</a> block.A step input starts at 0.1 s with an offset of 0.1, in order to demonstratethe initialization options. This step input drives 4 filter blocks thathave identical parameters, with the only exception of the used analog filter type(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main optionscan be set via parameters and are then applied to all the 4 filters.The default setting uses low pass filters of order 3 with a cut-off frequency of2 Hz resulting in the following outputs:</p><img src=\"modelica://Modelica/Resources/Images/Blocks/Filter1.png\"     alt=\"Filter1.png\"></html>"));end Filter;|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 删除模型

POST /model/package/delete

```text
暂无描述
```

> Body 请求参数

```json
{
  "parent_name": "",
  "model_name": "PID_Controller1",
  "package_id": "351b15ab-3f23-4cfa-928c-b46414f935a4"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» parent_name|body|string| 是 | 父节点名称|none|
|» model_name|body|string| 是 | 模型全名|none|
|» package_id|body|string| 是 | 包id|none|

> 返回示例

> 200 Response

```json
{
  "data": null,
  "msg": "string",
  "status": 0,
  "err": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## GET 判断模型是否存在

GET /model/exists

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|package_id|query|string| 否 ||none|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    true
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[boolean]|true|none|返回结果|示例：true|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none||示例：-|
|» err|string|true|none||示例：-|

## GET 获取左侧模型树子节点

GET /model/list_library

```text
获取左侧模型列表接口， 此接口获取系统模型和用户上传模型的子节点节点列表(需用传入父节点名称，返回子节点列表)，暂时没有图标信息
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|package_id|query|string| 是 ||模型包id值|
|model_name|query|string| 是 ||模型名称|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.PID_Controller",
      "name": "PID_Controller"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.Filter",
      "name": "Filter"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.FilterWithDifferentiation",
      "name": "FilterWithDifferentiation"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.FilterWithRiseTime",
      "name": "FilterWithRiseTime"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.SlewRateLimiter",
      "name": "SlewRateLimiter"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.InverseModel",
      "name": "InverseModel"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.ShowLogicalSources",
      "name": "ShowLogicalSources"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.LogicalNetwork1",
      "name": "LogicalNetwork1"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.RealNetwork1",
      "name": "RealNetwork1"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.IntegerNetwork1",
      "name": "IntegerNetwork1"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.BooleanNetwork1",
      "name": "BooleanNetwork1"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.Interaction1",
      "name": "Interaction1"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.BusUsage",
      "name": "BusUsage"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.Rectifier6pulseFFT",
      "name": "Rectifier6pulseFFT"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.Rectifier12pulseFFT",
      "name": "Rectifier12pulseFFT"
    },
    {
      "haschild": false,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.TotalHarmonicDistortion",
      "name": "TotalHarmonicDistortion"
    },
    {
      "haschild": true,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.NoiseExamples",
      "name": "NoiseExamples"
    },
    {
      "haschild": true,
      "image": "",
      "model_name": "Modelica.Blocks.Examples.BusUsage_Utilities",
      "name": "BusUsage_Utilities"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|返回数据对象|示例：-|
|»» haschild|boolean|true|none|是否有子节点|示例：-|
|»» image|string|true|none|图标的base64编码|none|
|»» model_name|string|true|none|模型名称|示例：Modelica.Blocks.Examples.PID_Controller|
|»» name|string|true|none|模型简称|示例：PID_Controller|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 获取模型文档

GET /model/document/get

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": {
    "document": "<html><head></head><body><p>\nThis is a simple drive train controlled by a PID controller:\n</p>\n\n<ul>\n<li> The two blocks &#34;kinematic_PTP&#34; and &#34;integrator&#34; are used to generate\n     the reference speed (= constant acceleration phase, constant speed phase,\n     constant deceleration phase until inertia is at rest). To check\n     whether the system starts in steady state, the reference speed is\n     zero until time = 0.5 s and then follows the sketched trajectory.</li>\n\n<li> The block &#34;PI&#34; is an instance of &#34;Blocks.Continuous.LimPID&#34; which is\n     a PID controller where several practical important aspects, such as\n     anti-windup-compensation has been added. In this case, the control block\n     is used as PI controller.</li>\n\n<li> The output of the controller is a torque that drives a motor inertia\n     &#34;inertia1&#34;. Via a compliant spring/damper component, the load\n     inertia &#34;inertia2&#34; is attached. A constant external torque of 10 Nm\n     is acting on the load inertia.</li>\n</ul>\n\n<p>\nThe PI controller settings included &#34;limitAtInit=false&#34;, in order that\nthe controller output limits of 12 Nm are removed from the initialization\nproblem.\n</p>\n\n<p>\nThe PI controller is initialized in steady state (initType=SteadyState)\nand the drive shall also be initialized in steady state.\nHowever, it is not possible to initialize &#34;inertia1&#34; in SteadyState, because\n&#34;der(inertia1.phi)=inertia1.w=0&#34; is an input to the PI controller that\ndefines that the derivative of the integrator state is zero (= the same\ncondition that was already defined by option SteadyState of the PI controller).\nFurthermore, one initial condition is missing, because the absolute position\nof inertia1 or inertia2 is not defined. The solution shown in this examples is\nto initialize the angle and the angular acceleration of &#34;inertia1&#34;.\n</p>\n\n<p>\nIn the following figure, results of a typical simulation are shown:\n</p>\n\n<img src=\"data:image/jpeg;base64,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\" alt=\"PID_controller.png\" href=\"模型缩略图\"/><br/>\n\n<img src=\"data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAfUAAADjBAMAAABunpAQAAAAD1BMVEX////AwMAAAAD/AAAAAP8RT9RwAAAEHElEQVR4Xu3dcU7rOBCA8Q97DvDesgeY9Qm86wsMku9/ppUeiAqEQmra1PHM9yepkX5MaFQlcvFdFEVRlIomUAVQ3fXqZUqkBG/o7+2k9eyqpVAKRVMCRbfsii5lLwlVTaSkq9r7x+z9/11f7aUkLrj2sfr66mXO+wSJix0Sm3NnXfurS9OmHV3Irm/2QtH0h75lVxS3ebYnZiuKoihKBbcp6tjutxLnvMOKKkrx01dzVzb6dfzBJ+5yMH2+xrm1Q9jD/l7Ywx723PzaG82rPQPNsZ3m014BaI7tNMf2XP3ZM2/l5tcOrfq1k1trnuyVj7V91SXsQwLOZw972MOeWd8e9rCHPex1MXvYU7mn/Xlqu6Iu7GGPcz7mXnSXPbPs3Esp//ze6nnr4PbCufpr4xrna+5wH3vYf9VJ7WEPe9jHBWEPe9jDnie1hz3sYQ/7+KexsIc97NTp7WEPe9jDHvY6ZA97nsQe9rCHXUnFqT0pivq06/3t1EntOLSHvWjMnVJK2+rf3z/ow+LnsV893ifJf1zS++/xknkvt/p5Zb3r3HOr259lDrO3CofaW4U57A2OtTeYxN442J6Zxd442l5nsed6tD3zeDv1Qj/SXmew5wv9QHtmBjut1Y0/zNp24Hh79WvPeLSHvfq1Zzzaw17ntoedeg97xqM97NWjPewZj/awV4/2sGdOYs+3tIe9hn3P8zar2fPA8zbO7Krg1Y5ze2Ixe939vA3KZY+XB/Tcbt3ePV6UxGPnLree+wu7z/lS9KF27LZ2GXiWeBW7+bULfu3m1y74tZtfu+DXbn7twsnscju7nc1O737t0G9jF05op9/Ebqe0029g75zTTv+xXTirXeyndjutHbGf2YXz2hH7kd3ObEds3I5wajvSx+12cjv0UbtwervYoN3Ob0fMrx2xEbuwgh2xAbutYUfsaruwiB2xa+22jB2x6+zCOnb6dXYbsyuk+ex0u8Iuw/vbUCa0Iy/77TZmVyWlGe2I7bULY3amtT+J7bTbuF0ntSO2yy7j+9tQyuU7leaq71nw98vwdyopzDp36DvmLozMfX47/Xu7jdsB5rXTv7MLy9rp23ZsYTt90y6sbKdv8Wxl+/ZNWmFl+/ZNWlvZvn2TVljfTv+aZx7s9K94ggs7/QueObGLObNv36QVvNgR82tHul87dMd2umM75tgu5teOmF87Yn7tiPm1I/bKM4d2xPza6Y7tdPNrfxLza0fMrx3pB3yf1KR2ZMM+zf42U3+fFD7tKKV4tbs954u+2Yufrpr7gw4m7n5QSQV/9kthD/t2k6xMjBR2z0VRVPT4N6xy/ErSVz9Ko4Jhu6KHr6Tc0q7Kia6O6aZz5xH2wtnticPnrrPY9Xh7KTqFPXH8uyRpjmtcKTrBNc5DUfQ/afA5HzRH1awAAAAASUVORK5CYII=\" alt=\"PID_controller2.png\" href=\"模型缩略图\"/>\n\n<p>\nIn the upper figure the reference speed (= integrator.y) and\nthe actual speed (= inertia1.w) are shown. As can be seen,\nthe system initializes in steady state, since no transients\nare present. The inertia follows the reference speed quite good\nuntil the end of the constant speed phase. Then there is a deviation.\nIn the lower figure the reason can be seen: The output of the\ncontroller (PI.y) is in its limits. The anti-windup compensation\nworks reasonably, since the input to the limiter (PI.limiter.u)\nis forced back to its limit after a transient phase.\n</p>\n\n</body></html>",
    "revisions": ""
  },
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none|返回数据对象|示例：-|
|»» document|string|true|none||文档内容，html格式|
|»» revisions|string|true|none|修订版本号|示例：-|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 获取新增组件名称

GET /model/component/name

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|class_name|query|string| 是 ||组件名称|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    "sin"
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[string]|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 更新连接器起始位置名称

POST /model/connection/name

```text
暂无描述
```

> Body 请求参数

```json
"{\n  \"from_name\": \"sum1.y\",\n  \"to_name\": \"sum2.y\",\n  \"from_name_new\": \"sum1new.y\",\n  \"to_name_new\": \"sum2new.y\",\n  \"model_name\": \"PID_Controller1\",\n  \"package_id\": \"351b15ab-3f23-4cfa-928c-b46414f935a4\",\n}"
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» from_name|body|string| 是 | 起始位置|none|
|» to_name|body|string| 是 | 终点位置|none|
|» from_name_new|body|string| 是 | 新起始位置|none|
|» to_name_new|body|string| 是 | 新终点位置|none|
|» model_name|body|string| 是 | 模型全名|none|
|» package_id|body|string| 是 | 包id|none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "更新连接器成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 创建组件连线

POST /model/connection/create

```text
暂无描述
```

> Body 请求参数

```json
{
  "package_name": "test1235",
  "package_id": "c20fb619-9e95-4bf1-a4c1-5db2f3d729a8",
  "model_name": "test1234",
  "connect_start": "abs.y",
  "connect_end": "sign.u",
  "line_points": [
    "95.08,-93.08",
    "126.88,-93.08"
  ],
  "color": "0,0,127"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_name|body|string| 是 ||none|
|» package_id|body|string| 是 ||模型包id值|
|» model_name|body|string| 是 ||要删除模型的名称|
|» connect_start|body|string| 是 ||连线起始位置|
|» connect_end|body|string| 是 ||连线终点位置|
|» line_points|body|[string]| 是 ||连线拐点，包含起始位置|
|» color|body|string| 是 ||连线颜色|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "连接组件成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 单位换算

POST /model/units/convert

```text
暂无描述
```

> Body 请求参数

```json
"{\n  \"s1\": \"N.m\",\n  \"s2\": \"N.m\",\n}"
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» s1|body|string| 是 ||需要转换成什么单位名称|
|» s2|body|string| 是 ||被转换的单位名称|

> 返回示例

> 成功

```json
{
  "data": {
    "offset": 0,
    "scale_factor": 1,
    "units_compatible": true
  },
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none|返回数据对象|示例：-|
|»» offset|integer|true|none|偏移量|示例：-|
|»» scale_factor|integer|true|none|换算比例|示例：1|
|»» units_compatible|boolean|true|none|两个单位是否兼容|示例：true|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 设置模型属性

POST /model/properties/set

```text
暂无描述
```

> Body 请求参数

```json
{
  "model_name": "Filter1",
  "old_component_name": "step",
  "new_component_name": "step",
  "final": false,
  "protected": false,
  "replaceable": false,
  "variability": "unspecified",
  "inner": true,
  "outer": true,
  "package_id": "b282a18a-de5a-469f-84c1-3f9acfd9f048",
  "causality": "input"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» model_name|body|string| 是 ||模型名称|
|» old_component_name|body|string| 是 ||模型组件原名称|
|» new_component_name|body|string| 是 ||模型组件新名称|
|» final|body|boolean| 是 ||组件属性|
|» protected|body|boolean| 是 ||组件属性|
|» replaceable|body|boolean| 是 ||组件属性|
|» variability|body|string| 是 ||组件属性|
|» inner|body|boolean| 是 ||组件属性|
|» outer|body|boolean| 是 ||组件属性|
|» package_id|body|string| 是 ||模型包id值|
|» causality|body|string| 是 ||组件属性|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "设置完成",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## GET 获取模型属性

GET /model/properties/get

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|component_name|query|string| 是 ||组件名称|
|package_id|query|string| 是 ||模型包id值|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": {
    "Causality": "unspecified",
    "Inner/Outer": "none",
    "Properties": [
      "false",
      "public",
      "false"
    ],
    "Variability": "unspecified",
    "annotation": "",
    "component_name": "PI",
    "dimension": "[]",
    "model_name": "Modelica.Blocks.Examples.PID_Controller",
    "path": "Modelica.Blocks.Continuous.LimPID"
  },
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» Causality|string|true|none||none|
|»» Inner|string|false|none||none|
|»» Properties|[string]|true|none||none|
|»» Variability|string|true|none||none|
|»» annotation|string|true|none||none|
|»» component_name|string|true|none||none|
|»» dimension|string|true|none||none|
|»» model_name|string|true|none||none|
|»» path|string|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 设置模型参数

POST /model/parameters/set

```text
暂无描述
```

> Body 请求参数

```json
{
  "model_name": "PID_Controller1",
  "parameter_value": {
    "PI.controllerType": "Modelica.Blocks.Types.SimpleController.PI",
    "PI.k": "200",
    "PI.Ti": "0.1",
    "PI.Td": "0.1",
    "PI.yMax": "12",
    "PI.yMin": "",
    "PI.wp": "",
    "PI.wd": "",
    "PI.Ni": "0.1",
    "PI.Nd": "",
    "PI.withFeedForward": "false",
    "PI.kFF": "",
    "PI.initType": "Modelica.Blocks.Types.InitPID.SteadyState",
    "PI.xi_start": "",
    "PI.xd_start": "",
    "PI.y_start": "",
    "PI.homotopyType": "",
    "PI.strict": "false",
    "PI.limitsAtInit": "false"
  },
  "package_id": "351b15ab-3f23-4cfa-928c-b46414f935a4"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» model_name|body|string| 是 ||模型名称|
|» parameter_value|body|object| 是 ||参数对象，包含组件所有参数信息|
|»» PI.controllerType|body|string| 是 ||none|
|»» PI.k|body|string| 是 ||none|
|»» PI.Ti|body|string| 是 ||none|
|»» PI.Td|body|string| 是 ||none|
|»» PI.yMax|body|string| 是 ||none|
|»» PI.yMin|body|string| 是 ||none|
|»» PI.wp|body|string| 是 ||none|
|»» PI.wd|body|string| 是 ||none|
|»» PI.Ni|body|string| 是 ||none|
|»» PI.Nd|body|string| 是 ||none|
|»» PI.withFeedForward|body|string| 是 ||none|
|»» PI.kFF|body|string| 是 ||none|
|»» PI.initType|body|string| 是 ||none|
|»» PI.xi_start|body|string| 是 ||none|
|»» PI.xd_start|body|string| 是 ||none|
|»» PI.y_start|body|string| 是 ||none|
|»» PI.homotopyType|body|string| 是 ||none|
|»» PI.strict|body|string| 是 ||none|
|»» PI.limitsAtInit|body|string| 是 ||none|
|» package_id|body|string| 是 ||模型包id值|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "设置完成",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 获取图形数据

POST /model/graphics

```text
获取模型的画图数据
```

> Body 请求参数

```json
{
  "model_name": "Modelica.Blocks.Examples.PID_Controller",
  "package_id": "6d602d98-01b2-4625-a0cd-940b9a91980b"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» model_name|body|string| 是 ||模型名称|
|» package_id|body|string| 是 ||模型包id值|

> 返回示例

> 200 Response

```json
{
  "data": [
    [
      {
        "ID": "string",
        "classname": "string",
        "extent1Diagram": "string",
        "extent2Diagram": "string",
        "graphType": "string",
        "inputOutputs": [
          {
            "ID": "string",
            "classname": "string",
            "extent1Diagram": "string",
            "extent2Diagram": "string",
            "graphType": "string",
            "inputOutputs": [
              null
            ],
            "name": "string",
            "originDiagram": "string",
            "original_name": "string",
            "output_type": "string",
            "parent": "string",
            "rotateAngle": "string",
            "rotation": "string",
            "subShapes": [
              null
            ],
            "type": "string",
            "visible": "string"
          }
        ],
        "name": "string",
        "originDiagram": "string",
        "original_name": "string",
        "output_type": "string",
        "parent": "string",
        "rotateAngle": "string",
        "rotation": "string",
        "subShapes": [
          {
            "arrow": "string",
            "arrowSize": "string",
            "color": "string",
            "linePattern": "string",
            "lineThickness": "string",
            "originalPoint": "string",
            "points": [
              null
            ],
            "rotation": "string",
            "smooth": "string",
            "type": "string",
            "visible": "string",
            "fillColor": "string",
            "fillPattern": "string",
            "polygonPoints": [
              null
            ],
            "extentsPoints": [
              null
            ],
            "fontName": "string",
            "fontSize": "string",
            "horizontalAlignment": "string",
            "originalTextString": "string",
            "textColor": "string",
            "textStyles": [
              null
            ],
            "borderPattern": "string",
            "radius": "string"
          }
        ],
        "type": "string",
        "visible": "string"
      }
    ]
  ],
  "msg": "string",
  "status": 0,
  "err": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[array]|true|none||none|
|»» ID|string|true|none||none|
|»» classname|string|true|none||none|
|»» extent1Diagram|string|true|none||none|
|»» extent2Diagram|string|true|none||none|
|»» graphType|string|true|none||none|
|»» inputOutputs|[object]|true|none||none|
|»»» ID|string|true|none||none|
|»»» classname|string|true|none||none|
|»»» extent1Diagram|string|true|none||none|
|»»» extent2Diagram|string|true|none||none|
|»»» graphType|string|true|none||none|
|»»» inputOutputs|[string]|true|none||none|
|»»» name|string|true|none||none|
|»»» originDiagram|string|true|none||none|
|»»» original_name|string|true|none||none|
|»»» output_type|string|true|none||none|
|»»» parent|string|true|none||none|
|»»» rotateAngle|string|true|none||none|
|»»» rotation|string|true|none||none|
|»»» subShapes|[object]|true|none||none|
|»»»» color|string|true|none||none|
|»»»» fillColor|string|true|none||none|
|»»»» fillPattern|string|true|none||none|
|»»»» linePattern|string|true|none||none|
|»»»» lineThickness|string|true|none||none|
|»»»» originalPoint|string|true|none||none|
|»»»» polygonPoints|[string]|true|none||none|
|»»»» rotation|string|true|none||none|
|»»»» smooth|string|true|none||none|
|»»»» type|string|true|none||none|
|»»»» visible|string|true|none||none|
|»»» type|string|true|none||none|
|»»» visible|string|true|none||none|
|»» name|string|true|none||none|
|»» originDiagram|string|true|none||none|
|»» original_name|string|true|none||none|
|»» output_type|string|true|none||none|
|»» parent|string|true|none||none|
|»» rotateAngle|string|true|none||none|
|»» rotation|string|true|none||none|
|»» subShapes|[object]|true|none||none|
|»»» arrow|string|true|none||none|
|»»» arrowSize|string|true|none||none|
|»»» color|string|true|none||none|
|»»» linePattern|string|true|none||none|
|»»» lineThickness|string|true|none||none|
|»»» originalPoint|string|true|none||none|
|»»» points|[string]|true|none||none|
|»»» rotation|string|true|none||none|
|»»» smooth|string|true|none||none|
|»»» type|string|true|none||none|
|»»» visible|string|true|none||none|
|»»» fillColor|string|true|none||none|
|»»» fillPattern|string|true|none||none|
|»»» polygonPoints|[string]|true|none||none|
|»»» extentsPoints|[string]|true|none||none|
|»»» fontName|string|true|none||none|
|»»» fontSize|string|true|none||none|
|»»» horizontalAlignment|string|true|none||none|
|»»» originalTextString|string|true|none||none|
|»»» textColor|string|true|none||none|
|»»» textStyles|[string]|true|none||none|
|»»» borderPattern|string|true|none||none|
|»»» radius|string|true|none||none|
|»» type|string|true|none||none|
|»» visible|string|true|none||none|
|msg|string|true|none||none|
|status|integer|true|none||none|
|err|string|true|none||none|

## GET 获取模型的全部组件描述数据

GET /model/components/get

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|package_id|query|string| 否 ||none|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "component_description": "",
      "component_model_name": "Modelica.Blocks.Continuous.LimPID",
      "component_name": "PI"
    },
    {
      "component_description": "",
      "component_model_name": "Modelica.Mechanics.Rotational.Components.Inertia",
      "component_name": "inertia1"
    },
    {
      "component_description": "",
      "component_model_name": "Modelica.Mechanics.Rotational.Sources.Torque",
      "component_name": "torque"
    },
    {
      "component_description": "",
      "component_model_name": "Modelica.Mechanics.Rotational.Components.SpringDamper",
      "component_name": "spring"
    },
    {
      "component_description": "",
      "component_model_name": "Modelica.Mechanics.Rotational.Components.Inertia",
      "component_name": "inertia2"
    },
    {
      "component_description": "",
      "component_model_name": "Modelica.Blocks.Sources.KinematicPTP",
      "component_name": "kinematicPTP"
    },
    {
      "component_description": "",
      "component_model_name": "Modelica.Blocks.Continuous.Integrator",
      "component_name": "integrator"
    },
    {
      "component_description": "",
      "component_model_name": "Modelica.Mechanics.Rotational.Sensors.SpeedSensor",
      "component_name": "speedSensor"
    },
    {
      "component_description": "",
      "component_model_name": "Modelica.Mechanics.Rotational.Sources.ConstantTorque",
      "component_name": "loadTorque"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|返回数据对象|示例：-|
|»» component_description|string|true|none|组件描述|示例：-|
|»» component_model_name|string|true|none|组件的模型名称|示例：Modelica.Blocks.Continuous.LimPID|
|»» component_name|string|true|none|模型组件名称|示例：PI|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 获取模型参数

GET /model/parameters/get

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|class_name|query|string| 否 ||模型组件的模型全名|
|component_name|query|string| 否 ||模型组件别名|
|package_id|query|string| 是 ||模型包id值|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": {
    "parameters": [
      {
        "comment": "Integrator gain",
        "defaultvalue": "1",
        "group": "参数",
        "name": "k",
        "tab": "通用设置",
        "type": "Normal",
        "unit": "",
        "value": ""
      },
      {
        "checked": "false",
        "comment": "= true,if reset port enabled",
        "defaultvalue": "false",
        "group": "参数",
        "name": "use_reset",
        "tab": "通用设置",
        "type": "CheckBox",
        "unit": "",
        "value": "false"
      },
      {
        "checked": "false",
        "comment": "= true,if set port enabled and used as reinitialization value when reset",
        "defaultvalue": "false",
        "group": "参数",
        "name": "use_set",
        "tab": "通用设置",
        "type": "CheckBox",
        "unit": "",
        "value": "false"
      },
      {
        "comment": "Type of initialization (1: no init,2: steady state,3,4: initial output)",
        "defaultvalue": "Init.InitialState",
        "group": "参数",
        "name": "initType",
        "options": [
          "Modelica.Blocks.Types.Init.NoInit",
          "Modelica.Blocks.Types.Init.SteadyState",
          "Modelica.Blocks.Types.Init.InitialState",
          "Modelica.Blocks.Types.Init.InitialOutput"
        ],
        "tab": "通用设置",
        "type": "Enumeration",
        "unit": "",
        "value": "Modelica.Blocks.Types.Init.InitialState"
      },
      {
        "comment": "Initial or guess value of output (= state)",
        "defaultvalue": "0",
        "group": "参数",
        "name": "y_start",
        "tab": "通用设置",
        "type": "Normal",
        "unit": "",
        "value": ""
      },
      {
        "comment": "y",
        "defaultvalue": false,
        "group": "初始化",
        "name": "y.fixed",
        "tab": "通用设置",
        "type": "fixed",
        "unit": null,
        "value": false
      },
      {
        "comment": "y",
        "defaultvalue": "y_start",
        "group": "初始化",
        "name": "y.start",
        "tab": "通用设置",
        "tab1": "",
        "type": "Normal",
        "unit": null,
        "value": ""
      }
    ],
    "properties": {
      "Causality": "unspecified",
      "Inner/Outer": "none",
      "Properties": [
        "false",
        "public",
        "false"
      ],
      "Variability": "unspecified",
      "annotation": "",
      "component_name": "integrator",
      "dimension": "[]",
      "model_name": "Modelica.Blocks.Examples.PID_Controller",
      "path": "Modelica.Blocks.Continuous.Integrator"
    }
  },
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none|返回数据对象数组|示例：-|
|»» parameters|[object]|true|none||示例：-|
|»»» comment|string|true|none||示例：Integrator gain|
|»»» defaultvalue|string|true|none||示例：1|
|»»» group|string|true|none||示例：参数|
|»»» name|string|true|none|模型名称|示例：k|
|»»» tab|string|true|none||示例：通用设置|
|»»» type|string|true|none||示例：Normal|
|»»» unit|string¦null|true|none|变量单位|示例：-|
|»»» value|string|true|none||示例：-|
|»»» checked|string|true|none||none|
|»»» options|[string]|false|none||none|
|»»» tab1|string|false|none||none|
|»» properties|object|true|none||示例：-|
|»»» Causality|string|true|none||示例：unspecified|
|»»» Inner|string|false|none||示例：none|
|»»» Properties|[string]|true|none||示例：false|
|»»» Variability|string|true|none||示例：unspecified|
|»»» annotation|string|true|none||示例：-|
|»»» component_name|string|true|none|模型组件别名|示例：integrator|
|»»» dimension|string|true|none||示例：[]|
|»»» model_name|string|true|none|模型名称|示例：Modelica.Blocks.Examples.PID_Controller|
|»»» path|string|true|none||示例：Modelica.Blocks.Continuous.Integrator|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码|示例：-|
|» err|string|true|none|cuo|示例：-|

## POST 删除组件连线

POST /model/connection/delete

```text
暂无描述
```

> Body 请求参数

```json
"{\n  \"connect_start\": \"Bessel.u\",\n  \"connect_end\": \"step.y\",\n  \"model_name\": \"PID_Controller1\",\n  \"package_id\": \"351b15ab-3f23-4cfa-928c-b46414f935a4\",\n}"
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» connect_start|body|string| 是 | 起点名称|none|
|» connect_end|body|string| 是 | 终点名称|none|
|» model_name|body|string| 是 | 模型全称|none|
|» package_id|body|string| 是 | 包id|none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "删除连线成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 创建模型组件

POST /model/component/add

```text
暂无描述
```

> Body 请求参数

```json
{
  "extent": [
    "-10,-10",
    "10,10"
  ],
  "model_name": "test1235",
  "new_component_name": "sign",
  "old_component_name": "Modelica.Blocks.Math.Sign",
  "origin": "140.8800048828125,-93.0800048828125",
  "package_id": "c20fb619-9e95-4bf1-a4c1-5db2f3d729a8",
  "rotation": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» extent|body|[string]| 是 ||范围坐标|
|» model_name|body|string| 是 ||模型名称|
|» new_component_name|body|string| 是 ||新增后的组件别名|
|» old_component_name|body|string| 是 ||组件的模型全名|
|» origin|body|string| 是 ||原点|
|» package_id|body|string| 是 ||模型包id值|
|» rotation|body|integer| 是 ||旋转角度|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "新增组件成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 删除模型组件

POST /model/component/delete

```text
暂无描述
```

> Body 请求参数

```json
{
  "package_name": "test2",
  "package_id": "a8c3a58d-1ae3-48e0-9982-a77fd2f520f7",
  "component_list": [
    {
      "delete_type": "component",
      "component_name": "sin2",
      "model_name": "test2.testModel"
    }
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 是 ||用户名|
|space_id|header|string| 是 ||用户空间id|
|Authorization|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_name|body|string| 是 ||none|
|» package_id|body|string| 是 ||模型包id值|
|» component_list|body|[object]| 是 ||需要删除的对象数组|
|»» delete_type|body|string| 否 ||删除类型，是组件还是连线|
|»» component_name|body|string| 否 ||模型组件别名|
|»» model_name|body|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "删除成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 更新模型组件

POST /model/component/update

```text
暂无描述
```

> Body 请求参数

```json
"{\n  \"extent\": [\"-10,-10\", \"10,10\"],\n  \"rotation\": \"0\",\n  \"origin\": \"0,0\",\n  \"component_class_name\": \"Modelica.Blocks.Continuous.LimPID\",\n  \"component_name\": \"limPID\",\n  \"model_name\": \"PID_Controller1\",\n  \"package_id\": \"351b15ab-3f23-4cfa-928c-b46414f935a4\",\n}"
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» extent|body|[string]| 是 | 范围坐标|none|
|» rotation|body|string| 是 | 旋转角度|none|
|» origin|body|string| 是 | 原点|none|
|» component_class_name|body|string| 是 | 组件模型名称|none|
|» component_name|body|string| 是 | 组件名称|none|
|» model_name|body|string| 是 | 模型全名|none|
|» package_id|body|string| 是 | 包id|none|
|» connection_list|body|[object]| 是 ||none|
|»» model_name|body|string| 是 | 模型全名|none|
|»» connect_start|body|string| 是 | 连线起始位置的名称|none|
|»» connect_end|body|string| 是 | 连线终止位置的名称|none|
|»» color|body|string| 是 | 颜色|none|
|»» line_points|body|[string]| 是 | 中间的拐点坐标|none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "更新组件成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## POST 更新组件连线

POST /model/connection/update

```text
暂无描述
```

> Body 请求参数

```json
{
  "package_id": "351b15ab-3f23-4cfa-928c-b46414f935a4",
  "model_name": "we.Filter1",
  "connect_start": "ChebyshevI.u",
  "connect_end": "step.y",
  "line_points": [
    "-10,-70",
    "-34,-70",
    "-10,36",
    "-34,50",
    "-50,50"
  ],
  "color": "0,0,127"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_id|body|string| 是 ||模型包id值|
|» model_name|body|string| 是 ||模型名称|
|» connect_start|body|string| 是 ||连线起始位置|
|» connect_end|body|string| 是 ||连线终点位置|
|» line_points|body|[string]| 是 ||连线拐点，包含起始位置|
|» color|body|string| 是 ||连线颜色|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "连接组件成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## GET 获取左侧系统模型树根节点

GET /model/root_library/sys

```text
获取左侧模型列表接口， 此接口获取系统模型的根节点列表，暂时没有图标信息
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "haschild": true,
      "image": "",
      "package_id": "6d602d98-01b2-4625-a0cd-940b9a91980b",
      "package_name": "Modelica"
    },
    {
      "haschild": false,
      "image": "",
      "package_id": "b282a18a-de5a-469f-84c1-3f9acfd9f048",
      "package_name": "SolarPower"
    },
    {
      "haschild": true,
      "image": "",
      "package_id": "d942f1a3-c5b3-45cd-9cdb-c951784e01ce",
      "package_name": "Buildings"
    },
    {
      "haschild": false,
      "image": "",
      "package_id": "fea10d0d-0c33-4cbd-8f6e-784b5ba81545",
      "package_name": "WindPowerSystem"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取左侧用户模型树根节点 

GET /model/root_library/user

```text
获取左侧模型列表接口， 此接口获取用户上传模型的根节点列表，暂时没有图标信息
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "haschild": false,
      "image": "",
      "package_id": "6a928649-5aa2-4f3f-a381-37099780624b",
      "package_name": "PID_Controller1"
    },
    {
      "haschild": false,
      "image": "",
      "package_id": "ff2ff7d5-8047-4782-97f8-9bfca069231c",
      "package_name": "PID_Controller"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 新增模型收藏

POST /model/collection/create

> Body 请求参数

```json
{
  "package_id": "0e066cfb-7553-4db7-83cf-b76cb0c96339",
  "model_name": "X_p.X_m"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "创建成功",
  "status": 0,
  "err": ""
}
```

```json
"Creation failed"
```

```json
{
  "data": null,
  "msg": "",
  "status": 2,
  "err": "名称已存在，请修改后再试。"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取收藏模型列表

GET /model/collection/get

和接口“获取左侧用户模型树根节点”相比，data中多了一个字段（”id“），删除收藏模型的时候，传入这个 ”id“就行。

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "haschild": false,
      "id": "edffc445-de7f-430e-a2c0-5107e6abf06b",
      "image": "",
      "model_name": "Xxxs",
      "package_id": "58a4835e-95cb-47bd-9385-6e122c9f5f13",
      "type": "model"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

```json
{
  "data": null,
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 删除收藏的模型

GET /model/collection/delete

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|query|string| 否 ||none|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "删除成功",
  "status": 0,
  "err": ""
}
```

```json
"not found"
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 关键字搜索模型

GET /model/search

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|keywords|query|string| 否 ||需要搜索的关键字|
|parent|query|string| 否 ||在哪个节点下进行搜索， 为空字符串时全局搜索|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "删除成功",
  "status": 0,
  "err": ""
}
```

```json
"not found"
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# yssim-go/用户

## GET 获取用户空间列表

GET /user/userspace/get

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "id": "1",
      "name": "test"
    },
    {
      "id": "26689c0e-667b-44f1-b072-ab5e8f298d9e",
      "name": ""
    },
    {
      "id": "973c6041-ee5f-491f-8a10-22e49fce1c13",
      "name": "test2"
    },
    {
      "id": "692fe630-1055-4fc4-ab07-af667f51ab52",
      "name": "test1"
    },
    {
      "id": "5ec43df7-0353-48ca-83aa-dfbc6b509563",
      "name": "test3"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|返回数据对象|示例：-|
|»» id|string|true|none|用户空间id|示例：1|
|»» name|string|true|none|用户空间名称|示例：test|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 创建用户空间

POST /user/userspace/create

```text
暂无描述
```

> Body 请求参数

```json
{
  "space_name": "test1"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» space_name|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "data": {
    "id": "string",
    "name": "string"
  },
  "msg": "string",
  "status": 0,
  "err": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||none|
|»» id|string|true|none||none|
|»» name|string|true|none||none|
|» msg|string|true|none||none|
|» status|integer|true|none||none|
|» err|string|true|none||none|

## GET 获取模型示例数据

GET /user/examples

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "id": "1",
      "image": "/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAIBAQEBAQIBAQECAgICAgQDAgICAgUEBAMEBgUGBgYFBgYGBwkIBgcJBwYGCAsICQoKCgoKBggLDAsKDAkKCgr/2wBDAQICAgICAgUDAwUKBwYHCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgr/wAARCAMSAxIDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooJAOM0UAFFGR60UAFFGR6iigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKG6H6UUUAQsTnGaepLD5a8V/bK/bx/Zp/Yl8PWOo/G34p+HNK1HVL21g0rQ9U16G1ubpJrqK3adVbLCGIyb5JduxFRizKASOI/bV/aM8JfEv/AIJu/Fr44/smftF6bqD+HfBmq32meLPh/wCIrW+jgvbS1acR+dEZI8jCbl6gN1Bwa6KWExNRxk42jJ2T79/uuh2Pp3zCeQeKeAScCvNvA/7SPwR8T+NJPg94e+MnhXVPGOl6bFLrPhSy8R2suqWK7VJaa1RzLGPnXllA+YeoqH47fG/4JeDbOH4Y/Ej9obQvA2t+K7aSDw5Hd+KLOw1K6lyFBs0uD++cMyjCq/JA2nNJ4erGqoNbhY9M3nOM0+vmX9jfxp+0F4+/4JsfDTx14V8W6fqXjrVfA2mXtzq3jaOW8ivZ2t0MjSmCSJtzHnIOOeFrrf8A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ObP9of4t689orhdM8UfEe81GykLKVy8EpKkgE4PGDXa/Bnxb8W/G/hP+1fi78ILTwbq32hlOkRa3HqIEYC4fzo1VeSWG3HG33rs4bbZklFB74FcNeviXFwlLX/t0G2TDkc0UDgYorkWwgooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigApGOBmlpGIKkZoA89b9qj9mrP/JwHgn/AMKa1/8Ai6P+GqP2av8Ao4DwT/4U1r/8VWe/7Hn7IhYqP2Xfh8een/CHWX/xqmf8MV/siHn/AIZT+Gx+vg6y/wDjVdX+w/3h6GrD+1H+zfdTLZ2vx68FyTycRRL4mtcufb567uKQSoHAx6j0PpXmVt+xn+yZZ3Ed7Zfsu/Dm3uIJBJbzxeDrIPFIPuup8rII9Rg+9eiwKIWEMYCqvAVeAKyqRoS/ht/MReIBopF+7S1kAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQBHQSB1OKBycV498fv2gvGngL4g+Gfgp8Gfh1pnivxt4rsNQ1KwsPEPiN9H0220+xktY7m4nu47W6kUiS9tY0jjgkZ2mGdiK7rUITqT5YrowPXRcxscKRmneZ7V81/A/9rT9pT4v2fjpL39k7Q9K1HwNrC6LPZW/xI+0pdaog33UQkNimyJIXtp45SpMiXIDJDIkka1fgd+2/8fP2hPhz4J+LXws/Y9+2aFrN/Z23jN7rx7BDNowlmWOeSzH2cx6ktpmRpz5lvzGUh899yL2PLcZGm6jUbK1/fj126gfTp2E5KCkACjAFeA61+1h8dPF3jnxR4c/Zm/Zq0zxfpvge+Fh4m1XxD46bRGuL0RRzSWmmx/YbhbuSOOWMM00ltCJX8rzcpKY6S/ty6v4q1/4E6v8ADr4e2sng74veJLvRtTv9Z1GW31fRLyHTb+7Nq9h5BTesljLDKWnVo3UhY5A29cpYHFNKX5NPpfXtomOx9FFQTkilrw7Sv2lvizL+0b8Vvgnqnwp0QReCPA+keIvC8tj4kmkl1pL19RiEdxutVWyYSaeRhfPG2QOW6oOH8N/8FEviRrPwb+Hf7T+r/sxJpXw18cJoMVzqF54wJ1nT7jVp4be2dLBbIxT2oluIVaZrmKXaWdYGUAtTy/F1ErLst+sldAfUt/qsGl2z3l7NHFBFGzyyyybQoFcz8Hfjh8MPjv4PHxA+EPi+013RWvLi0TUbEsY2mglaKVAWUZ2yIyEjjKkdjXm/xP8A2ofik/xzm/Zx/Zs+EGmeK/EGl+HrbW/FN94l8USaRpml2lxJNHBEZYrO7kluZGt5yIhEFCx5aRcqDzX/AATI1HU9W+DvjO/1nQrnS7t/jT43+1abdzRySWkn/CQ32Yi0bMhKjHKnHp6A+oOnh5VqujVmtejv0+QdD6eU5APtS0ifcH0pa41qhBRRRTAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAI5BgECvij9vzwjrnxr/ar+HPwe8Aa7pJ1vSPDeq67JpOq+MNR8MXF3Ezw24nsdW0pheRTQ5lSSKMNE0V5mUZEIP2qwO7IFcf8X/gJ8D/ANoLQ4vDHx0+Dvhfxnp0Ewmg0/xV4ft9RhjkAwHVLhHUMATyBmuvAY1YLEqpa+jXpdWKXwtdz55/Zi+M03wt8AeOvgb4X/ZMe78WfDq/tV13w58L/HkPiBNTutRRpTK+qa02nvNe4JmufthWXEkZ3ymRM9H/AMEzYPi34S/Zy0n4PfF39nLxV4B1Dwrp8Vu03iHUtHuItTd3mZzA2nX1y2I/l3ecsXEiBN4DbPbPhp8HvhX8GPDEHgj4R/DnRPC2jWxJt9H8PaTDZWsRPUrFCqoucDoK33hCvuVe1XWxilTnFR+Jptu97q/m+5jUg9LHy7pOm/tGfsjeJ/HHhL4dfszax8R/DfjLxde+IvD2s6H4k022bS7q/IkuLXUVvZ4pEgjud8iTW63LeTIE8ndCBJh3v7LPxk+Dfwi+Dev+G/DDeOPE3w9+KupeOPGOg6VqkEEuoS6pa6wt7b6bJeNDCRFcasTEs8kKvFEwLRswWvsRVYLjBoCbei4q/r0+RJRS7766W1183t3KUtD5L8K2v7Ui/tS/Ej9oXW/2Wb+10LxR8HtG0vSNEXxBp39tG+sJ9XmNq8fn/ZBI4vUGftRhU7P3py/lcDqnw7/aktf+CUnwg+BMf7Ifi658ceFbjwXaa54Xh1/w/wCdBHod/Y3Fxdec2pi3aGRLR1jCytLumj3xoN5T7zwfQ0YPoaueYyk4tQircv8ANryqy+12/IHLQ+YtT8PfG/4G/tReI/j14J/Z91nxjo3xN8MaPZ6jpmh3+mQahoGoWX2r95cPeXkMUlq0d0FIhZ5I3gYrHKJNy7H/AAT9+HHx3+GHgvxxoHx28F2WkXOo/FfxLrelvY6otwlza3+pz3aPtCgx4Eqj5vmYDJVDlR9C4PoaFU7hkVFbHTrUXCUVdqKv191WXXtoRBcl/MmQ5QfSlpF6ClrzjUKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBCqntRsX0paKAE2L6UuB6CiigBNi+lGxfSlooATYvpRsX0paKAE2L6UbV9KWigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAP/Zl2fWAiBwxRo=",
      "name": "Modelica.Blocks.Examples.TotalHarmonicDistortion"
    },
    {
      "id": "2",
      "image": "/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAIBAQEBAQIBAQECAgICAgQDAgICAgUEBAMEBgUGBgYFBgYGBwkIBgcJBwYGCAsICQoKCgoKBggLDAsKDAkKCgr/2wBDAQICAgICAgUDAwUKBwYHCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgr/wAARCAPvA/ADASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAoooo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KKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAjr47/AOCEH/KMLwd/2OfjX/1LdXr7Er47/wCCEH/KMLwd/wBjn41/9S3V69Ol/wAiir/18p/+kVipH2On3B9KWkT7g+lLXmEhRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRkeooyPUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUj/dNLSP8AdNADK+O/+CEH/KMLwd/2OfjX/wBS3V6+xK+O/wDghB/yjC8Hf9jn41/9S3V69Ol/yKKv/Xyn/wCkVipH2On3B9KWkT7g+lLXmEhRRRQAUUUUAFFFFABRRQTgZoAKKrteYJGDx7ULdljtH8qdmGxYyPUUZHqKjHIzRU3QElFR+cfejzj70XQElFR+cfesLxl8Tvh58PFtpPHnjjSdGF5cRwWn9qajFB50sjhERN7DczOyqAOSSBR70pKMVdsDoaKKKYBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFF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xqMl7dRxwEJOpuLozgHpvivwp4W8eeFtS8DeOfDWn61omtafNY6xo+rWaXNrfWsyGOWCaKQFJY3RmVkYFWViCCDX5UfsaeHvA/j34L/8EzfBXi7/AIJyX/hrTvDOoWF9oPjzxHpfhSbTb+6k+HuualLcWSWWo3F9DJdahbW2qb5raBmmtI5pSlxHGK/VvxZ4f1HxH4W1Pw9o/i/UNBu7/T5re113SYrd7rTpHQqtzCtzFNC0kZIdRLFJGWUb0dcqfA/B/wDwTd0DwP8AD/8AZ6+GuhftR/E4ab+zfqC3HhEypoDSaxHHYT6XBbaix0r544tNurqyU24t5GjuGkd3uFjuEAOO+Kvj/wD4Kk6F8adF+EHgj4z/AADj1rxp4g1q/wDDnh25+D+t3iaT4PsNSgil1O+1L/hIbZZrmO0vtNQ28NqGlvr1IlCWqXF9b/YNfL/jH/gmv4r8VeO/iN8QNK/4KOfHvw3d/E7MOuf8I2PCcM1hYrFJFbafp97LoEl9Y21skszQLFcBo5ri4uQxubieeT6goAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigCOvlD/gif/yjy0D/ALKF49/9TDWa+r6+UP8Agif/AMo8tA/7KF49/wDUw1mvRpf8iur/AI6f/pNQmJ9Zr0H0ooXoPpRXnFBRRRQAUUUUAFFFFADNrjp/Ogq56/zp9FABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAEdfKH/BE//lHloH/ZQvHv/qYazX1fXyh/wRP/AOUeWgf9lC8e/wDqYazXo0v+RXV/x0//AEmoTE+s16D6UUL0H0orzigooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAjr5Q/wCCJ/8Ayjy0D/soXj3/ANTDWa+r6+UP+CJ//KPLQP8AsoXj3/1MNZr0aX/Irq/46f8A6TUJifWSfdFLSJ90UtecUFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAEdfKH/AARP/wCUeWgf9lC8e/8AqYazX1fXyh/wRP8A+UeWgf8AZQvHv/qYazXo0v8AkV1f8dP/ANJqExPrJPuilpE+6KWvOKCiiigAooooAKKxPHnjnQfh54P1Txv4kg1GSx0m0a5vE0nSLnULlkXkiK2tY5Jpm9EjRmPYV4Z4G/4KpfsVfED/AIR650nx34k0+y8WTwQeG9d8TfDPxDo+lajLOMwJFf31hDbM0mQEHmfOxCrknFb0MJjcTd0qbkl1Sb11evyTfon2C6W59G7F9KNi+lQi7UcA0pugpwTWAEuxfSjYvpUIuwTgGj7WOuaV0BPRUH2tfX9KPta+v6U7oCeioPta+v6Uq3Bb7vNK6AmoqITkgMOhGRSNchThjRdATUVALsE4Bo+1rjOeDT6XAnorz/47/tF+Cf2etD0bxB44sdVmttc8T6doVrJpmnPOsNze3MdtA8xUYiiMssalzwC4HJIFd2khdA4J5Gat05qCm1o9vluBLRQOgoqACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAjr5Q/4In/APKPLQP+yhePf/Uw1mvq+vlD/gif/wAo8tA/7KF49/8AUw1mvRpf8iur/jp/+k1CYn1kn3RS0ifdFLXnFBRRRQAUUUUAZHiCES6ZfQ4+9bN/Kvy18Eaz8cdT/wCCQfwm0X9oD4Y6B4f+Btn4f8L33irxr4a8XS6nrS6ZZ3lpdRSyadLY2yWkfmQxmd47i6aKJXCRSEh1/T/4mP42i8E6o/w30HS9T102jLpthrOryWFrNKeAslxHb3DxL/tCJyP7pr4e0X9i7/goP4u/YF0H/gm18Q9E+FnhPw0PCNp4Z8VePvD3j/UNYv59NVVjuRa6fNo9rFHLLGpVXe5YRbs7ZMYP0+QY7D4al+9ko2q05Xbd4qKneSt1jdNaPXoyJq6+8+gPhp8U/jbrf7XnxU+CviXxboM2gaJ4J8N6z4MFl4faOWz+3yatFILuQ3Mn2ps2CsHQQKAduzILHB/Zt8Z/ta/tB/sqXOq2fxp8G6R8QIPHXiTRbjxLcfDu4u9NMOna7fWCtFp66jE8ZaO2TG66k2nO7eat+KvgV+0H8MP2kbr46fs4aR4S8Q2PiLwZo3hvXNH8WeILrS5tMXTZ76SG8t5oLW6+1lxqMoa2dYAPIQib5iB0H7DHwJ+L/wCz58N9Z8HfFjxVoGrS3njfXtZsbnQ7KaDEWoapdXxSVZHb5w1weFJCj5cuV3ty1KmGWGdSm439xpNK7aT5tHfS9rp6PzHU/h2XeP8A6Sv1OA/4JPp+1ZrH7J3w++IPxq+N3hPxHpOs+CrS7stP0vwJNYalFPKokaS4vDqU0U+dzZCWsOSQflxtrkNE0H/goBqH/BRvxr4StP2pfhsILf4baDfwQXHwYu5IIrOfUNUXyV266snnfuPnlZyj/u9sUext/r37GPwc/aN/Z5+HOg/AHx5N4ZXwr4D046ZoGs6PqktxfeIYFJWGW5t5bWNNPKJ96OOa43sVIdFUq8PxI+DXx+0D9qG4/aF+BkfhPUE8S+ELHw3r9v4m1S5tH0eO0uLqaG9tlhtpxeE/bJQ9s5t93lpidcnBiMbF5pippwcanNy+5F7yTWlrJtdlp17JPSlL5fmjI/ba/bu8J/s5/ELw18CU+Nnw3+H+t+J9HvtVk8YfFDUkTTdJtbZokUC2NxbNezTSShViWeIKkU0hfKLHJQ/ZH/bV8YftbeE/iJ4J+G3xL+GPiTxb4C1izsrbxv4QklvfD2twXMazpcrAlw0ls4QTQtAZ5dksWfMZTiuv+NfwC+Lz/Frwj+1L8CL/AEa/8aeGtAvdC1jSPFlzJa2fiLTro28hVrm3ilNjNHPbRypIkEq4eZDH86svTaddfteaX8M77xTfeBvAur+N7iWE2Hg5fF11Z6RZRbgsinVf7Nknnfbuff8AY41OEiCIFMrZp5bHBRioKU2t+baSnfVWs04pK11HW+6aNFZnkn7N/wAbP2/v2mfgr8Pfjj4X0f4caNa3M9oPGOl6zZ3pbWIPP8u+l0+RJc2QiRXMAnSb7QwG4wIVkbzbW/28/wBqH4cfETw54H+KfxY+Glqp/aM/4Vzr9tbfDvUIpntLiwjvdPaBhq8ix3EkckO6Rw0Y+1KfLXyCJfff2APhz+0v8F/gRpvwS/aD8D+CtLHhuyjg0zVPB/je51Y35Z5GcyxT6dafZ8fIFCtLuyfuBRny/wDbE/4Jf6v+0v8AEH4p+O9L+ISaWnjD4c2lj4ahlEsT6T4kgnEv9r+bGN8TKlppyoUOR5LnYDy3pU6uRzzuvCqlGhd8jir6OWmj7Rd97u1u1jT+vS36mx4Z/aE/a8i8DeDvBWoav4C8Q/ED4oeKNVTwdqmn6Dd6dpGleH7RWk/tO7t5L+eW6YxLE4himiLNewRExhJJq634efGX9pPRPirrH7MPxru/CN/4uk8Iya/4O8X+HtFubLS9UijlWCeKaxluria3e3kmtd2LiQSpcKVKMGQXPiv+y74ntrH4d+Mf2dp9FsfFfwnimtfDOn680qWF/p09ultc6fLLEGe3V1ihdZkSQxyW8RKOu9HpeEvgv+0TqXxP1T9q/wCL/hnwV/wndh4Tn0HwH4N0fxTdyaXYwSyRz3BuNTexSWR7mW3tNzLZYgS2UKspZmPBWqZdaPs0tpN3W0uZ8tttOTl+fMYxupHCfC/4/f8ABQv47eAfFqfDew+F+meJ/BXjfVtFk1HXdK1B9K1o21y0cdvCsdwJbfMYVpLs+csbOEWGUh9qft92/wC2HD8d/gq3wo/aC8GeH9C174kJp9tourfDa41GVbxdG1OfzZ549Ut/PhzDkQokDBtrNK4Uo/W/sMfDf9rr4U6r4o0T4+fC34f6Zouv+I9V8Q2upeFPiJe6pcLcXd0JFtmgn0i0XYFLkzebnKriPDEr2f7XHwW+I3xTsPBHjb4PLot14l+HnjWPxDpej+Ir2W0stTzZ3VlLby3EUUz2+YryR1kEUu141+Q5yG8XSw+Z+5CCjZrbmTvFLaV1vr5dLII80qM16/oeZf8ABQCL4qaB+xXpK+OdZ0TxT4utPih4HZ30LS5NIs76b/hLdM2IkU1xdNbggopLSycgngHaOv8AAnxm/aW8EftU+HvgJ8e9S8IazZeOvDGravoF14V0C4sH0h9PkshNaTtNdz/bdy36FbhUtv8AUnMXzjbkftXfBz9r74wfs3aL4T8J+GvA194tHjrQvEGr6dqfjC6sdL06DTdWtdTFrb3MenTS3BY2iQGV4YyTNJNtUKsFafxB+G/7Rut/tu/Cv416V4L8Gt4T8JeFNa07xDNP4yuk1BLjUhZM/kW404xypEbFArPNGZPPYlY/LAkulPD1MNGlVcbfvZeafIuRLprJdl9xqruKv2R9HDpRSRsXRXIxkA4pa+eWwBRRRTAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAoo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PxN4B0fR9A05Tcabp11pkduzKoJ+13YYAooGTHbtgRj55R5m2O2NO8CeCNcmXxv4k8B6Po+gabG82m6Zd6XFAXUDP2y7DKDGFXJjgbAiB8yUebsS2d2gPEvEP7CH7N3gPWp/Hv7FnxvtvgV4hll8908CalA2g302c/6ZoUr/AGGZSeS0SQz9dsy1s/Az9vu90b4hWX7N/wC2pB4Z8K+Nr93Twv4q0DVvN8M+MQg3Ys55Dm1vNvzNYTEuB80TzoC49UsfAfgXXmXxn4m8AaPpGgacrXGm6dd6ZHbsyqCftd2GAKKBkx27YEY+eUeZtjtvOf2pf2avhH+0p+z74xT4m/CnTrXw/D4du59E06KwFnei4SCRoNRkdAssEsbgPBGCrRMBI487Ylvsq8WuWrr+YH0lb3O5VKuGVhkHPBHrVgEEZFfLn7Cnxo+JNlqWtfsb/tGao198QPh5p9pc2fiNgEHi3w9cl1stWC8BZwY3trpFG1LiFiMJLGK+nI95A3fjWVWPs6nL0AnobofpQOgoqQIWYg4Br5n/AGYLm1n/AOCi37T1rFcI0kP/AAhfmxq4LJnSZcZHbI9a+mq/PLwR+xz8J/2kf+CnX7VnizWdU13wv4y8Pa14PHhrx74M1mTTdX00SeGbffGsqZSeFtq7re4SWB8fNG1ejgYxlQxF3b3F/wClwA/QoSBSFZsZ6VMpBUEelfC/wG/bq+O+g/tSW/7Ib3OmftI6ZBP5OqfE34a6atlP4VYD/V68oJ0zzs5BNtcxzMUbFiuQT90jHauOrRnRtzdVdenpuvRpMAooorIAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAjr5q/4JVD/jD6IlcgfE/wAf9v8AqctZr6Vr5w/4JPgH9jqPI/5qf4//APUy1mvQp/8AIqq/46f/AKTUJifRoKAY2/pRuT+7+lTUV5ZRDuT+7+lGUPG2pqKAPyb/AODw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QAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAR184/8Env+TOo/wDsp/j/AP8AUy1mvo6vnH/gk9/yZ1H/ANlP8f8A/qZazXoUv+RVV/x0/wD0moTE+kqKKK85bFBRRRTAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAjr5x/wCCT3/JnUf/AGU/x/8A+plrNfR1fOP/AASe/wCTOo/+yn+P/wD1MtZr0KX/ACKqv+On/wCk1CYn0lRRRXnLYoKKKKYBRRRQAUUUUAFFFFABRRRSl8IBRRkdM0ULYAooopgFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAR184/8Env+TOo/wDsp/j/AP8AUy1mvo6vnH/gk9/yZ1H/ANlP8f8A/qZazXoUv+RVV/x0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kkHmhVt+U+NlnfeO/g74t8R+I7WW20yy8NXtzo2h3Ee2SR1gdo7u5RhkMCA0UDD92QJJB5oRLfpJZ/8AhJFbxp42H9laDpeLvTdP1AeVynzLfXitjZtK7ooW4iwJZP3wRLblfjFpt142+EXirxL4gtJbfSrbw1ezaNpFzGUluJBA7R3lzGwyu1l3Q27f6s4kkHm7EtwDzP8A4KTWNz4O/Zi0L9q/QEaPVPgZruneNIZIlz/xLoQbfVoio6rJpVxfpx0JU4OMH6js7q3vbeK9tJlkimRXikU5DKRkEfga5Lxt8PdK+LvwW1j4W+IIFlsPE3hOfSr1JBlWiuLdonz/AMBc15//AME0vHusfEr9gL4P+LPE07SawfAOnWeul2yw1C1iFtdA+4nikBrsnJywqf8AK7fJ7L7038xnu6nIzS0ijCilrjER18cfs8/AL4MX3/BVr9oz416h8N9JvfGVpD4QgtfEd9YpNc2UbaQyOsDsCYdyxxhtpG4IA2QAK+x6+av2aP8AlIZ+03/108G/+mmWu/LpyhRxNn9hf+nIAfSkYVUCoMAdAO1LQvQfSiuACSiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAMD0owPQUUUAGB6V8yaf/AMperz/s3iH/ANPktfTeR0zXzJp//KXq8/7N4h/9Pkta0vtejIluj1P9qH/k37xB/uWv/pRDWoi/8JF5njLxs407QtMP2qztL8+UMRjf9tut+PLZMbo4m/1WBJJ+92pb5f7UP/Jv3iD/AHLX/wBKIavO58VlfGHjGE6ZoGmn7VpmmX/7rd5fzi+vA+Nm3G+OFv8AV4Esn73alvitilsSGZ/Fyp4z8ZQNpmgaeBd6bpWojyiTH84vbwN9zbjfHC3+rwJZP3uxLeG68rX5B418bj7Bo+mMbrTLC9cxb9g3C8uw2Nm0jckTf6vAkf8Ae7Vt5XkbXmbxd4yUWOi2B+02NjfN5edg3i9ug2PLKEbo4m/1eBI/73aluwNP4jiTxr4+Yafo+m/6Xpml3o8nGz5lvLsNjaVxujhbHlYEjgShFt2Mcgbxe/8Awnfjpf7K8O6V/pemadqP7nHl/ML67D48srjdHE2PJA8yTEu1LYknXxKjeNvG+NK0DSwLvTdO1L90Mp8yX14rY2bSu6KFuI9okk/fBEtkklPipG8aeM9ml6Dpai70zTNTPkglDuS+vQ2NgUrvjhbiPaJZP3oVLaOOO48TXZ8YeOLc6f4dsD9s07TtSPlFig3C+uw+NhXbujhf/U4Ej4lCJbAAA3ir/itPGQGmaBpo+2aZpupnyclPmW+vA2Nm3G6OFsCIL5jgShFt+Z+L+lX3xA+E/irxDrtvNb6ZZeHb250fRrmLa80iQO0d3cowypUgNDA3+rOJHHmhEt+lvrhvEu7xh4wxYaJpn+lafYXreUDsG5by6VsbNuC0cTf6vAkf97tWDA+MMF948+EHizxb4psJbTRrbw3ey6JpFwjRzzyCByl5cocFWDANFAwzGVWRx52xLcA9J8Llj4fs2b732KPP5V89/wDBNFT4X8C/FD4INx/wr/46eKrCCHvDbX96dbtY/wALbVYAP9nb0xgfQfhUg+HLIjOPsMXX6V8+fs8svgv/AIKS/tDfDj7kXiTQfB3jiEZwJZZba80WYj1Krolrn2dPx6qS58PUj2Sf3NL9WB9ML0H0ooorkTQEdfNX7NH/ACkM/ab/AOung3/00y19KkY4NfNX7NH/ACkM/ab/AOung3/00y13YH+DiP8AB/7fAD6YT7opaRPuilriAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAobofpRRQBA7sHCg8Zr5r0/8A5S9Xn/ZvEP8A6fJa+l6+aLD/AJS9Xn/ZvEP/AKfJa1pfa9GZyXvJ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ooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAwPSjA9BRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAf/2Zdn1gIgQNUa",
      "name": "Modelica.Blocks.Examples.Filter"
    },
    {
      "id": "3",
      "image": "/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAIBAQEBAQIBAQECAgICAgQDAgICAgUEBAMEBgUGBgYFBgYGBwkIBgcJBwYGCAsICQoKCgoKBggLDAsKDAkKCgr/2wBDAQICAgICAgUDAwUKBwYHCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgr/wAARCAOpA6sDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD93fht/wAi/df9jBqv/pfcVvL0H0rA+G3/ACL91/2MGq/+l9xW+vQfSgBaKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAOf8A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LMUyGbA3EjI5OM14R/wk3/BX3/oj37Of/hxNc/8AlVXpH7F/x31L9qT9kf4aftIax4fh0m78d+CNN1250u3ufOS0e6tkmaIPgb9pfGcDOOgrLFVcZWoU3OScaSUI2tom5StotdXJ6idNLVGH+y1+zp8c/gj4f0jwZ8VP2jbLxlpHhfRYtK8Mxaf4PfSJ3gjRYkl1GT7bOt7OI0ADRpbx5eRjESY/L4e58EfDz9sr9rXwZ8U0+G/jbR5vgpqmtQ39/wCI/D1/o8GotMI4oreIXESJqVs8sEV6ssbOsbWUBODKBX0+yZBDDGetKoC/dGM+lTDH4qFeVZO02mtElpJWei02bWgRXLoEEAQbmHP8qmVccmhV7mnVwxjbVjCiiirAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAwPhr/wAi7c/9jBq3/pwuK3h2+lYPw1/5F25/7GDVv/ThcVvDt9KAFooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKT2A+av+ClX/IlfCT/s4nwP/wCneGvo3yv9n9a+cv8AgpV/yJXwk/7OJ8D/APp3hr6RrWX8KPz/AENG7IasSjkjmnUUVlZEN3CiiiiyC7CvKv2tvFf9kfD6Lwrbyss+t3SxsV/54RkPJz2z8o/GvVa8P/aX+FvjvxdqUvjq21KxGm6Vp7eXZu8gmAALuduwjJx6jgD0oshrU7r4E+Mh44+Gml6zdPm7jX7Lec/xpwD75XafxNdoYkzyv615r8CfhN40+Fbajpuu6jY3NndtHLbNbTOWjcE7sqyADOexPSvS/m7kflRZDDYvpRsX0pcv/eoy/wDeosguJsX0o2L6UuX/AL1GX/vUWQXE2L6V84f8Ef8An/glj+zyGP8AzR/QP/SGKvpCvm7/AIJAkj/gld+z1j/ojvh7/wBIYq64pfUan+KH5TMqrtY+kcD0owOuKVup+tJXE3dkklFFFWaBRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFAB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",
      "name": "Modelica.Blocks.Examples.NoiseExamples.Densities"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|返回数据对象数组|示例：-|
|»» id|string|true|none|序号|示例：1|
|»» image|string|true|none|图片base64数据|示例：/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAIBAQEBAQIBAQECAgICAgQDAgICAgUEBAMEBgUGBgYFBgYGBwkIBgcJBwYGCAsICQoKCgoKBggLDAsKDAkKCgr/2wBDAQICAgICAgUDAwUKBwYHCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgoKCgr/wAARCAMSAxIDASIAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwD9/KKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooJAOM0UAFFGR60UAFFGR6iigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKG6H6UUUAQsTnGaepLD5a8V/bK/bx/Zp/Yl8PWOo/G34p+HNK1HVL21g0rQ9U16G1ubpJrqK3adVbLCGIyb5JduxFRizKASOI/bV/aM8JfEv/AIJu/Fr44/smftF6bqD+HfBmq32meLPh/wCIrW+jgvbS1acR+dEZI8jCbl6gN1Bwa6KWExNRxk42jJ2T79/uuh2Pp3zCeQeKeAScCvNvA/7SPwR8T+NJPg94e+MnhXVPGOl6bFLrPhSy8R2suqWK7VJaa1RzLGPnXllA+YeoqH47fG/4JeDbOH4Y/Ej9obQvA2t+K7aSDw5Hd+KLOw1K6lyFBs0uD++cMyjCq/JA2nNJ4erGqoNbhY9M3nOM0+vmX9jfxp+0F4+/4JsfDTx14V8W6fqXjrVfA2mXtzq3jaOW8ivZ2t0MjSmCSJtzHnIOOeFrrf8A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ObP9of4t689orhdM8UfEe81GykLKVy8EpKkgE4PGDXa/Bnxb8W/G/hP+1fi78ILTwbq32hlOkRa3HqIEYC4fzo1VeSWG3HG33rs4bbZklFB74FcNeviXFwlLX/t0G2TDkc0UDgYorkWwgooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigApGOBmlpGIKkZoA89b9qj9mrP/JwHgn/AMKa1/8Ai6P+GqP2av8Ao4DwT/4U1r/8VWe/7Hn7IhYqP2Xfh8een/CHWX/xqmf8MV/siHn/AIZT+Gx+vg6y/wDjVdX+w/3h6GrD+1H+zfdTLZ2vx68FyTycRRL4mtcufb567uKQSoHAx6j0PpXmVt+xn+yZZ3Ed7Zfsu/Dm3uIJBJbzxeDrIPFIPuup8rII9Rg+9eiwKIWEMYCqvAVeAKyqRoS/ht/MReIBopF+7S1kAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQAUUUUAFFFFABRRRQBHQSB1OKBycV498fv2gvGngL4g+Gfgp8Gfh1pnivxt4rsNQ1KwsPEPiN9H0220+xktY7m4nu47W6kUiS9tY0jjgkZ2mGdiK7rUITqT5YrowPXRcxscKRmneZ7V81/A/9rT9pT4v2fjpL39k7Q9K1HwNrC6LPZW/xI+0pdaog33UQkNimyJIXtp45SpMiXIDJDIkka1fgd+2/8fP2hPhz4J+LXws/Y9+2aFrN/Z23jN7rx7BDNowlmWOeSzH2cx6ktpmRpz5lvzGUh899yL2PLcZGm6jUbK1/fj126gfTp2E5KCkACjAFeA61+1h8dPF3jnxR4c/Zm/Zq0zxfpvge+Fh4m1XxD46bRGuL0RRzSWmmx/YbhbuSOOWMM00ltCJX8rzcpKY6S/ty6v4q1/4E6v8ADr4e2sng74veJLvRtTv9Z1GW31fRLyHTb+7Nq9h5BTesljLDKWnVo3UhY5A29cpYHFNKX5NPpfXtomOx9FFQTkilrw7Sv2lvizL+0b8Vvgnqnwp0QReCPA+keIvC8tj4kmkl1pL19RiEdxutVWyYSaeRhfPG2QOW6oOH8N/8FEviRrPwb+Hf7T+r/sxJpXw18cJoMVzqF54wJ1nT7jVp4be2dLBbIxT2oluIVaZrmKXaWdYGUAtTy/F1ErLst+sldAfUt/qsGl2z3l7NHFBFGzyyyybQoFcz8Hfjh8MPjv4PHxA+EPi+013RWvLi0TUbEsY2mglaKVAWUZ2yIyEjjKkdjXm/xP8A2ofik/xzm/Zx/Zs+EGmeK/EGl+HrbW/FN94l8USaRpml2lxJNHBEZYrO7kluZGt5yIhEFCx5aRcqDzX/AATI1HU9W+DvjO/1nQrnS7t/jT43+1abdzRySWkn/CQ32Yi0bMhKjHKnHp6A+oOnh5VqujVmtejv0+QdD6eU5APtS0ifcH0pa41qhBRRRTAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAI5BgECvij9vzwjrnxr/ar+HPwe8Aa7pJ1vSPDeq67JpOq+MNR8MXF3Ezw24nsdW0pheRTQ5lSSKMNE0V5mUZEIP2qwO7IFcf8X/gJ8D/ANoLQ4vDHx0+Dvhfxnp0Ewmg0/xV4ft9RhjkAwHVLhHUMATyBmuvAY1YLEqpa+jXpdWKXwtdz55/Zi+M03wt8AeOvgb4X/ZMe78WfDq/tV13w58L/HkPiBNTutRRpTK+qa02nvNe4JmufthWXEkZ3ymRM9H/AMEzYPi34S/Zy0n4PfF39nLxV4B1Dwrp8Vu03iHUtHuItTd3mZzA2nX1y2I/l3ecsXEiBN4DbPbPhp8HvhX8GPDEHgj4R/DnRPC2jWxJt9H8PaTDZWsRPUrFCqoucDoK33hCvuVe1XWxilTnFR+Jptu97q/m+5jUg9LHy7pOm/tGfsjeJ/HHhL4dfszax8R/DfjLxde+IvD2s6H4k022bS7q/IkuLXUVvZ4pEgjud8iTW63LeTIE8ndCBJh3v7LPxk+Dfwi+Dev+G/DDeOPE3w9+KupeOPGOg6VqkEEuoS6pa6wt7b6bJeNDCRFcasTEs8kKvFEwLRswWvsRVYLjBoCbei4q/r0+RJRS7766W1183t3KUtD5L8K2v7Ui/tS/Ej9oXW/2Wb+10LxR8HtG0vSNEXxBp39tG+sJ9XmNq8fn/ZBI4vUGftRhU7P3py/lcDqnw7/aktf+CUnwg+BMf7Ifi658ceFbjwXaa54Xh1/w/wCdBHod/Y3Fxdec2pi3aGRLR1jCytLumj3xoN5T7zwfQ0YPoaueYyk4tQircv8ANryqy+12/IHLQ+YtT8PfG/4G/tReI/j14J/Z91nxjo3xN8MaPZ6jpmh3+mQahoGoWX2r95cPeXkMUlq0d0FIhZ5I3gYrHKJNy7H/AAT9+HHx3+GHgvxxoHx28F2WkXOo/FfxLrelvY6otwlza3+pz3aPtCgx4Eqj5vmYDJVDlR9C4PoaFU7hkVFbHTrUXCUVdqKv191WXXtoRBcl/MmQ5QfSlpF6ClrzjUKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigBCqntRsX0paKAE2L6UuB6CiigBNi+lGxfSlooATYvpRsX0paKAE2L6UbV9KWigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKAP/Zl2fWAiBwxRo=|
|»» name|string|true|none|模型名称|示例：Modelica.Blocks.Examples.TotalHarmonicDistortion|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 进入用户空间

POST /user/userspace/login

```text
暂无描述
```

> Body 请求参数

```json
{
  "space_id": "e05f38e5-e598-4e6c-a05c-c2cada524c7a"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» space_id|body|string| 是 ||用户空间id|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "初始化完成",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none|返回数据对象|示例：-|
|» msg|string|true|none|消息字段|示例：创建成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 删除用户空间

POST /user/userspace/delete

```text
暂无描述
```

> Body 请求参数

```json
{
  "space_id": "d0ee5315-706a-4f51-9a69-8ee236cfa236"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» space_id|body|string| 是 ||none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "删除成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none|返回数据对象|示例：-|
|» msg|string|true|none|消息字段|示例：删除成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 获取用户最近工作空间

GET /user/userspace/recent

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "id": "5e98a09f-78ac-4f8e-b85f-48d9a00b1ebd",
      "name": "test1"
    },
    {
      "id": "94f0a4b7-8b04-40e4-9adc-74d9472b6600",
      "name": "test2"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|返回数据对象|none|
|»» id|string|true|none||用户空间id|
|»» name|string|true|none||用户空间名称|
|» msg|string|true|none||消息字段|
|» status|integer|true|none||状态码字段|
|» err|string|true|none||错误提示字段|

## GET 用户登录

GET /auth/test/login/{name}/{password}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|name|path|string| 是 ||none|
|password|path|string| 是 ||none|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# yssim-go/仿真

## GET 获取仿真参数

GET /simulation/options/get

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取仿真状态

GET /simulation/state/get

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|package_id|query|string| 是 ||模型包id|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 设置仿真参数

POST /simulation/options/set

```text
暂无描述
```

> Body 请求参数

```json
{
  "model_name": "we.Filter1",
  "package_id": "dd7f135b-441c-42e1-ad7b-a6cd1407e05f",
  "experiment": {
    "startTime": "0.0",
    "stopTime": "0.9",
    "tolerance": "1e-06",
    "numberOfIntervals": "500",
    "interval": "0.0018",
    "simulate_type": "OM",
    "method": "dassl"
  }
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» model_name|body|string| 是 ||模型名称|
|» package_id|body|string| 是 ||模型包id|
|» experiment|body|object| 是 ||-|
|»» startTime|body|string| 是 ||横坐标的开始时间|
|»» stopTime|body|string| 是 ||横坐标的结束时间|
|»» tolerance|body|string| 是 ||保留小数点后几位|
|»» numberOfIntervals|body|string| 是 ||步长|
|»» interval|body|string| 是 ||步长间隔|
|»» simulate_type|body|string| 是 ||仿真编译器类型|
|»» method|body|string| 是 ||仿真算法|

> 返回示例

> 成功

```json
{
  "data": [],
  "err": "",
  "msg": "设置成功",
  "status": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[string]¦null|true|none|返回数据对象数组|示例：{}|
|» err|string|true|none|错误消息|示例：-|
|» msg|string|true|none|消息字段|示例：设置成功|
|» status|integer|true|none|状态码|示例：-|

## POST 仿真计算

POST /simulation/simulate

```text
暂无描述
```

> Body 请求参数

```json
{
  "package_id": "6d602d98-01b2-4625-a0cd-940b9a91980b",
  "model_name": "Modelica.Blocks.Examples.PID_Controller",
  "simulate_type": "OM",
  "startTime": "0",
  "stopTime": "10",
  "tolerance": "1e6",
  "numberOfIntervals": "500",
  "method": "dassl",
  "experiment_id": "",
  "username": "wanghailong"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_id|body|string| 是 ||模型包id|
|» model_name|body|string| 是 ||模型名称|
|» simulate_type|body|string| 是 ||仿真编译器类型|
|» startTime|body|string| 是 ||开始时间|
|» stopTime|body|string| 是 ||结束时间|
|» tolerance|body|string| 是 ||保留小数点后几位|
|» numberOfIntervals|body|string| 是 ||步长总数|
|» method|body|string| 是 ||仿真算法|
|» experiment_id|body|string| 是 ||实验id|
|» username|body|string| 是 ||用户名|

> 返回示例

> 200 Response

```json
{
  "data": {
    "id": "string"
  },
  "msg": "string",
  "status": 0,
  "err": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none|返回数据对象|示例：-|
|»» id|string|true|none|用户空间id|示例：26664238-31d7-408a-9f09-ef5aa44f7adb|
|» msg|string|true|none|消息字段|示例：仿真任务正在准备，请等待仿真完成|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 获取多条仿真结果

POST /simulation/result

```text
暂无描述
```

> Body 请求参数

```json
{
  "variable": "inertia1.phi",
  "id": [
    "8210fcc6-c772-4243-b4b9-a4defd14b4df",
    "67815a7a-2436-4751-90d3-aab2aad0c9b5"
  ],
  "s1": "rad",
  "s2": "deg"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» variable|body|string| 是 | 变量名称|数组形式，支持一次查询多个|
|» id|body|string| 是 | 序号|none|
|» s1|body|string| 是 | 变量单位|可以为空字符|
|» s2|body|string| 是 | 变量单位|可以为空字符|

> 返回示例

> 200 Response

```json
{
  "data": [
    "string"
  ],
  "msg": "string",
  "status": 0,
  "err": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[string]|true|none|返回数据对象|示例：-|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 获取仿真结果列表

GET /simulation/record/list

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|page_num|query|string| 否 ||none|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "data": [
    {
      "id": "string",
      "simulate_end_time": "string",
      "simulate_start_time": "string",
      "simulate_status": "string"
    }
  ],
  "msg": "string",
  "status": 0,
  "err": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|返回数据对象|示例：-|
|»» id|string|false|none|序号|示例：f25c9e45-f674-4c72-bc6f-63e4aa8c6f83|
|»» simulate_end_time|string|false|none|仿真结束时间|示例：2022-08-30 16:20:32|
|»» simulate_start_time|string|false|none|仿真开始时间|示例：2022-08-30 16:20:25|
|»» simulate_status|string|false|none|仿真状态|示例：仿真完成|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 修改仿真实验记录

POST /simulation/experiment/edit

```text
暂无描述
```

> Body 请求参数

```json
{
  "experiment_id": "af6625b0-e596-4d6b-b21d-d2654ddedb14",
  "experiment_name": "test1",
  "model_var_data": {},
  "simulate_var_data": {
    "startTime": "0.0",
    "stopTime": "11",
    "tolerance": "1e-06",
    "numberOfIntervals": "500",
    "interval": "0.0018",
    "simulate_type": "OM",
    "method": "dassl"
  }
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» experiment_id|body|string| 是 ||实验记录id|
|» experiment_name|body|string| 是 ||实验名称|
|» model_var_data|body|object| 是 ||模型组件参数数据|
|» simulate_var_data|body|object| 是 ||仿真参数数据|
|»» startTime|body|string| 是 ||横坐标的开始时间|
|»» stopTime|body|string| 是 ||横坐标的结束时间|
|»» tolerance|body|string| 是 ||保留小数点后几位|
|»» numberOfIntervals|body|string| 是 ||步长总数|
|»» interval|body|string| 是 ||步长间隔|
|»» simulate_type|body|string| 是 ||仿真编译器类型|
|»» method|body|string| 是 ||仿真算法|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "实验记录已更新",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none|返回数据对象|示例：-|
|» msg|string|true|none|消息字段|示例：实验记录创建成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 创建仿真实验记录

POST /simulation/experiment/create

```text
暂无描述
```

> Body 请求参数

```json
{
  "package_id": "6d602d98-01b2-4625-a0cd-940b9a91980b",
  "model_name": "Modelica.Blocks.Examples.Filter",
  "model_var_data": {},
  "simulate_var_data": {
    "startTime": "0.0",
    "stopTime": "0.9",
    "tolerance": "1e-06",
    "numberOfIntervals": "500",
    "interval": "0.0018",
    "simulate_type": "OM",
    "method": "dassl"
  },
  "experiment_name": "test3"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_id|body|string| 是 ||模型包id|
|» model_name|body|string| 是 ||模型名称|
|» model_var_data|body|object| 是 ||模型组件参数数据|
|» simulate_var_data|body|object| 是 ||仿真参数数据|
|»» startTime|body|string| 是 ||横坐标的开始时间|
|»» stopTime|body|string| 是 ||横坐标的结束时间|
|»» tolerance|body|string| 是 ||保留小数点后几位|
|»» numberOfIntervals|body|string| 是 ||步长总数|
|»» interval|body|string| 是 ||步长间隔|
|»» simulate_type|body|string| 是 ||仿真编译器类型|
|»» method|body|string| 是 ||仿真算法|
|» experiment_name|body|string| 是 ||实验名称|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "实验记录创建成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none|返回数据对象|示例：-|
|» msg|string|true|none|消息字段|示例：实验记录创建成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 获取仿真结果树

GET /simulation/record/tree

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|record_id|query|string| 是 ||仿真记录id|
|parent_node|query|string| 否 ||父节点名称|
|key_words|query|string| 否 ||需要搜索的关键字|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 删除仿真实验记录

POST /simulation/experiment/delete

```text
暂无描述
```

> Body 请求参数

```json
{
  "experiment_id": "6bb51bb5-d9c3-486f-b73b-9db0e7b46e20"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» experiment_id|body|string| 是 | 实验记录id|none|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "实验记录创建成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none|返回数据对象|示例：-|
|» msg|string|true|none|消息字段|示例：实验记录创建成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 获取仿真实验列表

GET /simulation/experiment/list

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|model_name|query|string| 是 ||模型名称|
|package_id|query|string| 是 ||模型包id|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "experiment_name": "test",
      "id": "584e33fe-2fca-4962-8405-2a8449e0fc70",
      "interval": "0.0018",
      "method": "dassl",
      "numberOfIntervals": "",
      "simulate_type": "OM",
      "startTime": "0.0",
      "stopTime": "0.9",
      "tolerance": "1e-06"
    },
    {
      "experiment_name": "test1",
      "id": "6acd6f17-cd78-45eb-98c4-0c070b2d5ce4",
      "interval": "0.0018",
      "method": "dassl",
      "numberOfIntervals": "",
      "simulate_type": "OM",
      "startTime": "0.0",
      "stopTime": "0.9",
      "tolerance": "1e-06"
    },
    {
      "experiment_name": "test2",
      "id": "af6625b0-e596-4d6b-b21d-d2654ddedb14",
      "interval": "0.0018",
      "method": "dassl",
      "numberOfIntervals": "",
      "simulate_type": "OM",
      "startTime": "0.0",
      "stopTime": "0.9",
      "tolerance": "1e-06"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|返回数据对象|示例：-|
|»» experiment_name|string|true|none|实验名称|示例：test|
|»» id|string|true|none|序号|示例：584e33fe-2fca-4962-8405-2a8449e0fc70|
|»» interval|string|true|none|步长间隔|示例：0.0018|
|»» method|string|true|none|仿真算法|示例：dassl|
|»» numberOfIntervals|string|true|none|步长|示例：-|
|»» simulate_type|string|true|none|仿真编译器类型|示例：OM|
|»» startTime|string|true|none|横坐标的开始时间|示例：0.0|
|»» stopTime|string|true|none|横坐标的结束时间|示例：0.9|
|»» tolerance|string|true|none|保留小数点后几位|示例：1e-06|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 删除仿真结果记录

GET /simulation/record/delete

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|record_id|query|string| 否 ||none|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 获取单条仿真结果

POST /simulation/result/singular

```text
暂无描述
```

> Body 请求参数

```json
{
  "variable": "inertia1.phi",
  "id": "8210fcc6-c772-4243-b4b9-a4defd14b4df",
  "s1": "rad",
  "s2": "deg"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» variable|body|string| 是 | 变量名称|支持一次查询一个|
|» id|body|string| 是 | 序号|none|
|» s1|body|string| 是 | 变量单位|可以为空字符|
|» s2|body|string| 是 | 变量单位|可以为空字符|

> 返回示例

> 200 Response

```json
{
  "data": {
    "abscissa": [
      0
    ],
    "ordinate": [
      0
    ],
    "startTime": "string",
    "stopTime": "string"
  },
  "msg": "string",
  "status": 0,
  "err": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none|返回数据对象|示例：-|
|»» abscissa|[oneOf]|true|none|横坐标时间，单位是秒|示例：{}|

*oneOf*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|»»» *anonymous*|integer|false|none||none|

*xor*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|»»» *anonymous*|number|false|none||none|

*continued*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|»» ordinate|[integer]|true|none|纵坐标数据|示例：{}|
|»» startTime|string|true|none|横坐标的开始时间|示例：0|
|»» stopTime|string|true|none|横坐标的结束时间|示例：10|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 获取仿真结果记录详情

GET /simulation/record/details

```text
暂无描述
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|id|query|string| 是 ||模型名称|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "data": [
    {
      "id": "string",
      "simulate_end_time": "string",
      "simulate_start_time": "string",
      "simulate_status": "string"
    }
  ],
  "msg": "string",
  "status": 0,
  "err": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|返回数据对象|示例：-|
|»» id|string|false|none|序号|示例：f25c9e45-f674-4c72-bc6f-63e4aa8c6f83|
|»» simulate_end_time|string|false|none|仿真结束时间|示例：2022-08-30 16:20:32|
|»» simulate_start_time|string|false|none|仿真开始时间|示例：2022-08-30 16:20:25|
|»» simulate_status|string|false|none|仿真状态|示例：仿真完成|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## GET 获取仿真实验中保存的组件参数

GET /simulation/experiment/parameters

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|experiment_id|query|string| 是 ||none|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": {
    "parameters": [
      {
        "defaultvalue": "0.4",
        "name": "PI.Td"
      },
      {
        "defaultvalue": "0.3",
        "name": "PI.Ti"
      },
      {
        "defaultvalue": "3",
        "name": "PI.wp"
      },
      {
        "defaultvalue": "2",
        "name": "PI.kFF"
      },
      {
        "defaultvalue": "",
        "name": "PI.yMax"
      },
      {
        "defaultvalue": "Modelica.Blocks.Types.SimpleController.PID",
        "name": "PI.controllerType"
      },
      {
        "defaultvalue": false,
        "name": "PI.withFeedForward"
      },
      {
        "defaultvalue": "111",
        "name": "PI.k"
      }
    ]
  },
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# yssim-go/文件

## POST 保存模型源码

POST /file/code/save

```text
暂无描述
```

> Body 请求参数

```json
{
  "package_id": "1c05ac48-d8a4-4364-90e1-7350e8916c37"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_id|body|string| 是 ||模型包id|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "模型已保存",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none|返回数据对象|示例：-|
|» msg|string|true|none|消息字段|示例：实验记录创建成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 上传模型包文件

POST /file/upload/package

```text
暂无描述
```

> Body 请求参数

```yaml
file: C:\Users\simtek\Desktop\CompanyRelated\模型文件\Applications.mo

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» file|body|string(binary)| 是 ||模型文件|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "Applications 包已上传成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none|返回数据对象|示例：-|
|» msg|string|true|none|消息字段|示例：Applications 包已上传成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 下载筛选变量结果文件

POST /file/result/filter/get

```text
暂无描述
```

> Body 请求参数

```json
{
  "record_id": "098ebe63-5038-466e-8067-931d47033ff9",
  "var_list": [
    "inertia1.phi",
    "inertia1.a",
    "inertia1.w"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» record_id|body|string| 是 ||仿真记录id|
|» var_list|body|[string]| 是 ||变量列表|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "初始化完成",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none|返回数据对象|示例：-|
|» msg|string|true|none|消息字段|示例：创建成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 导出fmu

POST /file/fmu/export

```text
暂无描述
```

> Body 请求参数

```json
{
  "package_id": "6d602d98-01b2-4625-a0cd-940b9a91980b",
  "model_name": "Modelica.Blocks.Examples.Filter",
  "package_name": "Modelica",
  "fmu_name": "test",
  "fmuPar": {
    "fmiVersion": "2",
    "fmiType": "all",
    "includeSource": false,
    "includeImage": 0,
    "storeResult": false
  },
  "download_local": true
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_id|body|string| 是 | 模型包id|none|
|» model_name|body|string| 是 | 模型名称|none|
|» package_name|body|string| 是 | 模型包名称|none|
|» fmu_name|body|string| 是 | 导出fmu的文件名字|none|
|» fmuPar|body|object| 是 | 导出参数，本接口不使用，需dymola导出接口开发人员解释|none|
|»» fmiVersion|body|string| 是 ||-|
|»» fmiType|body|string| 是 ||-|
|»» includeSource|body|boolean| 是 ||-|
|»» includeImage|body|integer| 是 ||-|
|»» storeResult|body|boolean| 是 ||-|
|» download_local|body|boolean| 是 | 是否下载到本地|none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 更新模型源码并写入文件

POST /file/update/package

```text
# 获取模型的源码数据
```

> Body 请求参数

```json
{
  "package_name": "Filter1",
  "model_str": "model Filter1 \"12312312312323Demonstrates the Continuous.Filter block with various options\"\n  extends Modelica.Icons.Example;\n  parameter Integer order = 3 \"Number of order of filter\";\n  parameter SI.Frequency f_cut = 2 \"Cut-off frequency\";\n  parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass \"Type of filter (LowPass/HighPass)\";\n  parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState \"Type of initialization (no init/steady state/initial state/initial output)\";\n  parameter Boolean normalized = true \"= true, if amplitude at f_cut = -3db, otherwise unmodified filter\";\n  Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(\n    Placement(transformation(extent = {{-60, 40}, {-40, 60}})));\n  Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, 40}, {0, 60}})));\n  Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, 0}, {0, 20}})));\n  Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, -40}, {0, -20}})));\n  Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, -80}, {0, -60}})));\nequation\n  connect(step.y, CriticalDamping.u) annotation(\n    Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));\n  connect(step.y, Bessel.u) annotation(\n    Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));\n  connect(Butterworth.u, step.y) annotation(\n    Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n  connect(ChebyshevI.u, step.y) annotation(\n    Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n  annotation(\n    experiment(StopTime = 0.9),\n    Documentation(info = \"<html>\n\n<p>\nThis example demonstrates various options of the\n<a href=\\\"modelica://Modelica.Blocks.Continuous.Filter\\\">Filter</a> block.\nA step input starts at 0.1 s with an offset of 0.1, in order to demonstrate\nthe initialization options. This step input drives 4 filter blocks that\nhave identical parameters, with the only exception of the used analog filter type\n(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main options\ncan be set via parameters and are then applied to all the 4 filters.\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\n2 Hz resulting in the following outputs:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/Examples/Filter1.png\\\"\n     alt=\\\"Filter1.png\\\">\n</html>\"));\nend Filter1;",
  "package_id": "144aa9a8-5850-4044-9d37-28b0f3c043c6"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_name|body|string| 是 ||模型包名称|
|» model_str|body|string| 是 ||模型源码|
|» package_id|body|string| 是 ||模型包id|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 创建模型和模型包mo文件

POST /file/create/package

```text
# 获取模型的源码数据
```

> Body 请求参数

```json
{
  "package_name": "test1235",
  "str_type": "model",
  "package_id": "",
  "comment": "",
  "vars": {
    "expand": "",
    "insert_to": "",
    "partial": false,
    "encapsulated": false,
    "state": false
  }
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_name|body|string| 是 ||模型包名称|
|» str_type|body|string| 是 ||类型，是包类型还是模型|
|» package_id|body|string| 是 ||模型包id|
|» comment|body|string| 是 ||注释|
|» vars|body|object| 是 ||-|
|»» expand|body|string| 是 ||继承某模型或包|
|»» insert_to|body|string| 是 ||插入到某模型或包|
|»» partial|body|boolean| 是 ||-|
|»» encapsulated|body|boolean| 是 ||-|
|»» state|body|boolean| 是 ||-|

> 返回示例

> 成功

```json
{
  "data": {
    "id": "c20fb619-9e95-4bf1-a4c1-5db2f3d729a8",
    "model_str": "package test3\nend test3;"
  },
  "msg": "创建成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|true|none||示例：-|
|»» id|string|true|none|包id|示例：c20fb619-9e95-4bf1-a4c1-5db2f3d729a8|
|»» model_str|string|true|none|模型源码|示例：package test3end test3;|
|» msg|string|true|none|消息字段|示例：创建成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 上传模型图标

POST /file/upload/icon

> Body 请求参数

```yaml
package_id: dd7f135b-441c-42e1-ad7b-a6cd1407e05f
model_name: test1
file: C:\Users\simtek\Downloads\Filter.jpg

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_id|body|string| 是 ||模型包id|
|» model_name|body|string| 是 ||模型名称|
|» file|body|string(binary)| 是 ||上传的图标文件|

> 返回示例

> 成功

```json
{
  "data": null,
  "msg": "图标上传成功",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|null|true|none|返回数据对象数组|示例：-|
|» msg|string|true|none|消息字段|示例：图标上传成功|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST mo文件下载

POST /file/package/get

```text
# 获取模型的源码数据
```

> Body 请求参数

```json
{
  "package_id": "41b464e1-d5a1-4e70-a81c-508e04950a54"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» package_id|body|string| 是 | 包id|none|

> 返回示例

> 成功

```json
"package test5\n  model Filter2 \"Demonstrates the Continuous.Filter block with various options\"\n    extends Modelica.Icons.Example;\n    parameter Integer order = 3 \"Number of order of filter\";\n    parameter SI.Frequency f_cut = 2 \"Cut-off frequency\";\n    parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass \"Type of filter (LowPass/HighPass)\";\n    parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState \"Type of initialization (no init/steady state/initial state/initial output)\";\n    parameter Boolean normalized = true \"= true, if amplitude at f_cut = -3db, otherwise unmodified filter\";\n    Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(\n      Placement(transformation(extent = {{-60, 40}, {-40, 60}})));\n    Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 40}, {0, 60}})));\n    Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 0}, {0, 20}})));\n    Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -40}, {0, -20}})));\n    Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -80}, {0, -60}})));\n  equation\n    connect(step.y, CriticalDamping.u) annotation(\n      Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));\n    connect(step.y, Bessel.u) annotation(\n      Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));\n    connect(Butterworth.u, step.y) annotation(\n      Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    connect(ChebyshevI.u, step.y) annotation(\n      Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    annotation(\n      experiment(StopTime = 0.9),\n      Documentation(info = \"<html>\n\n<p>\nThis example demonstrates various options of the\n<a href=\\\"modelica://Modelica.Blocks.Continuous.Filter\\\">Filter</a> block.\nA step input starts at 0.1 s with an offset of 0.1, in order to demonstrate\nthe initialization options. This step input drives 4 filter blocks that\nhave identical parameters, with the only exception of the used analog filter type\n(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main options\ncan be set via parameters and are then applied to all the 4 filters.\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\n2 Hz resulting in the following outputs:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/Examples/Filter1.png\\\"\n     alt=\\\"Filter1.png\\\">\n</html>\"));\n  end Filter2;\n\n  model Filter3 \"Demonstrates the Continuous.Filter block with various options\"\n    extends Modelica.Icons.Example;\n    parameter Integer order = 3 \"Number of order of filter\";\n    parameter SI.Frequency f_cut = 2 \"Cut-off frequency\";\n    parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass \"Type of filter (LowPass/HighPass)\";\n    parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState \"Type of initialization (no init/steady state/initial state/initial output)\";\n    parameter Boolean normalized = true \"= true, if amplitude at f_cut = -3db, otherwise unmodified filter\";\n    Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(\n      Placement(transformation(extent = {{-60, 40}, {-40, 60}})));\n    Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 40}, {0, 60}})));\n    Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 0}, {0, 20}})));\n    Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -40}, {0, -20}})));\n    Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -80}, {0, -60}})));\n  equation\n    connect(step.y, CriticalDamping.u) annotation(\n      Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));\n    connect(step.y, Bessel.u) annotation(\n      Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));\n    connect(Butterworth.u, step.y) annotation(\n      Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    connect(ChebyshevI.u, step.y) annotation(\n      Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    annotation(\n      experiment(StopTime = 0.9),\n      Documentation(info = \"<html>\n\n<p>\nThis example demonstrates various options of the\n<a href=\\\"modelica://Modelica.Blocks.Continuous.Filter\\\">Filter</a> block.\nA step input starts at 0.1 s with an offset of 0.1, in order to demonstrate\nthe initialization options. This step input drives 4 filter blocks that\nhave identical parameters, with the only exception of the used analog filter type\n(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main options\ncan be set via parameters and are then applied to all the 4 filters.\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\n2 Hz resulting in the following outputs:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/Examples/Filter1.png\\\"\n     alt=\\\"Filter1.png\\\">\n</html>\"));\n  end Filter3;\n\n  model Filter5 \"Demonstrates the Continuous.Filter block with various options\"\n    extends Modelica.Icons.Example;\n    parameter Integer order = 3 \"Number of order of filter\";\n    parameter SI.Frequency f_cut = 2 \"Cut-off frequency\";\n    parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass \"Type of filter (LowPass/HighPass)\";\n    parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState \"Type of initialization (no init/steady state/initial state/initial output)\";\n    parameter Boolean normalized = true \"= true, if amplitude at f_cut = -3db, otherwise unmodified filter\";\n    Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(\n      Placement(transformation(extent = {{-60, 40}, {-40, 60}})));\n    Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 40}, {0, 60}})));\n    Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 0}, {0, 20}})));\n    Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -40}, {0, -20}})));\n    Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -80}, {0, -60}})));\n  equation\n    connect(step.y, CriticalDamping.u) annotation(\n      Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));\n    connect(step.y, Bessel.u) annotation(\n      Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));\n    connect(Butterworth.u, step.y) annotation(\n      Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    connect(ChebyshevI.u, step.y) annotation(\n      Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    annotation(\n      experiment(StopTime = 0.9),\n      Documentation(info = \"<html>\n\n<p>\nThis example demonstrates various options of the\n<a href=\\\"modelica://Modelica.Blocks.Continuous.Filter\\\">Filter</a> block.\nA step input starts at 0.1 s with an offset of 0.1, in order to demonstrate\nthe initialization options. This step input drives 4 filter blocks that\nhave identical parameters, with the only exception of the used analog filter type\n(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main options\ncan be set via parameters and are then applied to all the 4 filters.\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\n2 Hz resulting in the following outputs:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/Examples/Filter1.png\\\"\n     alt=\\\"Filter1.png\\\">\n</html>\"));\n  end Filter5;\n\n  model Filter6 \"Demonstrates the Continuous.Filter block with various options\"\n    extends Modelica.Icons.Example;\n    parameter Integer order = 3 \"Number of order of filter\";\n    parameter SI.Frequency f_cut = 2 \"Cut-off frequency\";\n    parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass \"Type of filter (LowPass/HighPass)\";\n    parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState \"Type of initialization (no init/steady state/initial state/initial output)\";\n    parameter Boolean normalized = true \"= true, if amplitude at f_cut = -3db, otherwise unmodified filter\";\n    Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(\n      Placement(transformation(extent = {{-60, 40}, {-40, 60}})));\n    Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 40}, {0, 60}})));\n    Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 0}, {0, 20}})));\n    Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -40}, {0, -20}})));\n    Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -80}, {0, -60}})));\n  equation\n    connect(step.y, CriticalDamping.u) annotation(\n      Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));\n    connect(step.y, Bessel.u) annotation(\n      Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));\n    connect(Butterworth.u, step.y) annotation(\n      Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    connect(ChebyshevI.u, step.y) annotation(\n      Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    annotation(\n      experiment(StopTime = 0.9),\n      Documentation(info = \"<html>\n\n<p>\nThis example demonstrates various options of the\n<a href=\\\"modelica://Modelica.Blocks.Continuous.Filter\\\">Filter</a> block.\nA step input starts at 0.1 s with an offset of 0.1, in order to demonstrate\nthe initialization options. This step input drives 4 filter blocks that\nhave identical parameters, with the only exception of the used analog filter type\n(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main options\ncan be set via parameters and are then applied to all the 4 filters.\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\n2 Hz resulting in the following outputs:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/Examples/Filter1.png\\\"\n     alt=\\\"Filter1.png\\\">\n</html>\"));\n  end Filter6;\n\n  model Filter7 \"Demonstrates the Continuous.Filter block with various options\"\n    extends Modelica.Icons.Example;\n    parameter Integer order = 3 \"Number of order of filter\";\n    parameter SI.Frequency f_cut = 2 \"Cut-off frequency\";\n    parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass \"Type of filter (LowPass/HighPass)\";\n    parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState \"Type of initialization (no init/steady state/initial state/initial output)\";\n    parameter Boolean normalized = true \"= true, if amplitude at f_cut = -3db, otherwise unmodified filter\";\n    Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(\n      Placement(transformation(extent = {{-60, 40}, {-40, 60}})));\n    Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 40}, {0, 60}})));\n    Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 0}, {0, 20}})));\n    Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -40}, {0, -20}})));\n    Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -80}, {0, -60}})));\n  equation\n    connect(step.y, CriticalDamping.u) annotation(\n      Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));\n    connect(step.y, Bessel.u) annotation(\n      Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));\n    connect(Butterworth.u, step.y) annotation(\n      Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    connect(ChebyshevI.u, step.y) annotation(\n      Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    annotation(\n      experiment(StopTime = 0.9),\n      Documentation(info = \"<html>\n\n<p>\nThis example demonstrates various options of the\n<a href=\\\"modelica://Modelica.Blocks.Continuous.Filter\\\">Filter</a> block.\nA step input starts at 0.1 s with an offset of 0.1, in order to demonstrate\nthe initialization options. This step input drives 4 filter blocks that\nhave identical parameters, with the only exception of the used analog filter type\n(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main options\ncan be set via parameters and are then applied to all the 4 filters.\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\n2 Hz resulting in the following outputs:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/Examples/Filter1.png\\\"\n     alt=\\\"Filter1.png\\\">\n</html>\"));\n  end Filter7;\n\n  model Filter8 \"Demonstrates the Continuous.Filter block with various options\"\n    extends Modelica.Icons.Example;\n    parameter Integer order = 3 \"Number of order of filter\";\n    parameter SI.Frequency f_cut = 2 \"Cut-off frequency\";\n    parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass \"Type of filter (LowPass/HighPass)\";\n    parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState \"Type of initialization (no init/steady state/initial state/initial output)\";\n    parameter Boolean normalized = true \"= true, if amplitude at f_cut = -3db, otherwise unmodified filter\";\n    Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(\n      Placement(transformation(extent = {{-60, 40}, {-40, 60}})));\n    Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 40}, {0, 60}})));\n    Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, 0}, {0, 20}})));\n    Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -40}, {0, -20}})));\n    Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n      Placement(transformation(extent = {{-20, -80}, {0, -60}})));\n  equation\n    connect(step.y, CriticalDamping.u) annotation(\n      Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));\n    connect(step.y, Bessel.u) annotation(\n      Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));\n    connect(Butterworth.u, step.y) annotation(\n      Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    connect(ChebyshevI.u, step.y) annotation(\n      Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n    annotation(\n      experiment(StopTime = 0.9),\n      Documentation(info = \"<html>\n\n<p>\nThis example demonstrates various options of the\n<a href=\\\"modelica://Modelica.Blocks.Continuous.Filter\\\">Filter</a> block.\nA step input starts at 0.1 s with an offset of 0.1, in order to demonstrate\nthe initialization options. This step input drives 4 filter blocks that\nhave identical parameters, with the only exception of the used analog filter type\n(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main options\ncan be set via parameters and are then applied to all the 4 filters.\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\n2 Hz resulting in the following outputs:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/Examples/Filter1.png\\\"\n     alt=\\\"Filter1.png\\\">\n</html>\"));\n  end Filter8;\nend test5;"
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» 返回文件|string|false|none|文件内容|示例：-|

## GET 获取mo文件信息接口

GET /file/package/list

```text
# 获取模型的源码数据
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|

> 返回示例

> 成功

```json
{
  "data": [
    {
      "create_time": "2022-09-02T15:27:10+08:00",
      "id": 0,
      "package_id": "41b464e1-d5a1-4e70-a81c-508e04950a54",
      "package_name": "test5",
      "space_name": "test1",
      "update_time": "2022-09-02T15:27:10+08:00"
    },
    {
      "create_time": "2022-09-02T15:21:29+08:00",
      "id": 1,
      "package_id": "4500ea39-bd84-4a98-ae61-f24070ee363d",
      "package_name": "Filter1",
      "space_name": "test1",
      "update_time": "2022-09-02T15:21:29+08:00"
    },
    {
      "create_time": "2022-09-01T19:46:41+08:00",
      "id": 2,
      "package_id": "c20fb619-9e95-4bf1-a4c1-5db2f3d729a8",
      "package_name": "test3",
      "space_name": "test1",
      "update_time": "2022-09-01T19:46:41+08:00"
    },
    {
      "create_time": "2022-09-01T19:34:41+08:00",
      "id": 3,
      "package_id": "dd7f135b-441c-42e1-ad7b-a6cd1407e05f",
      "package_name": "test1",
      "space_name": "test1",
      "update_time": "2022-09-01T19:34:41+08:00"
    },
    {
      "create_time": "2022-09-01T18:31:21+08:00",
      "id": 4,
      "package_id": "711f4b23-105e-4543-bb66-4fe570ff8af8",
      "package_name": "Applications",
      "space_name": "test1",
      "update_time": "2022-09-01T18:31:21+08:00"
    }
  ],
  "msg": "",
  "status": 0,
  "err": ""
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|[object]|true|none|返回数据对象数组|示例：-|
|»» create_time|string|true|none|上传时间|示例：2022-09-02T15:27:10+08:00|
|»» id|integer|true|none|序号|示例：-|
|»» package_id|string|true|none|模型包id|示例：41b464e1-d5a1-4e70-a81c-508e04950a54|
|»» package_name|string|true|none|模型包名称|示例：test5|
|»» space_name|string|true|none|所属用户空间|示例：test1|
|»» update_time|string|true|none|更新时间|示例：2022-09-02T15:27:10+08:00|
|» msg|string|true|none|消息字段|示例：-|
|» status|integer|true|none|状态码字段|示例：-|
|» err|string|true|none|错误提示字段|示例：-|

## POST 仿真结果文件下载

POST /file/result/all/get

```text
# 获取模型的源码数据
```

> Body 请求参数

```json
{
  "record_id": "6e4e7cb3-bf82-422d-9006-6b699b82263b"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|username|header|string| 否 ||none|
|Authorization|header|string| 否 ||none|
|space_id|header|string| 否 ||none|
|body|body|object| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_Tag">Tag</h2>

<a id="schematag"></a>
<a id="schema_Tag"></a>
<a id="tocStag"></a>
<a id="tocstag"></a>

```json
{
  "id": 1,
  "name": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|false|none||标签ID编号|
|name|string|false|none||标签名称|

<h2 id="tocS_Category">Category</h2>

<a id="schemacategory"></a>
<a id="schema_Category"></a>
<a id="tocScategory"></a>
<a id="tocscategory"></a>

```json
{
  "id": 1,
  "name": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|false|none||分组ID编号|
|name|string|false|none||分组名称|

<h2 id="tocS_Pet">Pet</h2>

<a id="schemapet"></a>
<a id="schema_Pet"></a>
<a id="tocSpet"></a>
<a id="tocspet"></a>

```json
{
  "id": 1,
  "category": {
    "id": 1,
    "name": "string"
  },
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "tags": [
    {
      "id": 1,
      "name": "string"
    }
  ],
  "status": "available"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|true|none||宠物ID编号|
|category|[Category](#schemacategory)|true|none||分组|
|name|string|true|none||名称|
|photoUrls|[string]|true|none||照片URL|
|tags|[[Tag](#schematag)]|true|none||标签|
|status|string|true|none||宠物销售状态|

#### 枚举值

|属性|值|
|---|---|
|status|available|
|status|pending|
|status|sold|
