# -- coding: utf-8 --
import time


def CdataToPYdata (data):
    b = data.replace(", ", ",")

    b = b.replace("\\\"", "")
    b = b.replace('\"true\"', 'True')
    b = b.replace('\"false\"', 'False')
    b = b.replace('true\n', 'True')
    b = b.replace('false\n', 'False')
    b = b.replace("\n", "\\n")
    b = b.replace("\r", "\\r")
    b = b.removesuffix("\\n")
    c = ""
    n = 0
    if b == "\"\"":
        b = b.replace("\"", "")
        return b
    l_b = len(b)
    try:
        for i in range(l_b):
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
                if i == l_b - 1 and b[i - 1] in ["}", ")", ","]:
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
                if i == l_b -1:
                    c += '"]'
                elif b[i + 1] in [")", "}"] and b[i - 1] == ",":
                    c += ']'
                elif b[i + 1] in [',', '}']:
                    c += '"]'
                elif b[i + 1] == " " and l_b != i + 2 and b[i + 2] == "-":
                    c += '"],"'
                else:
                    c += b[i]
            else:
                c += b[i]
            n += 1
            # py_data = ['']
        c = c.replace('[\"\"]', '[\"\"\"\"]')
        c = c.replace('\"\"', '"')
        c = c.replace(',\",', '"",')
        # c = c.replace(',\"]', ',\"\"]')
        # c = c.replace('[\",', '[\"\",')
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
    s = time.time()
    # a = CdataToPYdata('{{},{Evaluate=true, HideResult=true, choices(true,false)},{Dialog("General","Parameters",false,false,false,-,-,-,-,"",false), Evaluate=true, HideResult=true, choices(true,false)},{Evaluate=true, Dialog("General","Initialization",true,false,false,-,-,-,-,"",false)},{Dialog("General","Initialization",true,false,false,-,-,-,-,"",false)},{Placement(true,60.0,-120.0,-20.0,-20.0,20.0,20.0,90.0,-,-,-,-,-,-,)},{Placement(true,60.0,120.0,-20.0,-20.0,20.0,20.0,270.0,-,-,-,-,-,-,)},{HideResult=true},{HideResult=true}}')
    a = CdataToPYdata('("block","Wrap angle to interval ]-pi,pi] or [0,2*pi[",false,false,false,"D:/OpenModelica/lib/omlibrary/Modelica 3.2.3/Blocks/Math.mo",false,2476,3,2509,16,{},false,false,"","",false,"")')
    print(a)
    # for i in range(1000):
        # CdataToPYdata('{Modelica.Icons.Record}')
    # print(time.time() - s)
