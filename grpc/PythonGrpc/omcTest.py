from db_config.omc import omc

print("1",omc.sendExpression("clearProgram()"))
print("2",omc.sendExpression("setModelicaPath(\"/usr/lib/omlibrary\")"))
print("Buildings:",omc.sendExpression("loadModel(Buildings, {\"9.1.0\"},true,\"\",false)"))
print("Modelica:",omc.sendExpression("loadModel(Modelica, {\"4.0.0\"},true,\"\",false)"))
print("SolarPower:",omc.sendExpression("loadModel(SolarPower, {\"\"},true,\"\",false)"))
print("WindPowerSystem:",omc.sendExpression("loadModel(WindPowerSystem, {\"\"},true,\"\",false)"))
print("5",omc.loadFile("/yssim-go/public/UserFiles/UploadFile/xuqingda/ChillerStage/20230216111124/ChillerStage.mo"))
print("6",omc.buildModelFmu(className="Modelica.Blocks.Examples.PID_Controller",fileNamePrefix="ssss"))
print("7",omc.buildModelFmu(className="ChillerStage",fileNamePrefix="xxxx"))

