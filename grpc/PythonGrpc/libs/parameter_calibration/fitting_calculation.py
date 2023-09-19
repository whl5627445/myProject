from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
import pandas as pd
from libs.function.grpc_log import log


def prediction (data):
    # 读取数据
    # data = pd.read_excel(data_file)

    values = data.values
    row_num = len(data.values)  # 数据行数
    col_num = len(data.values[0])  # 数据列数
    X_all = []  # 保存特征
    y_all = []  # 保存第一列作为预测标签
    for i in range(1, row_num):
        X = []
        y = []
        for j in range(col_num):
            if j == 0:  # 第0列作为标签
                y.append(values[i][j])
            else:  # 其他列是输入特征
                X.append(values[i][j])
        X_all.append(X)
        y_all.append(y)

    # 取90%的数据作为训练数据，10%用于判断模型效果
    X_train, X_test, y_train, y_test = train_test_split(X_all, y_all, test_size=0.2, random_state=42)

    # 简单线性回归

    # 训练模型
    model = LinearRegression(fit_intercept=False)  # 取消截距
    model.fit(X_train, y_train)

    # 预测
    # predictions = []

    # for i, item in enumerate(X_test):
    #     pred = model.predict([item])[0]
    #     coef = model.coef_[0]
    #     pred = sum([coef[i] * item[i] for i in range(len(coef))])
    #     predictions.append(pred)
    # print(model.score(X_all, y_all))
    return model.coef_[0], model.score(X_all, y_all)


#
# def get_columns_value_dict (data):
#     d = {}
#     for i in data.columns:
#         d[i] = data[i].tolist()
#     return d


def get_formula_operation (value_dict, formula_list):
    data_dict = {}
    data = pd.DataFrame(value_dict)
    approach = data.columns[0]
    approach_data = data[approach].tolist()
    try:
        data_dict[approach] = [float(f) for f in approach_data]
    except Exception as e:
        log.info("approach数据解析失败， 含有非浮点型字符")
        return None, None
    l = len(approach_data)
    for formula_dict in formula_list:
        # formula = "Twb * Tr^2 * LGRatio^2"
        formula = formula_dict["formula"].replace(" ", "")
        f_str_list = formula.split("*")
        if formula == "1":
            data_dict[formula] = [1] * l
            continue
        f_dict = {}
        for f_str in f_str_list:
            index = f_str.find("^")
            name = f_str if f_str.find("^") == -1 else f_str[:index]
            if value_dict[name] is not None:
                f_dict[name] = value_dict[name]
            else:
                return None, None
        formula_data_list = []
        f_list = f_dict.keys()

        for i in range(0, l):
            replace_formula = formula.replace("^", "**")
            for f in f_list:
                d = str(f_dict[f][i])
                if d == "":
                    return None
                replace_formula = replace_formula.replace(f, d)
            value = eval(replace_formula)
            formula_data_list.append(value)

        data_dict[formula] = formula_data_list
    return pd.DataFrame(data_dict)


def get_coefficient_score(actual_data, formula_list):
    data = get_formula_operation(actual_data, formula_list)
    if data is None:
        return None, None, "实测数据含有空值，本次拟合失败"
    predictions_coefficient, predictions_score = prediction(data)
    return [str(c) for c in predictions_coefficient], str(predictions_score), None



# predictions, score = prediction((r"test.xlsx"))
# print(predictions)
# print(len(predictions))
# a = {
#     "formula": [
#         {"formula": "1", "coefficient": "Coeff(1)"},
#         {"formula": "Twb", "coefficient": "Coeff(2)"},
#         {"formula": "Twb^2", "coefficient": "Coeff(3)"},
#         {"formula": "Tr", "coefficient": "Coeff(4)"},
#         {"formula": "Twb * Tr", "coefficient": "Coeff(5)"},
#         {"formula": "Twb^2 * Tr", "coefficient": "Coeff(6)"},
#         {"formula": "Tr^2", "coefficient": "Coeff(7)"},
#         {"formula": "Twb * Tr^2", "coefficient": "Coeff(8)"},
#         {"formula": "Twb^2 * Tr^2", "coefficient": "Coeff(9)"},
#         {"formula": "LGRatio", "coefficient": "Coeff(10)"},
#         {"formula": "Twb * LGRatio", "coefficient": "Coeff(11)"},
#         {"formula": "Twb^2 * LGRatio", "coefficient": "Coeff(12)"},
#         {"formula": "Tr * LGRatio", "coefficient": "Coeff(13)"},
#         {"formula": "Twb * Tr * LGRatio", "coefficient": "Coeff(14)"},
#         {"formula": "Twb^2 * Tr * LGRatio", "coefficient": "Coeff(15)"},
#         {"formula": "Tr^2 * LGRatio", "coefficient": "Coeff(16)"},
#         {"formula": "Twb * Tr^2 * LGRatio", "coefficient": "Coeff(17)"},
#         {"formula": "Twb^2 * Tr^2 * LGRatio", "coefficient": "Coeff(18)"},
#         {"formula": "LGRatio^2", "coefficient": "Coeff(19)"},
#         {"formula": "Twb * LGRatio^2", "coefficient": "Coeff(20)"},
#         {"formula": "Twb^2 * LGRatio^2", "coefficient": "Coeff(21)"},
#         {"formula": "Tr * LGRatio^2", "coefficient": "Coeff(22)"},
#         {"formula": "Twb * Tr * LGRatio^2", "coefficient": "Coeff(23)"},
#         {"formula": "Twb^2 * Tr * LGRatio^2", "coefficient": "Coeff(24)"},
#         {"formula": "Tr^2 * LGRatio^2", "coefficient": "Coeff(25)"},
#         {"formula": "Twb * Tr^2 * LGRatio^2", "coefficient": "Coeff(26)"},
#         {"formula": "Twb^2 * Tr^2 * LGRatio^2", "coefficient": "Coeff(27)"}],
#     "variable": ["Tr", "LGRatio", "Twb"]
# }


if __name__ == '__main__':
    data = pd.read_excel("test.xlsx")
    c = data.columns
    for i in c:
        print(data[i].tolist())
    # value_dict = get_columns_value_dict(data)
    #
    # x = get_formula_operation(data, value_dict, a["formula"])
    # x.to_excel("test1.xlsx", index=False)
    # predictions, score = prediction(x)
    # predictions1, score1 = prediction(pd.read_excel("Cooling.xlsx"))
    # print(predictions.tolist())
    # print(predictions1.tolist())
    # # print(len(predictions))
