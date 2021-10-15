failed_modules={}
loaded_modules=[]
try:
  from casadi_core import *
  loaded_modules.append('casadi_core')
except Exception as e:
  failed_modules['casadi_core'] = str(e)
