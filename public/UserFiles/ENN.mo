﻿package ENN "Version 2.0"
  package Examples
    model Scenario1_Status
      extends Modelica.Icons.Example;
      parameter Media.Medium medium = ENN.Media.Water() "Cooling medium" annotation(
        choicesAllMatching = true);
      parameter Modelica.SIunits.Temperature Tamb(displayUnit = "degC") = 293.15 "Ambient temperature";
      Machines.Boiler.WasteSteamBoiler wasteSteamBoiler(medium = medium, Tamb = Tamb) annotation(
        Placement(transformation(extent = {{10, -20}, {50, 20}})));
      Sinks.Sink steamSink annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {30, -70})));
      Pipes.SteamPipe steamPipe annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = -90, origin = {30, -42})));
      Sources.Source source(constantAmbientTemperature = Tamb, medium = medium, constantAmbientPressure(displayUnit = "MPa") = 100000) annotation(
        Placement(transformation(extent = {{-60, 60}, {-80, 80}})));
      Machines.Pump.Pump pump_VF(medium = medium, idealPump(V_flow0 = 0.0075)) "变频水泵" annotation(
        Placement(transformation(extent = {{-40, 60}, {-20, 80}})));
      Pipes.Pipe coldWaterPipe(medium = medium, m = 0.1, T0 = Tamb, h_g = 0) annotation(
        Placement(transformation(extent = {{0, 60}, {20, 80}})));
      Sensors.MassFlowSensor massFlowSensor(medium = medium) annotation(
        Placement(transformation(extent = {{10, -10}, {-10, 10}}, rotation = 90, origin = {30, 40})));
      Modelica.Blocks.Continuous.LimPID PID(controllerType = Modelica.Blocks.Types.SimpleController.PI, k = 100, Ti = 0.1, yMax = 50, yMin = 0) annotation(
        Placement(transformation(extent = {{60, 80}, {40, 100}})));
      Sinks.Chimney chimney annotation(
        Placement(transformation(extent = {{70, -10}, {90, 10}})));
      Furnace.Components.GasGlassFurnance gasGlassFurnance(ele_factor = carbonFactor.factor) annotation(
        Placement(transformation(extent = {{-80, -20}, {-40, 20}})));
      Valves.MatValve Mat annotation(
        Placement(transformation(extent = {{-120, 4}, {-100, 24}})));
      Valves.GasValve Gas annotation(
        Placement(transformation(extent = {{-120, -20}, {-100, 0}})));
      Valves.AirValve Air annotation(
        Placement(transformation(extent = {{-120, -44}, {-100, -24}})));
      Sinks.ProductSink productSink annotation(
        Placement(transformation(extent = {{0, 10}, {-20, 30}})));
      Sinks.ThermalSink thermalSink annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {-60, -50})));
      Pipes.FluePipe fluePipe annotation(
        Placement(transformation(extent = {{-20, -10}, {0, 10}})));
      Blocks.Production.proGlass proGlass annotation(
        Placement(transformation(extent = {{-200, 4}, {-180, 24}})));
      Modelica.Blocks.Continuous.LimPID PID1(controllerType = Modelica.Blocks.Types.SimpleController.PI, k = 100, Ti = 0.1, yMax = 3, yMin = 0) annotation(
        Placement(transformation(extent = {{-160, 24}, {-140, 4}})));
      Sinks.Sink airSink annotation(
        Placement(transformation(extent = {{0, -30}, {-20, -10}})));
      Machines.Compressor.EletricalCompressor eletricalCompressor annotation(
        Placement(transformation(extent = {{-180, -80}, {-160, -60}})));
      Sources.AirSource1 airSource1_1 annotation(
        Placement(transformation(extent = {{-220, -80}, {-200, -60}})));
      Machines.Compressor.EletricalCompressor eletricalCompressor1 annotation(
        Placement(transformation(extent = {{-180, -100}, {-160, -80}})));
      Sources.AirSource1 airSource1_2 annotation(
        Placement(transformation(extent = {{-220, -100}, {-200, -80}})));
      Blocks.Controller.CompressorOperation compressorOperation(A = {1, 1, 1, 1}, N = {1, 1, 1, 1}, n = 4) annotation(
        Placement(transformation(extent = {{-260, -88}, {-240, -68}})));
      Pipes.AirPipe airPipe annotation(
        Placement(transformation(extent = {{-140, -80}, {-120, -60}})));
      Sources.AirSource1 airSource1_3 annotation(
        Placement(transformation(extent = {{-220, -120}, {-200, -100}})));
      Machines.Compressor.EletricalCompressor eletricalCompressor2 annotation(
        Placement(transformation(extent = {{-180, -120}, {-160, -100}})));
      Sources.AirSource1 airSource1_4 annotation(
        Placement(transformation(extent = {{-220, -140}, {-200, -120}})));
      Machines.Compressor.EletricalCompressor eletricalCompressor3 annotation(
        Placement(transformation(extent = {{-180, -140}, {-160, -120}})));
      Modelica.Blocks.Tables.CombiTable1D m_comAir(table = [0.1, 0.21369167; 0.6, 1.165590926; 1.2, 2.119256229; 1.8, 2.889894857; 2.4, 3.502902857; 3, 3.980571429; 3.6, 4.776685714]) annotation(
        Placement(transformation(extent = {{-292, -88}, {-272, -68}})));
      RenewableEnergy.Components.Internal.ACbus aCbusBar annotation(
        Placement(transformation(extent = {{-280, 40}, {-260, 60}})));
      Electrical.ElectricGrid electricGrid(V_ref(displayUnit = "kV")) annotation(
        Placement(transformation(extent = {{-340, 40}, {-320, 60}})));
      Electrical.Transformer transformer(V_ref = 380) annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {-250, 50})));
      inner Environment environment annotation(
        Placement(transformation(extent = {{-340, 80}, {-320, 100}})));
      Modelica.Electrical.Analog.Ideal.IdealOpeningSwitch switch annotation(
        Placement(transformation(extent = {{-306, 60}, {-286, 40}})));
      Electrical.GridControl gridControl(P_renew = 0, P_load = transformer.power_req) annotation(
        Placement(transformation(extent = {{-330, 10}, {-310, 30}})));
      RenewableEnergy.Components.Internal.ACbus aCbusBar1 annotation(
        Placement(transformation(extent = {{-240, 40}, {-220, 60}})));
      CarbonCalc.CarbonFactor carbonFactor(zones = ENN.Utilities.Types.Zones.North, P_grid = electricGrid.P_grid) annotation(
        Placement(transformation(extent = {{-340, -140}, {-320, -120}})));
    equation
      connect(wasteSteamBoiler.steam_out, steamPipe.steam_in) annotation(
        Line(points = {{30, -20}, {30, -32}}, color = {255, 0, 0}));
      connect(steamPipe.steam_out, steamSink.generalFlowPort) annotation(
        Line(points = {{30, -52}, {30, -60}}, color = {255, 0, 0}));
      connect(pump_VF.port_a, source.flowPort) annotation(
        Line(points = {{-40, 70}, {-60, 70}}, color = {85, 255, 85}));
      connect(wasteSteamBoiler.m_flow, PID.u_s) annotation(
        Line(points = {{52, 12}, {96, 12}, {96, 90}, {62, 90}}, color = {0, 0, 127}));
      connect(PID.y, pump_VF.f) annotation(
        Line(points = {{39, 90}, {-30, 90}, {-30, 82}}, color = {0, 0, 127}));
      connect(pump_VF.port_b, coldWaterPipe.port_a) annotation(
        Line(points = {{-20, 70}, {0, 70}}, color = {85, 255, 85}));
      connect(wasteSteamBoiler.flue_out, chimney.flue_in) annotation(
        Line(points = {{50, 0}, {70, 0}}, color = {95, 95, 95}, thickness = 0.5));
      connect(massFlowSensor.y, PID.u_m) annotation(
        Line(points = {{41, 40}, {50, 40}, {50, 78}}, color = {0, 0, 127}));
      connect(Mat.mat_out, gasGlassFurnance.mat_in) annotation(
        Line(points = {{-100, 14}, {-80, 14}}, color = {255, 170, 255}));
      connect(Gas.gas_out, gasGlassFurnance.gas_in) annotation(
        Line(points = {{-100, -10}, {-90, -10}, {-90, 0}, {-80, 0}}, color = {255, 170, 85}));
      connect(gasGlassFurnance.pro_out, productSink.port_a) annotation(
        Line(points = {{-40, 14}, {-40, 20}, {-20, 20}}, color = {255, 170, 255}));
      connect(thermalSink.heatPort, gasGlassFurnance.heatPort) annotation(
        Line(points = {{-60, -40}, {-60, -20}, {-60.4, -20}}, color = {191, 0, 0}));
      connect(gasGlassFurnance.flue_out, fluePipe.flue_in) annotation(
        Line(points = {{-40, 0}, {-20, 0}}, color = {95, 95, 95}, thickness = 0.5));
      connect(fluePipe.flue_out, wasteSteamBoiler.flue_in) annotation(
        Line(points = {{0, 0}, {10, 0}}, color = {95, 95, 95}, thickness = 0.5));
      connect(proGlass.proGlass, PID1.u_s) annotation(
        Line(points = {{-179, 14}, {-162, 14}}, color = {0, 0, 127}));
      connect(gasGlassFurnance.m_pro, PID1.u_m) annotation(
        Line(points = {{-48, 22}, {-48, 34}, {-150, 34}, {-150, 26}}, color = {0, 0, 127}));
      connect(PID1.y, Mat.load) annotation(
        Line(points = {{-139, 14}, {-122, 14}}, color = {0, 0, 127}));
      connect(gasGlassFurnance.com_air_out, airSink.generalFlowPort) annotation(
        Line(points = {{-40.4, -18}, {-40.4, -20}, {-20, -20}}, color = {0, 127, 255}));
      connect(airSource1_1.port_a, eletricalCompressor.air_in) annotation(
        Line(points = {{-200, -70}, {-180, -70}}, color = {0, 127, 255}));
      connect(airSource1_2.port_a, eletricalCompressor1.air_in) annotation(
        Line(points = {{-200, -90}, {-180, -90}}, color = {0, 127, 255}));
      connect(compressorOperation.B[1], airSource1_1.m_in) annotation(
        Line(points = {{-239, -72.75}, {-234, -72.75}, {-234, -70}, {-222, -70}}, color = {0, 0, 127}));
      connect(compressorOperation.B[2], airSource1_2.m_in) annotation(
        Line(points = {{-239, -72.25}, {-234, -72.25}, {-234, -90}, {-222, -90}}, color = {0, 0, 127}));
      connect(eletricalCompressor.air_out, airPipe.air_in) annotation(
        Line(points = {{-160, -70}, {-140, -70}}, color = {0, 127, 255}));
      connect(eletricalCompressor1.air_out, airPipe.air_in) annotation(
        Line(points = {{-160, -90}, {-148, -90}, {-148, -70}, {-140, -70}}, color = {0, 127, 255}));
      connect(eletricalCompressor2.air_out, airPipe.air_in) annotation(
        Line(points = {{-160, -110}, {-148, -110}, {-148, -70}, {-140, -70}}, color = {0, 127, 255}));
      connect(eletricalCompressor3.air_out, airPipe.air_in) annotation(
        Line(points = {{-160, -130}, {-148, -130}, {-148, -70}, {-140, -70}}, color = {0, 127, 255}));
      connect(airSource1_3.port_a, eletricalCompressor2.air_in) annotation(
        Line(points = {{-200, -110}, {-180, -110}}, color = {0, 127, 255}));
      connect(airSource1_4.port_a, eletricalCompressor3.air_in) annotation(
        Line(points = {{-200, -130}, {-180, -130}}, color = {0, 127, 255}));
      connect(compressorOperation.B[3], airSource1_3.m_in) annotation(
        Line(points = {{-239, -71.75}, {-234, -71.75}, {-234, -110}, {-222, -110}}, color = {0, 0, 127}));
      connect(compressorOperation.B[4], airSource1_4.m_in) annotation(
        Line(points = {{-239, -71.25}, {-234, -71.25}, {-234, -130}, {-222, -130}}, color = {0, 0, 127}));
      connect(compressorOperation.Q, m_comAir.y[1]) annotation(
        Line(points = {{-262, -78}, {-271, -78}}, color = {0, 0, 127}));
      connect(airPipe.air_out, gasGlassFurnance.com_air_in) annotation(
        Line(points = {{-120, -70}, {-80, -70}, {-80, -18}}, color = {0, 127, 255}));
      connect(coldWaterPipe.port_b, massFlowSensor.port_a) annotation(
        Line(points = {{20, 70}, {30, 70}, {30, 50}}, color = {85, 255, 85}));
      connect(massFlowSensor.port_b, wasteSteamBoiler.water_in) annotation(
        Line(points = {{30, 30}, {30, 30}, {30, 20}}, color = {85, 255, 85}));
      connect(electricGrid.p, switch.p) annotation(
        Line(points = {{-320, 50}, {-306, 50}}, color = {0, 140, 72}));
      connect(gridControl.y, switch.control) annotation(
        Line(points = {{-309, 20}, {-296, 20}, {-296, 38}}, color = {255, 0, 255}));
      connect(switch.n, aCbusBar.term) annotation(
        Line(points = {{-286, 50}, {-270, 50}}, color = {0, 0, 255}));
      connect(aCbusBar.term, transformer.p) annotation(
        Line(points = {{-270, 50}, {-259.8, 50}}, color = {0, 140, 72}));
      connect(transformer.n, aCbusBar1.term) annotation(
        Line(points = {{-240, 50}, {-230, 50}}, color = {0, 140, 72}));
      connect(gasGlassFurnance.pin, aCbusBar1.term) annotation(
        Line(points = {{-60, 20}, {-60, 50}, {-230, 50}}, color = {0, 140, 72}));
      connect(Mat.m_mat, Gas.dmMat) annotation(
        Line(points = {{-99, 19}, {-92, 19}, {-92, 2}, {-128, 2}, {-128, -10}, {-122, -10}}, color = {0, 0, 127}));
      connect(Air.dmMat, Gas.dmMat) annotation(
        Line(points = {{-122, -34}, {-128, -34}, {-128, -10}, {-122, -10}}, color = {0, 0, 127}));
      connect(eletricalCompressor.pin, aCbusBar1.term) annotation(
        Line(points = {{-170, -80}, {-154, -80}, {-154, -20}, {-220, -20}, {-220, 50}, {-230, 50}}, color = {0, 140, 72}));
      connect(eletricalCompressor1.pin, aCbusBar1.term) annotation(
        Line(points = {{-170, -100}, {-154, -100}, {-154, -20}, {-220, -20}, {-220, 50}, {-230, 50}}, color = {0, 140, 72}));
      connect(eletricalCompressor2.pin, aCbusBar1.term) annotation(
        Line(points = {{-170, -120}, {-154, -120}, {-154, -20}, {-220, -20}, {-220, 50}, {-230, 50}}, color = {0, 140, 72}));
      connect(eletricalCompressor3.pin, aCbusBar1.term) annotation(
        Line(points = {{-170, -140}, {-154, -140}, {-154, -20}, {-220, -20}, {-220, 50}, {-230, 50}}, color = {0, 140, 72}));
      connect(pump_VF.pin, aCbusBar1.term) annotation(
        Line(points = {{-30, 60}, {-30, 50}, {-230, 50}}, color = {0, 140, 72}));
      connect(Air.air_out, gasGlassFurnance.air_in) annotation(
        Line(points = {{-100, -34}, {-86, -34}, {-86, -9.6}, {-80, -9.6}}, color = {0, 127, 255}));
      connect(gasGlassFurnance.m_pro, m_comAir.u[1]) annotation(
        Line(points = {{-48, 22}, {-48, 34}, {-294, 34}, {-294, -78}}, color = {0, 0, 127}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-340, -140}, {120, 100}})),
        Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-340, -140}, {120, 100}})),
        experiment(StopTime = 3.1536e+07, StartTime = 0, Tolerance = 1e-06, Interval = 63072),
  __OpenModelica_simulationFlags( ls = "lapack", lv = "LOG_STATS", nls = "hybrid", s = "dassl"));
    end Scenario1_Status;
  end Examples;

  package Blocks
    package Production "产量模型"
      model proGlass "燃气玻璃窑炉产能计算，输出燃气玻璃窑炉逐时产能（kg/s）"
        replaceable Base.proBase pro_GasFurnace "燃气玻璃窑炉逐时产能（kg/s）" annotation(
          Placement(transformation(extent = {{0, -20}, {40, 20}})));
        Modelica.Blocks.Interfaces.RealOutput proGlass(unit = "kg/s") "燃气玻璃窑炉逐时产能" annotation(
          Placement(transformation(extent = {{100, -10}, {120, 10}})));
        Modelica.Blocks.Sources.CombiTimeTable day(table = date.table_day) "公历日" annotation(
          Placement(transformation(extent = {{-80, 20}, {-60, 40}})));
        Modelica.Blocks.Sources.CombiTimeTable month(table = date.table_month) "公历月" annotation(
          Placement(transformation(extent = {{-80, -10}, {-60, 10}})));
        Modelica.Blocks.Sources.CombiTimeTable year(table = date.table_year) "公历年" annotation(
          Placement(transformation(extent = {{-80, -40}, {-60, -20}})));
        replaceable Date.Date2021 date constrainedby Date.DateBase annotation(
           Placement(transformation(extent = {{-80, -80}, {-60, -60}})),
           __Dymola_choicesAllMatching = true);
        Base.RealToInteger realToInteger annotation(
          Placement(transformation(extent = {{-40, 20}, {-20, 40}})));
        Base.RealToInteger realToInteger1 annotation(
          Placement(transformation(extent = {{-40, -10}, {-20, 10}})));
        Base.RealToInteger realToInteger2 annotation(
          Placement(transformation(extent = {{-40, -40}, {-20, -20}})));
      equation
        connect(pro_GasFurnace.proGlass, proGlass) annotation(
          Line(points = {{42, 0}, {110, 0}}, color = {0, 0, 127}));
        connect(day.y[1], realToInteger.u) annotation(
          Line(points = {{-59, 30}, {-42, 30}}, color = {0, 0, 127}));
        connect(realToInteger.y, pro_GasFurnace.day) annotation(
          Line(points = {{-19, 30}, {-12, 30}, {-12, 12}, {-4, 12}}, color = {255, 127, 0}));
        connect(month.y[1], realToInteger1.u) annotation(
          Line(points = {{-59, 0}, {-42, 0}}, color = {0, 0, 127}));
        connect(realToInteger1.y, pro_GasFurnace.month) annotation(
          Line(points = {{-19, 0}, {-4, 0}}, color = {255, 127, 0}));
        connect(year.y[1], realToInteger2.u) annotation(
          Line(points = {{-59, -30}, {-42, -30}}, color = {0, 0, 127}));
        connect(realToInteger2.y, pro_GasFurnace.year) annotation(
          Line(points = {{-19, -30}, {-10, -30}, {-10, -12}, {-4, -12}}, color = {255, 127, 0}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 20), Line(origin = {0.061, 4.184}, points = {{81.939, 36.056}, {65.362, 36.056}, {14.39, -26.199}, {-29.966, 113.485}, {-65.374, -61.217}, {-78.061, -78.184}}, color = {95, 95, 95}, smooth = Smooth.Bezier)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end proGlass;

      package Base
        model proBase "燃气玻璃窑炉产能计算，输出燃气玻璃窑炉逐时产能（kg/s）"
          parameter Integer dayDaily_0 = 1 "平季开始日" annotation(
            Dialog(group = "平季起止时间"));
          parameter Integer monthDaily_0 = 1 "平季开始月" annotation(
            Dialog(group = "平季起止时间"));
          parameter Integer dayDaily_1 = 31 "平季结束日" annotation(
            Dialog(group = "平季起止时间"));
          parameter Integer monthDaily_1 = 3 "平季结束月" annotation(
            Dialog(group = "平季起止时间"));
          parameter Integer dayLow_0 = 1 "淡季开始日" annotation(
            Dialog(group = "淡季起止时间"));
          parameter Integer monthLow_0 = 4 "淡季开始月" annotation(
            Dialog(group = "淡季起止时间"));
          parameter Integer dayLow_1 = 31 "淡季结束日" annotation(
            Dialog(group = "淡季起止时间"));
          parameter Integer monthLow_1 = 7 "淡季结束月" annotation(
            Dialog(group = "淡季起止时间"));
          parameter Integer dayHigh_0 = 1 "旺季开始日" annotation(
            Dialog(group = "旺季起止时间"));
          parameter Integer monthHigh_0 = 8 "旺季开始月" annotation(
            Dialog(group = "旺季起止时间"));
          parameter Integer dayHigh_1 = 31 "旺季结束日" annotation(
            Dialog(group = "旺季起止时间"));
          parameter Integer monthHigh_1 = 12 "旺季结束月" annotation(
            Dialog(group = "旺季起止时间"));
          parameter Real furnace_nominal = 250 "燃气玻璃窑炉日产能，单位：t/d";
          parameter Real furnace_PJ = furnace_nominal * 0.85 "平季日产量，单位：t/d";
          parameter Real furnace_DJ = furnace_nominal * 0.65 "淡季日产量，单位：t/d";
          parameter Real furnace_WJ = furnace_nominal "旺季日产量，单位：t/d";
          Modelica.Blocks.Interfaces.IntegerInput day "日" annotation(
            Placement(transformation(extent = {{-140, 40}, {-100, 80}})));
          Modelica.Blocks.Interfaces.IntegerInput month "月" annotation(
            Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
          Modelica.Blocks.Interfaces.IntegerInput year "年" annotation(
            Placement(transformation(extent = {{-140, -80}, {-100, -40}})));
          Modelica.Blocks.Interfaces.RealOutput proGlass(unit = "t/h") "燃气玻璃窑炉逐时产能" annotation(
            Placement(transformation(extent = {{100, -10}, {120, 10}})));
        protected
          Boolean High;
          Boolean Low;
          Boolean Daily;
          Real proGlass_d "燃气玻璃窑炉日产能";
        equation
          High = if ENN.Utilities.Functions.dayOfTheYear(day, month, year) >= ENN.Utilities.Functions.dayOfTheYear(dayHigh_0, monthHigh_0, year) and ENN.Utilities.Functions.dayOfTheYear(day, month, year) <= ENN.Utilities.Functions.dayOfTheYear(dayHigh_1, monthHigh_1, year) then true else false;
          Low = if ENN.Utilities.Functions.dayOfTheYear(day, month, year) >= ENN.Utilities.Functions.dayOfTheYear(dayLow_0, monthLow_0, year) and ENN.Utilities.Functions.dayOfTheYear(day, month, year) <= ENN.Utilities.Functions.dayOfTheYear(dayLow_1, monthLow_1, year) then true else false;
          Daily = if ENN.Utilities.Functions.dayOfTheYear(day, month, year) >= ENN.Utilities.Functions.dayOfTheYear(dayDaily_0, monthDaily_0, year) and ENN.Utilities.Functions.dayOfTheYear(day, month, year) <= ENN.Utilities.Functions.dayOfTheYear(dayDaily_1, monthDaily_1, year) then true else false;
          proGlass_d = if High then furnace_WJ elseif Daily then furnace_PJ else furnace_DJ;
          proGlass = proGlass_d / 24 / 3600 * 1000;
          annotation(
            Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 20), Line(origin = {0.061, 4.184}, points = {{81.939, 36.056}, {65.362, 36.056}, {14.39, -26.199}, {-29.966, 113.485}, {-65.374, -61.217}, {-78.061, -78.184}}, color = {95, 95, 95}, smooth = Smooth.Bezier)}),
            Diagram(coordinateSystem(preserveAspectRatio = false)));
        end proBase;

        block RealToInteger "Convert Real to Integer signal,Returns the smallest integer not less than u"
          extends Modelica.Blocks.Icons.IntegerBlock;
        public
          Modelica.Blocks.Interfaces.RealInput u "Connector of Real input signal" annotation(
            Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
          Modelica.Blocks.Interfaces.IntegerOutput y "Connector of Integer output signal" annotation(
            Placement(transformation(extent = {{100, -10}, {120, 10}})));
        equation
          y = integer(ceil(u));
          annotation(
            Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100.0, -100.0}, {100.0, 100.0}}), graphics = {Text(lineColor = {0, 0, 127}, extent = {{-100.0, -40.0}, {0.0, 40.0}}, textString = "R"), Text(lineColor = {255, 127, 0}, extent = {{20.0, -40.0}, {120.0, 40.0}}, textString = "I"), Polygon(lineColor = {255, 127, 0}, fillColor = {255, 127, 0}, fillPattern = FillPattern.Solid, points = {{50.0, 0.0}, {30.0, 20.0}, {30.0, 10.0}, {0.0, 10.0}, {0.0, -10.0}, {30.0, -10.0}, {30.0, -20.0}, {50.0, 0.0}})}),
            Documentation(info = "<html>
<p>This block computes the output <b>y</b> as <i>nearest integer value</i> of the input <b>u</b>: </p>
<p><span style=\"font-family: Courier New;\">    y = ceil(u);</span></p>
</html>"));
        end RealToInteger;
      end Base;

      package Date "公历数表"
        record Date2021 "2021年公历表"
          extends DateBase(table_year = [0, 2021; 31536000, 2021], table_month = [0, 1; 2678400, 1; 2764800, 2; 5097600, 2; 5184000, 3; 7776000, 3; 7862400, 4; 10368000, 4; 10454400, 5; 12960000, 5; 13046400, 5; 13132800, 6; 15638400, 6; 15724800, 7; 18057600, 7; 18316800, 7; 18403200, 8; 20995200, 8; 21081600, 9; 23587200, 9; 23673600, 10; 26265600, 10; 26352000, 11; 28857600, 11; 28944000, 12; 31536000, 12], table_day = [86400, 1; 172800, 2; 259200, 3; 345600, 4; 432000, 5; 518400, 6; 604800, 7; 691200, 8; 777600, 9; 864000, 10; 950400, 11; 1036800, 12; 1123200, 13; 1209600, 14; 1296000, 15; 1382400, 16; 1468800, 17; 1555200, 18; 1641600, 19; 1728000, 20; 1814400, 21; 1900800, 22; 1987200, 23; 2073600, 24; 2160000, 25; 2246400, 26; 2332800, 27; 2419200, 28; 2505600, 29; 2592000, 30; 2678400, 31; 2764800, 1; 2851200, 2; 2937600, 3; 3024000, 4; 3110400, 5; 3196800, 6; 3283200, 7; 3369600, 8; 3456000, 9; 3542400, 10; 3628800, 11; 3715200, 12; 3801600, 13; 3888000, 14; 3974400, 15; 4060800, 16; 4147200, 17; 4233600, 18; 4320000, 19; 4406400, 20; 4492800, 21; 4579200, 22; 4665600, 23; 4752000, 24; 4838400, 25; 4924800, 26; 5011200, 27; 5097600, 28; 5184000, 1; 5270400, 2; 5356800, 3; 5443200, 4; 5529600, 5; 5616000, 6; 5702400, 7; 5788800, 8; 5875200, 9; 5961600, 10; 6048000, 11; 6134400, 12; 6220800, 13; 6307200, 14; 6393600, 15; 6480000, 16; 6566400, 17; 6652800, 18; 6739200, 19; 6825600, 20; 6912000, 21; 6998400, 22; 7084800, 23; 7171200, 24; 7257600, 25; 7344000, 26; 7430400, 27; 7516800, 28; 7603200, 29; 7689600, 30; 7776000, 31; 7862400, 1; 7948800, 2; 8035200, 3; 8121600, 4; 8208000, 5; 8294400, 6; 8380800, 7; 8467200, 8; 8553600, 9; 8640000, 10; 8726400, 11; 8812800, 12; 8899200, 13; 8985600, 14; 9072000, 15; 9158400, 16; 9244800, 17; 9331200, 18; 9417600, 19; 9504000, 20; 9590400, 21; 9676800, 22; 9763200, 23; 9849600, 24; 9936000, 25; 10022400, 26; 10108800, 27; 10195200, 28; 10281600, 29; 10368000, 30; 10454400, 1; 10540800, 2; 10627200, 3; 10713600, 4; 10800000, 5; 10886400, 6; 10972800, 7; 11059200, 8; 11145600, 9; 11232000, 10; 11318400, 11; 11404800, 12; 11491200, 13; 11577600, 14; 11664000, 15; 11750400, 16; 11836800, 17; 11923200, 18; 12009600, 19; 12096000, 20; 12182400, 21; 12268800, 22; 12355200, 23; 12441600, 24; 12528000, 25; 12614400, 26; 12700800, 27; 12787200, 28; 12873600, 29; 12960000, 30; 13046400, 31; 13132800, 1; 13219200, 2; 13305600, 3; 13392000, 4; 13478400, 5; 13564800, 6; 13651200, 7; 13737600, 8; 13824000, 9; 13910400, 10; 13996800, 11; 14083200, 12; 14169600, 13; 14256000, 14; 14342400, 15; 14428800, 16; 14515200, 17; 14601600, 18; 14688000, 19; 14774400, 20; 14860800, 21; 14947200, 22; 15033600, 23; 15120000, 24; 15206400, 25; 15292800, 26; 15379200, 27; 15465600, 28; 15552000, 29; 15638400, 30; 15724800, 1; 15811200, 2; 15897600, 3; 15984000, 4; 16070400, 5; 16156800, 6; 16243200, 7; 16329600, 8; 16416000, 9; 16502400, 10; 16588800, 11; 16675200, 12; 16761600, 13; 16848000, 14; 16934400, 15; 17020800, 16; 17107200, 17; 17193600, 18; 17280000, 19; 17366400, 20; 17452800, 21; 17539200, 22; 17625600, 23; 17712000, 24; 17798400, 25; 17884800, 26; 17971200, 27; 18057600, 28; 18144000, 29; 18230400, 30; 18316800, 31; 18403200, 1; 18489600, 2; 18576000, 3; 18662400, 4; 18748800, 5; 18835200, 6; 18921600, 7; 19008000, 8; 19094400, 9; 19180800, 10; 19267200, 11; 19353600, 12; 19440000, 13; 19526400, 14; 19612800, 15; 19699200, 16; 19785600, 17; 19872000, 18; 19958400, 19; 20044800, 20; 20131200, 21; 20217600, 22; 20304000, 23; 20390400, 24; 20476800, 25; 20563200, 26; 20649600, 27; 20736000, 28; 20822400, 29; 20908800, 30; 20995200, 31; 21081600, 1; 21168000, 2; 21254400, 3; 21340800, 4; 21427200, 5; 21513600, 6; 21600000, 7; 21686400, 8; 21772800, 9; 21859200, 10; 21945600, 11; 22032000, 12; 22118400, 13; 22204800, 14; 22291200, 15; 22377600, 16; 22464000, 17; 22550400, 18; 22636800, 19; 22723200, 20; 22809600, 21; 22896000, 22; 22982400, 23; 23068800, 24; 23155200, 25; 23241600, 26; 23328000, 27; 23414400, 28; 23500800, 29; 23587200, 30; 23673600, 1; 23760000, 2; 23846400, 3; 23932800, 4; 24019200, 5; 24105600, 6; 24192000, 7; 24278400, 8; 24364800, 9; 24451200, 10; 24537600, 11; 24624000, 12; 24710400, 13; 24796800, 14; 24883200, 15; 24969600, 16; 25056000, 17; 25142400, 18; 25228800, 19; 25315200, 20; 25401600, 21; 25488000, 22; 25574400, 23; 25660800, 24; 25747200, 25; 25833600, 26; 25920000, 27; 26006400, 28; 26092800, 29; 26179200, 30; 26265600, 31; 26352000, 1; 26438400, 2; 26524800, 3; 26611200, 4; 26697600, 5; 26784000, 6; 26870400, 7; 26956800, 8; 27043200, 9; 27129600, 10; 27216000, 11; 27302400, 12; 27388800, 13; 27475200, 14; 27561600, 15; 27648000, 16; 27734400, 17; 27820800, 18; 27907200, 19; 27993600, 20; 28080000, 21; 28166400, 22; 28252800, 23; 28339200, 24; 28425600, 25; 28512000, 26; 28598400, 27; 28684800, 28; 28771200, 29; 28857600, 30; 28944000, 1; 29030400, 2; 29116800, 3; 29203200, 4; 29289600, 5; 29376000, 6; 29462400, 7; 29548800, 8; 29635200, 9; 29721600, 10; 29808000, 11; 29894400, 12; 29980800, 13; 30067200, 14; 30153600, 15; 30240000, 16; 30326400, 17; 30412800, 18; 30499200, 19; 30585600, 20; 30672000, 21; 30758400, 22; 30844800, 23; 30931200, 24; 31017600, 25; 31104000, 26; 31190400, 27; 31276800, 28; 31363200, 29; 31449600, 30; 31536000, 31]);
        end Date2021;

        partial record DateBase
          extends Modelica.Icons.Record;
          parameter Real table_year[:, :] "公历年";
          parameter Real table_month[:, :] "公历月";
          parameter Real table_day[:, :] "公历日";
        end DateBase;
      end Date;
    end Production;

    package Controller "空压机控制单元"
      model CompressorOperation "空压机运行规律"
        parameter Real A[n] = {2, 2, 2, 2} "每种装机规模的质量流量组成的数组，从大到小依次排序";
        parameter Integer N[n] = {1, 1, 1, 1} "每种规模的数量";
        parameter Real AN[n](each unit = "kg/s") = A .* N "每种装机规模组成的数组{A1*N1,A2*N2...}";
        parameter Integer n = 4 "装机规模种类";
        Real sumA[n] "总的装机规模";
        Real SAN[n];
        //  Integer BN[n];
        Integer k "1<k<=Nn";
        Real m "临界值";
        Modelica.Blocks.Interfaces.RealInput Q(unit = "kg/s") "空气需求" annotation(
          Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
        Modelica.Blocks.Interfaces.RealOutput B[n] "设定流量" annotation(
          Placement(transformation(extent = {{100, 50}, {120, 70}})));
        Modelica.Blocks.Interfaces.RealOutput Bn[n] "每种装机规模开启的数量" annotation(
          Placement(transformation(extent = {{100, -70}, {120, -50}})));
      algorithm
        for i in 1:n loop
          sumA[i] := A[1:i] * N[1:i];
        end for;
        for i in 1:n - 1 loop
          SAN[i] := sumA[n] - sumA[n - i];
        end for;
        SAN[n] := sumA[n];
//全不开
        if Q <= 0 then
          B[1:n] := zeros(n);
          k := 0;
          Bn[1:n] := zeros(n);
        elseif Q <= SAN[1] then
          m := Q / A[n];
          k := if m > integer(m) then integer(m) + 1 else integer(m);
          B[1:n - 1] := fill(0, n - 1);
          B[n] := Q / A[n] / k;
          Bn[1:n - 1] := fill(0, n - 1);
          Bn[n] := k;
        elseif Q > SAN[n] then
          B[1:n] := ones(n);
          Bn := N;
        elseif Q > SAN[n - 1] then
          m := (Q - SAN[n - 1]) / A[1];
          k := if m > integer(m) then integer(m) + 1 else integer(m);
          B[1:n] := fill(Q / (SAN[n - 1] + k * A[1]), n);
          Bn[1] := k;
          Bn[2:n] := N[2:n];
        else
          for i in 2:n - 1 loop
            if Q > SAN[i - 1] and Q <= SAN[i] then
              m := (Q - SAN[i - 1]) / A[n - i + 1];
              k := if m > integer(m) then integer(m) + 1 else integer(m);
              B[1:n - i] := fill(0, n - i);
              B[n - i + 1:n] := fill(Q / (SAN[i - 1] + k * A[n - i + 1]), i);
              Bn[1:n - i] := fill(0, n - i);
              Bn[n - i + 1] := k;
              Bn[n - i + 2:n] := N[n - i + 2:n];
            end if;
          end for;
        end if;
//   else
//     B[1:n] := zeros(n);
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 20), Line(points = {{-84, 0}, {-54, 0}, {-54, 40}, {-24, 40}, {-24, -70}, {6, -70}, {6, 80}, {36, 80}, {36, -20}, {66, -20}, {66, 60}})}),
          Diagram(coordinateSystem(preserveAspectRatio = false)),
          Documentation(info = "<html>
<p><span style=\"font-family: 宋体;\">空气质量流量逐时负荷：Q kg/s</span> </p>
<p><span style=\"font-family: 宋体;\">单台电力空压机装机规模：</span>A1<span style=\"font-family: 宋体;\">，</span>...<span style=\"font-family: 宋体;\">，</span>Ai<span style=\"font-family: 宋体;\">，&hellip;，</span>An<span style=\"font-family: 宋体;\">（从小到大依次排序）</span>kW </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机装机规模对应的台数：</span>N1<span style=\"font-family: 宋体;\">（</span>N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1<span style=\"font-family: 宋体;\">）</span> </p>
<p><span style=\"font-family: 宋体;\">），&hellip;，</span>Ni(Ni-1, Ni-2,<span style=\"font-family: 宋体;\">&hellip;</span> Ni- Ni)<span style=\"font-family: 宋体;\">&hellip;，</span>Nn(Nn-1, Nn-2,<span style=\"font-family: 宋体;\">&hellip;</span> Nn- Nn) </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机质量流量：</span>B1(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1), <span style=\"font-family: 宋体;\">&hellip;，</span>Bi(Ni-1, Ni-2,<span style=\"font-family: 宋体;\">&hellip;</span> Ni- Ni)<span style=\"font-family: 宋体;\">，&hellip;，</span>Bn(Nn-1, Nn-2,<span style=\"font-family: 宋体;\">&hellip;</span> Nn- Nn) </p>
<p><span style=\"font-family: 宋体;\">电力空压机总装机规模（kg/s）</span>=A1*N1+<span style=\"font-family: 宋体;\">&hellip;</span>+Ai*Ni+<span style=\"font-family: 宋体;\">&hellip;</span>+An*Nn </p>
<p>1<span style=\"font-family: 宋体;\">、满开逻辑</span> </p>
<p>If&nbsp; Q&gt; A1*N1+&hellip;+Ai*Ni+&hellip;+An*Nn </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机质量流量：</span>B1(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1)= Bi(Ni-1, Ni-2,<span style=\"font-family: 宋体;\">&hellip;</span> Ni- Ni) = Bn(Nn-1, Nn-2,<span style=\"font-family: 宋体;\">&hellip;</span> Nn- Nn)=1 </p>
<p>2<span style=\"font-family: 宋体;\">、非满开</span> </p>
<p><span style=\"font-family: 宋体;\">第一种设备基本逻辑</span> </p>
<p>else if&nbsp; Q &gt; A1*(N1-1)+&hellip;+Ai*Ni+&hellip;+An*Nn </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机质量流量：</span>B1(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1)= Bi(Ni-1, Ni-2,<span style=\"font-family: 宋体;\">&hellip;</span> Ni- Ni) = Bn(Nn-1, Nn-2,<span style=\"font-family: 宋体;\">&hellip;</span> Nn- Nn)= Q/( A1*N1+<span style=\"font-family: 宋体;\">&hellip;</span>+Ai*Ni+<span style=\"font-family: 宋体;\">&hellip;</span>+An*Nn) </p>
<p>else if&nbsp; Q &gt; A1*(N1- M)+&hellip;+Ai*Ni+&hellip;+An*Nn </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机质量流量：</span>B1(N1- 1, <span style=\"font-family: 宋体;\">&hellip;</span>N1- (M-1))=0<span style=\"font-family: 宋体;\">，</span>B1(N1-M,<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1<span style=\"font-family: 宋体;\">）</span>) =Bi(Ni-1, Ni-2,<span style=\"font-family: 宋体;\">&hellip;</span>Ni- Ni) =Bn(Nn-1, Nn-2, <span style=\"font-family: 宋体;\">&hellip;</span>Nn- Nn)= Q/( A1*(N1- M+1)+<span style=\"font-family: 宋体;\">&hellip;</span>+Ai*Ni+<span style=\"font-family: 宋体;\">&hellip;</span>+An*Nn) </p>
<p><span style=\"font-family: 宋体;\">中间设备基本逻辑</span> </p>
<p>else if&nbsp; Q &gt;Ai*(Ni-1)+&hellip;+An*Nn </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机质量流量：</span>B1(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1)=0<span style=\"font-family: 宋体;\">；</span>Bi(Ni-1, Ni-2,<span style=\"font-family: 宋体;\">&hellip;</span> Ni- Ni) = Bn(Nn-1, Nn-2,<span style=\"font-family: 宋体;\">&hellip;</span> Nn- Nn)= Q/(Ai*Ni+<span style=\"font-family: 宋体;\">&hellip;</span>+An*Nn) </p>
<p>else if&nbsp; Q &gt; Ai*<span style=\"font-family: 宋体;\">（</span>Ni- M)+<span style=\"font-family: 宋体;\">&hellip;</span>+An*Nn </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机质量流量：</span>B1(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1)=0<span style=\"font-family: 宋体;\">；</span>Bi(N1- 1, <span style=\"font-family: 宋体;\">&hellip;</span>N1- (M-1))=0<span style=\"font-family: 宋体;\">，</span>Bi(N1-M,<span style=\"font-family: 宋体;\">&hellip;</span>N1-Ni<span style=\"font-family: 宋体;\">）</span>)= Bn(Nn-1, Nn-2, <span style=\"font-family: 宋体;\">&hellip;</span>Nn- Nn)= Q/(Ai*(Ni-M+1)+<span style=\"font-family: 宋体;\">&hellip;</span>+An*Nn) </p>
<p>&nbsp; </p>
<p><span style=\"font-family: 宋体;\">最后设备基本逻辑</span> </p>
<p>else if&nbsp; Q &gt; An*(Nn-1) </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机质量流量：</span>B1(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1)=0<span style=\"font-family: 宋体;\">；</span>Bi(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-Ni)=0<span style=\"font-family: 宋体;\">；</span> Bn(Nn-1, Nn-2,<span style=\"font-family: 宋体;\">&hellip;</span> Nn- Nn)= Q/(Ai*Ni+<span style=\"font-family: 宋体;\">&hellip;</span>+An*Nn) </p>
<p>else if&nbsp; Q &gt; An*<span style=\"font-family: 宋体;\">（</span>Nn-M<span style=\"font-family: 宋体;\">）</span> </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机质量流量：</span>B1(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1)=0<span style=\"font-family: 宋体;\">；</span>Bi(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-Ni)=0<span style=\"font-family: 宋体;\">；</span>Bn(N1- 1, <span style=\"font-family: 宋体;\">&hellip;</span>N1- (M-1))=0<span style=\"font-family: 宋体;\">，</span>Bn(N1-M,<span style=\"font-family: 宋体;\">&hellip;</span>N1-Nn<span style=\"font-family: 宋体;\">）</span>)= Q/( An*(Nn-M+1)) </p>
<p>&nbsp; </p>
<p>3<span style=\"font-family: 宋体;\">、不开</span> </p>
<p>else </p>
<p><span style=\"font-family: 宋体;\">每种电力空压机质量流量：</span>B1(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-N1)= Bi(N1-1,N1-2<span style=\"font-family: 宋体;\">&hellip;</span>N1-Ni)= Bn(Nn-1, Nn-2,<span style=\"font-family: 宋体;\">&hellip;</span> Nn- Nn)=0 </p>
</html>"));
      end CompressorOperation;
    end Controller;
  end Blocks;

  package Electrical
    model Transformer "变压器"
      ENN.Interfaces.Electrical.Pin_AC p annotation(
        Placement(transformation(extent = {{-10, 88}, {10, 108}}), iconTransformation(extent = {{-10, 88}, {10, 108}})));
      ENN.Interfaces.Electrical.Pin_AC n annotation(
        Placement(transformation(extent = {{-10, -110}, {10, -90}}), iconTransformation(extent = {{-10, -110}, {10, -90}})));
      parameter Modelica.SIunits.Power PowerCapacity = 8000000 "变压器容量";
      parameter Modelica.SIunits.Voltage V_ref = 220 "输出端电压";
      //parameter Real efficiency=0.99 "变压器效率";
      parameter Real lossEta = 0.4 "变压器损耗系数";
      Modelica.SIunits.Power power_load "输出功率";
      Modelica.SIunits.Power power_req "输入功率";
      Modelica.SIunits.Power loss "变压器损耗";
      //能流计算
      Real ele_in(unit = "kwh") "累计输入电量";
      Real ele_out(unit = "kwh") "累计输出电量";
      Real ele_loss(unit = "kwh") "累计损耗量";
    equation
      n.v = V_ref;
      power_load = -n.v * n.i;
      power_load = power_req * (1 - lossEta);
      loss = power_req * lossEta;
      -p.i = power_req / p.v;
      der(ele_in) = power_req / 1000 / 3600;
      der(ele_out) = power_load / 1000 / 3600;
      der(ele_loss) = loss / 1000 / 3600;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Text(extent = {{-100, 10}, {100, -10}}, lineColor = {0, 0, 0}, textString = "%name", origin = {-82, 0}, rotation = 90), Rectangle(extent = {{-60, -80}, {60, -94}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-54, 40}, {54, -68}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Polygon(points = {{-6, 22}, {12, 22}, {4, -2}, {26, -2}, {-2, -44}, {2, -14}, {-20, -14}, {-6, 22}}, lineColor = {255, 255, 255}, pattern = LinePattern.None, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-6, 84}, {6, 46}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-10, 56}, {10, 52}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-10, 68}, {10, 62}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-10, 78}, {10, 74}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-50, 78}, {-30, 74}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-50, 68}, {-30, 62}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-50, 56}, {-30, 52}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-46, 84}, {-34, 46}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{34, 84}, {46, 46}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{30, 56}, {50, 52}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{30, 68}, {50, 62}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{30, 78}, {50, 74}}, lineColor = {0, 0, 0}, pattern = LinePattern.None, fillColor = {0, 0, 0}, fillPattern = FillPattern.Solid)}),
        Documentation(revisions = "<html>
<hr><p><font color=\"#E72614\"><b>Copyright &copy; 2004-2020, MODELON AB</b></font> <font color=\"#AFAFAF\"><br /><br /> The use of this software component is regulated by the licensing conditions for Modelon Libraries. <br /> This copyright notice must, unaltered, accompany all components that are derived from, copied from, <br /> or by other means have their origin from any Modelon Library. </font></p>
</html>"));
    end Transformer;

    model Inverter "Inverter"
      parameter Modelica.SIunits.Voltage V_ref = 48 "Reference DC source voltage";
      parameter Real efficiency = 0.95 "Efficiency";
      Modelica.SIunits.Power power_in;
      Modelica.SIunits.Power power_out;
      Modelica.SIunits.Power loss;
      ENN.Interfaces.Electrical.Pin_AC p annotation(
        Placement(transformation(extent = {{70, -10}, {90, 10}}), iconTransformation(extent = {{70, -10}, {90, 10}})));
      Modelica.Electrical.Analog.Interfaces.NegativePin n annotation(
        Placement(transformation(extent = {{-90, -10}, {-70, 10}}), iconTransformation(extent = {{-90, -10}, {-70, 10}})));
    equation
      n.v = V_ref;
      power_in = n.v * n.i;
      power_out = power_in * efficiency;
      loss = power_out * (1 - efficiency);
      -p.i = power_out / p.v;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-80, 60}, {80, -60}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 5), Line(points = {{-74, -56}, {76, 56}}, color = {0, 0, 0}), Text(extent = {{4, -10}, {84, -50}}, lineColor = {0, 140, 72}, textString = "~"), Text(extent = {{-86, 56}, {-6, 16}}, lineColor = {0, 0, 255}, textString = "="), Text(extent = {{-80, -66}, {80, -80}}, lineColor = {0, 0, 0}, textString = "%name")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)),
        Documentation(revisions = "<html>
<hr><p><font color=\"#E72614\"><b>Copyright &copy; 2004-2020, MODELON AB</b></font> <font color=\"#AFAFAF\"><br /><br /> The use of this software component is regulated by the licensing conditions for Modelon Libraries. <br /> This copyright notice must, unaltered, accompany all components that are derived from, copied from, <br /> or by other means have their origin from any Modelon Library. </font></p>
</html>", info = "<html>
Simple model of DC-AC inverter.
</html>"));
    end Inverter;

    model SpecifiedPower "特定功率的电气附件"
      parameter Modelica.SIunits.Voltage u_min(min = Modelica.Constants.eps) "电气负载工作的最低电压";
      Modelica.SIunits.Voltage u;
      Modelica.SIunits.Current i;
      Modelica.Electrical.Analog.Interfaces.PositivePin p "Positive pin" annotation(
        Placement(transformation(extent = {{-110, 50}, {-90, 70}})));
      Modelica.Electrical.Analog.Interfaces.NegativePin n "Negative pin" annotation(
        Placement(transformation(extent = {{-110, -70}, {-90, -50}})));
      Modelica.Blocks.Interfaces.RealInput power "Power consumed by this component" annotation(
        Placement(transformation(extent = {{20, -20}, {-20, 20}}, rotation = 270, origin = {0, -120})));
      Modelica.Blocks.Interfaces.RealOutput power_act "Immediate power" annotation(
        Placement(transformation(extent = {{100, -10}, {120, 10}})));
    equation
      i = if noEvent(u < u_min) then power / u_min ^ 2 * u else power / u;
      u = p.v - n.v;
      i = p.i;
      0 = p.i + n.i;
      power_act = power;
      annotation(
        Diagram(graphics),
        Icon(graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {238, 238, 238}, fillPattern = FillPattern.Solid), Text(extent = {{-100, 90}, {100, 60}}, lineColor = {0, 0, 0}, textString = "%name"), Line(points = {{-70, 10}, {-24, 10}}, color = {95, 95, 95}, thickness = 0.5), Line(points = {{-70, -6}, {-24, -6}}, color = {95, 95, 95}, thickness = 0.5), Text(extent = {{-22, 22}, {10, -22}}, lineColor = {95, 95, 95}, lineThickness = 0.5, fillColor = {95, 95, 95}, fillPattern = FillPattern.Solid, textString = "P"), Line(points = {{18, 6}, {48, 6}, {48, 16}, {66, 0}, {48, -14}, {48, -4}, {18, -4}, {18, 6}}, color = {95, 95, 95}, thickness = 0.5)}),
        Documentation(revisions = "<html>
</html>", info = "<html>
<p>该模块模拟电气负载。</p>
<p>参数u_min定义最低的工作电压，可避免当负载开关闭合时，产生瞬时大电流。</p>
<p>输入信号P即power负载的功率消耗，负载按照如下公式生成电流：</p>
</html>"));
    end SpecifiedPower;

    model ElectricGrid "Electric grid"
      parameter Modelica.SIunits.Voltage V_ref = 10000 "Grid reference AC voltage";
      Modelica.Blocks.Interfaces.RealOutput P_grid(unit = "W") "Power output from grid" annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {44, 110}), iconTransformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {44, 110})));
      ENN.Interfaces.Electrical.Pin_AC p annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}}), iconTransformation(extent = {{90, -10}, {110, 10}})));
      // protected
      //   Modelica.SIunits.Power P_grid "Power output from grid";
    equation
      p.v = V_ref;
      P_grid = p.v * p.i;
//  P_net = P_grid;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 20), Text(extent = {{-100, -110}, {100, -130}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.None, textString = "%name"), Bitmap(extent = {{-78, -68}, {74, 68}}, imageSource = "/9j/4AAQSkZJRgABAQEASABIAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDL/2wBDAQkJCQwLDBgNDRgyIRwhMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjL/wAARCAIAAgADASIAAhEBAxEB/8QAHAABAAMBAQEBAQAAAAAAAAAAAAYHCAUEAwIB/8QAVhAAAQMCAwMGBgwLAwsFAQAAAQACAwQFBgcREiExE0FRYXGBCBQVIjKRFhgjQlJWcoKSlKHSFyQzN1Vik6Kxs9F1stM1NkNEU1RjdMHC4SVzlcPj8P/EABQBAQAAAAAAAAAAAAAAAAAAAAD/xAAUEQEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIRAxEAPwC/0REBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBfKpqqeippKmqniggjG0+WV4a1o6STuC+qqHwjJZI8vKNjHlrZLnG14Hvhych0PeAe5BbFLV01dTMqaSoiqIJBqyWJ4e1w6iNxX2WVMl8xzhO8+R7nNpZq5485x3U8p3B/U07ge482/VfEahAREQEREBERAREQEREBERAXNvWIbRh2lZU3i409FC92wx0z9No9AHEr03CvpbVb6ivrZmwUtPGZJZHcGtHFY2zExxVY7xPJXybUdFFrHRwE/k49eJ/WPE+rgAg2bTVMFZTRVNNNHNBK0PjkjcHNc08CCOIX1UCyWe5+UdiLnEnZmGpPMJ5APsU9QEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBU94R/wCb63/2rH/KlVwqnvCP/N9b/wC1Y/5UqDP0OFa+pwXNienbylJT1hpahoG+LzWlrj1Eu06jp0q+sjMx/LNvbha6za3CkZ+KSPO+aIe963NHrHYV8/B6pKevy9vdHVwsmp5658csbxqHNMTAQVVGPMJXLK/G0UtDNKyn5Txi3VY46A+iT8Jv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MD0+BMLRUI2X182ktZMPfyacAfgt4DvPOUEnoKGmtlBT0NFC2Gmp4xHFG3g1oGgC9CIgIiICIiAiIgIiICjmOMIUeNsL1NoqtGSOG3TzaamGUei7s5j0glSNEGOMI4huuVWPpG1kMjRFIaa4Uvw2a8R0kek08/YVr+hraa5UEFdRzNmpqiMSRSNO5zSNQVUWeuXnlu1HE9sh1uFFH+NMYN80I5+st/hr0BQbKPN2nwhb57NfzUSW4e6Ur4m7bonE+czTX0Tx6jr0oNDYrsEGKcLXGyz6BtVCWtcR6Dxva7ucAe5ZXywv8+BczIYq/WCKSV1BXMcfQ1dpqfkvAPYCrt9sDgf4dx+rf+VQ2Z95sGIsZTXjD5mENWxr52SxbBEo3Ega8CAD2koNmIoTlRir2W4AoaqV+3WUw8VqiTvL2Aece1pae0lTZAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREGXc+bpcKTMl0VNXVUMficR2I5nNGu/mBVU1NZVVr2vqqmadzRoDLIXEDvW1L3gHC2I7ga+72eCrqiwM5R7nA7I4DcQud+CPAXxbpfpv+8gxqi2V+CPAXxbpfpv8AvJ+CPAXxbpfpv+8gxxDNLTytlglfFI3e17HFpHYQvb5dvH6VrvrD/wCq11+CPAXxbpfpv+8n4I8BfFul+m/7yD35dSyTZc4elle6SR1DEXOcdSTs8SVJl56CgpbXQQUNFC2Glp2COKNuujWjgN69CAiIgIiICIiAiIgKrM/6qoo8u4ZKaeWCTyhENqJ5adNl+7UK01zb5h+1YloBQ3iiZV0weJBG8kDaAIB3EdJQYhN9u5BButcQebxh/wDVeBbK/BHgL4t0v03/AHk/BHgL4t0v03/eQY1RbK/BHgL4t0v03/eT8EeAvi3S/Tf95Bj6muNdRMLKWsqIGuOpbFK5oJ7itA+DhXVlbDiTxuqnn2HU2zyshds68rrpqrA/BHgL4t0v03/eXbw/hKw4WFQLJbYqIVGzyvJlx29nXTiTw1PrQdpERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREBERAREQEREH/2Q==", fileName = "modelica://ENN/Resources/icons/市电.jpg")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)),
        Documentation(info = "<html>
<h4>Description</h4>
<p>Ideal electric grid model with no limitations on the power flow.</p>
</html>", revisions = "<html>
<hr><p><font color=\"#E72614\"><b>Copyright &copy; 2004-2020, MODELON AB</b></font> <font color=\"#AFAFAF\"><br /><br /> The use of this software component is regulated by the licensing conditions for Modelon Libraries. <br /> This copyright notice must, unaltered, accompany all components that are derived from, copied from, <br /> or by other means have their origin from any Modelon Library. </font></p>
</html>"));
    end ElectricGrid;

    model GridControl "State graph-controller for microgrid"
      input Modelica.SIunits.Power P_renew "风电与光伏发电功率总和" annotation(
        Dialog(group = "Production"));
      input Modelica.SIunits.Power P_load "包括损耗在内的用电功率" annotation(
        Dialog(group = "Consumption"));
      Modelica.SIunits.Power P_net = P_renew - P_load;
      Modelica.Blocks.Interfaces.BooleanOutput y "断开电网" annotation(
        Placement(transformation(extent = {{100, -10}, {120, 10}})));
      Modelica.Blocks.Logical.Hysteresis hysteresis(uLow = 0, uHigh = 1) annotation(
        Placement(transformation(extent = {{40, -10}, {60, 10}})));
      Modelica.Blocks.Sources.RealExpression realExpression(y = P_net) annotation(
        Placement(transformation(extent = {{0, -10}, {20, 10}})));
    equation
      y = if P_net >= 0 then true else false;
      connect(hysteresis.u, realExpression.y) annotation(
        Line(points = {{38, 0}, {21, 0}}, color = {0, 0, 127}));
      annotation(
        Documentation(info = "<html>
</html>", revisions = "<html>
</html>"),
        Icon(graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {135, 135, 135}, fillColor = {135, 135, 135}, fillPattern = FillPattern.Solid), Text(extent = {{36, 92}, {-42, -94}}, lineColor = {255, 255, 255}, textString = "C")}));
    end GridControl;

    model Load
      ENN.Interfaces.Electrical.Pin_AC n annotation(
        Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
      parameter Modelica.SIunits.Power P_load = 100e3 "负载功率";
    equation
      n.i = P_load / n.v;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 20), Rectangle(extent = {{-60, -24}, {56, -62}}, lineColor = {175, 175, 175}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid, radius = 20), Rectangle(extent = {{-60, 14}, {56, -44}}, lineColor = {175, 175, 175}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-80, 28}, {76, 6}}, lineColor = {175, 175, 175}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid, radius = 10), Rectangle(extent = {{-31, 8}, {31, -8}}, lineColor = {0, 128, 255}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid, radius = 10, origin = {-28, 51}, rotation = 90, pattern = LinePattern.None), Rectangle(extent = {{-31, 8}, {31, -8}}, lineColor = {175, 175, 175}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid, radius = 10, origin = {28, 51}, rotation = 90), Text(extent = {{-100, -110}, {100, -130}}, lineColor = {0, 0, 0}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid, textString = "%name")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end Load;

    model SimpleGenerator "发电机"
      Modelica.Mechanics.Rotational.Interfaces.Flange_a flange_a "机械轴，输入机械功" annotation(
        Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
      Interfaces.Electrical.Pin_AC pin "电接口，输出电能" annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      parameter Modelica.SIunits.Voltage V_ref = 380 "额定电压";
      Modelica.SIunits.Power P_out "输出功率";
    equation
      pin.v * pin.i = flange_a.tau * max(der(flange_a.phi), 1e-3);
      pin.v = V_ref;
      P_out = pin.v * pin.i;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(lineColor = {128, 128, 128}, extent = {{-100.0, -100.0}, {100.0, 100.0}}, radius = 25.0), Rectangle(origin = {2.835, 10}, fillColor = {0, 128, 255}, fillPattern = FillPattern.HorizontalCylinder, extent = {{-60, -60}, {60, 60}}), Rectangle(origin = {2.835, 10}, fillColor = {128, 128, 128}, fillPattern = FillPattern.HorizontalCylinder, extent = {{-80, -60}, {-60, 60}}), Rectangle(origin = {2.835, 10}, fillColor = {95, 95, 95}, fillPattern = FillPattern.HorizontalCylinder, extent = {{60, -10}, {80, 10}}), Rectangle(origin = {2.835, 10}, lineColor = {95, 95, 95}, fillColor = {95, 95, 95}, fillPattern = FillPattern.Solid, extent = {{-60, 50}, {20, 70}}), Polygon(origin = {2.835, 10}, fillPattern = FillPattern.Solid, points = {{-70, -90}, {-60, -90}, {-30, -20}, {20, -20}, {50, -90}, {60, -90}, {60, -100}, {-70, -100}, {-70, -90}})}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end SimpleGenerator;
  end Electrical;

  package Furnace "窑炉"
    model TestGasGlassFurnace "燃气窑炉"
      extends Modelica.Icons.Example;
      ENN.Valves.MatValve solidValve annotation(
        Placement(transformation(extent = {{-50, 40}, {-30, 60}})));
      Components.GasGlassFurnance gasGlassFurnance "燃气玻璃窑炉" annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}})));
      ENN.Valves.GasValve gasValve annotation(
        Placement(transformation(extent = {{-50, -10}, {-30, 10}})));
      Modelica.Blocks.Sources.TimeTable Load(table = [0, 0.03; 2, 0.2; 3, 0.4; 4, 0.6; 5, 0.8; 6, 1; 7, 1.2; 8, 1.2]) annotation(
        Placement(transformation(extent = {{-80, 40}, {-60, 60}})));
      ENN.Valves.AirValve airValve annotation(
        Placement(transformation(extent = {{-50, -40}, {-30, -20}})));
      Sinks.ElectricalSink electricalSink annotation(
        Placement(transformation(extent = {{20, 40}, {0, 60}})));
      Sinks.ThermalSink thermalSink annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {0, -50})));
      Sinks.Sink sinkNone5 annotation(
        Placement(transformation(extent = {{40, -20}, {20, 0}})));
      Sinks.ProductSink productSink annotation(
        Placement(transformation(extent = {{42, 10}, {22, 30}})));
      Valves.ComAirValve comAirValve annotation(
        Placement(transformation(extent = {{-50, -70}, {-30, -50}})));
    equation
      connect(solidValve.m_mat, gasValve.dmMat) annotation(
        Line(points = {{-29, 55}, {-24, 55}, {-24, 32}, {-60, 32}, {-60, 0}, {-52, 0}}, color = {0, 0, 127}));
      connect(Load.y, solidValve.load) annotation(
        Line(points = {{-59, 50}, {-60, 50}, {-60, 50}, {-52, 50}}, color = {0, 0, 127}));
      connect(airValve.dmMat, solidValve.m_mat) annotation(
        Line(points = {{-52, -30}, {-60, -30}, {-60, 32}, {-24, 32}, {-24, 55}, {-29, 55}}, color = {0, 0, 127}));
      connect(solidValve.mat_out, gasGlassFurnance.mat_in) annotation(
        Line(points = {{-30, 50}, {-18, 50}, {-18, 7}, {-10, 7}}, color = {255, 170, 255}));
      connect(gasValve.gas_out, gasGlassFurnance.gas_in) annotation(
        Line(points = {{-30, 0}, {-10, 0}}, color = {255, 170, 85}));
      connect(gasGlassFurnance.flue_out, sinkNone5.generalFlowPort) annotation(
        Line(points = {{10, 0}, {14, 0}, {14, -10}, {20, -10}}, color = {95, 95, 95}, thickness = 0.5));
      connect(electricalSink.pin, gasGlassFurnance.pin) annotation(
        Line(points = {{0, 50}, {0, 10}}, color = {0, 0, 255}));
      connect(thermalSink.heatPort, gasGlassFurnance.heatPort) annotation(
        Line(points = {{6.66134e-16, -40}, {-0.2, -40}, {-0.2, -10}}, color = {191, 0, 0}));
      connect(gasGlassFurnance.pro_out, productSink.port_a) annotation(
        Line(points = {{10, 7}, {14, 7}, {14, 20}, {22, 20}}, color = {255, 170, 255}));
      connect(airValve.air_out, gasGlassFurnance.air_in) annotation(
        Line(points = {{-30, -30}, {-22, -30}, {-22, -4.8}, {-10, -4.8}}, color = {0, 127, 255}));
      connect(solidValve.m_mat, comAirValve.m_in) annotation(
        Line(points = {{-29, 55}, {-24, 55}, {-24, 32}, {-60, 32}, {-60, -60}, {-52, -60}}, color = {0, 0, 127}));
      connect(comAirValve.air_out, gasGlassFurnance.com_air_in) annotation(
        Line(points = {{-30, -60}, {-16, -60}, {-16, -9}, {-10, -9}}, color = {0, 127, 255}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false)),
        Diagram(coordinateSystem(preserveAspectRatio = false)),
        experiment(StopTime = 10));
    end TestGasGlassFurnace;

    package Components "组件"
      model GasGlassFurnance "燃气玻璃窑炉"
        parameter Real gas_perMat(unit = "m3/kg") = 0.14 "100%负荷下单位kg原料耗燃气量";
        parameter Real air_perGas(unit = "m3/m3") = 11.43 "单位体积燃气燃烧耗空气量(传递系数)";
        parameter Real pwr_perMat(unit = "kWh/kg") = 0.0629 "100%负荷率下单位质量原料耗电量";
        parameter Real comair_perMat(unit = "m3/kg") = 1.028571 "100%负荷率下单位质量原料耗压缩空气量";
        parameter Real h_gas(unit = "kJ/m3") = 34000 "天然气热值";
        parameter Modelica.SIunits.SpecificEnthalpy h_flue = 550.892e3 "烟气焓值（J/kg）";
        parameter Modelica.SIunits.SpecificEnthalpy h_glass = 180e3 "单位质量玻璃吸热量（J/kg）";
        parameter Modelica.SIunits.Temperature T_flue = 733.15 "烟气温度";
        //能流计算需求
        Real Q_flue(unit = "kWh") "年烟气累计热量";
        Real W_mech(unit = "kWh") "机械功年累计量";
        //SCOP计算
        input Real ele_factor(unit = "kgCO2/kWh") = 1 "加权平均后的电碳排放因子" annotation(
          Dialog(group = "SCOP计算"));
        parameter Real ratio = 1 "烟气碳排放占比" annotation(
          Dialog(group = "SCOP计算"));
        parameter Real gas_factor(unit = "kgCO2/m3") = 2.08280490528603 "天然气碳排放因子" annotation(
          Dialog(group = "SCOP计算"));
        //Real SCOP2_ele_0( unit="kgCO2/s") "SCOP2_电_燃气玻璃窑炉";
        Real SCOP2_ele(unit = "kgC02") "SCOP2_电_燃气玻璃窑炉年累计量";
        Real SCOP1_gas(unit = "kgC02") "SCOP1_天然气_燃气玻璃窑炉年累计量";
        //累计量计算
        Real gas_consumption(unit = "m3") "燃气年累计耗量_窑炉";
        Real ele_consumption(unit = "kWh") "电年累计耗量_窑炉";
        Interfaces.Materials.Port_a mat_in "原料输入接口" annotation(
          Placement(transformation(extent = {{-110, 60}, {-90, 80}})));
        Interfaces.Materials.Port_b pro_out "成品输出接口" annotation(
          Placement(transformation(extent = {{90, 60}, {110, 80}})));
        Interfaces.Gas.Port_a gas_in "天然气输入接口" annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Air.Port_a com_air_in "压缩空气" annotation(
          Placement(transformation(extent = {{-110, -100}, {-90, -80}})));
        Interfaces.FlueGas.Port_b flue_out "烟气输出接口" annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        Interfaces.Electrical.Pin_AC pin annotation(
          Placement(transformation(extent = {{-10, 90}, {10, 110}})));
        Sources.FlueGasSource2 flueSource_H annotation(
          Placement(transformation(extent = {{60, -30}, {80, -10}})));
        Sources.MatSource2 productSource_H annotation(
          Placement(transformation(extent = {{60, 40}, {80, 60}})));
        Modelica.Blocks.Tables.CombiTable1D m_FlueGas(table = [0.015081976, 0.232375621; 0.089154538, 1.373648992; 0.175673967, 2.706697521; 0.2596167, 4.000045598; 0.34104, 5.254575498; 0.42, 6.471152092; 0.504, 7.76538251]) "输出烟气质量流量kg/s，天然气流量(m3/s) vs 烟气质量流量（kg/s）" annotation(
          Placement(transformation(extent = {{20, 10}, {40, 30}})));
        Sources.FlueGasSource1 flueSource_m annotation(
          Placement(transformation(extent = {{60, 10}, {80, 30}})));
        Sensors.MassFLowSensor matSensor annotation(
          Placement(transformation(extent = {{-100, 80}, {-80, 100}})));
        Sources.ThermalSource thermalSource annotation(
          Placement(transformation(extent = {{-32, -80}, {-52, -60}})));
        Sources.ThermalSource thermalSource1 annotation(
          Placement(transformation(extent = {{40, -80}, {20, -60}})));
        Sources.ElectricalSource electricalSource annotation(
          Placement(transformation(extent = {{-10, 60}, {10, 80}})));
        Modelica.Blocks.Tables.CombiTable1D pwr_Ele(table = [0.1, 33942.85714; 0.6, 190080; 1.2, 353005.7143; 1.8, 488777.1429; 2.4, 597394.2857; 3, 678857.1429; 3.6, 814628.5714]) "总电耗计算，物料流量（kg/s）vs总耗电量（W）" annotation(
          Placement(transformation(extent = {{-40, 60}, {-20, 80}})));
        Modelica.Blocks.Tables.CombiTable1D pwr_Cooling(table = [0.015081976, 92147.99515; 0.089154538, 537534.9467; 0.175673967, 1044817.629; 0.2596167, 1522518.664; 0.34104, 1971295.453; 0.42, 2391792.43; 0.504, 2870150.916]) "输出空气、水带走热量(W)" annotation(
          Placement(transformation(extent = {{72, -80}, {52, -60}})));
        Modelica.Blocks.Tables.CombiTable1D pwr_Loss(table = [0.015081976, 112625.3274; 0.089154538, 656987.1571; 0.175673967, 1276999.324; 0.2596167, 1860856.144; 0.34104, 2409361.109; 0.42, 2923301.859; 0.504, 3507962.23]) "输出窑体散热量（W）" annotation(
          Placement(transformation(extent = {{0, -80}, {-20, -60}})));
        Sinks.Sink gasSink annotation(
          Placement(transformation(extent = {{-20, -10}, {-40, 10}})));
        Sensors.MassFLowSensor gasSensor annotation(
          Placement(transformation(extent = {{-78, -10}, {-58, 10}})));
        Modelica.Blocks.Math.Gain h_FlueGas(k = h_flue) "烟气焓（J/kg）" annotation(
          Placement(transformation(extent = {{20, -30}, {40, -10}})));
        Interfaces.Thermal.HeatPort heatPort annotation(
          Placement(transformation(extent = {{-12, -110}, {8, -90}})));
        Sources.GasSource2 gasSource_H annotation(
          Placement(transformation(extent = {{-60, -40}, {-80, -20}})));
        Modelica.Blocks.Tables.CombiTable1D V_gas(table = [0.1, 0.015081976; 0.6, 0.089154538; 1.2, 0.175673967; 1.8, 0.2596167; 2.4, 0.34104; 3, 0.42; 3.6, 0.504]) "输入天然气流量m3/s，原料流量（kg/s) vs 需求天然气流量（m3/s）" annotation(
          Placement(transformation(extent = {{-60, 10}, {-40, 30}})));
        Modelica.Blocks.Math.Gain h_Glass(k = h_glass) "单位质量玻璃吸热量（J/kg）" annotation(
          Placement(transformation(extent = {{20, 40}, {40, 60}})));
        Modelica.Blocks.Tables.CombiTable1D H_gas(table = [0.015081976, 512787.1858; 0.089154538, 3031254.301; 0.175673967, 5972914.878; 0.2596167, 8826967.8; 0.34104, 11595360; 0.42, 14280000; 0.504, 17136000]) "输入天然气热功率W，输入天然气流量（m3/s）vs 天然气热功率（W）" annotation(
          Placement(transformation(extent = {{-30, -40}, {-50, -20}})));
        Modelica.Blocks.Interfaces.RealOutput m_pro "产品流量" annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {60, 110})));
        Interfaces.Air.Port_b com_air_out "压缩空气输出" annotation(
          Placement(transformation(extent = {{88, -100}, {108, -80}})));
        Interfaces.Air.Port_a air_in annotation(
          Placement(transformation(extent = {{-110, -58}, {-90, -38}})));
        Sinks.Sink airSink annotation(
          Placement(transformation(extent = {{-60, -60}, {-80, -40}})));
        Modelica.Blocks.Interfaces.RealOutput W_out "机械功输出" annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {90, 110})));
      equation
        der(gas_consumption) = V_gas.y[1];
        der(ele_consumption) = pwr_Ele.y[1] / 1000 / 3600;
        SCOP1_gas = gas_consumption * gas_factor;
        SCOP2_ele = ele_consumption * ele_factor;
        der(Q_flue) = flueSource_H.H_in / 1000 / 3600;
        der(W_mech) = electricalSource.P_in / 1000 / 3600 * 0.13;
        W_out = W_mech;
        connect(m_FlueGas.y[1], flueSource_m.m_in) annotation(
          Line(points = {{41, 20}, {58, 20}}, color = {0, 0, 127}));
        connect(electricalSource.P_in, pwr_Ele.y[1]) annotation(
          Line(points = {{-12, 70}, {-19, 70}}, color = {0, 0, 127}));
        connect(pwr_Cooling.y[1], thermalSource1.Q_in) annotation(
          Line(points = {{51, -70}, {42, -70}}, color = {0, 0, 127}));
        connect(pwr_Loss.y[1], thermalSource.Q_in) annotation(
          Line(points = {{-21, -70}, {-30, -70}}, color = {0, 0, 127}));
        connect(electricalSource.pin, pin) annotation(
          Line(points = {{10, 70}, {20, 70}, {20, 84}, {0, 84}, {0, 100}}, color = {0, 0, 255}));
        connect(matSensor.Out, pro_out) annotation(
          Line(points = {{-80, 90}, {-22, 90}, {-22, 90}, {80, 90}, {80, 70}, {100, 70}}, color = {0, 0, 0}));
        connect(gasSensor.Out, gasSink.generalFlowPort) annotation(
          Line(points = {{-58, 0}, {-40, 0}}, color = {0, 0, 0}));
        connect(productSource_H.port_a, pro_out) annotation(
          Line(points = {{80, 50}, {90, 50}, {90, 70}, {100, 70}}, color = {255, 170, 255}));
        connect(flueSource_m.port_a, flue_out) annotation(
          Line(points = {{80, 20}, {90, 20}, {90, 0}, {100, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(flueSource_H.port_a, flue_out) annotation(
          Line(points = {{80, -20}, {90, -20}, {90, 0}, {100, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(gas_in, gasSensor.In) annotation(
          Line(points = {{-100, 0}, {-100, 0}, {-78, 0}}, color = {255, 170, 85}));
        connect(mat_in, matSensor.In) annotation(
          Line(points = {{-100, 70}, {-100, 90}}, color = {255, 170, 255}));
        connect(matSensor.m_flow, pwr_Ele.u[1]) annotation(
          Line(points = {{-79, 95.2}, {-58, 95.2}, {-58, 70}, {-42, 70}}, color = {0, 0, 127}));
        connect(h_FlueGas.y, flueSource_H.H_in) annotation(
          Line(points = {{41, -20}, {58, -20}}, color = {0, 0, 127}));
        connect(h_FlueGas.u, m_FlueGas.y[1]) annotation(
          Line(points = {{18, -20}, {8, -20}, {8, 0}, {52, 0}, {52, 20}, {41, 20}}, color = {0, 0, 127}));
        connect(thermalSource.heatPort, heatPort) annotation(
          Line(points = {{-52, -70}, {-56, -70}, {-56, -100}, {-2, -100}}, color = {191, 0, 0}));
        connect(gasSource_H.port_a, gasSink.generalFlowPort) annotation(
          Line(points = {{-80, -30}, {-80, -12}, {-40, -12}, {-40, 0}}, color = {255, 170, 85}));
        connect(V_gas.u[1], matSensor.m_flow) annotation(
          Line(points = {{-62, 20}, {-70, 20}, {-70, 95.2}, {-79, 95.2}}, color = {0, 0, 127}));
        connect(V_gas.y[1], m_FlueGas.u[1]) annotation(
          Line(points = {{-39, 20}, {18, 20}}, color = {0, 0, 127}));
        connect(matSensor.m_flow, h_Glass.u) annotation(
          Line(points = {{-79, 95.2}, {-70, 95.2}, {-70, 50}, {18, 50}}, color = {0, 0, 127}));
        connect(h_Glass.y, productSource_H.H_in) annotation(
          Line(points = {{41, 50}, {58, 50}}, color = {0, 0, 127}));
        connect(gasSource_H.H_in, H_gas.y[1]) annotation(
          Line(points = {{-58, -30}, {-51, -30}}, color = {0, 0, 127}));
        connect(H_gas.u[1], V_gas.y[1]) annotation(
          Line(points = {{-28, -30}, {-16, -30}, {-16, 20}, {-39, 20}}, color = {0, 0, 127}));
        connect(pwr_Loss.u[1], V_gas.y[1]) annotation(
          Line(points = {{2, -70}, {4, -70}, {4, 20}, {-39, 20}}, color = {0, 0, 127}));
        connect(pwr_Cooling.u[1], V_gas.y[1]) annotation(
          Line(points = {{74, -70}, {84, -70}, {84, -40}, {4, -40}, {4, 20}, {-39, 20}}, color = {0, 0, 127}));
        connect(matSensor.m_flow, m_pro) annotation(
          Line(points = {{-79, 95.2}, {60, 95.2}, {60, 110}}, color = {0, 0, 127}));
        connect(thermalSource1.heatPort, heatPort) annotation(
          Line(points = {{20, -70}, {16, -70}, {16, -100}, {-2, -100}}, color = {191, 0, 0}));
        connect(air_in, airSink.generalFlowPort) annotation(
          Line(points = {{-100, -48}, {-90, -48}, {-90, -50}, {-80, -50}}, color = {0, 127, 255}));
        connect(com_air_in, com_air_out) annotation(
          Line(points = {{-100, -90}, {-80, -90}, {-80, -60}, {80, -60}, {80, -90}, {98, -90}}, color = {0, 127, 255}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 15)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end GasGlassFurnance;
    end Components;
  end Furnace;

  package Machines
    package Boiler "锅炉"
      model TestFurnacewithBoiler
        extends Modelica.Icons.Example;
        parameter Media.Medium medium = ENN.Media.Water() "Cooling medium" annotation(
          choicesAllMatching = true);
        parameter Modelica.SIunits.Temperature Tamb(displayUnit = "degC") = 293.15 "Ambient temperature";
        WasteSteamBoiler wasteSteamBoiler(medium = medium, Tamb = Tamb) annotation(
          Placement(transformation(extent = {{10, -20}, {50, 20}})));
        Sinks.Sink sinkNone1 annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {30, -100})));
        Pipes.SteamPipe steamPipe annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = -90, origin = {30, -42})));
        Modelica.Thermal.FluidHeatFlow.Sources.Ambient ambient(constantAmbientTemperature = Tamb, medium = medium, constantAmbientPressure(displayUnit = "MPa") = 100000) annotation(
          Placement(transformation(extent = {{-60, 60}, {-80, 80}})));
        Pump.Pump pump_VF(medium = medium, idealPump(V_flow0 = 0.0075)) annotation(
          Placement(transformation(extent = {{-40, 60}, {-20, 80}})));
        Pipes.Pipe pipe(medium = medium, m = 0.1, T0 = Tamb, h_g = 0) annotation(
          Placement(transformation(extent = {{0, 60}, {20, 80}})));
        Sinks.ElectricalSink electricalSink annotation(
          Placement(transformation(extent = {{-100, 60}, {-80, 40}})));
        Modelica.Thermal.FluidHeatFlow.Sensors.MassFlowSensor massFlowSensor(medium = medium) annotation(
          Placement(transformation(extent = {{10, -10}, {-10, 10}}, rotation = 90, origin = {30, 40})));
        Modelica.Blocks.Continuous.LimPID PID(controllerType = Modelica.Blocks.Types.SimpleController.PI, k = 100, Ti = 0.1, yMax = 50, yMin = 0) annotation(
          Placement(transformation(extent = {{60, 80}, {40, 100}})));
        Sinks.Chimney chimney annotation(
          Placement(transformation(extent = {{60, -10}, {80, 10}})));
        Furnace.Components.GasGlassFurnance gasGlassFurnance annotation(
          Placement(transformation(extent = {{-80, -20}, {-40, 20}})));
        Valves.MatValve solidValve annotation(
          Placement(transformation(extent = {{-140, 4}, {-120, 24}})));
        Valves.GasValve fuelValve annotation(
          Placement(transformation(extent = {{-140, -26}, {-120, -6}})));
        Valves.AirValve airComValve annotation(
          Placement(transformation(extent = {{-140, -50}, {-120, -30}})));
        Sinks.ProductSink productSink annotation(
          Placement(transformation(extent = {{0, 10}, {-20, 30}})));
        Sinks.ThermalSink thermalSink annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {-60, -50})));
        Pipes.FluePipe fluePipe annotation(
          Placement(transformation(extent = {{-20, -10}, {0, 10}})));
        Blocks.Production.proGlass proGlass annotation(
          Placement(transformation(extent = {{-200, 4}, {-180, 24}})));
        Modelica.Blocks.Continuous.LimPID PID1(controllerType = Modelica.Blocks.Types.SimpleController.PI, k = 100, Ti = 0.1, yMax = 50, yMin = 0) annotation(
          Placement(transformation(extent = {{-170, 24}, {-150, 4}})));
        Sinks.Sink sink annotation(
          Placement(transformation(extent = {{0, -30}, {-20, -10}})));
        Turbine.SteamTurbine_nomech steamTurbine annotation(
          Placement(transformation(extent = {{10, 10}, {-10, -10}}, rotation = 90, origin = {30, -70})));
      equation
        connect(wasteSteamBoiler.steam_out, steamPipe.steam_in) annotation(
          Line(points = {{30, -20}, {30, -32}}, color = {255, 0, 0}));
        connect(pump_VF.port_a, ambient.flowPort) annotation(
          Line(points = {{-40, 70}, {-60, 70}}, color = {85, 255, 85}));
        connect(electricalSink.pin, pump_VF.pin) annotation(
          Line(points = {{-80, 50}, {-30, 50}, {-30, 60}}, color = {0, 0, 255}));
        connect(wasteSteamBoiler.water_in, massFlowSensor.flowPort_b) annotation(
          Line(points = {{30, 20}, {30, 30}}, color = {85, 255, 85}));
        connect(wasteSteamBoiler.m_flow, PID.u_s) annotation(
          Line(points = {{52, 12}, {96, 12}, {96, 90}, {62, 90}}, color = {0, 0, 127}));
        connect(PID.y, pump_VF.f) annotation(
          Line(points = {{39, 90}, {-30, 90}, {-30, 82}}, color = {0, 0, 127}));
        connect(pipe.port_b, massFlowSensor.flowPort_a) annotation(
          Line(points = {{20, 70}, {30, 70}, {30, 50}}, color = {85, 255, 85}));
        connect(pump_VF.port_b, pipe.port_a) annotation(
          Line(points = {{-20, 70}, {0, 70}}, color = {85, 255, 85}));
        connect(wasteSteamBoiler.flue_out, chimney.flue_in) annotation(
          Line(points = {{50, 0}, {60, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(massFlowSensor.y, PID.u_m) annotation(
          Line(points = {{41, 40}, {50, 40}, {50, 78}}, color = {0, 0, 127}));
        connect(gasGlassFurnance.pin, electricalSink.pin) annotation(
          Line(points = {{-60, 20}, {-60, 50}, {-80, 50}}, color = {0, 0, 255}));
        connect(solidValve.mat_out, gasGlassFurnance.mat_in) annotation(
          Line(points = {{-120, 14}, {-80, 14}}, color = {255, 170, 255}));
        connect(fuelValve.gas_out, gasGlassFurnance.gas_in) annotation(
          Line(points = {{-120, -16}, {-104, -16}, {-104, 0}, {-80, 0}}, color = {255, 170, 85}));
        connect(airComValve.air_out, gasGlassFurnance.com_air_in) annotation(
          Line(points = {{-120, -40}, {-88, -40}, {-88, -18}, {-80, -18}}, color = {0, 127, 255}));
        connect(gasGlassFurnance.pro_out, productSink.port_a) annotation(
          Line(points = {{-40, 14}, {-40, 20}, {-20, 20}}, color = {255, 170, 255}));
        connect(thermalSink.heatPort, gasGlassFurnance.heatPort) annotation(
          Line(points = {{-60, -40}, {-60, -20}, {-60.4, -20}}, color = {191, 0, 0}));
        connect(gasGlassFurnance.flue_out, fluePipe.flue_in) annotation(
          Line(points = {{-40, 0}, {-20, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(fluePipe.flue_out, wasteSteamBoiler.flue_in) annotation(
          Line(points = {{0, 0}, {10, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(proGlass.proGlass, PID1.u_s) annotation(
          Line(points = {{-179, 14}, {-172, 14}}, color = {0, 0, 127}));
        connect(gasGlassFurnance.m_pro, PID1.u_m) annotation(
          Line(points = {{-48, 22}, {-48, 34}, {-160, 34}, {-160, 26}}, color = {0, 0, 127}));
        connect(PID1.y, solidValve.load) annotation(
          Line(points = {{-149, 14}, {-142, 14}}, color = {0, 0, 127}));
        connect(solidValve.m_mat, fuelValve.dmMat) annotation(
          Line(points = {{-119, 19}, {-110, 19}, {-110, 0}, {-150, 0}, {-150, -16}, {-142, -16}}, color = {0, 0, 127}));
        connect(solidValve.m_mat, airComValve.dmMat) annotation(
          Line(points = {{-119, 19}, {-110, 19}, {-110, 0}, {-150, 0}, {-150, -40}, {-142, -40}}, color = {0, 0, 127}));
        connect(gasGlassFurnance.com_air_out, sink.generalFlowPort) annotation(
          Line(points = {{-40.4, -18}, {-40.4, -20}, {-20, -20}}, color = {0, 127, 255}));
        connect(steamPipe.steam_out, steamTurbine.steam_in) annotation(
          Line(points = {{30, -52}, {30, -60}}, color = {255, 0, 0}));
        connect(steamTurbine.steam_out, sinkNone1.generalFlowPort) annotation(
          Line(points = {{30, -80}, {30, -90}}, color = {255, 0, 0}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-200, -100}, {120, 100}})),
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-200, -100}, {120, 100}})),
          experiment(StopTime = 31536000));
      end TestFurnacewithBoiler;

      model TestBoiler
        extends Modelica.Icons.Example;
        parameter Media.Medium medium = ENN.Media.Water() "Cooling medium" annotation(
          choicesAllMatching = true);
        parameter Modelica.SIunits.Temperature Tamb(displayUnit = "degC") = 293.15 "Ambient temperature";
        WasteSteamBoiler wasteSteamBoiler(medium = medium, Tamb = Tamb) annotation(
          Placement(transformation(extent = {{-20, -20}, {20, 20}})));
        Sinks.Sink sinkNone1 annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {0, -70})));
        Sources.FlueGasSource2 smokeSource2_1 annotation(
          Placement(transformation(extent = {{-60, -10}, {-40, 10}})));
        Modelica.Blocks.Sources.Constant const(k = 4277.9e3) annotation(
          Placement(transformation(extent = {{-100, -10}, {-80, 10}})));
        Sources.FlueGasSource1 smokeSource1_1 annotation(
          Placement(transformation(extent = {{-60, -50}, {-40, -30}})));
        Modelica.Blocks.Sources.Constant const1(k = 7.77) annotation(
          Placement(transformation(extent = {{-100, -50}, {-80, -30}})));
        Pipes.SteamPipe steamPipe annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = -90, origin = {0, -42})));
        Modelica.Thermal.FluidHeatFlow.Sources.Ambient ambient(constantAmbientTemperature = Tamb, medium = medium, constantAmbientPressure(displayUnit = "MPa") = 100000) annotation(
          Placement(transformation(extent = {{-60, 60}, {-80, 80}})));
        Pump.Pump pump_VF(medium = medium, idealPump(V_flow0 = 0.0075)) annotation(
          Placement(transformation(extent = {{-40, 60}, {-20, 80}})));
        Pipes.Pipe pipe(medium = medium, m = 0.1, T0 = Tamb, h_g = 0) annotation(
          Placement(transformation(extent = {{0, 60}, {20, 80}})));
        Sinks.ElectricalSink electricalSink annotation(
          Placement(transformation(extent = {{-60, 48}, {-40, 28}})));
        Modelica.Thermal.FluidHeatFlow.Sensors.MassFlowSensor massFlowSensor(medium = medium) annotation(
          Placement(transformation(extent = {{10, -10}, {-10, 10}}, rotation = 90, origin = {0, 40})));
        Modelica.Blocks.Continuous.LimPID PID(controllerType = Modelica.Blocks.Types.SimpleController.PI, k = 100, Ti = 0.1, yMax = 50, yMin = 0) annotation(
          Placement(transformation(extent = {{60, 80}, {40, 100}})));
        Sinks.Chimney chimney annotation(
          Placement(transformation(extent = {{60, -10}, {80, 10}})));
      equation
        connect(smokeSource2_1.port_a, wasteSteamBoiler.flue_in) annotation(
          Line(points = {{-40, 0}, {-20, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(const.y, smokeSource2_1.H_in) annotation(
          Line(points = {{-79, 0}, {-62, 0}}, color = {0, 0, 127}));
        connect(const1.y, smokeSource1_1.m_in) annotation(
          Line(points = {{-79, -40}, {-62, -40}}, color = {0, 0, 127}));
        connect(smokeSource1_1.port_a, wasteSteamBoiler.flue_in) annotation(
          Line(points = {{-40, -40}, {-30, -40}, {-30, 0}, {-20, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(wasteSteamBoiler.steam_out, steamPipe.steam_in) annotation(
          Line(points = {{0, -20}, {0, -26}, {1.77636e-15, -26}, {1.77636e-15, -32}}, color = {255, 0, 0}));
        connect(steamPipe.steam_out, sinkNone1.generalFlowPort) annotation(
          Line(points = {{-1.77636e-15, -52}, {0, -60}, {6.66134e-16, -60}}, color = {255, 0, 0}));
        connect(pump_VF.port_a, ambient.flowPort) annotation(
          Line(points = {{-40, 70}, {-60, 70}}, color = {85, 255, 85}));
        connect(electricalSink.pin, pump_VF.pin) annotation(
          Line(points = {{-40, 38}, {-30, 38}, {-30, 60}}, color = {0, 0, 255}));
        connect(wasteSteamBoiler.water_in, massFlowSensor.flowPort_b) annotation(
          Line(points = {{0, 20}, {0, 30}, {-4.44089e-16, 30}}, color = {85, 255, 85}));
        connect(wasteSteamBoiler.m_flow, PID.u_s) annotation(
          Line(points = {{22, 12}, {96, 12}, {96, 90}, {62, 90}}, color = {0, 0, 127}));
        connect(PID.y, pump_VF.f) annotation(
          Line(points = {{39, 90}, {-30, 90}, {-30, 82}}, color = {0, 0, 127}));
        connect(pipe.port_b, massFlowSensor.flowPort_a) annotation(
          Line(points = {{20, 70}, {38, 70}, {38, 50}, {6.66134e-16, 50}}, color = {85, 255, 85}));
        connect(pump_VF.port_b, pipe.port_a) annotation(
          Line(points = {{-20, 70}, {0, 70}}, color = {85, 255, 85}));
        connect(wasteSteamBoiler.flue_out, chimney.flue_in) annotation(
          Line(points = {{20, 0}, {60, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(massFlowSensor.y, PID.u_m) annotation(
          Line(points = {{11, 40}, {50, 40}, {50, 78}}, color = {0, 0, 127}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false)),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end TestBoiler;

      model TestBoiler_nowater
        extends Modelica.Icons.Example;
        WasteSteamBoiler_nowater wasteSteamBoiler_nowater annotation(
          Placement(transformation(extent = {{-20, -20}, {20, 20}})));
        Sinks.Sink sinkNone annotation(
          Placement(transformation(extent = {{80, -10}, {60, 10}})));
        Sinks.Sink sinkNone1 annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {0, -70})));
        Sources.FlueGasSource2 smokeSource2_1 annotation(
          Placement(transformation(extent = {{-60, -10}, {-40, 10}})));
        Modelica.Blocks.Sources.Constant const(k = 4277.9e3) annotation(
          Placement(transformation(extent = {{-100, -10}, {-80, 10}})));
        Sources.FlueGasSource1 smokeSource1_1 annotation(
          Placement(transformation(extent = {{-60, -50}, {-40, -30}})));
        Modelica.Blocks.Sources.Constant const1(k = 7.77) annotation(
          Placement(transformation(extent = {{-100, -50}, {-80, -30}})));
        Pipes.SteamPipe steamPipe annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = -90, origin = {0, -42})));
      equation
        connect(wasteSteamBoiler_nowater.flue_out, sinkNone.generalFlowPort) annotation(
          Line(points = {{20, 0}, {60, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(smokeSource2_1.port_a, wasteSteamBoiler_nowater.flue_in) annotation(
          Line(points = {{-40, 0}, {-20, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(const.y, smokeSource2_1.H_in) annotation(
          Line(points = {{-79, 0}, {-62, 0}}, color = {0, 0, 127}));
        connect(const1.y, smokeSource1_1.m_in) annotation(
          Line(points = {{-79, -40}, {-62, -40}}, color = {0, 0, 127}));
        connect(smokeSource1_1.port_a, wasteSteamBoiler_nowater.flue_in) annotation(
          Line(points = {{-40, -40}, {-30, -40}, {-30, 0}, {-20, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(wasteSteamBoiler_nowater.steam_out, steamPipe.steam_in) annotation(
          Line(points = {{0, -20}, {0, -26}, {1.77636e-15, -26}, {1.77636e-15, -32}}, color = {255, 0, 0}));
        connect(steamPipe.steam_out, sinkNone1.generalFlowPort) annotation(
          Line(points = {{-1.77636e-15, -52}, {0, -60}, {6.66134e-16, -60}}, color = {255, 0, 0}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false)),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end TestBoiler_nowater;

      model WasteSteamBoiler "余热蒸汽锅炉"
        parameter Media.Medium medium = Media.Medium() "Medium";
        parameter Modelica.SIunits.Temperature Tamb = 293.15 "Ambient temperature";
        //内置参数表
        parameter Modelica.SIunits.Pressure p_steam = 800000 "蒸汽压力" annotation(
          Dialog(group = "蒸汽参数"));
        parameter Modelica.SIunits.Temperature T_steam = 453.15 "蒸汽温度" annotation(
          Dialog(group = "蒸汽参数"));
        parameter Modelica.SIunits.SpecificEnthalpy h_steam = 2792.43e3 "蒸汽焓值" annotation(
          Dialog(group = "蒸汽参数"));
        parameter Modelica.SIunits.Temperature T_water = 453.15 "锅炉凝结水温度" annotation(
          Dialog(group = "蒸汽参数"));
        parameter Real Pollution_rate = 0.07 "余热蒸汽锅炉排污率";
        parameter Modelica.SIunits.Temperature T_inlet = 733.15 "烟气入口温度" annotation(
          Dialog(group = "烟气参数"));
        parameter Modelica.SIunits.Temperature T_outlet = 423.15 "烟气出口温度" annotation(
          Dialog(group = "烟气参数"));
        parameter Modelica.SIunits.SpecificEnthalpy h_flue = 179.64e3 "烟气焓值" annotation(
          Dialog(group = "烟气参数"));
        parameter Modelica.SIunits.SpecificHeatCapacity cp_flue = 1.1976e3 "烟气比热" annotation(
          Dialog(group = "烟气参数"));
        parameter Modelica.SIunits.SpecificEnthalpy h_satWater = 763.188e3 "蒸汽温度下饱和水焓值";
        //能流计算需求
        Real Q_inlet(unit = "kWh") "累计入口烟气热量";
        Real Q_outlet(unit = "kWh") "累计出口烟气热量";
        Interfaces.FlueGas.Port_a flue_in "烟气输入接口" annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.FlueGas.Port_b flue_out "烟气输出接口" annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        Interfaces.Water_media.Port_a water_in(medium = medium) "水输入接口" annotation(
          Placement(transformation(extent = {{-10, 90}, {10, 110}})));
        Interfaces.Steam.Port_b steam_out "蒸汽输出接口" annotation(
          Placement(transformation(extent = {{-10, -110}, {10, -90}})));
        Sensors.MassFLowSensor massFLowSensor annotation(
          Placement(transformation(extent = {{-80, -10}, {-60, 10}})));
        Sources.FlueGasSource2 flueSource_H annotation(
          Placement(transformation(extent = {{0, 20}, {20, 40}})));
        Modelica.Blocks.Math.Gain gain(k = h_flue) "烟气焓值（J/kg）" annotation(
          Placement(transformation(extent = {{-40, 20}, {-20, 40}})));
        Modelica.Blocks.Math.Feedback feedback annotation(
          Placement(transformation(extent = {{-20, -20}, {0, -40}})));
        Sources.SteamSource1 steamSource_m annotation(
          Placement(transformation(extent = {{20, -70}, {40, -50}})));
        Sources.SteamSource2 steamSource_H annotation(
          Placement(transformation(extent = {{20, -40}, {40, -20}})));
        Sensors.MassFLowSensor massFLowSensor1 annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 0, origin = {64, -60})));
        Modelica.Blocks.Interfaces.RealOutput m_flow annotation(
          Placement(transformation(extent = {{100, 50}, {120, 70}})));
        Sources.Source sink(constantAmbientTemperature = Tamb, medium = medium, constantAmbientPressure(displayUnit = "MPa") = 1000000) annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = -90, origin = {0, 58})));
        Modelica.Thermal.FluidHeatFlow.Sensors.TemperatureSensor temperatureSensor(medium = medium) annotation(
          Placement(transformation(extent = {{2, 70}, {22, 90}})));
        Modelica.Blocks.Tables.CombiTable1D Enthalpy_water(table = [10, 41800; 12, 50160; 14, 58520; 16, 66880; 18, 75240; 20, 83600; 22, 91960; 24, 100320; 26, 108680; 28, 117040; 30, 125400; 32, 133760; 34, 142120; 36, 150480; 38, 158840; 40, 167200; 42, 175560; 44, 183920; 46, 192280; 48, 200640; 50, 209000; 52, 217360; 54, 225720; 56, 234080; 58, 242440; 60, 250800; 62, 259160; 64, 267520; 66, 275880; 68, 284240; 70, 292600; 72, 300960; 74, 309320; 76, 317680; 78, 326040; 80, 334400; 82, 342760; 84, 351120; 86, 359480; 88, 367840; 90, 376200; 92, 384560; 94, 392920; 96, 401280; 98, 409640; 100, 418000; 102, 426360; 104, 434720]) "输入锅炉补水温度degC，输出锅炉补水焓值（J/kg）" annotation(
          Placement(transformation(extent = {{40, 70}, {60, 90}})));
        Modelica.Blocks.Math.Division division annotation(
          Placement(transformation(extent = {{-40, -70}, {-20, -50}})));
        Modelica.Blocks.Sources.RealExpression dh(y = h_steam - Enthalpy_water.y[1]) "蒸汽焓值-锅炉补水焓值（J/kg）" annotation(
          Placement(transformation(extent = {{-100, -76}, {-68, -56}})));
        Sources.FlueGasSource1 flueSource_m annotation(
          Placement(transformation(extent = {{0, -10}, {20, 10}})));
        Sinks.Sink sinkNone annotation(
          Placement(transformation(extent = {{-20, -10}, {-40, 10}})));
        Modelica.Blocks.Sources.RealExpression H_flow(y = flue_in.H_flow) "蒸汽焓值-锅炉补水焓值（J/kg）" annotation(
          Placement(transformation(extent = {{-68, -40}, {-40, -20}})));
      equation
        der(Q_inlet) = flue_in.H_flow / 1000 / 3600;
        der(Q_outlet) = flueSource_H.H_in / 1000 / 3600;
        connect(flue_in, massFLowSensor.In) annotation(
          Line(points = {{-100, 0}, {-80, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(gain.y, flueSource_H.H_in) annotation(
          Line(points = {{-19, 30}, {-2, 30}}, color = {0, 0, 127}));
        connect(gain.u, massFLowSensor.m_flow) annotation(
          Line(points = {{-42, 30}, {-48, 30}, {-48, 5.2}, {-59, 5.2}}, color = {0, 0, 127}));
        connect(flueSource_H.port_a, flue_out) annotation(
          Line(points = {{20, 30}, {60, 30}, {60, 0}, {100, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(flue_in, flue_in) annotation(
          Line(points = {{-100, 0}, {-100, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(feedback.u2, gain.y) annotation(
          Line(points = {{-10, -22}, {-10, 30}, {-19, 30}}, color = {0, 0, 127}));
        connect(feedback.y, steamSource_H.H_in) annotation(
          Line(points = {{-1, -30}, {18, -30}}, color = {0, 0, 127}));
        connect(steamSource_m.port_a, massFLowSensor1.In) annotation(
          Line(points = {{40, -60}, {54, -60}}, color = {255, 0, 0}));
        connect(massFLowSensor1.Out, steam_out) annotation(
          Line(points = {{74, -60}, {74, -100}, {0, -100}}));
        connect(massFLowSensor1.m_flow, m_flow) annotation(
          Line(points = {{75, -54.8}, {80, -54.8}, {80, 60}, {110, 60}}, color = {0, 0, 127}));
        connect(water_in, sink.flowPort) annotation(
          Line(points = {{0, 100}, {0, 68}, {1.77636e-15, 68}}, color = {85, 255, 85}));
        connect(water_in, temperatureSensor.flowPort) annotation(
          Line(points = {{0, 100}, {0, 80}, {2, 80}}, color = {85, 255, 85}));
        connect(temperatureSensor.y, Enthalpy_water.u[1]) annotation(
          Line(points = {{23, 80}, {38, 80}}, color = {0, 0, 127}));
        connect(division.y, steamSource_m.m_in) annotation(
          Line(points = {{-19, -60}, {18, -60}}, color = {0, 0, 127}));
        connect(feedback.y, division.u1) annotation(
          Line(points = {{-1, -30}, {6, -30}, {6, -42}, {-60, -42}, {-60, -54}, {-42, -54}}, color = {0, 0, 127}));
        connect(dh.y, division.u2) annotation(
          Line(points = {{-66.4, -66}, {-42, -66}}, color = {0, 0, 127}));
        connect(steamSource_H.port_a, steam_out) annotation(
          Line(points = {{40, -30}, {48, -30}, {48, -100}, {0, -100}}, color = {255, 0, 0}));
        connect(massFLowSensor.Out, sinkNone.generalFlowPort) annotation(
          Line(points = {{-60, 0}, {-40, 0}}, color = {0, 0, 0}));
        connect(H_flow.y, feedback.u1) annotation(
          Line(points = {{-38.6, -30}, {-18, -30}}, color = {0, 0, 127}));
        connect(flueSource_m.m_in, massFLowSensor.m_flow) annotation(
          Line(points = {{-2, 0}, {-4, 0}, {-4, 14}, {-48, 14}, {-48, 5.2}, {-59, 5.2}}, color = {0, 0, 127}));
        connect(flueSource_m.port_a, flue_out) annotation(
          Line(points = {{20, 0}, {100, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(m_flow, m_flow) annotation(
          Line(points = {{110, 60}, {110, 60}}, color = {0, 0, 127}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-60, 80}, {60, -80}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end WasteSteamBoiler;

      model WasteSteamBoiler_nowater "余热蒸汽锅炉"
        parameter Media.Medium medium = Media.Medium() "Medium";
        parameter Modelica.SIunits.Temperature Tamb = 293.15 "Ambient temperature";
        //内置参数表
        parameter Modelica.SIunits.Pressure p_steam = 800000 "蒸汽压力" annotation(
          Dialog(group = "蒸汽参数"));
        parameter Modelica.SIunits.Temperature T_steam = 453.15 "蒸汽温度" annotation(
          Dialog(group = "蒸汽参数"));
        parameter Modelica.SIunits.SpecificEnthalpy h_steam = 2792.43e3 "蒸汽焓值" annotation(
          Dialog(group = "蒸汽参数"));
        parameter Modelica.SIunits.Temperature T_water = 453.15 "锅炉凝结水温度" annotation(
          Dialog(group = "蒸汽参数"));
        parameter Real Pollution_rate = 0.07 "余热蒸汽锅炉排污率";
        parameter Modelica.SIunits.Temperature T_outlet = 423.15 "烟气出口温度" annotation(
          Dialog(group = "烟气参数"));
        parameter Modelica.SIunits.SpecificEnthalpy h_flue = 179.64e3 "烟气焓值" annotation(
          Dialog(group = "烟气参数"));
        parameter Modelica.SIunits.SpecificHeatCapacity cp_flue = 1.1976e3 "烟气比热" annotation(
          Dialog(group = "烟气参数"));
        Interfaces.FlueGas.Port_a flue_in "烟气输入接口" annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.FlueGas.Port_b flue_out "烟气输出接口" annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        Interfaces.Steam.Port_b steam_out "蒸汽输出接口" annotation(
          Placement(transformation(extent = {{-10, -110}, {10, -90}})));
        Sensors.MassFLowSensor massFLowSensor annotation(
          Placement(transformation(extent = {{-78, -10}, {-58, 10}})));
        Sources.FlueGasSource2 smokeSource2_1 annotation(
          Placement(transformation(extent = {{0, 20}, {20, 40}})));
        Modelica.Blocks.Math.Gain gain(k = h_flue) "烟气焓值（J/kg）" annotation(
          Placement(transformation(extent = {{-40, 20}, {-20, 40}})));
        Modelica.Blocks.Math.Feedback feedback annotation(
          Placement(transformation(extent = {{-20, -20}, {0, -40}})));
        Sources.SteamSource1 steamSource1 annotation(
          Placement(transformation(extent = {{20, -70}, {40, -50}})));
        Sources.SteamSource2 steamSource2 annotation(
          Placement(transformation(extent = {{20, -40}, {40, -20}})));
        Sensors.MassFLowSensor massFLowSensor1 annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 0, origin = {64, -60})));
        Modelica.Blocks.Interfaces.RealOutput m_flow annotation(
          Placement(transformation(extent = {{100, -70}, {120, -50}})));
        Modelica.Blocks.Math.Division division annotation(
          Placement(transformation(extent = {{-40, -70}, {-20, -50}})));
        Modelica.Blocks.Sources.RealExpression dh(y = h_steam - 83.6e3) "蒸汽焓值-锅炉补水焓值（J/kg）" annotation(
          Placement(transformation(extent = {{-100, -76}, {-68, -56}})));
        Sources.FlueGasSource1 smokeSource1_1 annotation(
          Placement(transformation(extent = {{0, -10}, {20, 10}})));
        Sinks.Sink sinkNone annotation(
          Placement(transformation(extent = {{-20, -20}, {-40, 0}})));
        Modelica.Blocks.Sources.RealExpression H_flow(y = flue_in.H_flow) "蒸汽焓值-锅炉补水焓值（J/kg）" annotation(
          Placement(transformation(extent = {{-50, -40}, {-30, -20}})));
      equation
        connect(flue_in, massFLowSensor.In) annotation(
          Line(points = {{-100, 0}, {-78, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(gain.y, smokeSource2_1.H_in) annotation(
          Line(points = {{-19, 30}, {-2, 30}}, color = {0, 0, 127}));
        connect(gain.u, massFLowSensor.m_flow) annotation(
          Line(points = {{-42, 30}, {-48, 30}, {-48, 5.2}, {-57, 5.2}}, color = {0, 0, 127}));
        connect(smokeSource2_1.port_a, flue_out) annotation(
          Line(points = {{20, 30}, {60, 30}, {60, 0}, {100, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(flue_in, flue_in) annotation(
          Line(points = {{-100, 0}, {-100, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(feedback.u2, gain.y) annotation(
          Line(points = {{-10, -22}, {-10, 30}, {-19, 30}}, color = {0, 0, 127}));
        connect(feedback.y, steamSource2.H_in) annotation(
          Line(points = {{-1, -30}, {18, -30}}, color = {0, 0, 127}));
        connect(steamSource1.port_a, massFLowSensor1.In) annotation(
          Line(points = {{40, -60}, {54, -60}}, color = {255, 0, 0}));
        connect(massFLowSensor1.Out, steam_out) annotation(
          Line(points = {{74, -60}, {74, -100}, {0, -100}}));
        connect(massFLowSensor1.m_flow, m_flow) annotation(
          Line(points = {{75, -54.8}, {80, -54.8}, {80, -60}, {110, -60}}, color = {0, 0, 127}));
        connect(division.y, steamSource1.m_in) annotation(
          Line(points = {{-19, -60}, {18, -60}}, color = {0, 0, 127}));
        connect(feedback.y, division.u1) annotation(
          Line(points = {{-1, -30}, {6, -30}, {6, -42}, {-60, -42}, {-60, -54}, {-42, -54}}, color = {0, 0, 127}));
        connect(dh.y, division.u2) annotation(
          Line(points = {{-66.4, -66}, {-42, -66}}, color = {0, 0, 127}));
        connect(steamSource2.port_a, steam_out) annotation(
          Line(points = {{40, -30}, {48, -30}, {48, -100}, {0, -100}}, color = {255, 0, 0}));
        connect(massFLowSensor.m_flow, smokeSource1_1.m_in) annotation(
          Line(points = {{-57, 5.2}, {-6, 5.2}, {-6, 0}, {-2, 0}}, color = {0, 0, 127}));
        connect(smokeSource1_1.port_a, flue_out) annotation(
          Line(points = {{20, 0}, {100, 0}}, color = {95, 95, 95}, thickness = 0.5));
        connect(sinkNone.generalFlowPort, massFLowSensor.Out) annotation(
          Line(points = {{-40, -10}, {-58, -10}, {-58, 0}}, color = {0, 0, 0}));
        connect(H_flow.y, feedback.u1) annotation(
          Line(points = {{-29, -30}, {-18, -30}}, color = {0, 0, 127}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-60, 80}, {60, -80}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end WasteSteamBoiler_nowater;
    end Boiler;

    package Compressor
      model TestCompressor
        extends Modelica.Icons.Example;
        EletricalCompressor eletricalCompressor annotation(
          Placement(transformation(extent = {{10, 10}, {30, 30}})));
        parameter Modelica.Thermal.FluidHeatFlow.Media.Medium medium = Modelica.Thermal.FluidHeatFlow.Media.Medium() "Medium in the component";
        parameter Modelica.SIunits.Temperature Tamb = 293.15 "Initial temperature of medium";
        Sources.AirSource1 airSource1_1 annotation(
          Placement(transformation(extent = {{-30, 10}, {-10, 30}})));
        Modelica.Blocks.Sources.Ramp air(height = 1.5, duration = 5, startTime = 5) annotation(
          Placement(transformation(extent = {{-100, 50}, {-80, 70}})));
        EletricalCompressor eletricalCompressor1 annotation(
          Placement(transformation(extent = {{10, -50}, {30, -30}})));
        Sinks.ElectricalSink electricalSink1 annotation(
          Placement(transformation(extent = {{80, -80}, {60, -60}})));
        Sources.AirSource1 airSource1_2 annotation(
          Placement(transformation(extent = {{-30, -50}, {-10, -30}})));
        Sinks.Sink sinkNone1 annotation(
          Placement(transformation(extent = {{100, 10}, {80, 30}})));
        Blocks.Controller.CompressorOperation compressorOperation(A = {1, 0.5}, N = {1, 1}, n = 2) annotation(
          Placement(transformation(extent = {{-70, 50}, {-50, 70}})));
        Pipes.AirPipe airPipe1_1 annotation(
          Placement(transformation(extent = {{50, 10}, {70, 30}})));
      equation
        connect(airSource1_1.port_a, eletricalCompressor.air_in) annotation(
          Line(points = {{-10, 20}, {10, 20}}, color = {0, 127, 255}));
        connect(electricalSink1.pin, eletricalCompressor1.pin) annotation(
          Line(points = {{60, -70}, {20, -70}, {20, -50}}, color = {0, 0, 255}));
        connect(airSource1_2.port_a, eletricalCompressor1.air_in) annotation(
          Line(points = {{-10, -40}, {10, -40}}, color = {0, 127, 255}));
        connect(eletricalCompressor.pin, electricalSink1.pin) annotation(
          Line(points = {{20, 10}, {20, 0}, {50, 0}, {50, -70}, {60, -70}}, color = {0, 0, 255}));
        connect(air.y, compressorOperation.Q) annotation(
          Line(points = {{-79, 60}, {-72, 60}}, color = {0, 0, 127}));
        connect(compressorOperation.B[1], airSource1_1.m_in) annotation(
          Line(points = {{-49, 65.5}, {-42, 65.5}, {-42, 20}, {-32, 20}}, color = {0, 0, 127}));
        connect(compressorOperation.B[2], airSource1_2.m_in) annotation(
          Line(points = {{-49, 66.5}, {-42, 66.5}, {-42, -40}, {-32, -40}}, color = {0, 0, 127}));
        connect(airPipe1_1.air_out, sinkNone1.generalFlowPort) annotation(
          Line(points = {{70, 20}, {80, 20}}, color = {0, 127, 255}));
        connect(eletricalCompressor.air_out, airPipe1_1.air_in) annotation(
          Line(points = {{30, 20}, {50, 20}}, color = {0, 127, 255}));
        connect(eletricalCompressor1.air_out, airPipe1_1.air_in) annotation(
          Line(points = {{30, -40}, {40, -40}, {40, 20}, {50, 20}}, color = {0, 127, 255}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false)),
          Diagram(coordinateSystem(preserveAspectRatio = false)),
          experiment(StopTime = 10));
      end TestCompressor;

      model Compressor
        parameter Modelica.SIunits.Density rho_air = 1.29 "空气密度";
        Modelica.Blocks.Tables.CombiTable1D pwrChar(table = [1.03, 1617.408; 1.08, 1683.4584; 1.14, 1747.3248; 1.2, 1809.0072; 1.25, 1868.5056; 1.31, 1925.82; 1.36, 1980.9504; 1.42, 2033.8968; 1.47, 2084.6592; 1.53, 2133.2376]) "空压机能耗曲线，流量（m3/s) vs 电功率（W）" annotation(
          Placement(transformation(extent = {{-40, -70}, {-20, -50}})));
        Sources.ElectricalSource electricalSource annotation(
          Placement(transformation(extent = {{0, -70}, {20, -50}})));
        Sensors.MassFLowSensor dmSensor annotation(
          Placement(transformation(extent = {{-80, -10}, {-60, 10}})));
        Sources.AirSource2 airSource2_1 annotation(
          Placement(transformation(extent = {{40, -30}, {60, -10}})));
        Modelica.Blocks.Tables.CombiTable1D dpChar(table = [1.03, 700000; 1.08, 690000; 1.14, 680000; 1.2, 670000; 1.25, 660000; 1.31, 650000; 1.36, 640000; 1.42, 630000; 1.47, 620000; 1.53, 610000]) "空压机压差曲线，流量（m3/s) vs dp（pa）" annotation(
          Placement(transformation(extent = {{-6, 40}, {14, 60}})));
        Modelica.Blocks.Math.Gain dV_dm(k = 1 / rho_air) "默认空气密度=1.29kg/m3" annotation(
          Placement(transformation(extent = {{-50, 10}, {-30, 30}})));
        Modelica.Blocks.Math.Gain etaChar(k = 0.5) "空压机效率曲线，流量（m3/s) vs 效率" annotation(
          Placement(transformation(extent = {{-8, 10}, {12, 30}})));
        Modelica.Blocks.Math.Product product annotation(
          Placement(transformation(extent = {{0, -10}, {20, -30}})));
        Interfaces.Air.Port_a air_in annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Air.Port_b air_out annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        Interfaces.Electrical.Pin_AC pin annotation(
          Placement(transformation(extent = {{-10, -110}, {10, -90}})));
        Modelica.Blocks.Sources.RealExpression p_in(y = 0.1) "入口压力" annotation(
          Placement(transformation(extent = {{-10, 66}, {10, 86}})));
        Modelica.Blocks.Math.Add p_out "出口压力" annotation(
          Placement(transformation(extent = {{40, 60}, {60, 80}})));
      equation
        connect(pwrChar.y[1], electricalSource.P_in) annotation(
          Line(points = {{-19, -60}, {-2, -60}}, color = {0, 0, 127}));
        connect(electricalSource.pin, pin) annotation(
          Line(points = {{20, -60}, {26, -60}, {26, -100}, {0, -100}}, color = {0, 0, 255}));
        connect(air_in, dmSensor.In) annotation(
          Line(points = {{-100, 0}, {-90, 0}, {-90, 0}, {-80, 0}}, color = {0, 127, 255}));
        connect(dmSensor.Out, air_out) annotation(
          Line(points = {{-60, 0}, {16, 0}, {16, 0}, {100, 0}, {100, 0}}, color = {0, 0, 0}));
        connect(airSource2_1.port_a, air_out) annotation(
          Line(points = {{60, -20}, {100, -20}, {100, 0}}, color = {0, 127, 255}));
        connect(dmSensor.m_flow, dV_dm.u) annotation(
          Line(points = {{-59, 5.2}, {-59, 20}, {-52, 20}}, color = {0, 0, 127}));
        connect(dV_dm.y, dpChar.u[1]) annotation(
          Line(points = {{-29, 20}, {-20, 20}, {-20, 50}, {-8, 50}}, color = {0, 0, 127}));
        connect(dV_dm.y, pwrChar.u[1]) annotation(
          Line(points = {{-29, 20}, {-24, 20}, {-24, -30}, {-60, -30}, {-60, -60}, {-42, -60}}, color = {0, 0, 127}));
        connect(product.y, airSource2_1.H_in) annotation(
          Line(points = {{21, -20}, {38, -20}}, color = {0, 0, 127}));
        connect(pwrChar.y[1], product.u1) annotation(
          Line(points = {{-19, -60}, {-19, -26}, {-2, -26}}, color = {0, 0, 127}));
        connect(p_in.y, p_out.u1) annotation(
          Line(points = {{11, 76}, {38, 76}}, color = {0, 0, 127}));
        connect(dpChar.y[1], p_out.u2) annotation(
          Line(points = {{15, 50}, {26, 50}, {26, 64}, {38, 64}}, color = {0, 0, 127}));
        connect(dV_dm.y, etaChar.u) annotation(
          Line(points = {{-29, 20}, {-10, 20}}, color = {0, 0, 127}));
        connect(etaChar.y, product.u2) annotation(
          Line(points = {{13, 20}, {26, 20}, {26, -6}, {-14, -6}, {-14, -14}, {-2, -14}}, color = {0, 0, 127}));
        annotation(
          Icon(graphics = {Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Line(points = {{-44, 66}, {64, 46}}, color = {28, 108, 200}, smooth = Smooth.Bezier), Line(points = {{-44, -66}, {64, -46}}, color = {28, 108, 200}, smooth = Smooth.Bezier)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end Compressor;

      model EletricalCompressor "电动空压机"
        parameter Modelica.SIunits.Density rho_air = 1.29 "空气密度";
        parameter Real eta = 0.7 "空压机效率";
        //能流计算需求
        Real WP_out(unit = "kWh") "累计势能";
        Real Q_out(unit = "kWh") "累计耗电";
        Modelica.Blocks.Tables.CombiTable1D pwrChar(table = [0.04, 1354.162416; 0.25, 8004.900979; 0.5, 15773.2039; 0.74, 23310.15355; 0.97, 30620.89136; 1.2, 37710.45734; 1.44, 45252.5488]) "电动空压机能耗曲线，流量（m3/s) vs 电功率（W）" annotation(
          Placement(transformation(extent = {{-40, -70}, {-20, -50}})));
        Sources.ElectricalSource electricalSource annotation(
          Placement(transformation(extent = {{0, -70}, {20, -50}})));
        Sensors.MassFLowSensor dmSensor annotation(
          Placement(transformation(extent = {{-80, -10}, {-60, 10}})));
        Sources.AirSource2 airSource_H annotation(
          Placement(transformation(extent = {{40, -30}, {60, -10}})));
        Modelica.Blocks.Sources.Constant dpChar(k = 0.8e6) "空压机压差曲线 dp（pa）" annotation(
          Placement(transformation(extent = {{-8, 40}, {12, 60}})));
        Modelica.Blocks.Math.Gain dV_dm(k = 1 / rho_air) "默认空气密度=1.29kg/m3" annotation(
          Placement(transformation(extent = {{-50, 10}, {-30, 30}})));
        Modelica.Blocks.Math.Gain etaChar(k = eta) "空压机效率曲线，流量（m3/s) vs 效率" annotation(
          Placement(transformation(extent = {{-8, 10}, {12, 30}})));
        Modelica.Blocks.Math.Product product annotation(
          Placement(transformation(extent = {{0, -10}, {20, -30}})));
        Interfaces.Air.Port_a air_in annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Air.Port_b air_out annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        Interfaces.Electrical.Pin_AC pin annotation(
          Placement(transformation(extent = {{-10, -110}, {10, -90}})));
        Modelica.Blocks.Sources.RealExpression p_in(y = 0.1e6) "入口压力" annotation(
          Placement(transformation(extent = {{-10, 66}, {10, 86}})));
        Modelica.Blocks.Math.Add p_out "出口压力" annotation(
          Placement(transformation(extent = {{40, 60}, {60, 80}})));
      equation
        der(WP_out) = airSource_H.H_in / 1000 / 3600;
        der(Q_out) = electricalSource.P_in / 1000 / 3600;
        connect(pwrChar.y[1], electricalSource.P_in) annotation(
          Line(points = {{-19, -60}, {-2, -60}}, color = {0, 0, 127}));
        connect(air_in, dmSensor.In) annotation(
          Line(points = {{-100, 0}, {-90, 0}, {-90, 0}, {-80, 0}}, color = {0, 127, 255}));
        connect(dmSensor.Out, air_out) annotation(
          Line(points = {{-60, 0}, {16, 0}, {16, 0}, {100, 0}, {100, 0}}, color = {0, 0, 0}));
        connect(airSource_H.port_a, air_out) annotation(
          Line(points = {{60, -20}, {100, -20}, {100, 0}}, color = {0, 127, 255}));
        connect(dmSensor.m_flow, dV_dm.u) annotation(
          Line(points = {{-59, 5.2}, {-59, 20}, {-52, 20}}, color = {0, 0, 127}));
        connect(dV_dm.y, pwrChar.u[1]) annotation(
          Line(points = {{-29, 20}, {-28, 20}, {-28, -32}, {-60, -32}, {-60, -60}, {-42, -60}}, color = {0, 0, 127}));
        connect(product.y, airSource_H.H_in) annotation(
          Line(points = {{21, -20}, {38, -20}}, color = {0, 0, 127}));
        connect(pwrChar.y[1], product.u1) annotation(
          Line(points = {{-19, -60}, {-19, -26}, {-2, -26}}, color = {0, 0, 127}));
        connect(p_in.y, p_out.u1) annotation(
          Line(points = {{11, 76}, {38, 76}}, color = {0, 0, 127}));
        connect(dV_dm.y, etaChar.u) annotation(
          Line(points = {{-29, 20}, {-10, 20}}, color = {0, 0, 127}));
        connect(etaChar.y, product.u2) annotation(
          Line(points = {{13, 20}, {30, 20}, {30, -4}, {-20, -4}, {-20, -14}, {-2, -14}}, color = {0, 0, 127}));
        connect(dpChar.y, p_out.u2) annotation(
          Line(points = {{13, 50}, {26, 50}, {26, 64}, {38, 64}}, color = {0, 0, 127}));
        connect(electricalSource.pin, pin) annotation(
          Line(points = {{20, -60}, {34, -60}, {34, -100}, {0, -100}}, color = {0, 140, 72}));
        annotation(
          Icon(graphics = {Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Line(points = {{-44, 66}, {64, 46}}, color = {28, 108, 200}, smooth = Smooth.Bezier), Line(points = {{-44, -66}, {64, -46}}, color = {28, 108, 200}, smooth = Smooth.Bezier)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end EletricalCompressor;

      model DragCompressor "汽拖空压机"
        parameter Modelica.SIunits.Density rho_air = 1.29 "空气密度";
        Interfaces.Air.Port_a air_in annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Air.Port_b air_out annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        Modelica.Mechanics.Rotational.Interfaces.Flange_a flange_a annotation(
          Placement(transformation(extent = {{-10, 90}, {10, 110}})));
        annotation(
          Icon(graphics = {Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Line(points = {{-44, 66}, {64, 46}}, color = {28, 108, 200}, smooth = Smooth.Bezier), Line(points = {{-44, -66}, {64, -46}}, color = {28, 108, 200}, smooth = Smooth.Bezier)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end DragCompressor;
    end Compressor;

    package Turbine "汽轮机"
      model SteamTurbine "抽凝式汽轮机"
        parameter Modelica.SIunits.SpecificEnthalpy h_steam = 2573.68e3 "蒸汽焓值（J/kg）";
        parameter Modelica.SIunits.Temperature T_steam = 313.15 "蒸汽温度";
        parameter Modelica.SIunits.Pressure p_steam = 0.007e6 "蒸汽压力";
        Interfaces.Steam.Port_a steam_in "蒸汽输入接口" annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Steam.Port_b steam_out "蒸汽输出接口" annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        Sensors.MassFLowSensor massFLowSensor annotation(
          Placement(transformation(extent = {{-80, -10}, {-60, 10}})));
        Sinks.Sink sink annotation(
          Placement(transformation(extent = {{0, -10}, {-20, 10}})));
        Sources.SteamSource1 steamSource1_1 annotation(
          Placement(transformation(extent = {{60, 10}, {80, 30}})));
        Sources.SteamSource2 steamSource2_1 annotation(
          Placement(transformation(extent = {{60, -30}, {80, -10}})));
        Modelica.Blocks.Math.Gain h_Steam(k = h_steam) "蒸汽焓（J/kg）" annotation(
          Placement(transformation(extent = {{20, -30}, {40, -10}})));
        Modelica.SIunits.Power shaftPower "轴功率";
        Sensors.EnthalpyFLowSensor enthalpyFLowSensor annotation(
          Placement(transformation(extent = {{-50, -10}, {-30, 10}})));
        Modelica.Mechanics.Rotational.Sources.Speed speed annotation(
          Placement(transformation(extent = {{-40, 50}, {-20, 70}})));
        Modelica.Blocks.Sources.Constant const annotation(
          Placement(transformation(extent = {{-80, 50}, {-60, 70}})));
        Modelica.Mechanics.Rotational.Interfaces.Flange_b flange "Flange of shaft" annotation(
          Placement(transformation(extent = {{-10, 90}, {10, 110}})));
      equation
        flange.tau = shaftPower / speed.w;
        shaftPower = enthalpyFLowSensor.H_flow + steam_out.H_flow;
        connect(massFLowSensor.In, steam_in) annotation(
          Line(points = {{-80, 0}, {-100, 0}}, color = {0, 0, 0}));
        connect(massFLowSensor.m_flow, steamSource1_1.m_in) annotation(
          Line(points = {{-59, 5.2}, {-59, 20}, {58, 20}}, color = {0, 0, 127}));
        connect(h_Steam.y, steamSource2_1.H_in) annotation(
          Line(points = {{41, -20}, {58, -20}}, color = {0, 0, 127}));
        connect(massFLowSensor.m_flow, h_Steam.u) annotation(
          Line(points = {{-59, 5.2}, {-59, 20}, {10, 20}, {10, -20}, {18, -20}}, color = {0, 0, 127}));
        connect(steamSource2_1.port_a, steam_out) annotation(
          Line(points = {{80, -20}, {80, 0}, {100, 0}}, color = {255, 0, 0}));
        connect(const.y, speed.w_ref) annotation(
          Line(points = {{-59, 60}, {-42, 60}}, color = {0, 0, 127}));
        connect(speed.flange, flange) annotation(
          Line(points = {{-20, 60}, {0, 60}, {0, 100}}, color = {0, 0, 0}));
        connect(massFLowSensor.Out, enthalpyFLowSensor.In) annotation(
          Line(points = {{-60, 0}, {-50, 0}}));
        connect(enthalpyFLowSensor.Out, sink.generalFlowPort) annotation(
          Line(points = {{-30, 0}, {-20, 0}}));
        connect(steamSource1_1.port_a, steam_out) annotation(
          Line(points = {{80, 20}, {80, 0}, {100, 0}}, color = {255, 0, 0}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end SteamTurbine;

      model SteamTurbine_nomech
        parameter Modelica.SIunits.SpecificEnthalpy h_steam = 2573.68e3 "蒸汽焓值（J/kg）";
        Sensors.MassFLowSensor massFLowSensor annotation(
          Placement(transformation(extent = {{-88, -10}, {-68, 10}})));
        Sources.SteamSource2 steamSource2_1 annotation(
          Placement(transformation(extent = {{40, -40}, {60, -20}})));
        Modelica.Blocks.Math.Gain h_Steam(k = h_steam) "蒸汽焓（J/kg）" annotation(
          Placement(transformation(extent = {{10, -40}, {30, -20}})));
        Interfaces.Steam.Port_a steam_in "蒸汽输入接口" annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Steam.Port_b steam_out "蒸汽输出接口" annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        Sources.SteamSource1 steamSource1_1 annotation(
          Placement(transformation(extent = {{40, 20}, {60, 40}})));
        Sinks.Sink sink annotation(
          Placement(transformation(extent = {{-8, -10}, {-28, 10}})));
        Sensors.EnthalpyFLowSensor enthalpyFLowSensor annotation(
          Placement(transformation(extent = {{-60, -10}, {-40, 10}})));
      equation
        connect(massFLowSensor.In, steam_in) annotation(
          Line(points = {{-88, 0}, {-100, 0}}, color = {0, 0, 0}));
        connect(h_Steam.y, steamSource2_1.H_in) annotation(
          Line(points = {{31, -30}, {38, -30}}, color = {0, 0, 127}));
        connect(steamSource2_1.port_a, steam_out) annotation(
          Line(points = {{60, -30}, {60, 0}, {100, 0}}, color = {255, 0, 0}));
        connect(steamSource1_1.port_a, steam_out) annotation(
          Line(points = {{60, 30}, {60, 0}, {100, 0}}, color = {255, 0, 0}));
        connect(massFLowSensor.m_flow, steamSource1_1.m_in) annotation(
          Line(points = {{-67, 5.2}, {-67, 30}, {38, 30}}, color = {0, 0, 127}));
        connect(massFLowSensor.m_flow, h_Steam.u) annotation(
          Line(points = {{-67, 5.2}, {-67, 30}, {0, 30}, {0, -30}, {8, -30}}, color = {0, 0, 127}));
        connect(massFLowSensor.Out, enthalpyFLowSensor.In) annotation(
          Line(points = {{-68, 0}, {-60, 0}}));
        connect(enthalpyFLowSensor.Out, sink.generalFlowPort) annotation(
          Line(points = {{-40, 0}, {-28, 0}}, color = {0, 0, 0}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-148, 132}, {152, 92}}, textString = "%name", lineColor = {0, 0, 255})}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end SteamTurbine_nomech;
    end Turbine;

    package MicroTurbine "微燃机"
      model MicroTurbine
        Interfaces.Gas.Port_a gas_in "天然气输入接口" annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Air.Port_a air_in "空气输入接口" annotation(
          Placement(transformation(extent = {{-110, -80}, {-90, -60}})));
        Interfaces.Electrical.Pin_AC pin "电力输出接口" annotation(
          Placement(transformation(extent = {{-10, -110}, {10, -90}})));
        Interfaces.FlueGas.Port_b flue_out "烟气输出接口" annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end MicroTurbine;
    end MicroTurbine;

    package OxygenGenerator "制氧机"
      model OxygenGenerator "制氧机"
        parameter Modelica.SIunits.Density rho_air = 1.29 "空气密度";
        Interfaces.Electrical.Pin_AC pin "电力输入接口" annotation(
          Placement(transformation(extent = {{-10, -110}, {10, -90}})));
        Interfaces.Air.Port_a air_out "纯氧输出接口" annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Sensors.MassFLowSensor dmSensor annotation(
          Placement(transformation(extent = {{-80, -10}, {-60, 10}})));
        Modelica.Blocks.Tables.CombiTable1D pwrChar(table = [0.089, 28800; 0.177, 57600; 0.266, 86400; 0.354, 115200; 0.443, 144000; 0.886, 288000; 1.329, 432000; 1.771, 576000; 2.214, 720000; 2.657, 864000; 3.1, 1008000; 3.543, 1152000; 3.986, 1296000]) "制氧机功率曲线，流量（m3/s）vs 功率（W）" annotation(
          Placement(transformation(extent = {{-12, 20}, {8, 40}})));
        Sources.ElectricalSource electricalSource annotation(
          Placement(transformation(extent = {{28, 20}, {48, 40}})));
        Modelica.Blocks.Math.Gain dV_dm(k = -1 / rho_air) "默认空气密度=1.29kg/m3" annotation(
          Placement(transformation(extent = {{-48, 20}, {-28, 40}})));
        Sinks.Sink sinkNone annotation(
          Placement(transformation(extent = {{-20, -10}, {-40, 10}})));
      equation
        connect(air_out, dmSensor.In) annotation(
          Line(points = {{-100, 0}, {-90, 0}, {-90, 0}, {-80, 0}}, color = {0, 127, 255}));
        connect(pwrChar.y[1], electricalSource.P_in) annotation(
          Line(points = {{9, 30}, {26, 30}}, color = {0, 0, 127}));
        connect(electricalSource.pin, pin) annotation(
          Line(points = {{48, 30}, {58, 30}, {58, -100}, {0, -100}}, color = {0, 0, 255}));
        connect(dV_dm.y, pwrChar.u[1]) annotation(
          Line(points = {{-27, 30}, {-14, 30}}, color = {0, 0, 127}));
        connect(dV_dm.u, dmSensor.m_flow) annotation(
          Line(points = {{-50, 30}, {-59, 30}, {-59, 5.2}}, color = {0, 0, 127}));
        connect(dmSensor.Out, sinkNone.generalFlowPort) annotation(
          Line(points = {{-60, 0}, {-40, 0}}, color = {0, 0, 0}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end OxygenGenerator;
    end OxygenGenerator;

    package Fan "风机"
      model TestFan
        extends Modelica.Icons.Example;
        Fan fan annotation(
          Placement(transformation(extent = {{-10, -10}, {10, 10}})));
        parameter Modelica.Thermal.FluidHeatFlow.Media.Medium medium = Modelica.Thermal.FluidHeatFlow.Media.Medium() "Medium in the component";
        parameter Modelica.SIunits.Temperature Tamb = 293.15 "Initial temperature of medium";
        Sinks.ElectricalSink electricalSink annotation(
          Placement(transformation(extent = {{-30, -40}, {-10, -20}})));
        Sources.AirSource1 airSource1_1 annotation(
          Placement(transformation(extent = {{-50, -10}, {-30, 10}})));
        Modelica.Blocks.Sources.RealExpression realExpression(y = 169.9) annotation(
          Placement(transformation(extent = {{-90, -10}, {-70, 10}})));
        Sinks.Sink sinkNone annotation(
          Placement(transformation(extent = {{82, -10}, {62, 10}})));
        Sensors.EnthalpyFLowSensor enthalpyFLowSensor annotation(
          Placement(transformation(extent = {{24, -10}, {44, 10}})));
      equation
        connect(electricalSink.pin, fan.pin) annotation(
          Line(points = {{-10, -30}, {0, -30}, {0, -10}}, color = {0, 0, 255}));
        connect(airSource1_1.port_a, fan.air_in) annotation(
          Line(points = {{-30, 0}, {-10, 0}}, color = {0, 127, 255}));
        connect(realExpression.y, airSource1_1.m_in) annotation(
          Line(points = {{-69, 0}, {-52, 0}}, color = {0, 0, 127}));
        connect(fan.air_out, enthalpyFLowSensor.In) annotation(
          Line(points = {{10, 0}, {24, 0}}, color = {0, 127, 255}));
        connect(enthalpyFLowSensor.Out, sinkNone.generalFlowPort) annotation(
          Line(points = {{44, 0}, {62, 0}}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false)),
          Diagram(coordinateSystem(preserveAspectRatio = false)),
          experiment(StopTime = 10));
      end TestFan;

      model Fan "风机"
        parameter Modelica.SIunits.Density rho_air(displayUnit = "kg/m3") = 1.29 "空气密度";
        Interfaces.Air.Port_a air_in annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Air.Port_b air_out annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        Interfaces.Electrical.Pin_AC pin annotation(
          Placement(transformation(extent = {{-10, -110}, {10, -90}})));
        Sources.ElectricalSource electricalSource annotation(
          Placement(transformation(extent = {{0, -70}, {20, -50}})));
        Sensors.MassFLowSensor dmSensor annotation(
          Placement(transformation(extent = {{-80, -10}, {-60, 10}})));
        Sources.AirSource2 airSource2_1 annotation(
          Placement(transformation(extent = {{40, -30}, {60, -10}})));
        Modelica.Blocks.Math.Gain dV_dm(k = 1 / rho_air) "默认空气密度=1.29kg/m3" annotation(
          Placement(transformation(extent = {{-50, 10}, {-30, 30}})));
        Modelica.Blocks.Sources.Constant dpChar(k = dp) "空压机压差曲线 dp（Pa）" annotation(
          Placement(transformation(extent = {{-8, 40}, {12, 60}})));
        Modelica.Blocks.Sources.RealExpression p_in(y = 0.1e6) "入口压力（Pa）" annotation(
          Placement(transformation(extent = {{-10, 66}, {10, 86}})));
        Modelica.Blocks.Math.Add p_out "出口压力（Mpa）" annotation(
          Placement(transformation(extent = {{40, 60}, {60, 80}})));
        Modelica.Blocks.Math.Gain etaChar(k = 1 / eta) "风机效率" annotation(
          Placement(transformation(extent = {{-42, -70}, {-22, -50}})));
        Modelica.Blocks.Math.MultiProduct multiProduct(nu = 3) annotation(
          Placement(transformation(extent = {{6, -26}, {18, -14}})));
        Modelica.Blocks.Sources.RealExpression const(y = 2.73 * 3600 / 1000) annotation(
          Placement(transformation(extent = {{-60, -28}, {-20, -8}})));
        parameter Real eta = 0.7 "风机效率";
        parameter Modelica.SIunits.PressureDifference dp = 300000 "风机扬程";
      equation
        connect(air_in, dmSensor.In) annotation(
          Line(points = {{-100, 0}, {-90, 0}, {-90, 0}, {-80, 0}}, color = {0, 127, 255}));
        connect(dmSensor.Out, air_out) annotation(
          Line(points = {{-60, 0}, {16, 0}, {16, 0}, {100, 0}, {100, 0}}, color = {0, 0, 0}));
        connect(airSource2_1.port_a, air_out) annotation(
          Line(points = {{60, -20}, {100, -20}, {100, 0}}, color = {0, 127, 255}));
        connect(dmSensor.m_flow, dV_dm.u) annotation(
          Line(points = {{-59, 5.2}, {-59, 20}, {-52, 20}}, color = {0, 0, 127}));
        connect(p_in.y, p_out.u1) annotation(
          Line(points = {{11, 76}, {38, 76}}, color = {0, 0, 127}));
        connect(dpChar.y, p_out.u2) annotation(
          Line(points = {{13, 50}, {26, 50}, {26, 64}, {38, 64}}, color = {0, 0, 127}));
        connect(const.y, multiProduct.u[1]) annotation(
          Line(points = {{-18, -18}, {-16, -18}, {-16, -17.2}, {6, -17.2}}, color = {0, 0, 127}));
        connect(dV_dm.y, multiProduct.u[2]) annotation(
          Line(points = {{-29, 20}, {-10, 20}, {-10, -20}, {6, -20}}, color = {0, 0, 127}));
        connect(dpChar.y, multiProduct.u[3]) annotation(
          Line(points = {{13, 50}, {20, 50}, {20, 20}, {6, 20}, {6, -22.8}}, color = {0, 0, 127}));
        connect(multiProduct.y, airSource2_1.H_in) annotation(
          Line(points = {{19.02, -20}, {38, -20}}, color = {0, 0, 127}));
        connect(multiProduct.y, etaChar.u) annotation(
          Line(points = {{19.02, -20}, {24, -20}, {24, -40}, {-70, -40}, {-70, -60}, {-44, -60}, {-44, -60}}, color = {0, 0, 127}));
        connect(etaChar.y, electricalSource.P_in) annotation(
          Line(points = {{-21, -60}, {-2, -60}}, color = {0, 0, 127}));
        connect(electricalSource.pin, pin) annotation(
          Line(points = {{20, -60}, {28, -60}, {28, -100}, {0, -100}}, color = {0, 140, 72}));
        annotation(
          Icon(graphics = {Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {0, 83, 134}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Line(points = {{-44, 66}, {64, 46}}, color = {0, 83, 134}, smooth = Smooth.Bezier), Line(points = {{410, -10}}, color = {135, 135, 135}), Line(points = {{238, 158}}, color = {135, 135, 135}), Polygon(points = {{92, 228}, {92, 228}}, lineColor = {135, 135, 135}), Line(points = {{320, 80}}, color = {135, 135, 135}), Polygon(points = {{-4, 0}, {-14, 46}, {0, 56}, {14, 46}, {4, 0}, {-4, 0}}, lineColor = {0, 83, 134}, origin = {0, 0}, rotation = 0, smooth = Smooth.Bezier, fillColor = {0, 83, 134}, fillPattern = FillPattern.Solid), Polygon(points = {{-4, 0}, {-14, 46}, {0, 56}, {14, 46}, {4, 0}, {-4, 0}}, lineColor = {0, 83, 134}, origin = {0, 0}, rotation = 120, smooth = Smooth.Bezier, fillColor = {0, 83, 134}, fillPattern = FillPattern.Solid), Polygon(points = {{-4, 0}, {-14, 46}, {0, 56}, {14, 46}, {4, 0}, {-4, 0}}, lineColor = {0, 83, 134}, origin = {0, 0}, rotation = 240, smooth = Smooth.Bezier, fillColor = {0, 83, 134}, fillPattern = FillPattern.Solid), Line(points = {{-132, 72}}, color = {135, 135, 135}, smooth = Smooth.Bezier), Line(points = {{-124, 56}}, color = {135, 135, 135}, smooth = Smooth.Bezier), Line(points = {{-44, -66}, {64, -46}}, color = {0, 83, 134}, smooth = Smooth.Bezier)}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end Fan;
    end Fan;

    package GearBox "齿轮箱"
      model IdealGear "Ideal gear without inertia"
        extends Modelica.Mechanics.Rotational.Icons.Gear;
        extends Modelica.Mechanics.Rotational.Interfaces.PartialElementaryTwoFlangesAndSupport2;
        parameter Real ratio(start = 1) = 0.9 "Transmission ratio (flange_a.phi/flange_b.phi)";
        Modelica.SIunits.Angle phi_a "Angle between left shaft flange and support";
        Modelica.SIunits.Angle phi_b "Angle between right shaft flange and support";
      equation
        phi_a = flange_a.phi - phi_support;
        phi_b = flange_b.phi - phi_support;
        phi_a = ratio * phi_b;
        0 = ratio * flange_a.tau + flange_b.tau;
        annotation(
          Documentation(info = "<html>
<p>
This element characterizes any type of gear box which is fixed in the
ground and which has one driving shaft and one driven shaft.
The gear is <strong>ideal</strong>, i.e., it does not have inertia, elasticity, damping
or backlash. If these effects have to be considered, the gear has to be
connected to other elements in an appropriate way.
</p>

</html>"),
          Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Text(extent = {{-153, 145}, {147, 105}}, lineColor = {0, 0, 255}, textString = "%name"), Text(extent = {{-146, -49}, {154, -79}}, textString = "ratio=%ratio")}));
      end IdealGear;
    end GearBox;

    package Pump
      model TestPump
        extends Modelica.Icons.Example;
        parameter Media.Medium medium = ENN.Media.Water() "Cooling medium" annotation(
          choicesAllMatching = true);
        parameter Modelica.SIunits.Temperature TAmb(displayUnit = "degC") = 293.15 "Ambient temperature";
        Modelica.Thermal.FluidHeatFlow.Sources.Ambient ambient(constantAmbientTemperature = TAmb, medium = medium, constantAmbientPressure(displayUnit = "MPa") = 100000) annotation(
          Placement(transformation(extent = {{-40, -40}, {-60, -20}})));
        Modelica.Thermal.FluidHeatFlow.Sources.Ambient ambient1(constantAmbientTemperature = TAmb, medium = medium, constantAmbientPressure(displayUnit = "MPa") = 1000000) annotation(
          Placement(transformation(extent = {{72, -40}, {92, -20}})));
        Pump pump_VF(medium = medium) annotation(
          Placement(transformation(extent = {{-20, -40}, {0, -20}})));
        Modelica.Blocks.Sources.Constant const(k = 48) annotation(
          Placement(transformation(extent = {{-60, 0}, {-40, 20}})));
        Pipes.Pipe pipe(medium = medium, m = 0.1, T0 = TAmb, h_g = 0) annotation(
          Placement(transformation(extent = {{20, -40}, {40, -20}})));
        Sinks.ElectricalSink electricalSink annotation(
          Placement(transformation(extent = {{-60, -80}, {-40, -60}})));
      equation
        connect(const.y, pump_VF.f) annotation(
          Line(points = {{-39, 10}, {-10, 10}, {-10, -18}}, color = {0, 0, 127}));
        connect(pump_VF.port_a, ambient.flowPort) annotation(
          Line(points = {{-20, -30}, {-40, -30}}, color = {85, 255, 85}));
        connect(electricalSink.pin, pump_VF.pin) annotation(
          Line(points = {{-40, -70}, {-10, -70}, {-10, -40}}, color = {0, 0, 255}));
        connect(pump_VF.port_b, pipe.port_a) annotation(
          Line(points = {{0, -30}, {20, -30}}, color = {85, 255, 85}));
        connect(pipe.port_b, ambient1.flowPort) annotation(
          Line(points = {{40, -30}, {72, -30}}, color = {85, 255, 85}));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false)),
          Diagram(coordinateSystem(preserveAspectRatio = false)),
          experiment(StopTime = 50));
      end TestPump;

      package Base
        partial model PumpBase "Partial model of two port"
          parameter Modelica.Thermal.FluidHeatFlow.Media.Medium medium = Modelica.Thermal.FluidHeatFlow.Media.Medium() "Medium in the component" annotation(
            choicesAllMatching = true);
          parameter Modelica.SIunits.Mass m(start = 1) "Mass of medium";
          parameter Modelica.SIunits.Temperature T0(start = 293.15, displayUnit = "degC") "Initial temperature of medium" annotation(
            Dialog(enable = m > Modelica.Constants.small));
          parameter Boolean T0fixed = false "Initial temperature guess value or fixed" annotation(
            choices(checkBox = true),
            Dialog(enable = m > Modelica.Constants.small));
          parameter Real tapT(final min = 0, final max = 1) = 1 "Defines temperature of heatPort between inlet and outlet temperature";
          Modelica.SIunits.Pressure dp "Pressure drop a->b";
          Modelica.SIunits.VolumeFlowRate V_flow(start = 0) "Volume flow a->b";
          Modelica.SIunits.MassFlowRate m_flow "Mass flow rate";
          Modelica.SIunits.HeatFlowRate Q_flow "Heat exchange with ambient";
          output Modelica.SIunits.Temperature T(start = T0, fixed = T0fixed) "Outlet temperature of medium";
          output Modelica.SIunits.Temperature T_a "Temperature at flowPort_a";
          output Modelica.SIunits.Temperature T_b "Temperature at flowPort_b";
          output Modelica.SIunits.TemperatureDifference dT "Temperature increase of coolant in flow direction";
          Modelica.SIunits.Temperature T_q "Temperature relevant for heat exchange with ambient";
          Modelica.SIunits.Height head "Pump head";
          parameter Real eta = 1 "Efficiency";
          Modelica.SIunits.Power P_tot "Total power consumption";
          //Constants
          constant Modelica.SIunits.Acceleration g = Modelica.Constants.g_n "Gravitational acceleration";
          Modelica.SIunits.SpecificEnthalpy h "Medium's specific enthalpy";
        public
          Interfaces.Water_media.Port_a port_a(final medium = medium) annotation(
            Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
          Interfaces.Water_media.Port_b port_b(final medium = medium) annotation(
            Placement(transformation(extent = {{90, -10}, {110, 10}})));
        equation
          dp = port_b.p - port_a.p;
          V_flow = port_a.m_flow / medium.rho;
          T_a = port_a.h / medium.cp;
          T_b = port_b.h / medium.cp;
          dT = if noEvent(V_flow >= 0) then T - T_a else T_b - T;
          h = medium.cp * T;
          T_q = T - noEvent(sign(V_flow)) * (1 - tapT) * dT;
// mass balance
          port_a.m_flow + port_b.m_flow = 0;
// energy balance
//    if m>Modelica.Constants.small then
//     port_a.H_flow + port_b.H_flow + Q_flow = m*medium.cv*der(T);
//   else
//     port_a.H_flow + port_b.H_flow + Q_flow = 0;
//   end if;
//Fluid
          dp = head * medium.rho * g;
          P_tot = dp * V_flow * eta;
// mass flow a->b mixing rule at a, energy flow at b defined by medium's temperature
// mass flow b->a mixing rule at b, energy flow at a defined by medium's temperature
          port_a.H_flow = semiLinear(port_a.m_flow, port_a.h, h);
          port_b.H_flow = semiLinear(port_b.m_flow, port_b.h, h);
          annotation(
            Documentation(info = "<html>
<p>Partial model with two flowPorts.</p>
<p>Possible heat exchange with the ambient is defined by Q_flow; setting this = 0 means no energy exchange.</p>
<p>
Setting parameter m (mass of medium within pipe) to zero
leads to neglect of temperature transient cv*m*der(T).</p>
<p>Mixing rule is applied.</p>
<p>Parameter 0 &lt; tapT &lt; 1 defines temperature of heatPort between medium's inlet and outlet temperature.</p>
</html>"));
        end PumpBase;

        model IdealPump "Model of an ideal pump"
          extends PumpBase(final tapT = 1);
          parameter Modelica.SIunits.AngularVelocity wNominal(start = 1, displayUnit = "rev/min") "Nominal speed" annotation(
            Dialog(group = "Pump characteristic"));
          parameter Modelica.SIunits.Pressure dp0(start = 2) "Max. pressure increase @ V_flow=0" annotation(
            Dialog(group = "Pump characteristic"));
          parameter Modelica.SIunits.VolumeFlowRate V_flow0(start = 2) "Max. volume flow rate @ dp=0" annotation(
            Dialog(group = "Pump characteristic"));
          Modelica.SIunits.AngularVelocity w = der(flange_a.phi) "Speed";
        protected
          Modelica.SIunits.Pressure dp1;
          Modelica.SIunits.VolumeFlowRate V_flow1;
        public
          Modelica.Mechanics.Rotational.Interfaces.Flange_a flange_a annotation(
            Placement(transformation(extent = {{-10, -110}, {10, -90}})));
        equation
// pump characteristic
          dp1 = dp0 * sign(w / wNominal) * (w / wNominal) ^ 2;
          V_flow1 = V_flow0 * (w / wNominal);
          if noEvent(abs(w) < Modelica.Constants.small) then
            dp = 0;
            flange_a.tau = 0;
          else
            dp = -dp1;
            flange_a.tau * w = -dp * V_flow;
          end if;
// no energy exchange with medium
          Q_flow = 0;
          annotation(
            Documentation(info = "<html>
<p>
Simple fan resp. pump where characteristic is dependent on shaft's speed, <br>
torque * speed = pressure increase * volume flow (without losses)<br>
Pressure increase versus volume flow is defined by a linear function,
from dp0(V_flow=0) to V_flow0(dp=0).<br>
The axis intersections vary with speed as follows:
</p>
<ul>
<li>dp prop. speed^2</li>
<li>V_flow prop. speed</li>
</ul>
<p>
Coolant's temperature and enthalpy flow are not affected.<br>
Setting parameter m (mass of medium within fan/pump) to zero
leads to neglection of temperature transient cv*m*der(T).<br>
Thermodynamic equations are defined by Partials.TwoPort.
</p>
</html>"),
            Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 150}, {150, 90}}, lineColor = {0, 0, 255}, textString = "%name"), Rectangle(extent = {{-10, -40}, {10, -100}}, fillPattern = FillPattern.VerticalCylinder, fillColor = {175, 175, 175}), Polygon(points = {{-60, 68}, {90, 10}, {90, -10}, {-60, -68}, {-60, 68}}, fillPattern = FillPattern.HorizontalCylinder, fillColor = {0, 0, 255})}));
        end IdealPump;
      end Base;

      model Pump "水泵模型"
        parameter Modelica.SIunits.Temperature TAmb(displayUnit = "degC") = 293.15 "Ambient temperature";
        //能流计算需求
        Real WP_out(unit = "kWh") "累计势能";
        Real Q_out(unit = "kWh") "累计耗电";
        Modelica.Mechanics.Rotational.Sources.Speed speed annotation(
          Placement(transformation(extent = {{-40, 50}, {-20, 70}})));
        Modelica.Blocks.Interfaces.RealInput f "Reference angular velocity of flange with respect to support as input signal" annotation(
          Placement(transformation(extent = {{-140, 40}, {-100, 80}}), iconTransformation(extent = {{-20, -20}, {20, 20}}, rotation = -90, origin = {0, 120})));
        Modelica.Blocks.Math.Gain gain(k = 2 * 3.14) annotation(
          Placement(transformation(extent = {{-80, 50}, {-60, 70}})));
        Interfaces.Water_media.Port_a port_a(medium = medium) annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Water_media.Port_b port_b(medium = medium) annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        parameter Media.Medium medium = Media.Medium() "Medium in the component";
        replaceable Modelica.Thermal.FluidHeatFlow.Sources.IdealPump idealPump(medium = medium, m = 0, T0 = TAmb, wNominal(displayUnit = "rpm") = 308.923277603, dp0(displayUnit = "MPa") = 1020000, V_flow0 = 0.0075) constrainedby Base.IdealPump(medium = medium, m = 0, T0 = TAmb, wNominal(displayUnit = "rpm") = 308.923277603, dp0(displayUnit = "MPa") = 1020000, V_flow0 = 0.002) annotation(
           Placement(transformation(extent = {{-10, 10}, {10, -10}})));
        constant Modelica.SIunits.Acceleration g = Modelica.Constants.g_n;
        parameter Real eta = 0.33 "efficient ratio";
        Modelica.SIunits.Height head;
        Modelica.SIunits.Power pwr0;
        Modelica.Thermal.FluidHeatFlow.Sensors.PressureSensor pressureSensor(medium = medium) annotation(
          Placement(transformation(extent = {{10, -10}, {-10, 10}}, rotation = -90, origin = {60, 18})));
        Modelica.Blocks.Sources.RealExpression realExpression(y = pwr0) annotation(
          Placement(transformation(extent = {{-60, -50}, {-40, -30}})));
        Sources.ElectricalSource electricalSource annotation(
          Placement(transformation(extent = {{-20, -50}, {0, -30}})));
        Interfaces.Electrical.Pin_AC pin annotation(
          Placement(transformation(extent = {{-10, -110}, {10, -90}})));
      equation
        -idealPump.dp = head * medium.rho * g;
        pwr0 = 2.73 * port_a.m_flow * head / 1000 / eta * 3600;
        der(WP_out) = abs(port_b.H_flow) / 1000 / 3600;
        der(Q_out) = electricalSource.P_in / 1000 / 3600;
        connect(gain.u, f) annotation(
          Line(points = {{-82, 60}, {-120, 60}}, color = {0, 0, 127}));
        connect(gain.y, speed.w_ref) annotation(
          Line(points = {{-59, 60}, {-42, 60}}, color = {0, 0, 127}));
        connect(speed.flange, idealPump.flange_a) annotation(
          Line(points = {{-20, 60}, {0, 60}, {0, 10}}, color = {0, 0, 0}));
        connect(port_b, pressureSensor.flowPort) annotation(
          Line(points = {{100, 0}, {60, 0}, {60, 8}}, color = {85, 255, 85}));
        connect(port_a, idealPump.flowPort_a) annotation(
          Line(points = {{-100, 0}, {-10, 0}}, color = {85, 255, 85}));
        connect(port_b, idealPump.flowPort_b) annotation(
          Line(points = {{100, 0}, {10, 0}}, color = {85, 255, 85}));
        connect(realExpression.y, electricalSource.P_in) annotation(
          Line(points = {{-39, -40}, {-22, -40}}, color = {0, 0, 127}));
        connect(electricalSource.pin, pin) annotation(
          Line(points = {{0, -40}, {0, -100}}, color = {0, 0, 255}));
        annotation(
          experiment(StopTime = 30),
          Icon(graphics = {Polygon(points = {{20, -70}, {60, -85}, {20, -100}, {20, -70}}, lineColor = {0, 128, 255}, fillColor = {0, 128, 255}, fillPattern = FillPattern.Solid, visible = showDesignFlowDirection), Polygon(points = {{20, -75}, {50, -85}, {20, -95}, {20, -75}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, visible = allowFlowReversal), Line(points = {{55, -85}, {-60, -85}}, color = {0, 128, 255}, visible = showDesignFlowDirection), Text(extent = {{-149, -114}, {151, -154}}, lineColor = {0, 0, 255}, textString = "%name"), Rectangle(extent = {{-100, 46}, {100, -46}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.HorizontalCylinder), Polygon(points = {{-48, -60}, {-72, -100}, {72, -100}, {48, -60}, {-48, -60}}, lineColor = {0, 0, 255}, pattern = LinePattern.None, fillPattern = FillPattern.VerticalCylinder), Ellipse(extent = {{-80, 80}, {80, -80}}, fillPattern = FillPattern.Sphere, fillColor = {0, 100, 199}), Polygon(points = {{-28, 30}, {-28, -30}, {50, -2}, {-28, 30}}, pattern = LinePattern.None, fillPattern = FillPattern.HorizontalCylinder, fillColor = {255, 255, 255}), Rectangle(extent = {{-10, 100}, {10, 78}}, fillPattern = FillPattern.VerticalCylinder, fillColor = {95, 95, 95})}));
      end Pump;
    end Pump;

    package CoolingTower
      model CoolingTower "冷却塔"
        Interfaces.Air.Port_a port_a annotation(
          Placement(transformation(extent = {{-110, -70}, {-90, -50}})));
        Interfaces.Air.Port_b port_b annotation(
          Placement(transformation(extent = {{90, -70}, {110, -50}})));
        Interfaces.HotWater.Port_a port_a1 annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Interfaces.Water_media.Port_b port_b1 annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false)),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end CoolingTower;
    end CoolingTower;

    package HeatPump "热泵"
      model HeatPump "电驱离心式热泵"
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false)),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end HeatPump;
    end HeatPump;

    package Condenser "凝汽器"
      model Condenser "凝汽器"
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false)),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end Condenser;
    end Condenser;
  end Machines;

  package RenewableEnergy "可再生能源"
    model ElectricalNet_withGrid
      extends Modelica.Icons.Example;
      Components.Internal.ACbus aCbusBar annotation(
        Placement(transformation(extent = {{-30, 0}, {-10, 20}})));
      Electrical.ElectricGrid electricGrid(V_ref(displayUnit = "kV")) annotation(
        Placement(transformation(extent = {{-100, 0}, {-80, 20}})));
      Electrical.Transformer transformer(V_ref = 380) annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {0, 10})));
      inner Environment environment annotation(
        Placement(transformation(extent = {{-100, 80}, {-80, 100}})));
      Modelica.Electrical.Analog.Ideal.IdealOpeningSwitch switch annotation(
        Placement(transformation(extent = {{-60, 20}, {-40, 0}})));
      //Modelica.SIunits.Power pAll= windPower.P_wind+electricGrid.P_grid+solarPower.P_solar;
      Electrical.GridControl gridControl(P_renew = 0, P_load = transformer.power_req) annotation(
        Placement(transformation(extent = {{-92, -30}, {-72, -10}})));
      Sources.ElectricalSource electricalSource annotation(
        Placement(transformation(extent = {{60, 0}, {40, 20}})));
      Modelica.Blocks.Sources.Constant pwr1(k = 100e3) annotation(
        Placement(transformation(extent = {{100, 0}, {80, 20}})));
      Components.Internal.ACbus aCbusBar1 annotation(
        Placement(transformation(extent = {{10, 0}, {30, 20}})));
      Sources.ElectricalSource electricalSource1 annotation(
        Placement(transformation(extent = {{60, -40}, {40, -20}})));
      Modelica.Blocks.Sources.Constant pwr2(k = 50e3) annotation(
        Placement(transformation(extent = {{100, -40}, {80, -20}})));
    equation
      connect(electricGrid.p, switch.p) annotation(
        Line(points = {{-80, 10}, {-60, 10}}, color = {0, 140, 72}));
      connect(gridControl.y, switch.control) annotation(
        Line(points = {{-71, -20}, {-50, -20}, {-50, -2}}, color = {255, 0, 255}));
      connect(switch.n, aCbusBar.term) annotation(
        Line(points = {{-40, 10}, {-20, 10}}, color = {0, 0, 255}));
      connect(aCbusBar.term, transformer.p) annotation(
        Line(points = {{-20, 10}, {-9.8, 10}}, color = {0, 140, 72}));
      connect(electricalSource.P_in, pwr1.y) annotation(
        Line(points = {{62, 10}, {79, 10}}, color = {0, 0, 127}));
      connect(aCbusBar1.term, transformer.n) annotation(
        Line(points = {{20, 10}, {10, 10}}, color = {0, 140, 72}));
      connect(aCbusBar1.term, electricalSource.pin) annotation(
        Line(points = {{20, 10}, {40, 10}}, color = {0, 140, 72}));
      connect(electricalSource1.P_in, pwr2.y) annotation(
        Line(points = {{62, -30}, {79, -30}}, color = {0, 0, 127}));
      connect(electricalSource1.pin, aCbusBar1.term) annotation(
        Line(points = {{40, -30}, {30, -30}, {30, 10}, {20, 10}}, color = {0, 140, 72}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false)),
        Diagram(coordinateSystem(preserveAspectRatio = false)),
        experiment(StopTime = 86400, __Dymola_Algorithm = "Dassl"));
    end ElectricalNet_withGrid;

    model ElectricalNet
      extends Modelica.Icons.Example;
      Components.SolarPower solarPower annotation(
        Placement(transformation(extent = {{-80, 40}, {-60, 60}})));
      Components.WindPower windPower annotation(
        Placement(transformation(extent = {{-80, 0}, {-60, 20}})));
      Electrical.Inverter inverter(V_ref = 100) annotation(
        Placement(transformation(extent = {{-40, 40}, {-20, 60}})));
      Components.Internal.ACbus aCbusBar annotation(
        Placement(transformation(extent = {{0, 0}, {20, 20}})));
      Electrical.ElectricGrid electricGrid(V_ref(displayUnit = "kV")) annotation(
        Placement(transformation(extent = {{-80, -40}, {-60, -20}})));
      Electrical.Transformer transformer(V_ref = 380) annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {30, 10})));
      Electrical.Load load(P_load(displayUnit = "kW") = 1000000) annotation(
        Placement(transformation(extent = {{60, 0}, {80, 20}})));
      inner Environment environment annotation(
        Placement(transformation(extent = {{-100, 80}, {-80, 100}})));
      Modelica.Electrical.Analog.Ideal.IdealOpeningSwitch switch annotation(
        Placement(transformation(extent = {{-46, -20}, {-26, -40}})));
      Modelica.SIunits.Power pAll = windPower.P_wind + electricGrid.P_grid + solarPower.P_solar;
      Electrical.GridControl gridControl(P_renew = solarPower.P_solar + windPower.P_wind, P_load = transformer.power_req) annotation(
        Placement(transformation(extent = {{-80, -80}, {-60, -60}})));
    equation
      connect(solarPower.p, inverter.n) annotation(
        Line(points = {{-60, 50}, {-38, 50}}, color = {0, 0, 255}));
      connect(electricGrid.p, switch.p) annotation(
        Line(points = {{-60, -30}, {-46, -30}}, color = {0, 140, 72}));
      connect(gridControl.y, switch.control) annotation(
        Line(points = {{-59, -70}, {-36, -70}, {-36, -42}}, color = {255, 0, 255}));
      connect(inverter.p, aCbusBar.term) annotation(
        Line(points = {{-22, 50}, {-14, 50}, {-14, 10}, {10, 10}}, color = {0, 140, 72}));
      connect(windPower.p, aCbusBar.term) annotation(
        Line(points = {{-60, 10}, {10, 10}}, color = {0, 140, 72}));
      connect(switch.n, aCbusBar.term) annotation(
        Line(points = {{-26, -30}, {-14, -30}, {-14, 10}, {10, 10}}, color = {0, 0, 255}));
      connect(aCbusBar.term, transformer.p) annotation(
        Line(points = {{10, 10}, {16, 10}, {16, 10}, {20.2, 10}}, color = {0, 140, 72}));
      connect(transformer.n, load.n) annotation(
        Line(points = {{40, 10}, {60, 10}}, color = {0, 140, 72}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false)),
        Diagram(coordinateSystem(preserveAspectRatio = false)),
        experiment(StopTime = 86400, __Dymola_Algorithm = "Dassl"));
    end ElectricalNet;

    model ElectricalNet_1
      extends Modelica.Icons.Example;
      Components.SolarPower solarPower annotation(
        Placement(transformation(extent = {{-80, 40}, {-60, 60}})));
      Components.WindPower windPower annotation(
        Placement(transformation(extent = {{-80, 0}, {-60, 20}})));
      Electrical.Inverter inverter(V_ref = 380) annotation(
        Placement(transformation(extent = {{-40, 40}, {-20, 60}})));
      Components.Internal.ACbus aCbusBar annotation(
        Placement(transformation(extent = {{0, 0}, {20, 20}})));
      Electrical.ElectricGrid electricGrid(V_ref(displayUnit = "kV")) annotation(
        Placement(transformation(extent = {{-80, -40}, {-60, -20}})));
      Electrical.Transformer transformer(V_ref = 380) annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}}, rotation = 90, origin = {30, 10})));
      Electrical.Load load(P_load(displayUnit = "kW") = 1000000) annotation(
        Placement(transformation(extent = {{60, 0}, {80, 20}})));
      inner Environment environment annotation(
        Placement(transformation(extent = {{-100, 80}, {-80, 100}})));
      Modelica.Electrical.Analog.Ideal.IdealOpeningSwitch switch annotation(
        Placement(transformation(extent = {{-46, -20}, {-26, -40}})));
      Modelica.SIunits.Power pAll = windPower.P_wind + electricGrid.P_grid + solarPower.P_solar;
      Electrical.GridControl gridControl(P_renew = solarPower.P_solar + windPower.P_wind, P_load = transformer.power_req) annotation(
        Placement(transformation(extent = {{-80, -80}, {-60, -60}})));
      Components.Internal.ACbus aCbusBar1 annotation(
        Placement(transformation(extent = {{40, 0}, {60, 20}})));
    equation
      connect(solarPower.p, inverter.n) annotation(
        Line(points = {{-60, 50}, {-38, 50}}, color = {0, 0, 255}));
      connect(electricGrid.p, switch.p) annotation(
        Line(points = {{-60, -30}, {-46, -30}}, color = {0, 140, 72}));
      connect(gridControl.y, switch.control) annotation(
        Line(points = {{-59, -70}, {-36, -70}, {-36, -42}}, color = {255, 0, 255}));
      connect(aCbusBar.term, transformer.p) annotation(
        Line(points = {{10, 10}, {16, 10}, {16, 10}, {20.2, 10}}, color = {0, 140, 72}));
      connect(windPower.p, aCbusBar.term) annotation(
        Line(points = {{-60, 10}, {10, 10}}, color = {0, 140, 72}));
      connect(switch.n, aCbusBar.term) annotation(
        Line(points = {{-26, -30}, {-10, -30}, {-10, 10}, {10, 10}}, color = {0, 0, 255}));
      connect(inverter.p, aCbusBar1.term) annotation(
        Line(points = {{-22, 50}, {46, 50}, {46, 10}, {50, 10}}, color = {0, 140, 72}));
      connect(transformer.n, aCbusBar1.term) annotation(
        Line(points = {{40, 10}, {50, 10}}, color = {0, 140, 72}));
      connect(load.n, aCbusBar1.term) annotation(
        Line(points = {{60, 10}, {50, 10}}, color = {0, 140, 72}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false)),
        Diagram(coordinateSystem(preserveAspectRatio = false)),
        experiment(StopTime = 86400, __Dymola_Algorithm = "Dassl"));
    end ElectricalNet_1;

    package Components
      model SolarPower "光伏发电"
        Modelica.Electrical.Analog.Interfaces.PositivePin p annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
        parameter Modelica.SIunits.Power P_rat = 300000 "额定功率";
        parameter Real eta_PV = 0.935 "光伏降额因子，考虑老化、灰尘等影响因素";
        Modelica.SIunits.Power P_solar "输出的发电功率";
        Modelica.SIunits.Power loss;
        Modelica.Blocks.Sources.TimeTable powerTable(table = [0, 0; 1800, 0; 5400, 0; 9000, 0; 12600, 0; 16200, 0; 19800, 0; 23400, 5794.176452; 27000, 28338.43317; 30600, 53483.19485; 34200, 74004.40013; 37800, 81813.77151; 41400, 81004.55093; 45000, 91143.37766; 48600, 68771.96417; 52200, 42900.47528; 55800, 12342.57791; 59400, 0; 63000, 0; 66600, 0; 70200, 0; 73800, 0; 77400, 0; 81000, 0; 84600, 0; 88200, 0; 91800, 0; 95400, 0; 99000, 0; 102600, 0; 106200, 0; 109800, 5794.176452; 113400, 28338.43317; 117000, 53483.19485; 120600, 74004.40013; 124200, 81813.77151; 127800, 81004.55093; 131400, 91143.37766; 135000, 68771.96417; 138600, 42900.47528; 142200, 12342.57791; 145800, 0; 149400, 0; 153000, 0; 156600, 0; 160200, 0; 163800, 0; 167400, 0; 171000, 0; 174600, 0; 178200, 0; 181800, 0; 185400, 0; 189000, 0; 192600, 0; 196200, 5794.176452; 199800, 28338.43317; 203400, 53483.19485; 207000, 74004.40013; 210600, 81813.77151; 214200, 81004.55093; 217800, 91143.37766; 221400, 68771.96417; 225000, 42900.47528; 228600, 12342.57791; 232200, 0; 235800, 0; 239400, 0; 243000, 0; 246600, 0; 250200, 0; 253800, 0; 257400, 0; 261000, 0; 264600, 0; 268200, 0; 271800, 0; 275400, 0; 279000, 0; 282600, 5794.176452; 286200, 28338.43317; 289800, 53483.19485; 293400, 74004.40013; 297000, 81813.77151; 300600, 81004.55093; 304200, 91143.37766; 307800, 68771.96417; 311400, 42900.47528; 315000, 12342.57791; 318600, 0; 322200, 0; 325800, 0; 329400, 0; 333000, 0; 336600, 0; 340200, 0; 343800, 0; 347400, 0; 351000, 0; 354600, 0; 358200, 0; 361800, 0; 365400, 0; 369000, 5794.176452; 372600, 28338.43317; 376200, 53483.19485; 379800, 74004.40013; 383400, 81813.77151; 387000, 81004.55093; 390600, 91143.37766; 394200, 68771.96417; 397800, 42900.47528; 401400, 12342.57791; 405000, 0; 408600, 0; 412200, 0; 415800, 0; 419400, 0; 423000, 0; 426600, 0; 430200, 0; 433800, 0; 437400, 0; 441000, 0; 444600, 0; 448200, 0; 451800, 0; 455400, 5794.176452; 459000, 28338.43317; 462600, 53483.19485; 466200, 74004.40013; 469800, 81813.77151; 473400, 81004.55093; 477000, 91143.37766; 480600, 68771.96417; 484200, 42900.47528; 487800, 12342.57791; 491400, 0; 495000, 0; 498600, 0; 502200, 0; 505800, 0; 509400, 0; 513000, 0; 516600, 0; 520200, 0; 523800, 0; 527400, 0; 531000, 0; 534600, 0; 538200, 0; 541800, 5794.176452; 545400, 28338.43317; 549000, 53483.19485; 552600, 74004.40013; 556200, 81813.77151; 559800, 81004.55093; 563400, 91143.37766; 567000, 68771.96417; 570600, 42900.47528; 574200, 12342.57791; 577800, 0; 581400, 0; 585000, 0; 588600, 0; 592200, 0; 595800, 0; 599400, 0; 603000, 0; 606600, 0; 610200, 0; 613800, 0; 617400, 0; 621000, 0; 624600, 0; 628200, 5794.176452; 631800, 28338.43317; 635400, 53483.19485; 639000, 74004.40013; 642600, 81813.77151; 646200, 81004.55093; 649800, 91143.37766; 653400, 68771.96417; 657000, 42900.47528; 660600, 12342.57791; 664200, 0; 667800, 0; 671400, 0; 675000, 0; 678600, 0; 682200, 0; 685800, 0; 689400, 0; 693000, 0; 696600, 0; 700200, 0; 703800, 0; 707400, 0; 711000, 0; 714600, 5794.176452; 718200, 28338.43317; 721800, 53483.19485; 725400, 74004.40013; 729000, 81813.77151; 732600, 81004.55093; 736200, 91143.37766; 739800, 68771.96417; 743400, 42900.47528; 747000, 12342.57791; 750600, 0; 754200, 0; 757800, 0; 761400, 0; 765000, 0; 768600, 0; 772200, 0; 775800, 0; 779400, 0; 783000, 0; 786600, 0; 790200, 0; 793800, 0; 797400, 0; 801000, 5794.176452; 804600, 28338.43317; 808200, 53483.19485; 811800, 74004.40013; 815400, 81813.77151; 819000, 81004.55093; 822600, 91143.37766; 826200, 68771.96417; 829800, 42900.47528; 833400, 12342.57791; 837000, 0; 840600, 0; 844200, 0; 847800, 0; 851400, 0; 855000, 0; 858600, 0; 862200, 0; 865800, 0; 869400, 0; 873000, 0; 876600, 0; 880200, 0; 883800, 0; 887400, 5794.176452; 891000, 28338.43317; 894600, 53483.19485; 898200, 74004.40013; 901800, 81813.77151; 905400, 81004.55093; 909000, 91143.37766; 912600, 68771.96417; 916200, 42900.47528; 919800, 12342.57791; 923400, 0; 927000, 0; 930600, 0; 934200, 0; 937800, 0; 941400, 0; 945000, 0; 948600, 0; 952200, 0; 955800, 0; 959400, 0; 963000, 0; 966600, 0; 970200, 0; 973800, 5794.176452; 977400, 28338.43317; 981000, 53483.19485; 984600, 74004.40013; 988200, 81813.77151; 991800, 81004.55093; 995400, 91143.37766; 999000, 68771.96417; 1002600, 42900.47528; 1006200, 12342.57791; 1009800, 0; 1013400, 0; 1017000, 0; 1020600, 0; 1024200, 0; 1027800, 0; 1031400, 0; 1035000, 0; 1038600, 0; 1042200, 0; 1045800, 0; 1049400, 0; 1053000, 0; 1056600, 0; 1060200, 5794.176452; 1063800, 28338.43317; 1067400, 53483.19485; 1071000, 74004.40013; 1074600, 81813.77151; 1078200, 81004.55093; 1081800, 91143.37766; 1085400, 68771.96417; 1089000, 42900.47528; 1092600, 12342.57791; 1096200, 0; 1099800, 0; 1103400, 0; 1107000, 0; 1110600, 0; 1114200, 0; 1117800, 0; 1121400, 0; 1125000, 0; 1128600, 0; 1132200, 0; 1135800, 0; 1139400, 0; 1143000, 0; 1146600, 5794.176452; 1150200, 28338.43317; 1153800, 53483.19485; 1157400, 74004.40013; 1161000, 81813.77151; 1164600, 81004.55093; 1168200, 91143.37766; 1171800, 68771.96417; 1175400, 42900.47528; 1179000, 12342.57791; 1182600, 0; 1186200, 0; 1189800, 0; 1193400, 0; 1197000, 0; 1200600, 0; 1204200, 0; 1207800, 0; 1211400, 0; 1215000, 0; 1218600, 0; 1222200, 0; 1225800, 0; 1229400, 0; 1233000, 5794.176452; 1236600, 28338.43317; 1240200, 53483.19485; 1243800, 74004.40013; 1247400, 81813.77151; 1251000, 81004.55093; 1254600, 91143.37766; 1258200, 68771.96417; 1261800, 42900.47528; 1265400, 12342.57791; 1269000, 0; 1272600, 0; 1276200, 0; 1279800, 0; 1283400, 0; 1287000, 0; 1290600, 0; 1294200, 0; 1297800, 0; 1301400, 0; 1305000, 0; 1308600, 0; 1312200, 0; 1315800, 0; 1319400, 5794.176452; 1323000, 28338.43317; 1326600, 53483.19485; 1330200, 74004.40013; 1333800, 81813.77151; 1337400, 81004.55093; 1341000, 91143.37766; 1344600, 68771.96417; 1348200, 42900.47528; 1351800, 12342.57791; 1355400, 0; 1359000, 0; 1362600, 0; 1366200, 0; 1369800, 0; 1373400, 0; 1377000, 0; 1380600, 0; 1384200, 0; 1387800, 0; 1391400, 0; 1395000, 0; 1398600, 0; 1402200, 0; 1405800, 5794.176452; 1409400, 28338.43317; 1413000, 53483.19485; 1416600, 74004.40013; 1420200, 81813.77151; 1423800, 81004.55093; 1427400, 91143.37766; 1431000, 68771.96417; 1434600, 42900.47528; 1438200, 12342.57791; 1441800, 0; 1445400, 0; 1449000, 0; 1452600, 0; 1456200, 0; 1459800, 0; 1463400, 0; 1467000, 0; 1470600, 0; 1474200, 0; 1477800, 0; 1481400, 0; 1485000, 0; 1488600, 0; 1492200, 5794.176452; 1495800, 28338.43317; 1499400, 53483.19485; 1503000, 74004.40013; 1506600, 81813.77151; 1510200, 81004.55093; 1513800, 91143.37766; 1517400, 68771.96417; 1521000, 42900.47528; 1524600, 12342.57791; 1528200, 0; 1531800, 0; 1535400, 0; 1539000, 0; 1542600, 0; 1546200, 0; 1549800, 0; 1553400, 0; 1557000, 0; 1560600, 0; 1564200, 0; 1567800, 0; 1571400, 0; 1575000, 0; 1578600, 5794.176452; 1582200, 28338.43317; 1585800, 53483.19485; 1589400, 74004.40013; 1593000, 81813.77151; 1596600, 81004.55093; 1600200, 91143.37766; 1603800, 68771.96417; 1607400, 42900.47528; 1611000, 12342.57791; 1614600, 0; 1618200, 0; 1621800, 0; 1625400, 0; 1629000, 0; 1632600, 0; 1636200, 0; 1639800, 0; 1643400, 0; 1647000, 0; 1650600, 0; 1654200, 0; 1657800, 0; 1661400, 0; 1665000, 5794.176452; 1668600, 28338.43317; 1672200, 53483.19485; 1675800, 74004.40013; 1679400, 81813.77151; 1683000, 81004.55093; 1686600, 91143.37766; 1690200, 68771.96417; 1693800, 42900.47528; 1697400, 12342.57791; 1701000, 0; 1704600, 0; 1708200, 0; 1711800, 0; 1715400, 0; 1719000, 0; 1722600, 0; 1726200, 0; 1729800, 0; 1733400, 0; 1737000, 0; 1740600, 0; 1744200, 0; 1747800, 0; 1751400, 5794.176452; 1755000, 28338.43317; 1758600, 53483.19485; 1762200, 74004.40013; 1765800, 81813.77151; 1769400, 81004.55093; 1773000, 91143.37766; 1776600, 68771.96417; 1780200, 42900.47528; 1783800, 12342.57791; 1787400, 0; 1791000, 0; 1794600, 0; 1798200, 0; 1801800, 0; 1805400, 0; 1809000, 0; 1812600, 0; 1816200, 0; 1819800, 0; 1823400, 0; 1827000, 0; 1830600, 0; 1834200, 0; 1837800, 5794.176452; 1841400, 28338.43317; 1845000, 53483.19485; 1848600, 74004.40013; 1852200, 81813.77151; 1855800, 81004.55093; 1859400, 91143.37766; 1863000, 68771.96417; 1866600, 42900.47528; 1870200, 12342.57791; 1873800, 0; 1877400, 0; 1881000, 0; 1884600, 0; 1888200, 0; 1891800, 0; 1895400, 0; 1899000, 0; 1902600, 0; 1906200, 0; 1909800, 0; 1913400, 0; 1917000, 0; 1920600, 0; 1924200, 5794.176452; 1927800, 28338.43317; 1931400, 53483.19485; 1935000, 74004.40013; 1938600, 81813.77151; 1942200, 81004.55093; 1945800, 91143.37766; 1949400, 68771.96417; 1953000, 42900.47528; 1956600, 12342.57791; 1960200, 0; 1963800, 0; 1967400, 0; 1971000, 0; 1974600, 0; 1978200, 0; 1981800, 0; 1985400, 0; 1989000, 0; 1992600, 0; 1996200, 0; 1999800, 0; 2003400, 0; 2007000, 0; 2010600, 5794.176452; 2014200, 28338.43317; 2017800, 53483.19485; 2021400, 74004.40013; 2025000, 81813.77151; 2028600, 81004.55093; 2032200, 91143.37766; 2035800, 68771.96417; 2039400, 42900.47528; 2043000, 12342.57791; 2046600, 0; 2050200, 0; 2053800, 0; 2057400, 0; 2061000, 0; 2064600, 0; 2068200, 0; 2071800, 0; 2075400, 0; 2079000, 0; 2082600, 0; 2086200, 0; 2089800, 0; 2093400, 0; 2097000, 5794.176452; 2100600, 28338.43317; 2104200, 53483.19485; 2107800, 74004.40013; 2111400, 81813.77151; 2115000, 81004.55093; 2118600, 91143.37766; 2122200, 68771.96417; 2125800, 42900.47528; 2129400, 12342.57791; 2133000, 0; 2136600, 0; 2140200, 0; 2143800, 0; 2147400, 0; 2151000, 0; 2154600, 0; 2158200, 0; 2161800, 0; 2165400, 0; 2169000, 0; 2172600, 0; 2176200, 0; 2179800, 0; 2183400, 5794.176452; 2187000, 28338.43317; 2190600, 53483.19485; 2194200, 74004.40013; 2197800, 81813.77151; 2201400, 81004.55093; 2205000, 91143.37766; 2208600, 68771.96417; 2212200, 42900.47528; 2215800, 12342.57791; 2219400, 0; 2223000, 0; 2226600, 0; 2230200, 0; 2233800, 0; 2237400, 0; 2241000, 0; 2244600, 0; 2248200, 0; 2251800, 0; 2255400, 0; 2259000, 0; 2262600, 0; 2266200, 0; 2269800, 5794.176452; 2273400, 28338.43317; 2277000, 53483.19485; 2280600, 74004.40013; 2284200, 81813.77151; 2287800, 81004.55093; 2291400, 91143.37766; 2295000, 68771.96417; 2298600, 42900.47528; 2302200, 12342.57791; 2305800, 0; 2309400, 0; 2313000, 0; 2316600, 0; 2320200, 0; 2323800, 0; 2327400, 0; 2331000, 0; 2334600, 0; 2338200, 0; 2341800, 0; 2345400, 0; 2349000, 0; 2352600, 0; 2356200, 5794.176452; 2359800, 28338.43317; 2363400, 53483.19485; 2367000, 74004.40013; 2370600, 81813.77151; 2374200, 81004.55093; 2377800, 91143.37766; 2381400, 68771.96417; 2385000, 42900.47528; 2388600, 12342.57791; 2392200, 0; 2395800, 0; 2399400, 0; 2403000, 0; 2406600, 0; 2410200, 0; 2413800, 0; 2417400, 0; 2421000, 0; 2424600, 0; 2428200, 0; 2431800, 0; 2435400, 0; 2439000, 0; 2442600, 5794.176452; 2446200, 28338.43317; 2449800, 53483.19485; 2453400, 74004.40013; 2457000, 81813.77151; 2460600, 81004.55093; 2464200, 91143.37766; 2467800, 68771.96417; 2471400, 42900.47528; 2475000, 12342.57791; 2478600, 0; 2482200, 0; 2485800, 0; 2489400, 0; 2493000, 0; 2496600, 0; 2500200, 0; 2503800, 0; 2507400, 0; 2511000, 0; 2514600, 0; 2518200, 0; 2521800, 0; 2525400, 0; 2529000, 5794.176452; 2532600, 28338.43317; 2536200, 53483.19485; 2539800, 74004.40013; 2543400, 81813.77151; 2547000, 81004.55093; 2550600, 91143.37766; 2554200, 68771.96417; 2557800, 42900.47528; 2561400, 12342.57791; 2565000, 0; 2568600, 0; 2572200, 0; 2575800, 0; 2579400, 0; 2583000, 0; 2586600, 0; 2590200, 0; 2593800, 0; 2597400, 0; 2601000, 0; 2604600, 0; 2608200, 0; 2611800, 0; 2615400, 5794.176452; 2619000, 28338.43317; 2622600, 53483.19485; 2626200, 74004.40013; 2629800, 81813.77151; 2633400, 81004.55093; 2637000, 91143.37766; 2640600, 68771.96417; 2644200, 42900.47528; 2647800, 12342.57791; 2651400, 0; 2655000, 0; 2658600, 0; 2662200, 0; 2665800, 0; 2669400, 0; 2673000, 0; 2676600, 0; 2680200, 0; 2683800, 0; 2687400, 0; 2691000, 0; 2694600, 0; 2698200, 0; 2701800, 12681.49151; 2705400, 39997.0883; 2709000, 65528.0381; 2712600, 86297.12983; 2716200, 96971.41838; 2719800, 98810.65457; 2723400, 93326.53936; 2727000, 77285.71235; 2730600, 54135.69158; 2734200, 23813.48953; 2737800, 0; 2741400, 0; 2745000, 0; 2748600, 0; 2752200, 0; 2755800, 0; 2759400, 0; 2763000, 0; 2766600, 0; 2770200, 0; 2773800, 0; 2777400, 0; 2781000, 0; 2784600, 0; 2788200, 12681.49151; 2791800, 39997.0883; 2795400, 65528.0381; 2799000, 86297.12983; 2802600, 96971.41838; 2806200, 98810.65457; 2809800, 93326.53936; 2813400, 77285.71235; 2817000, 54135.69158; 2820600, 23813.48953; 2824200, 0; 2827800, 0; 2831400, 0; 2835000, 0; 2838600, 0; 2842200, 0; 2845800, 0; 2849400, 0; 2853000, 0; 2856600, 0; 2860200, 0; 2863800, 0; 2867400, 0; 2871000, 0; 2874600, 12681.49151; 2878200, 39997.0883; 2881800, 65528.0381; 2885400, 86297.12983; 2889000, 96971.41838; 2892600, 98810.65457; 2896200, 93326.53936; 2899800, 77285.71235; 2903400, 54135.69158; 2907000, 23813.48953; 2910600, 0; 2914200, 0; 2917800, 0; 2921400, 0; 2925000, 0; 2928600, 0; 2932200, 0; 2935800, 0; 2939400, 0; 2943000, 0; 2946600, 0; 2950200, 0; 2953800, 0; 2957400, 0; 2961000, 12681.49151; 2964600, 39997.0883; 2968200, 65528.0381; 2971800, 86297.12983; 2975400, 96971.41838; 2979000, 98810.65457; 2982600, 93326.53936; 2986200, 77285.71235; 2989800, 54135.69158; 2993400, 23813.48953; 2997000, 0; 3000600, 0; 3004200, 0; 3007800, 0; 3011400, 0; 3015000, 0; 3018600, 0; 3022200, 0; 3025800, 0; 3029400, 0; 3033000, 0; 3036600, 0; 3040200, 0; 3043800, 0; 3047400, 12681.49151; 3051000, 39997.0883; 3054600, 65528.0381; 3058200, 86297.12983; 3061800, 96971.41838; 3065400, 98810.65457; 3069000, 93326.53936; 3072600, 77285.71235; 3076200, 54135.69158; 3079800, 23813.48953; 3083400, 0; 3087000, 0; 3090600, 0; 3094200, 0; 3097800, 0; 3101400, 0; 3105000, 0; 3108600, 0; 3112200, 0; 3115800, 0; 3119400, 0; 3123000, 0; 3126600, 0; 3130200, 0; 3133800, 12681.49151; 3137400, 39997.0883; 3141000, 65528.0381; 3144600, 86297.12983; 3148200, 96971.41838; 3151800, 98810.65457; 3155400, 93326.53936; 3159000, 77285.71235; 3162600, 54135.69158; 3166200, 23813.48953; 3169800, 0; 3173400, 0; 3177000, 0; 3180600, 0; 3184200, 0; 3187800, 0; 3191400, 0; 3195000, 0; 3198600, 0; 3202200, 0; 3205800, 0; 3209400, 0; 3213000, 0; 3216600, 0; 3220200, 12681.49151; 3223800, 39997.0883; 3227400, 65528.0381; 3231000, 86297.12983; 3234600, 96971.41838; 3238200, 98810.65457; 3241800, 93326.53936; 3245400, 77285.71235; 3249000, 54135.69158; 3252600, 23813.48953; 3256200, 0; 3259800, 0; 3263400, 0; 3267000, 0; 3270600, 0; 3274200, 0; 3277800, 0; 3281400, 0; 3285000, 0; 3288600, 0; 3292200, 0; 3295800, 0; 3299400, 0; 3303000, 0; 3306600, 12681.49151; 3310200, 39997.0883; 3313800, 65528.0381; 3317400, 86297.12983; 3321000, 96971.41838; 3324600, 98810.65457; 3328200, 93326.53936; 3331800, 77285.71235; 3335400, 54135.69158; 3339000, 23813.48953; 3342600, 0; 3346200, 0; 3349800, 0; 3353400, 0; 3357000, 0; 3360600, 0; 3364200, 0; 3367800, 0; 3371400, 0; 3375000, 0; 3378600, 0; 3382200, 0; 3385800, 0; 3389400, 0; 3393000, 12681.49151; 3396600, 39997.0883; 3400200, 65528.0381; 3403800, 86297.12983; 3407400, 96971.41838; 3411000, 98810.65457; 3414600, 93326.53936; 3418200, 77285.71235; 3421800, 54135.69158; 3425400, 23813.48953; 3429000, 0; 3432600, 0; 3436200, 0; 3439800, 0; 3443400, 0; 3447000, 0; 3450600, 0; 3454200, 0; 3457800, 0; 3461400, 0; 3465000, 0; 3468600, 0; 3472200, 0; 3475800, 0; 3479400, 12681.49151; 3483000, 39997.0883; 3486600, 65528.0381; 3490200, 86297.12983; 3493800, 96971.41838; 3497400, 98810.65457; 3501000, 93326.53936; 3504600, 77285.71235; 3508200, 54135.69158; 3511800, 23813.48953; 3515400, 0; 3519000, 0; 3522600, 0; 3526200, 0; 3529800, 0; 3533400, 0; 3537000, 0; 3540600, 0; 3544200, 0; 3547800, 0; 3551400, 0; 3555000, 0; 3558600, 0; 3562200, 0; 3565800, 12681.49151; 3569400, 39997.0883; 3573000, 65528.0381; 3576600, 86297.12983; 3580200, 96971.41838; 3583800, 98810.65457; 3587400, 93326.53936; 3591000, 77285.71235; 3594600, 54135.69158; 3598200, 23813.48953; 3601800, 0; 3605400, 0; 3609000, 0; 3612600, 0; 3616200, 0; 3619800, 0; 3623400, 0; 3627000, 0; 3630600, 0; 3634200, 0; 3637800, 0; 3641400, 0; 3645000, 0; 3648600, 0; 3652200, 12681.49151; 3655800, 39997.0883; 3659400, 65528.0381; 3663000, 86297.12983; 3666600, 96971.41838; 3670200, 98810.65457; 3673800, 93326.53936; 3677400, 77285.71235; 3681000, 54135.69158; 3684600, 23813.48953; 3688200, 0; 3691800, 0; 3695400, 0; 3699000, 0; 3702600, 0; 3706200, 0; 3709800, 0; 3713400, 0; 3717000, 0; 3720600, 0; 3724200, 0; 3727800, 0; 3731400, 0; 3735000, 0; 3738600, 12681.49151; 3742200, 39997.0883; 3745800, 65528.0381; 3749400, 86297.12983; 3753000, 96971.41838; 3756600, 98810.65457; 3760200, 93326.53936; 3763800, 77285.71235; 3767400, 54135.69158; 3771000, 23813.48953; 3774600, 0; 3778200, 0; 3781800, 0; 3785400, 0; 3789000, 0; 3792600, 0; 3796200, 0; 3799800, 0; 3803400, 0; 3807000, 0; 3810600, 0; 3814200, 0; 3817800, 0; 3821400, 0; 3825000, 12681.49151; 3828600, 39997.0883; 3832200, 65528.0381; 3835800, 86297.12983; 3839400, 96971.41838; 3843000, 98810.65457; 3846600, 93326.53936; 3850200, 77285.71235; 3853800, 54135.69158; 3857400, 23813.48953; 3861000, 0; 3864600, 0; 3868200, 0; 3871800, 0; 3875400, 0; 3879000, 0; 3882600, 0; 3886200, 0; 3889800, 0; 3893400, 0; 3897000, 0; 3900600, 0; 3904200, 0; 3907800, 0; 3911400, 12681.49151; 3915000, 39997.0883; 3918600, 65528.0381; 3922200, 86297.12983; 3925800, 96971.41838; 3929400, 98810.65457; 3933000, 93326.53936; 3936600, 77285.71235; 3940200, 54135.69158; 3943800, 23813.48953; 3947400, 0; 3951000, 0; 3954600, 0; 3958200, 0; 3961800, 0; 3965400, 0; 3969000, 0; 3972600, 0; 3976200, 0; 3979800, 0; 3983400, 0; 3987000, 0; 3990600, 0; 3994200, 0; 3997800, 12681.49151; 4001400, 39997.0883; 4005000, 65528.0381; 4008600, 86297.12983; 4012200, 96971.41838; 4015800, 98810.65457; 4019400, 93326.53936; 4023000, 77285.71235; 4026600, 54135.69158; 4030200, 23813.48953; 4033800, 0; 4037400, 0; 4041000, 0; 4044600, 0; 4048200, 0; 4051800, 0; 4055400, 0; 4059000, 0; 4062600, 0; 4066200, 0; 4069800, 0; 4073400, 0; 4077000, 0; 4080600, 0; 4084200, 12681.49151; 4087800, 39997.0883; 4091400, 65528.0381; 4095000, 86297.12983; 4098600, 96971.41838; 4102200, 98810.65457; 4105800, 93326.53936; 4109400, 77285.71235; 4113000, 54135.69158; 4116600, 23813.48953; 4120200, 0; 4123800, 0; 4127400, 0; 4131000, 0; 4134600, 0; 4138200, 0; 4141800, 0; 4145400, 0; 4149000, 0; 4152600, 0; 4156200, 0; 4159800, 0; 4163400, 0; 4167000, 0; 4170600, 12681.49151; 4174200, 39997.0883; 4177800, 65528.0381; 4181400, 86297.12983; 4185000, 96971.41838; 4188600, 98810.65457; 4192200, 93326.53936; 4195800, 77285.71235; 4199400, 54135.69158; 4203000, 23813.48953; 4206600, 0; 4210200, 0; 4213800, 0; 4217400, 0; 4221000, 0; 4224600, 0; 4228200, 0; 4231800, 0; 4235400, 0; 4239000, 0; 4242600, 0; 4246200, 0; 4249800, 0; 4253400, 0; 4257000, 12681.49151; 4260600, 39997.0883; 4264200, 65528.0381; 4267800, 86297.12983; 4271400, 96971.41838; 4275000, 98810.65457; 4278600, 93326.53936; 4282200, 77285.71235; 4285800, 54135.69158; 4289400, 23813.48953; 4293000, 0; 4296600, 0; 4300200, 0; 4303800, 0; 4307400, 0; 4311000, 0; 4314600, 0; 4318200, 0; 4321800, 0; 4325400, 0; 4329000, 0; 4332600, 0; 4336200, 0; 4339800, 0; 4343400, 12681.49151; 4347000, 39997.0883; 4350600, 65528.0381; 4354200, 86297.12983; 4357800, 96971.41838; 4361400, 98810.65457; 4365000, 93326.53936; 4368600, 77285.71235; 4372200, 54135.69158; 4375800, 23813.48953; 4379400, 0; 4383000, 0; 4386600, 0; 4390200, 0; 4393800, 0; 4397400, 0; 4401000, 0; 4404600, 0; 4408200, 0; 4411800, 0; 4415400, 0; 4419000, 0; 4422600, 0; 4426200, 0; 4429800, 12681.49151; 4433400, 39997.0883; 4437000, 65528.0381; 4440600, 86297.12983; 4444200, 96971.41838; 4447800, 98810.65457; 4451400, 93326.53936; 4455000, 77285.71235; 4458600, 54135.69158; 4462200, 23813.48953; 4465800, 0; 4469400, 0; 4473000, 0; 4476600, 0; 4480200, 0; 4483800, 0; 4487400, 0; 4491000, 0; 4494600, 0; 4498200, 0; 4501800, 0; 4505400, 0; 4509000, 0; 4512600, 0; 4516200, 12681.49151; 4519800, 39997.0883; 4523400, 65528.0381; 4527000, 86297.12983; 4530600, 96971.41838; 4534200, 98810.65457; 4537800, 93326.53936; 4541400, 77285.71235; 4545000, 54135.69158; 4548600, 23813.48953; 4552200, 0; 4555800, 0; 4559400, 0; 4563000, 0; 4566600, 0; 4570200, 0; 4573800, 0; 4577400, 0; 4581000, 0; 4584600, 0; 4588200, 0; 4591800, 0; 4595400, 0; 4599000, 0; 4602600, 12681.49151; 4606200, 39997.0883; 4609800, 65528.0381; 4613400, 86297.12983; 4617000, 96971.41838; 4620600, 98810.65457; 4624200, 93326.53936; 4627800, 77285.71235; 4631400, 54135.69158; 4635000, 23813.48953; 4638600, 0; 4642200, 0; 4645800, 0; 4649400, 0; 4653000, 0; 4656600, 0; 4660200, 0; 4663800, 0; 4667400, 0; 4671000, 0; 4674600, 0; 4678200, 0; 4681800, 0; 4685400, 0; 4689000, 12681.49151; 4692600, 39997.0883; 4696200, 65528.0381; 4699800, 86297.12983; 4703400, 96971.41838; 4707000, 98810.65457; 4710600, 93326.53936; 4714200, 77285.71235; 4717800, 54135.69158; 4721400, 23813.48953; 4725000, 0; 4728600, 0; 4732200, 0; 4735800, 0; 4739400, 0; 4743000, 0; 4746600, 0; 4750200, 0; 4753800, 0; 4757400, 0; 4761000, 0; 4764600, 0; 4768200, 0; 4771800, 0; 4775400, 12681.49151; 4779000, 39997.0883; 4782600, 65528.0381; 4786200, 86297.12983; 4789800, 96971.41838; 4793400, 98810.65457; 4797000, 93326.53936; 4800600, 77285.71235; 4804200, 54135.69158; 4807800, 23813.48953; 4811400, 0; 4815000, 0; 4818600, 0; 4822200, 0; 4825800, 0; 4829400, 0; 4833000, 0; 4836600, 0; 4840200, 0; 4843800, 0; 4847400, 0; 4851000, 0; 4854600, 0; 4858200, 0; 4861800, 12681.49151; 4865400, 39997.0883; 4869000, 65528.0381; 4872600, 86297.12983; 4876200, 96971.41838; 4879800, 98810.65457; 4883400, 93326.53936; 4887000, 77285.71235; 4890600, 54135.69158; 4894200, 23813.48953; 4897800, 0; 4901400, 0; 4905000, 0; 4908600, 0; 4912200, 0; 4915800, 0; 4919400, 0; 4923000, 0; 4926600, 0; 4930200, 0; 4933800, 0; 4937400, 0; 4941000, 0; 4944600, 0; 4948200, 12681.49151; 4951800, 39997.0883; 4955400, 65528.0381; 4959000, 86297.12983; 4962600, 96971.41838; 4966200, 98810.65457; 4969800, 93326.53936; 4973400, 77285.71235; 4977000, 54135.69158; 4980600, 23813.48953; 4984200, 0; 4987800, 0; 4991400, 0; 4995000, 0; 4998600, 0; 5002200, 0; 5005800, 0; 5009400, 0; 5013000, 0; 5016600, 0; 5020200, 0; 5023800, 0; 5027400, 0; 5031000, 0; 5034600, 12681.49151; 5038200, 39997.0883; 5041800, 65528.0381; 5045400, 86297.12983; 5049000, 96971.41838; 5052600, 98810.65457; 5056200, 93326.53936; 5059800, 77285.71235; 5063400, 54135.69158; 5067000, 23813.48953; 5070600, 0; 5074200, 0; 5077800, 0; 5081400, 0; 5085000, 0; 5088600, 0; 5092200, 0; 5095800, 0; 5099400, 0; 5103000, 0; 5106600, 0; 5110200, 0; 5113800, 0; 5117400, 1865.92123; 5121000, 27222.80869; 5124600, 56818.28353; 5128200, 81043.83348; 5131800, 95228.76309; 5135400, 103258.1168; 5139000, 104994.4056; 5142600, 94835.93757; 5146200, 79562.88126; 5149800, 57709.99746; 5153400, 31763.87172; 5157000, 4827.825668; 5160600, 0; 5164200, 0; 5167800, 0; 5171400, 0; 5175000, 0; 5178600, 0; 5182200, 0; 5185800, 0; 5189400, 0; 5193000, 0; 5196600, 0; 5200200, 0; 5203800, 1865.92123; 5207400, 27222.80869; 5211000, 56818.28353; 5214600, 81043.83348; 5218200, 95228.76309; 5221800, 103258.1168; 5225400, 104994.4056; 5229000, 94835.93757; 5232600, 79562.88126; 5236200, 57709.99746; 5239800, 31763.87172; 5243400, 4827.825668; 5247000, 0; 5250600, 0; 5254200, 0; 5257800, 0; 5261400, 0; 5265000, 0; 5268600, 0; 5272200, 0; 5275800, 0; 5279400, 0; 5283000, 0; 5286600, 0; 5290200, 1865.92123; 5293800, 27222.80869; 5297400, 56818.28353; 5301000, 81043.83348; 5304600, 95228.76309; 5308200, 103258.1168; 5311800, 104994.4056; 5315400, 94835.93757; 5319000, 79562.88126; 5322600, 57709.99746; 5326200, 31763.87172; 5329800, 4827.825668; 5333400, 0; 5337000, 0; 5340600, 0; 5344200, 0; 5347800, 0; 5351400, 0; 5355000, 0; 5358600, 0; 5362200, 0; 5365800, 0; 5369400, 0; 5373000, 0; 5376600, 1865.92123; 5380200, 27222.80869; 5383800, 56818.28353; 5387400, 81043.83348; 5391000, 95228.76309; 5394600, 103258.1168; 5398200, 104994.4056; 5401800, 94835.93757; 5405400, 79562.88126; 5409000, 57709.99746; 5412600, 31763.87172; 5416200, 4827.825668; 5419800, 0; 5423400, 0; 5427000, 0; 5430600, 0; 5434200, 0; 5437800, 0; 5441400, 0; 5445000, 0; 5448600, 0; 5452200, 0; 5455800, 0; 5459400, 0; 5463000, 1865.92123; 5466600, 27222.80869; 5470200, 56818.28353; 5473800, 81043.83348; 5477400, 95228.76309; 5481000, 103258.1168; 5484600, 104994.4056; 5488200, 94835.93757; 5491800, 79562.88126; 5495400, 57709.99746; 5499000, 31763.87172; 5502600, 4827.825668; 5506200, 0; 5509800, 0; 5513400, 0; 5517000, 0; 5520600, 0; 5524200, 0; 5527800, 0; 5531400, 0; 5535000, 0; 5538600, 0; 5542200, 0; 5545800, 0; 5549400, 1865.92123; 5553000, 27222.80869; 5556600, 56818.28353; 5560200, 81043.83348; 5563800, 95228.76309; 5567400, 103258.1168; 5571000, 104994.4056; 5574600, 94835.93757; 5578200, 79562.88126; 5581800, 57709.99746; 5585400, 31763.87172; 5589000, 4827.825668; 5592600, 0; 5596200, 0; 5599800, 0; 5603400, 0; 5607000, 0; 5610600, 0; 5614200, 0; 5617800, 0; 5621400, 0; 5625000, 0; 5628600, 0; 5632200, 0; 5635800, 1865.92123; 5639400, 27222.80869; 5643000, 56818.28353; 5646600, 81043.83348; 5650200, 95228.76309; 5653800, 103258.1168; 5657400, 104994.4056; 5661000, 94835.93757; 5664600, 79562.88126; 5668200, 57709.99746; 5671800, 31763.87172; 5675400, 4827.825668; 5679000, 0; 5682600, 0; 5686200, 0; 5689800, 0; 5693400, 0; 5697000, 0; 5700600, 0; 5704200, 0; 5707800, 0; 5711400, 0; 5715000, 0; 5718600, 0; 5722200, 1865.92123; 5725800, 27222.80869; 5729400, 56818.28353; 5733000, 81043.83348; 5736600, 95228.76309; 5740200, 103258.1168; 5743800, 104994.4056; 5747400, 94835.93757; 5751000, 79562.88126; 5754600, 57709.99746; 5758200, 31763.87172; 5761800, 4827.825668; 5765400, 0; 5769000, 0; 5772600, 0; 5776200, 0; 5779800, 0; 5783400, 0; 5787000, 0; 5790600, 0; 5794200, 0; 5797800, 0; 5801400, 0; 5805000, 0; 5808600, 1865.92123; 5812200, 27222.80869; 5815800, 56818.28353; 5819400, 81043.83348; 5823000, 95228.76309; 5826600, 103258.1168; 5830200, 104994.4056; 5833800, 94835.93757; 5837400, 79562.88126; 5841000, 57709.99746; 5844600, 31763.87172; 5848200, 4827.825668; 5851800, 0; 5855400, 0; 5859000, 0; 5862600, 0; 5866200, 0; 5869800, 0; 5873400, 0; 5877000, 0; 5880600, 0; 5884200, 0; 5887800, 0; 5891400, 0; 5895000, 1865.92123; 5898600, 27222.80869; 5902200, 56818.28353; 5905800, 81043.83348; 5909400, 95228.76309; 5913000, 103258.1168; 5916600, 104994.4056; 5920200, 94835.93757; 5923800, 79562.88126; 5927400, 57709.99746; 5931000, 31763.87172; 5934600, 4827.825668; 5938200, 0; 5941800, 0; 5945400, 0; 5949000, 0; 5952600, 0; 5956200, 0; 5959800, 0; 5963400, 0; 5967000, 0; 5970600, 0; 5974200, 0; 5977800, 0; 5981400, 1865.92123; 5985000, 27222.80869; 5988600, 56818.28353; 5992200, 81043.83348; 5995800, 95228.76309; 5999400, 103258.1168; 6003000, 104994.4056; 6006600, 94835.93757; 6010200, 79562.88126; 6013800, 57709.99746; 6017400, 31763.87172; 6021000, 4827.825668; 6024600, 0; 6028200, 0; 6031800, 0; 6035400, 0; 6039000, 0; 6042600, 0; 6046200, 0; 6049800, 0; 6053400, 0; 6057000, 0; 6060600, 0; 6064200, 0; 6067800, 1865.92123; 6071400, 27222.80869; 6075000, 56818.28353; 6078600, 81043.83348; 6082200, 95228.76309; 6085800, 103258.1168; 6089400, 104994.4056; 6093000, 94835.93757; 6096600, 79562.88126; 6100200, 57709.99746; 6103800, 31763.87172; 6107400, 4827.825668; 6111000, 0; 6114600, 0; 6118200, 0; 6121800, 0; 6125400, 0; 6129000, 0; 6132600, 0; 6136200, 0; 6139800, 0; 6143400, 0; 6147000, 0; 6150600, 0; 6154200, 1865.92123; 6157800, 27222.80869; 6161400, 56818.28353; 6165000, 81043.83348; 6168600, 95228.76309; 6172200, 103258.1168; 6175800, 104994.4056; 6179400, 94835.93757; 6183000, 79562.88126; 6186600, 57709.99746; 6190200, 31763.87172; 6193800, 4827.825668; 6197400, 0; 6201000, 0; 6204600, 0; 6208200, 0; 6211800, 0; 6215400, 0; 6219000, 0; 6222600, 0; 6226200, 0; 6229800, 0; 6233400, 0; 6237000, 0; 6240600, 1865.92123; 6244200, 27222.80869; 6247800, 56818.28353; 6251400, 81043.83348; 6255000, 95228.76309; 6258600, 103258.1168; 6262200, 104994.4056; 6265800, 94835.93757; 6269400, 79562.88126; 6273000, 57709.99746; 6276600, 31763.87172; 6280200, 4827.825668; 6283800, 0; 6287400, 0; 6291000, 0; 6294600, 0; 6298200, 0; 6301800, 0; 6305400, 0; 6309000, 0; 6312600, 0; 6316200, 0; 6319800, 0; 6323400, 0; 6327000, 1865.92123; 6330600, 27222.80869; 6334200, 56818.28353; 6337800, 81043.83348; 6341400, 95228.76309; 6345000, 103258.1168; 6348600, 104994.4056; 6352200, 94835.93757; 6355800, 79562.88126; 6359400, 57709.99746; 6363000, 31763.87172; 6366600, 4827.825668; 6370200, 0; 6373800, 0; 6377400, 0; 6381000, 0; 6384600, 0; 6388200, 0; 6391800, 0; 6395400, 0; 6399000, 0; 6402600, 0; 6406200, 0; 6409800, 0; 6413400, 1865.92123; 6417000, 27222.80869; 6420600, 56818.28353; 6424200, 81043.83348; 6427800, 95228.76309; 6431400, 103258.1168; 6435000, 104994.4056; 6438600, 94835.93757; 6442200, 79562.88126; 6445800, 57709.99746; 6449400, 31763.87172; 6453000, 4827.825668; 6456600, 0; 6460200, 0; 6463800, 0; 6467400, 0; 6471000, 0; 6474600, 0; 6478200, 0; 6481800, 0; 6485400, 0; 6489000, 0; 6492600, 0; 6496200, 0; 6499800, 1865.92123; 6503400, 27222.80869; 6507000, 56818.28353; 6510600, 81043.83348; 6514200, 95228.76309; 6517800, 103258.1168; 6521400, 104994.4056; 6525000, 94835.93757; 6528600, 79562.88126; 6532200, 57709.99746; 6535800, 31763.87172; 6539400, 4827.825668; 6543000, 0; 6546600, 0; 6550200, 0; 6553800, 0; 6557400, 0; 6561000, 0; 6564600, 0; 6568200, 0; 6571800, 0; 6575400, 0; 6579000, 0; 6582600, 0; 6586200, 1865.92123; 6589800, 27222.80869; 6593400, 56818.28353; 6597000, 81043.83348; 6600600, 95228.76309; 6604200, 103258.1168; 6607800, 104994.4056; 6611400, 94835.93757; 6615000, 79562.88126; 6618600, 57709.99746; 6622200, 31763.87172; 6625800, 4827.825668; 6629400, 0; 6633000, 0; 6636600, 0; 6640200, 0; 6643800, 0; 6647400, 0; 6651000, 0; 6654600, 0; 6658200, 0; 6661800, 0; 6665400, 0; 6669000, 0; 6672600, 1865.92123; 6676200, 27222.80869; 6679800, 56818.28353; 6683400, 81043.83348; 6687000, 95228.76309; 6690600, 103258.1168; 6694200, 104994.4056; 6697800, 94835.93757; 6701400, 79562.88126; 6705000, 57709.99746; 6708600, 31763.87172; 6712200, 4827.825668; 6715800, 0; 6719400, 0; 6723000, 0; 6726600, 0; 6730200, 0; 6733800, 0; 6737400, 0; 6741000, 0; 6744600, 0; 6748200, 0; 6751800, 0; 6755400, 0; 6759000, 1865.92123; 6762600, 27222.80869; 6766200, 56818.28353; 6769800, 81043.83348; 6773400, 95228.76309; 6777000, 103258.1168; 6780600, 104994.4056; 6784200, 94835.93757; 6787800, 79562.88126; 6791400, 57709.99746; 6795000, 31763.87172; 6798600, 4827.825668; 6802200, 0; 6805800, 0; 6809400, 0; 6813000, 0; 6816600, 0; 6820200, 0; 6823800, 0; 6827400, 0; 6831000, 0; 6834600, 0; 6838200, 0; 6841800, 0; 6845400, 1865.92123; 6849000, 27222.80869; 6852600, 56818.28353; 6856200, 81043.83348; 6859800, 95228.76309; 6863400, 103258.1168; 6867000, 104994.4056; 6870600, 94835.93757; 6874200, 79562.88126; 6877800, 57709.99746; 6881400, 31763.87172; 6885000, 4827.825668; 6888600, 0; 6892200, 0; 6895800, 0; 6899400, 0; 6903000, 0; 6906600, 0; 6910200, 0; 6913800, 0; 6917400, 0; 6921000, 0; 6924600, 0; 6928200, 0; 6931800, 1865.92123; 6935400, 27222.80869; 6939000, 56818.28353; 6942600, 81043.83348; 6946200, 95228.76309; 6949800, 103258.1168; 6953400, 104994.4056; 6957000, 94835.93757; 6960600, 79562.88126; 6964200, 57709.99746; 6967800, 31763.87172; 6971400, 4827.825668; 6975000, 0; 6978600, 0; 6982200, 0; 6985800, 0; 6989400, 0; 6993000, 0; 6996600, 0; 7000200, 0; 7003800, 0; 7007400, 0; 7011000, 0; 7014600, 0; 7018200, 1865.92123; 7021800, 27222.80869; 7025400, 56818.28353; 7029000, 81043.83348; 7032600, 95228.76309; 7036200, 103258.1168; 7039800, 104994.4056; 7043400, 94835.93757; 7047000, 79562.88126; 7050600, 57709.99746; 7054200, 31763.87172; 7057800, 4827.825668; 7061400, 0; 7065000, 0; 7068600, 0; 7072200, 0; 7075800, 0; 7079400, 0; 7083000, 0; 7086600, 0; 7090200, 0; 7093800, 0; 7097400, 0; 7101000, 0; 7104600, 1865.92123; 7108200, 27222.80869; 7111800, 56818.28353; 7115400, 81043.83348; 7119000, 95228.76309; 7122600, 103258.1168; 7126200, 104994.4056; 7129800, 94835.93757; 7133400, 79562.88126; 7137000, 57709.99746; 7140600, 31763.87172; 7144200, 4827.825668; 7147800, 0; 7151400, 0; 7155000, 0; 7158600, 0; 7162200, 0; 7165800, 0; 7169400, 0; 7173000, 0; 7176600, 0; 7180200, 0; 7183800, 0; 7187400, 0; 7191000, 1865.92123; 7194600, 27222.80869; 7198200, 56818.28353; 7201800, 81043.83348; 7205400, 95228.76309; 7209000, 103258.1168; 7212600, 104994.4056; 7216200, 94835.93757; 7219800, 79562.88126; 7223400, 57709.99746; 7227000, 31763.87172; 7230600, 4827.825668; 7234200, 0; 7237800, 0; 7241400, 0; 7245000, 0; 7248600, 0; 7252200, 0; 7255800, 0; 7259400, 0; 7263000, 0; 7266600, 0; 7270200, 0; 7273800, 0; 7277400, 1865.92123; 7281000, 27222.80869; 7284600, 56818.28353; 7288200, 81043.83348; 7291800, 95228.76309; 7295400, 103258.1168; 7299000, 104994.4056; 7302600, 94835.93757; 7306200, 79562.88126; 7309800, 57709.99746; 7313400, 31763.87172; 7317000, 4827.825668; 7320600, 0; 7324200, 0; 7327800, 0; 7331400, 0; 7335000, 0; 7338600, 0; 7342200, 0; 7345800, 0; 7349400, 0; 7353000, 0; 7356600, 0; 7360200, 0; 7363800, 1865.92123; 7367400, 27222.80869; 7371000, 56818.28353; 7374600, 81043.83348; 7378200, 95228.76309; 7381800, 103258.1168; 7385400, 104994.4056; 7389000, 94835.93757; 7392600, 79562.88126; 7396200, 57709.99746; 7399800, 31763.87172; 7403400, 4827.825668; 7407000, 0; 7410600, 0; 7414200, 0; 7417800, 0; 7421400, 0; 7425000, 0; 7428600, 0; 7432200, 0; 7435800, 0; 7439400, 0; 7443000, 0; 7446600, 0; 7450200, 1865.92123; 7453800, 27222.80869; 7457400, 56818.28353; 7461000, 81043.83348; 7464600, 95228.76309; 7468200, 103258.1168; 7471800, 104994.4056; 7475400, 94835.93757; 7479000, 79562.88126; 7482600, 57709.99746; 7486200, 31763.87172; 7489800, 4827.825668; 7493400, 0; 7497000, 0; 7500600, 0; 7504200, 0; 7507800, 0; 7511400, 0; 7515000, 0; 7518600, 0; 7522200, 0; 7525800, 0; 7529400, 0; 7533000, 0; 7536600, 1865.92123; 7540200, 27222.80869; 7543800, 56818.28353; 7547400, 81043.83348; 7551000, 95228.76309; 7554600, 103258.1168; 7558200, 104994.4056; 7561800, 94835.93757; 7565400, 79562.88126; 7569000, 57709.99746; 7572600, 31763.87172; 7576200, 4827.825668; 7579800, 0; 7583400, 0; 7587000, 0; 7590600, 0; 7594200, 0; 7597800, 0; 7601400, 0; 7605000, 0; 7608600, 0; 7612200, 0; 7615800, 0; 7619400, 0; 7623000, 1865.92123; 7626600, 27222.80869; 7630200, 56818.28353; 7633800, 81043.83348; 7637400, 95228.76309; 7641000, 103258.1168; 7644600, 104994.4056; 7648200, 94835.93757; 7651800, 79562.88126; 7655400, 57709.99746; 7659000, 31763.87172; 7662600, 4827.825668; 7666200, 0; 7669800, 0; 7673400, 0; 7677000, 0; 7680600, 0; 7684200, 0; 7687800, 0; 7691400, 0; 7695000, 0; 7698600, 0; 7702200, 0; 7705800, 0; 7709400, 1865.92123; 7713000, 27222.80869; 7716600, 56818.28353; 7720200, 81043.83348; 7723800, 95228.76309; 7727400, 103258.1168; 7731000, 104994.4056; 7734600, 94835.93757; 7738200, 79562.88126; 7741800, 57709.99746; 7745400, 31763.87172; 7749000, 4827.825668; 7752600, 0; 7756200, 0; 7759800, 0; 7763400, 0; 7767000, 0; 7770600, 0; 7774200, 0; 7777800, 0; 7781400, 0; 7785000, 0; 7788600, 0; 7792200, 0; 7795800, 16744.18788; 7799400, 46250.49133; 7803000, 78062.41871; 7806600, 107065.3817; 7810200, 126496.7581; 7813800, 129135.2362; 7817400, 121986.9901; 7821000, 120225.2986; 7824600, 96629.18608; 7828200, 67074.17226; 7831800, 42231.88624; 7835400, 13854.03957; 7839000, 0; 7842600, 0; 7846200, 0; 7849800, 0; 7853400, 0; 7857000, 0; 7860600, 0; 7864200, 0; 7867800, 0; 7871400, 0; 7875000, 0; 7878600, 0; 7882200, 16744.18788; 7885800, 46250.49133; 7889400, 78062.41871; 7893000, 107065.3817; 7896600, 126496.7581; 7900200, 129135.2362; 7903800, 121986.9901; 7907400, 120225.2986; 7911000, 96629.18608; 7914600, 67074.17226; 7918200, 42231.88624; 7921800, 13854.03957; 7925400, 0; 7929000, 0; 7932600, 0; 7936200, 0; 7939800, 0; 7943400, 0; 7947000, 0; 7950600, 0; 7954200, 0; 7957800, 0; 7961400, 0; 7965000, 0; 7968600, 16744.18788; 7972200, 46250.49133; 7975800, 78062.41871; 7979400, 107065.3817; 7983000, 126496.7581; 7986600, 129135.2362; 7990200, 121986.9901; 7993800, 120225.2986; 7997400, 96629.18608; 8001000, 67074.17226; 8004600, 42231.88624; 8008200, 13854.03957; 8011800, 0; 8015400, 0; 8019000, 0; 8022600, 0; 8026200, 0; 8029800, 0; 8033400, 0; 8037000, 0; 8040600, 0; 8044200, 0; 8047800, 0; 8051400, 0; 8055000, 16744.18788; 8058600, 46250.49133; 8062200, 78062.41871; 8065800, 107065.3817; 8069400, 126496.7581; 8073000, 129135.2362; 8076600, 121986.9901; 8080200, 120225.2986; 8083800, 96629.18608; 8087400, 67074.17226; 8091000, 42231.88624; 8094600, 13854.03957; 8098200, 0; 8101800, 0; 8105400, 0; 8109000, 0; 8112600, 0; 8116200, 0; 8119800, 0; 8123400, 0; 8127000, 0; 8130600, 0; 8134200, 0; 8137800, 0; 8141400, 16744.18788; 8145000, 46250.49133; 8148600, 78062.41871; 8152200, 107065.3817; 8155800, 126496.7581; 8159400, 129135.2362; 8163000, 121986.9901; 8166600, 120225.2986; 8170200, 96629.18608; 8173800, 67074.17226; 8177400, 42231.88624; 8181000, 13854.03957; 8184600, 0; 8188200, 0; 8191800, 0; 8195400, 0; 8199000, 0; 8202600, 0; 8206200, 0; 8209800, 0; 8213400, 0; 8217000, 0; 8220600, 0; 8224200, 0; 8227800, 16744.18788; 8231400, 46250.49133; 8235000, 78062.41871; 8238600, 107065.3817; 8242200, 126496.7581; 8245800, 129135.2362; 8249400, 121986.9901; 8253000, 120225.2986; 8256600, 96629.18608; 8260200, 67074.17226; 8263800, 42231.88624; 8267400, 13854.03957; 8271000, 0; 8274600, 0; 8278200, 0; 8281800, 0; 8285400, 0; 8289000, 0; 8292600, 0; 8296200, 0; 8299800, 0; 8303400, 0; 8307000, 0; 8310600, 0; 8314200, 16744.18788; 8317800, 46250.49133; 8321400, 78062.41871; 8325000, 107065.3817; 8328600, 126496.7581; 8332200, 129135.2362; 8335800, 121986.9901; 8339400, 120225.2986; 8343000, 96629.18608; 8346600, 67074.17226; 8350200, 42231.88624; 8353800, 13854.03957; 8357400, 0; 8361000, 0; 8364600, 0; 8368200, 0; 8371800, 0; 8375400, 0; 8379000, 0; 8382600, 0; 8386200, 0; 8389800, 0; 8393400, 0; 8397000, 0; 8400600, 16744.18788; 8404200, 46250.49133; 8407800, 78062.41871; 8411400, 107065.3817; 8415000, 126496.7581; 8418600, 129135.2362; 8422200, 121986.9901; 8425800, 120225.2986; 8429400, 96629.18608; 8433000, 67074.17226; 8436600, 42231.88624; 8440200, 13854.03957; 8443800, 0; 8447400, 0; 8451000, 0; 8454600, 0; 8458200, 0; 8461800, 0; 8465400, 0; 8469000, 0; 8472600, 0; 8476200, 0; 8479800, 0; 8483400, 0; 8487000, 16744.18788; 8490600, 46250.49133; 8494200, 78062.41871; 8497800, 107065.3817; 8501400, 126496.7581; 8505000, 129135.2362; 8508600, 121986.9901; 8512200, 120225.2986; 8515800, 96629.18608; 8519400, 67074.17226; 8523000, 42231.88624; 8526600, 13854.03957; 8530200, 0; 8533800, 0; 8537400, 0; 8541000, 0; 8544600, 0; 8548200, 0; 8551800, 0; 8555400, 0; 8559000, 0; 8562600, 0; 8566200, 0; 8569800, 0; 8573400, 16744.18788; 8577000, 46250.49133; 8580600, 78062.41871; 8584200, 107065.3817; 8587800, 126496.7581; 8591400, 129135.2362; 8595000, 121986.9901; 8598600, 120225.2986; 8602200, 96629.18608; 8605800, 67074.17226; 8609400, 42231.88624; 8613000, 13854.03957; 8616600, 0; 8620200, 0; 8623800, 0; 8627400, 0; 8631000, 0; 8634600, 0; 8638200, 0; 8641800, 0; 8645400, 0; 8649000, 0; 8652600, 0; 8656200, 0; 8659800, 16744.18788; 8663400, 46250.49133; 8667000, 78062.41871; 8670600, 107065.3817; 8674200, 126496.7581; 8677800, 129135.2362; 8681400, 121986.9901; 8685000, 120225.2986; 8688600, 96629.18608; 8692200, 67074.17226; 8695800, 42231.88624; 8699400, 13854.03957; 8703000, 0; 8706600, 0; 8710200, 0; 8713800, 0; 8717400, 0; 8721000, 0; 8724600, 0; 8728200, 0; 8731800, 0; 8735400, 0; 8739000, 0; 8742600, 0; 8746200, 16744.18788; 8749800, 46250.49133; 8753400, 78062.41871; 8757000, 107065.3817; 8760600, 126496.7581; 8764200, 129135.2362; 8767800, 121986.9901; 8771400, 120225.2986; 8775000, 96629.18608; 8778600, 67074.17226; 8782200, 42231.88624; 8785800, 13854.03957; 8789400, 0; 8793000, 0; 8796600, 0; 8800200, 0; 8803800, 0; 8807400, 0; 8811000, 0; 8814600, 0; 8818200, 0; 8821800, 0; 8825400, 0; 8829000, 0; 8832600, 16744.18788; 8836200, 46250.49133; 8839800, 78062.41871; 8843400, 107065.3817; 8847000, 126496.7581; 8850600, 129135.2362; 8854200, 121986.9901; 8857800, 120225.2986; 8861400, 96629.18608; 8865000, 67074.17226; 8868600, 42231.88624; 8872200, 13854.03957; 8875800, 0; 8879400, 0; 8883000, 0; 8886600, 0; 8890200, 0; 8893800, 0; 8897400, 0; 8901000, 0; 8904600, 0; 8908200, 0; 8911800, 0; 8915400, 0; 8919000, 16744.18788; 8922600, 46250.49133; 8926200, 78062.41871; 8929800, 107065.3817; 8933400, 126496.7581; 8937000, 129135.2362; 8940600, 121986.9901; 8944200, 120225.2986; 8947800, 96629.18608; 8951400, 67074.17226; 8955000, 42231.88624; 8958600, 13854.03957; 8962200, 0; 8965800, 0; 8969400, 0; 8973000, 0; 8976600, 0; 8980200, 0; 8983800, 0; 8987400, 0; 8991000, 0; 8994600, 0; 8998200, 0; 9001800, 0; 9005400, 16744.18788; 9009000, 46250.49133; 9012600, 78062.41871; 9016200, 107065.3817; 9019800, 126496.7581; 9023400, 129135.2362; 9027000, 121986.9901; 9030600, 120225.2986; 9034200, 96629.18608; 9037800, 67074.17226; 9041400, 42231.88624; 9045000, 13854.03957; 9048600, 0; 9052200, 0; 9055800, 0; 9059400, 0; 9063000, 0; 9066600, 0; 9070200, 0; 9073800, 0; 9077400, 0; 9081000, 0; 9084600, 0; 9088200, 0; 9091800, 16744.18788; 9095400, 46250.49133; 9099000, 78062.41871; 9102600, 107065.3817; 9106200, 126496.7581; 9109800, 129135.2362; 9113400, 121986.9901; 9117000, 120225.2986; 9120600, 96629.18608; 9124200, 67074.17226; 9127800, 42231.88624; 9131400, 13854.03957; 9135000, 0; 9138600, 0; 9142200, 0; 9145800, 0; 9149400, 0; 9153000, 0; 9156600, 0; 9160200, 0; 9163800, 0; 9167400, 0; 9171000, 0; 9174600, 0; 9178200, 16744.18788; 9181800, 46250.49133; 9185400, 78062.41871; 9189000, 107065.3817; 9192600, 126496.7581; 9196200, 129135.2362; 9199800, 121986.9901; 9203400, 120225.2986; 9207000, 96629.18608; 9210600, 67074.17226; 9214200, 42231.88624; 9217800, 13854.03957; 9221400, 0; 9225000, 0; 9228600, 0; 9232200, 0; 9235800, 0; 9239400, 0; 9243000, 0; 9246600, 0; 9250200, 0; 9253800, 0; 9257400, 0; 9261000, 0; 9264600, 16744.18788; 9268200, 46250.49133; 9271800, 78062.41871; 9275400, 107065.3817; 9279000, 126496.7581; 9282600, 129135.2362; 9286200, 121986.9901; 9289800, 120225.2986; 9293400, 96629.18608; 9297000, 67074.17226; 9300600, 42231.88624; 9304200, 13854.03957; 9307800, 0; 9311400, 0; 9315000, 0; 9318600, 0; 9322200, 0; 9325800, 0; 9329400, 0; 9333000, 0; 9336600, 0; 9340200, 0; 9343800, 0; 9347400, 0; 9351000, 16744.18788; 9354600, 46250.49133; 9358200, 78062.41871; 9361800, 107065.3817; 9365400, 126496.7581; 9369000, 129135.2362; 9372600, 121986.9901; 9376200, 120225.2986; 9379800, 96629.18608; 9383400, 67074.17226; 9387000, 42231.88624; 9390600, 13854.03957; 9394200, 0; 9397800, 0; 9401400, 0; 9405000, 0; 9408600, 0; 9412200, 0; 9415800, 0; 9419400, 0; 9423000, 0; 9426600, 0; 9430200, 0; 9433800, 0; 9437400, 16744.18788; 9441000, 46250.49133; 9444600, 78062.41871; 9448200, 107065.3817; 9451800, 126496.7581; 9455400, 129135.2362; 9459000, 121986.9901; 9462600, 120225.2986; 9466200, 96629.18608; 9469800, 67074.17226; 9473400, 42231.88624; 9477000, 13854.03957; 9480600, 0; 9484200, 0; 9487800, 0; 9491400, 0; 9495000, 0; 9498600, 0; 9502200, 0; 9505800, 0; 9509400, 0; 9513000, 0; 9516600, 0; 9520200, 0; 9523800, 16744.18788; 9527400, 46250.49133; 9531000, 78062.41871; 9534600, 107065.3817; 9538200, 126496.7581; 9541800, 129135.2362; 9545400, 121986.9901; 9549000, 120225.2986; 9552600, 96629.18608; 9556200, 67074.17226; 9559800, 42231.88624; 9563400, 13854.03957; 9567000, 0; 9570600, 0; 9574200, 0; 9577800, 0; 9581400, 0; 9585000, 0; 9588600, 0; 9592200, 0; 9595800, 0; 9599400, 0; 9603000, 0; 9606600, 0; 9610200, 16744.18788; 9613800, 46250.49133; 9617400, 78062.41871; 9621000, 107065.3817; 9624600, 126496.7581; 9628200, 129135.2362; 9631800, 121986.9901; 9635400, 120225.2986; 9639000, 96629.18608; 9642600, 67074.17226; 9646200, 42231.88624; 9649800, 13854.03957; 9653400, 0; 9657000, 0; 9660600, 0; 9664200, 0; 9667800, 0; 9671400, 0; 9675000, 0; 9678600, 0; 9682200, 0; 9685800, 0; 9689400, 0; 9693000, 0; 9696600, 16744.18788; 9700200, 46250.49133; 9703800, 78062.41871; 9707400, 107065.3817; 9711000, 126496.7581; 9714600, 129135.2362; 9718200, 121986.9901; 9721800, 120225.2986; 9725400, 96629.18608; 9729000, 67074.17226; 9732600, 42231.88624; 9736200, 13854.03957; 9739800, 0; 9743400, 0; 9747000, 0; 9750600, 0; 9754200, 0; 9757800, 0; 9761400, 0; 9765000, 0; 9768600, 0; 9772200, 0; 9775800, 0; 9779400, 0; 9783000, 16744.18788; 9786600, 46250.49133; 9790200, 78062.41871; 9793800, 107065.3817; 9797400, 126496.7581; 9801000, 129135.2362; 9804600, 121986.9901; 9808200, 120225.2986; 9811800, 96629.18608; 9815400, 67074.17226; 9819000, 42231.88624; 9822600, 13854.03957; 9826200, 0; 9829800, 0; 9833400, 0; 9837000, 0; 9840600, 0; 9844200, 0; 9847800, 0; 9851400, 0; 9855000, 0; 9858600, 0; 9862200, 0; 9865800, 0; 9869400, 16744.18788; 9873000, 46250.49133; 9876600, 78062.41871; 9880200, 107065.3817; 9883800, 126496.7581; 9887400, 129135.2362; 9891000, 121986.9901; 9894600, 120225.2986; 9898200, 96629.18608; 9901800, 67074.17226; 9905400, 42231.88624; 9909000, 13854.03957; 9912600, 0; 9916200, 0; 9919800, 0; 9923400, 0; 9927000, 0; 9930600, 0; 9934200, 0; 9937800, 0; 9941400, 0; 9945000, 0; 9948600, 0; 9952200, 0; 9955800, 16744.18788; 9959400, 46250.49133; 9963000, 78062.41871; 9966600, 107065.3817; 9970200, 126496.7581; 9973800, 129135.2362; 9977400, 121986.9901; 9981000, 120225.2986; 9984600, 96629.18608; 9988200, 67074.17226; 9991800, 42231.88624; 9995400, 13854.03957; 9999000, 0; 10002600, 0; 10006200, 0; 10009800, 0; 10013400, 0; 10017000, 0; 10020600, 0; 10024200, 0; 10027800, 0; 10031400, 0; 10035000, 0; 10038600, 0; 10042200, 16744.18788; 10045800, 46250.49133; 10049400, 78062.41871; 10053000, 107065.3817; 10056600, 126496.7581; 10060200, 129135.2362; 10063800, 121986.9901; 10067400, 120225.2986; 10071000, 96629.18608; 10074600, 67074.17226; 10078200, 42231.88624; 10081800, 13854.03957; 10085400, 0; 10089000, 0; 10092600, 0; 10096200, 0; 10099800, 0; 10103400, 0; 10107000, 0; 10110600, 0; 10114200, 0; 10117800, 0; 10121400, 0; 10125000, 0; 10128600, 16744.18788; 10132200, 46250.49133; 10135800, 78062.41871; 10139400, 107065.3817; 10143000, 126496.7581; 10146600, 129135.2362; 10150200, 121986.9901; 10153800, 120225.2986; 10157400, 96629.18608; 10161000, 67074.17226; 10164600, 42231.88624; 10168200, 13854.03957; 10171800, 0; 10175400, 0; 10179000, 0; 10182600, 0; 10186200, 0; 10189800, 0; 10193400, 0; 10197000, 0; 10200600, 0; 10204200, 0; 10207800, 0; 10211400, 0; 10215000, 16744.18788; 10218600, 46250.49133; 10222200, 78062.41871; 10225800, 107065.3817; 10229400, 126496.7581; 10233000, 129135.2362; 10236600, 121986.9901; 10240200, 120225.2986; 10243800, 96629.18608; 10247400, 67074.17226; 10251000, 42231.88624; 10254600, 13854.03957; 10258200, 0; 10261800, 0; 10265400, 0; 10269000, 0; 10272600, 0; 10276200, 0; 10279800, 0; 10283400, 0; 10287000, 0; 10290600, 0; 10294200, 0; 10297800, 0; 10301400, 16744.18788; 10305000, 46250.49133; 10308600, 78062.41871; 10312200, 107065.3817; 10315800, 126496.7581; 10319400, 129135.2362; 10323000, 121986.9901; 10326600, 120225.2986; 10330200, 96629.18608; 10333800, 67074.17226; 10337400, 42231.88624; 10341000, 13854.03957; 10344600, 0; 10348200, 0; 10351800, 0; 10355400, 0; 10359000, 0; 10362600, 0; 10366200, 0; 10369800, 0; 10373400, 0; 10377000, 0; 10380600, 0; 10384200, 1661.651959; 10387800, 30396.83891; 10391400, 66352.15895; 10395000, 96560.44161; 10398600, 118358.3298; 10402200, 137716.7716; 10405800, 124537.4753; 10409400, 126411.253; 10413000, 118260.1235; 10416600, 106141.4561; 10420200, 77795.16641; 10423800, 48435.38689; 10427400, 19833.76062; 10431000, 121.7759119; 10434600, 0; 10438200, 0; 10441800, 0; 10445400, 0; 10449000, 0; 10452600, 0; 10456200, 0; 10459800, 0; 10463400, 0; 10467000, 0; 10470600, 1661.651959; 10474200, 30396.83891; 10477800, 66352.15895; 10481400, 96560.44161; 10485000, 118358.3298; 10488600, 137716.7716; 10492200, 124537.4753; 10495800, 126411.253; 10499400, 118260.1235; 10503000, 106141.4561; 10506600, 77795.16641; 10510200, 48435.38689; 10513800, 19833.76062; 10517400, 121.7759119; 10521000, 0; 10524600, 0; 10528200, 0; 10531800, 0; 10535400, 0; 10539000, 0; 10542600, 0; 10546200, 0; 10549800, 0; 10553400, 0; 10557000, 1661.651959; 10560600, 30396.83891; 10564200, 66352.15895; 10567800, 96560.44161; 10571400, 118358.3298; 10575000, 137716.7716; 10578600, 124537.4753; 10582200, 126411.253; 10585800, 118260.1235; 10589400, 106141.4561; 10593000, 77795.16641; 10596600, 48435.38689; 10600200, 19833.76062; 10603800, 121.7759119; 10607400, 0; 10611000, 0; 10614600, 0; 10618200, 0; 10621800, 0; 10625400, 0; 10629000, 0; 10632600, 0; 10636200, 0; 10639800, 0; 10643400, 1661.651959; 10647000, 30396.83891; 10650600, 66352.15895; 10654200, 96560.44161; 10657800, 118358.3298; 10661400, 137716.7716; 10665000, 124537.4753; 10668600, 126411.253; 10672200, 118260.1235; 10675800, 106141.4561; 10679400, 77795.16641; 10683000, 48435.38689; 10686600, 19833.76062; 10690200, 121.7759119; 10693800, 0; 10697400, 0; 10701000, 0; 10704600, 0; 10708200, 0; 10711800, 0; 10715400, 0; 10719000, 0; 10722600, 0; 10726200, 0; 10729800, 1661.651959; 10733400, 30396.83891; 10737000, 66352.15895; 10740600, 96560.44161; 10744200, 118358.3298; 10747800, 137716.7716; 10751400, 124537.4753; 10755000, 126411.253; 10758600, 118260.1235; 10762200, 106141.4561; 10765800, 77795.16641; 10769400, 48435.38689; 10773000, 19833.76062; 10776600, 121.7759119; 10780200, 0; 10783800, 0; 10787400, 0; 10791000, 0; 10794600, 0; 10798200, 0; 10801800, 0; 10805400, 0; 10809000, 0; 10812600, 0; 10816200, 1661.651959; 10819800, 30396.83891; 10823400, 66352.15895; 10827000, 96560.44161; 10830600, 118358.3298; 10834200, 137716.7716; 10837800, 124537.4753; 10841400, 126411.253; 10845000, 118260.1235; 10848600, 106141.4561; 10852200, 77795.16641; 10855800, 48435.38689; 10859400, 19833.76062; 10863000, 121.7759119; 10866600, 0; 10870200, 0; 10873800, 0; 10877400, 0; 10881000, 0; 10884600, 0; 10888200, 0; 10891800, 0; 10895400, 0; 10899000, 0; 10902600, 1661.651959; 10906200, 30396.83891; 10909800, 66352.15895; 10913400, 96560.44161; 10917000, 118358.3298; 10920600, 137716.7716; 10924200, 124537.4753; 10927800, 126411.253; 10931400, 118260.1235; 10935000, 106141.4561; 10938600, 77795.16641; 10942200, 48435.38689; 10945800, 19833.76062; 10949400, 121.7759119; 10953000, 0; 10956600, 0; 10960200, 0; 10963800, 0; 10967400, 0; 10971000, 0; 10974600, 0; 10978200, 0; 10981800, 0; 10985400, 0; 10989000, 1661.651959; 10992600, 30396.83891; 10996200, 66352.15895; 10999800, 96560.44161; 11003400, 118358.3298; 11007000, 137716.7716; 11010600, 124537.4753; 11014200, 126411.253; 11017800, 118260.1235; 11021400, 106141.4561; 11025000, 77795.16641; 11028600, 48435.38689; 11032200, 19833.76062; 11035800, 121.7759119; 11039400, 0; 11043000, 0; 11046600, 0; 11050200, 0; 11053800, 0; 11057400, 0; 11061000, 0; 11064600, 0; 11068200, 0; 11071800, 0; 11075400, 1661.651959; 11079000, 30396.83891; 11082600, 66352.15895; 11086200, 96560.44161; 11089800, 118358.3298; 11093400, 137716.7716; 11097000, 124537.4753; 11100600, 126411.253; 11104200, 118260.1235; 11107800, 106141.4561; 11111400, 77795.16641; 11115000, 48435.38689; 11118600, 19833.76062; 11122200, 121.7759119; 11125800, 0; 11129400, 0; 11133000, 0; 11136600, 0; 11140200, 0; 11143800, 0; 11147400, 0; 11151000, 0; 11154600, 0; 11158200, 0; 11161800, 1661.651959; 11165400, 30396.83891; 11169000, 66352.15895; 11172600, 96560.44161; 11176200, 118358.3298; 11179800, 137716.7716; 11183400, 124537.4753; 11187000, 126411.253; 11190600, 118260.1235; 11194200, 106141.4561; 11197800, 77795.16641; 11201400, 48435.38689; 11205000, 19833.76062; 11208600, 121.7759119; 11212200, 0; 11215800, 0; 11219400, 0; 11223000, 0; 11226600, 0; 11230200, 0; 11233800, 0; 11237400, 0; 11241000, 0; 11244600, 0; 11248200, 1661.651959; 11251800, 30396.83891; 11255400, 66352.15895; 11259000, 96560.44161; 11262600, 118358.3298; 11266200, 137716.7716; 11269800, 124537.4753; 11273400, 126411.253; 11277000, 118260.1235; 11280600, 106141.4561; 11284200, 77795.16641; 11287800, 48435.38689; 11291400, 19833.76062; 11295000, 121.7759119; 11298600, 0; 11302200, 0; 11305800, 0; 11309400, 0; 11313000, 0; 11316600, 0; 11320200, 0; 11323800, 0; 11327400, 0; 11331000, 0; 11334600, 1661.651959; 11338200, 30396.83891; 11341800, 66352.15895; 11345400, 96560.44161; 11349000, 118358.3298; 11352600, 137716.7716; 11356200, 124537.4753; 11359800, 126411.253; 11363400, 118260.1235; 11367000, 106141.4561; 11370600, 77795.16641; 11374200, 48435.38689; 11377800, 19833.76062; 11381400, 121.7759119; 11385000, 0; 11388600, 0; 11392200, 0; 11395800, 0; 11399400, 0; 11403000, 0; 11406600, 0; 11410200, 0; 11413800, 0; 11417400, 0; 11421000, 1661.651959; 11424600, 30396.83891; 11428200, 66352.15895; 11431800, 96560.44161; 11435400, 118358.3298; 11439000, 137716.7716; 11442600, 124537.4753; 11446200, 126411.253; 11449800, 118260.1235; 11453400, 106141.4561; 11457000, 77795.16641; 11460600, 48435.38689; 11464200, 19833.76062; 11467800, 121.7759119; 11471400, 0; 11475000, 0; 11478600, 0; 11482200, 0; 11485800, 0; 11489400, 0; 11493000, 0; 11496600, 0; 11500200, 0; 11503800, 0; 11507400, 1661.651959; 11511000, 30396.83891; 11514600, 66352.15895; 11518200, 96560.44161; 11521800, 118358.3298; 11525400, 137716.7716; 11529000, 124537.4753; 11532600, 126411.253; 11536200, 118260.1235; 11539800, 106141.4561; 11543400, 77795.16641; 11547000, 48435.38689; 11550600, 19833.76062; 11554200, 121.7759119; 11557800, 0; 11561400, 0; 11565000, 0; 11568600, 0; 11572200, 0; 11575800, 0; 11579400, 0; 11583000, 0; 11586600, 0; 11590200, 0; 11593800, 1661.651959; 11597400, 30396.83891; 11601000, 66352.15895; 11604600, 96560.44161; 11608200, 118358.3298; 11611800, 137716.7716; 11615400, 124537.4753; 11619000, 126411.253; 11622600, 118260.1235; 11626200, 106141.4561; 11629800, 77795.16641; 11633400, 48435.38689; 11637000, 19833.76062; 11640600, 121.7759119; 11644200, 0; 11647800, 0; 11651400, 0; 11655000, 0; 11658600, 0; 11662200, 0; 11665800, 0; 11669400, 0; 11673000, 0; 11676600, 0; 11680200, 1661.651959; 11683800, 30396.83891; 11687400, 66352.15895; 11691000, 96560.44161; 11694600, 118358.3298; 11698200, 137716.7716; 11701800, 124537.4753; 11705400, 126411.253; 11709000, 118260.1235; 11712600, 106141.4561; 11716200, 77795.16641; 11719800, 48435.38689; 11723400, 19833.76062; 11727000, 121.7759119; 11730600, 0; 11734200, 0; 11737800, 0; 11741400, 0; 11745000, 0; 11748600, 0; 11752200, 0; 11755800, 0; 11759400, 0; 11763000, 0; 11766600, 1661.651959; 11770200, 30396.83891; 11773800, 66352.15895; 11777400, 96560.44161; 11781000, 118358.3298; 11784600, 137716.7716; 11788200, 124537.4753; 11791800, 126411.253; 11795400, 118260.1235; 11799000, 106141.4561; 11802600, 77795.16641; 11806200, 48435.38689; 11809800, 19833.76062; 11813400, 121.7759119; 11817000, 0; 11820600, 0; 11824200, 0; 11827800, 0; 11831400, 0; 11835000, 0; 11838600, 0; 11842200, 0; 11845800, 0; 11849400, 0; 11853000, 1661.651959; 11856600, 30396.83891; 11860200, 66352.15895; 11863800, 96560.44161; 11867400, 118358.3298; 11871000, 137716.7716; 11874600, 124537.4753; 11878200, 126411.253; 11881800, 118260.1235; 11885400, 106141.4561; 11889000, 77795.16641; 11892600, 48435.38689; 11896200, 19833.76062; 11899800, 121.7759119; 11903400, 0; 11907000, 0; 11910600, 0; 11914200, 0; 11917800, 0; 11921400, 0; 11925000, 0; 11928600, 0; 11932200, 0; 11935800, 0; 11939400, 1661.651959; 11943000, 30396.83891; 11946600, 66352.15895; 11950200, 96560.44161; 11953800, 118358.3298; 11957400, 137716.7716; 11961000, 124537.4753; 11964600, 126411.253; 11968200, 118260.1235; 11971800, 106141.4561; 11975400, 77795.16641; 11979000, 48435.38689; 11982600, 19833.76062; 11986200, 121.7759119; 11989800, 0; 11993400, 0; 11997000, 0; 12000600, 0; 12004200, 0; 12007800, 0; 12011400, 0; 12015000, 0; 12018600, 0; 12022200, 0; 12025800, 1661.651959; 12029400, 30396.83891; 12033000, 66352.15895; 12036600, 96560.44161; 12040200, 118358.3298; 12043800, 137716.7716; 12047400, 124537.4753; 12051000, 126411.253; 12054600, 118260.1235; 12058200, 106141.4561; 12061800, 77795.16641; 12065400, 48435.38689; 12069000, 19833.76062; 12072600, 121.7759119; 12076200, 0; 12079800, 0; 12083400, 0; 12087000, 0; 12090600, 0; 12094200, 0; 12097800, 0; 12101400, 0; 12105000, 0; 12108600, 0; 12112200, 1661.651959; 12115800, 30396.83891; 12119400, 66352.15895; 12123000, 96560.44161; 12126600, 118358.3298; 12130200, 137716.7716; 12133800, 124537.4753; 12137400, 126411.253; 12141000, 118260.1235; 12144600, 106141.4561; 12148200, 77795.16641; 12151800, 48435.38689; 12155400, 19833.76062; 12159000, 121.7759119; 12162600, 0; 12166200, 0; 12169800, 0; 12173400, 0; 12177000, 0; 12180600, 0; 12184200, 0; 12187800, 0; 12191400, 0; 12195000, 0; 12198600, 1661.651959; 12202200, 30396.83891; 12205800, 66352.15895; 12209400, 96560.44161; 12213000, 118358.3298; 12216600, 137716.7716; 12220200, 124537.4753; 12223800, 126411.253; 12227400, 118260.1235; 12231000, 106141.4561; 12234600, 77795.16641; 12238200, 48435.38689; 12241800, 19833.76062; 12245400, 121.7759119; 12249000, 0; 12252600, 0; 12256200, 0; 12259800, 0; 12263400, 0; 12267000, 0; 12270600, 0; 12274200, 0; 12277800, 0; 12281400, 0; 12285000, 1661.651959; 12288600, 30396.83891; 12292200, 66352.15895; 12295800, 96560.44161; 12299400, 118358.3298; 12303000, 137716.7716; 12306600, 124537.4753; 12310200, 126411.253; 12313800, 118260.1235; 12317400, 106141.4561; 12321000, 77795.16641; 12324600, 48435.38689; 12328200, 19833.76062; 12331800, 121.7759119; 12335400, 0; 12339000, 0; 12342600, 0; 12346200, 0; 12349800, 0; 12353400, 0; 12357000, 0; 12360600, 0; 12364200, 0; 12367800, 0; 12371400, 1661.651959; 12375000, 30396.83891; 12378600, 66352.15895; 12382200, 96560.44161; 12385800, 118358.3298; 12389400, 137716.7716; 12393000, 124537.4753; 12396600, 126411.253; 12400200, 118260.1235; 12403800, 106141.4561; 12407400, 77795.16641; 12411000, 48435.38689; 12414600, 19833.76062; 12418200, 121.7759119; 12421800, 0; 12425400, 0; 12429000, 0; 12432600, 0; 12436200, 0; 12439800, 0; 12443400, 0; 12447000, 0; 12450600, 0; 12454200, 0; 12457800, 1661.651959; 12461400, 30396.83891; 12465000, 66352.15895; 12468600, 96560.44161; 12472200, 118358.3298; 12475800, 137716.7716; 12479400, 124537.4753; 12483000, 126411.253; 12486600, 118260.1235; 12490200, 106141.4561; 12493800, 77795.16641; 12497400, 48435.38689; 12501000, 19833.76062; 12504600, 121.7759119; 12508200, 0; 12511800, 0; 12515400, 0; 12519000, 0; 12522600, 0; 12526200, 0; 12529800, 0; 12533400, 0; 12537000, 0; 12540600, 0; 12544200, 1661.651959; 12547800, 30396.83891; 12551400, 66352.15895; 12555000, 96560.44161; 12558600, 118358.3298; 12562200, 137716.7716; 12565800, 124537.4753; 12569400, 126411.253; 12573000, 118260.1235; 12576600, 106141.4561; 12580200, 77795.16641; 12583800, 48435.38689; 12587400, 19833.76062; 12591000, 121.7759119; 12594600, 0; 12598200, 0; 12601800, 0; 12605400, 0; 12609000, 0; 12612600, 0; 12616200, 0; 12619800, 0; 12623400, 0; 12627000, 0; 12630600, 1661.651959; 12634200, 30396.83891; 12637800, 66352.15895; 12641400, 96560.44161; 12645000, 118358.3298; 12648600, 137716.7716; 12652200, 124537.4753; 12655800, 126411.253; 12659400, 118260.1235; 12663000, 106141.4561; 12666600, 77795.16641; 12670200, 48435.38689; 12673800, 19833.76062; 12677400, 121.7759119; 12681000, 0; 12684600, 0; 12688200, 0; 12691800, 0; 12695400, 0; 12699000, 0; 12702600, 0; 12706200, 0; 12709800, 0; 12713400, 0; 12717000, 1661.651959; 12720600, 30396.83891; 12724200, 66352.15895; 12727800, 96560.44161; 12731400, 118358.3298; 12735000, 137716.7716; 12738600, 124537.4753; 12742200, 126411.253; 12745800, 118260.1235; 12749400, 106141.4561; 12753000, 77795.16641; 12756600, 48435.38689; 12760200, 19833.76062; 12763800, 121.7759119; 12767400, 0; 12771000, 0; 12774600, 0; 12778200, 0; 12781800, 0; 12785400, 0; 12789000, 0; 12792600, 0; 12796200, 0; 12799800, 0; 12803400, 1661.651959; 12807000, 30396.83891; 12810600, 66352.15895; 12814200, 96560.44161; 12817800, 118358.3298; 12821400, 137716.7716; 12825000, 124537.4753; 12828600, 126411.253; 12832200, 118260.1235; 12835800, 106141.4561; 12839400, 77795.16641; 12843000, 48435.38689; 12846600, 19833.76062; 12850200, 121.7759119; 12853800, 0; 12857400, 0; 12861000, 0; 12864600, 0; 12868200, 0; 12871800, 0; 12875400, 0; 12879000, 0; 12882600, 0; 12886200, 0; 12889800, 1661.651959; 12893400, 30396.83891; 12897000, 66352.15895; 12900600, 96560.44161; 12904200, 118358.3298; 12907800, 137716.7716; 12911400, 124537.4753; 12915000, 126411.253; 12918600, 118260.1235; 12922200, 106141.4561; 12925800, 77795.16641; 12929400, 48435.38689; 12933000, 19833.76062; 12936600, 121.7759119; 12940200, 0; 12943800, 0; 12947400, 0; 12951000, 0; 12954600, 0; 12958200, 0; 12961800, 0; 12965400, 0; 12969000, 0; 12972600, 0; 12976200, 1661.651959; 12979800, 30396.83891; 12983400, 66352.15895; 12987000, 96560.44161; 12990600, 118358.3298; 12994200, 137716.7716; 12997800, 124537.4753; 13001400, 126411.253; 13005000, 118260.1235; 13008600, 106141.4561; 13012200, 77795.16641; 13015800, 48435.38689; 13019400, 19833.76062; 13023000, 121.7759119; 13026600, 0; 13030200, 0; 13033800, 0; 13037400, 0; 13041000, 0; 13044600, 0; 13048200, 0; 13051800, 0; 13055400, 0; 13059000, 0; 13062600, 5228.245817; 13066200, 27030.19324; 13069800, 55192.90246; 13073400, 75667.49244; 13077000, 96824.02753; 13080600, 111737.5175; 13084200, 119429.696; 13087800, 120886.9477; 13091400, 108096.4178; 13095000, 90682.46238; 13098600, 69672.05838; 13102200, 48868.67344; 13105800, 24623.08938; 13109400, 2143.256049; 13113000, 0; 13116600, 0; 13120200, 0; 13123800, 0; 13127400, 0; 13131000, 0; 13134600, 0; 13138200, 0; 13141800, 0; 13145400, 0; 13149000, 5228.245817; 13152600, 27030.19324; 13156200, 55192.90246; 13159800, 75667.49244; 13163400, 96824.02753; 13167000, 111737.5175; 13170600, 119429.696; 13174200, 120886.9477; 13177800, 108096.4178; 13181400, 90682.46238; 13185000, 69672.05838; 13188600, 48868.67344; 13192200, 24623.08938; 13195800, 2143.256049; 13199400, 0; 13203000, 0; 13206600, 0; 13210200, 0; 13213800, 0; 13217400, 0; 13221000, 0; 13224600, 0; 13228200, 0; 13231800, 0; 13235400, 5228.245817; 13239000, 27030.19324; 13242600, 55192.90246; 13246200, 75667.49244; 13249800, 96824.02753; 13253400, 111737.5175; 13257000, 119429.696; 13260600, 120886.9477; 13264200, 108096.4178; 13267800, 90682.46238; 13271400, 69672.05838; 13275000, 48868.67344; 13278600, 24623.08938; 13282200, 2143.256049; 13285800, 0; 13289400, 0; 13293000, 0; 13296600, 0; 13300200, 0; 13303800, 0; 13307400, 0; 13311000, 0; 13314600, 0; 13318200, 0; 13321800, 5228.245817; 13325400, 27030.19324; 13329000, 55192.90246; 13332600, 75667.49244; 13336200, 96824.02753; 13339800, 111737.5175; 13343400, 119429.696; 13347000, 120886.9477; 13350600, 108096.4178; 13354200, 90682.46238; 13357800, 69672.05838; 13361400, 48868.67344; 13365000, 24623.08938; 13368600, 2143.256049; 13372200, 0; 13375800, 0; 13379400, 0; 13383000, 0; 13386600, 0; 13390200, 0; 13393800, 0; 13397400, 0; 13401000, 0; 13404600, 0; 13408200, 5228.245817; 13411800, 27030.19324; 13415400, 55192.90246; 13419000, 75667.49244; 13422600, 96824.02753; 13426200, 111737.5175; 13429800, 119429.696; 13433400, 120886.9477; 13437000, 108096.4178; 13440600, 90682.46238; 13444200, 69672.05838; 13447800, 48868.67344; 13451400, 24623.08938; 13455000, 2143.256049; 13458600, 0; 13462200, 0; 13465800, 0; 13469400, 0; 13473000, 0; 13476600, 0; 13480200, 0; 13483800, 0; 13487400, 0; 13491000, 0; 13494600, 5228.245817; 13498200, 27030.19324; 13501800, 55192.90246; 13505400, 75667.49244; 13509000, 96824.02753; 13512600, 111737.5175; 13516200, 119429.696; 13519800, 120886.9477; 13523400, 108096.4178; 13527000, 90682.46238; 13530600, 69672.05838; 13534200, 48868.67344; 13537800, 24623.08938; 13541400, 2143.256049; 13545000, 0; 13548600, 0; 13552200, 0; 13555800, 0; 13559400, 0; 13563000, 0; 13566600, 0; 13570200, 0; 13573800, 0; 13577400, 0; 13581000, 5228.245817; 13584600, 27030.19324; 13588200, 55192.90246; 13591800, 75667.49244; 13595400, 96824.02753; 13599000, 111737.5175; 13602600, 119429.696; 13606200, 120886.9477; 13609800, 108096.4178; 13613400, 90682.46238; 13617000, 69672.05838; 13620600, 48868.67344; 13624200, 24623.08938; 13627800, 2143.256049; 13631400, 0; 13635000, 0; 13638600, 0; 13642200, 0; 13645800, 0; 13649400, 0; 13653000, 0; 13656600, 0; 13660200, 0; 13663800, 0; 13667400, 5228.245817; 13671000, 27030.19324; 13674600, 55192.90246; 13678200, 75667.49244; 13681800, 96824.02753; 13685400, 111737.5175; 13689000, 119429.696; 13692600, 120886.9477; 13696200, 108096.4178; 13699800, 90682.46238; 13703400, 69672.05838; 13707000, 48868.67344; 13710600, 24623.08938; 13714200, 2143.256049; 13717800, 0; 13721400, 0; 13725000, 0; 13728600, 0; 13732200, 0; 13735800, 0; 13739400, 0; 13743000, 0; 13746600, 0; 13750200, 0; 13753800, 5228.245817; 13757400, 27030.19324; 13761000, 55192.90246; 13764600, 75667.49244; 13768200, 96824.02753; 13771800, 111737.5175; 13775400, 119429.696; 13779000, 120886.9477; 13782600, 108096.4178; 13786200, 90682.46238; 13789800, 69672.05838; 13793400, 48868.67344; 13797000, 24623.08938; 13800600, 2143.256049; 13804200, 0; 13807800, 0; 13811400, 0; 13815000, 0; 13818600, 0; 13822200, 0; 13825800, 0; 13829400, 0; 13833000, 0; 13836600, 0; 13840200, 5228.245817; 13843800, 27030.19324; 13847400, 55192.90246; 13851000, 75667.49244; 13854600, 96824.02753; 13858200, 111737.5175; 13861800, 119429.696; 13865400, 120886.9477; 13869000, 108096.4178; 13872600, 90682.46238; 13876200, 69672.05838; 13879800, 48868.67344; 13883400, 24623.08938; 13887000, 2143.256049; 13890600, 0; 13894200, 0; 13897800, 0; 13901400, 0; 13905000, 0; 13908600, 0; 13912200, 0; 13915800, 0; 13919400, 0; 13923000, 0; 13926600, 5228.245817; 13930200, 27030.19324; 13933800, 55192.90246; 13937400, 75667.49244; 13941000, 96824.02753; 13944600, 111737.5175; 13948200, 119429.696; 13951800, 120886.9477; 13955400, 108096.4178; 13959000, 90682.46238; 13962600, 69672.05838; 13966200, 48868.67344; 13969800, 24623.08938; 13973400, 2143.256049; 13977000, 0; 13980600, 0; 13984200, 0; 13987800, 0; 13991400, 0; 13995000, 0; 13998600, 0; 14002200, 0; 14005800, 0; 14009400, 0; 14013000, 5228.245817; 14016600, 27030.19324; 14020200, 55192.90246; 14023800, 75667.49244; 14027400, 96824.02753; 14031000, 111737.5175; 14034600, 119429.696; 14038200, 120886.9477; 14041800, 108096.4178; 14045400, 90682.46238; 14049000, 69672.05838; 14052600, 48868.67344; 14056200, 24623.08938; 14059800, 2143.256049; 14063400, 0; 14067000, 0; 14070600, 0; 14074200, 0; 14077800, 0; 14081400, 0; 14085000, 0; 14088600, 0; 14092200, 0; 14095800, 0; 14099400, 5228.245817; 14103000, 27030.19324; 14106600, 55192.90246; 14110200, 75667.49244; 14113800, 96824.02753; 14117400, 111737.5175; 14121000, 119429.696; 14124600, 120886.9477; 14128200, 108096.4178; 14131800, 90682.46238; 14135400, 69672.05838; 14139000, 48868.67344; 14142600, 24623.08938; 14146200, 2143.256049; 14149800, 0; 14153400, 0; 14157000, 0; 14160600, 0; 14164200, 0; 14167800, 0; 14171400, 0; 14175000, 0; 14178600, 0; 14182200, 0; 14185800, 5228.245817; 14189400, 27030.19324; 14193000, 55192.90246; 14196600, 75667.49244; 14200200, 96824.02753; 14203800, 111737.5175; 14207400, 119429.696; 14211000, 120886.9477; 14214600, 108096.4178; 14218200, 90682.46238; 14221800, 69672.05838; 14225400, 48868.67344; 14229000, 24623.08938; 14232600, 2143.256049; 14236200, 0; 14239800, 0; 14243400, 0; 14247000, 0; 14250600, 0; 14254200, 0; 14257800, 0; 14261400, 0; 14265000, 0; 14268600, 0; 14272200, 5228.245817; 14275800, 27030.19324; 14279400, 55192.90246; 14283000, 75667.49244; 14286600, 96824.02753; 14290200, 111737.5175; 14293800, 119429.696; 14297400, 120886.9477; 14301000, 108096.4178; 14304600, 90682.46238; 14308200, 69672.05838; 14311800, 48868.67344; 14315400, 24623.08938; 14319000, 2143.256049; 14322600, 0; 14326200, 0; 14329800, 0; 14333400, 0; 14337000, 0; 14340600, 0; 14344200, 0; 14347800, 0; 14351400, 0; 14355000, 0; 14358600, 5228.245817; 14362200, 27030.19324; 14365800, 55192.90246; 14369400, 75667.49244; 14373000, 96824.02753; 14376600, 111737.5175; 14380200, 119429.696; 14383800, 120886.9477; 14387400, 108096.4178; 14391000, 90682.46238; 14394600, 69672.05838; 14398200, 48868.67344; 14401800, 24623.08938; 14405400, 2143.256049; 14409000, 0; 14412600, 0; 14416200, 0; 14419800, 0; 14423400, 0; 14427000, 0; 14430600, 0; 14434200, 0; 14437800, 0; 14441400, 0; 14445000, 5228.245817; 14448600, 27030.19324; 14452200, 55192.90246; 14455800, 75667.49244; 14459400, 96824.02753; 14463000, 111737.5175; 14466600, 119429.696; 14470200, 120886.9477; 14473800, 108096.4178; 14477400, 90682.46238; 14481000, 69672.05838; 14484600, 48868.67344; 14488200, 24623.08938; 14491800, 2143.256049; 14495400, 0; 14499000, 0; 14502600, 0; 14506200, 0; 14509800, 0; 14513400, 0; 14517000, 0; 14520600, 0; 14524200, 0; 14527800, 0; 14531400, 5228.245817; 14535000, 27030.19324; 14538600, 55192.90246; 14542200, 75667.49244; 14545800, 96824.02753; 14549400, 111737.5175; 14553000, 119429.696; 14556600, 120886.9477; 14560200, 108096.4178; 14563800, 90682.46238; 14567400, 69672.05838; 14571000, 48868.67344; 14574600, 24623.08938; 14578200, 2143.256049; 14581800, 0; 14585400, 0; 14589000, 0; 14592600, 0; 14596200, 0; 14599800, 0; 14603400, 0; 14607000, 0; 14610600, 0; 14614200, 0; 14617800, 5228.245817; 14621400, 27030.19324; 14625000, 55192.90246; 14628600, 75667.49244; 14632200, 96824.02753; 14635800, 111737.5175; 14639400, 119429.696; 14643000, 120886.9477; 14646600, 108096.4178; 14650200, 90682.46238; 14653800, 69672.05838; 14657400, 48868.67344; 14661000, 24623.08938; 14664600, 2143.256049; 14668200, 0; 14671800, 0; 14675400, 0; 14679000, 0; 14682600, 0; 14686200, 0; 14689800, 0; 14693400, 0; 14697000, 0; 14700600, 0; 14704200, 5228.245817; 14707800, 27030.19324; 14711400, 55192.90246; 14715000, 75667.49244; 14718600, 96824.02753; 14722200, 111737.5175; 14725800, 119429.696; 14729400, 120886.9477; 14733000, 108096.4178; 14736600, 90682.46238; 14740200, 69672.05838; 14743800, 48868.67344; 14747400, 24623.08938; 14751000, 2143.256049; 14754600, 0; 14758200, 0; 14761800, 0; 14765400, 0; 14769000, 0; 14772600, 0; 14776200, 0; 14779800, 0; 14783400, 0; 14787000, 0; 14790600, 5228.245817; 14794200, 27030.19324; 14797800, 55192.90246; 14801400, 75667.49244; 14805000, 96824.02753; 14808600, 111737.5175; 14812200, 119429.696; 14815800, 120886.9477; 14819400, 108096.4178; 14823000, 90682.46238; 14826600, 69672.05838; 14830200, 48868.67344; 14833800, 24623.08938; 14837400, 2143.256049; 14841000, 0; 14844600, 0; 14848200, 0; 14851800, 0; 14855400, 0; 14859000, 0; 14862600, 0; 14866200, 0; 14869800, 0; 14873400, 0; 14877000, 5228.245817; 14880600, 27030.19324; 14884200, 55192.90246; 14887800, 75667.49244; 14891400, 96824.02753; 14895000, 111737.5175; 14898600, 119429.696; 14902200, 120886.9477; 14905800, 108096.4178; 14909400, 90682.46238; 14913000, 69672.05838; 14916600, 48868.67344; 14920200, 24623.08938; 14923800, 2143.256049; 14927400, 0; 14931000, 0; 14934600, 0; 14938200, 0; 14941800, 0; 14945400, 0; 14949000, 0; 14952600, 0; 14956200, 0; 14959800, 0; 14963400, 5228.245817; 14967000, 27030.19324; 14970600, 55192.90246; 14974200, 75667.49244; 14977800, 96824.02753; 14981400, 111737.5175; 14985000, 119429.696; 14988600, 120886.9477; 14992200, 108096.4178; 14995800, 90682.46238; 14999400, 69672.05838; 15003000, 48868.67344; 15006600, 24623.08938; 15010200, 2143.256049; 15013800, 0; 15017400, 0; 15021000, 0; 15024600, 0; 15028200, 0; 15031800, 0; 15035400, 0; 15039000, 0; 15042600, 0; 15046200, 0; 15049800, 5228.245817; 15053400, 27030.19324; 15057000, 55192.90246; 15060600, 75667.49244; 15064200, 96824.02753; 15067800, 111737.5175; 15071400, 119429.696; 15075000, 120886.9477; 15078600, 108096.4178; 15082200, 90682.46238; 15085800, 69672.05838; 15089400, 48868.67344; 15093000, 24623.08938; 15096600, 2143.256049; 15100200, 0; 15103800, 0; 15107400, 0; 15111000, 0; 15114600, 0; 15118200, 0; 15121800, 0; 15125400, 0; 15129000, 0; 15132600, 0; 15136200, 5228.245817; 15139800, 27030.19324; 15143400, 55192.90246; 15147000, 75667.49244; 15150600, 96824.02753; 15154200, 111737.5175; 15157800, 119429.696; 15161400, 120886.9477; 15165000, 108096.4178; 15168600, 90682.46238; 15172200, 69672.05838; 15175800, 48868.67344; 15179400, 24623.08938; 15183000, 2143.256049; 15186600, 0; 15190200, 0; 15193800, 0; 15197400, 0; 15201000, 0; 15204600, 0; 15208200, 0; 15211800, 0; 15215400, 0; 15219000, 0; 15222600, 5228.245817; 15226200, 27030.19324; 15229800, 55192.90246; 15233400, 75667.49244; 15237000, 96824.02753; 15240600, 111737.5175; 15244200, 119429.696; 15247800, 120886.9477; 15251400, 108096.4178; 15255000, 90682.46238; 15258600, 69672.05838; 15262200, 48868.67344; 15265800, 24623.08938; 15269400, 2143.256049; 15273000, 0; 15276600, 0; 15280200, 0; 15283800, 0; 15287400, 0; 15291000, 0; 15294600, 0; 15298200, 0; 15301800, 0; 15305400, 0; 15309000, 5228.245817; 15312600, 27030.19324; 15316200, 55192.90246; 15319800, 75667.49244; 15323400, 96824.02753; 15327000, 111737.5175; 15330600, 119429.696; 15334200, 120886.9477; 15337800, 108096.4178; 15341400, 90682.46238; 15345000, 69672.05838; 15348600, 48868.67344; 15352200, 24623.08938; 15355800, 2143.256049; 15359400, 0; 15363000, 0; 15366600, 0; 15370200, 0; 15373800, 0; 15377400, 0; 15381000, 0; 15384600, 0; 15388200, 0; 15391800, 0; 15395400, 5228.245817; 15399000, 27030.19324; 15402600, 55192.90246; 15406200, 75667.49244; 15409800, 96824.02753; 15413400, 111737.5175; 15417000, 119429.696; 15420600, 120886.9477; 15424200, 108096.4178; 15427800, 90682.46238; 15431400, 69672.05838; 15435000, 48868.67344; 15438600, 24623.08938; 15442200, 2143.256049; 15445800, 0; 15449400, 0; 15453000, 0; 15456600, 0; 15460200, 0; 15463800, 0; 15467400, 0; 15471000, 0; 15474600, 0; 15478200, 0; 15481800, 5228.245817; 15485400, 27030.19324; 15489000, 55192.90246; 15492600, 75667.49244; 15496200, 96824.02753; 15499800, 111737.5175; 15503400, 119429.696; 15507000, 120886.9477; 15510600, 108096.4178; 15514200, 90682.46238; 15517800, 69672.05838; 15521400, 48868.67344; 15525000, 24623.08938; 15528600, 2143.256049; 15532200, 0; 15535800, 0; 15539400, 0; 15543000, 0; 15546600, 0; 15550200, 0; 15553800, 0; 15557400, 0; 15561000, 0; 15564600, 0; 15568200, 5228.245817; 15571800, 27030.19324; 15575400, 55192.90246; 15579000, 75667.49244; 15582600, 96824.02753; 15586200, 111737.5175; 15589800, 119429.696; 15593400, 120886.9477; 15597000, 108096.4178; 15600600, 90682.46238; 15604200, 69672.05838; 15607800, 48868.67344; 15611400, 24623.08938; 15615000, 2143.256049; 15618600, 0; 15622200, 0; 15625800, 0; 15629400, 0; 15633000, 0; 15636600, 0; 15640200, 0; 15643800, 0; 15647400, 0; 15651000, 0; 15654600, 1194.189587; 15658200, 31029.288; 15661800, 66344.30244; 15665400, 103234.5472; 15669000, 133364.2648; 15672600, 153508.3576; 15676200, 155280.0007; 15679800, 153901.1831; 15683400, 140800.4519; 15687000, 119913.9189; 15690600, 96030.12715; 15694200, 63445.25009; 15697800, 30832.87524; 15701400, 2600.504957; 15705000, 0; 15708600, 0; 15712200, 0; 15715800, 0; 15719400, 0; 15723000, 0; 15726600, 0; 15730200, 0; 15733800, 0; 15737400, 0; 15741000, 1194.189587; 15744600, 31029.288; 15748200, 66344.30244; 15751800, 103234.5472; 15755400, 133364.2648; 15759000, 153508.3576; 15762600, 155280.0007; 15766200, 153901.1831; 15769800, 140800.4519; 15773400, 119913.9189; 15777000, 96030.12715; 15780600, 63445.25009; 15784200, 30832.87524; 15787800, 2600.504957; 15791400, 0; 15795000, 0; 15798600, 0; 15802200, 0; 15805800, 0; 15809400, 0; 15813000, 0; 15816600, 0; 15820200, 0; 15823800, 0; 15827400, 1194.189587; 15831000, 31029.288; 15834600, 66344.30244; 15838200, 103234.5472; 15841800, 133364.2648; 15845400, 153508.3576; 15849000, 155280.0007; 15852600, 153901.1831; 15856200, 140800.4519; 15859800, 119913.9189; 15863400, 96030.12715; 15867000, 63445.25009; 15870600, 30832.87524; 15874200, 2600.504957; 15877800, 0; 15881400, 0; 15885000, 0; 15888600, 0; 15892200, 0; 15895800, 0; 15899400, 0; 15903000, 0; 15906600, 0; 15910200, 0; 15913800, 1194.189587; 15917400, 31029.288; 15921000, 66344.30244; 15924600, 103234.5472; 15928200, 133364.2648; 15931800, 153508.3576; 15935400, 155280.0007; 15939000, 153901.1831; 15942600, 140800.4519; 15946200, 119913.9189; 15949800, 96030.12715; 15953400, 63445.25009; 15957000, 30832.87524; 15960600, 2600.504957; 15964200, 0; 15967800, 0; 15971400, 0; 15975000, 0; 15978600, 0; 15982200, 0; 15985800, 0; 15989400, 0; 15993000, 0; 15996600, 0; 16000200, 1194.189587; 16003800, 31029.288; 16007400, 66344.30244; 16011000, 103234.5472; 16014600, 133364.2648; 16018200, 153508.3576; 16021800, 155280.0007; 16025400, 153901.1831; 16029000, 140800.4519; 16032600, 119913.9189; 16036200, 96030.12715; 16039800, 63445.25009; 16043400, 30832.87524; 16047000, 2600.504957; 16050600, 0; 16054200, 0; 16057800, 0; 16061400, 0; 16065000, 0; 16068600, 0; 16072200, 0; 16075800, 0; 16079400, 0; 16083000, 0; 16086600, 1194.189587; 16090200, 31029.288; 16093800, 66344.30244; 16097400, 103234.5472; 16101000, 133364.2648; 16104600, 153508.3576; 16108200, 155280.0007; 16111800, 153901.1831; 16115400, 140800.4519; 16119000, 119913.9189; 16122600, 96030.12715; 16126200, 63445.25009; 16129800, 30832.87524; 16133400, 2600.504957; 16137000, 0; 16140600, 0; 16144200, 0; 16147800, 0; 16151400, 0; 16155000, 0; 16158600, 0; 16162200, 0; 16165800, 0; 16169400, 0; 16173000, 1194.189587; 16176600, 31029.288; 16180200, 66344.30244; 16183800, 103234.5472; 16187400, 133364.2648; 16191000, 153508.3576; 16194600, 155280.0007; 16198200, 153901.1831; 16201800, 140800.4519; 16205400, 119913.9189; 16209000, 96030.12715; 16212600, 63445.25009; 16216200, 30832.87524; 16219800, 2600.504957; 16223400, 0; 16227000, 0; 16230600, 0; 16234200, 0; 16237800, 0; 16241400, 0; 16245000, 0; 16248600, 0; 16252200, 0; 16255800, 0; 16259400, 1194.189587; 16263000, 31029.288; 16266600, 66344.30244; 16270200, 103234.5472; 16273800, 133364.2648; 16277400, 153508.3576; 16281000, 155280.0007; 16284600, 153901.1831; 16288200, 140800.4519; 16291800, 119913.9189; 16295400, 96030.12715; 16299000, 63445.25009; 16302600, 30832.87524; 16306200, 2600.504957; 16309800, 0; 16313400, 0; 16317000, 0; 16320600, 0; 16324200, 0; 16327800, 0; 16331400, 0; 16335000, 0; 16338600, 0; 16342200, 0; 16345800, 1194.189587; 16349400, 31029.288; 16353000, 66344.30244; 16356600, 103234.5472; 16360200, 133364.2648; 16363800, 153508.3576; 16367400, 155280.0007; 16371000, 153901.1831; 16374600, 140800.4519; 16378200, 119913.9189; 16381800, 96030.12715; 16385400, 63445.25009; 16389000, 30832.87524; 16392600, 2600.504957; 16396200, 0; 16399800, 0; 16403400, 0; 16407000, 0; 16410600, 0; 16414200, 0; 16417800, 0; 16421400, 0; 16425000, 0; 16428600, 0; 16432200, 1194.189587; 16435800, 31029.288; 16439400, 66344.30244; 16443000, 103234.5472; 16446600, 133364.2648; 16450200, 153508.3576; 16453800, 155280.0007; 16457400, 153901.1831; 16461000, 140800.4519; 16464600, 119913.9189; 16468200, 96030.12715; 16471800, 63445.25009; 16475400, 30832.87524; 16479000, 2600.504957; 16482600, 0; 16486200, 0; 16489800, 0; 16493400, 0; 16497000, 0; 16500600, 0; 16504200, 0; 16507800, 0; 16511400, 0; 16515000, 0; 16518600, 1194.189587; 16522200, 31029.288; 16525800, 66344.30244; 16529400, 103234.5472; 16533000, 133364.2648; 16536600, 153508.3576; 16540200, 155280.0007; 16543800, 153901.1831; 16547400, 140800.4519; 16551000, 119913.9189; 16554600, 96030.12715; 16558200, 63445.25009; 16561800, 30832.87524; 16565400, 2600.504957; 16569000, 0; 16572600, 0; 16576200, 0; 16579800, 0; 16583400, 0; 16587000, 0; 16590600, 0; 16594200, 0; 16597800, 0; 16601400, 0; 16605000, 1194.189587; 16608600, 31029.288; 16612200, 66344.30244; 16615800, 103234.5472; 16619400, 133364.2648; 16623000, 153508.3576; 16626600, 155280.0007; 16630200, 153901.1831; 16633800, 140800.4519; 16637400, 119913.9189; 16641000, 96030.12715; 16644600, 63445.25009; 16648200, 30832.87524; 16651800, 2600.504957; 16655400, 0; 16659000, 0; 16662600, 0; 16666200, 0; 16669800, 0; 16673400, 0; 16677000, 0; 16680600, 0; 16684200, 0; 16687800, 0; 16691400, 1194.189587; 16695000, 31029.288; 16698600, 66344.30244; 16702200, 103234.5472; 16705800, 133364.2648; 16709400, 153508.3576; 16713000, 155280.0007; 16716600, 153901.1831; 16720200, 140800.4519; 16723800, 119913.9189; 16727400, 96030.12715; 16731000, 63445.25009; 16734600, 30832.87524; 16738200, 2600.504957; 16741800, 0; 16745400, 0; 16749000, 0; 16752600, 0; 16756200, 0; 16759800, 0; 16763400, 0; 16767000, 0; 16770600, 0; 16774200, 0; 16777800, 1194.189587; 16781400, 31029.288; 16785000, 66344.30244; 16788600, 103234.5472; 16792200, 133364.2648; 16795800, 153508.3576; 16799400, 155280.0007; 16803000, 153901.1831; 16806600, 140800.4519; 16810200, 119913.9189; 16813800, 96030.12715; 16817400, 63445.25009; 16821000, 30832.87524; 16824600, 2600.504957; 16828200, 0; 16831800, 0; 16835400, 0; 16839000, 0; 16842600, 0; 16846200, 0; 16849800, 0; 16853400, 0; 16857000, 0; 16860600, 0; 16864200, 1194.189587; 16867800, 31029.288; 16871400, 66344.30244; 16875000, 103234.5472; 16878600, 133364.2648; 16882200, 153508.3576; 16885800, 155280.0007; 16889400, 153901.1831; 16893000, 140800.4519; 16896600, 119913.9189; 16900200, 96030.12715; 16903800, 63445.25009; 16907400, 30832.87524; 16911000, 2600.504957; 16914600, 0; 16918200, 0; 16921800, 0; 16925400, 0; 16929000, 0; 16932600, 0; 16936200, 0; 16939800, 0; 16943400, 0; 16947000, 0; 16950600, 1194.189587; 16954200, 31029.288; 16957800, 66344.30244; 16961400, 103234.5472; 16965000, 133364.2648; 16968600, 153508.3576; 16972200, 155280.0007; 16975800, 153901.1831; 16979400, 140800.4519; 16983000, 119913.9189; 16986600, 96030.12715; 16990200, 63445.25009; 16993800, 30832.87524; 16997400, 2600.504957; 17001000, 0; 17004600, 0; 17008200, 0; 17011800, 0; 17015400, 0; 17019000, 0; 17022600, 0; 17026200, 0; 17029800, 0; 17033400, 0; 17037000, 1194.189587; 17040600, 31029.288; 17044200, 66344.30244; 17047800, 103234.5472; 17051400, 133364.2648; 17055000, 153508.3576; 17058600, 155280.0007; 17062200, 153901.1831; 17065800, 140800.4519; 17069400, 119913.9189; 17073000, 96030.12715; 17076600, 63445.25009; 17080200, 30832.87524; 17083800, 2600.504957; 17087400, 0; 17091000, 0; 17094600, 0; 17098200, 0; 17101800, 0; 17105400, 0; 17109000, 0; 17112600, 0; 17116200, 0; 17119800, 0; 17123400, 1194.189587; 17127000, 31029.288; 17130600, 66344.30244; 17134200, 103234.5472; 17137800, 133364.2648; 17141400, 153508.3576; 17145000, 155280.0007; 17148600, 153901.1831; 17152200, 140800.4519; 17155800, 119913.9189; 17159400, 96030.12715; 17163000, 63445.25009; 17166600, 30832.87524; 17170200, 2600.504957; 17173800, 0; 17177400, 0; 17181000, 0; 17184600, 0; 17188200, 0; 17191800, 0; 17195400, 0; 17199000, 0; 17202600, 0; 17206200, 0; 17209800, 1194.189587; 17213400, 31029.288; 17217000, 66344.30244; 17220600, 103234.5472; 17224200, 133364.2648; 17227800, 153508.3576; 17231400, 155280.0007; 17235000, 153901.1831; 17238600, 140800.4519; 17242200, 119913.9189; 17245800, 96030.12715; 17249400, 63445.25009; 17253000, 30832.87524; 17256600, 2600.504957; 17260200, 0; 17263800, 0; 17267400, 0; 17271000, 0; 17274600, 0; 17278200, 0; 17281800, 0; 17285400, 0; 17289000, 0; 17292600, 0; 17296200, 1194.189587; 17299800, 31029.288; 17303400, 66344.30244; 17307000, 103234.5472; 17310600, 133364.2648; 17314200, 153508.3576; 17317800, 155280.0007; 17321400, 153901.1831; 17325000, 140800.4519; 17328600, 119913.9189; 17332200, 96030.12715; 17335800, 63445.25009; 17339400, 30832.87524; 17343000, 2600.504957; 17346600, 0; 17350200, 0; 17353800, 0; 17357400, 0; 17361000, 0; 17364600, 0; 17368200, 0; 17371800, 0; 17375400, 0; 17379000, 0; 17382600, 1194.189587; 17386200, 31029.288; 17389800, 66344.30244; 17393400, 103234.5472; 17397000, 133364.2648; 17400600, 153508.3576; 17404200, 155280.0007; 17407800, 153901.1831; 17411400, 140800.4519; 17415000, 119913.9189; 17418600, 96030.12715; 17422200, 63445.25009; 17425800, 30832.87524; 17429400, 2600.504957; 17433000, 0; 17436600, 0; 17440200, 0; 17443800, 0; 17447400, 0; 17451000, 0; 17454600, 0; 17458200, 0; 17461800, 0; 17465400, 0; 17469000, 1194.189587; 17472600, 31029.288; 17476200, 66344.30244; 17479800, 103234.5472; 17483400, 133364.2648; 17487000, 153508.3576; 17490600, 155280.0007; 17494200, 153901.1831; 17497800, 140800.4519; 17501400, 119913.9189; 17505000, 96030.12715; 17508600, 63445.25009; 17512200, 30832.87524; 17515800, 2600.504957; 17519400, 0; 17523000, 0; 17526600, 0; 17530200, 0; 17533800, 0; 17537400, 0; 17541000, 0; 17544600, 0; 17548200, 0; 17551800, 0; 17555400, 1194.189587; 17559000, 31029.288; 17562600, 66344.30244; 17566200, 103234.5472; 17569800, 133364.2648; 17573400, 153508.3576; 17577000, 155280.0007; 17580600, 153901.1831; 17584200, 140800.4519; 17587800, 119913.9189; 17591400, 96030.12715; 17595000, 63445.25009; 17598600, 30832.87524; 17602200, 2600.504957; 17605800, 0; 17609400, 0; 17613000, 0; 17616600, 0; 17620200, 0; 17623800, 0; 17627400, 0; 17631000, 0; 17634600, 0; 17638200, 0; 17641800, 1194.189587; 17645400, 31029.288; 17649000, 66344.30244; 17652600, 103234.5472; 17656200, 133364.2648; 17659800, 153508.3576; 17663400, 155280.0007; 17667000, 153901.1831; 17670600, 140800.4519; 17674200, 119913.9189; 17677800, 96030.12715; 17681400, 63445.25009; 17685000, 30832.87524; 17688600, 2600.504957; 17692200, 0; 17695800, 0; 17699400, 0; 17703000, 0; 17706600, 0; 17710200, 0; 17713800, 0; 17717400, 0; 17721000, 0; 17724600, 0; 17728200, 1194.189587; 17731800, 31029.288; 17735400, 66344.30244; 17739000, 103234.5472; 17742600, 133364.2648; 17746200, 153508.3576; 17749800, 155280.0007; 17753400, 153901.1831; 17757000, 140800.4519; 17760600, 119913.9189; 17764200, 96030.12715; 17767800, 63445.25009; 17771400, 30832.87524; 17775000, 2600.504957; 17778600, 0; 17782200, 0; 17785800, 0; 17789400, 0; 17793000, 0; 17796600, 0; 17800200, 0; 17803800, 0; 17807400, 0; 17811000, 0; 17814600, 1194.189587; 17818200, 31029.288; 17821800, 66344.30244; 17825400, 103234.5472; 17829000, 133364.2648; 17832600, 153508.3576; 17836200, 155280.0007; 17839800, 153901.1831; 17843400, 140800.4519; 17847000, 119913.9189; 17850600, 96030.12715; 17854200, 63445.25009; 17857800, 30832.87524; 17861400, 2600.504957; 17865000, 0; 17868600, 0; 17872200, 0; 17875800, 0; 17879400, 0; 17883000, 0; 17886600, 0; 17890200, 0; 17893800, 0; 17897400, 0; 17901000, 1194.189587; 17904600, 31029.288; 17908200, 66344.30244; 17911800, 103234.5472; 17915400, 133364.2648; 17919000, 153508.3576; 17922600, 155280.0007; 17926200, 153901.1831; 17929800, 140800.4519; 17933400, 119913.9189; 17937000, 96030.12715; 17940600, 63445.25009; 17944200, 30832.87524; 17947800, 2600.504957; 17951400, 0; 17955000, 0; 17958600, 0; 17962200, 0; 17965800, 0; 17969400, 0; 17973000, 0; 17976600, 0; 17980200, 0; 17983800, 0; 17987400, 1194.189587; 17991000, 31029.288; 17994600, 66344.30244; 17998200, 103234.5472; 18001800, 133364.2648; 18005400, 153508.3576; 18009000, 155280.0007; 18012600, 153901.1831; 18016200, 140800.4519; 18019800, 119913.9189; 18023400, 96030.12715; 18027000, 63445.25009; 18030600, 30832.87524; 18034200, 2600.504957; 18037800, 0; 18041400, 0; 18045000, 0; 18048600, 0; 18052200, 0; 18055800, 0; 18059400, 0; 18063000, 0; 18066600, 0; 18070200, 0; 18073800, 1194.189587; 18077400, 31029.288; 18081000, 66344.30244; 18084600, 103234.5472; 18088200, 133364.2648; 18091800, 153508.3576; 18095400, 155280.0007; 18099000, 153901.1831; 18102600, 140800.4519; 18106200, 119913.9189; 18109800, 96030.12715; 18113400, 63445.25009; 18117000, 30832.87524; 18120600, 2600.504957; 18124200, 0; 18127800, 0; 18131400, 0; 18135000, 0; 18138600, 0; 18142200, 0; 18145800, 0; 18149400, 0; 18153000, 0; 18156600, 0; 18160200, 1194.189587; 18163800, 31029.288; 18167400, 66344.30244; 18171000, 103234.5472; 18174600, 133364.2648; 18178200, 153508.3576; 18181800, 155280.0007; 18185400, 153901.1831; 18189000, 140800.4519; 18192600, 119913.9189; 18196200, 96030.12715; 18199800, 63445.25009; 18203400, 30832.87524; 18207000, 2600.504957; 18210600, 0; 18214200, 0; 18217800, 0; 18221400, 0; 18225000, 0; 18228600, 0; 18232200, 0; 18235800, 0; 18239400, 0; 18243000, 0; 18246600, 1194.189587; 18250200, 31029.288; 18253800, 66344.30244; 18257400, 103234.5472; 18261000, 133364.2648; 18264600, 153508.3576; 18268200, 155280.0007; 18271800, 153901.1831; 18275400, 140800.4519; 18279000, 119913.9189; 18282600, 96030.12715; 18286200, 63445.25009; 18289800, 30832.87524; 18293400, 2600.504957; 18297000, 0; 18300600, 0; 18304200, 0; 18307800, 0; 18311400, 0; 18315000, 0; 18318600, 0; 18322200, 0; 18325800, 0; 18329400, 0; 18333000, 0; 18336600, 21762.53393; 18340200, 51852.96893; 18343800, 83019.74586; 18347400, 111267.8292; 18351000, 128371.4524; 18354600, 136628.6449; 18358200, 135913.7024; 18361800, 124513.9058; 18365400, 102331.0485; 18369000, 76561.69427; 18372600, 49676.71554; 18376200, 21809.67299; 18379800, 58.92382833; 18383400, 0; 18387000, 0; 18390600, 0; 18394200, 0; 18397800, 0; 18401400, 0; 18405000, 0; 18408600, 0; 18412200, 0; 18415800, 0; 18419400, 0; 18423000, 21762.53393; 18426600, 51852.96893; 18430200, 83019.74586; 18433800, 111267.8292; 18437400, 128371.4524; 18441000, 136628.6449; 18444600, 135913.7024; 18448200, 124513.9058; 18451800, 102331.0485; 18455400, 76561.69427; 18459000, 49676.71554; 18462600, 21809.67299; 18466200, 58.92382833; 18469800, 0; 18473400, 0; 18477000, 0; 18480600, 0; 18484200, 0; 18487800, 0; 18491400, 0; 18495000, 0; 18498600, 0; 18502200, 0; 18505800, 0; 18509400, 21762.53393; 18513000, 51852.96893; 18516600, 83019.74586; 18520200, 111267.8292; 18523800, 128371.4524; 18527400, 136628.6449; 18531000, 135913.7024; 18534600, 124513.9058; 18538200, 102331.0485; 18541800, 76561.69427; 18545400, 49676.71554; 18549000, 21809.67299; 18552600, 58.92382833; 18556200, 0; 18559800, 0; 18563400, 0; 18567000, 0; 18570600, 0; 18574200, 0; 18577800, 0; 18581400, 0; 18585000, 0; 18588600, 0; 18592200, 0; 18595800, 21762.53393; 18599400, 51852.96893; 18603000, 83019.74586; 18606600, 111267.8292; 18610200, 128371.4524; 18613800, 136628.6449; 18617400, 135913.7024; 18621000, 124513.9058; 18624600, 102331.0485; 18628200, 76561.69427; 18631800, 49676.71554; 18635400, 21809.67299; 18639000, 58.92382833; 18642600, 0; 18646200, 0; 18649800, 0; 18653400, 0; 18657000, 0; 18660600, 0; 18664200, 0; 18667800, 0; 18671400, 0; 18675000, 0; 18678600, 0; 18682200, 21762.53393; 18685800, 51852.96893; 18689400, 83019.74586; 18693000, 111267.8292; 18696600, 128371.4524; 18700200, 136628.6449; 18703800, 135913.7024; 18707400, 124513.9058; 18711000, 102331.0485; 18714600, 76561.69427; 18718200, 49676.71554; 18721800, 21809.67299; 18725400, 58.92382833; 18729000, 0; 18732600, 0; 18736200, 0; 18739800, 0; 18743400, 0; 18747000, 0; 18750600, 0; 18754200, 0; 18757800, 0; 18761400, 0; 18765000, 0; 18768600, 21762.53393; 18772200, 51852.96893; 18775800, 83019.74586; 18779400, 111267.8292; 18783000, 128371.4524; 18786600, 136628.6449; 18790200, 135913.7024; 18793800, 124513.9058; 18797400, 102331.0485; 18801000, 76561.69427; 18804600, 49676.71554; 18808200, 21809.67299; 18811800, 58.92382833; 18815400, 0; 18819000, 0; 18822600, 0; 18826200, 0; 18829800, 0; 18833400, 0; 18837000, 0; 18840600, 0; 18844200, 0; 18847800, 0; 18851400, 0; 18855000, 21762.53393; 18858600, 51852.96893; 18862200, 83019.74586; 18865800, 111267.8292; 18869400, 128371.4524; 18873000, 136628.6449; 18876600, 135913.7024; 18880200, 124513.9058; 18883800, 102331.0485; 18887400, 76561.69427; 18891000, 49676.71554; 18894600, 21809.67299; 18898200, 58.92382833; 18901800, 0; 18905400, 0; 18909000, 0; 18912600, 0; 18916200, 0; 18919800, 0; 18923400, 0; 18927000, 0; 18930600, 0; 18934200, 0; 18937800, 0; 18941400, 21762.53393; 18945000, 51852.96893; 18948600, 83019.74586; 18952200, 111267.8292; 18955800, 128371.4524; 18959400, 136628.6449; 18963000, 135913.7024; 18966600, 124513.9058; 18970200, 102331.0485; 18973800, 76561.69427; 18977400, 49676.71554; 18981000, 21809.67299; 18984600, 58.92382833; 18988200, 0; 18991800, 0; 18995400, 0; 18999000, 0; 19002600, 0; 19006200, 0; 19009800, 0; 19013400, 0; 19017000, 0; 19020600, 0; 19024200, 0; 19027800, 21762.53393; 19031400, 51852.96893; 19035000, 83019.74586; 19038600, 111267.8292; 19042200, 128371.4524; 19045800, 136628.6449; 19049400, 135913.7024; 19053000, 124513.9058; 19056600, 102331.0485; 19060200, 76561.69427; 19063800, 49676.71554; 19067400, 21809.67299; 19071000, 58.92382833; 19074600, 0; 19078200, 0; 19081800, 0; 19085400, 0; 19089000, 0; 19092600, 0; 19096200, 0; 19099800, 0; 19103400, 0; 19107000, 0; 19110600, 0; 19114200, 21762.53393; 19117800, 51852.96893; 19121400, 83019.74586; 19125000, 111267.8292; 19128600, 128371.4524; 19132200, 136628.6449; 19135800, 135913.7024; 19139400, 124513.9058; 19143000, 102331.0485; 19146600, 76561.69427; 19150200, 49676.71554; 19153800, 21809.67299; 19157400, 58.92382833; 19161000, 0; 19164600, 0; 19168200, 0; 19171800, 0; 19175400, 0; 19179000, 0; 19182600, 0; 19186200, 0; 19189800, 0; 19193400, 0; 19197000, 0; 19200600, 21762.53393; 19204200, 51852.96893; 19207800, 83019.74586; 19211400, 111267.8292; 19215000, 128371.4524; 19218600, 136628.6449; 19222200, 135913.7024; 19225800, 124513.9058; 19229400, 102331.0485; 19233000, 76561.69427; 19236600, 49676.71554; 19240200, 21809.67299; 19243800, 58.92382833; 19247400, 0; 19251000, 0; 19254600, 0; 19258200, 0; 19261800, 0; 19265400, 0; 19269000, 0; 19272600, 0; 19276200, 0; 19279800, 0; 19283400, 0; 19287000, 21762.53393; 19290600, 51852.96893; 19294200, 83019.74586; 19297800, 111267.8292; 19301400, 128371.4524; 19305000, 136628.6449; 19308600, 135913.7024; 19312200, 124513.9058; 19315800, 102331.0485; 19319400, 76561.69427; 19323000, 49676.71554; 19326600, 21809.67299; 19330200, 58.92382833; 19333800, 0; 19337400, 0; 19341000, 0; 19344600, 0; 19348200, 0; 19351800, 0; 19355400, 0; 19359000, 0; 19362600, 0; 19366200, 0; 19369800, 0; 19373400, 21762.53393; 19377000, 51852.96893; 19380600, 83019.74586; 19384200, 111267.8292; 19387800, 128371.4524; 19391400, 136628.6449; 19395000, 135913.7024; 19398600, 124513.9058; 19402200, 102331.0485; 19405800, 76561.69427; 19409400, 49676.71554; 19413000, 21809.67299; 19416600, 58.92382833; 19420200, 0; 19423800, 0; 19427400, 0; 19431000, 0; 19434600, 0; 19438200, 0; 19441800, 0; 19445400, 0; 19449000, 0; 19452600, 0; 19456200, 0; 19459800, 21762.53393; 19463400, 51852.96893; 19467000, 83019.74586; 19470600, 111267.8292; 19474200, 128371.4524; 19477800, 136628.6449; 19481400, 135913.7024; 19485000, 124513.9058; 19488600, 102331.0485; 19492200, 76561.69427; 19495800, 49676.71554; 19499400, 21809.67299; 19503000, 58.92382833; 19506600, 0; 19510200, 0; 19513800, 0; 19517400, 0; 19521000, 0; 19524600, 0; 19528200, 0; 19531800, 0; 19535400, 0; 19539000, 0; 19542600, 0; 19546200, 21762.53393; 19549800, 51852.96893; 19553400, 83019.74586; 19557000, 111267.8292; 19560600, 128371.4524; 19564200, 136628.6449; 19567800, 135913.7024; 19571400, 124513.9058; 19575000, 102331.0485; 19578600, 76561.69427; 19582200, 49676.71554; 19585800, 21809.67299; 19589400, 58.92382833; 19593000, 0; 19596600, 0; 19600200, 0; 19603800, 0; 19607400, 0; 19611000, 0; 19614600, 0; 19618200, 0; 19621800, 0; 19625400, 0; 19629000, 0; 19632600, 21762.53393; 19636200, 51852.96893; 19639800, 83019.74586; 19643400, 111267.8292; 19647000, 128371.4524; 19650600, 136628.6449; 19654200, 135913.7024; 19657800, 124513.9058; 19661400, 102331.0485; 19665000, 76561.69427; 19668600, 49676.71554; 19672200, 21809.67299; 19675800, 58.92382833; 19679400, 0; 19683000, 0; 19686600, 0; 19690200, 0; 19693800, 0; 19697400, 0; 19701000, 0; 19704600, 0; 19708200, 0; 19711800, 0; 19715400, 0; 19719000, 21762.53393; 19722600, 51852.96893; 19726200, 83019.74586; 19729800, 111267.8292; 19733400, 128371.4524; 19737000, 136628.6449; 19740600, 135913.7024; 19744200, 124513.9058; 19747800, 102331.0485; 19751400, 76561.69427; 19755000, 49676.71554; 19758600, 21809.67299; 19762200, 58.92382833; 19765800, 0; 19769400, 0; 19773000, 0; 19776600, 0; 19780200, 0; 19783800, 0; 19787400, 0; 19791000, 0; 19794600, 0; 19798200, 0; 19801800, 0; 19805400, 21762.53393; 19809000, 51852.96893; 19812600, 83019.74586; 19816200, 111267.8292; 19819800, 128371.4524; 19823400, 136628.6449; 19827000, 135913.7024; 19830600, 124513.9058; 19834200, 102331.0485; 19837800, 76561.69427; 19841400, 49676.71554; 19845000, 21809.67299; 19848600, 58.92382833; 19852200, 0; 19855800, 0; 19859400, 0; 19863000, 0; 19866600, 0; 19870200, 0; 19873800, 0; 19877400, 0; 19881000, 0; 19884600, 0; 19888200, 0; 19891800, 21762.53393; 19895400, 51852.96893; 19899000, 83019.74586; 19902600, 111267.8292; 19906200, 128371.4524; 19909800, 136628.6449; 19913400, 135913.7024; 19917000, 124513.9058; 19920600, 102331.0485; 19924200, 76561.69427; 19927800, 49676.71554; 19931400, 21809.67299; 19935000, 58.92382833; 19938600, 0; 19942200, 0; 19945800, 0; 19949400, 0; 19953000, 0; 19956600, 0; 19960200, 0; 19963800, 0; 19967400, 0; 19971000, 0; 19974600, 0; 19978200, 21762.53393; 19981800, 51852.96893; 19985400, 83019.74586; 19989000, 111267.8292; 19992600, 128371.4524; 19996200, 136628.6449; 19999800, 135913.7024; 20003400, 124513.9058; 20007000, 102331.0485; 20010600, 76561.69427; 20014200, 49676.71554; 20017800, 21809.67299; 20021400, 58.92382833; 20025000, 0; 20028600, 0; 20032200, 0; 20035800, 0; 20039400, 0; 20043000, 0; 20046600, 0; 20050200, 0; 20053800, 0; 20057400, 0; 20061000, 0; 20064600, 21762.53393; 20068200, 51852.96893; 20071800, 83019.74586; 20075400, 111267.8292; 20079000, 128371.4524; 20082600, 136628.6449; 20086200, 135913.7024; 20089800, 124513.9058; 20093400, 102331.0485; 20097000, 76561.69427; 20100600, 49676.71554; 20104200, 21809.67299; 20107800, 58.92382833; 20111400, 0; 20115000, 0; 20118600, 0; 20122200, 0; 20125800, 0; 20129400, 0; 20133000, 0; 20136600, 0; 20140200, 0; 20143800, 0; 20147400, 0; 20151000, 21762.53393; 20154600, 51852.96893; 20158200, 83019.74586; 20161800, 111267.8292; 20165400, 128371.4524; 20169000, 136628.6449; 20172600, 135913.7024; 20176200, 124513.9058; 20179800, 102331.0485; 20183400, 76561.69427; 20187000, 49676.71554; 20190600, 21809.67299; 20194200, 58.92382833; 20197800, 0; 20201400, 0; 20205000, 0; 20208600, 0; 20212200, 0; 20215800, 0; 20219400, 0; 20223000, 0; 20226600, 0; 20230200, 0; 20233800, 0; 20237400, 21762.53393; 20241000, 51852.96893; 20244600, 83019.74586; 20248200, 111267.8292; 20251800, 128371.4524; 20255400, 136628.6449; 20259000, 135913.7024; 20262600, 124513.9058; 20266200, 102331.0485; 20269800, 76561.69427; 20273400, 49676.71554; 20277000, 21809.67299; 20280600, 58.92382833; 20284200, 0; 20287800, 0; 20291400, 0; 20295000, 0; 20298600, 0; 20302200, 0; 20305800, 0; 20309400, 0; 20313000, 0; 20316600, 0; 20320200, 0; 20323800, 21762.53393; 20327400, 51852.96893; 20331000, 83019.74586; 20334600, 111267.8292; 20338200, 128371.4524; 20341800, 136628.6449; 20345400, 135913.7024; 20349000, 124513.9058; 20352600, 102331.0485; 20356200, 76561.69427; 20359800, 49676.71554; 20363400, 21809.67299; 20367000, 58.92382833; 20370600, 0; 20374200, 0; 20377800, 0; 20381400, 0; 20385000, 0; 20388600, 0; 20392200, 0; 20395800, 0; 20399400, 0; 20403000, 0; 20406600, 0; 20410200, 21762.53393; 20413800, 51852.96893; 20417400, 83019.74586; 20421000, 111267.8292; 20424600, 128371.4524; 20428200, 136628.6449; 20431800, 135913.7024; 20435400, 124513.9058; 20439000, 102331.0485; 20442600, 76561.69427; 20446200, 49676.71554; 20449800, 21809.67299; 20453400, 58.92382833; 20457000, 0; 20460600, 0; 20464200, 0; 20467800, 0; 20471400, 0; 20475000, 0; 20478600, 0; 20482200, 0; 20485800, 0; 20489400, 0; 20493000, 0; 20496600, 21762.53393; 20500200, 51852.96893; 20503800, 83019.74586; 20507400, 111267.8292; 20511000, 128371.4524; 20514600, 136628.6449; 20518200, 135913.7024; 20521800, 124513.9058; 20525400, 102331.0485; 20529000, 76561.69427; 20532600, 49676.71554; 20536200, 21809.67299; 20539800, 58.92382833; 20543400, 0; 20547000, 0; 20550600, 0; 20554200, 0; 20557800, 0; 20561400, 0; 20565000, 0; 20568600, 0; 20572200, 0; 20575800, 0; 20579400, 0; 20583000, 21762.53393; 20586600, 51852.96893; 20590200, 83019.74586; 20593800, 111267.8292; 20597400, 128371.4524; 20601000, 136628.6449; 20604600, 135913.7024; 20608200, 124513.9058; 20611800, 102331.0485; 20615400, 76561.69427; 20619000, 49676.71554; 20622600, 21809.67299; 20626200, 58.92382833; 20629800, 0; 20633400, 0; 20637000, 0; 20640600, 0; 20644200, 0; 20647800, 0; 20651400, 0; 20655000, 0; 20658600, 0; 20662200, 0; 20665800, 0; 20669400, 21762.53393; 20673000, 51852.96893; 20676600, 83019.74586; 20680200, 111267.8292; 20683800, 128371.4524; 20687400, 136628.6449; 20691000, 135913.7024; 20694600, 124513.9058; 20698200, 102331.0485; 20701800, 76561.69427; 20705400, 49676.71554; 20709000, 21809.67299; 20712600, 58.92382833; 20716200, 0; 20719800, 0; 20723400, 0; 20727000, 0; 20730600, 0; 20734200, 0; 20737800, 0; 20741400, 0; 20745000, 0; 20748600, 0; 20752200, 0; 20755800, 21762.53393; 20759400, 51852.96893; 20763000, 83019.74586; 20766600, 111267.8292; 20770200, 128371.4524; 20773800, 136628.6449; 20777400, 135913.7024; 20781000, 124513.9058; 20784600, 102331.0485; 20788200, 76561.69427; 20791800, 49676.71554; 20795400, 21809.67299; 20799000, 58.92382833; 20802600, 0; 20806200, 0; 20809800, 0; 20813400, 0; 20817000, 0; 20820600, 0; 20824200, 0; 20827800, 0; 20831400, 0; 20835000, 0; 20838600, 0; 20842200, 21762.53393; 20845800, 51852.96893; 20849400, 83019.74586; 20853000, 111267.8292; 20856600, 128371.4524; 20860200, 136628.6449; 20863800, 135913.7024; 20867400, 124513.9058; 20871000, 102331.0485; 20874600, 76561.69427; 20878200, 49676.71554; 20881800, 21809.67299; 20885400, 58.92382833; 20889000, 0; 20892600, 0; 20896200, 0; 20899800, 0; 20903400, 0; 20907000, 0; 20910600, 0; 20914200, 0; 20917800, 0; 20921400, 0; 20925000, 0; 20928600, 21762.53393; 20932200, 51852.96893; 20935800, 83019.74586; 20939400, 111267.8292; 20943000, 128371.4524; 20946600, 136628.6449; 20950200, 135913.7024; 20953800, 124513.9058; 20957400, 102331.0485; 20961000, 76561.69427; 20964600, 49676.71554; 20968200, 21809.67299; 20971800, 58.92382833; 20975400, 0; 20979000, 0; 20982600, 0; 20986200, 0; 20989800, 0; 20993400, 0; 20997000, 0; 21000600, 0; 21004200, 0; 21007800, 0; 21011400, 0; 21015000, 13123.3841; 21018600, 43104.61361; 21022200, 74494.38449; 21025800, 103615.0642; 21029400, 120830.119; 21033000, 127048.8089; 21036600, 124085.595; 21040200, 117249.9072; 21043800, 93714.68258; 21047400, 70337.7667; 21051000, 37470.44809; 21054600, 7363.383472; 21058200, 0; 21061800, 0; 21065400, 0; 21069000, 0; 21072600, 0; 21076200, 0; 21079800, 0; 21083400, 0; 21087000, 0; 21090600, 0; 21094200, 0; 21097800, 0; 21101400, 13123.3841; 21105000, 43104.61361; 21108600, 74494.38449; 21112200, 103615.0642; 21115800, 120830.119; 21119400, 127048.8089; 21123000, 124085.595; 21126600, 117249.9072; 21130200, 93714.68258; 21133800, 70337.7667; 21137400, 37470.44809; 21141000, 7363.383472; 21144600, 0; 21148200, 0; 21151800, 0; 21155400, 0; 21159000, 0; 21162600, 0; 21166200, 0; 21169800, 0; 21173400, 0; 21177000, 0; 21180600, 0; 21184200, 0; 21187800, 13123.3841; 21191400, 43104.61361; 21195000, 74494.38449; 21198600, 103615.0642; 21202200, 120830.119; 21205800, 127048.8089; 21209400, 124085.595; 21213000, 117249.9072; 21216600, 93714.68258; 21220200, 70337.7667; 21223800, 37470.44809; 21227400, 7363.383472; 21231000, 0; 21234600, 0; 21238200, 0; 21241800, 0; 21245400, 0; 21249000, 0; 21252600, 0; 21256200, 0; 21259800, 0; 21263400, 0; 21267000, 0; 21270600, 0; 21274200, 13123.3841; 21277800, 43104.61361; 21281400, 74494.38449; 21285000, 103615.0642; 21288600, 120830.119; 21292200, 127048.8089; 21295800, 124085.595; 21299400, 117249.9072; 21303000, 93714.68258; 21306600, 70337.7667; 21310200, 37470.44809; 21313800, 7363.383472; 21317400, 0; 21321000, 0; 21324600, 0; 21328200, 0; 21331800, 0; 21335400, 0; 21339000, 0; 21342600, 0; 21346200, 0; 21349800, 0; 21353400, 0; 21357000, 0; 21360600, 13123.3841; 21364200, 43104.61361; 21367800, 74494.38449; 21371400, 103615.0642; 21375000, 120830.119; 21378600, 127048.8089; 21382200, 124085.595; 21385800, 117249.9072; 21389400, 93714.68258; 21393000, 70337.7667; 21396600, 37470.44809; 21400200, 7363.383472; 21403800, 0; 21407400, 0; 21411000, 0; 21414600, 0; 21418200, 0; 21421800, 0; 21425400, 0; 21429000, 0; 21432600, 0; 21436200, 0; 21439800, 0; 21443400, 0; 21447000, 13123.3841; 21450600, 43104.61361; 21454200, 74494.38449; 21457800, 103615.0642; 21461400, 120830.119; 21465000, 127048.8089; 21468600, 124085.595; 21472200, 117249.9072; 21475800, 93714.68258; 21479400, 70337.7667; 21483000, 37470.44809; 21486600, 7363.383472; 21490200, 0; 21493800, 0; 21497400, 0; 21501000, 0; 21504600, 0; 21508200, 0; 21511800, 0; 21515400, 0; 21519000, 0; 21522600, 0; 21526200, 0; 21529800, 0; 21533400, 13123.3841; 21537000, 43104.61361; 21540600, 74494.38449; 21544200, 103615.0642; 21547800, 120830.119; 21551400, 127048.8089; 21555000, 124085.595; 21558600, 117249.9072; 21562200, 93714.68258; 21565800, 70337.7667; 21569400, 37470.44809; 21573000, 7363.383472; 21576600, 0; 21580200, 0; 21583800, 0; 21587400, 0; 21591000, 0; 21594600, 0; 21598200, 0; 21601800, 0; 21605400, 0; 21609000, 0; 21612600, 0; 21616200, 0; 21619800, 13123.3841; 21623400, 43104.61361; 21627000, 74494.38449; 21630600, 103615.0642; 21634200, 120830.119; 21637800, 127048.8089; 21641400, 124085.595; 21645000, 117249.9072; 21648600, 93714.68258; 21652200, 70337.7667; 21655800, 37470.44809; 21659400, 7363.383472; 21663000, 0; 21666600, 0; 21670200, 0; 21673800, 0; 21677400, 0; 21681000, 0; 21684600, 0; 21688200, 0; 21691800, 0; 21695400, 0; 21699000, 0; 21702600, 0; 21706200, 13123.3841; 21709800, 43104.61361; 21713400, 74494.38449; 21717000, 103615.0642; 21720600, 120830.119; 21724200, 127048.8089; 21727800, 124085.595; 21731400, 117249.9072; 21735000, 93714.68258; 21738600, 70337.7667; 21742200, 37470.44809; 21745800, 7363.383472; 21749400, 0; 21753000, 0; 21756600, 0; 21760200, 0; 21763800, 0; 21767400, 0; 21771000, 0; 21774600, 0; 21778200, 0; 21781800, 0; 21785400, 0; 21789000, 0; 21792600, 13123.3841; 21796200, 43104.61361; 21799800, 74494.38449; 21803400, 103615.0642; 21807000, 120830.119; 21810600, 127048.8089; 21814200, 124085.595; 21817800, 117249.9072; 21821400, 93714.68258; 21825000, 70337.7667; 21828600, 37470.44809; 21832200, 7363.383472; 21835800, 0; 21839400, 0; 21843000, 0; 21846600, 0; 21850200, 0; 21853800, 0; 21857400, 0; 21861000, 0; 21864600, 0; 21868200, 0; 21871800, 0; 21875400, 0; 21879000, 13123.3841; 21882600, 43104.61361; 21886200, 74494.38449; 21889800, 103615.0642; 21893400, 120830.119; 21897000, 127048.8089; 21900600, 124085.595; 21904200, 117249.9072; 21907800, 93714.68258; 21911400, 70337.7667; 21915000, 37470.44809; 21918600, 7363.383472; 21922200, 0; 21925800, 0; 21929400, 0; 21933000, 0; 21936600, 0; 21940200, 0; 21943800, 0; 21947400, 0; 21951000, 0; 21954600, 0; 21958200, 0; 21961800, 0; 21965400, 13123.3841; 21969000, 43104.61361; 21972600, 74494.38449; 21976200, 103615.0642; 21979800, 120830.119; 21983400, 127048.8089; 21987000, 124085.595; 21990600, 117249.9072; 21994200, 93714.68258; 21997800, 70337.7667; 22001400, 37470.44809; 22005000, 7363.383472; 22008600, 0; 22012200, 0; 22015800, 0; 22019400, 0; 22023000, 0; 22026600, 0; 22030200, 0; 22033800, 0; 22037400, 0; 22041000, 0; 22044600, 0; 22048200, 0; 22051800, 13123.3841; 22055400, 43104.61361; 22059000, 74494.38449; 22062600, 103615.0642; 22066200, 120830.119; 22069800, 127048.8089; 22073400, 124085.595; 22077000, 117249.9072; 22080600, 93714.68258; 22084200, 70337.7667; 22087800, 37470.44809; 22091400, 7363.383472; 22095000, 0; 22098600, 0; 22102200, 0; 22105800, 0; 22109400, 0; 22113000, 0; 22116600, 0; 22120200, 0; 22123800, 0; 22127400, 0; 22131000, 0; 22134600, 0; 22138200, 13123.3841; 22141800, 43104.61361; 22145400, 74494.38449; 22149000, 103615.0642; 22152600, 120830.119; 22156200, 127048.8089; 22159800, 124085.595; 22163400, 117249.9072; 22167000, 93714.68258; 22170600, 70337.7667; 22174200, 37470.44809; 22177800, 7363.383472; 22181400, 0; 22185000, 0; 22188600, 0; 22192200, 0; 22195800, 0; 22199400, 0; 22203000, 0; 22206600, 0; 22210200, 0; 22213800, 0; 22217400, 0; 22221000, 0; 22224600, 13123.3841; 22228200, 43104.61361; 22231800, 74494.38449; 22235400, 103615.0642; 22239000, 120830.119; 22242600, 127048.8089; 22246200, 124085.595; 22249800, 117249.9072; 22253400, 93714.68258; 22257000, 70337.7667; 22260600, 37470.44809; 22264200, 7363.383472; 22267800, 0; 22271400, 0; 22275000, 0; 22278600, 0; 22282200, 0; 22285800, 0; 22289400, 0; 22293000, 0; 22296600, 0; 22300200, 0; 22303800, 0; 22307400, 0; 22311000, 13123.3841; 22314600, 43104.61361; 22318200, 74494.38449; 22321800, 103615.0642; 22325400, 120830.119; 22329000, 127048.8089; 22332600, 124085.595; 22336200, 117249.9072; 22339800, 93714.68258; 22343400, 70337.7667; 22347000, 37470.44809; 22350600, 7363.383472; 22354200, 0; 22357800, 0; 22361400, 0; 22365000, 0; 22368600, 0; 22372200, 0; 22375800, 0; 22379400, 0; 22383000, 0; 22386600, 0; 22390200, 0; 22393800, 0; 22397400, 13123.3841; 22401000, 43104.61361; 22404600, 74494.38449; 22408200, 103615.0642; 22411800, 120830.119; 22415400, 127048.8089; 22419000, 124085.595; 22422600, 117249.9072; 22426200, 93714.68258; 22429800, 70337.7667; 22433400, 37470.44809; 22437000, 7363.383472; 22440600, 0; 22444200, 0; 22447800, 0; 22451400, 0; 22455000, 0; 22458600, 0; 22462200, 0; 22465800, 0; 22469400, 0; 22473000, 0; 22476600, 0; 22480200, 0; 22483800, 13123.3841; 22487400, 43104.61361; 22491000, 74494.38449; 22494600, 103615.0642; 22498200, 120830.119; 22501800, 127048.8089; 22505400, 124085.595; 22509000, 117249.9072; 22512600, 93714.68258; 22516200, 70337.7667; 22519800, 37470.44809; 22523400, 7363.383472; 22527000, 0; 22530600, 0; 22534200, 0; 22537800, 0; 22541400, 0; 22545000, 0; 22548600, 0; 22552200, 0; 22555800, 0; 22559400, 0; 22563000, 0; 22566600, 0; 22570200, 13123.3841; 22573800, 43104.61361; 22577400, 74494.38449; 22581000, 103615.0642; 22584600, 120830.119; 22588200, 127048.8089; 22591800, 124085.595; 22595400, 117249.9072; 22599000, 93714.68258; 22602600, 70337.7667; 22606200, 37470.44809; 22609800, 7363.383472; 22613400, 0; 22617000, 0; 22620600, 0; 22624200, 0; 22627800, 0; 22631400, 0; 22635000, 0; 22638600, 0; 22642200, 0; 22645800, 0; 22649400, 0; 22653000, 0; 22656600, 13123.3841; 22660200, 43104.61361; 22663800, 74494.38449; 22667400, 103615.0642; 22671000, 120830.119; 22674600, 127048.8089; 22678200, 124085.595; 22681800, 117249.9072; 22685400, 93714.68258; 22689000, 70337.7667; 22692600, 37470.44809; 22696200, 7363.383472; 22699800, 0; 22703400, 0; 22707000, 0; 22710600, 0; 22714200, 0; 22717800, 0; 22721400, 0; 22725000, 0; 22728600, 0; 22732200, 0; 22735800, 0; 22739400, 0; 22743000, 13123.3841; 22746600, 43104.61361; 22750200, 74494.38449; 22753800, 103615.0642; 22757400, 120830.119; 22761000, 127048.8089; 22764600, 124085.595; 22768200, 117249.9072; 22771800, 93714.68258; 22775400, 70337.7667; 22779000, 37470.44809; 22782600, 7363.383472; 22786200, 0; 22789800, 0; 22793400, 0; 22797000, 0; 22800600, 0; 22804200, 0; 22807800, 0; 22811400, 0; 22815000, 0; 22818600, 0; 22822200, 0; 22825800, 0; 22829400, 13123.3841; 22833000, 43104.61361; 22836600, 74494.38449; 22840200, 103615.0642; 22843800, 120830.119; 22847400, 127048.8089; 22851000, 124085.595; 22854600, 117249.9072; 22858200, 93714.68258; 22861800, 70337.7667; 22865400, 37470.44809; 22869000, 7363.383472; 22872600, 0; 22876200, 0; 22879800, 0; 22883400, 0; 22887000, 0; 22890600, 0; 22894200, 0; 22897800, 0; 22901400, 0; 22905000, 0; 22908600, 0; 22912200, 0; 22915800, 13123.3841; 22919400, 43104.61361; 22923000, 74494.38449; 22926600, 103615.0642; 22930200, 120830.119; 22933800, 127048.8089; 22937400, 124085.595; 22941000, 117249.9072; 22944600, 93714.68258; 22948200, 70337.7667; 22951800, 37470.44809; 22955400, 7363.383472; 22959000, 0; 22962600, 0; 22966200, 0; 22969800, 0; 22973400, 0; 22977000, 0; 22980600, 0; 22984200, 0; 22987800, 0; 22991400, 0; 22995000, 0; 22998600, 0; 23002200, 13123.3841; 23005800, 43104.61361; 23009400, 74494.38449; 23013000, 103615.0642; 23016600, 120830.119; 23020200, 127048.8089; 23023800, 124085.595; 23027400, 117249.9072; 23031000, 93714.68258; 23034600, 70337.7667; 23038200, 37470.44809; 23041800, 7363.383472; 23045400, 0; 23049000, 0; 23052600, 0; 23056200, 0; 23059800, 0; 23063400, 0; 23067000, 0; 23070600, 0; 23074200, 0; 23077800, 0; 23081400, 0; 23085000, 0; 23088600, 13123.3841; 23092200, 43104.61361; 23095800, 74494.38449; 23099400, 103615.0642; 23103000, 120830.119; 23106600, 127048.8089; 23110200, 124085.595; 23113800, 117249.9072; 23117400, 93714.68258; 23121000, 70337.7667; 23124600, 37470.44809; 23128200, 7363.383472; 23131800, 0; 23135400, 0; 23139000, 0; 23142600, 0; 23146200, 0; 23149800, 0; 23153400, 0; 23157000, 0; 23160600, 0; 23164200, 0; 23167800, 0; 23171400, 0; 23175000, 13123.3841; 23178600, 43104.61361; 23182200, 74494.38449; 23185800, 103615.0642; 23189400, 120830.119; 23193000, 127048.8089; 23196600, 124085.595; 23200200, 117249.9072; 23203800, 93714.68258; 23207400, 70337.7667; 23211000, 37470.44809; 23214600, 7363.383472; 23218200, 0; 23221800, 0; 23225400, 0; 23229000, 0; 23232600, 0; 23236200, 0; 23239800, 0; 23243400, 0; 23247000, 0; 23250600, 0; 23254200, 0; 23257800, 0; 23261400, 13123.3841; 23265000, 43104.61361; 23268600, 74494.38449; 23272200, 103615.0642; 23275800, 120830.119; 23279400, 127048.8089; 23283000, 124085.595; 23286600, 117249.9072; 23290200, 93714.68258; 23293800, 70337.7667; 23297400, 37470.44809; 23301000, 7363.383472; 23304600, 0; 23308200, 0; 23311800, 0; 23315400, 0; 23319000, 0; 23322600, 0; 23326200, 0; 23329800, 0; 23333400, 0; 23337000, 0; 23340600, 0; 23344200, 0; 23347800, 13123.3841; 23351400, 43104.61361; 23355000, 74494.38449; 23358600, 103615.0642; 23362200, 120830.119; 23365800, 127048.8089; 23369400, 124085.595; 23373000, 117249.9072; 23376600, 93714.68258; 23380200, 70337.7667; 23383800, 37470.44809; 23387400, 7363.383472; 23391000, 0; 23394600, 0; 23398200, 0; 23401800, 0; 23405400, 0; 23409000, 0; 23412600, 0; 23416200, 0; 23419800, 0; 23423400, 0; 23427000, 0; 23430600, 0; 23434200, 13123.3841; 23437800, 43104.61361; 23441400, 74494.38449; 23445000, 103615.0642; 23448600, 120830.119; 23452200, 127048.8089; 23455800, 124085.595; 23459400, 117249.9072; 23463000, 93714.68258; 23466600, 70337.7667; 23470200, 37470.44809; 23473800, 7363.383472; 23477400, 0; 23481000, 0; 23484600, 0; 23488200, 0; 23491800, 0; 23495400, 0; 23499000, 0; 23502600, 0; 23506200, 0; 23509800, 0; 23513400, 0; 23517000, 0; 23520600, 13123.3841; 23524200, 43104.61361; 23527800, 74494.38449; 23531400, 103615.0642; 23535000, 120830.119; 23538600, 127048.8089; 23542200, 124085.595; 23545800, 117249.9072; 23549400, 93714.68258; 23553000, 70337.7667; 23556600, 37470.44809; 23560200, 7363.383472; 23563800, 0; 23567400, 0; 23571000, 0; 23574600, 0; 23578200, 0; 23581800, 0; 23585400, 0; 23589000, 0; 23592600, 0; 23596200, 0; 23599800, 0; 23603400, 0; 23607000, 4764.973584; 23610600, 35000.75403; 23614200, 62078.21727; 23617800, 88153.97543; 23621400, 106353.5819; 23625000, 121909.4726; 23628600, 115082.165; 23632200, 107669.5474; 23635800, 82367.65549; 23639400, 49877.05655; 23643000, 19075.60736; 23646600, 0; 23650200, 0; 23653800, 0; 23657400, 0; 23661000, 0; 23664600, 0; 23668200, 0; 23671800, 0; 23675400, 0; 23679000, 0; 23682600, 0; 23686200, 0; 23689800, 0; 23693400, 4764.973584; 23697000, 35000.75403; 23700600, 62078.21727; 23704200, 88153.97543; 23707800, 106353.5819; 23711400, 121909.4726; 23715000, 115082.165; 23718600, 107669.5474; 23722200, 82367.65549; 23725800, 49877.05655; 23729400, 19075.60736; 23733000, 0; 23736600, 0; 23740200, 0; 23743800, 0; 23747400, 0; 23751000, 0; 23754600, 0; 23758200, 0; 23761800, 0; 23765400, 0; 23769000, 0; 23772600, 0; 23776200, 0; 23779800, 4764.973584; 23783400, 35000.75403; 23787000, 62078.21727; 23790600, 88153.97543; 23794200, 106353.5819; 23797800, 121909.4726; 23801400, 115082.165; 23805000, 107669.5474; 23808600, 82367.65549; 23812200, 49877.05655; 23815800, 19075.60736; 23819400, 0; 23823000, 0; 23826600, 0; 23830200, 0; 23833800, 0; 23837400, 0; 23841000, 0; 23844600, 0; 23848200, 0; 23851800, 0; 23855400, 0; 23859000, 0; 23862600, 0; 23866200, 4764.973584; 23869800, 35000.75403; 23873400, 62078.21727; 23877000, 88153.97543; 23880600, 106353.5819; 23884200, 121909.4726; 23887800, 115082.165; 23891400, 107669.5474; 23895000, 82367.65549; 23898600, 49877.05655; 23902200, 19075.60736; 23905800, 0; 23909400, 0; 23913000, 0; 23916600, 0; 23920200, 0; 23923800, 0; 23927400, 0; 23931000, 0; 23934600, 0; 23938200, 0; 23941800, 0; 23945400, 0; 23949000, 0; 23952600, 4764.973584; 23956200, 35000.75403; 23959800, 62078.21727; 23963400, 88153.97543; 23967000, 106353.5819; 23970600, 121909.4726; 23974200, 115082.165; 23977800, 107669.5474; 23981400, 82367.65549; 23985000, 49877.05655; 23988600, 19075.60736; 23992200, 0; 23995800, 0; 23999400, 0; 24003000, 0; 24006600, 0; 24010200, 0; 24013800, 0; 24017400, 0; 24021000, 0; 24024600, 0; 24028200, 0; 24031800, 0; 24035400, 0; 24039000, 4764.973584; 24042600, 35000.75403; 24046200, 62078.21727; 24049800, 88153.97543; 24053400, 106353.5819; 24057000, 121909.4726; 24060600, 115082.165; 24064200, 107669.5474; 24067800, 82367.65549; 24071400, 49877.05655; 24075000, 19075.60736; 24078600, 0; 24082200, 0; 24085800, 0; 24089400, 0; 24093000, 0; 24096600, 0; 24100200, 0; 24103800, 0; 24107400, 0; 24111000, 0; 24114600, 0; 24118200, 0; 24121800, 0; 24125400, 4764.973584; 24129000, 35000.75403; 24132600, 62078.21727; 24136200, 88153.97543; 24139800, 106353.5819; 24143400, 121909.4726; 24147000, 115082.165; 24150600, 107669.5474; 24154200, 82367.65549; 24157800, 49877.05655; 24161400, 19075.60736; 24165000, 0; 24168600, 0; 24172200, 0; 24175800, 0; 24179400, 0; 24183000, 0; 24186600, 0; 24190200, 0; 24193800, 0; 24197400, 0; 24201000, 0; 24204600, 0; 24208200, 0; 24211800, 4764.973584; 24215400, 35000.75403; 24219000, 62078.21727; 24222600, 88153.97543; 24226200, 106353.5819; 24229800, 121909.4726; 24233400, 115082.165; 24237000, 107669.5474; 24240600, 82367.65549; 24244200, 49877.05655; 24247800, 19075.60736; 24251400, 0; 24255000, 0; 24258600, 0; 24262200, 0; 24265800, 0; 24269400, 0; 24273000, 0; 24276600, 0; 24280200, 0; 24283800, 0; 24287400, 0; 24291000, 0; 24294600, 0; 24298200, 4764.973584; 24301800, 35000.75403; 24305400, 62078.21727; 24309000, 88153.97543; 24312600, 106353.5819; 24316200, 121909.4726; 24319800, 115082.165; 24323400, 107669.5474; 24327000, 82367.65549; 24330600, 49877.05655; 24334200, 19075.60736; 24337800, 0; 24341400, 0; 24345000, 0; 24348600, 0; 24352200, 0; 24355800, 0; 24359400, 0; 24363000, 0; 24366600, 0; 24370200, 0; 24373800, 0; 24377400, 0; 24381000, 0; 24384600, 4764.973584; 24388200, 35000.75403; 24391800, 62078.21727; 24395400, 88153.97543; 24399000, 106353.5819; 24402600, 121909.4726; 24406200, 115082.165; 24409800, 107669.5474; 24413400, 82367.65549; 24417000, 49877.05655; 24420600, 19075.60736; 24424200, 0; 24427800, 0; 24431400, 0; 24435000, 0; 24438600, 0; 24442200, 0; 24445800, 0; 24449400, 0; 24453000, 0; 24456600, 0; 24460200, 0; 24463800, 0; 24467400, 0; 24471000, 4764.973584; 24474600, 35000.75403; 24478200, 62078.21727; 24481800, 88153.97543; 24485400, 106353.5819; 24489000, 121909.4726; 24492600, 115082.165; 24496200, 107669.5474; 24499800, 82367.65549; 24503400, 49877.05655; 24507000, 19075.60736; 24510600, 0; 24514200, 0; 24517800, 0; 24521400, 0; 24525000, 0; 24528600, 0; 24532200, 0; 24535800, 0; 24539400, 0; 24543000, 0; 24546600, 0; 24550200, 0; 24553800, 0; 24557400, 4764.973584; 24561000, 35000.75403; 24564600, 62078.21727; 24568200, 88153.97543; 24571800, 106353.5819; 24575400, 121909.4726; 24579000, 115082.165; 24582600, 107669.5474; 24586200, 82367.65549; 24589800, 49877.05655; 24593400, 19075.60736; 24597000, 0; 24600600, 0; 24604200, 0; 24607800, 0; 24611400, 0; 24615000, 0; 24618600, 0; 24622200, 0; 24625800, 0; 24629400, 0; 24633000, 0; 24636600, 0; 24640200, 0; 24643800, 4764.973584; 24647400, 35000.75403; 24651000, 62078.21727; 24654600, 88153.97543; 24658200, 106353.5819; 24661800, 121909.4726; 24665400, 115082.165; 24669000, 107669.5474; 24672600, 82367.65549; 24676200, 49877.05655; 24679800, 19075.60736; 24683400, 0; 24687000, 0; 24690600, 0; 24694200, 0; 24697800, 0; 24701400, 0; 24705000, 0; 24708600, 0; 24712200, 0; 24715800, 0; 24719400, 0; 24723000, 0; 24726600, 0; 24730200, 4764.973584; 24733800, 35000.75403; 24737400, 62078.21727; 24741000, 88153.97543; 24744600, 106353.5819; 24748200, 121909.4726; 24751800, 115082.165; 24755400, 107669.5474; 24759000, 82367.65549; 24762600, 49877.05655; 24766200, 19075.60736; 24769800, 0; 24773400, 0; 24777000, 0; 24780600, 0; 24784200, 0; 24787800, 0; 24791400, 0; 24795000, 0; 24798600, 0; 24802200, 0; 24805800, 0; 24809400, 0; 24813000, 0; 24816600, 4764.973584; 24820200, 35000.75403; 24823800, 62078.21727; 24827400, 88153.97543; 24831000, 106353.5819; 24834600, 121909.4726; 24838200, 115082.165; 24841800, 107669.5474; 24845400, 82367.65549; 24849000, 49877.05655; 24852600, 19075.60736; 24856200, 0; 24859800, 0; 24863400, 0; 24867000, 0; 24870600, 0; 24874200, 0; 24877800, 0; 24881400, 0; 24885000, 0; 24888600, 0; 24892200, 0; 24895800, 0; 24899400, 0; 24903000, 4764.973584; 24906600, 35000.75403; 24910200, 62078.21727; 24913800, 88153.97543; 24917400, 106353.5819; 24921000, 121909.4726; 24924600, 115082.165; 24928200, 107669.5474; 24931800, 82367.65549; 24935400, 49877.05655; 24939000, 19075.60736; 24942600, 0; 24946200, 0; 24949800, 0; 24953400, 0; 24957000, 0; 24960600, 0; 24964200, 0; 24967800, 0; 24971400, 0; 24975000, 0; 24978600, 0; 24982200, 0; 24985800, 0; 24989400, 4764.973584; 24993000, 35000.75403; 24996600, 62078.21727; 25000200, 88153.97543; 25003800, 106353.5819; 25007400, 121909.4726; 25011000, 115082.165; 25014600, 107669.5474; 25018200, 82367.65549; 25021800, 49877.05655; 25025400, 19075.60736; 25029000, 0; 25032600, 0; 25036200, 0; 25039800, 0; 25043400, 0; 25047000, 0; 25050600, 0; 25054200, 0; 25057800, 0; 25061400, 0; 25065000, 0; 25068600, 0; 25072200, 0; 25075800, 4764.973584; 25079400, 35000.75403; 25083000, 62078.21727; 25086600, 88153.97543; 25090200, 106353.5819; 25093800, 121909.4726; 25097400, 115082.165; 25101000, 107669.5474; 25104600, 82367.65549; 25108200, 49877.05655; 25111800, 19075.60736; 25115400, 0; 25119000, 0; 25122600, 0; 25126200, 0; 25129800, 0; 25133400, 0; 25137000, 0; 25140600, 0; 25144200, 0; 25147800, 0; 25151400, 0; 25155000, 0; 25158600, 0; 25162200, 4764.973584; 25165800, 35000.75403; 25169400, 62078.21727; 25173000, 88153.97543; 25176600, 106353.5819; 25180200, 121909.4726; 25183800, 115082.165; 25187400, 107669.5474; 25191000, 82367.65549; 25194600, 49877.05655; 25198200, 19075.60736; 25201800, 0; 25205400, 0; 25209000, 0; 25212600, 0; 25216200, 0; 25219800, 0; 25223400, 0; 25227000, 0; 25230600, 0; 25234200, 0; 25237800, 0; 25241400, 0; 25245000, 0; 25248600, 4764.973584; 25252200, 35000.75403; 25255800, 62078.21727; 25259400, 88153.97543; 25263000, 106353.5819; 25266600, 121909.4726; 25270200, 115082.165; 25273800, 107669.5474; 25277400, 82367.65549; 25281000, 49877.05655; 25284600, 19075.60736; 25288200, 0; 25291800, 0; 25295400, 0; 25299000, 0; 25302600, 0; 25306200, 0; 25309800, 0; 25313400, 0; 25317000, 0; 25320600, 0; 25324200, 0; 25327800, 0; 25331400, 0; 25335000, 4764.973584; 25338600, 35000.75403; 25342200, 62078.21727; 25345800, 88153.97543; 25349400, 106353.5819; 25353000, 121909.4726; 25356600, 115082.165; 25360200, 107669.5474; 25363800, 82367.65549; 25367400, 49877.05655; 25371000, 19075.60736; 25374600, 0; 25378200, 0; 25381800, 0; 25385400, 0; 25389000, 0; 25392600, 0; 25396200, 0; 25399800, 0; 25403400, 0; 25407000, 0; 25410600, 0; 25414200, 0; 25417800, 0; 25421400, 4764.973584; 25425000, 35000.75403; 25428600, 62078.21727; 25432200, 88153.97543; 25435800, 106353.5819; 25439400, 121909.4726; 25443000, 115082.165; 25446600, 107669.5474; 25450200, 82367.65549; 25453800, 49877.05655; 25457400, 19075.60736; 25461000, 0; 25464600, 0; 25468200, 0; 25471800, 0; 25475400, 0; 25479000, 0; 25482600, 0; 25486200, 0; 25489800, 0; 25493400, 0; 25497000, 0; 25500600, 0; 25504200, 0; 25507800, 4764.973584; 25511400, 35000.75403; 25515000, 62078.21727; 25518600, 88153.97543; 25522200, 106353.5819; 25525800, 121909.4726; 25529400, 115082.165; 25533000, 107669.5474; 25536600, 82367.65549; 25540200, 49877.05655; 25543800, 19075.60736; 25547400, 0; 25551000, 0; 25554600, 0; 25558200, 0; 25561800, 0; 25565400, 0; 25569000, 0; 25572600, 0; 25576200, 0; 25579800, 0; 25583400, 0; 25587000, 0; 25590600, 0; 25594200, 4764.973584; 25597800, 35000.75403; 25601400, 62078.21727; 25605000, 88153.97543; 25608600, 106353.5819; 25612200, 121909.4726; 25615800, 115082.165; 25619400, 107669.5474; 25623000, 82367.65549; 25626600, 49877.05655; 25630200, 19075.60736; 25633800, 0; 25637400, 0; 25641000, 0; 25644600, 0; 25648200, 0; 25651800, 0; 25655400, 0; 25659000, 0; 25662600, 0; 25666200, 0; 25669800, 0; 25673400, 0; 25677000, 0; 25680600, 4764.973584; 25684200, 35000.75403; 25687800, 62078.21727; 25691400, 88153.97543; 25695000, 106353.5819; 25698600, 121909.4726; 25702200, 115082.165; 25705800, 107669.5474; 25709400, 82367.65549; 25713000, 49877.05655; 25716600, 19075.60736; 25720200, 0; 25723800, 0; 25727400, 0; 25731000, 0; 25734600, 0; 25738200, 0; 25741800, 0; 25745400, 0; 25749000, 0; 25752600, 0; 25756200, 0; 25759800, 0; 25763400, 0; 25767000, 4764.973584; 25770600, 35000.75403; 25774200, 62078.21727; 25777800, 88153.97543; 25781400, 106353.5819; 25785000, 121909.4726; 25788600, 115082.165; 25792200, 107669.5474; 25795800, 82367.65549; 25799400, 49877.05655; 25803000, 19075.60736; 25806600, 0; 25810200, 0; 25813800, 0; 25817400, 0; 25821000, 0; 25824600, 0; 25828200, 0; 25831800, 0; 25835400, 0; 25839000, 0; 25842600, 0; 25846200, 0; 25849800, 0; 25853400, 4764.973584; 25857000, 35000.75403; 25860600, 62078.21727; 25864200, 88153.97543; 25867800, 106353.5819; 25871400, 121909.4726; 25875000, 115082.165; 25878600, 107669.5474; 25882200, 82367.65549; 25885800, 49877.05655; 25889400, 19075.60736; 25893000, 0; 25896600, 0; 25900200, 0; 25903800, 0; 25907400, 0; 25911000, 0; 25914600, 0; 25918200, 0; 25921800, 0; 25925400, 0; 25929000, 0; 25932600, 0; 25936200, 0; 25939800, 4764.973584; 25943400, 35000.75403; 25947000, 62078.21727; 25950600, 88153.97543; 25954200, 106353.5819; 25957800, 121909.4726; 25961400, 115082.165; 25965000, 107669.5474; 25968600, 82367.65549; 25972200, 49877.05655; 25975800, 19075.60736; 25979400, 0; 25983000, 0; 25986600, 0; 25990200, 0; 25993800, 0; 25997400, 0; 26001000, 0; 26004600, 0; 26008200, 0; 26011800, 0; 26015400, 0; 26019000, 0; 26022600, 0; 26026200, 4764.973584; 26029800, 35000.75403; 26033400, 62078.21727; 26037000, 88153.97543; 26040600, 106353.5819; 26044200, 121909.4726; 26047800, 115082.165; 26051400, 107669.5474; 26055000, 82367.65549; 26058600, 49877.05655; 26062200, 19075.60736; 26065800, 0; 26069400, 0; 26073000, 0; 26076600, 0; 26080200, 0; 26083800, 0; 26087400, 0; 26091000, 0; 26094600, 0; 26098200, 0; 26101800, 0; 26105400, 0; 26109000, 0; 26112600, 4764.973584; 26116200, 35000.75403; 26119800, 62078.21727; 26123400, 88153.97543; 26127000, 106353.5819; 26130600, 121909.4726; 26134200, 115082.165; 26137800, 107669.5474; 26141400, 82367.65549; 26145000, 49877.05655; 26148600, 19075.60736; 26152200, 0; 26155800, 0; 26159400, 0; 26163000, 0; 26166600, 0; 26170200, 0; 26173800, 0; 26177400, 0; 26181000, 0; 26184600, 0; 26188200, 0; 26191800, 0; 26195400, 0; 26199000, 4764.973584; 26202600, 35000.75403; 26206200, 62078.21727; 26209800, 88153.97543; 26213400, 106353.5819; 26217000, 121909.4726; 26220600, 115082.165; 26224200, 107669.5474; 26227800, 82367.65549; 26231400, 49877.05655; 26235000, 19075.60736; 26238600, 0; 26242200, 0; 26245800, 0; 26249400, 0; 26253000, 0; 26256600, 0; 26260200, 0; 26263800, 0; 26267400, 0; 26271000, 0; 26274600, 0; 26278200, 0; 26281800, 0; 26285400, 0; 26289000, 19869.76962; 26292600, 49156.87643; 26296200, 76166.77368; 26299800, 95464.19652; 26303400, 99397.55847; 26307000, 91607.95931; 26310600, 91944.87267; 26314200, 63786.22264; 26317800, 33155.52161; 26321400, 3811.586042; 26325000, 0; 26328600, 0; 26332200, 0; 26335800, 0; 26339400, 0; 26343000, 0; 26346600, 0; 26350200, 0; 26353800, 0; 26357400, 0; 26361000, 0; 26364600, 0; 26368200, 0; 26371800, 0; 26375400, 19869.76962; 26379000, 49156.87643; 26382600, 76166.77368; 26386200, 95464.19652; 26389800, 99397.55847; 26393400, 91607.95931; 26397000, 91944.87267; 26400600, 63786.22264; 26404200, 33155.52161; 26407800, 3811.586042; 26411400, 0; 26415000, 0; 26418600, 0; 26422200, 0; 26425800, 0; 26429400, 0; 26433000, 0; 26436600, 0; 26440200, 0; 26443800, 0; 26447400, 0; 26451000, 0; 26454600, 0; 26458200, 0; 26461800, 19869.76962; 26465400, 49156.87643; 26469000, 76166.77368; 26472600, 95464.19652; 26476200, 99397.55847; 26479800, 91607.95931; 26483400, 91944.87267; 26487000, 63786.22264; 26490600, 33155.52161; 26494200, 3811.586042; 26497800, 0; 26501400, 0; 26505000, 0; 26508600, 0; 26512200, 0; 26515800, 0; 26519400, 0; 26523000, 0; 26526600, 0; 26530200, 0; 26533800, 0; 26537400, 0; 26541000, 0; 26544600, 0; 26548200, 19869.76962; 26551800, 49156.87643; 26555400, 76166.77368; 26559000, 95464.19652; 26562600, 99397.55847; 26566200, 91607.95931; 26569800, 91944.87267; 26573400, 63786.22264; 26577000, 33155.52161; 26580600, 3811.586042; 26584200, 0; 26587800, 0; 26591400, 0; 26595000, 0; 26598600, 0; 26602200, 0; 26605800, 0; 26609400, 0; 26613000, 0; 26616600, 0; 26620200, 0; 26623800, 0; 26627400, 0; 26631000, 0; 26634600, 19869.76962; 26638200, 49156.87643; 26641800, 76166.77368; 26645400, 95464.19652; 26649000, 99397.55847; 26652600, 91607.95931; 26656200, 91944.87267; 26659800, 63786.22264; 26663400, 33155.52161; 26667000, 3811.586042; 26670600, 0; 26674200, 0; 26677800, 0; 26681400, 0; 26685000, 0; 26688600, 0; 26692200, 0; 26695800, 0; 26699400, 0; 26703000, 0; 26706600, 0; 26710200, 0; 26713800, 0; 26717400, 0; 26721000, 19869.76962; 26724600, 49156.87643; 26728200, 76166.77368; 26731800, 95464.19652; 26735400, 99397.55847; 26739000, 91607.95931; 26742600, 91944.87267; 26746200, 63786.22264; 26749800, 33155.52161; 26753400, 3811.586042; 26757000, 0; 26760600, 0; 26764200, 0; 26767800, 0; 26771400, 0; 26775000, 0; 26778600, 0; 26782200, 0; 26785800, 0; 26789400, 0; 26793000, 0; 26796600, 0; 26800200, 0; 26803800, 0; 26807400, 19869.76962; 26811000, 49156.87643; 26814600, 76166.77368; 26818200, 95464.19652; 26821800, 99397.55847; 26825400, 91607.95931; 26829000, 91944.87267; 26832600, 63786.22264; 26836200, 33155.52161; 26839800, 3811.586042; 26843400, 0; 26847000, 0; 26850600, 0; 26854200, 0; 26857800, 0; 26861400, 0; 26865000, 0; 26868600, 0; 26872200, 0; 26875800, 0; 26879400, 0; 26883000, 0; 26886600, 0; 26890200, 0; 26893800, 19869.76962; 26897400, 49156.87643; 26901000, 76166.77368; 26904600, 95464.19652; 26908200, 99397.55847; 26911800, 91607.95931; 26915400, 91944.87267; 26919000, 63786.22264; 26922600, 33155.52161; 26926200, 3811.586042; 26929800, 0; 26933400, 0; 26937000, 0; 26940600, 0; 26944200, 0; 26947800, 0; 26951400, 0; 26955000, 0; 26958600, 0; 26962200, 0; 26965800, 0; 26969400, 0; 26973000, 0; 26976600, 0; 26980200, 19869.76962; 26983800, 49156.87643; 26987400, 76166.77368; 26991000, 95464.19652; 26994600, 99397.55847; 26998200, 91607.95931; 27001800, 91944.87267; 27005400, 63786.22264; 27009000, 33155.52161; 27012600, 3811.586042; 27016200, 0; 27019800, 0; 27023400, 0; 27027000, 0; 27030600, 0; 27034200, 0; 27037800, 0; 27041400, 0; 27045000, 0; 27048600, 0; 27052200, 0; 27055800, 0; 27059400, 0; 27063000, 0; 27066600, 19869.76962; 27070200, 49156.87643; 27073800, 76166.77368; 27077400, 95464.19652; 27081000, 99397.55847; 27084600, 91607.95931; 27088200, 91944.87267; 27091800, 63786.22264; 27095400, 33155.52161; 27099000, 3811.586042; 27102600, 0; 27106200, 0; 27109800, 0; 27113400, 0; 27117000, 0; 27120600, 0; 27124200, 0; 27127800, 0; 27131400, 0; 27135000, 0; 27138600, 0; 27142200, 0; 27145800, 0; 27149400, 0; 27153000, 19869.76962; 27156600, 49156.87643; 27160200, 76166.77368; 27163800, 95464.19652; 27167400, 99397.55847; 27171000, 91607.95931; 27174600, 91944.87267; 27178200, 63786.22264; 27181800, 33155.52161; 27185400, 3811.586042; 27189000, 0; 27192600, 0; 27196200, 0; 27199800, 0; 27203400, 0; 27207000, 0; 27210600, 0; 27214200, 0; 27217800, 0; 27221400, 0; 27225000, 0; 27228600, 0; 27232200, 0; 27235800, 0; 27239400, 19869.76962; 27243000, 49156.87643; 27246600, 76166.77368; 27250200, 95464.19652; 27253800, 99397.55847; 27257400, 91607.95931; 27261000, 91944.87267; 27264600, 63786.22264; 27268200, 33155.52161; 27271800, 3811.586042; 27275400, 0; 27279000, 0; 27282600, 0; 27286200, 0; 27289800, 0; 27293400, 0; 27297000, 0; 27300600, 0; 27304200, 0; 27307800, 0; 27311400, 0; 27315000, 0; 27318600, 0; 27322200, 0; 27325800, 19869.76962; 27329400, 49156.87643; 27333000, 76166.77368; 27336600, 95464.19652; 27340200, 99397.55847; 27343800, 91607.95931; 27347400, 91944.87267; 27351000, 63786.22264; 27354600, 33155.52161; 27358200, 3811.586042; 27361800, 0; 27365400, 0; 27369000, 0; 27372600, 0; 27376200, 0; 27379800, 0; 27383400, 0; 27387000, 0; 27390600, 0; 27394200, 0; 27397800, 0; 27401400, 0; 27405000, 0; 27408600, 0; 27412200, 19869.76962; 27415800, 49156.87643; 27419400, 76166.77368; 27423000, 95464.19652; 27426600, 99397.55847; 27430200, 91607.95931; 27433800, 91944.87267; 27437400, 63786.22264; 27441000, 33155.52161; 27444600, 3811.586042; 27448200, 0; 27451800, 0; 27455400, 0; 27459000, 0; 27462600, 0; 27466200, 0; 27469800, 0; 27473400, 0; 27477000, 0; 27480600, 0; 27484200, 0; 27487800, 0; 27491400, 0; 27495000, 0; 27498600, 19869.76962; 27502200, 49156.87643; 27505800, 76166.77368; 27509400, 95464.19652; 27513000, 99397.55847; 27516600, 91607.95931; 27520200, 91944.87267; 27523800, 63786.22264; 27527400, 33155.52161; 27531000, 3811.586042; 27534600, 0; 27538200, 0; 27541800, 0; 27545400, 0; 27549000, 0; 27552600, 0; 27556200, 0; 27559800, 0; 27563400, 0; 27567000, 0; 27570600, 0; 27574200, 0; 27577800, 0; 27581400, 0; 27585000, 19869.76962; 27588600, 49156.87643; 27592200, 76166.77368; 27595800, 95464.19652; 27599400, 99397.55847; 27603000, 91607.95931; 27606600, 91944.87267; 27610200, 63786.22264; 27613800, 33155.52161; 27617400, 3811.586042; 27621000, 0; 27624600, 0; 27628200, 0; 27631800, 0; 27635400, 0; 27639000, 0; 27642600, 0; 27646200, 0; 27649800, 0; 27653400, 0; 27657000, 0; 27660600, 0; 27664200, 0; 27667800, 0; 27671400, 19869.76962; 27675000, 49156.87643; 27678600, 76166.77368; 27682200, 95464.19652; 27685800, 99397.55847; 27689400, 91607.95931; 27693000, 91944.87267; 27696600, 63786.22264; 27700200, 33155.52161; 27703800, 3811.586042; 27707400, 0; 27711000, 0; 27714600, 0; 27718200, 0; 27721800, 0; 27725400, 0; 27729000, 0; 27732600, 0; 27736200, 0; 27739800, 0; 27743400, 0; 27747000, 0; 27750600, 0; 27754200, 0; 27757800, 19869.76962; 27761400, 49156.87643; 27765000, 76166.77368; 27768600, 95464.19652; 27772200, 99397.55847; 27775800, 91607.95931; 27779400, 91944.87267; 27783000, 63786.22264; 27786600, 33155.52161; 27790200, 3811.586042; 27793800, 0; 27797400, 0; 27801000, 0; 27804600, 0; 27808200, 0; 27811800, 0; 27815400, 0; 27819000, 0; 27822600, 0; 27826200, 0; 27829800, 0; 27833400, 0; 27837000, 0; 27840600, 0; 27844200, 19869.76962; 27847800, 49156.87643; 27851400, 76166.77368; 27855000, 95464.19652; 27858600, 99397.55847; 27862200, 91607.95931; 27865800, 91944.87267; 27869400, 63786.22264; 27873000, 33155.52161; 27876600, 3811.586042; 27880200, 0; 27883800, 0; 27887400, 0; 27891000, 0; 27894600, 0; 27898200, 0; 27901800, 0; 27905400, 0; 27909000, 0; 27912600, 0; 27916200, 0; 27919800, 0; 27923400, 0; 27927000, 0; 27930600, 19869.76962; 27934200, 49156.87643; 27937800, 76166.77368; 27941400, 95464.19652; 27945000, 99397.55847; 27948600, 91607.95931; 27952200, 91944.87267; 27955800, 63786.22264; 27959400, 33155.52161; 27963000, 3811.586042; 27966600, 0; 27970200, 0; 27973800, 0; 27977400, 0; 27981000, 0; 27984600, 0; 27988200, 0; 27991800, 0; 27995400, 0; 27999000, 0; 28002600, 0; 28006200, 0; 28009800, 0; 28013400, 0; 28017000, 19869.76962; 28020600, 49156.87643; 28024200, 76166.77368; 28027800, 95464.19652; 28031400, 99397.55847; 28035000, 91607.95931; 28038600, 91944.87267; 28042200, 63786.22264; 28045800, 33155.52161; 28049400, 3811.586042; 28053000, 0; 28056600, 0; 28060200, 0; 28063800, 0; 28067400, 0; 28071000, 0; 28074600, 0; 28078200, 0; 28081800, 0; 28085400, 0; 28089000, 0; 28092600, 0; 28096200, 0; 28099800, 0; 28103400, 19869.76962; 28107000, 49156.87643; 28110600, 76166.77368; 28114200, 95464.19652; 28117800, 99397.55847; 28121400, 91607.95931; 28125000, 91944.87267; 28128600, 63786.22264; 28132200, 33155.52161; 28135800, 3811.586042; 28139400, 0; 28143000, 0; 28146600, 0; 28150200, 0; 28153800, 0; 28157400, 0; 28161000, 0; 28164600, 0; 28168200, 0; 28171800, 0; 28175400, 0; 28179000, 0; 28182600, 0; 28186200, 0; 28189800, 19869.76962; 28193400, 49156.87643; 28197000, 76166.77368; 28200600, 95464.19652; 28204200, 99397.55847; 28207800, 91607.95931; 28211400, 91944.87267; 28215000, 63786.22264; 28218600, 33155.52161; 28222200, 3811.586042; 28225800, 0; 28229400, 0; 28233000, 0; 28236600, 0; 28240200, 0; 28243800, 0; 28247400, 0; 28251000, 0; 28254600, 0; 28258200, 0; 28261800, 0; 28265400, 0; 28269000, 0; 28272600, 0; 28276200, 19869.76962; 28279800, 49156.87643; 28283400, 76166.77368; 28287000, 95464.19652; 28290600, 99397.55847; 28294200, 91607.95931; 28297800, 91944.87267; 28301400, 63786.22264; 28305000, 33155.52161; 28308600, 3811.586042; 28312200, 0; 28315800, 0; 28319400, 0; 28323000, 0; 28326600, 0; 28330200, 0; 28333800, 0; 28337400, 0; 28341000, 0; 28344600, 0; 28348200, 0; 28351800, 0; 28355400, 0; 28359000, 0; 28362600, 19869.76962; 28366200, 49156.87643; 28369800, 76166.77368; 28373400, 95464.19652; 28377000, 99397.55847; 28380600, 91607.95931; 28384200, 91944.87267; 28387800, 63786.22264; 28391400, 33155.52161; 28395000, 3811.586042; 28398600, 0; 28402200, 0; 28405800, 0; 28409400, 0; 28413000, 0; 28416600, 0; 28420200, 0; 28423800, 0; 28427400, 0; 28431000, 0; 28434600, 0; 28438200, 0; 28441800, 0; 28445400, 0; 28449000, 19869.76962; 28452600, 49156.87643; 28456200, 76166.77368; 28459800, 95464.19652; 28463400, 99397.55847; 28467000, 91607.95931; 28470600, 91944.87267; 28474200, 63786.22264; 28477800, 33155.52161; 28481400, 3811.586042; 28485000, 0; 28488600, 0; 28492200, 0; 28495800, 0; 28499400, 0; 28503000, 0; 28506600, 0; 28510200, 0; 28513800, 0; 28517400, 0; 28521000, 0; 28524600, 0; 28528200, 0; 28531800, 0; 28535400, 19869.76962; 28539000, 49156.87643; 28542600, 76166.77368; 28546200, 95464.19652; 28549800, 99397.55847; 28553400, 91607.95931; 28557000, 91944.87267; 28560600, 63786.22264; 28564200, 33155.52161; 28567800, 3811.586042; 28571400, 0; 28575000, 0; 28578600, 0; 28582200, 0; 28585800, 0; 28589400, 0; 28593000, 0; 28596600, 0; 28600200, 0; 28603800, 0; 28607400, 0; 28611000, 0; 28614600, 0; 28618200, 0; 28621800, 19869.76962; 28625400, 49156.87643; 28629000, 76166.77368; 28632600, 95464.19652; 28636200, 99397.55847; 28639800, 91607.95931; 28643400, 91944.87267; 28647000, 63786.22264; 28650600, 33155.52161; 28654200, 3811.586042; 28657800, 0; 28661400, 0; 28665000, 0; 28668600, 0; 28672200, 0; 28675800, 0; 28679400, 0; 28683000, 0; 28686600, 0; 28690200, 0; 28693800, 0; 28697400, 0; 28701000, 0; 28704600, 0; 28708200, 19869.76962; 28711800, 49156.87643; 28715400, 76166.77368; 28719000, 95464.19652; 28722600, 99397.55847; 28726200, 91607.95931; 28729800, 91944.87267; 28733400, 63786.22264; 28737000, 33155.52161; 28740600, 3811.586042; 28744200, 0; 28747800, 0; 28751400, 0; 28755000, 0; 28758600, 0; 28762200, 0; 28765800, 0; 28769400, 0; 28773000, 0; 28776600, 0; 28780200, 0; 28783800, 0; 28787400, 0; 28791000, 0; 28794600, 19869.76962; 28798200, 49156.87643; 28801800, 76166.77368; 28805400, 95464.19652; 28809000, 99397.55847; 28812600, 91607.95931; 28816200, 91944.87267; 28819800, 63786.22264; 28823400, 33155.52161; 28827000, 3811.586042; 28830600, 0; 28834200, 0; 28837800, 0; 28841400, 0; 28845000, 0; 28848600, 0; 28852200, 0; 28855800, 0; 28859400, 0; 28863000, 0; 28866600, 0; 28870200, 0; 28873800, 0; 28877400, 0; 28881000, 9950.270477; 28884600, 39039.0004; 28888200, 67126.02523; 28891800, 80855.27723; 28895400, 85698.81592; 28899000, 89155.68052; 28902600, 83699.33401; 28906200, 58197.10111; 28909800, 30408.62367; 28913400, 2298.029305; 28917000, 0; 28920600, 0; 28924200, 0; 28927800, 0; 28931400, 0; 28935000, 0; 28938600, 0; 28942200, 0; 28945800, 0; 28949400, 0; 28953000, 0; 28956600, 0; 28960200, 0; 28963800, 0; 28967400, 9950.270477; 28971000, 39039.0004; 28974600, 67126.02523; 28978200, 80855.27723; 28981800, 85698.81592; 28985400, 89155.68052; 28989000, 83699.33401; 28992600, 58197.10111; 28996200, 30408.62367; 28999800, 2298.029305; 29003400, 0; 29007000, 0; 29010600, 0; 29014200, 0; 29017800, 0; 29021400, 0; 29025000, 0; 29028600, 0; 29032200, 0; 29035800, 0; 29039400, 0; 29043000, 0; 29046600, 0; 29050200, 0; 29053800, 9950.270477; 29057400, 39039.0004; 29061000, 67126.02523; 29064600, 80855.27723; 29068200, 85698.81592; 29071800, 89155.68052; 29075400, 83699.33401; 29079000, 58197.10111; 29082600, 30408.62367; 29086200, 2298.029305; 29089800, 0; 29093400, 0; 29097000, 0; 29100600, 0; 29104200, 0; 29107800, 0; 29111400, 0; 29115000, 0; 29118600, 0; 29122200, 0; 29125800, 0; 29129400, 0; 29133000, 0; 29136600, 0; 29140200, 9950.270477; 29143800, 39039.0004; 29147400, 67126.02523; 29151000, 80855.27723; 29154600, 85698.81592; 29158200, 89155.68052; 29161800, 83699.33401; 29165400, 58197.10111; 29169000, 30408.62367; 29172600, 2298.029305; 29176200, 0; 29179800, 0; 29183400, 0; 29187000, 0; 29190600, 0; 29194200, 0; 29197800, 0; 29201400, 0; 29205000, 0; 29208600, 0; 29212200, 0; 29215800, 0; 29219400, 0; 29223000, 0; 29226600, 9950.270477; 29230200, 39039.0004; 29233800, 67126.02523; 29237400, 80855.27723; 29241000, 85698.81592; 29244600, 89155.68052; 29248200, 83699.33401; 29251800, 58197.10111; 29255400, 30408.62367; 29259000, 2298.029305; 29262600, 0; 29266200, 0; 29269800, 0; 29273400, 0; 29277000, 0; 29280600, 0; 29284200, 0; 29287800, 0; 29291400, 0; 29295000, 0; 29298600, 0; 29302200, 0; 29305800, 0; 29309400, 0; 29313000, 9950.270477; 29316600, 39039.0004; 29320200, 67126.02523; 29323800, 80855.27723; 29327400, 85698.81592; 29331000, 89155.68052; 29334600, 83699.33401; 29338200, 58197.10111; 29341800, 30408.62367; 29345400, 2298.029305; 29349000, 0; 29352600, 0; 29356200, 0; 29359800, 0; 29363400, 0; 29367000, 0; 29370600, 0; 29374200, 0; 29377800, 0; 29381400, 0; 29385000, 0; 29388600, 0; 29392200, 0; 29395800, 0; 29399400, 9950.270477; 29403000, 39039.0004; 29406600, 67126.02523; 29410200, 80855.27723; 29413800, 85698.81592; 29417400, 89155.68052; 29421000, 83699.33401; 29424600, 58197.10111; 29428200, 30408.62367; 29431800, 2298.029305; 29435400, 0; 29439000, 0; 29442600, 0; 29446200, 0; 29449800, 0; 29453400, 0; 29457000, 0; 29460600, 0; 29464200, 0; 29467800, 0; 29471400, 0; 29475000, 0; 29478600, 0; 29482200, 0; 29485800, 9950.270477; 29489400, 39039.0004; 29493000, 67126.02523; 29496600, 80855.27723; 29500200, 85698.81592; 29503800, 89155.68052; 29507400, 83699.33401; 29511000, 58197.10111; 29514600, 30408.62367; 29518200, 2298.029305; 29521800, 0; 29525400, 0; 29529000, 0; 29532600, 0; 29536200, 0; 29539800, 0; 29543400, 0; 29547000, 0; 29550600, 0; 29554200, 0; 29557800, 0; 29561400, 0; 29565000, 0; 29568600, 0; 29572200, 9950.270477; 29575800, 39039.0004; 29579400, 67126.02523; 29583000, 80855.27723; 29586600, 85698.81592; 29590200, 89155.68052; 29593800, 83699.33401; 29597400, 58197.10111; 29601000, 30408.62367; 29604600, 2298.029305; 29608200, 0; 29611800, 0; 29615400, 0; 29619000, 0; 29622600, 0; 29626200, 0; 29629800, 0; 29633400, 0; 29637000, 0; 29640600, 0; 29644200, 0; 29647800, 0; 29651400, 0; 29655000, 0; 29658600, 9950.270477; 29662200, 39039.0004; 29665800, 67126.02523; 29669400, 80855.27723; 29673000, 85698.81592; 29676600, 89155.68052; 29680200, 83699.33401; 29683800, 58197.10111; 29687400, 30408.62367; 29691000, 2298.029305; 29694600, 0; 29698200, 0; 29701800, 0; 29705400, 0; 29709000, 0; 29712600, 0; 29716200, 0; 29719800, 0; 29723400, 0; 29727000, 0; 29730600, 0; 29734200, 0; 29737800, 0; 29741400, 0; 29745000, 9950.270477; 29748600, 39039.0004; 29752200, 67126.02523; 29755800, 80855.27723; 29759400, 85698.81592; 29763000, 89155.68052; 29766600, 83699.33401; 29770200, 58197.10111; 29773800, 30408.62367; 29777400, 2298.029305; 29781000, 0; 29784600, 0; 29788200, 0; 29791800, 0; 29795400, 0; 29799000, 0; 29802600, 0; 29806200, 0; 29809800, 0; 29813400, 0; 29817000, 0; 29820600, 0; 29824200, 0; 29827800, 0; 29831400, 9950.270477; 29835000, 39039.0004; 29838600, 67126.02523; 29842200, 80855.27723; 29845800, 85698.81592; 29849400, 89155.68052; 29853000, 83699.33401; 29856600, 58197.10111; 29860200, 30408.62367; 29863800, 2298.029305; 29867400, 0; 29871000, 0; 29874600, 0; 29878200, 0; 29881800, 0; 29885400, 0; 29889000, 0; 29892600, 0; 29896200, 0; 29899800, 0; 29903400, 0; 29907000, 0; 29910600, 0; 29914200, 0; 29917800, 9950.270477; 29921400, 39039.0004; 29925000, 67126.02523; 29928600, 80855.27723; 29932200, 85698.81592; 29935800, 89155.68052; 29939400, 83699.33401; 29943000, 58197.10111; 29946600, 30408.62367; 29950200, 2298.029305; 29953800, 0; 29957400, 0; 29961000, 0; 29964600, 0; 29968200, 0; 29971800, 0; 29975400, 0; 29979000, 0; 29982600, 0; 29986200, 0; 29989800, 0; 29993400, 0; 29997000, 0; 30000600, 0; 30004200, 9950.270477; 30007800, 39039.0004; 30011400, 67126.02523; 30015000, 80855.27723; 30018600, 85698.81592; 30022200, 89155.68052; 30025800, 83699.33401; 30029400, 58197.10111; 30033000, 30408.62367; 30036600, 2298.029305; 30040200, 0; 30043800, 0; 30047400, 0; 30051000, 0; 30054600, 0; 30058200, 0; 30061800, 0; 30065400, 0; 30069000, 0; 30072600, 0; 30076200, 0; 30079800, 0; 30083400, 0; 30087000, 0; 30090600, 9950.270477; 30094200, 39039.0004; 30097800, 67126.02523; 30101400, 80855.27723; 30105000, 85698.81592; 30108600, 89155.68052; 30112200, 83699.33401; 30115800, 58197.10111; 30119400, 30408.62367; 30123000, 2298.029305; 30126600, 0; 30130200, 0; 30133800, 0; 30137400, 0; 30141000, 0; 30144600, 0; 30148200, 0; 30151800, 0; 30155400, 0; 30159000, 0; 30162600, 0; 30166200, 0; 30169800, 0; 30173400, 0; 30177000, 9950.270477; 30180600, 39039.0004; 30184200, 67126.02523; 30187800, 80855.27723; 30191400, 85698.81592; 30195000, 89155.68052; 30198600, 83699.33401; 30202200, 58197.10111; 30205800, 30408.62367; 30209400, 2298.029305; 30213000, 0; 30216600, 0; 30220200, 0; 30223800, 0; 30227400, 0; 30231000, 0; 30234600, 0; 30238200, 0; 30241800, 0; 30245400, 0; 30249000, 0; 30252600, 0; 30256200, 0; 30259800, 0; 30263400, 9950.270477; 30267000, 39039.0004; 30270600, 67126.02523; 30274200, 80855.27723; 30277800, 85698.81592; 30281400, 89155.68052; 30285000, 83699.33401; 30288600, 58197.10111; 30292200, 30408.62367; 30295800, 2298.029305; 30299400, 0; 30303000, 0; 30306600, 0; 30310200, 0; 30313800, 0; 30317400, 0; 30321000, 0; 30324600, 0; 30328200, 0; 30331800, 0; 30335400, 0; 30339000, 0; 30342600, 0; 30346200, 0; 30349800, 9950.270477; 30353400, 39039.0004; 30357000, 67126.02523; 30360600, 80855.27723; 30364200, 85698.81592; 30367800, 89155.68052; 30371400, 83699.33401; 30375000, 58197.10111; 30378600, 30408.62367; 30382200, 2298.029305; 30385800, 0; 30389400, 0; 30393000, 0; 30396600, 0; 30400200, 0; 30403800, 0; 30407400, 0; 30411000, 0; 30414600, 0; 30418200, 0; 30421800, 0; 30425400, 0; 30429000, 0; 30432600, 0; 30436200, 9950.270477; 30439800, 39039.0004; 30443400, 67126.02523; 30447000, 80855.27723; 30450600, 85698.81592; 30454200, 89155.68052; 30457800, 83699.33401; 30461400, 58197.10111; 30465000, 30408.62367; 30468600, 2298.029305; 30472200, 0; 30475800, 0; 30479400, 0; 30483000, 0; 30486600, 0; 30490200, 0; 30493800, 0; 30497400, 0; 30501000, 0; 30504600, 0; 30508200, 0; 30511800, 0; 30515400, 0; 30519000, 0; 30522600, 9950.270477; 30526200, 39039.0004; 30529800, 67126.02523; 30533400, 80855.27723; 30537000, 85698.81592; 30540600, 89155.68052; 30544200, 83699.33401; 30547800, 58197.10111; 30551400, 30408.62367; 30555000, 2298.029305; 30558600, 0; 30562200, 0; 30565800, 0; 30569400, 0; 30573000, 0; 30576600, 0; 30580200, 0; 30583800, 0; 30587400, 0; 30591000, 0; 30594600, 0; 30598200, 0; 30601800, 0; 30605400, 0; 30609000, 9950.270477; 30612600, 39039.0004; 30616200, 67126.02523; 30619800, 80855.27723; 30623400, 85698.81592; 30627000, 89155.68052; 30630600, 83699.33401; 30634200, 58197.10111; 30637800, 30408.62367; 30641400, 2298.029305; 30645000, 0; 30648600, 0; 30652200, 0; 30655800, 0; 30659400, 0; 30663000, 0; 30666600, 0; 30670200, 0; 30673800, 0; 30677400, 0; 30681000, 0; 30684600, 0; 30688200, 0; 30691800, 0; 30695400, 9950.270477; 30699000, 39039.0004; 30702600, 67126.02523; 30706200, 80855.27723; 30709800, 85698.81592; 30713400, 89155.68052; 30717000, 83699.33401; 30720600, 58197.10111; 30724200, 30408.62367; 30727800, 2298.029305; 30731400, 0; 30735000, 0; 30738600, 0; 30742200, 0; 30745800, 0; 30749400, 0; 30753000, 0; 30756600, 0; 30760200, 0; 30763800, 0; 30767400, 0; 30771000, 0; 30774600, 0; 30778200, 0; 30781800, 9950.270477; 30785400, 39039.0004; 30789000, 67126.02523; 30792600, 80855.27723; 30796200, 85698.81592; 30799800, 89155.68052; 30803400, 83699.33401; 30807000, 58197.10111; 30810600, 30408.62367; 30814200, 2298.029305; 30817800, 0; 30821400, 0; 30825000, 0; 30828600, 0; 30832200, 0; 30835800, 0; 30839400, 0; 30843000, 0; 30846600, 0; 30850200, 0; 30853800, 0; 30857400, 0; 30861000, 0; 30864600, 0; 30868200, 9950.270477; 30871800, 39039.0004; 30875400, 67126.02523; 30879000, 80855.27723; 30882600, 85698.81592; 30886200, 89155.68052; 30889800, 83699.33401; 30893400, 58197.10111; 30897000, 30408.62367; 30900600, 2298.029305; 30904200, 0; 30907800, 0; 30911400, 0; 30915000, 0; 30918600, 0; 30922200, 0; 30925800, 0; 30929400, 0; 30933000, 0; 30936600, 0; 30940200, 0; 30943800, 0; 30947400, 0; 30951000, 0; 30954600, 9950.270477; 30958200, 39039.0004; 30961800, 67126.02523; 30965400, 80855.27723; 30969000, 85698.81592; 30972600, 89155.68052; 30976200, 83699.33401; 30979800, 58197.10111; 30983400, 30408.62367; 30987000, 2298.029305; 30990600, 0; 30994200, 0; 30997800, 0; 31001400, 0; 31005000, 0; 31008600, 0; 31012200, 0; 31015800, 0; 31019400, 0; 31023000, 0; 31026600, 0; 31030200, 0; 31033800, 0; 31037400, 0; 31041000, 9950.270477; 31044600, 39039.0004; 31048200, 67126.02523; 31051800, 80855.27723; 31055400, 85698.81592; 31059000, 89155.68052; 31062600, 83699.33401; 31066200, 58197.10111; 31069800, 30408.62367; 31073400, 2298.029305; 31077000, 0; 31080600, 0; 31084200, 0; 31087800, 0; 31091400, 0; 31095000, 0; 31098600, 0; 31102200, 0; 31105800, 0; 31109400, 0; 31113000, 0; 31116600, 0; 31120200, 0; 31123800, 0; 31127400, 9950.270477; 31131000, 39039.0004; 31134600, 67126.02523; 31138200, 80855.27723; 31141800, 85698.81592; 31145400, 89155.68052; 31149000, 83699.33401; 31152600, 58197.10111; 31156200, 30408.62367; 31159800, 2298.029305; 31163400, 0; 31167000, 0; 31170600, 0; 31174200, 0; 31177800, 0; 31181400, 0; 31185000, 0; 31188600, 0; 31192200, 0; 31195800, 0; 31199400, 0; 31203000, 0; 31206600, 0; 31210200, 0; 31213800, 9950.270477; 31217400, 39039.0004; 31221000, 67126.02523; 31224600, 80855.27723; 31228200, 85698.81592; 31231800, 89155.68052; 31235400, 83699.33401; 31239000, 58197.10111; 31242600, 30408.62367; 31246200, 2298.029305; 31249800, 0; 31253400, 0; 31257000, 0; 31260600, 0; 31264200, 0; 31267800, 0; 31271400, 0; 31275000, 0; 31278600, 0; 31282200, 0; 31285800, 0; 31289400, 0; 31293000, 0; 31296600, 0; 31300200, 9950.270477; 31303800, 39039.0004; 31307400, 67126.02523; 31311000, 80855.27723; 31314600, 85698.81592; 31318200, 89155.68052; 31321800, 83699.33401; 31325400, 58197.10111; 31329000, 30408.62367; 31332600, 2298.029305; 31336200, 0; 31339800, 0; 31343400, 0; 31347000, 0; 31350600, 0; 31354200, 0; 31357800, 0; 31361400, 0; 31365000, 0; 31368600, 0; 31372200, 0; 31375800, 0; 31379400, 0; 31383000, 0; 31386600, 9950.270477; 31390200, 39039.0004; 31393800, 67126.02523; 31397400, 80855.27723; 31401000, 85698.81592; 31404600, 89155.68052; 31408200, 83699.33401; 31411800, 58197.10111; 31415400, 30408.62367; 31419000, 2298.029305; 31422600, 0; 31426200, 0; 31429800, 0; 31433400, 0; 31437000, 0; 31440600, 0; 31444200, 0; 31447800, 0; 31451400, 0; 31455000, 0; 31458600, 0; 31462200, 0; 31465800, 0; 31469400, 0; 31473000, 9950.270477; 31476600, 39039.0004; 31480200, 67126.02523; 31483800, 80855.27723; 31487400, 85698.81592; 31491000, 89155.68052; 31494600, 83699.33401; 31498200, 58197.10111; 31501800, 30408.62367; 31505400, 2298.029305; 31509000, 0; 31512600, 0; 31516200, 0; 31519800, 0; 31523400, 0; 31527000, 0; 31530600, 0]) annotation(
          Placement(transformation(extent = {{-12, -10}, {8, 10}})));
      equation
        P_solar = powerTable.y;
        loss = P_solar * (1 - eta_PV);
        -p.i * p.v = P_solar;
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 20), Text(extent = {{-100, -110}, {100, -126}}, lineColor = {0, 0, 0}, textString = "%name"), Polygon(points = {{-90, 24}, {88, 24}, {88, 18}, {-92, 18}, {-90, 24}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Polygon(points = {{-82, 38}, {96, 38}, {96, 34}, {-84, 34}, {-82, 38}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Polygon(points = {{-86, 46}, {92, 46}, {92, 44}, {-88, 44}, {-86, 46}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Polygon(points = {{-26, 54}, {-58, 0}, {-42, 0}, {-18, 54}, {-26, 54}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Polygon(points = {{-8, 0}, {-4, 54}, {4, 54}, {8, -2}, {-8, 0}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Polygon(points = {{42, 0}, {18, 56}, {22, 56}, {58, 0}, {42, 0}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Bitmap(extent = {{-70, -68}, {64, 60}}, fileName = "modelica://ENN/Resources/icons/光伏.png")}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end SolarPower;

      model WindPower "风力发电"
        parameter Modelica.SIunits.Power P_rat = 2000000 "额定功率" annotation(
          Dialog(group = "风机额定参数"));
        parameter Real e = 0.1977, f = -10.464, g = 183.18, h = -1062.8, i = 2035.4;
        Real eta "风机的总体效率" annotation(
          Dialog(group = "风机额定参数"));
        Modelica.SIunits.Power P_wind "发电功率";
        Modelica.SIunits.Velocity speed;
        ENN.Interfaces.Electrical.Pin_AC p annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}}), iconTransformation(extent = {{90, -10}, {110, 10}})));
        outer ENN.Environment environment annotation(
          Placement(transformation(extent = {{-100, 80}, {-80, 100}})));
      protected
        parameter Real alpha = 0.3 "空气的摩擦系数" annotation(
          Dialog(group = "环境系数"));
        parameter Modelica.SIunits.Length h_ref = 5 "风速的参考高度" annotation(
          Dialog(group = "环境系数"));
      equation
        speed = environment.windSpeed;
        P_wind / 1000 = e * speed ^ 4 + f * speed ^ 3 + g * speed ^ 2 + h * speed + i;
        P_wind = P_rat * eta;
        p.i * p.v = -P_wind;
        annotation(
          Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 45), Bitmap(extent = {{-88, -86}, {88, 86}}, imageSource = "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", fileName = "modelica://RenewableEnergy/windturbine2.png")}),
          Diagram(coordinateSystem(preserveAspectRatio = false)));
      end WindPower;

      package Internal
        model DCbus "直流总线"
          output Modelica.SIunits.Voltage v(stateSelect = StateSelect.never);
          Modelica.Electrical.Analog.Interfaces.PositivePin term "Terminal" annotation(
            Placement(transformation(extent = {{-7, -60}, {7, 60}}, rotation = 0)));
        equation
          term.i = 0;
          v = term.v;
          annotation(
            defaultComponentName = "bus1",
            Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}, grid = {2, 2}), graphics = {Text(extent = {{-100, -98}, {100, -120}}, lineColor = {0, 0, 0}, textString = "%name"), Rectangle(extent = {{-10, 80}, {10, -80}}, lineColor = {0, 0, 255}, pattern = LinePattern.None, fillColor = {0, 0, 255}, fillPattern = FillPattern.Solid)}),
            Documentation(revisions = "<html>
</html>"));
        end DCbus;

        model ACbus "交流总线"
          output Modelica.SIunits.Voltage v(stateSelect = StateSelect.never);
          ENN.Interfaces.Electrical.Pin_AC term "Terminal" annotation(
            Placement(transformation(extent = {{-7, -60}, {7, 60}}, rotation = 0)));
        equation
          term.i = 0;
          v = term.v;
          annotation(
            Icon(graphics = {Rectangle(extent = {{-10, 80}, {10, -80}}, lineColor = {0, 140, 72}, fillColor = {0, 140, 72}, fillPattern = FillPattern.Solid)}),
            Documentation(revisions = "<html>
</html>"));
        end ACbus;
      end Internal;
    end Components;
  end RenewableEnergy;

  package Valves
    model MatValve "物料阀"
      Interfaces.Materials.Port_a mat_out annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      replaceable Modelica.Blocks.Tables.CombiTable1D m_flow_mat(table = [0.03, 0.1; 0.2, 0.6; 0.4, 1.2; 0.6, 1.8; 0.8, 2.4; 1, 3; 1.2, 3.6]) "计算原料质量流量（kg/s），负载率 vs 原料质量流量（kg/s）" annotation(
        Placement(transformation(extent = {{-52, -10}, {-32, 10}})));
      Modelica.Blocks.Interfaces.RealInput load "Connector of Real input signals" annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
      Modelica.Blocks.Interfaces.RealOutput m_mat "输出原料质量流量（kg/s）" annotation(
        Placement(transformation(extent = {{100, 40}, {120, 60}})));
      Sources.MatSource1 solidSource1_1 annotation(
        Placement(transformation(extent = {{40, -40}, {60, -20}})));
    equation
      connect(load, m_flow_mat.u[1]) annotation(
        Line(points = {{-120, 0}, {-54, 0}}, color = {0, 0, 127}));
      connect(m_flow_mat.y[1], m_mat) annotation(
        Line(points = {{-31, 0}, {0, 0}, {0, 50}, {110, 50}}, color = {0, 0, 127}));
      connect(solidSource1_1.port_a, mat_out) annotation(
        Line(points = {{60, -30}, {80, -30}, {80, 0}, {100, 0}}, color = {255, 170, 255}));
      connect(m_flow_mat.y[1], solidSource1_1.m_in) annotation(
        Line(points = {{-31, 0}, {0, 0}, {0, -30}, {38, -30}}, color = {0, 0, 127}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {28, 108, 200}), Text(extent = {{-80, 36}, {80, -24}}, lineColor = {28, 108, 200}, textString = "%name")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end MatValve;

    model ComAirValve "压缩空气阀"
      Interfaces.Air.Port_a air_out annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Sources.AirSource1 solidSink annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}})));
      Modelica.Blocks.Interfaces.RealInput m_in "Connector of Real input signals" annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
      replaceable Modelica.Blocks.Tables.CombiTable1D m_flow_O2(table = [0.01508, 0.1724; 0.08915, 1.019; 0.17567, 2.008; 0.25962, 2.9674; 0.34104, 3.8981; 0.42, 4.8006; 0.504, 5.7607]) "计算纯氧质量流量（kg/s），天然气质量流量（kg/s）vs纯氧质量流量（kg/s）" annotation(
        Placement(transformation(extent = {{-70, -10}, {-50, 10}})));
    equation
      connect(m_flow_O2.u[1], m_in) annotation(
        Line(points = {{-72, 0}, {-120, 0}}, color = {0, 0, 127}));
      connect(solidSink.port_a, air_out) annotation(
        Line(points = {{10, 0}, {100, 0}}, color = {0, 127, 255}));
      connect(m_flow_O2.y[1], solidSink.m_in) annotation(
        Line(points = {{-49, 0}, {-12, 0}}, color = {0, 0, 127}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {28, 108, 200}), Text(extent = {{-80, 36}, {80, -24}}, lineColor = {28, 108, 200}, textString = "%name")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end ComAirValve;

    model GasValve "燃料阀"
      Interfaces.Gas.Port_a gas_out annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput dmMat "Connector of Real input signals" annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
      replaceable Modelica.Blocks.Tables.CombiTable1D m_flow_gas(table = [0.1, 0.009996409; 0.6, 0.059092072; 1.2, 0.11643758; 1.8, 0.172075242; 2.4, 0.22604301; 3, 0.278378092; 3.6, 0.33405371]) "计算天然气质量流量（kg/s），原料质量流量（kg/s）vs天然气质量流量（kg/s）" annotation(
        Placement(transformation(extent = {{-72, -10}, {-52, 10}})));
      Modelica.Blocks.Interfaces.RealOutput m_gas "天然气质量流量" annotation(
        Placement(transformation(extent = {{100, 40}, {120, 60}})));
      Sources.GasSource1 gasSource1_1 annotation(
        Placement(transformation(extent = {{-12, -10}, {8, 10}})));
    equation
      connect(m_flow_gas.u[1], dmMat) annotation(
        Line(points = {{-74, 0}, {-120, 0}}, color = {0, 0, 127}));
      connect(m_flow_gas.y[1], m_gas) annotation(
        Line(points = {{-51, 0}, {-40, 0}, {-40, 50}, {110, 50}}, color = {0, 0, 127}));
      connect(gasSource1_1.port_a, gas_out) annotation(
        Line(points = {{8, 0}, {100, 0}}, color = {255, 170, 85}));
      connect(m_flow_gas.y[1], gasSource1_1.m_in) annotation(
        Line(points = {{-51, 0}, {-14, 0}}, color = {0, 0, 127}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {28, 108, 200}), Text(extent = {{-80, 36}, {80, -24}}, lineColor = {28, 108, 200}, textString = "%name")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end GasValve;

    model AirValve "空气阀"
      Interfaces.Air.Port_a air_out annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Sources.AirSource1 solidSink annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}})));
      Modelica.Blocks.Interfaces.RealInput dmMat "Connector of Real input signals" annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
      replaceable Modelica.Blocks.Tables.CombiTable1D m_flow_air(table = [0.1, 0.222379212; 0.6, 1.31455692; 1.2, 2.590259941; 1.8, 3.827970356; 2.4, 5.028532488; 3, 6.192774; 3.6, 7.4313288]) "计算空气质量流量（kg/s），原料质量流量（kg/s）vs 空气质量流量（kg/s）" annotation(
        Placement(transformation(extent = {{-60, -10}, {-40, 10}})));
    equation
      connect(m_flow_air.u[1], dmMat) annotation(
        Line(points = {{-62, 0}, {-120, 0}}, color = {0, 0, 127}));
      connect(solidSink.port_a, air_out) annotation(
        Line(points = {{10, 0}, {100, 0}}, color = {0, 127, 255}));
      connect(m_flow_air.y[1], solidSink.m_in) annotation(
        Line(points = {{-39, 0}, {-12, 0}}, color = {0, 0, 127}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {28, 108, 200}), Text(extent = {{-80, 36}, {80, -24}}, lineColor = {28, 108, 200}, textString = "%name")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end AirValve;
  end Valves;

  package Pipes
    model AirPipe "理想空气管道"
      //parameter Integer n_ports=1   "Number of fluid ports" annotation(Evaluate=true, Dialog(connectorSizing=true, tab="General",group="Ports"));
      Interfaces.Air.Port_a air_in "空气输入接口" annotation(
        Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
      Interfaces.Air.Port_b air_out "空气输出接口" annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
    equation
      air_in.m_flow + air_out.m_flow = 0;
      air_in.H_flow + air_out.H_flow = 0;
      air_in.p = air_out.p;
      air_in.h = air_out.h;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-80, 20}, {80, -20}}, lineColor = {28, 108, 200}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 80}, {150, 40}}, textString = "%name", lineColor = {0, 0, 255})}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end AirPipe;

    model Pipe "Pipe with optional heat exchange"
      extends Interfaces.Water_media.TwoPort;
      parameter Boolean useHeatPort = false "=true, if HeatPort is enabled" annotation(
        Evaluate = true,
        HideResult = true,
        choices(checkBox = true));
      parameter Modelica.SIunits.Length h_g(start = 0) "Geodetic height (height difference from flowPort_a to flowPort_b)";
      parameter Modelica.SIunits.Acceleration g(final min = 0) = Modelica.Constants.g_n "Gravitation";
      Interfaces.Thermal.HeatPort heatPort(T = T_q, Q_flow = Q_flowHeatPort) if useHeatPort annotation(
        Placement(transformation(extent = {{-10, -110}, {10, -90}})));
      Modelica.SIunits.Pressure pressureDrop;
      Real power_in(unit = "kWh") "进口年累计能量";
      Real power_out(unit = "kWh") "出口年累计能量";
    protected
      Modelica.SIunits.HeatFlowRate Q_flowHeatPort "Heat flow at conditional heatPort";
    equation
      pressureDrop = 10366 * V_flow ^ 2;
      if not useHeatPort then
        Q_flowHeatPort = 0;
      end if;
// coupling with FrictionModel
      dp = pressureDrop + medium.rho * g * h_g;
// energy exchange with medium
      Q_flow = Q_flowHeatPort;
      power_out = power_in;
      der(power_in) = abs(port_a.H_flow / 1000);
      annotation(
        Documentation(info = "<html>
<p>Pipe with optional heat exchange.</p>
<p>
Thermodynamic equations are defined by Partials.TwoPort.
Q_flow is defined by heatPort.Q_flow (useHeatPort=true) or zero (useHeatPort=false).</p>
<p>
<strong>Note:</strong> Setting parameter m (mass of medium within pipe) to zero
leads to neglect of temperature transient cv*m*der(T).
</p>
<p>
<strong>Note:</strong> Injecting heat into a pipe with zero mass flow causes
temperature rise defined by storing heat in medium's mass.
</p>
</html>"),
        Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Polygon(visible = useHeatPort, points = {{-10, -90}, {-10, -40}, {0, -20}, {10, -40}, {10, -90}, {-10, -90}}, lineColor = {255, 0, 0}), Text(extent = {{-150, 80}, {150, 40}}, textString = "%name", lineColor = {0, 0, 255}), Rectangle(extent = {{-80, 20}, {80, -20}}, lineColor = {85, 255, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
    end Pipe;

    model SteamPipe "理想蒸汽管道"
      //parameter Integer n_ports=1   "Number of fluid ports" annotation(Evaluate=true, Dialog(connectorSizing=true, tab="General",group="Ports"));
      Interfaces.Steam.Port_a steam_in "蒸汽输入接口" annotation(
        Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
      Interfaces.Steam.Port_b steam_out "蒸汽输出接口" annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Real power_in(unit = "kWh") "进口年累计能量";
      Real power_out(unit = "kWh") "出口年累计能量";
    equation
      steam_in.m_flow + steam_out.m_flow = 0;
      steam_in.H_flow + steam_out.H_flow = 0;
      steam_in.p = steam_out.p;
      steam_in.h = steam_out.h;
      power_out = power_in;
      der(power_in) = abs(steam_in.H_flow / 1000);
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-80, 20}, {80, -20}}, lineColor = {238, 46, 47}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 80}, {150, 40}}, textString = "%name", lineColor = {0, 0, 255})}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end SteamPipe;

    model FluePipe "理想烟气管道"
      //parameter Integer n_ports=1   "Number of fluid ports" annotation(Evaluate=true, Dialog(connectorSizing=true, tab="General",group="Ports"));
      Interfaces.FlueGas.Port_a flue_in "烟气输入接口" annotation(
        Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
      Interfaces.FlueGas.Port_b flue_out "烟气输出接口" annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Real power_in(unit = "kWh") "进口年累计能量";
      Real power_out(unit = "kWh") "出口年累计能量";
    equation
      flue_in.m_flow + flue_out.m_flow = 0;
      flue_in.H_flow + flue_out.H_flow = 0;
      flue_in.p = flue_out.p;
      flue_in.h = flue_out.h;
      power_out = power_in;
      der(power_in) = abs(flue_in.H_flow / 1000);
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-80, 20}, {80, -20}}, lineColor = {135, 135, 135}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 80}, {150, 40}}, textString = "%name", lineColor = {0, 0, 255})}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end FluePipe;
  end Pipes;

  package Tanks
    model Tank "水箱"
      extends Tanks.Internal.SinglePortBottom(final Exchange = true);
      parameter Modelica.SIunits.Area ATank(start = 1) "Cross section of tank";
      parameter Modelica.SIunits.Length hTank(start = 1) "Height of tank";
      parameter Modelica.SIunits.Pressure pAmbient(start = 0) "Ambient pressure";
      parameter Modelica.SIunits.Acceleration g(final min = 0) = Modelica.Constants.g_n "Gravitation";
      parameter Boolean useHeatPort = false "=true, if HeatPort is enabled" annotation(
        Evaluate = true,
        HideResult = true,
        choices(checkBox = true));
      Modelica.SIunits.Mass m "Mass of medium in tank";
    protected
      Modelica.SIunits.Enthalpy H "Enthalpy of medium";
      Modelica.SIunits.HeatFlowRate Q_flow "Heat flow at the optional heatPort";
    public
      Interfaces.Thermal.HeatPort heatPort(final T = T, final Q_flow = Q_flow) if useHeatPort "Optional port for cooling or heating the medium in the tank" annotation(
        Placement(transformation(extent = {{-110, -10}, {-90, 10}}), iconTransformation(extent = {{-110, -10}, {-90, 10}})));
      Modelica.Blocks.Interfaces.RealOutput level(quantity = "Length", unit = "m", start = 0) "Level of medium in tank" annotation(
        Placement(transformation(extent = {{10, -10}, {-10, 10}}, rotation = 180, origin = {110, 0})));
      Modelica.Blocks.Interfaces.RealOutput TTank(quantity = "Temperature", unit = "K", displayUnit = "degC") "Temperature of medium in tank" annotation(
        Placement(transformation(extent = {{10, -10}, {-10, 10}}, rotation = 180, origin = {110, -60})));
    equation
//output medium temperature
      TTank = T;
//optional heating/cooling
      if not useHeatPort then
        Q_flow = 0;
      end if;
//check level
      assert(level >= 0, "Tank got empty!");
      assert(level <= hTank, "Tank got full!");
//mass balance
      m = medium.rho * ATank * level;
      der(m) = flowPort.m_flow;
//energy balance
      H = m * h;
      der(H) = flowPort.H_flow + Q_flow;
//pressure at bottom
      flowPort.p = pAmbient + m * g / ATank;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-70, 80}, {70, -80}}, lineColor = {135, 135, 135}, fillColor = {215, 215, 215}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-70, 6}, {70, -82}}, lineColor = {135, 135, 135}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid)}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end Tank;

    package Internal
      partial model SinglePortBottom "Partial model of a single port at the bottom"
        parameter Modelica.Thermal.FluidHeatFlow.Media.Medium medium = Modelica.Thermal.FluidHeatFlow.Media.Medium() "Medium" annotation(
          choicesAllMatching = true);
        parameter Modelica.SIunits.Temperature T0(start = 293.15, displayUnit = "degC") "Initial temperature of medium";
        parameter Boolean T0fixed = false "Initial temperature guess value or fixed" annotation(
          choices(checkBox = true));
        output Modelica.SIunits.Temperature T_port "Temperature at flowPort_a";
        output Modelica.SIunits.Temperature T(start = T0, fixed = T0fixed) "Outlet temperature of medium";
        Interfaces.Water_media.Port_a flowPort(final medium = medium) annotation(
          Placement(transformation(extent = {{-10, -110}, {10, -90}})));
      protected
        constant Boolean Exchange = true "Exchange of medium via flowport" annotation(
          HideResult = true);
        Modelica.SIunits.SpecificEnthalpy h "Specific enthalpy in the volume";
      equation
        T_port = flowPort.h / medium.cp;
        T = h / medium.cp;
// mass flow -> ambient: mixing rule
// mass flow <- ambient: energy flow defined by ambient's temperature
        if Exchange then
          flowPort.H_flow = semiLinear(flowPort.m_flow, flowPort.h, h);
        else
          h = flowPort.h;
        end if;
        annotation(
          Documentation(info = "<html>
<p>
Partial model of single port at the bottom, defining the medium and the temperature at the port.
</p>
</html>"),
          Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Text(extent = {{-150, 140}, {150, 100}}, lineColor = {0, 0, 255}, textString = "%name")}));
      end SinglePortBottom;
    end Internal;
  end Tanks;

  package Sensors "传感器"
    model MassFLowSensor
      ENN.Interfaces.GeneralFlowPort In annotation(
        Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
      Modelica.Blocks.Interfaces.RealOutput m_flow annotation(
        Placement(transformation(extent = {{100, 42}, {120, 62}})));
      ENN.Interfaces.GeneralFlowPort Out annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
    equation
      In.p = Out.p;
      In.h = Out.h;
      In.m_flow = m_flow;
      In.m_flow + Out.m_flow = 0;
      In.H_flow + Out.H_flow = 0;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(fillColor = {245, 245, 245}, fillPattern = FillPattern.Solid, extent = {{-70.0, -70.0}, {70.0, 70.0}}), Line(points = {{0.0, 70.0}, {0.0, 40.0}}), Line(points = {{22.9, 32.8}, {40.2, 57.3}}), Line(points = {{-22.9, 32.8}, {-40.2, 57.3}}), Line(points = {{37.6, 13.7}, {65.8, 23.9}}), Line(points = {{-37.6, 13.7}, {-65.8, 23.9}}), Ellipse(lineColor = {64, 64, 64}, fillColor = {255, 255, 255}, extent = {{-12.0, -12.0}, {12.0, 12.0}}), Polygon(rotation = -17.5, fillColor = {64, 64, 64}, pattern = LinePattern.None, fillPattern = FillPattern.Solid, points = {{-5.0, 0.0}, {-2.0, 60.0}, {0.0, 65.0}, {2.0, 60.0}, {5.0, 0.0}}), Ellipse(fillColor = {64, 64, 64}, pattern = LinePattern.None, fillPattern = FillPattern.Solid, extent = {{-7.0, -7.0}, {7.0, 7.0}}), Line(points = {{0, -100}, {0, -70}}), Text(extent = {{-30, -20}, {30, -60}}, textString = "dm", lineColor = {0, 0, 0})}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end MassFLowSensor;

    model EnthalpyFLowSensor
      ENN.Interfaces.GeneralFlowPort In annotation(
        Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
      Modelica.Blocks.Interfaces.RealOutput H_flow annotation(
        Placement(transformation(extent = {{100, 42}, {120, 62}})));
      ENN.Interfaces.GeneralFlowPort Out annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
    equation
      In.p = Out.p;
      In.h = Out.h;
      In.H_flow = H_flow;
      In.m_flow + Out.m_flow = 0;
      In.H_flow + Out.H_flow = 0;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(fillColor = {245, 245, 245}, fillPattern = FillPattern.Solid, extent = {{-70.0, -70.0}, {70.0, 70.0}}), Line(points = {{0.0, 70.0}, {0.0, 40.0}}), Line(points = {{22.9, 32.8}, {40.2, 57.3}}), Line(points = {{-22.9, 32.8}, {-40.2, 57.3}}), Line(points = {{37.6, 13.7}, {65.8, 23.9}}), Line(points = {{-37.6, 13.7}, {-65.8, 23.9}}), Ellipse(lineColor = {64, 64, 64}, fillColor = {255, 255, 255}, extent = {{-12.0, -12.0}, {12.0, 12.0}}), Polygon(rotation = -17.5, fillColor = {64, 64, 64}, pattern = LinePattern.None, fillPattern = FillPattern.Solid, points = {{-5.0, 0.0}, {-2.0, 60.0}, {0.0, 65.0}, {2.0, 60.0}, {5.0, 0.0}}), Ellipse(fillColor = {64, 64, 64}, pattern = LinePattern.None, fillPattern = FillPattern.Solid, extent = {{-7.0, -7.0}, {7.0, 7.0}}), Line(points = {{0, -100}, {0, -70}}), Text(extent = {{-30, -20}, {30, -60}}, lineColor = {0, 0, 0}, textString = "dH")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end EnthalpyFLowSensor;

    model MassFlowSensor "Mass flow sensor"
      extends Interfaces.Water_media.FlowSensor(y(unit = "kg/s") "Mass flow as output signal");
    equation
      y = V_flow * medium.rho;
      annotation(
        Documentation(info = "<html>
<p>The MassFlowSensor measures the mass flow rate.</p>
<p>Thermodynamic equations are defined by Partials.FlowSensor.</p>
</html>"),
        Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Text(extent = {{-30, -20}, {30, -60}}, textString = "m")}));
    end MassFlowSensor;
  end Sensors;

  package Sinks "汇"
    model MatSink "物料汇"
      Interfaces.Materials.Port_b port_b annotation(
        Placement(transformation(extent = {{110, 10}, {90, -10}})));
      Modelica.Blocks.Interfaces.RealInput p_in annotation(
        Placement(transformation(extent = {{-140, 30}, {-100, 70}})));
      Modelica.Blocks.Interfaces.RealInput h_in annotation(
        Placement(transformation(extent = {{-140, -70}, {-100, -30}})));
    equation
      p_in = port_b.p;
      h_in = port_b.h;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "p h"), Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 170, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end MatSink;

    model AirSink "空气汇"
      Interfaces.Air.Port_b port_b annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput p_in annotation(
        Placement(transformation(extent = {{-140, 30}, {-100, 70}})));
      Modelica.Blocks.Interfaces.RealInput h_in annotation(
        Placement(transformation(extent = {{-140, -70}, {-100, -30}})));
    equation
      p_in = port_b.p;
      h_in = port_b.h;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {0, 0, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "p h")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end AirSink;

    model FuelSink "燃料汇"
      Interfaces.Gas.Port_b port_b annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput p_in annotation(
        Placement(transformation(extent = {{-140, 30}, {-100, 70}})));
      Modelica.Blocks.Interfaces.RealInput h_in annotation(
        Placement(transformation(extent = {{-140, -70}, {-100, -30}})));
    equation
      p_in = port_b.p;
      h_in = port_b.h;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 170, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "p h")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end FuelSink;

    model FlueGasSink "烟气汇"
      Interfaces.FlueGas.Port_b port_b annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput p_in annotation(
        Placement(transformation(extent = {{-140, 30}, {-100, 70}})));
      Modelica.Blocks.Interfaces.RealInput h_in annotation(
        Placement(transformation(extent = {{-140, -70}, {-100, -30}})));
    equation
      p_in = port_b.p;
      h_in = port_b.h;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {95, 95, 95}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "p h")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end FlueGasSink;

    model WaterSink "水汇"
      Interfaces.ColdWater.Port_b port_b annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput p_in annotation(
        Placement(transformation(extent = {{-140, 30}, {-100, 70}})));
      Modelica.Blocks.Interfaces.RealInput h_in annotation(
        Placement(transformation(extent = {{-140, -70}, {-100, -30}})));
    equation
      p_in = port_b.p;
      h_in = port_b.h;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {85, 255, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "p h")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end WaterSink;

    model Sink
      ENN.Interfaces.GeneralFlowPort generalFlowPort annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
    equation
      0 = generalFlowPort.p;
      0 = generalFlowPort.h;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {95, 95, 95}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 134}, {150, 94}}, textString = "%name", lineColor = {0, 0, 255})}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end Sink;

    model ElectricalSink "电汇"
      Interfaces.Electrical.Pin_AC pin annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      parameter Modelica.SIunits.Voltage V_ref = 380 "额定电压";
    equation
      pin.v = V_ref;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-80, 60}, {80, -60}}, lineColor = {0, 0, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end ElectricalSink;

    model ThermalSink "热汇"
      Interfaces.Thermal.HeatPort heatPort annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
    equation
      298.15 = heatPort.T;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {238, 46, 47}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 132}, {150, 92}}, textString = "%name", lineColor = {0, 0, 255})}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end ThermalSink;

    model Chimney "烟囱"
      Interfaces.FlueGas.Port_a flue_in "烟气输入接口" annotation(
        Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
    equation
      0 = flue_in.p;
      0 = flue_in.h;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Polygon(points = {{-20, 40}, {20, 40}, {40, -80}, {-40, -80}, {-20, 40}}, lineColor = {135, 135, 135}, fillColor = {215, 215, 215}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-20, 60}, {20, 40}}, lineColor = {135, 135, 135}, fillColor = {215, 215, 215}, fillPattern = FillPattern.Solid)}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end Chimney;

    model ProductSink
      Interfaces.Materials.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
    equation
      0 = port_a.p;
      0 = port_a.h;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {95, 95, 95}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 132}, {150, 92}}, textString = "%name", lineColor = {0, 0, 255})}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end ProductSink;
  end Sinks;

  package Sources "源"
    model MatSource1 "物料源1"
      Interfaces.Materials.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput m_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      -m_in = port_a.m_flow;
      0 = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 170, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dm")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end MatSource1;

    model MatSource2 "物料源2"
      Interfaces.Materials.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput H_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      0 = port_a.m_flow;
      -H_in = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 170, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dh")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end MatSource2;

    model AirSource1 "空气源1"
      Interfaces.Air.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput m_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      -m_in = port_a.m_flow;
      0 = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {0, 128, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dm")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end AirSource1;

    model AirSource2 "空气源2"
      Interfaces.Air.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput H_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      0 = port_a.m_flow;
      -H_in = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {0, 128, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dh")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end AirSource2;

    model GasSource1 "燃料源1"
      Interfaces.Gas.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput m_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      -m_in = port_a.m_flow;
      0 = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 170, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dm")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end GasSource1;

    model GasSource2 "燃料源2"
      Interfaces.Gas.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput H_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      0 = port_a.m_flow;
      -H_in = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 170, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dh")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end GasSource2;

    model FlueGasSource1 "烟气源1"
      Interfaces.FlueGas.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput m_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      -m_in = port_a.m_flow;
      0 = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {95, 95, 95}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dm")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end FlueGasSource1;

    model FlueGasSource2 "烟气源2"
      Interfaces.FlueGas.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput H_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      0 = port_a.m_flow;
      -H_in = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {95, 95, 95}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dh")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end FlueGasSource2;

    model waterSource1 "水源1"
      Interfaces.ColdWater.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput m_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      -m_in = port_a.m_flow;
      0 = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {85, 255, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dm")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end waterSource1;

    model waterSource2 "水源2"
      Interfaces.ColdWater.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput H_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      0 = port_a.m_flow;
      -H_in = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {85, 255, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dh")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end waterSource2;

    model SourceNone "无流源"
      Interfaces.GeneralFlowPort Port annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
    equation
      0 = Port.m_flow;
      0 = Port.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {215, 215, 215}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dm=0
dh=0")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end SourceNone;

    model ElectricalSource "电源"
      Interfaces.Electrical.Pin_AC pin annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput P_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      P_in = pin.v * pin.i;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-80, 60}, {80, -60}}, lineColor = {0, 0, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Rectangle(extent = {{-46, 24}, {48, -26}}, lineColor = {0, 0, 255}, fillColor = {0, 0, 255}, fillPattern = FillPattern.Solid)}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end ElectricalSource;

    model ThermalSource "热源"
      ENN.Interfaces.Thermal.HeatPort heatPort annotation(
        Placement(transformation(extent = {{110, -10}, {90, 10}})));
      Modelica.Blocks.Interfaces.RealInput Q_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      Q_in = heatPort.Q_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Text(extent = {{-150, 100}, {150, 60}}, textString = "%name", lineColor = {0, 0, 255}), Text(extent = {{-150, -55}, {150, -85}}, textString = "Q_flow=%Q_flow"), Line(points = {{-100, -20}, {48, -20}}, color = {191, 0, 0}, thickness = 0.5), Line(points = {{-100, 20}, {46, 20}}, color = {191, 0, 0}, thickness = 0.5), Polygon(points = {{40, 0}, {40, 40}, {70, 20}, {40, 0}}, lineColor = {191, 0, 0}, fillColor = {191, 0, 0}, fillPattern = FillPattern.Solid), Polygon(points = {{40, -40}, {40, 0}, {70, -20}, {40, -40}}, lineColor = {191, 0, 0}, fillColor = {191, 0, 0}, fillPattern = FillPattern.Solid), Rectangle(extent = {{70, 40}, {90, -40}}, lineColor = {191, 0, 0}, fillColor = {191, 0, 0}, fillPattern = FillPattern.Solid)}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end ThermalSource;

    model FuelSource "燃料源"
      Interfaces.Gas.Port_a port_a annotation(
        Placement(transformation(extent = {{110, -10}, {90, 10}})));
      Modelica.Blocks.Interfaces.RealInput m_in annotation(
        Placement(transformation(extent = {{-140, 30}, {-100, 70}})));
      Modelica.Blocks.Interfaces.RealInput H_in annotation(
        Placement(transformation(extent = {{-140, -70}, {-100, -30}})));
    equation
      m_in = port_a.m_flow;
      H_in = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 128, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dm")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end FuelSource;

    model SteamSource1 "蒸汽源1"
      Interfaces.Steam.Port_a port_a annotation(
        Placement(transformation(extent = {{110, -10}, {90, 10}})));
      Modelica.Blocks.Interfaces.RealInput m_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      -m_in = port_a.m_flow;
      0 = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dm")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end SteamSource1;

    model SteamSource2 "蒸汽源2"
      Interfaces.Steam.Port_a port_a annotation(
        Placement(transformation(extent = {{90, -10}, {110, 10}})));
      Modelica.Blocks.Interfaces.RealInput H_in annotation(
        Placement(transformation(extent = {{-140, -20}, {-100, 20}})));
    equation
      0 = port_a.m_flow;
      -H_in = port_a.H_flow;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-60, 60}, {60, -60}}, lineColor = {28, 108, 200}, fillColor = {28, 108, 200}, fillPattern = FillPattern.Solid, textString = "dh")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end SteamSource2;

    model Source "Ambient with constant properties"
      extends Interfaces.Water_media.SinglePortLeft(final Exchange = true, final T0 = 293.15, final T0fixed = false);
      parameter Boolean usePressureInput = false "Enable / disable pressure input" annotation(
        Evaluate = true,
        choices(checkBox = true));
      parameter Modelica.SIunits.Pressure constantAmbientPressure(start = 0) "Ambient pressure" annotation(
        Dialog(enable = not usePressureInput));
      parameter Boolean useTemperatureInput = false "Enable / disable temperature input" annotation(
        Evaluate = true,
        choices(checkBox = true));
      parameter Modelica.SIunits.Temperature constantAmbientTemperature(start = 293.15, displayUnit = "degC") "Ambient temperature" annotation(
        Dialog(enable = not useTemperatureInput));
      Modelica.Blocks.Interfaces.RealInput ambientPressure = pAmbient if usePressureInput annotation(
        Placement(transformation(extent = {{-20, -20}, {20, 20}}, rotation = 180, origin = {100, 60}), iconTransformation(extent = {{-20, -20}, {20, 20}}, rotation = 180, origin = {100, 60})));
      Modelica.Blocks.Interfaces.RealInput ambientTemperature = TAmbient if useTemperatureInput annotation(
        Placement(transformation(extent = {{-20, -20}, {20, 20}}, rotation = 180, origin = {100, -60}), iconTransformation(extent = {{-20, -20}, {20, 20}}, rotation = 180, origin = {100, -60})));
    protected
      Modelica.SIunits.Pressure pAmbient;
      Modelica.SIunits.Temperature TAmbient;
    equation
      if not usePressureInput then
        pAmbient = constantAmbientPressure;
      end if;
      if not useTemperatureInput then
        TAmbient = constantAmbientTemperature;
      end if;
      flowPort.p = pAmbient;
      T = TAmbient;
      annotation(
        Documentation(info = "<html>
<p>(Infinite) ambient with constant pressure and temperature.</p>
<p>Thermodynamic equations are defined by Partials.Ambient.</p>
</html>"),
        Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-90, 90}, {90, -90}}, lineColor = {255, 0, 0}, fillColor = {0, 0, 255}, fillPattern = FillPattern.Solid), Text(extent = {{20, 80}, {80, 20}}, textString = "p"), Text(extent = {{20, -20}, {80, -80}}, textString = "T")}));
    end Source;
  end Sources;

  package Interfaces
    connector FlowPort "Connector flow port"
      parameter Modelica.Thermal.FluidHeatFlow.Media.Medium medium "Medium in the connector";
      Modelica.SIunits.Pressure p;
      flow Modelica.SIunits.MassFlowRate m_flow;
      Modelica.SIunits.SpecificEnthalpy h;
      flow Modelica.SIunits.EnthalpyFlowRate H_flow;
      annotation(
        Documentation(info = "<html>
Basic definition of the connector.<br>
<strong>Variables:</strong>
<ul>
<li>Pressure p</li>
<li>flow MassFlowRate m_flow</li>
<li>Specific Enthalpy h</li>
<li>flow EnthaplyFlowRate H_flow</li>
</ul>
If ports with different media are connected, the simulation is asserted due to the check of parameter.
</html>"));
    end FlowPort;

    connector GeneralFlowPort "一般流接口"
      Units.Pressure p;
      flow Units.MassFlowRate m_flow;
      Units.SpecificEnthalpy h;
      flow Modelica.SIunits.EnthalpyFlowRate H_flow;
    end GeneralFlowPort;

    package Gas "天然气"
      connector Port_b "Hollow flow port (used downstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_b",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {255, 170, 85}, fillPattern = FillPattern.Solid, lineColor = {255, 170, 85}), Ellipse(extent = {{-30, 30}, {30, -30}}, lineColor = {255, 170, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {255, 170, 85}, fillColor = {255, 170, 85}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {255, 170, 85}, fillPattern = FillPattern.Solid, lineColor = {255, 170, 85}), Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
      end Port_b;

      connector Port_a "Filled flow port (used upstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_a",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {255, 170, 85}, fillPattern = FillPattern.Solid, lineColor = {255, 170, 85}), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {255, 170, 85}, fillColor = {255, 170, 85}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {255, 170, 85}, fillPattern = FillPattern.Solid, lineColor = {255, 170, 85})}));
      end Port_a;
    end Gas;

    package Materials "配料"
      connector Port_b "Hollow flow port (used downstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_b",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {255, 170, 255}, fillPattern = FillPattern.Solid, lineColor = {255, 170, 255}), Ellipse(extent = {{-30, 30}, {30, -30}}, lineColor = {255, 170, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {255, 170, 255}, fillColor = {255, 170, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {255, 170, 255}, fillPattern = FillPattern.Solid, lineColor = {255, 170, 255}), Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
      end Port_b;

      connector Port_a "Filled flow port (used upstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_a",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {255, 170, 255}, fillPattern = FillPattern.Solid, lineColor = {255, 170, 255}), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {255, 170, 255}, fillColor = {255, 170, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {255, 170, 255}, fillPattern = FillPattern.Solid, lineColor = {255, 170, 255})}));
      end Port_a;
    end Materials;

    package Steam "蒸汽"
      connector Port_b "Hollow flow port (used downstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_b",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {255, 0, 0}, fillPattern = FillPattern.Solid, lineColor = {255, 0, 0}), Ellipse(extent = {{-30, 30}, {30, -30}}, lineColor = {255, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {255, 0, 0}, fillColor = {255, 0, 0}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {255, 0, 0}, fillPattern = FillPattern.Solid, lineColor = {255, 0, 0}), Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
      end Port_b;

      connector Port_a "Filled flow port (used upstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_a",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {255, 0, 0}, fillPattern = FillPattern.Solid, lineColor = {255, 0, 0}), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {255, 0, 0}, fillColor = {255, 0, 0}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {255, 0, 0}, fillPattern = FillPattern.Solid, lineColor = {255, 0, 0})}));
      end Port_a;
    end Steam;

    package ColdWater "冷水"
      connector Port_b "Hollow flow port (used downstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_b",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85}), Ellipse(extent = {{-30, 30}, {30, -30}}, lineColor = {85, 255, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {85, 255, 85}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85}), Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
      end Port_b;

      connector Port_a "Filled flow port (used upstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_a",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85}), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {85, 255, 85}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85})}));
      end Port_a;
    end ColdWater;

    package FlueGas "烟气"
      connector Port_b "Hollow flow port (used downstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_b",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid, lineColor = {175, 175, 175}), Ellipse(extent = {{-30, 30}, {30, -30}}, lineColor = {215, 215, 215}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {95, 95, 95}, fillColor = {95, 95, 95}, fillPattern = FillPattern.Solid, thickness = 0.5), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid, lineColor = {175, 175, 175}), Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {215, 215, 215}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
      end Port_b;

      connector Port_a "Filled flow port (used upstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_a",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid, lineColor = {175, 175, 175}, startAngle = 0, endAngle = 360), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {95, 95, 95}, fillColor = {95, 95, 95}, fillPattern = FillPattern.Solid, thickness = 0.5), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {175, 175, 175}, fillPattern = FillPattern.Solid, lineColor = {175, 175, 175})}));
      end Port_a;
    end FlueGas;

    package Air "空气"
      connector Port_b "Hollow flow port (used downstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_b",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-30, 30}, {30, -30}}, lineColor = {0, 127, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 127, 255}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {0, 127, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
      end Port_b;

      connector Port_a "Filled flow port (used upstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_a",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 127, 255}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid)}));
      end Port_a;

      package Air_withMedia
        connector Port_a "Filled flow port (used upstream)"
          extends Interfaces.FlowPort;
          annotation(
            Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
            defaultComponentName = "port_a",
            Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
            Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 127, 255}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid)}));
        end Port_a;

        connector Port_b "Hollow flow port (used downstream)"
          extends Interfaces.FlowPort;
          annotation(
            Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
            defaultComponentName = "port_b",
            Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-30, 30}, {30, -30}}, lineColor = {0, 127, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
            Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 127, 255}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {0, 127, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
        end Port_b;
      end Air_withMedia;
    end Air;

    package Water_media "水"
      connector Port_b "Hollow flow port (used downstream)"
        extends Interfaces.FlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_b",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85}), Ellipse(extent = {{-30, 30}, {30, -30}}, lineColor = {85, 255, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {85, 255, 85}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85}), Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
      end Port_b;

      connector Port_a "Filled flow port (used upstream)"
        extends Interfaces.FlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_a",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85}), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {85, 255, 85}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85})}));
      end Port_a;

      partial model TwoPort "Partial model of two port"
        parameter Modelica.Thermal.FluidHeatFlow.Media.Medium medium = Modelica.Thermal.FluidHeatFlow.Media.Medium() "Medium in the component" annotation(
          choicesAllMatching = true);
        parameter Modelica.SIunits.Mass m(start = 1) "Mass of medium";
        parameter Modelica.SIunits.Temperature T0(start = 293.15, displayUnit = "degC") "Initial temperature of medium" annotation(
          Dialog(enable = m > Modelica.Constants.small));
        parameter Boolean T0fixed = false "Initial temperature guess value or fixed" annotation(
          choices(checkBox = true),
          Dialog(enable = m > Modelica.Constants.small));
        parameter Real tapT(final min = 0, final max = 1) = 1 "Defines temperature of heatPort between inlet and outlet temperature";
        Modelica.SIunits.Pressure dp "Pressure drop a->b";
        Modelica.SIunits.VolumeFlowRate V_flow(start = 0) "Volume flow a->b";
        Modelica.SIunits.HeatFlowRate Q_flow "Heat exchange with ambient";
        output Modelica.SIunits.Temperature T(start = T0, fixed = T0fixed) "Outlet temperature of medium";
        output Modelica.SIunits.Temperature T_a "Temperature at flowPort_a";
        output Modelica.SIunits.Temperature T_b "Temperature at flowPort_b";
        output Modelica.SIunits.TemperatureDifference dT "Temperature increase of coolant in flow direction";
        Modelica.SIunits.Temperature T_q "Temperature relevant for heat exchange with ambient";
      protected
        Modelica.SIunits.SpecificEnthalpy h(start = medium.cp * T0) "Medium's specific enthalpy";
      public
        Port_a port_a(final medium = medium) annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
        Port_b port_b(final medium = medium) annotation(
          Placement(transformation(extent = {{90, -10}, {110, 10}})));
      equation
        dp = port_a.p - port_b.p;
        V_flow = port_a.m_flow / medium.rho;
        T_a = port_a.h / medium.cp;
        T_b = port_b.h / medium.cp;
        dT = if noEvent(V_flow >= 0) then T - T_a else T_b - T;
        h = medium.cp * T;
        T_q = T - noEvent(sign(V_flow)) * (1 - tapT) * dT;
// mass balance
        port_a.m_flow + port_b.m_flow = 0;
// energy balance
        if m > Modelica.Constants.small then
          port_a.H_flow + port_b.H_flow + Q_flow = m * medium.cv * der(T);
        else
          port_a.H_flow + port_b.H_flow + Q_flow = 0;
        end if;
// mass flow a->b mixing rule at a, energy flow at b defined by medium's temperature
// mass flow b->a mixing rule at b, energy flow at a defined by medium's temperature
        port_a.H_flow = semiLinear(port_a.m_flow, port_a.h, h);
        port_b.H_flow = semiLinear(port_b.m_flow, port_b.h, h);
        annotation(
          Documentation(info = "<html>
<p>Partial model with two flowPorts.</p>
<p>Possible heat exchange with the ambient is defined by Q_flow; setting this = 0 means no energy exchange.</p>
<p>
Setting parameter m (mass of medium within pipe) to zero
leads to neglect of temperature transient cv*m*der(T).</p>
<p>Mixing rule is applied.</p>
<p>Parameter 0 &lt; tapT &lt; 1 defines temperature of heatPort between medium's inlet and outlet temperature.</p>
</html>"));
      end TwoPort;

      partial model SinglePortLeft "Partial model of a single port at the left"
        parameter Modelica.Thermal.FluidHeatFlow.Media.Medium medium = Modelica.Thermal.FluidHeatFlow.Media.Medium() "Medium" annotation(
          choicesAllMatching = true);
        parameter Modelica.SIunits.Temperature T0(start = 293.15, displayUnit = "degC") "Initial temperature of medium";
        parameter Boolean T0fixed = false "Initial temperature guess value or fixed" annotation(
          choices(checkBox = true));
        output Modelica.SIunits.Temperature T_port "Temperature at flowPort_a";
        output Modelica.SIunits.Temperature T(start = T0, fixed = T0fixed) "Outlet temperature of medium";
        Port_a flowPort(final medium = medium) annotation(
          Placement(transformation(extent = {{-110, -10}, {-90, 10}})));
      protected
        constant Boolean Exchange = true "Exchange of medium via flowport" annotation(
          HideResult = true);
        Modelica.SIunits.SpecificEnthalpy h "Specific enthalpy in the volume";
      equation
        T_port = flowPort.h / medium.cp;
        T = h / medium.cp;
// mass flow -> ambient: mixing rule
// mass flow <- ambient: energy flow defined by ambient's temperature
        if Exchange then
          flowPort.H_flow = semiLinear(flowPort.m_flow, flowPort.h, h);
        else
          h = flowPort.h;
        end if;
        annotation(
          Documentation(info = "<html>
<p>
Partial model of single port at the left, defining the medium and the temperature at the port.
</p>
</html>"),
          Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Text(extent = {{-150, 140}, {150, 100}}, lineColor = {0, 0, 255}, textString = "%name")}));
      end SinglePortLeft;

      partial model FlowSensor "Partial model of flow sensor"
        extends Modelica.Icons.RotationalSensor;
        extends Interfaces.Water_media.TwoPort(final m = 0, final T0 = 293.15, final T0fixed = false, final tapT = 1);
        Modelica.Blocks.Interfaces.RealOutput y annotation(
          Placement(transformation(origin = {0, -110}, extent = {{10, -10}, {-10, 10}}, rotation = 90)));
      equation
// no pressure drop
        dp = 0;
// no energy exchange
        Q_flow = 0;
        annotation(
          Documentation(info = "<html>
<p>Partial model for a flow sensor (mass flow/heat flow).</p>
<p>Pressure, mass flow, temperature and enthalpy flow of medium are not affected, but mixing rule is applied.</p>
</html>"),
          Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Line(points = {{-70, 0}, {-90, 0}}), Line(points = {{70, 0}, {90, 0}}), Line(points = {{0, -100}, {0, -70}}), Text(extent = {{-150, 100}, {150, 140}}, lineColor = {28, 108, 200}, textString = "%name")}));
      end FlowSensor;
    end Water_media;

    package Electrical
      connector Pin "电接口"
        Units.Frequency f "频率";
        flow Units.Power P "功率";
        annotation(
          Diagram(graphics = {Text(extent = {{-160, 110}, {40, 50}}, lineColor = {0, 0, 255}, textString = "%name"), Ellipse(extent = {{-40, 40}, {40, -40}}, lineColor = {0, 0, 255}, fillColor = {0, 0, 255}, fillPattern = FillPattern.Solid)}),
          Icon(graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 255}, fillColor = {0, 0, 255}, fillPattern = FillPattern.Solid)}));
      end Pin;

      connector Pin_AC "Pin of an electrical component"
        Modelica.SIunits.Voltage v "Potential at the pin" annotation(
          unassignedMessage = "An electrical potential cannot be uniquely calculated.
The reason could be that
- a ground object is missing (Modelica.Electrical.Analog.Basic.Ground)
  to define the zero potential of the electrical circuit, or
- a connector of an electrical component is not connected.");
        flow Modelica.SIunits.Current i "Current flowing into the pin" annotation(
          unassignedMessage = "An electrical current cannot be uniquely calculated.
The reason could be that
- a ground object is missing (Modelica.Electrical.Analog.Basic.Ground)
  to define the zero potential of the electrical circuit, or
- a connector of an electrical component is not connected.");
        annotation(
          defaultComponentName = "pin",
          Icon(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{100, 100}, {-100, -100}}, lineColor = {0, 140, 72}, fillColor = {0, 140, 72}, fillPattern = FillPattern.Solid)}),
          Diagram(coordinateSystem(preserveAspectRatio = true, extent = {{-100, -100}, {100, 100}}), graphics = {Text(extent = {{-160, 110}, {40, 50}}, lineColor = {0, 0, 255}, textString = "%name"), Ellipse(extent = {{-40, 40}, {40, -40}}, lineColor = {0, 0, 255}, fillColor = {0, 140, 72}, fillPattern = FillPattern.Solid)}),
          Documentation(revisions = "<html>
</html>", info = "<html>
<p>Pin is the basic electric connector. It includes the voltage which consists between the pin and the ground node. The ground node is the node of (any) ground device (Modelica.Electrical.Basic.Ground).</p>
<p>Furthermore, the pin includes the current, which is considered to be <b>positive</b> if it is flowing at the pin <b>into the device</b>.</p>
</html>"));
      end Pin_AC;
    end Electrical;

    package Thermal
      connector HeatPort "热接口"
        Modelica.SIunits.Temperature T "温度";
        flow Modelica.SIunits.HeatFlowRate Q_flow "热流";
        annotation(
          Diagram(graphics = {Rectangle(extent = {{-50, 50}, {50, -50}}, lineColor = {191, 0, 0}, fillColor = {191, 0, 0}, fillPattern = FillPattern.Solid), Text(extent = {{-120, 120}, {100, 60}}, lineColor = {191, 0, 0}, textString = "%name")}),
          Icon(graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {191, 0, 0}, fillColor = {191, 0, 0}, fillPattern = FillPattern.Solid)}));
      end HeatPort;
    end Thermal;

    package HotWater "热水"
      connector Port_b "Hollow flow port (used downstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_b",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85}), Ellipse(extent = {{-30, 30}, {30, -30}}, lineColor = {85, 255, 85}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {85, 255, 85}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {255, 0, 0}, lineThickness = 0.5), Ellipse(extent = {{-80, 80}, {80, -80}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid)}));
      end Port_b;

      connector Port_a "Filled flow port (used upstream)"
        extends Interfaces.GeneralFlowPort;
        annotation(
          Documentation(info = "<html>
Same as FlowPort, but icon allows to differentiate direction of flow.
</html>"),
          defaultComponentName = "port_a",
          Diagram(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-40, 40}, {40, -40}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {85, 255, 85}), Text(extent = {{-150, 110}, {150, 50}}, textString = "%name")}),
          Icon(coordinateSystem(preserveAspectRatio = false, extent = {{-100, -100}, {100, 100}}), graphics = {Ellipse(extent = {{-100, 100}, {100, -100}}, lineColor = {85, 255, 85}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid), Ellipse(extent = {{-100, 100}, {100, -100}}, fillColor = {85, 255, 85}, fillPattern = FillPattern.Solid, lineColor = {255, 0, 0}, lineThickness = 0.5)}));
      end Port_a;
    end HotWater;
  end Interfaces;

  package Media
    record Medium "Record containing media properties"
      extends Modelica.Icons.Record;
      parameter Modelica.SIunits.Density rho = 1 "Density";
      parameter Modelica.SIunits.SpecificHeatCapacity cp = 1 "Specific heat capacity at constant pressure";
      parameter Modelica.SIunits.SpecificHeatCapacity cv = 1 "Specific heat capacity at constant volume";
      parameter Modelica.SIunits.ThermalConductivity lamda = 1 "Thermal conductivity";
      parameter Modelica.SIunits.KinematicViscosity nue = 1 "Kinematic viscosity";
      annotation(
        defaultComponentPrefixes = "parameter",
        Documentation(info = "<html>
Record containing (constant) medium properties.
</html>"));
    end Medium;

    record Air_30degC "properties of air at 30 degC and 1 bar"
      extends Medium(nue = 16.3E-6, lamda = 0.0264, cv = 720, cp = 1007, rho = 1.149);
    end Air_30degC;

    record Water "properties of water at 30 degC and 1 bar"
      extends Medium(nue = 0.8E-6, lamda = 0.615, cv = 4177, cp = 4177, rho(displayUnit = "kg/dm3") = 995.6);
    end Water;
  end Media;

  package Economy
    block Analyser "经济性分析模块"
      replaceable Modelica.Blocks.Sources.RealExpression ele annotation(
        Placement(transformation(extent = {{-70, 50}, {-50, 70}})));
      replaceable Modelica.Blocks.Sources.RealExpression gas annotation(
        Placement(transformation(extent = {{-70, 10}, {-50, 30}})));
      replaceable Modelica.Blocks.Sources.RealExpression water annotation(
        Placement(transformation(extent = {{-70, -30}, {-50, -10}})));
      replaceable Modelica.Blocks.Sources.RealExpression materials annotation(
        Placement(transformation(extent = {{-70, -70}, {-50, -50}})));
      Modelica.Blocks.Interfaces.RealOutput profit annotation(
        Placement(transformation(extent = {{100, -10}, {120, 10}})));
      Modelica.Blocks.Math.Sum sum(nin = 4) annotation(
        Placement(transformation(extent = {{-10, -10}, {10, 10}})));
      replaceable EquipmentEconomy equipmentEconomy annotation(
        Placement(transformation(extent = {{70, -90}, {90, -70}})));
    equation
      connect(ele.y, sum.u[1]) annotation(
        Line(points = {{-49, 60}, {-28, 60}, {-28, -1.5}, {-12, -1.5}}, color = {0, 0, 127}));
      connect(gas.y, sum.u[2]) annotation(
        Line(points = {{-49, 20}, {-36, 20}, {-36, -0.5}, {-12, -0.5}}, color = {0, 0, 127}));
      connect(water.y, sum.u[3]) annotation(
        Line(points = {{-49, -20}, {-34, -20}, {-34, 0.5}, {-12, 0.5}}, color = {0, 0, 127}));
      connect(materials.y, sum.u[4]) annotation(
        Line(points = {{-49, -60}, {-28, -60}, {-28, 1.5}, {-12, 1.5}}, color = {0, 0, 127}));
      connect(sum.y, profit) annotation(
        Line(points = {{11, 0}, {110, 0}}, color = {0, 0, 127}));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 45)}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end Analyser;

    model EquipmentEconomy "设备固定投资"
      perEquipment perE "设备价格表" annotation(
        Placement(transformation(extent = {{-90, -70}, {-70, -50}})));
      parameter Real n1(unit = "台数") = 0 "余热热水锅炉台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n2(unit = "台数") = 0 "电力空压机台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n3(unit = "台数") = 0 "汽轮机台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n4(unit = "台数") = 0 "微燃机台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n5(unit = "台数") = 0 "余热蒸汽锅炉台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n6(unit = "台数") = 0 "电驱动空压机台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n7(unit = "台数") = 0 "凝汽机台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n8(unit = "台数") = 0 "冷却塔台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n9(unit = "台数") = 0 "变压器台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n10(unit = "台数") = 0 "热泵台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n11(unit = "台数") = 0 "水泵台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n12(unit = "台数") = 0 "风机台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n13(unit = "台数") = 0 "风电台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      parameter Real n14(unit = "台数") = 0 "光伏台数" annotation(
        Dialog(tab = "内置参数", group = "台数设置"));
      replaceable Modelica.Blocks.Sources.RealExpression wasteHotwaterBoiler "余热热水锅炉" annotation(
        Placement(transformation(extent = {{-80, 58}, {-60, 78}})));
      replaceable Modelica.Blocks.Sources.RealExpression eleTurbine "电力空压机" annotation(
        Placement(transformation(extent = {{-80, 32}, {-60, 52}})));
      replaceable Modelica.Blocks.Sources.RealExpression gasTurbine "汽轮机" annotation(
        Placement(transformation(extent = {{-80, 6}, {-60, 26}})));
      replaceable Modelica.Blocks.Sources.RealExpression microTurbine "微燃机" annotation(
        Placement(transformation(extent = {{-80, -20}, {-60, 0}})));
      replaceable Modelica.Blocks.Sources.RealExpression wasteSteamBoiler "余热蒸汽锅炉" annotation(
        Placement(transformation(extent = {{-40, 58}, {-20, 78}})));
      replaceable Modelica.Blocks.Sources.RealExpression compressor "电驱动空压机" annotation(
        Placement(transformation(extent = {{-40, 32}, {-20, 52}})));
      replaceable Modelica.Blocks.Sources.RealExpression condensingEngine "凝汽机" annotation(
        Placement(transformation(extent = {{-40, 6}, {-20, 26}})));
      replaceable Modelica.Blocks.Sources.RealExpression coolingTower "冷却塔" annotation(
        Placement(transformation(extent = {{-40, -20}, {-20, 0}})));
      replaceable Modelica.Blocks.Sources.RealExpression transformer "变压器" annotation(
        Placement(transformation(extent = {{42, 6}, {62, 26}})));
      replaceable Modelica.Blocks.Sources.RealExpression heatPump "热泵" annotation(
        Placement(transformation(extent = {{2, 58}, {22, 78}})));
      replaceable Modelica.Blocks.Sources.RealExpression pump "水泵" annotation(
        Placement(transformation(extent = {{2, 32}, {22, 52}})));
      replaceable Modelica.Blocks.Sources.RealExpression fan "风机" annotation(
        Placement(transformation(extent = {{2, 6}, {22, 26}})));
      replaceable Modelica.Blocks.Sources.RealExpression windPower "风电" annotation(
        Placement(transformation(extent = {{42, 58}, {62, 78}})));
      replaceable Modelica.Blocks.Sources.RealExpression solarPower "光伏" annotation(
        Placement(transformation(extent = {{42, 32}, {62, 52}})));
      Modelica.Blocks.Interfaces.RealOutput cost(unit = "元");
    equation
      cost = wasteHotwaterBoiler.y * perE.p1 * n1 + eleTurbine.y * perE.p2 * n2 + gasTurbine.y * perE.p3 * n3 + microTurbine.y * perE.p4 * n4 + wasteSteamBoiler.y * perE.p5 * n5 + compressor.y * perE.p6 * n6 + condensingEngine.y * perE.p7 * n7 + coolingTower.y * perE.p8 * n8 + transformer.y * perE.p9 * n9 + heatPump.y * perE.p10 * n10 + pump.y * perE.p11 * n11 + fan.y * perE.p12 * n12 + windPower.y * perE.p13 * n13 + solarPower.y * perE.p14 * n14;
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 45)}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end EquipmentEconomy;

    record perEquipment "单位装机规模设备价格表"
      extends Modelica.Icons.Record;
      parameter Real p1(unit = "元/t") = 1000 "余热热水锅炉单位装机";
      parameter Real p2(unit = "元/t") = 1000 "电力空压机单位装机";
      parameter Real p3(unit = "元/t") = 1000 "汽轮机单位装机";
      parameter Real p4(unit = "元/t") = 1000 "微燃机单位装机";
      parameter Real p5(unit = "元/t") = 1000 "余热蒸汽锅炉单位装机";
      parameter Real p6(unit = "元/t") = 1000 "电驱动空压机单位装机";
      parameter Real p7(unit = "元/t") = 1000 "凝汽机单位装机";
      parameter Real p8(unit = "元/t") = 1000 "冷却塔单位装机";
      parameter Real p9(unit = "元/t") = 1000 "变压器单位装机";
      parameter Real p10(unit = "元/t") = 1000 "热泵单位装机";
      parameter Real p11(unit = "元/t") = 1000 "水泵单位装机";
      parameter Real p12(unit = "元/t") = 1000 "风机单位装机";
      parameter Real p13(unit = "元/t") = 1000 "风电单位装机";
      parameter Real p14(unit = "元/t") = 1000 "光伏单位装机";
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false)),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end perEquipment;
  end Economy;

  package CarbonCalc "碳指标计算"
    model SCOP
      replaceable Modelica.Blocks.Sources.RealExpression ele annotation(
        Placement(transformation(extent = {{-70, 50}, {-50, 70}})));
      replaceable Modelica.Blocks.Sources.RealExpression gas annotation(
        Placement(transformation(extent = {{-70, 10}, {-50, 30}})));
      replaceable Modelica.Blocks.Sources.RealExpression water annotation(
        Placement(transformation(extent = {{-70, -30}, {-50, -10}})));
      replaceable Modelica.Blocks.Sources.RealExpression materials annotation(
        Placement(transformation(extent = {{-70, -70}, {-50, -50}})));
      Modelica.Blocks.Interfaces.RealOutput SCOP1 annotation(
        Placement(transformation(extent = {{100, 50}, {120, 70}})));
      Modelica.Blocks.Interfaces.RealOutput SCOP2 annotation(
        Placement(transformation(extent = {{100, -10}, {120, 10}})));
      Modelica.Blocks.Interfaces.RealOutput SCOP3 annotation(
        Placement(transformation(extent = {{100, -70}, {120, -50}})));
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 45)}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end SCOP;

    model CarbonFactor "加权碳排放因子计算模型"
      parameter ENN.Utilities.Types.Zones zones = Utilities.Types.Zones.East annotation(
        Dialog(group = "选择电网区域"),
        Evaluate = true);
      parameter Real factor_solar(unit = "kgCO2/kWh") = 1 "光伏发电电碳排放因子";
      parameter Real factor_wind(unit = "kgCO2/kWh") = 1 "风电电碳排放因子";
      Real factor_grid "市电碳排放因子";
      input Modelica.SIunits.Power P_solar = 0 "光伏发电功率" annotation(
        Dialog(group = "功率"));
      input Modelica.SIunits.Power P_wind = 0 "风电功率" annotation(
        Dialog(group = "功率"));
      input Modelica.SIunits.Power P_grid = 0 "电网功率" annotation(
        Dialog(group = "功率"));
      output Real factor(unit = "kgCO2/kWh") "加权平均度电碳排放因子";
    equation
      factor_grid = Utilities.Functions.CO2_emission_factor_Cal(zones);
      factor = (factor_grid * P_grid + factor_solar * P_solar + factor_wind * P_wind) / (P_solar + P_wind + P_grid);
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Rectangle(extent = {{-100, 100}, {100, -100}}, lineColor = {0, 0, 0}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, radius = 45), Text(extent = {{-100, 100}, {100, -100}}, lineColor = {28, 108, 200}, textString = "C")}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end CarbonFactor;
  end CarbonCalc;

  package Units "单位"
    type Frequency = Real(quantity = "Frequency", final unit = "Hz");
    type Power = Real(quantity = "Power", final unit = "W");
    type Pressure = Real(quantity = "Pressure", final unit = "Pa", displayUnit = "bar");
    type MassFlowRate = Real(quantity = "MassFlowRate", final unit = "kg/s");
    type SpecificEnthalpy = Real(quantity = "SpecificEnthalpy", final unit = "J/kg");
    type EnthalpyFlowRate = Real(quantity = "EnthalpyFlowRate", final unit = "W");
  end Units;

  package Icons
    model pump
      annotation(
        Icon(coordinateSystem(preserveAspectRatio = false), graphics = {Polygon(points = {{20, -70}, {60, -85}, {20, -100}, {20, -70}}, lineColor = {0, 128, 255}, fillColor = {0, 128, 255}, fillPattern = FillPattern.Solid, visible = showDesignFlowDirection), Polygon(points = {{20, -75}, {50, -85}, {20, -95}, {20, -75}}, lineColor = {255, 255, 255}, fillColor = {255, 255, 255}, fillPattern = FillPattern.Solid, visible = allowFlowReversal), Line(points = {{55, -85}, {-60, -85}}, color = {0, 128, 255}, visible = showDesignFlowDirection), Text(extent = {{-149, -114}, {151, -154}}, lineColor = {0, 0, 255}, textString = "%name"), Rectangle(extent = {{-100, 46}, {100, -46}}, fillColor = {0, 127, 255}, fillPattern = FillPattern.HorizontalCylinder), Polygon(points = {{-48, -60}, {-72, -100}, {72, -100}, {48, -60}, {-48, -60}}, lineColor = {0, 0, 255}, pattern = LinePattern.None, fillPattern = FillPattern.VerticalCylinder), Ellipse(extent = {{-80, 80}, {80, -80}}, fillPattern = FillPattern.Sphere, fillColor = {0, 100, 199}), Polygon(points = {{-28, 30}, {-28, -30}, {50, -2}, {-28, 30}}, pattern = LinePattern.None, fillPattern = FillPattern.HorizontalCylinder, fillColor = {255, 255, 255}), Rectangle(extent = {{-10, 100}, {10, 78}}, fillPattern = FillPattern.VerticalCylinder, fillColor = {95, 95, 95})}),
        Diagram(coordinateSystem(preserveAspectRatio = false)));
    end pump;
  end Icons;

  package Utilities
    package Functions
      function dayOfTheYear "Determined day of the year based on date"
        extends Modelica.Icons.Function;
        input Integer day "Day";
        input Integer month "Month";
        input Integer year "Year";
        output Integer dayOfYear "Day of the year indicated by day, month, year";
      protected
        Boolean leapYear "Indicates leap year";
      algorithm
        leapYear := if mod(year, 4) == 0 then true else false;
        dayOfYear := day;
        dayOfYear := dayOfYear + (if month > 1 then 31 else 0);
        dayOfYear := dayOfYear + (if month > 2 then 28 + (if leapYear then 1 else 0) else 0);
        dayOfYear := dayOfYear + (if month > 3 then 31 else 0);
        dayOfYear := dayOfYear + (if month > 4 then 30 else 0);
        dayOfYear := dayOfYear + (if month > 5 then 31 else 0);
        dayOfYear := dayOfYear + (if month > 6 then 30 else 0);
        dayOfYear := dayOfYear + (if month > 7 then 31 else 0);
        dayOfYear := dayOfYear + (if month > 8 then 31 else 0);
        dayOfYear := dayOfYear + (if month > 9 then 30 else 0);
        dayOfYear := dayOfYear + (if month > 10 then 31 else 0);
        dayOfYear := dayOfYear + (if month > 11 then 30 else 0);
        annotation(
          Documentation(info = "<html>
<p>Calculate the day of the year (between 1 and 365 or 366).</p >
</html>"));
      end dayOfTheYear;

      function CO2_emission_factor_Cal "市电碳排放指标，输出碳排放系数（kgCO2/kWh）"
        import ENN.Utilities.Types.Zones;
        input ENN.Utilities.Types.Zones zones;
        output Real factor;
      algorithm
        if zones == Zones.North then
          factor := 0.9419;
        elseif zones == Zones.Northeast then
          factor := 1.0826;
        elseif zones == Zones.East then
          factor := 0.7921;
        elseif zones == Zones.Central then
          factor := 0.8587;
        elseif zones == Zones.South then
          factor := 0.8042;
        else
          factor := 0.8922;
        end if;
        annotation(
          Documentation(info = "<html>
<h4>表1 各区域电网划分 </h4>
<table cellspacing=\"0\" cellpadding=\"0\" border=\"1\"><tr>
<td valign=\"top\"><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">电网名称</span> </p></td>
<td valign=\"top\"><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">覆盖省市</span> </p></td>
</tr>
<tr>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">华北区域电网</span> </p></td>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">北京市、天津市、河北省、山西省、山东省、内蒙古自治区</span> </p></td>
</tr>
<tr>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">东北区域电网</span> </p></td>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">辽宁省、吉林省、黑龙江省</span> </p></td>
</tr>
<tr>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">华东区域电网</span> </p></td>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">上海市、江苏省、浙江省、安徽省、福建省</span> </p></td>
</tr>
<tr>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">华中区域电网</span> </p></td>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">河南省、湖北省、湖南省、江西省、四川省、重庆市</span> </p></td>
</tr>
<tr>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">西北区域电网</span> </p></td>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">陕西省、甘肃省、青海省、宁夏自治区、新疆自治区</span> </p></td>
</tr>
<tr>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">南方区域电网</span> </p></td>
<td valign=\"top\"><p><span style=\"font-family: 仿宋; font-size: 11pt;\">广东省、广西自治区、云南省、贵州省、海南省</span></p></td>
</tr>
</table>
<p><br><br><h4>表2 不同区域电的碳排放因子</h4></p>
<table cellspacing=\"0\" cellpadding=\"0\" border=\"1\"><tr>
<td><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">区域</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">华北</span> </p><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">电网</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">东北</span> </p><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">电网</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">华中</span> </p><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">电网</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">华东</span> </p><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">电网</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">西北</span> </p><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">电网</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">南方</span> </p><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">电网</span> </p></td>
</tr>
<tr>
<td><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">碳排放系数</span> </p><p align=\"center\"><span style=\"font-family: 仿宋; font-size: 11pt;\">（</span></span><span style=\"font-family: Times New Roman,serif;\">kgCO<sub><span style=\"font-family: Times New Roman,serif;\">2</sub>/kWh</span><span style=\"font-family: 仿宋;\">）</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: Times New Roman,serif; font-size: 11pt;\">0.9419</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: Times New Roman,serif; font-size: 11pt;\">1.0826</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: Times New Roman,serif; font-size: 11pt;\">0.7921</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: Times New Roman,serif; font-size: 11pt;\">0.8587</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: Times New Roman,serif; font-size: 11pt;\">0.8922</span> </p></td>
<td><p align=\"center\"><span style=\"font-family: Times New Roman,serif; font-size: 11pt;\">0.8042</span> </p></td>
</tr>
</table>
</html>"));
      end CO2_emission_factor_Cal;
    end Functions;

    package Types
      type Zones = enumeration(North "华北地区，包含北京市、天津市、河北省、山西省、山东省、内蒙古自治区", Northeast "东北地区，包含辽宁省、吉林省、黑龙江省", East "华东地区，包含上海市、江苏省、浙江省、安徽省、福建省", Central "华中地区，包含河南省、湖北省、湖南省、江西省、四川省、重庆市", South "南方地区，包含广东省、广西自治区、云南省、贵州省、海南省", Northwest "西北地区，包含陕西省、甘肃省、青海省、宁夏自治区、新疆自治区");
    end Types;
  end Utilities;

  model Environment
    //   parameter Boolean use_solar = true "true:包含太阳模型" annotation(Dialog(group="Solar"),choices(checkBox=true));
    //   parameter Boolean use_wind = true "true：包含风速模型" annotation(Dialog(group="Wind"),choices(checkBox=true));
    parameter Modelica.SIunits.Pressure AirPressure = 100000 "大气压力";
    Modelica.Blocks.Interfaces.RealOutput T_amb(unit = "K") "环境温度" annotation(
      Placement(transformation(extent = {{100, 30}, {120, 50}})));
    Modelica.Blocks.Interfaces.RealOutput windSpeed "风速，m/s" annotation(
      Placement(transformation(extent = {{100, -10}, {120, 10}})));
    Modelica.Blocks.Sources.RealExpression contPressure(y = 1e5) annotation(
      Placement(transformation(extent = {{-20, -50}, {0, -30}})));
    Modelica.Blocks.Sources.Sine Tamb(amplitude = 7.5, freqHz = 1 / 3600 / 24, phase = -0.5235987755983, offset = 300) annotation(
      Dialog(group = "Solar"),
      choicesAllMatching,
      Placement(transformation(extent = {{-20, 30}, {0, 50}})));
    Modelica.Blocks.Sources.CombiTimeTable speedTable(table = [0, 9.1; 3600, 8.8; 7200, 8.4; 10800, 7.8; 14400, 6.5; 18000, 6.8; 21600, 7.3; 25200, 7; 28800, 6.3; 32400, 5.6; 36000, 5.1; 39600, 5; 43200, 5; 46800, 4.8; 50400, 4.5; 54000, 4.2; 57600, 3.7; 61200, 2.9; 64800, 2.1; 68400, 1.8; 72000, 2.1; 75600, 2.2; 79200, 1.8; 82800, 1.1; 86400, 0.7; 90000, 0.9; 93600, 1.1; 97200, 1.4; 100800, 1.6; 104400, 1.6; 108000, 1.7; 111600, 1.7; 115200, 1.6; 118800, 1.7; 122400, 2.5; 126000, 3.7; 129600, 4.9; 133200, 5.6; 136800, 5.9; 140400, 5.9; 144000, 5.8; 147600, 5.9; 151200, 5.7; 154800, 5.8; 158400, 6.2; 162000, 6.5; 165600, 6.7; 169200, 6.8; 172800, 7; 176400, 7.1; 180000, 7.3; 183600, 7.1; 187200, 6.7; 190800, 5.9; 194400, 5; 198000, 4.8; 201600, 4.9; 205200, 5; 208800, 5; 212400, 5.5; 216000, 5.8; 219600, 5.8; 223200, 5.8; 226800, 5.3; 230400, 4.3; 234000, 3.1; 237600, 3.1; 241200, 3.9; 244800, 4.7; 248400, 5.5; 252000, 6.2; 255600, 6.7; 259200, 7; 262800, 7.4; 266400, 7.8; 270000, 7.5; 273600, 7.2; 277200, 7.3; 280800, 6.9; 284400, 6.2; 288000, 5.6; 291600, 4.9; 295200, 4.3; 298800, 4; 302400, 3.9; 306000, 3.5; 309600, 2.8; 313200, 2.1; 316800, 1.5; 320400, 1.9; 324000, 3.8; 327600, 6.5; 331200, 8.9; 334800, 9.9; 338400, 9.2; 342000, 8.3; 345600, 7.5; 349200, 6.8; 352800, 6.2; 356400, 5.5; 360000, 4.9; 363600, 4.1; 367200, 3.8; 370800, 3.5; 374400, 3.2; 378000, 3.1; 381600, 2.8; 385200, 2.3; 388800, 2; 392400, 2; 396000, 2.1; 399600, 2.4; 403200, 2.8; 406800, 3.1; 410400, 3.6; 414000, 4.1; 417600, 4.5; 421200, 5.2; 424800, 5.9; 428400, 6.3; 432000, 6.3; 435600, 6.2; 439200, 6.2; 442800, 5.9; 446400, 5; 450000, 4.7; 453600, 4.4; 457200, 4.4; 460800, 4.5; 464400, 4.7; 468000, 4.6; 471600, 4; 475200, 2.9; 478800, 1.9; 482400, 1.9; 486000, 2.8; 489600, 4.2; 493200, 6; 496800, 7.3; 500400, 7.9; 504000, 8.3; 507600, 8.6; 511200, 8.9; 514800, 8.9; 518400, 8.9; 522000, 8.8; 525600, 8.6; 529200, 7.7; 532800, 7; 536400, 6.3; 540000, 5.5; 543600, 5.3; 547200, 5.4; 550800, 5.7; 554400, 5.9; 558000, 5.6; 561600, 6.2; 565200, 7.2; 568800, 8.2; 572400, 8.9; 576000, 9.5; 579600, 10.1; 583200, 10.2; 586800, 9.7; 590400, 9.1; 594000, 8.9; 597600, 8.8; 601200, 8.9; 604800, 8.5; 608400, 8.1; 612000, 7.8; 615600, 7; 619200, 6.5; 622800, 5.4; 626400, 4.3; 630000, 3.6; 633600, 3.2; 637200, 2.7; 640800, 2.2; 644400, 1.4; 648000, 0.5; 651600, 1.2; 655200, 2.4; 658800, 3.4; 662400, 4.4; 666000, 5.4; 669600, 6.2; 673200, 6.9; 676800, 7.3; 680400, 7.4; 684000, 7.3; 687600, 7; 691200, 6.5; 694800, 5.9; 698400, 5.5; 702000, 5; 705600, 4.3; 709200, 4.4; 712800, 5.2; 716400, 5.4; 720000, 5.3; 723600, 4.9; 727200, 4; 730800, 3.3; 734400, 2.6; 738000, 2; 741600, 1.9; 745200, 2.3; 748800, 2.9; 752400, 3.5; 756000, 4; 759600, 4; 763200, 3.3; 766800, 1.9; 770400, 0.8; 774000, 2.6; 777600, 4.7; 781200, 6.6; 784800, 8; 788400, 8.4; 792000, 8.2; 795600, 8.4; 799200, 8.3; 802800, 8.3; 806400, 8.5; 810000, 9; 813600, 9.4; 817200, 9.4; 820800, 9.2; 824400, 8.7; 828000, 8; 831600, 7.4; 835200, 6.4; 838800, 5.2; 842400, 4.2; 846000, 3.7; 849600, 4.4; 853200, 5.3; 856800, 6.2; 860400, 6.7; 864000, 7.2; 867600, 7.7; 871200, 8.3; 874800, 8.2; 878400, 7.9; 882000, 7; 885600, 6; 889200, 5.3; 892800, 4.7; 896400, 4.3; 900000, 4; 903600, 3.5; 907200, 3.2; 910800, 2.8; 914400, 2.6; 918000, 2.7; 921600, 3.3; 925200, 4.1; 928800, 5.1; 932400, 6.4; 936000, 7.3; 939600, 7.3; 943200, 7; 946800, 6.9; 950400, 7.1; 954000, 7.2; 957600, 7.1; 961200, 6.4; 964800, 6.1; 968400, 5.8; 972000, 5.5; 975600, 5.2; 979200, 5.2; 982800, 5.3; 986400, 5.5; 990000, 5.5; 993600, 5.7; 997200, 6.2; 1000800, 7.1; 1004400, 8.2; 1008000, 9.5; 1011600, 10.6; 1015200, 10.9; 1018800, 10.4; 1022400, 9.4; 1026000, 9.1; 1029600, 9.1; 1033200, 9; 1036800, 8.4; 1040400, 7.9; 1044000, 7.5; 1047600, 7; 1051200, 5.9; 1054800, 6.1; 1058400, 6.7; 1062000, 6.8; 1065600, 7; 1069200, 7.1; 1072800, 6.8; 1076400, 6; 1080000, 4.6; 1083600, 2.7; 1087200, 1.8; 1090800, 3.3; 1094400, 4.9; 1098000, 5.8; 1101600, 5.7; 1105200, 4.9; 1108800, 4.2; 1112400, 4.3; 1116000, 4.9; 1119600, 5.6; 1123200, 5.8; 1126800, 5.4; 1130400, 4.6; 1134000, 3.5; 1137600, 2.1; 1141200, 1; 1144800, 1.5; 1148400, 2.8; 1152000, 3.7; 1155600, 4.3; 1159200, 4.4; 1162800, 4.3; 1166400, 4.2; 1170000, 4.1; 1173600, 4.4; 1177200, 5.2; 1180800, 6; 1184400, 6.6; 1188000, 6.9; 1191600, 7; 1195200, 6.8; 1198800, 6.3; 1202400, 5.5; 1206000, 4.7; 1209600, 3.9; 1213200, 3.2; 1216800, 3; 1220400, 3; 1224000, 2.7; 1227600, 2.3; 1231200, 2.9; 1234800, 3.7; 1238400, 3.5; 1242000, 3; 1245600, 2.8; 1249200, 3.7; 1252800, 5.2; 1256400, 6.6; 1260000, 7.3; 1263600, 7.5; 1267200, 7.6; 1270800, 7.4; 1274400, 7.1; 1278000, 6.8; 1281600, 6.8; 1285200, 6.8; 1288800, 6.7; 1292400, 6.6; 1296000, 6.7; 1299600, 7; 1303200, 7.4; 1306800, 7.6; 1310400, 7.4; 1314000, 7.9; 1317600, 8; 1321200, 7.9; 1324800, 8.3; 1328400, 8.8; 1332000, 9.5; 1335600, 10; 1339200, 10.1; 1342800, 10; 1346400, 10; 1350000, 9.5; 1353600, 8.8; 1357200, 8.4; 1360800, 8.3; 1364400, 8.7; 1368000, 8.8; 1371600, 8.8; 1375200, 9.4; 1378800, 9.7; 1382400, 9.4; 1386000, 8.9; 1389600, 8.5; 1393200, 8.6; 1396800, 8.7; 1400400, 8.5; 1404000, 8; 1407600, 7.3; 1411200, 6.6; 1414800, 6.1; 1418400, 5.8; 1422000, 5.3; 1425600, 5.1; 1429200, 4.8; 1432800, 4.2; 1436400, 3.6; 1440000, 3.2; 1443600, 3.4; 1447200, 4.5; 1450800, 6.1; 1454400, 7.5; 1458000, 8.7; 1461600, 9.8; 1465200, 11.5; 1468800, 12.3; 1472400, 11.5; 1476000, 10.5; 1479600, 10.2; 1483200, 10.4; 1486800, 9.5; 1490400, 9.3; 1494000, 9.7; 1497600, 10.6; 1501200, 11; 1504800, 10.7; 1508400, 9.5; 1512000, 8.6; 1515600, 8.4; 1519200, 8.3; 1522800, 8.3; 1526400, 8.4; 1530000, 8.3; 1533600, 8.3; 1537200, 8.1; 1540800, 7.7; 1544400, 7.3; 1548000, 6.7; 1551600, 6.2; 1555200, 5.4; 1558800, 4.6; 1562400, 3.7; 1566000, 2.9; 1569600, 2.2; 1573200, 2.8; 1576800, 3.7; 1580400, 4; 1584000, 4; 1587600, 3.6; 1591200, 2.8; 1594800, 1.3; 1598400, 0.5; 1602000, 2; 1605600, 3.6; 1609200, 5; 1612800, 5.8; 1616400, 6.2; 1620000, 6.3; 1623600, 6.2; 1627200, 6; 1630800, 6.1; 1634400, 6; 1638000, 5.6; 1641600, 5.3; 1645200, 5.1; 1648800, 5.2; 1652400, 5; 1656000, 4.3; 1659600, 4.3; 1663200, 4.4; 1666800, 4.1; 1670400, 3.7; 1674000, 3.6; 1677600, 3.9; 1681200, 4.6; 1684800, 5.6; 1688400, 6.1; 1692000, 6; 1695600, 5.5; 1699200, 5.2; 1702800, 5.2; 1706400, 5.4; 1710000, 5.5; 1713600, 5.6; 1717200, 6; 1720800, 6.3; 1724400, 6.5; 1728000, 6.8; 1731600, 7; 1735200, 7.1; 1738800, 6.7; 1742400, 6.8; 1746000, 6.5; 1749600, 5.7; 1753200, 4.9; 1756800, 4.4; 1760400, 4.5; 1764000, 5.1; 1767600, 6.3; 1771200, 7.5; 1774800, 7.8; 1778400, 7.6; 1782000, 7.2; 1785600, 6.7; 1789200, 6.2; 1792800, 5.7; 1796400, 5.3; 1800000, 5.4; 1803600, 6; 1807200, 6.7; 1810800, 6.9; 1814400, 6.5; 1818000, 6.5; 1821600, 7.4; 1825200, 9; 1828800, 10.3; 1832400, 10.9; 1836000, 11; 1839600, 10.9; 1843200, 10.8; 1846800, 10.8; 1850400, 10.9; 1854000, 10.7; 1857600, 10.2; 1861200, 10; 1864800, 10; 1868400, 10.1; 1872000, 10; 1875600, 10; 1879200, 10; 1882800, 9.8; 1886400, 9.5; 1890000, 9.1; 1893600, 8.7; 1897200, 8.4; 1900800, 8.3; 1904400, 8.6; 1908000, 9.4; 1911600, 10.5; 1915200, 12.4; 1918800, 13.1; 1922400, 13; 1926000, 12.8; 1929600, 12.8; 1933200, 12.8; 1936800, 12.8; 1940400, 12.4; 1944000, 11.5; 1947600, 11.8; 1951200, 12.5; 1954800, 13.1; 1958400, 13.1; 1962000, 12.6; 1965600, 12.2; 1969200, 11.4; 1972800, 10.6; 1976400, 10.2; 1980000, 9.8; 1983600, 9.6; 1987200, 9.5; 1990800, 9.6; 1994400, 9.7; 1998000, 9.6; 2001600, 11; 2005200, 11.3; 2008800, 11.6; 2012400, 12.2; 2016000, 12.9; 2019600, 13.3; 2023200, 12.9; 2026800, 11.7; 2030400, 10.9; 2034000, 11; 2037600, 11; 2041200, 10.8; 2044800, 10.4; 2048400, 10; 2052000, 10; 2055600, 10; 2059200, 9.6; 2062800, 9.1; 2066400, 8.7; 2070000, 8.3; 2073600, 7.6; 2077200, 7; 2080800, 6.7; 2084400, 6.8; 2088000, 6.8; 2091600, 7.7; 2095200, 8; 2098800, 8.2; 2102400, 8.6; 2106000, 8.9; 2109600, 8.7; 2113200, 8; 2116800, 8.1; 2120400, 8.3; 2124000, 8.7; 2127600, 9.2; 2131200, 9.4; 2134800, 9.2; 2138400, 8.9; 2142000, 8.9; 2145600, 8.6; 2149200, 8.5; 2152800, 8.7; 2156400, 8.6; 2160000, 8.3; 2163600, 8.5; 2167200, 8.9; 2170800, 8.4; 2174400, 7.7; 2178000, 7.5; 2181600, 6.6; 2185200, 6; 2188800, 5.6; 2192400, 5.3; 2196000, 4.9; 2199600, 4.2; 2203200, 3.5; 2206800, 3.3; 2210400, 3.9; 2214000, 4.8; 2217600, 5.8; 2221200, 6.2; 2224800, 6.2; 2228400, 6.1; 2232000, 5.7; 2235600, 5.1; 2239200, 4.5; 2242800, 3.9; 2246400, 3.1; 2250000, 2; 2253600, 0.8; 2257200, 0.8; 2260800, 1.6; 2264400, 2.5; 2268000, 3.9; 2271600, 4.6; 2275200, 4.6; 2278800, 4.4; 2282400, 4.2; 2286000, 4.7; 2289600, 5.5; 2293200, 6.4; 2296800, 7.1; 2300400, 7.3; 2304000, 7.2; 2307600, 6.9; 2311200, 6.3; 2314800, 5.7; 2318400, 5.1; 2322000, 4.7; 2325600, 4.3; 2329200, 3.8; 2332800, 3.6; 2336400, 3.8; 2340000, 3.8; 2343600, 4.1; 2347200, 4.4; 2350800, 4.7; 2354400, 5.4; 2358000, 6.1; 2361600, 6.5; 2365200, 7; 2368800, 7.4; 2372400, 8.1; 2376000, 8.9; 2379600, 8.8; 2383200, 8.8; 2386800, 9.3; 2390400, 9.1; 2394000, 9.2; 2397600, 9.8; 2401200, 10.7; 2404800, 10.8; 2408400, 11.3; 2412000, 11.4; 2415600, 11.8; 2419200, 11.7; 2422800, 11.1; 2426400, 10.7; 2430000, 10.4; 2433600, 10.3; 2437200, 9.8; 2440800, 9.2; 2444400, 8.6; 2448000, 8.2; 2451600, 8.1; 2455200, 8; 2458800, 7.8; 2462400, 7.4; 2466000, 6.9; 2469600, 6.4; 2473200, 5.9; 2476800, 5.5; 2480400, 4.8; 2484000, 4.5; 2487600, 5; 2491200, 5.6; 2494800, 6.2; 2498400, 6.8; 2502000, 7.5; 2505600, 8; 2509200, 8.2; 2512800, 8.1; 2516400, 7.5; 2520000, 7.1; 2523600, 6.5; 2527200, 5.6; 2530800, 4.8; 2534400, 4.3; 2538000, 4; 2541600, 3.8; 2545200, 3.5; 2548800, 3; 2552400, 2.3; 2556000, 1.2; 2559600, 1.1; 2563200, 2.1; 2566800, 3; 2570400, 3.9; 2574000, 5.1; 2577600, 6.4; 2581200, 7.3; 2584800, 7.5; 2588400, 7.6; 2592000, 7.8; 2595600, 7.9; 2599200, 8.3; 2602800, 7.8; 2606400, 8.1; 2610000, 7.5; 2613600, 6.8; 2617200, 6.2; 2620800, 5.6; 2624400, 4.8; 2628000, 3.8; 2631600, 2.4; 2635200, 1.6; 2638800, 2.4; 2642400, 3.8; 2646000, 5.1; 2649600, 6.4; 2653200, 7.6; 2656800, 8.8; 2660400, 9.2; 2664000, 8.5; 2667600, 7.6; 2671200, 6.7; 2674800, 6; 2678400, 5.4; 2682000, 5; 2685600, 5; 2689200, 5.2; 2692800, 5.7; 2696400, 6.1; 2700000, 6.2; 2703600, 6.2; 2707200, 6.2; 2710800, 6.1; 2714400, 5.8; 2718000, 5; 2721600, 4.4; 2725200, 4.5; 2728800, 5.3; 2732400, 6.7; 2736000, 8; 2739600, 9; 2743200, 9.5; 2746800, 9; 2750400, 8.1; 2754000, 7.2; 2757600, 6.7; 2761200, 6.6; 2764800, 6.3; 2768400, 6; 2772000, 5.8; 2775600, 5.3; 2779200, 4.5; 2782800, 3.9; 2786400, 4.3; 2790000, 5; 2793600, 5.5; 2797200, 5.6; 2800800, 5.3; 2804400, 4.5; 2808000, 3.8; 2811600, 3; 2815200, 2.4; 2818800, 2.3; 2822400, 2.7; 2826000, 3.3; 2829600, 3.9; 2833200, 4.5; 2836800, 5.1; 2840400, 5.4; 2844000, 5.5; 2847600, 5.3; 2851200, 5; 2854800, 4.7; 2858400, 4.3; 2862000, 3.9; 2865600, 3.2; 2869200, 3.1; 2872800, 4.2; 2876400, 5.1; 2880000, 5.2; 2883600, 5; 2887200, 4.4; 2890800, 3.7; 2894400, 3.1; 2898000, 2.8; 2901600, 2.5; 2905200, 2; 2908800, 1.9; 2912400, 2.3; 2916000, 3.4; 2919600, 5.6; 2923200, 7.9; 2926800, 8.3; 2930400, 7.9; 2934000, 7.8; 2937600, 7.7; 2941200, 7.7; 2944800, 7.4; 2948400, 6.3; 2952000, 5.8; 2955600, 5.2; 2959200, 4.5; 2962800, 3.8; 2966400, 3.2; 2970000, 2.8; 2973600, 2.6; 2977200, 2.2; 2980800, 1.4; 2984400, 1.1; 2988000, 2; 2991600, 3.2; 2995200, 5; 2998800, 7.2; 3002400, 8.9; 3006000, 9.4; 3009600, 8.6; 3013200, 8; 3016800, 7.9; 3020400, 7.7; 3024000, 7.6; 3027600, 7.5; 3031200, 7.2; 3034800, 6.5; 3038400, 7.4; 3042000, 7.6; 3045600, 7.7; 3049200, 8.2; 3052800, 8.8; 3056400, 9.3; 3060000, 9.4; 3063600, 8.6; 3067200, 7.8; 3070800, 8.2; 3074400, 9.4; 3078000, 10.5; 3081600, 11.3; 3085200, 11.1; 3088800, 10.2; 3092400, 9.7; 3096000, 9.1; 3099600, 8.2; 3103200, 7.1; 3106800, 5.9; 3110400, 4.7; 3114000, 3.8; 3117600, 3.5; 3121200, 3.7; 3124800, 4.2; 3128400, 5.1; 3132000, 5.2; 3135600, 5; 3139200, 4.9; 3142800, 4.9; 3146400, 4.6; 3150000, 3.9; 3153600, 3.9; 3157200, 4.9; 3160800, 6; 3164400, 7; 3168000, 7.7; 3171600, 8; 3175200, 7.9; 3178800, 7.6; 3182400, 7.1; 3186000, 6.4; 3189600, 5.5; 3193200, 4.6; 3196800, 3.9; 3200400, 3.7; 3204000, 3.7; 3207600, 3.6; 3211200, 3.1; 3214800, 3.8; 3218400, 4.9; 3222000, 5.9; 3225600, 6.8; 3229200, 7.7; 3232800, 8.3; 3236400, 8.8; 3240000, 9.4; 3243600, 10.1; 3247200, 10.7; 3250800, 10.7; 3254400, 9.7; 3258000, 7.9; 3261600, 5.6; 3265200, 3.4; 3268800, 3; 3272400, 4.8; 3276000, 7; 3279600, 8.7; 3283200, 10.1; 3286800, 11.1; 3290400, 11.2; 3294000, 9.9; 3297600, 9.9; 3301200, 10.7; 3304800, 11.9; 3308400, 11; 3312000, 9.9; 3315600, 9; 3319200, 8.1; 3322800, 6.9; 3326400, 5.6; 3330000, 4.7; 3333600, 4.1; 3337200, 3.6; 3340800, 3.1; 3344400, 2.8; 3348000, 3.1; 3351600, 3.9; 3355200, 4.8; 3358800, 5.7; 3362400, 6.6; 3366000, 7; 3369600, 6.7; 3373200, 6.3; 3376800, 6.1; 3380400, 5.7; 3384000, 5.8; 3387600, 6.5; 3391200, 6.1; 3394800, 5.5; 3398400, 5; 3402000, 4.8; 3405600, 5.6; 3409200, 7.2; 3412800, 8.8; 3416400, 10.1; 3420000, 11; 3423600, 11.2; 3427200, 10.9; 3430800, 10.5; 3434400, 9.9; 3438000, 9.5; 3441600, 9.2; 3445200, 8.6; 3448800, 8; 3452400, 7.4; 3456000, 6.8; 3459600, 6.4; 3463200, 6; 3466800, 5.1; 3470400, 4.2; 3474000, 3.7; 3477600, 3; 3481200, 2.7; 3484800, 3.2; 3488400, 4.5; 3492000, 6.1; 3495600, 7.7; 3499200, 8.7; 3502800, 9.3; 3506400, 9.7; 3510000, 9.6; 3513600, 9.2; 3517200, 8.8; 3520800, 8.2; 3524400, 7.4; 3528000, 6.3; 3531600, 5.5; 3535200, 5; 3538800, 4.6; 3542400, 4.3; 3546000, 4; 3549600, 3.9; 3553200, 3.8; 3556800, 3.7; 3560400, 3.6; 3564000, 3.8; 3567600, 4.2; 3571200, 4.5; 3574800, 4.7; 3578400, 4.9; 3582000, 5.3; 3585600, 5.6; 3589200, 5.7; 3592800, 5.8; 3596400, 5.6; 3600000, 5.2; 3603600, 4.4; 3607200, 3.2; 3610800, 1.9; 3614400, 1.9; 3618000, 2.7; 3621600, 3.4; 3625200, 3.8; 3628800, 3.6; 3632400, 3.2; 3636000, 3; 3639600, 3; 3643200, 3; 3646800, 3.3; 3650400, 3.3; 3654000, 3.3; 3657600, 3.9; 3661200, 4.5; 3664800, 5.3; 3668400, 6.1; 3672000, 6.7; 3675600, 6.9; 3679200, 6.7; 3682800, 6.9; 3686400, 7.6; 3690000, 8.4; 3693600, 8.7; 3697200, 9.6; 3700800, 11; 3704400, 13.2; 3708000, 14.2; 3711600, 14.3; 3715200, 13.7; 3718800, 13.3; 3722400, 12.6; 3726000, 12.6; 3729600, 13.5; 3733200, 13; 3736800, 11.7; 3740400, 10.6; 3744000, 10; 3747600, 10; 3751200, 9.4; 3754800, 7.5; 3758400, 6.8; 3762000, 7.3; 3765600, 8.5; 3769200, 10; 3772800, 11.5; 3776400, 12.1; 3780000, 12.2; 3783600, 12.5; 3787200, 12.9; 3790800, 12.5; 3794400, 12.2; 3798000, 12; 3801600, 11.7; 3805200, 11.2; 3808800, 11; 3812400, 11.7; 3816000, 11.3; 3819600, 10.4; 3823200, 9.7; 3826800, 9.1; 3830400, 8.8; 3834000, 8.5; 3837600, 8; 3841200, 7.1; 3844800, 6.4; 3848400, 6.6; 3852000, 7.3; 3855600, 8.2; 3859200, 9.1; 3862800, 10; 3866400, 10.8; 3870000, 10.7; 3873600, 10.4; 3877200, 9.8; 3880800, 8.8; 3884400, 8.8; 3888000, 9.1; 3891600, 9.5; 3895200, 9.3; 3898800, 9.5; 3902400, 10; 3906000, 10.1; 3909600, 10.3; 3913200, 10.9; 3916800, 11.3; 3920400, 11.1; 3924000, 10.3; 3927600, 8.9; 3931200, 8.6; 3934800, 8.8; 3938400, 9.2; 3942000, 9.8; 3945600, 10.5; 3949200, 10.7; 3952800, 10.7; 3956400, 10.7; 3960000, 10.6; 3963600, 10.6; 3967200, 10.4; 3970800, 10.1; 3974400, 9.9; 3978000, 9.9; 3981600, 9.9; 3985200, 9.3; 3988800, 9.4; 3992400, 10.8; 3996000, 10.7; 3999600, 10.5; 4003200, 10.5; 4006800, 10.2; 4010400, 9; 4014000, 7.6; 4017600, 7.2; 4021200, 8; 4024800, 9.5; 4028400, 10.7; 4032000, 11.1; 4035600, 10.7; 4039200, 9.7; 4042800, 7.9; 4046400, 6.1; 4050000, 4.6; 4053600, 3.3; 4057200, 2.1; 4060800, 1.1; 4064400, 2.2; 4068000, 4.4; 4071600, 5.5; 4075200, 5.6; 4078800, 6.8; 4082400, 7.2; 4086000, 7.3; 4089600, 7.4; 4093200, 7.5; 4096800, 7.2; 4100400, 7; 4104000, 7; 4107600, 7.1; 4111200, 7; 4114800, 7.1; 4118400, 7.9; 4122000, 8.2; 4125600, 8; 4129200, 7.9; 4132800, 7.2; 4136400, 6.7; 4140000, 7.6; 4143600, 8.6; 4147200, 9.1; 4150800, 9.6; 4154400, 10.2; 4158000, 9.7; 4161600, 10.6; 4165200, 11.1; 4168800, 10.6; 4172400, 10; 4176000, 9.3; 4179600, 8.5; 4183200, 7.2; 4186800, 5.3; 4190400, 3.9; 4194000, 2.9; 4197600, 1.9; 4201200, 1.3; 4204800, 1.4; 4208400, 2.1; 4212000, 2.5; 4215600, 2.2; 4219200, 1.9; 4222800, 2.4; 4226400, 3.8; 4230000, 5.8; 4233600, 7.2; 4237200, 8; 4240800, 8.6; 4244400, 8.4; 4248000, 8.3; 4251600, 9; 4255200, 8.9; 4258800, 8.2; 4262400, 7.2; 4266000, 6.5; 4269600, 6.6; 4273200, 7.5; 4276800, 9.4; 4280400, 10.7; 4284000, 10.9; 4287600, 10.6; 4291200, 10.5; 4294800, 10.4; 4298400, 10.1; 4302000, 9.7; 4305600, 9.1; 4309200, 8.5; 4312800, 7.7; 4316400, 7; 4320000, 6.7; 4323600, 6.5; 4327200, 6.1; 4330800, 5.8; 4334400, 6.4; 4338000, 6.9; 4341600, 6.9; 4345200, 7; 4348800, 7; 4352400, 6.9; 4356000, 6.6; 4359600, 5.9; 4363200, 5; 4366800, 4.2; 4370400, 4.1; 4374000, 4.6; 4377600, 5.3; 4381200, 6; 4384800, 6.6; 4388400, 6.9; 4392000, 6.8; 4395600, 6.3; 4399200, 5.6; 4402800, 5; 4406400, 4.4; 4410000, 4; 4413600, 3.7; 4417200, 3.6; 4420800, 4.7; 4424400, 5.3; 4428000, 4.8; 4431600, 4.5; 4435200, 4.3; 4438800, 4.1; 4442400, 4.4; 4446000, 5; 4449600, 5.7; 4453200, 6.4; 4456800, 7.3; 4460400, 7.4; 4464000, 7; 4467600, 6.9; 4471200, 7.1; 4474800, 7.5; 4478400, 7.6; 4482000, 7.3; 4485600, 7; 4489200, 6.6; 4492800, 6.7; 4496400, 6.7; 4500000, 5.8; 4503600, 4.4; 4507200, 4.4; 4510800, 4.4; 4514400, 5; 4518000, 5.7; 4521600, 6.4; 4525200, 6.8; 4528800, 7; 4532400, 6.5; 4536000, 5.8; 4539600, 4.8; 4543200, 4.3; 4546800, 5.3; 4550400, 8.2; 4554000, 10.7; 4557600, 11; 4561200, 10.5; 4564800, 9.2; 4568400, 8; 4572000, 7.2; 4575600, 6.8; 4579200, 6.3; 4582800, 5.9; 4586400, 5.6; 4590000, 6.3; 4593600, 7; 4597200, 7.6; 4600800, 8.3; 4604400, 8.8; 4608000, 8.8; 4611600, 8.4; 4615200, 7.7; 4618800, 6.5; 4622400, 5.2; 4626000, 4.1; 4629600, 3.1; 4633200, 2.2; 4636800, 1.5; 4640400, 1.2; 4644000, 1.4; 4647600, 1.8; 4651200, 2.2; 4654800, 2.7; 4658400, 2.8; 4662000, 2.6; 4665600, 2.6; 4669200, 3.5; 4672800, 4.5; 4676400, 5.2; 4680000, 6.7; 4683600, 7.7; 4687200, 8.2; 4690800, 8.6; 4694400, 8.8; 4698000, 8.8; 4701600, 8.6; 4705200, 8; 4708800, 8.1; 4712400, 9; 4716000, 9.7; 4719600, 10.1; 4723200, 10.1; 4726800, 9.6; 4730400, 8.9; 4734000, 7.6; 4737600, 6.3; 4741200, 5.4; 4744800, 5.2; 4748400, 5.3; 4752000, 5.2; 4755600, 4.8; 4759200, 4.2; 4762800, 3.4; 4766400, 2.5; 4770000, 1.4; 4773600, 1; 4777200, 1.2; 4780800, 1.7; 4784400, 2.5; 4788000, 3.9; 4791600, 6.7; 4795200, 8.4; 4798800, 8.5; 4802400, 8; 4806000, 7; 4809600, 5.9; 4813200, 4.8; 4816800, 2.9; 4820400, 0.5; 4824000, 3.8; 4827600, 5.1; 4831200, 5.4; 4834800, 5.3; 4838400, 4.6; 4842000, 3.3; 4845600, 2; 4849200, 0.8; 4852800, 1.3; 4856400, 2.6; 4860000, 3.6; 4863600, 4.4; 4867200, 5.2; 4870800, 5.8; 4874400, 6; 4878000, 5.8; 4881600, 5.4; 4885200, 5; 4888800, 5.2; 4892400, 5.8; 4896000, 6.3; 4899600, 6.6; 4903200, 6; 4906800, 3.7; 4910400, 0.4; 4914000, 3.2; 4917600, 5.3; 4921200, 6.7; 4924800, 7.6; 4928400, 8.1; 4932000, 7.7; 4935600, 7.1; 4939200, 7.2; 4942800, 6.9; 4946400, 6.3; 4950000, 6; 4953600, 6.2; 4957200, 6.9; 4960800, 7.8; 4964400, 8.2; 4968000, 8.1; 4971600, 8.1; 4975200, 8; 4978800, 7.8; 4982400, 7.5; 4986000, 6.9; 4989600, 5.8; 4993200, 4.5; 4996800, 3.4; 5000400, 3.2; 5004000, 3.6; 5007600, 4.1; 5011200, 4.1; 5014800, 3.7; 5018400, 3.2; 5022000, 3.5; 5025600, 4.6; 5029200, 5.9; 5032800, 7.1; 5036400, 8.2; 5040000, 9.5; 5043600, 10.6; 5047200, 11.3; 5050800, 11.5; 5054400, 11.9; 5058000, 12.9; 5061600, 12.7; 5065200, 11.8; 5068800, 11.1; 5072400, 10.9; 5076000, 10.5; 5079600, 9.4; 5083200, 8.1; 5086800, 7.2; 5090400, 6.7; 5094000, 6.5; 5097600, 6.9; 5101200, 6.7; 5104800, 5.9; 5108400, 5; 5112000, 6.1; 5115600, 6.6; 5119200, 6.3; 5122800, 5.8; 5126400, 5.5; 5130000, 5.3; 5133600, 5; 5137200, 4.4; 5140800, 4.3; 5144400, 5; 5148000, 6.6; 5151600, 8.2; 5155200, 9.6; 5158800, 10.9; 5162400, 11.8; 5166000, 11.9; 5169600, 11.7; 5173200, 11.5; 5176800, 11.5; 5180400, 10.9; 5184000, 9.9; 5187600, 8.7; 5191200, 7.3; 5194800, 6.1; 5198400, 6.3; 5202000, 6.5; 5205600, 6; 5209200, 6.1; 5212800, 6.1; 5216400, 5.3; 5220000, 4.1; 5223600, 2.5; 5227200, 2.5; 5230800, 4.5; 5234400, 6.9; 5238000, 8.7; 5241600, 10.1; 5245200, 11.2; 5248800, 11.7; 5252400, 11.7; 5256000, 11.4; 5259600, 11; 5263200, 10.5; 5266800, 10.3; 5270400, 10.3; 5274000, 10.3; 5277600, 9.1; 5281200, 7.1; 5284800, 6.1; 5288400, 6.3; 5292000, 5.7; 5295600, 5.7; 5299200, 6.2; 5302800, 6.2; 5306400, 4.8; 5310000, 2.4; 5313600, 1.1; 5317200, 3.1; 5320800, 5.5; 5324400, 8.2; 5328000, 9.7; 5331600, 10.1; 5335200, 9.7; 5338800, 9.4; 5342400, 8.9; 5346000, 8.3; 5349600, 7.7; 5353200, 7.4; 5356800, 7.3; 5360400, 7.2; 5364000, 6.8; 5367600, 6; 5371200, 5.3; 5374800, 5; 5378400, 4.8; 5382000, 4.8; 5385600, 4.5; 5389200, 4.4; 5392800, 5.4; 5396400, 6.7; 5400000, 7.2; 5403600, 8.1; 5407200, 8.8; 5410800, 8.9; 5414400, 9; 5418000, 9.2; 5421600, 9; 5425200, 8.7; 5428800, 8.3; 5432400, 8.2; 5436000, 9.1; 5439600, 10.5; 5443200, 11.5; 5446800, 11.2; 5450400, 9.8; 5454000, 7.9; 5457600, 8.4; 5461200, 8.1; 5464800, 7.6; 5468400, 7.1; 5472000, 6.6; 5475600, 6.2; 5479200, 5.6; 5482800, 4.6; 5486400, 3.7; 5490000, 3; 5493600, 2.7; 5497200, 2.9; 5500800, 3.6; 5504400, 4.3; 5508000, 5.1; 5511600, 5.9; 5515200, 6.6; 5518800, 6.9; 5522400, 6.8; 5526000, 6.4; 5529600, 6; 5533200, 6; 5536800, 6.1; 5540400, 5.4; 5544000, 4.8; 5547600, 5.5; 5551200, 7; 5554800, 7.4; 5558400, 7.3; 5562000, 6.8; 5565600, 6; 5569200, 4.4; 5572800, 3.6; 5576400, 4; 5580000, 4.7; 5583600, 5.1; 5587200, 5.3; 5590800, 5.3; 5594400, 5.3; 5598000, 4.9; 5601600, 4.3; 5605200, 3.6; 5608800, 3.6; 5612400, 4.8; 5616000, 6.5; 5619600, 7.8; 5623200, 7.8; 5626800, 6.8; 5630400, 7.2; 5634000, 9; 5637600, 8.7; 5641200, 8.2; 5644800, 8; 5648400, 8.2; 5652000, 8.6; 5655600, 8.8; 5659200, 9.3; 5662800, 10.1; 5666400, 10.7; 5670000, 11.6; 5673600, 12.9; 5677200, 12.8; 5680800, 12; 5684400, 11.8; 5688000, 11.8; 5691600, 11.4; 5695200, 10.9; 5698800, 10.4; 5702400, 10.2; 5706000, 9.6; 5709600, 8.7; 5713200, 8.8; 5716800, 7.6; 5720400, 6.3; 5724000, 6; 5727600, 5.9; 5731200, 5.7; 5734800, 5.3; 5738400, 5; 5742000, 4.8; 5745600, 4.3; 5749200, 3.7; 5752800, 3.3; 5756400, 3.1; 5760000, 3.3; 5763600, 2.3; 5767200, 0.8; 5770800, 3.1; 5774400, 5.2; 5778000, 6.6; 5781600, 7.8; 5785200, 8.1; 5788800, 8; 5792400, 7.8; 5796000, 7.1; 5799600, 6.7; 5803200, 6.1; 5806800, 5.9; 5810400, 5.9; 5814000, 5.9; 5817600, 5.9; 5821200, 6.2; 5824800, 6.4; 5828400, 5.8; 5832000, 4.7; 5835600, 4.5; 5839200, 5.3; 5842800, 6.6; 5846400, 7.9; 5850000, 9.2; 5853600, 10.2; 5857200, 10.7; 5860800, 9.4; 5864400, 8.2; 5868000, 8; 5871600, 8; 5875200, 8.1; 5878800, 7.9; 5882400, 7.4; 5886000, 8; 5889600, 7.2; 5893200, 6.2; 5896800, 5.5; 5900400, 5.1; 5904000, 4.9; 5907600, 4.7; 5911200, 4.5; 5914800, 3.9; 5918400, 3.2; 5922000, 4.5; 5925600, 6.3; 5929200, 7.4; 5932800, 7.9; 5936400, 8; 5940000, 7.8; 5943600, 7.1; 5947200, 6.5; 5950800, 6.2; 5954400, 5.6; 5958000, 4.7; 5961600, 3.9; 5965200, 3.4; 5968800, 3.2; 5972400, 3.4; 5976000, 3.5; 5979600, 3.6; 5983200, 3.8; 5986800, 3.8; 5990400, 3.6; 5994000, 3.1; 5997600, 2.5; 6001200, 1.7; 6004800, 1.8; 6008400, 2.5; 6012000, 3.8; 6015600, 5.1; 6019200, 6.2; 6022800, 7; 6026400, 7.7; 6030000, 8.1; 6033600, 8.3; 6037200, 8.3; 6040800, 8.2; 6044400, 8; 6048000, 7.8; 6051600, 7.8; 6055200, 7.2; 6058800, 6; 6062400, 5.2; 6066000, 6.5; 6069600, 6.8; 6073200, 7; 6076800, 7.3; 6080400, 7.2; 6084000, 6.8; 6087600, 6; 6091200, 5; 6094800, 4.4; 6098400, 4.7; 6102000, 5.6; 6105600, 6.2; 6109200, 6.1; 6112800, 5.3; 6116400, 4.7; 6120000, 4.7; 6123600, 5; 6127200, 5.5; 6130800, 6.1; 6134400, 7; 6138000, 7.9; 6141600, 8; 6145200, 9.3; 6148800, 9.8; 6152400, 9; 6156000, 7.6; 6159600, 6; 6163200, 4.5; 6166800, 3.1; 6170400, 1.9; 6174000, 1.3; 6177600, 2.9; 6181200, 4.5; 6184800, 5.8; 6188400, 6.8; 6192000, 7.5; 6195600, 8; 6199200, 8.4; 6202800, 8.6; 6206400, 8.6; 6210000, 8.6; 6213600, 8.5; 6217200, 8.3; 6220800, 8.1; 6224400, 7.8; 6228000, 6.5; 6231600, 5; 6235200, 6.1; 6238800, 8.1; 6242400, 8.4; 6246000, 8.3; 6249600, 7.6; 6253200, 6.8; 6256800, 6.2; 6260400, 5.8; 6264000, 6.2; 6267600, 7; 6271200, 7.8; 6274800, 8.5; 6278400, 8.9; 6282000, 9.3; 6285600, 9.9; 6289200, 10.2; 6292800, 10.2; 6296400, 10; 6300000, 9.9; 6303600, 9.5; 6307200, 9.1; 6310800, 9; 6314400, 8.8; 6318000, 8.1; 6321600, 7.3; 6325200, 6.9; 6328800, 8; 6332400, 9; 6336000, 8.9; 6339600, 7.9; 6343200, 6.7; 6346800, 5.9; 6350400, 6.1; 6354000, 6.7; 6357600, 6.9; 6361200, 7; 6364800, 7.3; 6368400, 7.5; 6372000, 7.7; 6375600, 7.7; 6379200, 7.7; 6382800, 7.7; 6386400, 7.6; 6390000, 7.5; 6393600, 7; 6397200, 6.5; 6400800, 5.8; 6404400, 4.5; 6408000, 4.3; 6411600, 4.8; 6415200, 4.8; 6418800, 4.6; 6422400, 4.6; 6426000, 4.5; 6429600, 4.6; 6433200, 4.5; 6436800, 4.6; 6440400, 5.2; 6444000, 5.9; 6447600, 6.8; 6451200, 7.7; 6454800, 8.5; 6458400, 8.6; 6462000, 7.9; 6465600, 7.2; 6469200, 6.9; 6472800, 6.9; 6476400, 7.3; 6480000, 8; 6483600, 8.4; 6487200, 7.7; 6490800, 6.1; 6494400, 6.3; 6498000, 8.2; 6501600, 8; 6505200, 7.8; 6508800, 7.5; 6512400, 7.2; 6516000, 6.4; 6519600, 5.4; 6523200, 5.1; 6526800, 5.6; 6530400, 6.2; 6534000, 6.9; 6537600, 7; 6541200, 6.7; 6544800, 6.4; 6548400, 6.1; 6552000, 5.9; 6555600, 5.9; 6559200, 5.5; 6562800, 4.3; 6566400, 3; 6570000, 2; 6573600, 1.6; 6577200, 1.6; 6580800, 1.7; 6584400, 1.8; 6588000, 1.8; 6591600, 2.1; 6595200, 2.4; 6598800, 2.7; 6602400, 3.2; 6606000, 4.5; 6609600, 6.5; 6613200, 7.8; 6616800, 8.1; 6620400, 8.2; 6624000, 8.3; 6627600, 8.3; 6631200, 8.4; 6634800, 8.5; 6638400, 8.9; 6642000, 9.2; 6645600, 9.4; 6649200, 9.7; 6652800, 10.2; 6656400, 10.7; 6660000, 10; 6663600, 11.1; 6667200, 11.7; 6670800, 11.7; 6674400, 11.2; 6678000, 10.8; 6681600, 10.4; 6685200, 9.9; 6688800, 9.2; 6692400, 8.3; 6696000, 7; 6699600, 5.9; 6703200, 5.6; 6706800, 5.4; 6710400, 5.2; 6714000, 4.7; 6717600, 3.9; 6721200, 3; 6724800, 2.6; 6728400, 3.4; 6732000, 4.5; 6735600, 5.1; 6739200, 5.6; 6742800, 5.9; 6746400, 5.5; 6750000, 4.5; 6753600, 4.3; 6757200, 4.4; 6760800, 4.2; 6764400, 4; 6768000, 3.9; 6771600, 3.8; 6775200, 3.7; 6778800, 3.6; 6782400, 4.4; 6786000, 5.6; 6789600, 6.3; 6793200, 6.8; 6796800, 7.3; 6800400, 7.5; 6804000, 7.3; 6807600, 6.7; 6811200, 6; 6814800, 5.3; 6818400, 4.8; 6822000, 4.5; 6825600, 3.8; 6829200, 3.4; 6832800, 3.9; 6836400, 3.6; 6840000, 4.4; 6843600, 5.5; 6847200, 5.6; 6850800, 5.4; 6854400, 5.1; 6858000, 4.8; 6861600, 4.7; 6865200, 4.8; 6868800, 5.2; 6872400, 5.4; 6876000, 5.2; 6879600, 5.1; 6883200, 5.1; 6886800, 5.2; 6890400, 5; 6894000, 4.3; 6897600, 3.4; 6901200, 3; 6904800, 2.6; 6908400, 3; 6912000, 6; 6915600, 9.3; 6919200, 10.2; 6922800, 12.1; 6926400, 12.7; 6930000, 12.1; 6933600, 11.5; 6937200, 10.7; 6940800, 9.7; 6944400, 8.4; 6948000, 7; 6951600, 5.2; 6955200, 4.2; 6958800, 4.6; 6962400, 4.4; 6966000, 2.9; 6969600, 1.4; 6973200, 1.7; 6976800, 2.7; 6980400, 4.2; 6984000, 6.2; 6987600, 7.7; 6991200, 7.9; 6994800, 7.8; 6998400, 7.6; 7002000, 7.8; 7005600, 7.9; 7009200, 8.8; 7012800, 9.2; 7016400, 8.3; 7020000, 7.2; 7023600, 6.1; 7027200, 5.2; 7030800, 4.7; 7034400, 4.3; 7038000, 3.8; 7041600, 3.4; 7045200, 3.5; 7048800, 3.5; 7052400, 3.3; 7056000, 2.9; 7059600, 2.2; 7063200, 1.1; 7066800, 0.2; 7070400, 1.3; 7074000, 2.3; 7077600, 3.4; 7081200, 4.6; 7084800, 5.4; 7088400, 5.6; 7092000, 5.1; 7095600, 4.1; 7099200, 4.3; 7102800, 3.9; 7106400, 3.7; 7110000, 3.7; 7113600, 3.7; 7117200, 3.8; 7120800, 3.8; 7124400, 3.4; 7128000, 2.1; 7131600, 0.8; 7135200, 1.7; 7138800, 3; 7142400, 4.1; 7146000, 5.1; 7149600, 5.4; 7153200, 5.4; 7156800, 5.3; 7160400, 5.1; 7164000, 4.5; 7167600, 3.6; 7171200, 2.8; 7174800, 2.4; 7178400, 2; 7182000, 1.5; 7185600, 1.1; 7189200, 1.3; 7192800, 1.6; 7196400, 2.1; 7200000, 2.7; 7203600, 3.1; 7207200, 3.4; 7210800, 3.5; 7214400, 4; 7218000, 4.2; 7221600, 3.8; 7225200, 3.7; 7228800, 4.2; 7232400, 4.9; 7236000, 5.7; 7239600, 6.4; 7243200, 7.1; 7246800, 7.7; 7250400, 8.1; 7254000, 8; 7257600, 7.6; 7261200, 7.3; 7264800, 6.4; 7268400, 5.5; 7272000, 6.2; 7275600, 6.5; 7279200, 6.1; 7282800, 5.4; 7286400, 4.7; 7290000, 4.3; 7293600, 4.1; 7297200, 4.2; 7300800, 4.4; 7304400, 5.1; 7308000, 5.8; 7311600, 6.6; 7315200, 7.5; 7318800, 8.4; 7322400, 9.6; 7326000, 10.7; 7329600, 11.2; 7333200, 11; 7336800, 10.3; 7340400, 9.6; 7344000, 8.6; 7347600, 7.4; 7351200, 5.9; 7354800, 4.7; 7358400, 5.3; 7362000, 6; 7365600, 6.8; 7369200, 6.8; 7372800, 6.6; 7376400, 6.3; 7380000, 6.1; 7383600, 5.7; 7387200, 5.4; 7390800, 6.1; 7394400, 7.4; 7398000, 8.7; 7401600, 9.7; 7405200, 10.5; 7408800, 11; 7412400, 11.1; 7416000, 11; 7419600, 10.8; 7423200, 10.6; 7426800, 10.2; 7430400, 9.7; 7434000, 9.3; 7437600, 8.1; 7441200, 7.5; 7444800, 8.9; 7448400, 8.9; 7452000, 9.1; 7455600, 9.3; 7459200, 9.5; 7462800, 9.4; 7466400, 8.8; 7470000, 7.9; 7473600, 8; 7477200, 8.5; 7480800, 8.9; 7484400, 9.3; 7488000, 8.4; 7491600, 9.3; 7495200, 12.2; 7498800, 13.5; 7502400, 13.7; 7506000, 11.7; 7509600, 10.2; 7513200, 9.2; 7516800, 8.1; 7520400, 7.2; 7524000, 5.6; 7527600, 3.8; 7531200, 3; 7534800, 3.3; 7538400, 3.9; 7542000, 4.3; 7545600, 4.5; 7549200, 5; 7552800, 5.4; 7556400, 5.8; 7560000, 7.5; 7563600, 10; 7567200, 11.2; 7570800, 11.2; 7574400, 10.6; 7578000, 9.7; 7581600, 8.4; 7585200, 7.2; 7588800, 6.1; 7592400, 4.9; 7596000, 3.3; 7599600, 2.9; 7603200, 3.5; 7606800, 4.1; 7610400, 3.9; 7614000, 3.7; 7617600, 4.3; 7621200, 4.2; 7624800, 3.4; 7628400, 2.3; 7632000, 1.7; 7635600, 1.6; 7639200, 2; 7642800, 3.4; 7646400, 5.4; 7650000, 6.2; 7653600, 6.4; 7657200, 6.4; 7660800, 6.5; 7664400, 7.2; 7668000, 8.6; 7671600, 10; 7675200, 10.6; 7678800, 10.2; 7682400, 9.8; 7686000, 9.8; 7689600, 10; 7693200, 10; 7696800, 8.9; 7700400, 8.5; 7704000, 11; 7707600, 13; 7711200, 13; 7714800, 12.5; 7718400, 12.1; 7722000, 11.8; 7725600, 11.7; 7729200, 11.7; 7732800, 12; 7736400, 13; 7740000, 13.8; 7743600, 13.8; 7747200, 13.5; 7750800, 13.9; 7754400, 14.3; 7758000, 14.1; 7761600, 13.1; 7765200, 11.6; 7768800, 10; 7772400, 8.8; 7776000, 7.8; 7779600, 7.5; 7783200, 7.7; 7786800, 6; 7790400, 4.3; 7794000, 6.2; 7797600, 7.1; 7801200, 7.5; 7804800, 7.7; 7808400, 7.7; 7812000, 7.4; 7815600, 6.9; 7819200, 6.5; 7822800, 7.2; 7826400, 8.3; 7830000, 9.8; 7833600, 11.3; 7837200, 12.1; 7840800, 12.1; 7844400, 11.8; 7848000, 11.7; 7851600, 11.1; 7855200, 11; 7858800, 11.3; 7862400, 12.2; 7866000, 13.3; 7869600, 13.7; 7873200, 12.1; 7876800, 10.5; 7880400, 9.1; 7884000, 7.8; 7887600, 6.5; 7891200, 5.7; 7894800, 5.1; 7898400, 5.1; 7902000, 5.4; 7905600, 5.9; 7909200, 6.5; 7912800, 6.3; 7916400, 4.9; 7920000, 3.1; 7923600, 0.8; 7927200, 2.1; 7930800, 4.4; 7934400, 5.5; 7938000, 5.9; 7941600, 5.9; 7945200, 6.2; 7948800, 6.5; 7952400, 6.5; 7956000, 5.8; 7959600, 5.4; 7963200, 4.7; 7966800, 4.1; 7970400, 3.5; 7974000, 2.9; 7977600, 2.7; 7981200, 2.6; 7984800, 2.7; 7988400, 2.9; 7992000, 3.9; 7995600, 6; 7999200, 7.1; 8002800, 7.7; 8006400, 8.1; 8010000, 8.5; 8013600, 8.8; 8017200, 8.9; 8020800, 8.7; 8024400, 8.3; 8028000, 7.8; 8031600, 7.4; 8035200, 7.1; 8038800, 6.8; 8042400, 5.9; 8046000, 5.1; 8049600, 5.7; 8053200, 5.5; 8056800, 5.2; 8060400, 5; 8064000, 5; 8067600, 5; 8071200, 5.2; 8074800, 5.8; 8078400, 6.7; 8082000, 8; 8085600, 8.9; 8089200, 9.7; 8092800, 10.2; 8096400, 9.9; 8100000, 9.3; 8103600, 8.9; 8107200, 8.5; 8110800, 8.1; 8114400, 7.6; 8118000, 7; 8121600, 6.3; 8125200, 5.5; 8128800, 3.8; 8132400, 1.7; 8136000, 1.7; 8139600, 2.4; 8143200, 2.9; 8146800, 3; 8150400, 3.3; 8154000, 4.1; 8157600, 5.6; 8161200, 7.5; 8164800, 8.9; 8168400, 9.7; 8172000, 9.8; 8175600, 9.4; 8179200, 8.8; 8182800, 8.4; 8186400, 7.9; 8190000, 6.6; 8193600, 5.5; 8197200, 5; 8200800, 4.9; 8204400, 5; 8208000, 4.9; 8211600, 4.7; 8215200, 4; 8218800, 4; 8222400, 5.4; 8226000, 6.8; 8229600, 7.6; 8233200, 7.5; 8236800, 6.6; 8240400, 6.1; 8244000, 5.9; 8247600, 5.7; 8251200, 5.4; 8254800, 4.9; 8258400, 4.2; 8262000, 3.3; 8265600, 3.2; 8269200, 3.3; 8272800, 3; 8276400, 2.8; 8280000, 3.3; 8283600, 4.3; 8287200, 5.2; 8290800, 6.2; 8294400, 7.3; 8298000, 7.8; 8301600, 7.2; 8305200, 6.9; 8308800, 7.2; 8312400, 6.8; 8316000, 5.4; 8319600, 4.4; 8323200, 4.1; 8326800, 3.9; 8330400, 3.9; 8334000, 3.8; 8337600, 2.7; 8341200, 1.7; 8344800, 2.1; 8348400, 1.3; 8352000, 3.8; 8355600, 7.5; 8359200, 8.8; 8362800, 9.1; 8366400, 8.7; 8370000, 8.3; 8373600, 8; 8377200, 7.6; 8380800, 6.7; 8384400, 5.2; 8388000, 2.7; 8391600, 1; 8395200, 1.5; 8398800, 2; 8402400, 3.5; 8406000, 5.1; 8409600, 6.7; 8413200, 7.8; 8416800, 8.3; 8420400, 8.4; 8424000, 8.3; 8427600, 8.6; 8431200, 9.6; 8434800, 10.4; 8438400, 11.1; 8442000, 11.7; 8445600, 11.7; 8449200, 11.5; 8452800, 11.1; 8456400, 10.7; 8460000, 10.1; 8463600, 8.8; 8467200, 6; 8470800, 3.7; 8474400, 3.5; 8478000, 4.6; 8481600, 6.1; 8485200, 7.7; 8488800, 8.6; 8492400, 8.9; 8496000, 8.8; 8499600, 8.7; 8503200, 9; 8506800, 9.5; 8510400, 9.8; 8514000, 9.6; 8517600, 9.3; 8521200, 9.2; 8524800, 8.8; 8528400, 7.7; 8532000, 7.1; 8535600, 7.9; 8539200, 9.1; 8542800, 9.9; 8546400, 10.4; 8550000, 10.7; 8553600, 10.3; 8557200, 9.6; 8560800, 9.7; 8564400, 9.9; 8568000, 9.3; 8571600, 9.1; 8575200, 10; 8578800, 11.2; 8582400, 11.9; 8586000, 12.4; 8589600, 12.4; 8593200, 12.3; 8596800, 12; 8600400, 11.6; 8604000, 11.6; 8607600, 11.4; 8611200, 10.8; 8614800, 10.4; 8618400, 9.9; 8622000, 9.3; 8625600, 8.6; 8629200, 8.3; 8632800, 8.4; 8636400, 8.4; 8640000, 8.5; 8643600, 8.5; 8647200, 8.6; 8650800, 8.6; 8654400, 7.5; 8658000, 6.3; 8661600, 5.4; 8665200, 4.7; 8668800, 4.4; 8672400, 4.6; 8676000, 5.4; 8679600, 6.2; 8683200, 6.9; 8686800, 7.2; 8690400, 7.2; 8694000, 7; 8697600, 6.2; 8701200, 5.1; 8704800, 4.4; 8708400, 3.8; 8712000, 4; 8715600, 5.1; 8719200, 5.6; 8722800, 5.8; 8726400, 6.7; 8730000, 6.2; 8733600, 5.1; 8737200, 4.9; 8740800, 5; 8744400, 5.5; 8748000, 4.9; 8751600, 3.6; 8755200, 2.7; 8758800, 2; 8762400, 1.4; 8766000, 1.5; 8769600, 2.2; 8773200, 3.6; 8776800, 4.8; 8780400, 5.7; 8784000, 6.2; 8787600, 6.4; 8791200, 6.5; 8794800, 6.4; 8798400, 6.6; 8802000, 6.7; 8805600, 6.5; 8809200, 6; 8812800, 5.3; 8816400, 4.6; 8820000, 4; 8823600, 4.4; 8827200, 5.3; 8830800, 6.1; 8834400, 6.7; 8838000, 7.4; 8841600, 7.4; 8845200, 6.9; 8848800, 6.2; 8852400, 4.9; 8856000, 3.4; 8859600, 4; 8863200, 4.4; 8866800, 4.6; 8870400, 7.2; 8874000, 8.4; 8877600, 9.8; 8881200, 10.9; 8884800, 11.3; 8888400, 11; 8892000, 10.6; 8895600, 10.4; 8899200, 9.7; 8902800, 8.8; 8906400, 7.5; 8910000, 9.5; 8913600, 10; 8917200, 9.3; 8920800, 8.7; 8924400, 8.3; 8928000, 8.3; 8931600, 8.3; 8935200, 8; 8938800, 7.8; 8942400, 7.7; 8946000, 7.9; 8949600, 8.4; 8953200, 8.6; 8956800, 8.8; 8960400, 8.8; 8964000, 8.5; 8967600, 7.8; 8971200, 7.1; 8974800, 6.5; 8978400, 6.7; 8982000, 6.4; 8985600, 5.4; 8989200, 5.6; 8992800, 5.4; 8996400, 5.5; 9000000, 6.8; 9003600, 8.1; 9007200, 8.9; 9010800, 8.8; 9014400, 8.2; 9018000, 7.6; 9021600, 7.5; 9025200, 7.3; 9028800, 7; 9032400, 7.8; 9036000, 8.8; 9039600, 10; 9043200, 10.6; 9046800, 10.5; 9050400, 9.1; 9054000, 5.8; 9057600, 4.4; 9061200, 6.1; 9064800, 6.9; 9068400, 7.1; 9072000, 7.6; 9075600, 8.6; 9079200, 9.2; 9082800, 10.4; 9086400, 10.4; 9090000, 11.8; 9093600, 10.8; 9097200, 10.5; 9100800, 10.3; 9104400, 10.1; 9108000, 9.8; 9111600, 8.2; 9115200, 6.1; 9118800, 4.1; 9122400, 3.4; 9126000, 6.1; 9129600, 10.6; 9133200, 13.8; 9136800, 13.8; 9140400, 12.7; 9144000, 12; 9147600, 11.6; 9151200, 11.6; 9154800, 11.1; 9158400, 11; 9162000, 10.7; 9165600, 9.3; 9169200, 11.9; 9172800, 12.7; 9176400, 12.2; 9180000, 11.7; 9183600, 11.3; 9187200, 11.1; 9190800, 10.8; 9194400, 10.5; 9198000, 9.8; 9201600, 8.6; 9205200, 7; 9208800, 5.5; 9212400, 3.7; 9216000, 2.1; 9219600, 1.3; 9223200, 1.4; 9226800, 2.8; 9230400, 6.2; 9234000, 8.5; 9237600, 8.7; 9241200, 6.8; 9244800, 3.9; 9248400, 2.5; 9252000, 5.1; 9255600, 8.6; 9259200, 9.4; 9262800, 9.5; 9266400, 9.3; 9270000, 9; 9273600, 8.5; 9277200, 7.8; 9280800, 6.7; 9284400, 5.4; 9288000, 3.8; 9291600, 2.1; 9295200, 2.8; 9298800, 5.1; 9302400, 7.4; 9306000, 9; 9309600, 9.6; 9313200, 9.6; 9316800, 9.7; 9320400, 9.7; 9324000, 9.6; 9327600, 9.5; 9331200, 9.3; 9334800, 8.7; 9338400, 8.2; 9342000, 8.2; 9345600, 9.4; 9349200, 9.8; 9352800, 9.1; 9356400, 8.1; 9360000, 7.9; 9363600, 8.4; 9367200, 8.3; 9370800, 7.2; 9374400, 5.3; 9378000, 3.9; 9381600, 3.8; 9385200, 4.9; 9388800, 5.8; 9392400, 6; 9396000, 5.8; 9399600, 5.4; 9403200, 4.7; 9406800, 3.9; 9410400, 3.6; 9414000, 4.5; 9417600, 5.4; 9421200, 5.6; 9424800, 5.4; 9428400, 5.6; 9432000, 4.8; 9435600, 4; 9439200, 3.7; 9442800, 3.9; 9446400, 3.6; 9450000, 3.1; 9453600, 2.7; 9457200, 2.4; 9460800, 2.7; 9464400, 4.2; 9468000, 5.6; 9471600, 6.6; 9475200, 7.1; 9478800, 7.3; 9482400, 7.3; 9486000, 7.5; 9489600, 8; 9493200, 8.5; 9496800, 8.7; 9500400, 8.7; 9504000, 8.6; 9507600, 8.2; 9511200, 7; 9514800, 7.6; 9518400, 9.9; 9522000, 10.3; 9525600, 10.3; 9529200, 10.5; 9532800, 10.6; 9536400, 10.6; 9540000, 10.7; 9543600, 10.6; 9547200, 10.5; 9550800, 11.4; 9554400, 11.2; 9558000, 8.8; 9561600, 8.6; 9565200, 9.2; 9568800, 12.1; 9572400, 13.2; 9576000, 11.7; 9579600, 10.8; 9583200, 9.9; 9586800, 8.9; 9590400, 7.9; 9594000, 6.7; 9597600, 5.4; 9601200, 6.4; 9604800, 6; 9608400, 4.2; 9612000, 2.8; 9615600, 3.1; 9619200, 3.8; 9622800, 4.3; 9626400, 4.4; 9630000, 3.9; 9633600, 2.5; 9637200, 1; 9640800, 2.2; 9644400, 2.9; 9648000, 3.1; 9651600, 3.5; 9655200, 5; 9658800, 6.5; 9662400, 7.7; 9666000, 8.4; 9669600, 8.6; 9673200, 8.6; 9676800, 8.6; 9680400, 8.2; 9684000, 7.6; 9687600, 7.6; 9691200, 8.6; 9694800, 9.4; 9698400, 9.6; 9702000, 9.4; 9705600, 8.9; 9709200, 8.4; 9712800, 7.8; 9716400, 7.2; 9720000, 6.7; 9723600, 6.6; 9727200, 6.8; 9730800, 7; 9734400, 7.1; 9738000, 7.1; 9741600, 7.3; 9745200, 7.5; 9748800, 7.2; 9752400, 6.3; 9756000, 5.5; 9759600, 5.1; 9763200, 4.8; 9766800, 4.6; 9770400, 3.9; 9774000, 4.9; 9777600, 6.4; 9781200, 7.4; 9784800, 8; 9788400, 8.6; 9792000, 9.6; 9795600, 10.3; 9799200, 10.4; 9802800, 10.4; 9806400, 10.5; 9810000, 11.3; 9813600, 12.5; 9817200, 12.4; 9820800, 11.4; 9824400, 10.9; 9828000, 10.4; 9831600, 10; 9835200, 9.7; 9838800, 9.6; 9842400, 9.6; 9846000, 9.5; 9849600, 9.2; 9853200, 8; 9856800, 6.2; 9860400, 7.1; 9864000, 8.4; 9867600, 7.7; 9871200, 7; 9874800, 6.2; 9878400, 6.1; 9882000, 6.7; 9885600, 8; 9889200, 9.4; 9892800, 10.3; 9896400, 10.6; 9900000, 10.5; 9903600, 10.2; 9907200, 8.9; 9910800, 8; 9914400, 7.5; 9918000, 7; 9921600, 6.8; 9925200, 6.6; 9928800, 6.6; 9932400, 6.6; 9936000, 6.8; 9939600, 6.5; 9943200, 5.4; 9946800, 5; 9950400, 5.4; 9954000, 6.1; 9957600, 6.7; 9961200, 7; 9964800, 7; 9968400, 7; 9972000, 7; 9975600, 7.2; 9979200, 7.6; 9982800, 8.2; 9986400, 8.4; 9990000, 8; 9993600, 7.5; 9997200, 7; 10000800, 6.5; 10004400, 6.1; 10008000, 6; 10011600, 6.5; 10015200, 7.3; 10018800, 7.6; 10022400, 7.4; 10026000, 7; 10029600, 6.9; 10033200, 6.6; 10036800, 6.2; 10040400, 5.8; 10044000, 5.3; 10047600, 4.7; 10051200, 4.1; 10054800, 3.5; 10058400, 3; 10062000, 2.3; 10065600, 1.8; 10069200, 1.2; 10072800, 0.5; 10076400, 0.9; 10080000, 2.2; 10083600, 3.4; 10087200, 4.3; 10090800, 5; 10094400, 5.3; 10098000, 5.4; 10101600, 5.3; 10105200, 5.3; 10108800, 5.3; 10112400, 4.8; 10116000, 3.8; 10119600, 4; 10123200, 5.2; 10126800, 5.8; 10130400, 6.2; 10134000, 6.2; 10137600, 6.2; 10141200, 6.1; 10144800, 6.3; 10148400, 6.5; 10152000, 7.1; 10155600, 8.2; 10159200, 8.6; 10162800, 8.7; 10166400, 9; 10170000, 9; 10173600, 8.6; 10177200, 8.2; 10180800, 7.1; 10184400, 6.6; 10188000, 7.3; 10191600, 7.7; 10195200, 6.9; 10198800, 5.5; 10202400, 5; 10206000, 7.7; 10209600, 10; 10213200, 11.4; 10216800, 11.9; 10220400, 11.8; 10224000, 11.7; 10227600, 11.7; 10231200, 11.9; 10234800, 12.4; 10238400, 12.7; 10242000, 12.3; 10245600, 12.2; 10249200, 12.1; 10252800, 12; 10256400, 13.2; 10260000, 13.9; 10263600, 13.8; 10267200, 12.8; 10270800, 11.1; 10274400, 10.2; 10278000, 9.1; 10281600, 7.4; 10285200, 5.4; 10288800, 4.6; 10292400, 3.4; 10296000, 1.4; 10299600, 2.2; 10303200, 3.3; 10306800, 4; 10310400, 4.6; 10314000, 5; 10317600, 5.5; 10321200, 6.5; 10324800, 7.4; 10328400, 7.7; 10332000, 8.1; 10335600, 7.5; 10339200, 7.7; 10342800, 7.6; 10346400, 6.9; 10350000, 7.2; 10353600, 9.2; 10357200, 11.1; 10360800, 11.7; 10364400, 12.2; 10368000, 12; 10371600, 10.8; 10375200, 9.6; 10378800, 9.2; 10382400, 9.9; 10386000, 11.1; 10389600, 12.2; 10393200, 12.8; 10396800, 13.1; 10400400, 13.5; 10404000, 13.9; 10407600, 13.9; 10411200, 13.4; 10414800, 12.4; 10418400, 12; 10422000, 11.8; 10425600, 11; 10429200, 10.8; 10432800, 10.3; 10436400, 10.3; 10440000, 10.2; 10443600, 7.6; 10447200, 5.6; 10450800, 3.7; 10454400, 1.1; 10458000, 3.6; 10461600, 7.7; 10465200, 10.2; 10468800, 11.4; 10472400, 11.5; 10476000, 11; 10479600, 10.8; 10483200, 11.4; 10486800, 12.6; 10490400, 14.2; 10494000, 15.1; 10497600, 15.1; 10501200, 14.1; 10504800, 12.8; 10508400, 12.1; 10512000, 12.1; 10515600, 12.1; 10519200, 12.3; 10522800, 12.2; 10526400, 12.1; 10530000, 11.1; 10533600, 9.4; 10537200, 7.9; 10540800, 8; 10544400, 8.4; 10548000, 9.6; 10551600, 11.3; 10555200, 12.9; 10558800, 14.6; 10562400, 15.2; 10566000, 15.2; 10569600, 15.1; 10573200, 15.1; 10576800, 15; 10580400, 14.8; 10584000, 13.5; 10587600, 12.2; 10591200, 11.9; 10594800, 11.7; 10598400, 11.6; 10602000, 11.1; 10605600, 10.5; 10609200, 9.6; 10612800, 8.6; 10616400, 8; 10620000, 7.4; 10623600, 6.8; 10627200, 5.8; 10630800, 4.5; 10634400, 2.4; 10638000, 0.7; 10641600, 1; 10645200, 1.8; 10648800, 2.1; 10652400, 3; 10656000, 4.1; 10659600, 5.2; 10663200, 6.8; 10666800, 8.7; 10670400, 10.1; 10674000, 11; 10677600, 11.9; 10681200, 11.9; 10684800, 11.6; 10688400, 11.3; 10692000, 10.6; 10695600, 9.5; 10699200, 8.5; 10702800, 7.6; 10706400, 6.5; 10710000, 4.7; 10713600, 3.1; 10717200, 2.6; 10720800, 2.2; 10724400, 1.7; 10728000, 1.4; 10731600, 1.2; 10735200, 0.5; 10738800, 1.3; 10742400, 2.4; 10746000, 3.5; 10749600, 4.3; 10753200, 4.5; 10756800, 4.7; 10760400, 5.3; 10764000, 5.9; 10767600, 6.1; 10771200, 6.4; 10774800, 6.8; 10778400, 6; 10782000, 6.6; 10785600, 8.3; 10789200, 9.7; 10792800, 10.6; 10796400, 10.5; 10800000, 10.1; 10803600, 8.9; 10807200, 10.5; 10810800, 10.5; 10814400, 9.6; 10818000, 8.7; 10821600, 8.1; 10825200, 7.5; 10828800, 7.2; 10832400, 7.4; 10836000, 7.8; 10839600, 7.4; 10843200, 6.7; 10846800, 6.5; 10850400, 6.2; 10854000, 6.2; 10857600, 6.6; 10861200, 6.9; 10864800, 7.4; 10868400, 7.7; 10872000, 7.7; 10875600, 7.5; 10879200, 7.5; 10882800, 7.9; 10886400, 7.6; 10890000, 6.4; 10893600, 4.8; 10897200, 4.5; 10900800, 5.4; 10904400, 5.7; 10908000, 5.5; 10911600, 5.4; 10915200, 5.3; 10918800, 5.3; 10922400, 5.5; 10926000, 5.8; 10929600, 6.5; 10933200, 7.6; 10936800, 8.3; 10940400, 8.6; 10944000, 8.4; 10947600, 8; 10951200, 7.7; 10954800, 7.7; 10958400, 7.8; 10962000, 7.6; 10965600, 7.2; 10969200, 6.9; 10972800, 6.8; 10976400, 6; 10980000, 4.3; 10983600, 3.6; 10987200, 4.7; 10990800, 5.1; 10994400, 5.4; 10998000, 6; 11001600, 6.7; 11005200, 6.9; 11008800, 6.9; 11012400, 6.8; 11016000, 6.7; 11019600, 6.6; 11023200, 7.1; 11026800, 7.9; 11030400, 8.9; 11034000, 9.1; 11037600, 8.2; 11041200, 7.1; 11044800, 5.7; 11048400, 4.5; 11052000, 3.5; 11055600, 2.7; 11059200, 2; 11062800, 1.4; 11066400, 1; 11070000, 0.7; 11073600, 1.1; 11077200, 1.7; 11080800, 1.5; 11084400, 1.8; 11088000, 1.9; 11091600, 1.5; 11095200, 1.3; 11098800, 1.6; 11102400, 1.8; 11106000, 2.1; 11109600, 2.8; 11113200, 3.7; 11116800, 5; 11120400, 6.1; 11124000, 7.1; 11127600, 7.6; 11131200, 7.7; 11134800, 7.7; 11138400, 7.7; 11142000, 7.5; 11145600, 7; 11149200, 5.6; 11152800, 4.3; 11156400, 4.4; 11160000, 3.6; 11163600, 2.7; 11167200, 1.8; 11170800, 0.9; 11174400, 0.9; 11178000, 1.6; 11181600, 2.5; 11185200, 3.1; 11188800, 3.7; 11192400, 5.4; 11196000, 7.3; 11199600, 8.3; 11203200, 8.8; 11206800, 9.5; 11210400, 9.9; 11214000, 9.5; 11217600, 9.1; 11221200, 9; 11224800, 9.1; 11228400, 9.1; 11232000, 8.6; 11235600, 7.6; 11239200, 7.4; 11242800, 10.8; 11246400, 12; 11250000, 11.8; 11253600, 11.9; 11257200, 12.4; 11260800, 12.8; 11264400, 12.5; 11268000, 12.6; 11271600, 14.1; 11275200, 15.8; 11278800, 13.9; 11282400, 12.2; 11286000, 11; 11289600, 7.6; 11293200, 8.1; 11296800, 9.3; 11300400, 11.6; 11304000, 12.7; 11307600, 13; 11311200, 10.6; 11314800, 7.8; 11318400, 6.7; 11322000, 7.9; 11325600, 12; 11329200, 16.8; 11332800, 18.3; 11336400, 17.5; 11340000, 15.2; 11343600, 13.8; 11347200, 12.4; 11350800, 11.4; 11354400, 10.3; 11358000, 8.6; 11361600, 7.1; 11365200, 6.1; 11368800, 5.7; 11372400, 5.5; 11376000, 5.4; 11379600, 4.9; 11383200, 3.9; 11386800, 3.2; 11390400, 2.5; 11394000, 3.6; 11397600, 5; 11401200, 5.9; 11404800, 6.4; 11408400, 5.8; 11412000, 4.5; 11415600, 4.7; 11419200, 4.3; 11422800, 3.8; 11426400, 3.7; 11430000, 3.7; 11433600, 4.1; 11437200, 4.5; 11440800, 5; 11444400, 5.6; 11448000, 6.2; 11451600, 6.7; 11455200, 7.4; 11458800, 8.2; 11462400, 9.1; 11466000, 9.3; 11469600, 8.8; 11473200, 9.2; 11476800, 9.5; 11480400, 8; 11484000, 5.6; 11487600, 3.9; 11491200, 2.1; 11494800, 0.3; 11498400, 1.6; 11502000, 2.6; 11505600, 3.6; 11509200, 3.5; 11512800, 2.8; 11516400, 2.6; 11520000, 3.2; 11523600, 2.8; 11527200, 1.7; 11530800, 2; 11534400, 1.7; 11538000, 1.5; 11541600, 1.9; 11545200, 2.5; 11548800, 3.2; 11552400, 3.4; 11556000, 3.7; 11559600, 5.3; 11563200, 6; 11566800, 5.6; 11570400, 5.4; 11574000, 6.1; 11577600, 6.7; 11581200, 6.9; 11584800, 7.3; 11588400, 8.4; 11592000, 8.9; 11595600, 8.1; 11599200, 7; 11602800, 6.1; 11606400, 5.6; 11610000, 5.9; 11613600, 6.3; 11617200, 6.8; 11620800, 7; 11624400, 6.8; 11628000, 7; 11631600, 7.2; 11635200, 7; 11638800, 6.4; 11642400, 6.4; 11646000, 7.6; 11649600, 9; 11653200, 9.9; 11656800, 10.3; 11660400, 10.3; 11664000, 9.9; 11667600, 8.5; 11671200, 7.3; 11674800, 7.7; 11678400, 10.4; 11682000, 10.7; 11685600, 9.7; 11689200, 8.6; 11692800, 7.9; 11696400, 7.5; 11700000, 7.5; 11703600, 7.7; 11707200, 8; 11710800, 8.5; 11714400, 9.4; 11718000, 10.3; 11721600, 11.2; 11725200, 11.6; 11728800, 11.3; 11732400, 10.7; 11736000, 10.1; 11739600, 9.5; 11743200, 9; 11746800, 8.7; 11750400, 8.5; 11754000, 7.4; 11757600, 6.7; 11761200, 8.1; 11764800, 9.4; 11768400, 9.6; 11772000, 9; 11775600, 8.2; 11779200, 7.4; 11782800, 6.6; 11786400, 6; 11790000, 5.7; 11793600, 5.8; 11797200, 6.3; 11800800, 7; 11804400, 7.5; 11808000, 7.8; 11811600, 8; 11815200, 8.2; 11818800, 8.5; 11822400, 8.5; 11826000, 8.5; 11829600, 8.5; 11833200, 8.7; 11836800, 8.3; 11840400, 6.8; 11844000, 5.5; 11847600, 6.7; 11851200, 6.3; 11854800, 6; 11858400, 5.8; 11862000, 5.5; 11865600, 4.9; 11869200, 4.4; 11872800, 4.2; 11876400, 4.3; 11880000, 4.3; 11883600, 4.3; 11887200, 5.1; 11890800, 5.5; 11894400, 5.7; 11898000, 6.2; 11901600, 6.9; 11905200, 7.5; 11908800, 7.3; 11912400, 6.5; 11916000, 5.7; 11919600, 5.2; 11923200, 4.8; 11926800, 4; 11930400, 3; 11934000, 2.4; 11937600, 2.5; 11941200, 3.7; 11944800, 4.8; 11948400, 5.6; 11952000, 6; 11955600, 6.3; 11959200, 6.4; 11962800, 6.6; 11966400, 7.3; 11970000, 8.4; 11973600, 9.1; 11977200, 9.1; 11980800, 8.8; 11984400, 8.6; 11988000, 8.7; 11991600, 8.7; 11995200, 8.4; 11998800, 7.9; 12002400, 7.5; 12006000, 7.1; 12009600, 6.5; 12013200, 5.1; 12016800, 4.4; 12020400, 5; 12024000, 5.3; 12027600, 5.6; 12031200, 5.6; 12034800, 5.5; 12038400, 5.5; 12042000, 5.6; 12045600, 5.9; 12049200, 6.7; 12052800, 7.5; 12056400, 8.4; 12060000, 8.9; 12063600, 9.2; 12067200, 9.1; 12070800, 8.5; 12074400, 7.7; 12078000, 6.7; 12081600, 6; 12085200, 5.7; 12088800, 5.3; 12092400, 4.9; 12096000, 4.3; 12099600, 3.4; 12103200, 2.7; 12106800, 2.7; 12110400, 3.8; 12114000, 5.1; 12117600, 5.9; 12121200, 6.2; 12124800, 6.7; 12128400, 7.3; 12132000, 7.9; 12135600, 8.1; 12139200, 7.9; 12142800, 7.1; 12146400, 6.9; 12150000, 6.9; 12153600, 6.8; 12157200, 6.7; 12160800, 6.5; 12164400, 6.2; 12168000, 5.7; 12171600, 4.9; 12175200, 3.8; 12178800, 2.8; 12182400, 2.5; 12186000, 2.8; 12189600, 3; 12193200, 3.4; 12196800, 4; 12200400, 4.3; 12204000, 4.3; 12207600, 3.6; 12211200, 2.9; 12214800, 3.5; 12218400, 3.5; 12222000, 3.6; 12225600, 4.6; 12229200, 6; 12232800, 6.3; 12236400, 6; 12240000, 5.6; 12243600, 5.2; 12247200, 5; 12250800, 4.8; 12254400, 4.6; 12258000, 4.5; 12261600, 4.1; 12265200, 3.5; 12268800, 3; 12272400, 2.8; 12276000, 2.8; 12279600, 3; 12283200, 3.3; 12286800, 3.6; 12290400, 3.8; 12294000, 3.7; 12297600, 3.7; 12301200, 4.1; 12304800, 4.8; 12308400, 5.3; 12312000, 5.4; 12315600, 5.6; 12319200, 5.7; 12322800, 5.4; 12326400, 4.9; 12330000, 4.3; 12333600, 4.7; 12337200, 5.8; 12340800, 6.1; 12344400, 6.4; 12348000, 6.2; 12351600, 5.8; 12355200, 6.2; 12358800, 6.4; 12362400, 6.7; 12366000, 8; 12369600, 8.7; 12373200, 8.7; 12376800, 8.1; 12380400, 7.6; 12384000, 7.7; 12387600, 7.9; 12391200, 8.2; 12394800, 8.7; 12398400, 9.3; 12402000, 9.7; 12405600, 10.3; 12409200, 10.6; 12412800, 10.7; 12416400, 9.9; 12420000, 9.1; 12423600, 8.3; 12427200, 7.9; 12430800, 8; 12434400, 6.5; 12438000, 5.7; 12441600, 5.5; 12445200, 5.7; 12448800, 5.9; 12452400, 6.7; 12456000, 6.8; 12459600, 6.2; 12463200, 5.6; 12466800, 4.9; 12470400, 4.4; 12474000, 4; 12477600, 3.8; 12481200, 3.6; 12484800, 4.2; 12488400, 5.2; 12492000, 6.5; 12495600, 7.6; 12499200, 8.3; 12502800, 8.6; 12506400, 8.6; 12510000, 8.5; 12513600, 8.3; 12517200, 8; 12520800, 7.8; 12524400, 7.5; 12528000, 7; 12531600, 6; 12535200, 5.5; 12538800, 7.7; 12542400, 8.7; 12546000, 7.9; 12549600, 7.2; 12553200, 6.2; 12556800, 4.9; 12560400, 3.5; 12564000, 2.2; 12567600, 1.1; 12571200, 0.6; 12574800, 2; 12578400, 4.1; 12582000, 6; 12585600, 7.5; 12589200, 8.1; 12592800, 7.9; 12596400, 7.3; 12600000, 6.9; 12603600, 6.5; 12607200, 6.2; 12610800, 5.6; 12614400, 5; 12618000, 4.3; 12621600, 3.6; 12625200, 3; 12628800, 3.3; 12632400, 4.2; 12636000, 3.9; 12639600, 4.1; 12643200, 5.2; 12646800, 5.1; 12650400, 4.6; 12654000, 5.1; 12657600, 5.6; 12661200, 4.9; 12664800, 5.8; 12668400, 6.9; 12672000, 7.7; 12675600, 8.3; 12679200, 8.5; 12682800, 8.3; 12686400, 7.9; 12690000, 7.6; 12693600, 7.3; 12697200, 7.2; 12700800, 6.6; 12704400, 5.1; 12708000, 4.8; 12711600, 4.1; 12715200, 3.9; 12718800, 3.9; 12722400, 3.7; 12726000, 3.6; 12729600, 3.6; 12733200, 4; 12736800, 4.7; 12740400, 5.4; 12744000, 6.2; 12747600, 7.6; 12751200, 9; 12754800, 8.6; 12758400, 7.4; 12762000, 6.2; 12765600, 5.8; 12769200, 5.6; 12772800, 5.2; 12776400, 4.8; 12780000, 4.3; 12783600, 3.6; 12787200, 2.9; 12790800, 1.7; 12794400, 1.4; 12798000, 2.8; 12801600, 4.4; 12805200, 5.4; 12808800, 6.2; 12812400, 6.7; 12816000, 6.6; 12819600, 6.2; 12823200, 5.5; 12826800, 4.9; 12830400, 4.9; 12834000, 5.5; 12837600, 6.5; 12841200, 7.2; 12844800, 8; 12848400, 8.6; 12852000, 8.8; 12855600, 8.7; 12859200, 8.8; 12862800, 8.6; 12866400, 8.3; 12870000, 8.5; 12873600, 8.4; 12877200, 7.2; 12880800, 7; 12884400, 9.2; 12888000, 9.7; 12891600, 9.9; 12895200, 9.9; 12898800, 9.7; 12902400, 9; 12906000, 8.3; 12909600, 7.8; 12913200, 8.2; 12916800, 9.5; 12920400, 11.4; 12924000, 12.8; 12927600, 12.3; 12931200, 11; 12934800, 9.5; 12938400, 7.6; 12942000, 4.7; 12945600, 0.4; 12949200, 4.8; 12952800, 8.4; 12956400, 9.6; 12960000, 9.9; 12963600, 11.2; 12967200, 12.5; 12970800, 12.8; 12974400, 13.1; 12978000, 13.4; 12981600, 13.5; 12985200, 13.4; 12988800, 13; 12992400, 12.8; 12996000, 12; 12999600, 11.1; 13003200, 9.8; 13006800, 8.8; 13010400, 8.6; 13014000, 8.3; 13017600, 7.7; 13021200, 7.2; 13024800, 7.2; 13028400, 7.2; 13032000, 6.8; 13035600, 6.3; 13039200, 6; 13042800, 5.9; 13046400, 6.1; 13050000, 5.8; 13053600, 6.1; 13057200, 6.6; 13060800, 6.7; 13064400, 7.1; 13068000, 7.4; 13071600, 7.2; 13075200, 6.9; 13078800, 6.5; 13082400, 6.4; 13086000, 6; 13089600, 5.2; 13093200, 4.2; 13096800, 3.5; 13100400, 2.8; 13104000, 2.2; 13107600, 1.8; 13111200, 1.2; 13114800, 0.7; 13118400, 1.5; 13122000, 2.4; 13125600, 3.1; 13129200, 3.5; 13132800, 3.2; 13136400, 2.2; 13140000, 1.5; 13143600, 2; 13147200, 2.8; 13150800, 3.1; 13154400, 3; 13158000, 3.1; 13161600, 3.5; 13165200, 4; 13168800, 4.2; 13172400, 4.6; 13176000, 5.3; 13179600, 6.7; 13183200, 7.9; 13186800, 8.3; 13190400, 8.4; 13194000, 8.5; 13197600, 8.7; 13201200, 8.6; 13204800, 8.1; 13208400, 7.6; 13212000, 7.4; 13215600, 7.1; 13219200, 7; 13222800, 5.9; 13226400, 4.4; 13230000, 3.5; 13233600, 3.4; 13237200, 4; 13240800, 4.2; 13244400, 4; 13248000, 3.8; 13251600, 4; 13255200, 4.6; 13258800, 5.2; 13262400, 5.7; 13266000, 6.3; 13269600, 6.6; 13273200, 6.5; 13276800, 6.1; 13280400, 5.3; 13284000, 5.3; 13287600, 6.8; 13291200, 7.6; 13294800, 7.6; 13298400, 7.3; 13302000, 7.2; 13305600, 7; 13309200, 5.8; 13312800, 5.9; 13316400, 4.8; 13320000, 3.2; 13323600, 1.9; 13327200, 1.8; 13330800, 3.1; 13334400, 4.5; 13338000, 5.6; 13341600, 6.4; 13345200, 6.9; 13348800, 7.4; 13352400, 8; 13356000, 8.5; 13359600, 8.8; 13363200, 8.8; 13366800, 8.5; 13370400, 7.8; 13374000, 7.4; 13377600, 7.3; 13381200, 7.1; 13384800, 6.8; 13388400, 6.6; 13392000, 6.3; 13395600, 5.3; 13399200, 4.8; 13402800, 4.6; 13406400, 4.9; 13410000, 5.2; 13413600, 5.6; 13417200, 5.8; 13420800, 6.1; 13424400, 6.2; 13428000, 6.1; 13431600, 6; 13435200, 6.1; 13438800, 6.5; 13442400, 7.1; 13446000, 6.9; 13449600, 6.2; 13453200, 5; 13456800, 3.6; 13460400, 2; 13464000, 0.5; 13467600, 0.7; 13471200, 1.6; 13474800, 2.1; 13478400, 2.3; 13482000, 2.5; 13485600, 3.1; 13489200, 3; 13492800, 2.7; 13496400, 2.2; 13500000, 1.3; 13503600, 0.4; 13507200, 0.1; 13510800, 0.5; 13514400, 1.2; 13518000, 2.2; 13521600, 3.4; 13525200, 5.1; 13528800, 6.5; 13532400, 6.5; 13536000, 6; 13539600, 5.5; 13543200, 5; 13546800, 4.5; 13550400, 4; 13554000, 3.8; 13557600, 4.3; 13561200, 5.3; 13564800, 6.1; 13568400, 6.5; 13572000, 7; 13575600, 7; 13579200, 7; 13582800, 6.8; 13586400, 6.7; 13590000, 6.5; 13593600, 6.2; 13597200, 5.7; 13600800, 5.4; 13604400, 5; 13608000, 4.4; 13611600, 4; 13615200, 3.8; 13618800, 3.7; 13622400, 3.8; 13626000, 3.9; 13629600, 3.7; 13633200, 3.1; 13636800, 2.4; 13640400, 1.9; 13644000, 1.6; 13647600, 1.5; 13651200, 1.4; 13654800, 1.2; 13658400, 1.1; 13662000, 0.7; 13665600, 0.2; 13669200, 0.8; 13672800, 1.7; 13676400, 2.5; 13680000, 2.9; 13683600, 3.3; 13687200, 3.8; 13690800, 4.3; 13694400, 5; 13698000, 6.3; 13701600, 7.9; 13705200, 8.5; 13708800, 8.9; 13712400, 9; 13716000, 9; 13719600, 8.6; 13723200, 8.3; 13726800, 8.1; 13730400, 7.7; 13734000, 7.2; 13737600, 6.7; 13741200, 5.6; 13744800, 5; 13748400, 6.6; 13752000, 7.1; 13755600, 7.3; 13759200, 7.7; 13762800, 8.1; 13766400, 8.8; 13770000, 9.5; 13773600, 10.2; 13777200, 10.6; 13780800, 10.8; 13784400, 10.8; 13788000, 10.6; 13791600, 10.5; 13795200, 10.3; 13798800, 10.6; 13802400, 11.6; 13806000, 12; 13809600, 12.4; 13813200, 12.7; 13816800, 12.9; 13820400, 13.2; 13824000, 13.1; 13827600, 13.1; 13831200, 12.6; 13834800, 12.3; 13838400, 11.4; 13842000, 10.4; 13845600, 9.7; 13849200, 9.2; 13852800, 8.9; 13856400, 8.6; 13860000, 8.3; 13863600, 8; 13867200, 7.6; 13870800, 7.1; 13874400, 7.1; 13878000, 7.2; 13881600, 7.1; 13885200, 7; 13888800, 8.6; 13892400, 10; 13896000, 9.8; 13899600, 9.4; 13903200, 8.6; 13906800, 8.1; 13910400, 8; 13914000, 7.2; 13917600, 7.6; 13921200, 7.6; 13924800, 6.8; 13928400, 6.4; 13932000, 5.8; 13935600, 5.1; 13939200, 4.2; 13942800, 3.3; 13946400, 2.4; 13950000, 1.6; 13953600, 0.9; 13957200, 0.3; 13960800, 1.8; 13964400, 3.5; 13968000, 5; 13971600, 6.1; 13975200, 6.9; 13978800, 6.9; 13982400, 6.6; 13986000, 6; 13989600, 5.4; 13993200, 4.8; 13996800, 4.5; 14000400, 3.8; 14004000, 3.9; 14007600, 3.4; 14011200, 3.2; 14014800, 3.4; 14018400, 3.8; 14022000, 4.2; 14025600, 4.7; 14029200, 4.9; 14032800, 5.1; 14036400, 5.4; 14040000, 6; 14043600, 7; 14047200, 7.9; 14050800, 8.5; 14054400, 9.1; 14058000, 9.2; 14061600, 8.7; 14065200, 8.1; 14068800, 7.7; 14072400, 7.3; 14076000, 6.7; 14079600, 5.8; 14083200, 5.1; 14086800, 3.7; 14090400, 2.5; 14094000, 2.2; 14097600, 2.3; 14101200, 2.8; 14104800, 3; 14108400, 3.1; 14112000, 3.5; 14115600, 4.7; 14119200, 6; 14122800, 7; 14126400, 7.6; 14130000, 8; 14133600, 7.9; 14137200, 8.2; 14140800, 8.2; 14144400, 7.4; 14148000, 7; 14151600, 6.6; 14155200, 6.6; 14158800, 6.4; 14162400, 5.1; 14166000, 3.9; 14169600, 3.8; 14173200, 4.2; 14176800, 4.8; 14180400, 5.1; 14184000, 5; 14187600, 4.5; 14191200, 3.7; 14194800, 3.8; 14198400, 5.2; 14202000, 6.6; 14205600, 7.7; 14209200, 8.4; 14212800, 8.6; 14216400, 8.8; 14220000, 8.7; 14223600, 8.4; 14227200, 8.4; 14230800, 8.6; 14234400, 8.4; 14238000, 8.1; 14241600, 8.1; 14245200, 8; 14248800, 7.7; 14252400, 7.5; 14256000, 7.3; 14259600, 6.8; 14263200, 6.9; 14266800, 6.7; 14270400, 6.7; 14274000, 6.3; 14277600, 5.5; 14281200, 4.6; 14284800, 3.8; 14288400, 3.3; 14292000, 3.1; 14295600, 3.1; 14299200, 3.1; 14302800, 4.3; 14306400, 5.8; 14310000, 6.4; 14313600, 6.9; 14317200, 7.7; 14320800, 8.4; 14324400, 8.8; 14328000, 8.8; 14331600, 8.6; 14335200, 8.3; 14338800, 7.9; 14342400, 7.3; 14346000, 5.8; 14349600, 4.8; 14353200, 5.3; 14356800, 5.3; 14360400, 5.2; 14364000, 5; 14367600, 4.6; 14371200, 4.5; 14374800, 4.7; 14378400, 5.4; 14382000, 6.3; 14385600, 7.2; 14389200, 8.2; 14392800, 10; 14396400, 11.2; 14400000, 11.9; 14403600, 12.4; 14407200, 12.8; 14410800, 12.7; 14414400, 12.7; 14418000, 12; 14421600, 11.5; 14425200, 11.5; 14428800, 11.2; 14432400, 10.1; 14436000, 11.1; 14439600, 10.3; 14443200, 9.5; 14446800, 8.5; 14450400, 7.8; 14454000, 7.5; 14457600, 7.3; 14461200, 7; 14464800, 6.5; 14468400, 5.4; 14472000, 4.3; 14475600, 4.6; 14479200, 5.4; 14482800, 5.8; 14486400, 2.2; 14490000, 3.4; 14493600, 7; 14497200, 8.1; 14500800, 8.7; 14504400, 9.1; 14508000, 8.2; 14511600, 8; 14515200, 8.1; 14518800, 7.2; 14522400, 6.3; 14526000, 5.9; 14529600, 5.3; 14533200, 4.6; 14536800, 3.9; 14540400, 3.6; 14544000, 3.3; 14547600, 3.3; 14551200, 3.5; 14554800, 3.9; 14558400, 4.8; 14562000, 6; 14565600, 6.6; 14569200, 6.7; 14572800, 6.8; 14576400, 6.5; 14580000, 6.3; 14583600, 6.1; 14587200, 5.9; 14590800, 5.6; 14594400, 5.4; 14598000, 5; 14601600, 4.4; 14605200, 3.3; 14608800, 1.9; 14612400, 1.3; 14616000, 1.4; 14619600, 0.9; 14623200, 0.7; 14626800, 1.2; 14630400, 2; 14634000, 3; 14637600, 3.8; 14641200, 4.7; 14644800, 6.2; 14648400, 7.3; 14652000, 7.9; 14655600, 8.5; 14659200, 8; 14662800, 7.2; 14666400, 7.1; 14670000, 6.7; 14673600, 6; 14677200, 5.6; 14680800, 5.5; 14684400, 5.5; 14688000, 5.5; 14691600, 4.8; 14695200, 4.4; 14698800, 4.4; 14702400, 3.8; 14706000, 3; 14709600, 2.5; 14713200, 2.6; 14716800, 3; 14720400, 3.7; 14724000, 4.3; 14727600, 4.6; 14731200, 4.8; 14734800, 4.8; 14738400, 5.1; 14742000, 5.6; 14745600, 6; 14749200, 6.3; 14752800, 6.2; 14756400, 6; 14760000, 5.3; 14763600, 4.6; 14767200, 4.1; 14770800, 3.8; 14774400, 3.7; 14778000, 3.4; 14781600, 2.9; 14785200, 2.4; 14788800, 2.1; 14792400, 1.8; 14796000, 1.8; 14799600, 2; 14803200, 2.1; 14806800, 2.3; 14810400, 2.2; 14814000, 2.4; 14817600, 2.9; 14821200, 4; 14824800, 5; 14828400, 5.2; 14832000, 4.5; 14835600, 4; 14839200, 4.4; 14842800, 5; 14846400, 5.2; 14850000, 5; 14853600, 4.3; 14857200, 3.4; 14860800, 2.3; 14864400, 1.1; 14868000, 0.5; 14871600, 0.6; 14875200, 1.2; 14878800, 1.5; 14882400, 1.8; 14886000, 2.1; 14889600, 2.7; 14893200, 3.4; 14896800, 4.2; 14900400, 4.7; 14904000, 4.7; 14907600, 5; 14911200, 5.6; 14914800, 6.4; 14918400, 6.9; 14922000, 7.5; 14925600, 7.7; 14929200, 7.9; 14932800, 7.7; 14936400, 7; 14940000, 6.4; 14943600, 6.1; 14947200, 5.6; 14950800, 4.4; 14954400, 3.6; 14958000, 6.2; 14961600, 7.2; 14965200, 6.6; 14968800, 7.6; 14972400, 7.9; 14976000, 6.2; 14979600, 4.7; 14983200, 4.2; 14986800, 3.5; 14990400, 2.2; 14994000, 1.1; 14997600, 2.1; 15001200, 2.3; 15004800, 1.7; 15008400, 2.1; 15012000, 3.3; 15015600, 4.9; 15019200, 6; 15022800, 6.6; 15026400, 7.1; 15030000, 7.4; 15033600, 6.6; 15037200, 5.1; 15040800, 5.2; 15044400, 5.4; 15048000, 5.5; 15051600, 5.1; 15055200, 3.9; 15058800, 2.4; 15062400, 0.8; 15066000, 1.1; 15069600, 2.7; 15073200, 4.6; 15076800, 6.8; 15080400, 9.3; 15084000, 10.7; 15087600, 10.7; 15091200, 9.6; 15094800, 8.9; 15098400, 7.9; 15102000, 6.8; 15105600, 6.9; 15109200, 7.4; 15112800, 7.6; 15116400, 7.2; 15120000, 6.8; 15123600, 5.8; 15127200, 4.7; 15130800, 5.2; 15134400, 5.8; 15138000, 5.6; 15141600, 4.6; 15145200, 3.6; 15148800, 2.7; 15152400, 1.9; 15156000, 1.1; 15159600, 0.5; 15163200, 1.6; 15166800, 3.6; 15170400, 5.9; 15174000, 7.6; 15177600, 8.8; 15181200, 9.5; 15184800, 9.4; 15188400, 8.9; 15192000, 8.4; 15195600, 8.1; 15199200, 7.9; 15202800, 8.1; 15206400, 8.1; 15210000, 7; 15213600, 5.8; 15217200, 5.7; 15220800, 6.7; 15224400, 7.1; 15228000, 7; 15231600, 6.9; 15235200, 6.7; 15238800, 5.8; 15242400, 4.6; 15246000, 3.9; 15249600, 3.5; 15253200, 3.2; 15256800, 3.1; 15260400, 3.3; 15264000, 3.9; 15267600, 4.7; 15271200, 5.7; 15274800, 6.8; 15278400, 8.3; 15282000, 9; 15285600, 9.2; 15289200, 9.2; 15292800, 8.7; 15296400, 7.8; 15300000, 8.6; 15303600, 9.3; 15307200, 9.4; 15310800, 8.9; 15314400, 8.6; 15318000, 8.6; 15321600, 8.5; 15325200, 8.3; 15328800, 7.9; 15332400, 7.4; 15336000, 6.3; 15339600, 5.8; 15343200, 6.3; 15346800, 6.7; 15350400, 7; 15354000, 7.4; 15357600, 7.7; 15361200, 7.9; 15364800, 8.1; 15368400, 8.1; 15372000, 8.3; 15375600, 8.4; 15379200, 8.1; 15382800, 7.1; 15386400, 7.2; 15390000, 7; 15393600, 7.2; 15397200, 7.4; 15400800, 7.1; 15404400, 6.8; 15408000, 6.6; 15411600, 6.2; 15415200, 5.3; 15418800, 4; 15422400, 2.4; 15426000, 1.7; 15429600, 3.3; 15433200, 4.8; 15436800, 5.1; 15440400, 4.8; 15444000, 4.2; 15447600, 3.9; 15451200, 4; 15454800, 4.5; 15458400, 5.2; 15462000, 5.7; 15465600, 6; 15469200, 5.4; 15472800, 5.3; 15476400, 5.9; 15480000, 5.6; 15483600, 5; 15487200, 4.7; 15490800, 4.4; 15494400, 4; 15498000, 3.6; 15501600, 3; 15505200, 2.3; 15508800, 1.3; 15512400, 0.2; 15516000, 2.5; 15519600, 4.9; 15523200, 6.5; 15526800, 7.3; 15530400, 7.5; 15534000, 7.3; 15537600, 7; 15541200, 6.6; 15544800, 5.9; 15548400, 5.5; 15552000, 3.6; 15555600, 0.7; 15559200, 1.8; 15562800, 2.5; 15566400, 3.3; 15570000, 4.4; 15573600, 4.3; 15577200, 3.7; 15580800, 3.2; 15584400, 2.8; 15588000, 2.5; 15591600, 2.6; 15595200, 2.9; 15598800, 3.1; 15602400, 3.4; 15606000, 3.3; 15609600, 2.7; 15613200, 2.7; 15616800, 3.3; 15620400, 3.4; 15624000, 3.7; 15627600, 4.8; 15631200, 5.8; 15634800, 6.5; 15638400, 6.8; 15642000, 7.4; 15645600, 7.9; 15649200, 8.3; 15652800, 8.4; 15656400, 8.4; 15660000, 8.3; 15663600, 8.1; 15667200, 7.7; 15670800, 7.1; 15674400, 6.3; 15678000, 5.4; 15681600, 4.7; 15685200, 4.2; 15688800, 3.8; 15692400, 3.8; 15696000, 3.7; 15699600, 3.4; 15703200, 3.1; 15706800, 2.7; 15710400, 2.4; 15714000, 2.5; 15717600, 3; 15721200, 3.3; 15724800, 3.1; 15728400, 2.8; 15732000, 2.4; 15735600, 2.4; 15739200, 2.3; 15742800, 2.2; 15746400, 2; 15750000, 1.9; 15753600, 2.1; 15757200, 2.6; 15760800, 3; 15764400, 3.3; 15768000, 3.8; 15771600, 4.6; 15775200, 5.4; 15778800, 5.9; 15782400, 6.5; 15786000, 7.1; 15789600, 7.5; 15793200, 7.8; 15796800, 7.8; 15800400, 7.4; 15804000, 6.7; 15807600, 6; 15811200, 5.3; 15814800, 4; 15818400, 3; 15822000, 3.8; 15825600, 4.2; 15829200, 4.6; 15832800, 5; 15836400, 5.4; 15840000, 5.7; 15843600, 5.8; 15847200, 5.7; 15850800, 5.5; 15854400, 5.3; 15858000, 5.2; 15861600, 5.5; 15865200, 6; 15868800, 6.4; 15872400, 6.5; 15876000, 6.3; 15879600, 6; 15883200, 5.8; 15886800, 5.2; 15890400, 4.6; 15894000, 4.3; 15897600, 4; 15901200, 3.4; 15904800, 3.1; 15908400, 3.4; 15912000, 3.8; 15915600, 4.5; 15919200, 5.3; 15922800, 5.9; 15926400, 6.4; 15930000, 6.4; 15933600, 6.4; 15937200, 6.7; 15940800, 7; 15944400, 7.4; 15948000, 7.5; 15951600, 7.3; 15955200, 6.9; 15958800, 6.6; 15962400, 6.1; 15966000, 5.6; 15969600, 5.1; 15973200, 4.7; 15976800, 4.6; 15980400, 4.6; 15984000, 4.7; 15987600, 4.4; 15991200, 4.4; 15994800, 4.5; 15998400, 4.9; 16002000, 5.3; 16005600, 5.6; 16009200, 5.8; 16012800, 5.5; 16016400, 5; 16020000, 4.6; 16023600, 4.4; 16027200, 4.4; 16030800, 4.9; 16034400, 5.7; 16038000, 5.8; 16041600, 5.6; 16045200, 5.4; 16048800, 5.4; 16052400, 5.3; 16056000, 4.9; 16059600, 4.1; 16063200, 3.5; 16066800, 3.1; 16070400, 2.8; 16074000, 2.4; 16077600, 1.6; 16081200, 0.5; 16084800, 0.8; 16088400, 1.5; 16092000, 2; 16095600, 2.4; 16099200, 3.1; 16102800, 4; 16106400, 4.6; 16110000, 4.7; 16113600, 4.5; 16117200, 5; 16120800, 6.3; 16124400, 7.1; 16128000, 7.5; 16131600, 7.5; 16135200, 7.5; 16138800, 7.5; 16142400, 7.4; 16146000, 7.2; 16149600, 7; 16153200, 6.6; 16156800, 5.9; 16160400, 4.4; 16164000, 3.9; 16167600, 3.2; 16171200, 3.1; 16174800, 3.6; 16178400, 4.3; 16182000, 5; 16185600, 5.5; 16189200, 5.9; 16192800, 6.2; 16196400, 6.6; 16200000, 7.3; 16203600, 8.2; 16207200, 8.7; 16210800, 8.5; 16214400, 8.1; 16218000, 7.6; 16221600, 7.2; 16225200, 7; 16228800, 7; 16232400, 7; 16236000, 6.5; 16239600, 5.7; 16243200, 4.7; 16246800, 3.6; 16250400, 3.1; 16254000, 3.1; 16257600, 3.7; 16261200, 4.4; 16264800, 5; 16268400, 5.5; 16272000, 5.8; 16275600, 6.2; 16279200, 6.5; 16282800, 6.8; 16286400, 6.8; 16290000, 7.1; 16293600, 7.5; 16297200, 7.4; 16300800, 7.1; 16304400, 6.8; 16308000, 6.5; 16311600, 6.2; 16315200, 5.8; 16318800, 5.5; 16322400, 5.5; 16326000, 5.3; 16329600, 4.8; 16333200, 3.6; 16336800, 2.9; 16340400, 2.1; 16344000, 1.8; 16347600, 2.6; 16351200, 3.2; 16354800, 3.4; 16358400, 3.4; 16362000, 3.4; 16365600, 3.5; 16369200, 3.3; 16372800, 3; 16376400, 2.8; 16380000, 3.2; 16383600, 4.2; 16387200, 5.2; 16390800, 6.1; 16394400, 6.3; 16398000, 5.9; 16401600, 5.3; 16405200, 4.7; 16408800, 4.2; 16412400, 3.9; 16416000, 3.6; 16419600, 3.5; 16423200, 3.5; 16426800, 3.2; 16430400, 3; 16434000, 2.9; 16437600, 2.9; 16441200, 3.2; 16444800, 3.6; 16448400, 4; 16452000, 4.2; 16455600, 4.5; 16459200, 5; 16462800, 5.7; 16466400, 6.4; 16470000, 6.5; 16473600, 6.4; 16477200, 6.3; 16480800, 6.4; 16484400, 6.4; 16488000, 5.9; 16491600, 5.1; 16495200, 4.4; 16498800, 4.2; 16502400, 4; 16506000, 3.3; 16509600, 3; 16513200, 2.8; 16516800, 2.8; 16520400, 2.6; 16524000, 2.3; 16527600, 2.3; 16531200, 2.5; 16534800, 3; 16538400, 3.6; 16542000, 3.8; 16545600, 3.7; 16549200, 3.6; 16552800, 3.6; 16556400, 3.7; 16560000, 4.1; 16563600, 4.4; 16567200, 4.8; 16570800, 5; 16574400, 5.5; 16578000, 6; 16581600, 6.4; 16585200, 6.4; 16588800, 5.8; 16592400, 4.5; 16596000, 4.4; 16599600, 3.1; 16603200, 2.8; 16606800, 3.2; 16610400, 3.5; 16614000, 3.3; 16617600, 3; 16621200, 3.3; 16624800, 4.3; 16628400, 4.7; 16632000, 4.2; 16635600, 3.4; 16639200, 2.7; 16642800, 1.8; 16646400, 1.3; 16650000, 1.6; 16653600, 2.2; 16657200, 2.7; 16660800, 3.3; 16664400, 3.9; 16668000, 4.3; 16671600, 4.6; 16675200, 4.7; 16678800, 4.4; 16682400, 4.3; 16686000, 4.4; 16689600, 4.6; 16693200, 4.3; 16696800, 4.1; 16700400, 4.1; 16704000, 4.4; 16707600, 5; 16711200, 5.3; 16714800, 5.3; 16718400, 5.1; 16722000, 5; 16725600, 5.2; 16729200, 5.1; 16732800, 4.8; 16736400, 4.4; 16740000, 3.8; 16743600, 3.2; 16747200, 2.9; 16750800, 2.7; 16754400, 2.3; 16758000, 1.8; 16761600, 2.1; 16765200, 3.8; 16768800, 4.5; 16772400, 5.3; 16776000, 6; 16779600, 7; 16783200, 7.3; 16786800, 7; 16790400, 6.7; 16794000, 6.2; 16797600, 6; 16801200, 5.8; 16804800, 5.2; 16808400, 4.7; 16812000, 4.4; 16815600, 4; 16819200, 3.3; 16822800, 2.4; 16826400, 1.5; 16830000, 1.1; 16833600, 1.6; 16837200, 3.1; 16840800, 4.7; 16844400, 6.1; 16848000, 6.9; 16851600, 6.7; 16855200, 6; 16858800, 5.5; 16862400, 5.2; 16866000, 5.6; 16869600, 6; 16873200, 6.2; 16876800, 6.2; 16880400, 6.1; 16884000, 6; 16887600, 5.7; 16891200, 5.3; 16894800, 4.7; 16898400, 4.3; 16902000, 3.7; 16905600, 3; 16909200, 2.3; 16912800, 1.9; 16916400, 1.9; 16920000, 1.7; 16923600, 2; 16927200, 3.2; 16930800, 4.4; 16934400, 4.8; 16938000, 4.3; 16941600, 3.8; 16945200, 3.4; 16948800, 3.1; 16952400, 2.9; 16956000, 3.2; 16959600, 3.6; 16963200, 4.2; 16966800, 4.6; 16970400, 4.7; 16974000, 4.5; 16977600, 4.8; 16981200, 5.6; 16984800, 6.2; 16988400, 6; 16992000, 5.4; 16995600, 4.9; 16999200, 4.6; 17002800, 4.5; 17006400, 4.5; 17010000, 4.3; 17013600, 4.1; 17017200, 3.7; 17020800, 3.3; 17024400, 2.7; 17028000, 2.6; 17031600, 2.7; 17035200, 3; 17038800, 3.4; 17042400, 4; 17046000, 4.4; 17049600, 4.6; 17053200, 4.4; 17056800, 4.4; 17060400, 4.3; 17064000, 4.7; 17067600, 5.6; 17071200, 6.7; 17074800, 7.2; 17078400, 6.9; 17082000, 6.5; 17085600, 5.9; 17089200, 5.5; 17092800, 5.6; 17096400, 5.3; 17100000, 4.5; 17103600, 3.9; 17107200, 3.5; 17110800, 3.4; 17114400, 4.2; 17118000, 5.2; 17121600, 5.5; 17125200, 4.7; 17128800, 5.4; 17132400, 6.5; 17136000, 6.1; 17139600, 5.9; 17143200, 6; 17146800, 6.2; 17150400, 6.2; 17154000, 6; 17157600, 6.2; 17161200, 6.5; 17164800, 6.7; 17168400, 6.9; 17172000, 6.8; 17175600, 6.8; 17179200, 6.8; 17182800, 6.8; 17186400, 6.9; 17190000, 7.1; 17193600, 7; 17197200, 6.7; 17200800, 7.3; 17204400, 7.5; 17208000, 7.4; 17211600, 7.1; 17215200, 6.6; 17218800, 6.6; 17222400, 7.6; 17226000, 8.9; 17229600, 10.3; 17233200, 11; 17236800, 11.1; 17240400, 11.2; 17244000, 11.3; 17247600, 11.2; 17251200, 10.8; 17254800, 10.2; 17258400, 9.5; 17262000, 10.6; 17265600, 12.1; 17269200, 13.2; 17272800, 13.8; 17276400, 14.5; 17280000, 15.6; 17283600, 16.8; 17287200, 19; 17290800, 21.2; 17294400, 21.1; 17298000, 19.8; 17301600, 18.6; 17305200, 17.2; 17308800, 15.4; 17312400, 14.1; 17316000, 13.1; 17319600, 12.5; 17323200, 12; 17326800, 12.7; 17330400, 13.7; 17334000, 13.6; 17337600, 12; 17341200, 10.6; 17344800, 10.2; 17348400, 10.4; 17352000, 10; 17355600, 9.1; 17359200, 8.4; 17362800, 7.4; 17366400, 6; 17370000, 4.6; 17373600, 4.2; 17377200, 3.9; 17380800, 4.2; 17384400, 4.8; 17388000, 5.5; 17391600, 5.9; 17395200, 6.3; 17398800, 6.5; 17402400, 6.1; 17406000, 5.4; 17409600, 4.7; 17413200, 4.7; 17416800, 5; 17420400, 5.1; 17424000, 5.2; 17427600, 5.2; 17431200, 4.9; 17434800, 4.4; 17438400, 3.9; 17442000, 3.4; 17445600, 3; 17449200, 3.1; 17452800, 3.5; 17456400, 3.6; 17460000, 4; 17463600, 4.4; 17467200, 4.6; 17470800, 4.7; 17474400, 4.4; 17478000, 3.8; 17481600, 3.2; 17485200, 2.8; 17488800, 2.6; 17492400, 2.6; 17496000, 2.3; 17499600, 2.1; 17503200, 3.1; 17506800, 4.4; 17510400, 5.2; 17514000, 5.5; 17517600, 5.3; 17521200, 4.7; 17524800, 4.2; 17528400, 3.9; 17532000, 3.4; 17535600, 3.2; 17539200, 2.7; 17542800, 2.1; 17546400, 1.8; 17550000, 2.1; 17553600, 2.6; 17557200, 3; 17560800, 3.3; 17564400, 3.6; 17568000, 3.9; 17571600, 4.1; 17575200, 4.3; 17578800, 4.4; 17582400, 4.6; 17586000, 5; 17589600, 5.7; 17593200, 6.2; 17596800, 6.3; 17600400, 6.2; 17604000, 6; 17607600, 5.8; 17611200, 5.7; 17614800, 5.5; 17618400, 5.3; 17622000, 5; 17625600, 4.5; 17629200, 3.6; 17632800, 2.5; 17636400, 2.3; 17640000, 2.7; 17643600, 3.3; 17647200, 3.8; 17650800, 4.2; 17654400, 4.5; 17658000, 4.6; 17661600, 4.3; 17665200, 4; 17668800, 4; 17672400, 4.5; 17676000, 5.4; 17679600, 6.6; 17683200, 7.7; 17686800, 8.1; 17690400, 8; 17694000, 7.7; 17697600, 7.6; 17701200, 7.8; 17704800, 8.3; 17708400, 8.6; 17712000, 8.3; 17715600, 6.9; 17719200, 6.3; 17722800, 6.6; 17726400, 7.6; 17730000, 8.2; 17733600, 8.8; 17737200, 8.8; 17740800, 8.4; 17744400, 7.7; 17748000, 6.8; 17751600, 5.6; 17755200, 4.5; 17758800, 4.1; 17762400, 4; 17766000, 3.9; 17769600, 3.7; 17773200, 3.6; 17776800, 3.7; 17780400, 4.2; 17784000, 4.8; 17787600, 5.4; 17791200, 6.4; 17794800, 7.5; 17798400, 8.3; 17802000, 8.5; 17805600, 8.2; 17809200, 7.7; 17812800, 7; 17816400, 6; 17820000, 5.2; 17823600, 4.5; 17827200, 3.9; 17830800, 3.3; 17834400, 2.6; 17838000, 1.9; 17841600, 1.6; 17845200, 1.6; 17848800, 2.3; 17852400, 3.4; 17856000, 4.3; 17859600, 4.7; 17863200, 4.7; 17866800, 4.2; 17870400, 3.8; 17874000, 3.8; 17877600, 4.1; 17881200, 4.2; 17884800, 3.9; 17888400, 3.3; 17892000, 2.8; 17895600, 2.7; 17899200, 2.4; 17902800, 2.5; 17906400, 3.1; 17910000, 3.5; 17913600, 3.5; 17917200, 3.4; 17920800, 4; 17924400, 4.8; 17928000, 4.7; 17931600, 4.2; 17935200, 3.9; 17938800, 3.9; 17942400, 4.9; 17946000, 5.6; 17949600, 6.2; 17953200, 6.2; 17956800, 5.7; 17960400, 4.8; 17964000, 4.2; 17967600, 4.5; 17971200, 5.8; 17974800, 6.1; 17978400, 6.4; 17982000, 6.4; 17985600, 5.7; 17989200, 5; 17992800, 4.4; 17996400, 3.9; 18000000, 3.7; 18003600, 3.7; 18007200, 3.9; 18010800, 4.2; 18014400, 4.3; 18018000, 4.2; 18021600, 4.1; 18025200, 4.3; 18028800, 4.8; 18032400, 5.3; 18036000, 5.9; 18039600, 6.2; 18043200, 6.1; 18046800, 5.9; 18050400, 5.5; 18054000, 5.1; 18057600, 4.6; 18061200, 3.9; 18064800, 3.1; 18068400, 3.3; 18072000, 3.7; 18075600, 4.3; 18079200, 4.7; 18082800, 4.9; 18086400, 5; 18090000, 5.1; 18093600, 5.4; 18097200, 5.9; 18100800, 6.4; 18104400, 6.8; 18108000, 7.2; 18111600, 7.4; 18115200, 7.3; 18118800, 7.1; 18122400, 7; 18126000, 6.7; 18129600, 6.4; 18133200, 6; 18136800, 5.6; 18140400, 5.3; 18144000, 5.3; 18147600, 4.7; 18151200, 4; 18154800, 3.5; 18158400, 3.4; 18162000, 3.5; 18165600, 3.6; 18169200, 3.6; 18172800, 3.6; 18176400, 3.6; 18180000, 3.7; 18183600, 3.9; 18187200, 3.8; 18190800, 4; 18194400, 4.7; 18198000, 5.4; 18201600, 5.7; 18205200, 5.7; 18208800, 5.6; 18212400, 5.1; 18216000, 4.3; 18219600, 4.2; 18223200, 4.7; 18226800, 5.2; 18230400, 4.8; 18234000, 4.1; 18237600, 3.7; 18241200, 3.5; 18244800, 3.3; 18248400, 3.3; 18252000, 3.3; 18255600, 3.2; 18259200, 2.9; 18262800, 2.5; 18266400, 2.3; 18270000, 2.4; 18273600, 3; 18277200, 4.1; 18280800, 5.3; 18284400, 6; 18288000, 6; 18291600, 5.9; 18295200, 5.6; 18298800, 5.1; 18302400, 4.8; 18306000, 4.6; 18309600, 4.6; 18313200, 4.5; 18316800, 4.7; 18320400, 4.4; 18324000, 4; 18327600, 3.9; 18331200, 3.9; 18334800, 4; 18338400, 4.4; 18342000, 5.3; 18345600, 6; 18349200, 5.9; 18352800, 5.9; 18356400, 7.3; 18360000, 7.3; 18363600, 6.7; 18367200, 6.6; 18370800, 5.8; 18374400, 5.2; 18378000, 4.6; 18381600, 3.8; 18385200, 3.4; 18388800, 3.6; 18392400, 4.7; 18396000, 5.8; 18399600, 6; 18403200, 5.8; 18406800, 5.3; 18410400, 5; 18414000, 5.1; 18417600, 5.2; 18421200, 5.4; 18424800, 5.7; 18428400, 6.1; 18432000, 6.4; 18435600, 6.4; 18439200, 6.4; 18442800, 6.3; 18446400, 5.9; 18450000, 5.4; 18453600, 5.3; 18457200, 5; 18460800, 4.6; 18464400, 4.3; 18468000, 4.2; 18471600, 4; 18475200, 3.7; 18478800, 3.4; 18482400, 3.2; 18486000, 3.5; 18489600, 4.3; 18493200, 4.6; 18496800, 4.6; 18500400, 4.6; 18504000, 4.6; 18507600, 4.7; 18511200, 4.8; 18514800, 4.7; 18518400, 4.8; 18522000, 4.9; 18525600, 4.9; 18529200, 4.8; 18532800, 4.8; 18536400, 5; 18540000, 5.3; 18543600, 5.3; 18547200, 5; 18550800, 4.7; 18554400, 4.4; 18558000, 4.2; 18561600, 3.9; 18565200, 3.5; 18568800, 2.9; 18572400, 2.4; 18576000, 1.9; 18579600, 2.1; 18583200, 2.5; 18586800, 3; 18590400, 3.3; 18594000, 3.7; 18597600, 3.9; 18601200, 4; 18604800, 4.1; 18608400, 4.1; 18612000, 4.1; 18615600, 4.2; 18619200, 4.7; 18622800, 5.3; 18626400, 5.8; 18630000, 5.9; 18633600, 5.9; 18637200, 5.7; 18640800, 5.4; 18644400, 5.1; 18648000, 4.8; 18651600, 4.5; 18655200, 4.2; 18658800, 3.8; 18662400, 3.3; 18666000, 2.8; 18669600, 2.7; 18673200, 3.2; 18676800, 3.5; 18680400, 3.4; 18684000, 3.2; 18687600, 3.1; 18691200, 3; 18694800, 3; 18698400, 3.2; 18702000, 3.6; 18705600, 4.1; 18709200, 4.9; 18712800, 6.1; 18716400, 6.8; 18720000, 6.9; 18723600, 6.7; 18727200, 6.7; 18730800, 6.8; 18734400, 6.8; 18738000, 6.8; 18741600, 6.8; 18745200, 6.8; 18748800, 6.7; 18752400, 6; 18756000, 5.5; 18759600, 5.7; 18763200, 5.6; 18766800, 5.8; 18770400, 6.1; 18774000, 6.4; 18777600, 6.4; 18781200, 6.3; 18784800, 6.1; 18788400, 6.2; 18792000, 6.3; 18795600, 6.4; 18799200, 6.6; 18802800, 6.4; 18806400, 6.1; 18810000, 5.8; 18813600, 5.7; 18817200, 5.1; 18820800, 3.7; 18824400, 2.2; 18828000, 2.1; 18831600, 3.2; 18835200, 4.5; 18838800, 5.4; 18842400, 6.2; 18846000, 7.1; 18849600, 7.7; 18853200, 8; 18856800, 8.3; 18860400, 8.9; 18864000, 9.7; 18867600, 10.4; 18871200, 10.7; 18874800, 10.6; 18878400, 9.9; 18882000, 9.2; 18885600, 8.9; 18889200, 8.7; 18892800, 8.4; 18896400, 8; 18900000, 7.8; 18903600, 7.7; 18907200, 7.8; 18910800, 7.9; 18914400, 8.1; 18918000, 8.4; 18921600, 8.5; 18925200, 8.6; 18928800, 9.2; 18932400, 9.8; 18936000, 9.9; 18939600, 9.8; 18943200, 9.6; 18946800, 9.2; 18950400, 9; 18954000, 8.9; 18957600, 8.8; 18961200, 8.7; 18964800, 8.6; 18968400, 8.3; 18972000, 8.3; 18975600, 8.2; 18979200, 8; 18982800, 7.6; 18986400, 7; 18990000, 6.5; 18993600, 6.1; 18997200, 5.6; 19000800, 4.6; 19004400, 4.1; 19008000, 4.3; 19011600, 4.4; 19015200, 4.2; 19018800, 4; 19022400, 3.6; 19026000, 3; 19029600, 2.4; 19033200, 1.9; 19036800, 1.5; 19040400, 1.4; 19044000, 1.5; 19047600, 1.9; 19051200, 2.6; 19054800, 3.7; 19058400, 4.7; 19062000, 4.9; 19065600, 4.6; 19069200, 4.5; 19072800, 4.7; 19076400, 4.6; 19080000, 4.3; 19083600, 4.1; 19087200, 4.1; 19090800, 4.3; 19094400, 4.4; 19098000, 4.1; 19101600, 4.1; 19105200, 4.5; 19108800, 4.7; 19112400, 4.6; 19116000, 4.5; 19119600, 4.5; 19123200, 4.7; 19126800, 4.8; 19130400, 4.9; 19134000, 5; 19137600, 5.2; 19141200, 5.9; 19144800, 6.6; 19148400, 6.9; 19152000, 6.7; 19155600, 6.3; 19159200, 6; 19162800, 6; 19166400, 6; 19170000, 6.1; 19173600, 6.3; 19177200, 6.5; 19180800, 6.6; 19184400, 5.9; 19188000, 5.2; 19191600, 5.6; 19195200, 5.5; 19198800, 4.9; 19202400, 4.1; 19206000, 3.4; 19209600, 3.1; 19213200, 3.1; 19216800, 3.3; 19220400, 3.5; 19224000, 3.9; 19227600, 4.5; 19231200, 5.3; 19234800, 5.5; 19238400, 5.5; 19242000, 5.3; 19245600, 5.4; 19249200, 5.6; 19252800, 5.5; 19256400, 5.1; 19260000, 4.9; 19263600, 4.7; 19267200, 4.5; 19270800, 4.1; 19274400, 3.7; 19278000, 3.9; 19281600, 4.4; 19285200, 4.8; 19288800, 4.6; 19292400, 4.1; 19296000, 4; 19299600, 4.1; 19303200, 4.2; 19306800, 4.3; 19310400, 4.6; 19314000, 5.3; 19317600, 6; 19321200, 6.3; 19324800, 5.8; 19328400, 4.4; 19332000, 3.4; 19335600, 3.6; 19339200, 3.4; 19342800, 3.8; 19346400, 5.1; 19350000, 6; 19353600, 6.3; 19357200, 5.8; 19360800, 5.4; 19364400, 4.6; 19368000, 4; 19371600, 3.6; 19375200, 3.6; 19378800, 3.9; 19382400, 4; 19386000, 3.6; 19389600, 2.8; 19393200, 2.1; 19396800, 2.4; 19400400, 2.8; 19404000, 2.9; 19407600, 2.6; 19411200, 2.6; 19414800, 2.2; 19418400, 2; 19422000, 3.1; 19425600, 4.7; 19429200, 4.9; 19432800, 4; 19436400, 3.2; 19440000, 2.9; 19443600, 2.7; 19447200, 2.3; 19450800, 2.2; 19454400, 2.2; 19458000, 2.3; 19461600, 2.2; 19465200, 2.3; 19468800, 2.7; 19472400, 3; 19476000, 3.6; 19479600, 4.6; 19483200, 5.3; 19486800, 5.9; 19490400, 6; 19494000, 5.8; 19497600, 4.9; 19501200, 3.9; 19504800, 3.4; 19508400, 2.6; 19512000, 2; 19515600, 3.8; 19519200, 6; 19522800, 7.3; 19526400, 8.2; 19530000, 8.7; 19533600, 9.1; 19537200, 9.6; 19540800, 9.2; 19544400, 8.6; 19548000, 8.7; 19551600, 7.9; 19555200, 6.6; 19558800, 5.4; 19562400, 4.2; 19566000, 3.2; 19569600, 2.6; 19573200, 2.4; 19576800, 2.5; 19580400, 2.5; 19584000, 2.3; 19587600, 2.3; 19591200, 2.6; 19594800, 2.9; 19598400, 3.2; 19602000, 3.5; 19605600, 3.8; 19609200, 3.9; 19612800, 3.5; 19616400, 2.6; 19620000, 1.5; 19623600, 0.6; 19627200, 0.4; 19630800, 0.1; 19634400, 0.7; 19638000, 1.2; 19641600, 1; 19645200, 0.9; 19648800, 1.2; 19652400, 1.6; 19656000, 1.9; 19659600, 2.2; 19663200, 2.3; 19666800, 2.8; 19670400, 2.9; 19674000, 3.3; 19677600, 3.4; 19681200, 3.1; 19684800, 2.7; 19688400, 2.3; 19692000, 2.4; 19695600, 2.4; 19699200, 2.1; 19702800, 2.3; 19706400, 2.7; 19710000, 3.1; 19713600, 3.4; 19717200, 3.3; 19720800, 3.6; 19724400, 3.7; 19728000, 3.9; 19731600, 4; 19735200, 4.6; 19738800, 3.8; 19742400, 3.5; 19746000, 3.1; 19749600, 3.5; 19753200, 4.2; 19756800, 4.4; 19760400, 5.2; 19764000, 5.7; 19767600, 6.2; 19771200, 6.4; 19774800, 6.4; 19778400, 6.4; 19782000, 6.7; 19785600, 6.8; 19789200, 6.2; 19792800, 5.7; 19796400, 5.7; 19800000, 6.1; 19803600, 6.4; 19807200, 6.1; 19810800, 5.8; 19814400, 5.9; 19818000, 5.6; 19821600, 5.3; 19825200, 4.8; 19828800, 4.9; 19832400, 5.5; 19836000, 6.1; 19839600, 6.6; 19843200, 6.9; 19846800, 6.6; 19850400, 5.6; 19854000, 4.4; 19857600, 3.5; 19861200, 3.1; 19864800, 2.8; 19868400, 2.6; 19872000, 2.4; 19875600, 2.7; 19879200, 3.8; 19882800, 5.4; 19886400, 6.6; 19890000, 7.4; 19893600, 7.9; 19897200, 8.4; 19900800, 9; 19904400, 9.5; 19908000, 9.9; 19911600, 9.9; 19915200, 9.3; 19918800, 8.4; 19922400, 7.9; 19926000, 7.8; 19929600, 7.9; 19933200, 8; 19936800, 8.1; 19940400, 8; 19944000, 8.1; 19947600, 8.3; 19951200, 8.3; 19954800, 8.2; 19958400, 8; 19962000, 7.5; 19965600, 7.5; 19969200, 8.1; 19972800, 8.1; 19976400, 7.8; 19980000, 7.4; 19983600, 6.9; 19987200, 6.5; 19990800, 6.1; 19994400, 5.6; 19998000, 5.2; 20001600, 5.1; 20005200, 5.3; 20008800, 5.2; 20012400, 4.7; 20016000, 3.9; 20019600, 2.9; 20023200, 2.1; 20026800, 2; 20030400, 2.8; 20034000, 3.9; 20037600, 4.8; 20041200, 5.5; 20044800, 5.7; 20048400, 5.3; 20052000, 4.6; 20055600, 4.9; 20059200, 5.2; 20062800, 5.3; 20066400, 5.6; 20070000, 5.8; 20073600, 6; 20077200, 6; 20080800, 5.8; 20084400, 5.6; 20088000, 5.2; 20091600, 5.2; 20095200, 5.5; 20098800, 5.6; 20102400, 5.3; 20106000, 4.5; 20109600, 3.9; 20113200, 3.6; 20116800, 3.4; 20120400, 3.4; 20124000, 3.7; 20127600, 4.1; 20131200, 4.8; 20134800, 5.2; 20138400, 4.8; 20142000, 4.7; 20145600, 4.7; 20149200, 4.9; 20152800, 4.9; 20156400, 4.8; 20160000, 4.8; 20163600, 4.9; 20167200, 5; 20170800, 5.1; 20174400, 5.4; 20178000, 5.9; 20181600, 6.2; 20185200, 6.2; 20188800, 6.1; 20192400, 6.1; 20196000, 6; 20199600, 5.7; 20203200, 5.3; 20206800, 5.2; 20210400, 5.4; 20214000, 5.3; 20217600, 4.9; 20221200, 4.6; 20224800, 3.9; 20228400, 3.6; 20232000, 4.3; 20235600, 4.9; 20239200, 4.9; 20242800, 4.8; 20246400, 4.9; 20250000, 4.8; 20253600, 4.7; 20257200, 4.8; 20260800, 4.9; 20264400, 5.3; 20268000, 5.6; 20271600, 5.7; 20275200, 5.7; 20278800, 5.6; 20282400, 5.3; 20286000, 4.9; 20289600, 4.3; 20293200, 3.8; 20296800, 3.5; 20300400, 3.4; 20304000, 3.4; 20307600, 3.8; 20311200, 4.1; 20314800, 4.4; 20318400, 4.5; 20322000, 4.5; 20325600, 4.4; 20329200, 4.1; 20332800, 3.6; 20336400, 3.1; 20340000, 2.8; 20343600, 3; 20347200, 3.3; 20350800, 4.4; 20354400, 5.6; 20358000, 6.7; 20361600, 7.4; 20365200, 7.8; 20368800, 7.3; 20372400, 6.5; 20376000, 6.8; 20379600, 6.3; 20383200, 4.7; 20386800, 2.4; 20390400, 0.5; 20394000, 4.8; 20397600, 13.5; 20401200, 18; 20404800, 17.9; 20408400, 17.3; 20412000, 15.9; 20415600, 14.3; 20419200, 13.2; 20422800, 11.9; 20426400, 10.7; 20430000, 9.2; 20433600, 7.4; 20437200, 6.5; 20440800, 6.7; 20444400, 7.3; 20448000, 7.6; 20451600, 7.8; 20455200, 8.1; 20458800, 8.5; 20462400, 8.6; 20466000, 8.6; 20469600, 8.7; 20473200, 8.5; 20476800, 7.7; 20480400, 6.7; 20484000, 5.4; 20487600, 4.7; 20491200, 4.6; 20494800, 4.5; 20498400, 4.2; 20502000, 4; 20505600, 4; 20509200, 4.1; 20512800, 4.1; 20516400, 3.9; 20520000, 3.8; 20523600, 3.6; 20527200, 3.1; 20530800, 2.4; 20534400, 1.4; 20538000, 0.8; 20541600, 1.6; 20545200, 2.6; 20548800, 3.4; 20552400, 4; 20556000, 4.5; 20559600, 4.7; 20563200, 4.9; 20566800, 5.1; 20570400, 4.5; 20574000, 4.1; 20577600, 4.8; 20581200, 5.3; 20584800, 5.4; 20588400, 5.4; 20592000, 5.4; 20595600, 5.4; 20599200, 5.3; 20602800, 5; 20606400, 4.7; 20610000, 4.5; 20613600, 4.2; 20617200, 4.1; 20620800, 4.5; 20624400, 5.4; 20628000, 6.7; 20631600, 7.8; 20635200, 8.3; 20638800, 8.1; 20642400, 8; 20646000, 8.1; 20649600, 7.9; 20653200, 7.1; 20656800, 5.7; 20660400, 5.4; 20664000, 5.3; 20667600, 5.1; 20671200, 4.6; 20674800, 4.2; 20678400, 4; 20682000, 3.7; 20685600, 3.4; 20689200, 3; 20692800, 2.3; 20696400, 1.7; 20700000, 1.8; 20703600, 2.3; 20707200, 3; 20710800, 3.6; 20714400, 4.2; 20718000, 4.9; 20721600, 5.6; 20725200, 6.3; 20728800, 6.9; 20732400, 6.9; 20736000, 6; 20739600, 4.7; 20743200, 3.4; 20746800, 3.2; 20750400, 3.6; 20754000, 4.2; 20757600, 5; 20761200, 5.9; 20764800, 6.5; 20768400, 6.4; 20772000, 5.7; 20775600, 4.8; 20779200, 3.9; 20782800, 2.9; 20786400, 2.2; 20790000, 2.1; 20793600, 2.1; 20797200, 1.9; 20800800, 2.1; 20804400, 2.8; 20808000, 3.6; 20811600, 4.1; 20815200, 4.5; 20818800, 4.4; 20822400, 4.1; 20826000, 3.5; 20829600, 2.8; 20833200, 2.3; 20836800, 1.8; 20840400, 0.6; 20844000, 0.6; 20847600, 1.5; 20851200, 1.9; 20854800, 2.1; 20858400, 1.9; 20862000, 1.6; 20865600, 1.9; 20869200, 2.7; 20872800, 3.7; 20876400, 4.6; 20880000, 5.5; 20883600, 6.3; 20887200, 6.8; 20890800, 7.1; 20894400, 7.1; 20898000, 6.8; 20901600, 6.1; 20905200, 5.5; 20908800, 4.9; 20912400, 4.3; 20916000, 3.9; 20919600, 4.1; 20923200, 4.7; 20926800, 5.7; 20930400, 6.8; 20934000, 7.9; 20937600, 8.7; 20941200, 8.8; 20944800, 8.1; 20948400, 6.2; 20952000, 4.6; 20955600, 4.4; 20959200, 5; 20962800, 6; 20966400, 6.5; 20970000, 6.3; 20973600, 5.5; 20977200, 4.9; 20980800, 5.6; 20984400, 6.7; 20988000, 7.3; 20991600, 7.5; 20995200, 7.5; 20998800, 7.7; 21002400, 7.7; 21006000, 6.8; 21009600, 6.1; 21013200, 5.8; 21016800, 5.3; 21020400, 5.1; 21024000, 5.2; 21027600, 5.3; 21031200, 4.8; 21034800, 4.7; 21038400, 5.2; 21042000, 5.2; 21045600, 4.5; 21049200, 3.8; 21052800, 3.2; 21056400, 2.7; 21060000, 1.9; 21063600, 1.1; 21067200, 0.7; 21070800, 0.5; 21074400, 0.6; 21078000, 1; 21081600, 1.4; 21085200, 1.9; 21088800, 1.9; 21092400, 1.3; 21096000, 1.2; 21099600, 1.1; 21103200, 1.3; 21106800, 1.6; 21110400, 1.7; 21114000, 1.9; 21117600, 2.1; 21121200, 2.6; 21124800, 3.9; 21128400, 5.7; 21132000, 6.8; 21135600, 7.1; 21139200, 6.6; 21142800, 5.7; 21146400, 4.8; 21150000, 3.8; 21153600, 2.9; 21157200, 2; 21160800, 1.2; 21164400, 1.2; 21168000, 1.9; 21171600, 2.4; 21175200, 2.4; 21178800, 2.4; 21182400, 2.2; 21186000, 2.3; 21189600, 2.6; 21193200, 2.8; 21196800, 2.9; 21200400, 3.2; 21204000, 3.6; 21207600, 4.4; 21211200, 5.6; 21214800, 6.9; 21218400, 7.3; 21222000, 7.2; 21225600, 7.3; 21229200, 7.1; 21232800, 6.7; 21236400, 6.3; 21240000, 5.9; 21243600, 5.5; 21247200, 5.2; 21250800, 4.7; 21254400, 5.6; 21258000, 5.9; 21261600, 6.2; 21265200, 4.9; 21268800, 3.6; 21272400, 3.3; 21276000, 2.9; 21279600, 2.4; 21283200, 2.2; 21286800, 2.6; 21290400, 3.6; 21294000, 5; 21297600, 6.9; 21301200, 8.2; 21304800, 8.4; 21308400, 8.3; 21312000, 7.9; 21315600, 6.9; 21319200, 5.6; 21322800, 5.2; 21326400, 6.1; 21330000, 6; 21333600, 5; 21337200, 3.8; 21340800, 2.7; 21344400, 2; 21348000, 1.4; 21351600, 0.6; 21355200, 0.7; 21358800, 1.6; 21362400, 2.4; 21366000, 3; 21369600, 3.3; 21373200, 3.5; 21376800, 3.5; 21380400, 3.1; 21384000, 2.3; 21387600, 0.9; 21391200, 0.8; 21394800, 2.5; 21398400, 4; 21402000, 4.8; 21405600, 5.1; 21409200, 5.3; 21412800, 5.6; 21416400, 5.6; 21420000, 5.6; 21423600, 5.5; 21427200, 5.1; 21430800, 4.6; 21434400, 3.6; 21438000, 2.9; 21441600, 3.3; 21445200, 3.2; 21448800, 3; 21452400, 2.9; 21456000, 3.3; 21459600, 4.2; 21463200, 5.6; 21466800, 8.4; 21470400, 10.6; 21474000, 9.4; 21477600, 6.6; 21481200, 4.2; 21484800, 4; 21488400, 5.5; 21492000, 6.4; 21495600, 6.1; 21499200, 5.6; 21502800, 5.3; 21506400, 4.9; 21510000, 4.1; 21513600, 2.7; 21517200, 1; 21520800, 0.8; 21524400, 1.8; 21528000, 2.3; 21531600, 1.9; 21535200, 1.2; 21538800, 1.7; 21542400, 3; 21546000, 3.8; 21549600, 4.3; 21553200, 4.4; 21556800, 4.5; 21560400, 4.7; 21564000, 4.4; 21567600, 3.9; 21571200, 3.3; 21574800, 2.6; 21578400, 1.9; 21582000, 1.6; 21585600, 2.1; 21589200, 3.2; 21592800, 4.2; 21596400, 4.8; 21600000, 4.8; 21603600, 4.6; 21607200, 3.8; 21610800, 3; 21614400, 2.2; 21618000, 1.5; 21621600, 1; 21625200, 0.9; 21628800, 1.1; 21632400, 1.3; 21636000, 1.3; 21639600, 1.3; 21643200, 1.4; 21646800, 2.2; 21650400, 3.2; 21654000, 4.4; 21657600, 5; 21661200, 4.9; 21664800, 4.7; 21668400, 4.8; 21672000, 5.3; 21675600, 5.9; 21679200, 6.4; 21682800, 6.3; 21686400, 5.8; 21690000, 5.2; 21693600, 4; 21697200, 3; 21700800, 3.2; 21704400, 4; 21708000, 4.3; 21711600, 4.3; 21715200, 4; 21718800, 3.6; 21722400, 3; 21726000, 2.4; 21729600, 2.3; 21733200, 2.8; 21736800, 3.4; 21740400, 3.8; 21744000, 3.7; 21747600, 3.4; 21751200, 3; 21754800, 2.7; 21758400, 2.3; 21762000, 2.1; 21765600, 2.3; 21769200, 3.4; 21772800, 4.2; 21776400, 4.2; 21780000, 3.7; 21783600, 3.2; 21787200, 3; 21790800, 2.7; 21794400, 1.9; 21798000, 1.6; 21801600, 2.1; 21805200, 3; 21808800, 4.1; 21812400, 5.5; 21816000, 7.4; 21819600, 8.5; 21823200, 8.5; 21826800, 8.1; 21830400, 7.9; 21834000, 7.7; 21837600, 7.5; 21841200, 7.1; 21844800, 6.8; 21848400, 6.7; 21852000, 6.5; 21855600, 6.4; 21859200, 6.1; 21862800, 5.6; 21866400, 4.4; 21870000, 2.8; 21873600, 2.2; 21877200, 1.5; 21880800, 0.8; 21884400, 2.1; 21888000, 2.2; 21891600, 1.7; 21895200, 1.3; 21898800, 1.2; 21902400, 2.1; 21906000, 3.7; 21909600, 4.9; 21913200, 5.5; 21916800, 5.9; 21920400, 5.7; 21924000, 4.8; 21927600, 3.4; 21931200, 1.7; 21934800, 0.3; 21938400, 1.7; 21942000, 2.9; 21945600, 3.6; 21949200, 3.6; 21952800, 3.1; 21956400, 2.7; 21960000, 2.7; 21963600, 2.8; 21967200, 2.9; 21970800, 2.7; 21974400, 2.5; 21978000, 2.4; 21981600, 2.5; 21985200, 3.2; 21988800, 4.2; 21992400, 5.3; 21996000, 5.5; 21999600, 5.1; 22003200, 4.9; 22006800, 4.8; 22010400, 4.6; 22014000, 4.4; 22017600, 3.9; 22021200, 3.7; 22024800, 3.6; 22028400, 3.7; 22032000, 3.8; 22035600, 3.6; 22039200, 2.8; 22042800, 2; 22046400, 1; 22050000, 0.2; 22053600, 0.5; 22057200, 1.1; 22060800, 2; 22064400, 3; 22068000, 3.8; 22071600, 4.7; 22075200, 5.8; 22078800, 6.7; 22082400, 6.7; 22086000, 6.2; 22089600, 5.8; 22093200, 5.4; 22096800, 5.3; 22100400, 5.5; 22104000, 5.4; 22107600, 5.2; 22111200, 4.8; 22114800, 4.4; 22118400, 4.1; 22122000, 4; 22125600, 3.6; 22129200, 3.2; 22132800, 3.2; 22136400, 3; 22140000, 2.8; 22143600, 2.5; 22147200, 2.3; 22150800, 2.1; 22154400, 2.1; 22158000, 2.6; 22161600, 3.9; 22165200, 5.5; 22168800, 6.4; 22172400, 6.6; 22176000, 6.6; 22179600, 6.5; 22183200, 6.3; 22186800, 5.9; 22190400, 5.4; 22194000, 4.8; 22197600, 4.2; 22201200, 3.8; 22204800, 3.3; 22208400, 2.9; 22212000, 2.6; 22215600, 2.4; 22219200, 2.2; 22222800, 2; 22226400, 1.6; 22230000, 0.9; 22233600, 0.2; 22237200, 0.9; 22240800, 2; 22244400, 3.6; 22248000, 5.5; 22251600, 6.7; 22255200, 7; 22258800, 6.8; 22262400, 6.3; 22266000, 5.8; 22269600, 5.4; 22273200, 5.1; 22276800, 4.6; 22280400, 4.1; 22284000, 3.9; 22287600, 3.6; 22291200, 3.1; 22294800, 2.3; 22298400, 1.3; 22302000, 0.7; 22305600, 0.7; 22309200, 1.1; 22312800, 1.6; 22316400, 2; 22320000, 2.5; 22323600, 3; 22327200, 3.7; 22330800, 5; 22334400, 6.9; 22338000, 7.7; 22341600, 7.7; 22345200, 7.5; 22348800, 7.1; 22352400, 6.5; 22356000, 5.8; 22359600, 5; 22363200, 4.2; 22366800, 3.8; 22370400, 4.3; 22374000, 5.7; 22377600, 7.6; 22381200, 9.2; 22384800, 9.9; 22388400, 10.2; 22392000, 9.6; 22395600, 8.8; 22399200, 8.6; 22402800, 8.7; 22406400, 8.9; 22410000, 9.1; 22413600, 9.3; 22417200, 9.2; 22420800, 8.6; 22424400, 8.1; 22428000, 7.7; 22431600, 7.7; 22435200, 7.7; 22438800, 7.7; 22442400, 7.8; 22446000, 7.9; 22449600, 7.7; 22453200, 7.6; 22456800, 7.8; 22460400, 8.2; 22464000, 8.5; 22467600, 8.4; 22471200, 8.1; 22474800, 7.7; 22478400, 7.3; 22482000, 6.7; 22485600, 6.3; 22489200, 6.2; 22492800, 6.2; 22496400, 6; 22500000, 6; 22503600, 6.1; 22507200, 6.1; 22510800, 5.8; 22514400, 5.8; 22518000, 5.9; 22521600, 6.5; 22525200, 7.5; 22528800, 8; 22532400, 8.1; 22536000, 8; 22539600, 8.1; 22543200, 8; 22546800, 8; 22550400, 7.9; 22554000, 7.7; 22557600, 7; 22561200, 6.4; 22564800, 6.3; 22568400, 6.4; 22572000, 6.3; 22575600, 5.9; 22579200, 5.5; 22582800, 5.1; 22586400, 4.9; 22590000, 4.5; 22593600, 4; 22597200, 3.4; 22600800, 2.9; 22604400, 2.5; 22608000, 2.4; 22611600, 2.8; 22615200, 3.3; 22618800, 3.9; 22622400, 4.6; 22626000, 5.4; 22629600, 5.9; 22633200, 6; 22636800, 5.9; 22640400, 5.5; 22644000, 4.7; 22647600, 3.9; 22651200, 3.6; 22654800, 3.2; 22658400, 2.8; 22662000, 2.4; 22665600, 2.2; 22669200, 2.2; 22672800, 2.5; 22676400, 3.1; 22680000, 4.2; 22683600, 5.2; 22687200, 5.7; 22690800, 5.9; 22694400, 6.1; 22698000, 6.4; 22701600, 6.9; 22705200, 7.3; 22708800, 7.5; 22712400, 7.6; 22716000, 7.7; 22719600, 7.5; 22723200, 7.2; 22726800, 6.8; 22730400, 5.9; 22734000, 4.8; 22737600, 4.5; 22741200, 3.8; 22744800, 3.2; 22748400, 3; 22752000, 3.2; 22755600, 3.6; 22759200, 4; 22762800, 4.5; 22766400, 5.9; 22770000, 7.2; 22773600, 7.7; 22777200, 7.9; 22780800, 7.9; 22784400, 7.9; 22788000, 7.9; 22791600, 7.8; 22795200, 7.6; 22798800, 7.3; 22802400, 6.9; 22806000, 6.5; 22809600, 6.1; 22813200, 5.9; 22816800, 5.2; 22820400, 4.2; 22824000, 4.8; 22827600, 5.8; 22831200, 5.6; 22834800, 5.3; 22838400, 5.1; 22842000, 5; 22845600, 4.9; 22849200, 4.5; 22852800, 4.8; 22856400, 5.7; 22860000, 6.7; 22863600, 7.6; 22867200, 8.1; 22870800, 8.2; 22874400, 8; 22878000, 7.7; 22881600, 6.9; 22885200, 6.4; 22888800, 6.3; 22892400, 6.4; 22896000, 6.5; 22899600, 6.7; 22903200, 6.2; 22906800, 6.5; 22910400, 7.1; 22914000, 7.1; 22917600, 7; 22921200, 6.8; 22924800, 6.4; 22928400, 5.8; 22932000, 5.2; 22935600, 4.6; 22939200, 3.9; 22942800, 3.4; 22946400, 4.1; 22950000, 5.7; 22953600, 7.2; 22957200, 8.3; 22960800, 8.5; 22964400, 7.8; 22968000, 7.2; 22971600, 7; 22975200, 7.2; 22978800, 7.6; 22982400, 7.9; 22986000, 7.9; 22989600, 7.2; 22993200, 8.1; 22996800, 9.6; 23000400, 10.2; 23004000, 10.2; 23007600, 10.2; 23011200, 10.3; 23014800, 10.2; 23018400, 9.7; 23022000, 8.6; 23025600, 8.4; 23029200, 8.3; 23032800, 8; 23036400, 7.7; 23040000, 7.2; 23043600, 6.5; 23047200, 5.8; 23050800, 5.2; 23054400, 4.7; 23058000, 4.6; 23061600, 4.7; 23065200, 5; 23068800, 5.1; 23072400, 5; 23076000, 3.8; 23079600, 2.2; 23083200, 2.4; 23086800, 3.3; 23090400, 3.9; 23094000, 4.4; 23097600, 5.1; 23101200, 6.1; 23104800, 7.6; 23108400, 8.7; 23112000, 8.9; 23115600, 8.6; 23119200, 8.3; 23122800, 7.9; 23126400, 7.6; 23130000, 7; 23133600, 6.4; 23137200, 5.9; 23140800, 5.6; 23144400, 6; 23148000, 6.7; 23151600, 7; 23155200, 7.1; 23158800, 7; 23162400, 5.9; 23166000, 5.1; 23169600, 4.8; 23173200, 4.5; 23176800, 4; 23180400, 3.6; 23184000, 3.2; 23187600, 2.7; 23191200, 2.1; 23194800, 1.7; 23198400, 2.1; 23202000, 2.3; 23205600, 2; 23209200, 1.8; 23212800, 1; 23216400, 0.7; 23220000, 2.2; 23223600, 3.6; 23227200, 4.7; 23230800, 5.3; 23234400, 6; 23238000, 6.7; 23241600, 7.9; 23245200, 9.4; 23248800, 11.5; 23252400, 13.4; 23256000, 13.1; 23259600, 11.9; 23263200, 10.9; 23266800, 9.9; 23270400, 8.9; 23274000, 8.1; 23277600, 7.4; 23281200, 6.7; 23284800, 6.6; 23288400, 7.3; 23292000, 8.6; 23295600, 10.6; 23299200, 12.3; 23302800, 12.9; 23306400, 13; 23310000, 12.6; 23313600, 12.2; 23317200, 11.7; 23320800, 11.5; 23324400, 11.3; 23328000, 11.1; 23331600, 10.8; 23335200, 10.7; 23338800, 10.6; 23342400, 9.8; 23346000, 8.9; 23349600, 8; 23353200, 7.1; 23356800, 6.3; 23360400, 5.7; 23364000, 5.2; 23367600, 4.7; 23371200, 4.2; 23374800, 3.8; 23378400, 3.2; 23382000, 2.4; 23385600, 1.5; 23389200, 0.5; 23392800, 0.7; 23396400, 1.5; 23400000, 2.5; 23403600, 3.4; 23407200, 4; 23410800, 4.4; 23414400, 4.5; 23418000, 4.2; 23421600, 3.7; 23425200, 2.8; 23428800, 1.9; 23432400, 1.2; 23436000, 0.8; 23439600, 0.5; 23443200, 0.3; 23446800, 0.7; 23450400, 0.9; 23454000, 0.8; 23457600, 0.8; 23461200, 1.5; 23464800, 2.4; 23468400, 3.5; 23472000, 4.2; 23475600, 4.7; 23479200, 5.4; 23482800, 5.9; 23486400, 5.9; 23490000, 5.8; 23493600, 5.5; 23497200, 5.2; 23500800, 4.9; 23504400, 4.7; 23508000, 4.1; 23511600, 3.3; 23515200, 2.7; 23518800, 2; 23522400, 1.7; 23526000, 1.6; 23529600, 1.7; 23533200, 2; 23536800, 2.5; 23540400, 3.1; 23544000, 4.3; 23547600, 5.6; 23551200, 6.5; 23554800, 7.1; 23558400, 7.4; 23562000, 7.8; 23565600, 7.9; 23569200, 7.8; 23572800, 7.4; 23576400, 6.7; 23580000, 6.2; 23583600, 5.8; 23587200, 5.6; 23590800, 5.5; 23594400, 4.7; 23598000, 3.6; 23601600, 3.2; 23605200, 3.9; 23608800, 4.6; 23612400, 4.5; 23616000, 4.3; 23619600, 3.7; 23623200, 2.9; 23626800, 2; 23630400, 1.2; 23634000, 1.2; 23637600, 2.2; 23641200, 3.1; 23644800, 3.8; 23648400, 4.3; 23652000, 4.5; 23655600, 4.5; 23659200, 4.3; 23662800, 4; 23666400, 3.7; 23670000, 3.3; 23673600, 3.2; 23677200, 3.2; 23680800, 3.2; 23684400, 3; 23688000, 3.5; 23691600, 3.7; 23695200, 4.3; 23698800, 4.5; 23702400, 4.4; 23706000, 4.5; 23709600, 4.9; 23713200, 5.6; 23716800, 7.2; 23720400, 8.8; 23724000, 9.2; 23727600, 9; 23731200, 9.1; 23734800, 8.7; 23738400, 8.3; 23742000, 8.1; 23745600, 8; 23749200, 7.7; 23752800, 7.7; 23756400, 7.9; 23760000, 7.3; 23763600, 6.1; 23767200, 5.2; 23770800, 5.2; 23774400, 4.7; 23778000, 3.4; 23781600, 1.6; 23785200, 1.2; 23788800, 2.6; 23792400, 3.9; 23796000, 5.1; 23799600, 6.6; 23803200, 8.1; 23806800, 9.3; 23810400, 10.7; 23814000, 11.5; 23817600, 12.1; 23821200, 11.5; 23824800, 10.9; 23828400, 10.5; 23832000, 10.3; 23835600, 10.2; 23839200, 10.8; 23842800, 11.7; 23846400, 10.7; 23850000, 9; 23853600, 7.2; 23857200, 7.4; 23860800, 7.6; 23864400, 7.2; 23868000, 6.8; 23871600, 6.4; 23875200, 6; 23878800, 5.8; 23882400, 5.3; 23886000, 3.9; 23889600, 2.1; 23893200, 3.1; 23896800, 5.8; 23900400, 8.3; 23904000, 9.3; 23907600, 9; 23911200, 7.9; 23914800, 7.9; 23918400, 9; 23922000, 9.5; 23925600, 9; 23929200, 8.4; 23932800, 8.2; 23936400, 7.9; 23940000, 7.3; 23943600, 6.4; 23947200, 5.1; 23950800, 3.8; 23954400, 2.8; 23958000, 2.7; 23961600, 2.7; 23965200, 2.5; 23968800, 2.4; 23972400, 2.5; 23976000, 3.1; 23979600, 3.9; 23983200, 4.6; 23986800, 5; 23990400, 5.3; 23994000, 5.4; 23997600, 5; 24001200, 4.7; 24004800, 4.9; 24008400, 5.3; 24012000, 5.6; 24015600, 5.6; 24019200, 5.5; 24022800, 5.3; 24026400, 5.3; 24030000, 6.4; 24033600, 7.5; 24037200, 8; 24040800, 8.6; 24044400, 8.8; 24048000, 9.1; 24051600, 9.9; 24055200, 9.9; 24058800, 9.5; 24062400, 9.5; 24066000, 9.4; 24069600, 9.3; 24073200, 9.2; 24076800, 9.2; 24080400, 9.1; 24084000, 8.7; 24087600, 8.1; 24091200, 6.7; 24094800, 4.6; 24098400, 5.1; 24102000, 6.6; 24105600, 7; 24109200, 6.9; 24112800, 6.5; 24116400, 5.6; 24120000, 5.9; 24123600, 4.9; 24127200, 3.4; 24130800, 2.8; 24134400, 2.9; 24138000, 3; 24141600, 2.6; 24145200, 3.2; 24148800, 4; 24152400, 4.8; 24156000, 5.3; 24159600, 5.6; 24163200, 6.1; 24166800, 6.8; 24170400, 8.1; 24174000, 8.7; 24177600, 9.2; 24181200, 9.4; 24184800, 9.6; 24188400, 10.3; 24192000, 11.6; 24195600, 12.2; 24199200, 11.6; 24202800, 12.6; 24206400, 11.9; 24210000, 10.4; 24213600, 9.6; 24217200, 9.4; 24220800, 9.4; 24224400, 9.4; 24228000, 9.2; 24231600, 8.8; 24235200, 8.1; 24238800, 7.8; 24242400, 7.9; 24246000, 8.2; 24249600, 8.5; 24253200, 8.5; 24256800, 8.3; 24260400, 7.7; 24264000, 6.8; 24267600, 5.8; 24271200, 4.8; 24274800, 3.9; 24278400, 2.6; 24282000, 1.7; 24285600, 2.5; 24289200, 3.5; 24292800, 4.8; 24296400, 5.4; 24300000, 5.6; 24303600, 6; 24307200, 6.3; 24310800, 6.2; 24314400, 6.1; 24318000, 5.9; 24321600, 5.3; 24325200, 4.8; 24328800, 4.4; 24332400, 4.6; 24336000, 5; 24339600, 5.5; 24343200, 6.2; 24346800, 6.9; 24350400, 7.3; 24354000, 7.6; 24357600, 7.8; 24361200, 7.9; 24364800, 8.2; 24368400, 8.5; 24372000, 7.7; 24375600, 6.6; 24379200, 8.2; 24382800, 10; 24386400, 10.5; 24390000, 10.5; 24393600, 10.2; 24397200, 10; 24400800, 9.9; 24404400, 9; 24408000, 8.4; 24411600, 9.1; 24415200, 9.9; 24418800, 10.3; 24422400, 9.9; 24426000, 9.2; 24429600, 8.5; 24433200, 7.9; 24436800, 7.4; 24440400, 6.4; 24444000, 4.5; 24447600, 2.1; 24451200, 0.4; 24454800, 1.5; 24458400, 2.9; 24462000, 3.8; 24465600, 5.1; 24469200, 6.2; 24472800, 6.9; 24476400, 7.3; 24480000, 7.5; 24483600, 7.8; 24487200, 7.9; 24490800, 8.1; 24494400, 8.3; 24498000, 8.6; 24501600, 8.5; 24505200, 8.1; 24508800, 7.8; 24512400, 7.3; 24516000, 6.7; 24519600, 6.1; 24523200, 5.7; 24526800, 5.5; 24530400, 5.5; 24534000, 5.6; 24537600, 5.6; 24541200, 5.5; 24544800, 5.1; 24548400, 4.3; 24552000, 4; 24555600, 3.5; 24559200, 3; 24562800, 2.8; 24566400, 2.9; 24570000, 3.3; 24573600, 3.6; 24577200, 4; 24580800, 4.9; 24584400, 5.5; 24588000, 5.7; 24591600, 5.8; 24595200, 6.4; 24598800, 7; 24602400, 7; 24606000, 6.6; 24609600, 6.2; 24613200, 5.8; 24616800, 5.5; 24620400, 5.2; 24624000, 4.6; 24627600, 4.4; 24631200, 4.2; 24634800, 3.5; 24638400, 3.7; 24642000, 4.5; 24645600, 5; 24649200, 5; 24652800, 5.2; 24656400, 5.6; 24660000, 5.8; 24663600, 5.5; 24667200, 5.8; 24670800, 7.1; 24674400, 9.2; 24678000, 11; 24681600, 11.4; 24685200, 10.7; 24688800, 9.4; 24692400, 9; 24696000, 9; 24699600, 9; 24703200, 9; 24706800, 9; 24710400, 8.8; 24714000, 8.6; 24717600, 7.7; 24721200, 6.2; 24724800, 6; 24728400, 8.6; 24732000, 8.9; 24735600, 8.7; 24739200, 7.8; 24742800, 6.3; 24746400, 5.5; 24750000, 5.1; 24753600, 4.6; 24757200, 4; 24760800, 3.4; 24764400, 3; 24768000, 2.8; 24771600, 3.3; 24775200, 4.1; 24778800, 4.5; 24782400, 4.3; 24786000, 3.7; 24789600, 3.2; 24793200, 3; 24796800, 3.1; 24800400, 3.4; 24804000, 3.6; 24807600, 3.5; 24811200, 3.2; 24814800, 3.5; 24818400, 4; 24822000, 4.1; 24825600, 4.3; 24829200, 4.6; 24832800, 4.4; 24836400, 3.4; 24840000, 1.9; 24843600, 0.2; 24847200, 1.4; 24850800, 2.6; 24854400, 3.5; 24858000, 4.2; 24861600, 4.9; 24865200, 5.7; 24868800, 6.2; 24872400, 6.3; 24876000, 6; 24879600, 5.5; 24883200, 5; 24886800, 4.1; 24890400, 3; 24894000, 2.5; 24897600, 2.2; 24901200, 2.3; 24904800, 2.3; 24908400, 2.4; 24912000, 2.6; 24915600, 2.5; 24919200, 2.8; 24922800, 2.8; 24926400, 2.7; 24930000, 3.2; 24933600, 4.1; 24937200, 5.1; 24940800, 5.6; 24944400, 5.7; 24948000, 5.3; 24951600, 4.9; 24955200, 4.6; 24958800, 4.2; 24962400, 4.3; 24966000, 4.5; 24969600, 4.8; 24973200, 5.5; 24976800, 6; 24980400, 5.9; 24984000, 6.8; 24987600, 7.3; 24991200, 7.3; 24994800, 7; 24998400, 6.6; 25002000, 6.4; 25005600, 6.7; 25009200, 7.2; 25012800, 8; 25016400, 8.3; 25020000, 8.3; 25023600, 8.3; 25027200, 8.3; 25030800, 8.2; 25034400, 7.9; 25038000, 7.5; 25041600, 7.2; 25045200, 6.9; 25048800, 6.7; 25052400, 6.6; 25056000, 6.6; 25059600, 6.5; 25063200, 6; 25066800, 4.7; 25070400, 3.8; 25074000, 3; 25077600, 2.6; 25081200, 2.7; 25084800, 2.9; 25088400, 3; 25092000, 2.9; 25095600, 2.7; 25099200, 2.6; 25102800, 2.8; 25106400, 3.1; 25110000, 3.3; 25113600, 3.5; 25117200, 3.9; 25120800, 4.3; 25124400, 4.6; 25128000, 4.7; 25131600, 4.8; 25135200, 4.9; 25138800, 5; 25142400, 5.1; 25146000, 5.2; 25149600, 5.3; 25153200, 4.5; 25156800, 4.8; 25160400, 5.3; 25164000, 4.7; 25167600, 3.9; 25171200, 3.1; 25174800, 2.2; 25178400, 1.5; 25182000, 1.4; 25185600, 2.5; 25189200, 3.8; 25192800, 5.2; 25196400, 6.2; 25200000, 6.7; 25203600, 6.9; 25207200, 6.8; 25210800, 6.6; 25214400, 6.7; 25218000, 7.6; 25221600, 9; 25225200, 10.5; 25228800, 12.3; 25232400, 13.6; 25236000, 14.5; 25239600, 15.4; 25243200, 15.2; 25246800, 14.6; 25250400, 14.3; 25254000, 14.1; 25257600, 13.7; 25261200, 13.1; 25264800, 12.3; 25268400, 11.1; 25272000, 10.2; 25275600, 10.2; 25279200, 10.3; 25282800, 10.1; 25286400, 9.8; 25290000, 9.4; 25293600, 9; 25297200, 8.7; 25300800, 8.4; 25304400, 8.2; 25308000, 8; 25311600, 8; 25315200, 7.8; 25318800, 7.5; 25322400, 7.1; 25326000, 6.7; 25329600, 6.3; 25333200, 5.8; 25336800, 5.6; 25340400, 5.5; 25344000, 5.9; 25347600, 6.3; 25351200, 6.5; 25354800, 6.3; 25358400, 6.3; 25362000, 6.4; 25365600, 5.9; 25369200, 5; 25372800, 4.1; 25376400, 4.3; 25380000, 5.9; 25383600, 7.9; 25387200, 9.4; 25390800, 10.2; 25394400, 11; 25398000, 12.5; 25401600, 13.7; 25405200, 13.8; 25408800, 13.8; 25412400, 13.9; 25416000, 13.6; 25419600, 13.1; 25423200, 12.8; 25426800, 12.4; 25430400, 11.9; 25434000, 11.6; 25437600, 11.3; 25441200, 10.8; 25444800, 10.8; 25448400, 11.4; 25452000, 11.6; 25455600, 11.3; 25459200, 10.8; 25462800, 10.4; 25466400, 10.2; 25470000, 9.9; 25473600, 9.6; 25477200, 9.2; 25480800, 8.7; 25484400, 8.3; 25488000, 8.2; 25491600, 8; 25495200, 7.6; 25498800, 7; 25502400, 6.3; 25506000, 5.6; 25509600, 5.1; 25513200, 4.6; 25516800, 4.2; 25520400, 3.9; 25524000, 3.7; 25527600, 3.6; 25531200, 3.6; 25534800, 3.6; 25538400, 3.1; 25542000, 2.4; 25545600, 1.9; 25549200, 1.6; 25552800, 1.2; 25556400, 0.8; 25560000, 0.4; 25563600, 0.9; 25567200, 1.8; 25570800, 2.4; 25574400, 2.8; 25578000, 2.9; 25581600, 2.6; 25585200, 2.6; 25588800, 3.1; 25592400, 3; 25596000, 2.2; 25599600, 2.2; 25603200, 2.7; 25606800, 3.5; 25610400, 3.2; 25614000, 3.2; 25617600, 4.1; 25621200, 4.4; 25624800, 3.8; 25628400, 3.3; 25632000, 3.3; 25635600, 4.1; 25639200, 4.8; 25642800, 5.3; 25646400, 5.3; 25650000, 5.3; 25653600, 5.6; 25657200, 6.4; 25660800, 7.3; 25664400, 7.7; 25668000, 7.1; 25671600, 5.7; 25675200, 4.8; 25678800, 4.7; 25682400, 4.4; 25686000, 4.4; 25689600, 4.5; 25693200, 5; 25696800, 5.7; 25700400, 6.3; 25704000, 7; 25707600, 7.5; 25711200, 7.6; 25714800, 7.3; 25718400, 7; 25722000, 6.6; 25725600, 6.3; 25729200, 5.8; 25732800, 4.9; 25736400, 3.8; 25740000, 2.2; 25743600, 0.6; 25747200, 2.7; 25750800, 5; 25754400, 6.7; 25758000, 7.6; 25761600, 8.5; 25765200, 11.1; 25768800, 12.4; 25772400, 12.5; 25776000, 11.9; 25779600, 11.2; 25783200, 11.1; 25786800, 10.9; 25790400, 10.3; 25794000, 9.8; 25797600, 9.1; 25801200, 8.5; 25804800, 8; 25808400, 7.6; 25812000, 7; 25815600, 6.1; 25819200, 5.5; 25822800, 5.3; 25826400, 5.2; 25830000, 5.2; 25833600, 5.5; 25837200, 5.9; 25840800, 6; 25844400, 5.4; 25848000, 5.4; 25851600, 4.8; 25855200, 4.3; 25858800, 3.9; 25862400, 3.8; 25866000, 4.4; 25869600, 5.3; 25873200, 6.3; 25876800, 6.2; 25880400, 6.2; 25884000, 5.6; 25887600, 5.2; 25891200, 5.4; 25894800, 6.1; 25898400, 6.4; 25902000, 6.4; 25905600, 7; 25909200, 7.6; 25912800, 8; 25916400, 8.2; 25920000, 8.7; 25923600, 9.1; 25927200, 8.6; 25930800, 7.7; 25934400, 7.8; 25938000, 6.9; 25941600, 7.2; 25945200, 8.1; 25948800, 8.9; 25952400, 9.5; 25956000, 9.9; 25959600, 9.7; 25963200, 9.2; 25966800, 9; 25970400, 8.5; 25974000, 8.4; 25977600, 8.3; 25981200, 8.4; 25984800, 8.5; 25988400, 8.4; 25992000, 8.1; 25995600, 7.6; 25999200, 6.9; 26002800, 5.7; 26006400, 4.8; 26010000, 4.3; 26013600, 3.8; 26017200, 3.3; 26020800, 2.4; 26024400, 1.3; 26028000, 0.1; 26031600, 0.8; 26035200, 1.8; 26038800, 2.7; 26042400, 3.4; 26046000, 3.9; 26049600, 5; 26053200, 6.5; 26056800, 7.9; 26060400, 8.9; 26064000, 9.5; 26067600, 10; 26071200, 10.2; 26074800, 10.3; 26078400, 10.5; 26082000, 10.8; 26085600, 10.9; 26089200, 11; 26092800, 11.1; 26096400, 11; 26100000, 10.6; 26103600, 9.8; 26107200, 11.2; 26110800, 12.2; 26114400, 12.3; 26118000, 11.3; 26121600, 9.8; 26125200, 7.9; 26128800, 7; 26132400, 9.9; 26136000, 12.6; 26139600, 12.2; 26143200, 12.1; 26146800, 12.3; 26150400, 11.7; 26154000, 11.5; 26157600, 11.4; 26161200, 11.3; 26164800, 11.6; 26168400, 11.9; 26172000, 11.8; 26175600, 11.6; 26179200, 11.7; 26182800, 11.6; 26186400, 11.3; 26190000, 11.4; 26193600, 10.3; 26197200, 8.7; 26200800, 7.2; 26204400, 6.3; 26208000, 5.9; 26211600, 6.1; 26215200, 6.3; 26218800, 6.2; 26222400, 5.4; 26226000, 4.4; 26229600, 3.8; 26233200, 3.7; 26236800, 3.9; 26240400, 4.2; 26244000, 4.6; 26247600, 4.8; 26251200, 5.1; 26254800, 5.4; 26258400, 5.6; 26262000, 5.5; 26265600, 5; 26269200, 4.3; 26272800, 3.3; 26276400, 2.3; 26280000, 1.3; 26283600, 1.5; 26287200, 1.9; 26290800, 2.3; 26294400, 2.7; 26298000, 3.1; 26301600, 3.4; 26305200, 3.6; 26308800, 4.2; 26312400, 4.8; 26316000, 5.1; 26319600, 5.2; 26323200, 5.4; 26326800, 5.7; 26330400, 6; 26334000, 6.2; 26337600, 6.4; 26341200, 6.6; 26344800, 7; 26348400, 7.4; 26352000, 7.4; 26355600, 7.3; 26359200, 7.2; 26362800, 6.3; 26366400, 5.5; 26370000, 5.9; 26373600, 5.4; 26377200, 5.1; 26380800, 5.2; 26384400, 5.4; 26388000, 5.3; 26391600, 4.7; 26395200, 3.9; 26398800, 3.7; 26402400, 4.7; 26406000, 5.6; 26409600, 6.7; 26413200, 7.7; 26416800, 8.2; 26420400, 8.4; 26424000, 8.4; 26427600, 8.5; 26431200, 8.4; 26434800, 8.4; 26438400, 8.2; 26442000, 7.5; 26445600, 7; 26449200, 7; 26452800, 7.1; 26456400, 7.9; 26460000, 7.9; 26463600, 7.6; 26467200, 6.9; 26470800, 6.2; 26474400, 5.3; 26478000, 4.7; 26481600, 4.3; 26485200, 4.3; 26488800, 4.5; 26492400, 5; 26496000, 5.7; 26499600, 6.2; 26503200, 6.2; 26506800, 5.8; 26510400, 5.2; 26514000, 4.6; 26517600, 4; 26521200, 3.3; 26524800, 2.9; 26528400, 3.1; 26532000, 3.4; 26535600, 3.3; 26539200, 3.1; 26542800, 3; 26546400, 2.8; 26550000, 2.6; 26553600, 2.6; 26557200, 2.6; 26560800, 2.5; 26564400, 2.3; 26568000, 1.9; 26571600, 1.9; 26575200, 2.4; 26578800, 2.9; 26582400, 3.7; 26586000, 4.7; 26589600, 5.3; 26593200, 5.6; 26596800, 5.4; 26600400, 4.9; 26604000, 4; 26607600, 3; 26611200, 2.4; 26614800, 2.1; 26618400, 2; 26622000, 2.6; 26625600, 4.7; 26629200, 7.3; 26632800, 10.3; 26636400, 12.7; 26640000, 14; 26643600, 14.2; 26647200, 14.1; 26650800, 13.7; 26654400, 13.3; 26658000, 13.1; 26661600, 13.4; 26665200, 13.3; 26668800, 13.1; 26672400, 12.8; 26676000, 12.3; 26679600, 12.2; 26683200, 11.7; 26686800, 11.1; 26690400, 10.4; 26694000, 9.7; 26697600, 9.2; 26701200, 8.9; 26704800, 8.6; 26708400, 8.1; 26712000, 7.8; 26715600, 7.1; 26719200, 6.5; 26722800, 6.4; 26726400, 6.3; 26730000, 6.2; 26733600, 6.2; 26737200, 6.6; 26740800, 7.1; 26744400, 6.9; 26748000, 6.2; 26751600, 5.3; 26755200, 4.9; 26758800, 4.6; 26762400, 3.9; 26766000, 2.7; 26769600, 1.7; 26773200, 1.9; 26776800, 2.9; 26780400, 3.6; 26784000, 3.7; 26787600, 3.1; 26791200, 2.5; 26794800, 1.9; 26798400, 2.4; 26802000, 3.8; 26805600, 5.1; 26809200, 6.1; 26812800, 6.4; 26816400, 6.8; 26820000, 7.1; 26823600, 6.6; 26827200, 5.9; 26830800, 4.9; 26834400, 3.4; 26838000, 1.8; 26841600, 3.3; 26845200, 5.4; 26848800, 7.6; 26852400, 9; 26856000, 9.2; 26859600, 8.9; 26863200, 8.1; 26866800, 7.4; 26870400, 6.3; 26874000, 4.7; 26877600, 3.3; 26881200, 2.3; 26884800, 1.9; 26888400, 2.6; 26892000, 2.9; 26895600, 2.7; 26899200, 2.6; 26902800, 2.3; 26906400, 1.6; 26910000, 0.3; 26913600, 1.5; 26917200, 3; 26920800, 4.2; 26924400, 5.1; 26928000, 5.6; 26931600, 6; 26935200, 6.2; 26938800, 6.1; 26942400, 5.9; 26946000, 5.8; 26949600, 5.5; 26953200, 5.2; 26956800, 5.2; 26960400, 5.5; 26964000, 6.2; 26967600, 6.4; 26971200, 6.1; 26974800, 7.2; 26978400, 7.2; 26982000, 7; 26985600, 7.2; 26989200, 6.8; 26992800, 6.6; 26996400, 6.7; 27000000, 7.3; 27003600, 8; 27007200, 8.7; 27010800, 9.1; 27014400, 9.1; 27018000, 9; 27021600, 9; 27025200, 9.1; 27028800, 9.3; 27032400, 9.6; 27036000, 10.1; 27039600, 10.6; 27043200, 11.2; 27046800, 11.7; 27050400, 11.6; 27054000, 11.6; 27057600, 11.5; 27061200, 11.1; 27064800, 9.8; 27068400, 9.2; 27072000, 8.9; 27075600, 7.7; 27079200, 5.7; 27082800, 3.8; 27086400, 4; 27090000, 5.7; 27093600, 7.5; 27097200, 8.5; 27100800, 8.7; 27104400, 8; 27108000, 6.4; 27111600, 4.5; 27115200, 2.5; 27118800, 1.3; 27122400, 2.1; 27126000, 3.2; 27129600, 4; 27133200, 4.6; 27136800, 5.5; 27140400, 5.5; 27144000, 4.6; 27147600, 5.5; 27151200, 6; 27154800, 5.6; 27158400, 4.9; 27162000, 3.7; 27165600, 2.6; 27169200, 2.5; 27172800, 3.6; 27176400, 5.2; 27180000, 6.6; 27183600, 7.4; 27187200, 7.7; 27190800, 7.9; 27194400, 8.4; 27198000, 9.2; 27201600, 9.6; 27205200, 9.5; 27208800, 9.5; 27212400, 9.5; 27216000, 9.1; 27219600, 8.6; 27223200, 8; 27226800, 6.8; 27230400, 5.9; 27234000, 5.7; 27237600, 4.9; 27241200, 4.3; 27244800, 3.9; 27248400, 4.2; 27252000, 5; 27255600, 6.3; 27259200, 7.4; 27262800, 8.1; 27266400, 8.5; 27270000, 8.5; 27273600, 8.3; 27277200, 7.7; 27280800, 7.1; 27284400, 6.8; 27288000, 6.8; 27291600, 6.8; 27295200, 6.8; 27298800, 6.8; 27302400, 6.2; 27306000, 5.4; 27309600, 4.5; 27313200, 3.8; 27316800, 3.5; 27320400, 3.7; 27324000, 3.7; 27327600, 3.9; 27331200, 4.6; 27334800, 5.2; 27338400, 5; 27342000, 4; 27345600, 3.2; 27349200, 2.8; 27352800, 2.5; 27356400, 2.5; 27360000, 3; 27363600, 3.7; 27367200, 4.7; 27370800, 5.3; 27374400, 5; 27378000, 4.2; 27381600, 5.4; 27385200, 7; 27388800, 7.9; 27392400, 7.7; 27396000, 7.8; 27399600, 7.3; 27403200, 6.8; 27406800, 9.4; 27410400, 9.6; 27414000, 9; 27417600, 8.6; 27421200, 8.7; 27424800, 8.6; 27428400, 8.2; 27432000, 8.4; 27435600, 8.4; 27439200, 8.4; 27442800, 8.5; 27446400, 8.8; 27450000, 9; 27453600, 9; 27457200, 8.8; 27460800, 8.4; 27464400, 7.8; 27468000, 7.3; 27471600, 7; 27475200, 6.6; 27478800, 6.2; 27482400, 5.9; 27486000, 5.5; 27489600, 5.8; 27493200, 5.7; 27496800, 5; 27500400, 4.5; 27504000, 4.4; 27507600, 4.5; 27511200, 5; 27514800, 6.2; 27518400, 7.3; 27522000, 8; 27525600, 8.6; 27529200, 8.9; 27532800, 8.8; 27536400, 8.5; 27540000, 8.2; 27543600, 8; 27547200, 7.9; 27550800, 8.3; 27554400, 8.6; 27558000, 8.4; 27561600, 8; 27565200, 7.7; 27568800, 7.4; 27572400, 6.4; 27576000, 5.2; 27579600, 6.5; 27583200, 6.5; 27586800, 5.6; 27590400, 4.6; 27594000, 3.6; 27597600, 2.2; 27601200, 0.7; 27604800, 1.7; 27608400, 3.5; 27612000, 4.9; 27615600, 5.9; 27619200, 6.7; 27622800, 7.3; 27626400, 7.8; 27630000, 8.2; 27633600, 8.3; 27637200, 8.1; 27640800, 8; 27644400, 8.1; 27648000, 8.3; 27651600, 8.3; 27655200, 8.1; 27658800, 7.3; 27662400, 6.7; 27666000, 6; 27669600, 5; 27673200, 4.2; 27676800, 3.7; 27680400, 3.3; 27684000, 3.3; 27687600, 3.5; 27691200, 3.6; 27694800, 3.3; 27698400, 3; 27702000, 3; 27705600, 3.4; 27709200, 3.8; 27712800, 3.7; 27716400, 3.2; 27720000, 2.6; 27723600, 2.1; 27727200, 1.5; 27730800, 1.4; 27734400, 2.2; 27738000, 3.1; 27741600, 3.4; 27745200, 2.9; 27748800, 1.8; 27752400, 2.3; 27756000, 2.6; 27759600, 2.2; 27763200, 1.9; 27766800, 2; 27770400, 2.2; 27774000, 3; 27777600, 4.2; 27781200, 4.8; 27784800, 4.7; 27788400, 4.1; 27792000, 3.1; 27795600, 2; 27799200, 0.9; 27802800, 2.1; 27806400, 3.3; 27810000, 2.8; 27813600, 2.4; 27817200, 2.9; 27820800, 4.5; 27824400, 5.4; 27828000, 5.8; 27831600, 5.6; 27835200, 5.8; 27838800, 7.8; 27842400, 8.7; 27846000, 9; 27849600, 9.1; 27853200, 8.8; 27856800, 8.7; 27860400, 9.1; 27864000, 9.4; 27867600, 9.2; 27871200, 9.1; 27874800, 8.9; 27878400, 8.2; 27882000, 7.5; 27885600, 6.7; 27889200, 6.4; 27892800, 6.6; 27896400, 7.4; 27900000, 7.9; 27903600, 8.4; 27907200, 8.6; 27910800, 8.8; 27914400, 8.6; 27918000, 7.8; 27921600, 8.1; 27925200, 8.5; 27928800, 9.1; 27932400, 8.5; 27936000, 7.5; 27939600, 7.5; 27943200, 8.3; 27946800, 9.8; 27950400, 10.7; 27954000, 10.6; 27957600, 10.5; 27961200, 10.2; 27964800, 9.6; 27968400, 8.9; 27972000, 8.4; 27975600, 8.2; 27979200, 8.3; 27982800, 8.3; 27986400, 8.2; 27990000, 8.2; 27993600, 8.5; 27997200, 10.7; 28000800, 13; 28004400, 14; 28008000, 14.8; 28011600, 15.5; 28015200, 15.3; 28018800, 14.3; 28022400, 13.6; 28026000, 12.9; 28029600, 11.9; 28033200, 11; 28036800, 9.9; 28040400, 9.4; 28044000, 9.9; 28047600, 10.8; 28051200, 11.5; 28054800, 12.1; 28058400, 12.8; 28062000, 12.8; 28065600, 12.3; 28069200, 11.5; 28072800, 10.7; 28076400, 10.1; 28080000, 9.8; 28083600, 10.1; 28087200, 10.2; 28090800, 10.8; 28094400, 10.9; 28098000, 10.1; 28101600, 8.9; 28105200, 7.9; 28108800, 7.2; 28112400, 6.3; 28116000, 5.5; 28119600, 5.1; 28123200, 5.1; 28126800, 5; 28130400, 4.6; 28134000, 4.7; 28137600, 5.3; 28141200, 6.2; 28144800, 7.6; 28148400, 8.5; 28152000, 8.5; 28155600, 7.7; 28159200, 6.9; 28162800, 6.5; 28166400, 6.1; 28170000, 5.8; 28173600, 5.5; 28177200, 5.7; 28180800, 6.2; 28184400, 6.3; 28188000, 6.4; 28191600, 6.6; 28195200, 6.9; 28198800, 7.2; 28202400, 7.1; 28206000, 6.6; 28209600, 6.8; 28213200, 7.3; 28216800, 7.8; 28220400, 8; 28224000, 8; 28227600, 7.9; 28231200, 7.7; 28234800, 7.3; 28238400, 6.4; 28242000, 5.3; 28245600, 4.2; 28249200, 3.1; 28252800, 2.1; 28256400, 2.1; 28260000, 3.3; 28263600, 3.9; 28267200, 3.8; 28270800, 4; 28274400, 3.8; 28278000, 3.9; 28281600, 4.2; 28285200, 4.3; 28288800, 4.4; 28292400, 4.8; 28296000, 5.5; 28299600, 6.1; 28303200, 6.5; 28306800, 7; 28310400, 7.3; 28314000, 7.6; 28317600, 7.7; 28321200, 7.6; 28324800, 7.1; 28328400, 6.7; 28332000, 6.6; 28335600, 7; 28339200, 7.4; 28342800, 6.8; 28346400, 5.9; 28350000, 4.8; 28353600, 3.8; 28357200, 4.2; 28360800, 4.8; 28364400, 5.3; 28368000, 5.5; 28371600, 5.4; 28375200, 4.7; 28378800, 4.1; 28382400, 4.4; 28386000, 5.2; 28389600, 5.9; 28393200, 6.4; 28396800, 6.9; 28400400, 7.2; 28404000, 7.4; 28407600, 7.3; 28411200, 7.3; 28414800, 7.5; 28418400, 7.5; 28422000, 7.1; 28425600, 6.7; 28429200, 6.8; 28432800, 6.7; 28436400, 5.8; 28440000, 4.2; 28443600, 3.6; 28447200, 3.7; 28450800, 3.5; 28454400, 3.4; 28458000, 3.4; 28461600, 3.2; 28465200, 2.9; 28468800, 2.8; 28472400, 3.4; 28476000, 4.4; 28479600, 5.4; 28483200, 6.2; 28486800, 6.8; 28490400, 7.1; 28494000, 7.3; 28497600, 7.2; 28501200, 7; 28504800, 7; 28508400, 7.8; 28512000, 8.7; 28515600, 9.3; 28519200, 10.2; 28522800, 9.7; 28526400, 9.5; 28530000, 11.2; 28533600, 10.8; 28537200, 9.6; 28540800, 8.7; 28544400, 7.9; 28548000, 7; 28551600, 6.3; 28555200, 6; 28558800, 5.8; 28562400, 5.4; 28566000, 5; 28569600, 4.6; 28573200, 4.1; 28576800, 3.8; 28580400, 3.1; 28584000, 2.3; 28587600, 1.8; 28591200, 2; 28594800, 2.8; 28598400, 3.7; 28602000, 4.4; 28605600, 5; 28609200, 5; 28612800, 4.5; 28616400, 5.2; 28620000, 6.1; 28623600, 5.9; 28627200, 5.4; 28630800, 4.7; 28634400, 3.7; 28638000, 3; 28641600, 2.9; 28645200, 2.7; 28648800, 2.2; 28652400, 1.8; 28656000, 1.3; 28659600, 1; 28663200, 0.8; 28666800, 0.6; 28670400, 0.6; 28674000, 1.5; 28677600, 2.5; 28681200, 3.3; 28684800, 3.8; 28688400, 3.9; 28692000, 3.9; 28695600, 3.9; 28699200, 3.7; 28702800, 3.1; 28706400, 2.8; 28710000, 3.3; 28713600, 3.4; 28717200, 3.4; 28720800, 3.4; 28724400, 3.8; 28728000, 4.1; 28731600, 4.3; 28735200, 4.5; 28738800, 4.8; 28742400, 5.1; 28746000, 5.2; 28749600, 5.6; 28753200, 6.2; 28756800, 6.7; 28760400, 7.2; 28764000, 7.8; 28767600, 8.2; 28771200, 8.4; 28774800, 8.3; 28778400, 8.1; 28782000, 7.5; 28785600, 7.1; 28789200, 7.1; 28792800, 7.4; 28796400, 8.3; 28800000, 8.5; 28803600, 8.4; 28807200, 7.6; 28810800, 7.7; 28814400, 8; 28818000, 8.3; 28821600, 9; 28825200, 11.2; 28828800, 12.1; 28832400, 11; 28836000, 9.7; 28839600, 8.7; 28843200, 7.7; 28846800, 6.4; 28850400, 4.9; 28854000, 4; 28857600, 3.6; 28861200, 3.6; 28864800, 4.1; 28868400, 5.2; 28872000, 5.7; 28875600, 6.3; 28879200, 6.5; 28882800, 6; 28886400, 5.4; 28890000, 5; 28893600, 4.8; 28897200, 4.9; 28900800, 4.5; 28904400, 3.6; 28908000, 2.5; 28911600, 2; 28915200, 2.6; 28918800, 3.8; 28922400, 5; 28926000, 6.1; 28929600, 7.1; 28933200, 7.5; 28936800, 7.4; 28940400, 7; 28944000, 6.2; 28947600, 5.6; 28951200, 5.3; 28954800, 4.9; 28958400, 4; 28962000, 3.3; 28965600, 3.4; 28969200, 3.8; 28972800, 4.2; 28976400, 4.1; 28980000, 3.8; 28983600, 3.4; 28987200, 3.3; 28990800, 3.8; 28994400, 4.6; 28998000, 5.6; 29001600, 6.5; 29005200, 7.1; 29008800, 7.7; 29012400, 7.9; 29016000, 7.8; 29019600, 7.6; 29023200, 7.4; 29026800, 7.1; 29030400, 6.5; 29034000, 6.2; 29037600, 6.2; 29041200, 6.1; 29044800, 5.8; 29048400, 5.7; 29052000, 6.2; 29055600, 6.7; 29059200, 6.8; 29062800, 6.5; 29066400, 6; 29070000, 6.1; 29073600, 6.7; 29077200, 7; 29080800, 6.6; 29084400, 5.8; 29088000, 5; 29091600, 4.3; 29095200, 3.5; 29098800, 2.8; 29102400, 2.1; 29106000, 1.6; 29109600, 1.4; 29113200, 1.5; 29116800, 1.5; 29120400, 1.6; 29124000, 2.1; 29127600, 2.6; 29131200, 3; 29134800, 3.4; 29138400, 3.6; 29142000, 3.4; 29145600, 3; 29149200, 3.1; 29152800, 3.9; 29156400, 5.4; 29160000, 6.9; 29163600, 8.1; 29167200, 8.8; 29170800, 9; 29174400, 8.8; 29178000, 8.6; 29181600, 8.5; 29185200, 8.5; 29188800, 8.8; 29192400, 9.2; 29196000, 9.6; 29199600, 9.8; 29203200, 10.1; 29206800, 10.3; 29210400, 10.6; 29214000, 10.2; 29217600, 9; 29221200, 7.4; 29224800, 6.6; 29228400, 7.2; 29232000, 7.8; 29235600, 8.4; 29239200, 9.4; 29242800, 10; 29246400, 9.8; 29250000, 8.5; 29253600, 6.7; 29257200, 4.9; 29260800, 3.1; 29264400, 1.5; 29268000, 1.8; 29271600, 3.4; 29275200, 5.3; 29278800, 6.9; 29282400, 7.9; 29286000, 8.3; 29289600, 8.4; 29293200, 8.5; 29296800, 8.6; 29300400, 8; 29304000, 7.1; 29307600, 8; 29311200, 10.7; 29314800, 11; 29318400, 10.3; 29322000, 8.7; 29325600, 6.9; 29329200, 5.7; 29332800, 4.5; 29336400, 3.3; 29340000, 2.3; 29343600, 1.7; 29347200, 1.6; 29350800, 1.2; 29354400, 0.8; 29358000, 1.1; 29361600, 2.2; 29365200, 3.2; 29368800, 4; 29372400, 4.8; 29376000, 5.5; 29379600, 5.6; 29383200, 5.4; 29386800, 5.1; 29390400, 4.4; 29394000, 3.1; 29397600, 2; 29401200, 1.7; 29404800, 1.3; 29408400, 1; 29412000, 1.6; 29415600, 2.3; 29419200, 2.9; 29422800, 3.9; 29426400, 5.3; 29430000, 6.9; 29433600, 8.3; 29437200, 9.3; 29440800, 9.6; 29444400, 9.3; 29448000, 8.8; 29451600, 9; 29455200, 9.1; 29458800, 8.8; 29462400, 8.2; 29466000, 8.1; 29469600, 7.6; 29473200, 6.6; 29476800, 6; 29480400, 7.7; 29484000, 10.3; 29487600, 12.7; 29491200, 13.1; 29494800, 13.1; 29498400, 12.6; 29502000, 11.6; 29505600, 11.2; 29509200, 11.1; 29512800, 11; 29516400, 10.8; 29520000, 10.4; 29523600, 9.7; 29527200, 8.7; 29530800, 7.6; 29534400, 6.7; 29538000, 6.2; 29541600, 5.7; 29545200, 5.3; 29548800, 5.1; 29552400, 4.9; 29556000, 4.6; 29559600, 3.9; 29563200, 3.1; 29566800, 2.2; 29570400, 1.4; 29574000, 0.9; 29577600, 0.7; 29581200, 0.6; 29584800, 0.5; 29588400, 0.5; 29592000, 0.6; 29595600, 0.8; 29599200, 1; 29602800, 1; 29606400, 0.8; 29610000, 0.5; 29613600, 0.7; 29617200, 0.9; 29620800, 1.2; 29624400, 1.6; 29628000, 2.1; 29631600, 2.4; 29635200, 2.6; 29638800, 2.8; 29642400, 2.8; 29646000, 2.6; 29649600, 2.1; 29653200, 1.9; 29656800, 1.9; 29660400, 1.8; 29664000, 1.6; 29667600, 1.4; 29671200, 1.3; 29674800, 1.6; 29678400, 2.2; 29682000, 2.5; 29685600, 2.7; 29689200, 3.4; 29692800, 3.3; 29696400, 3.9; 29700000, 4.5; 29703600, 5.3; 29707200, 6; 29710800, 5.8; 29714400, 4.8; 29718000, 4.1; 29721600, 4.1; 29725200, 4.8; 29728800, 5.9; 29732400, 6.3; 29736000, 5.4; 29739600, 4.8; 29743200, 5.1; 29746800, 4.7; 29750400, 4.4; 29754000, 3.9; 29757600, 3.3; 29761200, 3.1; 29764800, 3.1; 29768400, 3.1; 29772000, 2.7; 29775600, 2.4; 29779200, 2.3; 29782800, 2; 29786400, 2.1; 29790000, 2.6; 29793600, 2.8; 29797200, 2.7; 29800800, 2.5; 29804400, 2.3; 29808000, 1.9; 29811600, 1.5; 29815200, 2; 29818800, 3; 29822400, 3.4; 29826000, 3.3; 29829600, 3.5; 29833200, 4.1; 29836800, 4.9; 29840400, 5.5; 29844000, 6.2; 29847600, 7.3; 29851200, 8.8; 29854800, 9.9; 29858400, 10.2; 29862000, 10.5; 29865600, 10.9; 29869200, 11.1; 29872800, 11; 29876400, 11; 29880000, 10.9; 29883600, 10.9; 29887200, 10.9; 29890800, 10.8; 29894400, 10.8; 29898000, 10.5; 29901600, 10.2; 29905200, 9.6; 29908800, 9.3; 29912400, 8.4; 29916000, 7.7; 29919600, 7.3; 29923200, 7; 29926800, 6.8; 29930400, 6.5; 29934000, 6.3; 29937600, 6.2; 29941200, 5.7; 29944800, 5.3; 29948400, 5.3; 29952000, 5.7; 29955600, 6.8; 29959200, 8; 29962800, 8.6; 29966400, 8.7; 29970000, 8.7; 29973600, 8.7; 29977200, 8.6; 29980800, 8.6; 29984400, 8.5; 29988000, 8.1; 29991600, 7.1; 29995200, 6.1; 29998800, 5.2; 30002400, 4.7; 30006000, 4.4; 30009600, 4.1; 30013200, 3.8; 30016800, 3.3; 30020400, 2.4; 30024000, 1.3; 30027600, 1.2; 30031200, 2.8; 30034800, 4.3; 30038400, 5.4; 30042000, 6.3; 30045600, 6.7; 30049200, 6.5; 30052800, 6.3; 30056400, 6.1; 30060000, 6.2; 30063600, 6.3; 30067200, 6.5; 30070800, 6.3; 30074400, 6.8; 30078000, 7.9; 30081600, 9.1; 30085200, 8.7; 30088800, 7.6; 30092400, 6.8; 30096000, 6.4; 30099600, 6.2; 30103200, 5.6; 30106800, 4.8; 30110400, 3.8; 30114000, 2.7; 30117600, 1.9; 30121200, 1.7; 30124800, 2.1; 30128400, 2.5; 30132000, 3.3; 30135600, 4.4; 30139200, 5.3; 30142800, 6; 30146400, 6.3; 30150000, 6.2; 30153600, 5.8; 30157200, 5.8; 30160800, 6.2; 30164400, 6.4; 30168000, 6; 30171600, 6.1; 30175200, 8.2; 30178800, 8.7; 30182400, 8.5; 30186000, 7.9; 30189600, 7.4; 30193200, 7.6; 30196800, 8.2; 30200400, 9; 30204000, 9.6; 30207600, 9.8; 30211200, 9.5; 30214800, 8.8; 30218400, 8.1; 30222000, 7.5; 30225600, 7.1; 30229200, 6.9; 30232800, 6.7; 30236400, 6.3; 30240000, 5.5; 30243600, 4.9; 30247200, 4.4; 30250800, 3.9; 30254400, 3.3; 30258000, 3.4; 30261600, 3.3; 30265200, 3; 30268800, 2.8; 30272400, 2.3; 30276000, 1.5; 30279600, 1.3; 30283200, 1.7; 30286800, 1.9; 30290400, 1.9; 30294000, 1.6; 30297600, 1.1; 30301200, 0.4; 30304800, 0.8; 30308400, 1.7; 30312000, 2.3; 30315600, 2.6; 30319200, 3; 30322800, 3.2; 30326400, 3.2; 30330000, 3; 30333600, 2.8; 30337200, 2.5; 30340800, 2.1; 30344400, 1.6; 30348000, 1.2; 30351600, 0.9; 30355200, 0.9; 30358800, 1.3; 30362400, 2.1; 30366000, 2.9; 30369600, 3.4; 30373200, 3.7; 30376800, 3.6; 30380400, 3.4; 30384000, 3.1; 30387600, 2.6; 30391200, 1.7; 30394800, 0.8; 30398400, 0.7; 30402000, 1; 30405600, 1.3; 30409200, 1.6; 30412800, 2; 30416400, 2.2; 30420000, 2.3; 30423600, 2.2; 30427200, 2; 30430800, 1.6; 30434400, 1.3; 30438000, 1; 30441600, 0.8; 30445200, 0.7; 30448800, 0.4; 30452400, 0.1; 30456000, 0.7; 30459600, 1.3; 30463200, 1.6; 30466800, 1.9; 30470400, 2; 30474000, 2.2; 30477600, 2.7; 30481200, 2.8; 30484800, 2.2; 30488400, 1.5; 30492000, 0.9; 30495600, 1.3; 30499200, 1.9; 30502800, 2.4; 30506400, 2.6; 30510000, 2.8; 30513600, 2.7; 30517200, 2.4; 30520800, 1.8; 30524400, 1; 30528000, 0.1; 30531600, 0.6; 30535200, 1.4; 30538800, 1.6; 30542400, 1.5; 30546000, 1.8; 30549600, 2.2; 30553200, 2.7; 30556800, 3.5; 30560400, 3.8; 30564000, 3.7; 30567600, 3.1; 30571200, 1.7; 30574800, 0.7; 30578400, 1.3; 30582000, 2; 30585600, 2.1; 30589200, 1.8; 30592800, 2.1; 30596400, 2.3; 30600000, 2.8; 30603600, 2.9; 30607200, 2.8; 30610800, 2.7; 30614400, 2.4; 30618000, 2.5; 30621600, 2.7; 30625200, 3.1; 30628800, 4; 30632400, 5; 30636000, 5.6; 30639600, 5.8; 30643200, 5.9; 30646800, 6.4; 30650400, 6.9; 30654000, 6.9; 30657600, 6.9; 30661200, 6.5; 30664800, 6.1; 30668400, 5.8; 30672000, 5.2; 30675600, 4.5; 30679200, 3.9; 30682800, 3.4; 30686400, 3.2; 30690000, 3.8; 30693600, 4.2; 30697200, 4.4; 30700800, 4.7; 30704400, 5.2; 30708000, 5.2; 30711600, 5.1; 30715200, 5.4; 30718800, 6; 30722400, 6.9; 30726000, 7.7; 30729600, 8.2; 30733200, 8.1; 30736800, 7.7; 30740400, 7.5; 30744000, 7.3; 30747600, 7.1; 30751200, 6.8; 30754800, 6.5; 30758400, 6.4; 30762000, 6.6; 30765600, 6.7; 30769200, 6.3; 30772800, 5.7; 30776400, 4.7; 30780000, 3.9; 30783600, 3.3; 30787200, 3; 30790800, 2.7; 30794400, 2.1; 30798000, 1.8; 30801600, 2.2; 30805200, 3; 30808800, 3.6; 30812400, 4; 30816000, 4.4; 30819600, 4.7; 30823200, 4.9; 30826800, 4.9; 30830400, 4.7; 30834000, 4.7; 30837600, 4.9; 30841200, 5.5; 30844800, 5.8; 30848400, 5.7; 30852000, 5.6; 30855600, 5.4; 30859200, 5.3; 30862800, 5.3; 30866400, 4.5; 30870000, 3.6; 30873600, 3; 30877200, 2.7; 30880800, 2.9; 30884400, 3.4; 30888000, 4.1; 30891600, 5.2; 30895200, 6.2; 30898800, 6.8; 30902400, 6.8; 30906000, 6.5; 30909600, 6; 30913200, 5.6; 30916800, 5.2; 30920400, 5.1; 30924000, 5.1; 30927600, 5; 30931200, 4.9; 30934800, 4.8; 30938400, 4.8; 30942000, 4.9; 30945600, 4.7; 30949200, 3.9; 30952800, 3.1; 30956400, 2.5; 30960000, 2.6; 30963600, 2.3; 30967200, 1.7; 30970800, 1.6; 30974400, 1.7; 30978000, 1.6; 30981600, 1.6; 30985200, 2; 30988800, 2.7; 30992400, 2.9; 30996000, 3.3; 30999600, 4.1; 31003200, 5.1; 31006800, 6.3; 31010400, 7.4; 31014000, 9.4; 31017600, 11.2; 31021200, 11.7; 31024800, 11.8; 31028400, 11.6; 31032000, 11.6; 31035600, 11; 31039200, 10.1; 31042800, 9; 31046400, 8.1; 31050000, 7.4; 31053600, 6.6; 31057200, 5.8; 31060800, 5.5; 31064400, 5.4; 31068000, 5.2; 31071600, 4.9; 31075200, 5; 31078800, 5.3; 31082400, 5.4; 31086000, 5.3; 31089600, 5.1; 31093200, 4.9; 31096800, 4.7; 31100400, 4.6; 31104000, 4.7; 31107600, 4.9; 31111200, 4.9; 31114800, 4.7; 31118400, 4.4; 31122000, 4.2; 31125600, 4.1; 31129200, 3.5; 31132800, 2.6; 31136400, 1.6; 31140000, 0.4; 31143600, 0.8; 31147200, 2; 31150800, 3.1; 31154400, 4.3; 31158000, 5.4; 31161600, 6.4; 31165200, 7.1; 31168800, 7.3; 31172400, 7.3; 31176000, 7.3; 31179600, 7.4; 31183200, 7.4; 31186800, 7.3; 31190400, 7.3; 31194000, 7.2; 31197600, 7; 31201200, 6.5; 31204800, 5.4; 31208400, 5.4; 31212000, 6; 31215600, 5.9; 31219200, 5.9; 31222800, 6.3; 31226400, 6.7; 31230000, 7.2; 31233600, 7.4; 31237200, 7.4; 31240800, 7.4; 31244400, 8.1; 31248000, 9; 31251600, 9.4; 31255200, 9.4; 31258800, 9.3; 31262400, 9.2; 31266000, 8.7; 31269600, 8.1; 31273200, 7.5; 31276800, 6.9; 31280400, 6.2; 31284000, 5.6; 31287600, 4.9; 31291200, 3.8; 31294800, 3.1; 31298400, 3; 31302000, 3.4; 31305600, 3.4; 31309200, 3.5; 31312800, 3.8; 31316400, 4.3; 31320000, 4.5; 31323600, 4.3; 31327200, 4.3; 31330800, 4.6; 31334400, 4.9; 31338000, 5.3; 31341600, 5.7; 31345200, 5.9; 31348800, 5.8; 31352400, 5.5; 31356000, 5.4; 31359600, 5.3; 31363200, 5.1; 31366800, 4.8; 31370400, 4.8; 31374000, 4.8; 31377600, 4.4; 31381200, 4.1; 31384800, 4.5; 31388400, 4.8; 31392000, 4.9; 31395600, 4.7; 31399200, 4; 31402800, 3.4; 31406400, 3.1; 31410000, 3.3; 31413600, 3.6; 31417200, 4.1; 31420800, 4.6; 31424400, 5.1; 31428000, 5.1; 31431600, 4.5; 31435200, 3.7; 31438800, 2.9; 31442400, 2.3; 31446000, 2; 31449600, 1.8; 31453200, 2; 31456800, 2.7; 31460400, 3.2; 31464000, 3.2; 31467600, 2.7; 31471200, 2.5; 31474800, 2.3; 31478400, 2.2; 31482000, 2.3; 31485600, 2.2; 31489200, 2.2; 31492800, 2.7; 31496400, 3.2; 31500000, 3.5; 31503600, 3.8; 31507200, 4.5; 31510800, 5.2; 31514400, 5.5; 31518000, 5.4; 31521600, 4.9; 31525200, 4.2; 31528800, 3.6; 31532400, 3]) annotation(
      Placement(transformation(extent = {{-20, -10}, {0, 10}})));
    Modelica.Blocks.Interfaces.RealOutput airPressure(unit = "Pa") "大气压力" annotation(
      Placement(transformation(extent = {{100, -50}, {120, -30}})));
  equation
    connect(Tamb.y, T_amb) annotation(
      Line(points = {{1, 40}, {110, 40}}, color = {0, 0, 127}));
    connect(speedTable.y[1], windSpeed) annotation(
      Line(points = {{1, 0}, {110, 0}}, color = {0, 0, 127}));
    connect(airPressure, contPressure.y) annotation(
      Line(points = {{110, -40}, {1, -40}}, color = {0, 0, 127}));
    annotation(
      Diagram(coordinateSystem(preserveAspectRatio = false)));
  end Environment;
  annotation(
    version = "2.0",
    uses(Modelica(version = "3.2.3")));
end ENN;