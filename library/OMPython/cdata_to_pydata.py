# -- coding: utf-8 --


def CdataToPYdata (data):
    b = data.replace(", ", ",")
    b = b.replace("\\\"", "")
    b = b.replace('\"true\"', 'True')
    b = b.replace('\"false\"', 'False')
    b = b.replace('true\n', 'True')
    b = b.replace('false\n', 'False')
    b = b.replace("\n", "")
    b = b.replace("\r", "\\r")

    c = ""
    n = 0
    if b == "\"\"":
        b = b.replace("\"", "")
        return b

    try:
        for i in range(len(b)):
            if i != n:
                continue
            if b[i] == '"':
                end_n = b[i + 1:].find('"')
                d = b[i:i + end_n + 2]
                if b[i - 1] == "=":
                    d = b[i + 1:i + end_n + 2]
                c += d
                n += (end_n + 2)
                continue
            if b[i] == "," and (b[i + 1] in ['(', ')']):
                c += '",'
            elif b[i] == "," and ((b[i + 1] in [')', '}']) and (b[i - 1] in [')', '}'])):
                pass
            elif b[i] == "," and (b[i + 1] == '{' and b[i - 1] == '}'):
                c += ','
            elif b[i] == "," and (b[i - 1] in ['(', ')']):
                c += ',"'
            elif b[i] == "," and (b[i + 1] in ['{', '}']):
                c += '",'
            elif b[i] == "," and (b[i - 1] in ['{', '}']):
                c += ',"'
            elif b[i] == "," and (b[i + 1] not in ['(', ')', '{', '}'] or b[i - 1] not in ['(', ')', '{', '}']):
                c += '","'
            elif b[i] == "{":
                if b[i + 1] == "{" or b[i + 1] == "(":
                    c += '['
                else:
                    c += '["'
            elif b[i] == "}":
                if i == len(b) - 1 and b[i - 1] in ["}", ")", ","]:
                    c += ']'
                elif b[i - 1] in ["}", ")", ","]:
                    c += ']'
                else:
                    c += '"]'
            elif b[i] == "(":
                if i == 0:
                    c += '["'
                else:
                    c += '",["'
            elif b[i] == ")":
                if i == len(b) -1:
                    c += '"]'
                elif b[i + 1] in [")", "}"] and b[i - 1] == ",":
                    c += ']'
                elif b[i + 1] in [',', '}']:
                    c += '"]'
                else:
                    c += b[i]
            else:
                c += b[i]
            n += 1
            # py_data = ['']
        c = c.replace('[\"\"]', '[\"\"\"\"]')
        c = c.replace('\"\"', '"')
        c = c.replace(',\",', '"",')
        c = c.replace(',\"]', ',\"\"]')
        c = c.replace('[\",', '[\"\",')
    except Exception as e:
        if c != "Error":
            print(e)
    if c.startswith("["):
        py_data = eval(c)
    else:
        py_data = c.replace("\"", "")
        if py_data in ['True', 'true', 'TRUE', 'False', 'false', 'FALSE']:
            py_data = eval(py_data)
    return py_data


if __name__ == '__main__':
    print(CdataToPYdata('{Modelica.Icons.Record}'))

