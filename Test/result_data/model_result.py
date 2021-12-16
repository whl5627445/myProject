test_getgraphicsdata_res = [
  [
    {
      "visible": "true",
      "originalPoint": "0.0,0.0",
      "rotation": "0",
      "type": "Line",
      "points": [
        "-39,50",
        "-22,50"
      ],
      "color": "0,0,127",
      "linePattern": "LinePattern.Solid",
      "lineThickness": "0.25",
      "arrow": "Arrow.None,Arrow.None",
      "arrowSize": "3",
      "smooth": "Smooth.None",
      "connectionfrom_original_name": "step.y",
      "connectionto_original_name": "CriticalDamping.u",
      "connectionfrom": "step.y",
      "connectionto": "CriticalDamping.u"
    },
    {
      "visible": "true",
      "originalPoint": "0.0,0.0",
      "rotation": "0",
      "type": "Line",
      "points": [
        "-39,50",
        "-32,50",
        "-32,10",
        "-22,10"
      ],
      "color": "0,0,127",
      "linePattern": "LinePattern.Solid",
      "lineThickness": "0.25",
      "arrow": "Arrow.None,Arrow.None",
      "arrowSize": "3",
      "smooth": "Smooth.None",
      "connectionfrom_original_name": "step.y",
      "connectionto_original_name": "Bessel.u",
      "connectionfrom": "step.y",
      "connectionto": "Bessel.u"
    },
    {
      "visible": "true",
      "originalPoint": "0.0,0.0",
      "rotation": "0",
      "type": "Line",
      "points": [
        "-22,-30",
        "-32,-30",
        "-32,50",
        "-39,50"
      ],
      "color": "0,0,127",
      "linePattern": "LinePattern.Solid",
      "lineThickness": "0.25",
      "arrow": "Arrow.None,Arrow.None",
      "arrowSize": "3",
      "smooth": "Smooth.None",
      "connectionfrom_original_name": "Butterworth.u",
      "connectionto_original_name": "step.y",
      "connectionfrom": "Butterworth.u",
      "connectionto": "step.y"
    },
    {
      "visible": "true",
      "originalPoint": "0.0,0.0",
      "rotation": "0",
      "type": "Line",
      "points": [
        "-22,-70",
        "-32,-70",
        "-32,50",
        "-39,50"
      ],
      "color": "0,0,127",
      "linePattern": "LinePattern.Solid",
      "lineThickness": "0.25",
      "arrow": "Arrow.None,Arrow.None",
      "arrowSize": "3",
      "smooth": "Smooth.None",
      "connectionfrom_original_name": "ChebyshevI.u",
      "connectionto_original_name": "step.y",
      "connectionfrom": "ChebyshevI.u",
      "connectionto": "step.y"
    }
  ],
  [
    {
      "type": "Transformation",
      "ID": "5",
      "original_name": "step",
      "name": "step",
      "parent": "",
      "classname": "Modelica.Blocks.Sources.Step",
      "visible": "true",
      "rotateAngle": "0",
      "originDiagram": "-,-",
      "extent1Diagram": "-60.0,40.0",
      "extent2Diagram": "-40.0,60.0",
      "rotation": "0",
      "output_type": "",
      "inputOutputs": [
        {
          "type": "Transformation",
          "ID": "0",
          "original_name": "y",
          "name": "y",
          "parent": "step",
          "classname": "Modelica.Blocks.Interfaces.RealOutput",
          "visible": "true",
          "rotateAngle": "0",
          "originDiagram": "-,-",
          "extent1Diagram": "100.0,-10.0",
          "extent2Diagram": "120.0,10.0",
          "rotation": "0",
          "output_type": "",
          "inputOutputs": [],
          "subShapes": [
            {
              "visible": "true",
              "originalPoint": "0.0,0.0",
              "rotation": "0",
              "type": "Polygon",
              "color": "0,0,127",
              "fillColor": "255,255,255",
              "linePattern": "LinePattern.Solid",
              "fillPattern": "FillPattern.Solid",
              "lineThickness": "0.25",
              "polygonPoints": [
                "-100.0,100.0",
                "100.0,0.0",
                "-100.0,-100.0"
              ],
              "smooth": "Smooth.None"
            }
          ]
        }
      ],
      "subShapes": [
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-80,68",
            "-80,-80"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "-80,90",
            "-88,68",
            "-72,68",
            "-80,90"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-90,-70",
            "82,-70"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "90,-70",
            "68,-62",
            "68,-78",
            "90,-70"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-80,-70",
            "0,-70",
            "0,50",
            "80,50"
          ],
          "color": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,0",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-150,-150",
            "150,-110"
          ],
          "originalTextString": "startTime=%startTime",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Rectangle",
          "color": "0,0,127",
          "fillColor": "255,255,255",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "borderPattern": "BorderPattern.None",
          "extentsPoints": [
            "-100,-100",
            "100,100"
          ],
          "radius": "0"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,255",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-150,150",
            "150,110"
          ],
          "originalTextString": "%name",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        }
      ]
    },
    {
      "type": "Transformation",
      "ID": "6",
      "original_name": "CriticalDamping",
      "name": "CriticalDamping",
      "parent": "",
      "classname": "Modelica.Blocks.Continuous.Filter",
      "visible": "true",
      "rotateAngle": "0",
      "originDiagram": "-,-",
      "extent1Diagram": "-20.0,40.0",
      "extent2Diagram": "0.0,60.0",
      "rotation": "0",
      "output_type": "",
      "inputOutputs": [
        {
          "type": "Transformation",
          "ID": "1",
          "original_name": "u",
          "name": "u",
          "parent": "CriticalDamping",
          "classname": "Modelica.Blocks.Interfaces.RealInput",
          "visible": "true",
          "rotateAngle": "0",
          "originDiagram": "-,-",
          "extent1Diagram": "-140.0,-20.0",
          "extent2Diagram": "-100.0,20.0",
          "rotation": "0",
          "output_type": "",
          "inputOutputs": [],
          "subShapes": [
            {
              "visible": "true",
              "originalPoint": "0.0,0.0",
              "rotation": "0",
              "type": "Polygon",
              "color": "0,0,127",
              "fillColor": "0,0,127",
              "linePattern": "LinePattern.Solid",
              "fillPattern": "FillPattern.Solid",
              "lineThickness": "0.25",
              "polygonPoints": [
                "-100.0,100.0",
                "100.0,0.0",
                "-100.0,-100.0"
              ],
              "smooth": "Smooth.None"
            }
          ]
        },
        {
          "type": "Transformation",
          "ID": "2",
          "original_name": "y",
          "name": "y",
          "parent": "CriticalDamping",
          "classname": "Modelica.Blocks.Interfaces.RealOutput",
          "visible": "true",
          "rotateAngle": "0",
          "originDiagram": "-,-",
          "extent1Diagram": "100.0,-10.0",
          "extent2Diagram": "120.0,10.0",
          "rotation": "0",
          "output_type": "",
          "inputOutputs": [],
          "subShapes": [
            {
              "visible": "true",
              "originalPoint": "0.0,0.0",
              "rotation": "0",
              "type": "Polygon",
              "color": "0,0,127",
              "fillColor": "255,255,255",
              "linePattern": "LinePattern.Solid",
              "fillPattern": "FillPattern.Solid",
              "lineThickness": "0.25",
              "polygonPoints": [
                "-100.0,100.0",
                "100.0,0.0",
                "-100.0,-100.0"
              ],
              "smooth": "Smooth.None"
            }
          ]
        }
      ],
      "subShapes": [
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-80.0,80.0",
            "-80.0,-88.0"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "-80.0,92.0",
            "-88.0,70.0",
            "-72.0,70.0",
            "-80.0,92.0"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-90.0,-78.0",
            "82.0,-78.0"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "90.0,-78.0",
            "68.0,-70.0",
            "68.0,-86.0",
            "90.0,-78.0"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "192,192,192",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-66.0,52.0",
            "88.0,90.0"
          ],
          "originalTextString": "%order",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,0",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-138.0,-140.0",
            "162.0,-110.0"
          ],
          "originalTextString": "f_cut=%f_cut",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Rectangle",
          "color": "160,160,164",
          "fillColor": "255,255,255",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Backward",
          "lineThickness": "0.25",
          "borderPattern": "BorderPattern.None",
          "extentsPoints": [
            "-80.0,-78.0",
            "22.0,10.0"
          ],
          "radius": "0"
        },
        {
          "visible": "true",
          "originalPoint": "3.333,-6.667",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-83.333,34.667",
            "24.667,34.667",
            "42.667,-71.333"
          ],
          "color": "0,0,127",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.Bezier"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Rectangle",
          "color": "0,0,127",
          "fillColor": "255,255,255",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "borderPattern": "BorderPattern.None",
          "extentsPoints": [
            "-100,-100",
            "100,100"
          ],
          "radius": "0"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,255",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-150,150",
            "150,110"
          ],
          "originalTextString": "%name",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        }
      ]
    },
    {
      "type": "Transformation",
      "ID": "7",
      "original_name": "Bessel",
      "name": "Bessel",
      "parent": "",
      "classname": "Modelica.Blocks.Continuous.Filter",
      "visible": "true",
      "rotateAngle": "0",
      "originDiagram": "-,-",
      "extent1Diagram": "-20.0,0.0",
      "extent2Diagram": "0.0,20.0",
      "rotation": "0",
      "output_type": "",
      "inputOutputs": [
        {
          "type": "Transformation",
          "ID": "1",
          "original_name": "u",
          "name": "u",
          "parent": "Bessel",
          "classname": "Modelica.Blocks.Interfaces.RealInput",
          "visible": "true",
          "rotateAngle": "0",
          "originDiagram": "-,-",
          "extent1Diagram": "-140.0,-20.0",
          "extent2Diagram": "-100.0,20.0",
          "rotation": "0",
          "output_type": "",
          "inputOutputs": [],
          "subShapes": [
            {
              "visible": "true",
              "originalPoint": "0.0,0.0",
              "rotation": "0",
              "type": "Polygon",
              "color": "0,0,127",
              "fillColor": "0,0,127",
              "linePattern": "LinePattern.Solid",
              "fillPattern": "FillPattern.Solid",
              "lineThickness": "0.25",
              "polygonPoints": [
                "-100.0,100.0",
                "100.0,0.0",
                "-100.0,-100.0"
              ],
              "smooth": "Smooth.None"
            }
          ]
        },
        {
          "type": "Transformation",
          "ID": "2",
          "original_name": "y",
          "name": "y",
          "parent": "Bessel",
          "classname": "Modelica.Blocks.Interfaces.RealOutput",
          "visible": "true",
          "rotateAngle": "0",
          "originDiagram": "-,-",
          "extent1Diagram": "100.0,-10.0",
          "extent2Diagram": "120.0,10.0",
          "rotation": "0",
          "output_type": "",
          "inputOutputs": [],
          "subShapes": [
            {
              "visible": "true",
              "originalPoint": "0.0,0.0",
              "rotation": "0",
              "type": "Polygon",
              "color": "0,0,127",
              "fillColor": "255,255,255",
              "linePattern": "LinePattern.Solid",
              "fillPattern": "FillPattern.Solid",
              "lineThickness": "0.25",
              "polygonPoints": [
                "-100.0,100.0",
                "100.0,0.0",
                "-100.0,-100.0"
              ],
              "smooth": "Smooth.None"
            }
          ]
        }
      ],
      "subShapes": [
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-80.0,80.0",
            "-80.0,-88.0"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "-80.0,92.0",
            "-88.0,70.0",
            "-72.0,70.0",
            "-80.0,92.0"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-90.0,-78.0",
            "82.0,-78.0"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "90.0,-78.0",
            "68.0,-70.0",
            "68.0,-86.0",
            "90.0,-78.0"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "192,192,192",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-66.0,52.0",
            "88.0,90.0"
          ],
          "originalTextString": "%order",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,0",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-138.0,-140.0",
            "162.0,-110.0"
          ],
          "originalTextString": "f_cut=%f_cut",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Rectangle",
          "color": "160,160,164",
          "fillColor": "255,255,255",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Backward",
          "lineThickness": "0.25",
          "borderPattern": "BorderPattern.None",
          "extentsPoints": [
            "-80.0,-78.0",
            "22.0,10.0"
          ],
          "radius": "0"
        },
        {
          "visible": "true",
          "originalPoint": "3.333,-6.667",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-83.333,34.667",
            "24.667,34.667",
            "42.667,-71.333"
          ],
          "color": "0,0,127",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.Bezier"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Rectangle",
          "color": "0,0,127",
          "fillColor": "255,255,255",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "borderPattern": "BorderPattern.None",
          "extentsPoints": [
            "-100,-100",
            "100,100"
          ],
          "radius": "0"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,255",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-150,150",
            "150,110"
          ],
          "originalTextString": "%name",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        }
      ]
    },
    {
      "type": "Transformation",
      "ID": "8",
      "original_name": "Butterworth",
      "name": "Butterworth",
      "parent": "",
      "classname": "Modelica.Blocks.Continuous.Filter",
      "visible": "true",
      "rotateAngle": "0",
      "originDiagram": "-,-",
      "extent1Diagram": "-20.0,-40.0",
      "extent2Diagram": "0.0,-20.0",
      "rotation": "0",
      "output_type": "",
      "inputOutputs": [
        {
          "type": "Transformation",
          "ID": "1",
          "original_name": "u",
          "name": "u",
          "parent": "Butterworth",
          "classname": "Modelica.Blocks.Interfaces.RealInput",
          "visible": "true",
          "rotateAngle": "0",
          "originDiagram": "-,-",
          "extent1Diagram": "-140.0,-20.0",
          "extent2Diagram": "-100.0,20.0",
          "rotation": "0",
          "output_type": "",
          "inputOutputs": [],
          "subShapes": [
            {
              "visible": "true",
              "originalPoint": "0.0,0.0",
              "rotation": "0",
              "type": "Polygon",
              "color": "0,0,127",
              "fillColor": "0,0,127",
              "linePattern": "LinePattern.Solid",
              "fillPattern": "FillPattern.Solid",
              "lineThickness": "0.25",
              "polygonPoints": [
                "-100.0,100.0",
                "100.0,0.0",
                "-100.0,-100.0"
              ],
              "smooth": "Smooth.None"
            }
          ]
        },
        {
          "type": "Transformation",
          "ID": "2",
          "original_name": "y",
          "name": "y",
          "parent": "Butterworth",
          "classname": "Modelica.Blocks.Interfaces.RealOutput",
          "visible": "true",
          "rotateAngle": "0",
          "originDiagram": "-,-",
          "extent1Diagram": "100.0,-10.0",
          "extent2Diagram": "120.0,10.0",
          "rotation": "0",
          "output_type": "",
          "inputOutputs": [],
          "subShapes": [
            {
              "visible": "true",
              "originalPoint": "0.0,0.0",
              "rotation": "0",
              "type": "Polygon",
              "color": "0,0,127",
              "fillColor": "255,255,255",
              "linePattern": "LinePattern.Solid",
              "fillPattern": "FillPattern.Solid",
              "lineThickness": "0.25",
              "polygonPoints": [
                "-100.0,100.0",
                "100.0,0.0",
                "-100.0,-100.0"
              ],
              "smooth": "Smooth.None"
            }
          ]
        }
      ],
      "subShapes": [
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-80.0,80.0",
            "-80.0,-88.0"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "-80.0,92.0",
            "-88.0,70.0",
            "-72.0,70.0",
            "-80.0,92.0"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-90.0,-78.0",
            "82.0,-78.0"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "90.0,-78.0",
            "68.0,-70.0",
            "68.0,-86.0",
            "90.0,-78.0"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "192,192,192",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-66.0,52.0",
            "88.0,90.0"
          ],
          "originalTextString": "%order",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,0",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-138.0,-140.0",
            "162.0,-110.0"
          ],
          "originalTextString": "f_cut=%f_cut",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Rectangle",
          "color": "160,160,164",
          "fillColor": "255,255,255",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Backward",
          "lineThickness": "0.25",
          "borderPattern": "BorderPattern.None",
          "extentsPoints": [
            "-80.0,-78.0",
            "22.0,10.0"
          ],
          "radius": "0"
        },
        {
          "visible": "true",
          "originalPoint": "3.333,-6.667",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-83.333,34.667",
            "24.667,34.667",
            "42.667,-71.333"
          ],
          "color": "0,0,127",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.Bezier"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Rectangle",
          "color": "0,0,127",
          "fillColor": "255,255,255",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "borderPattern": "BorderPattern.None",
          "extentsPoints": [
            "-100,-100",
            "100,100"
          ],
          "radius": "0"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,255",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-150,150",
            "150,110"
          ],
          "originalTextString": "%name",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        }
      ]
    },
    {
      "type": "Transformation",
      "ID": "9",
      "original_name": "ChebyshevI",
      "name": "ChebyshevI",
      "parent": "",
      "classname": "Modelica.Blocks.Continuous.Filter",
      "visible": "true",
      "rotateAngle": "0",
      "originDiagram": "-,-",
      "extent1Diagram": "-20.0,-80.0",
      "extent2Diagram": "0.0,-60.0",
      "rotation": "0",
      "output_type": "",
      "inputOutputs": [
        {
          "type": "Transformation",
          "ID": "1",
          "original_name": "u",
          "name": "u",
          "parent": "ChebyshevI",
          "classname": "Modelica.Blocks.Interfaces.RealInput",
          "visible": "true",
          "rotateAngle": "0",
          "originDiagram": "-,-",
          "extent1Diagram": "-140.0,-20.0",
          "extent2Diagram": "-100.0,20.0",
          "rotation": "0",
          "output_type": "",
          "inputOutputs": [],
          "subShapes": [
            {
              "visible": "true",
              "originalPoint": "0.0,0.0",
              "rotation": "0",
              "type": "Polygon",
              "color": "0,0,127",
              "fillColor": "0,0,127",
              "linePattern": "LinePattern.Solid",
              "fillPattern": "FillPattern.Solid",
              "lineThickness": "0.25",
              "polygonPoints": [
                "-100.0,100.0",
                "100.0,0.0",
                "-100.0,-100.0"
              ],
              "smooth": "Smooth.None"
            }
          ]
        },
        {
          "type": "Transformation",
          "ID": "2",
          "original_name": "y",
          "name": "y",
          "parent": "ChebyshevI",
          "classname": "Modelica.Blocks.Interfaces.RealOutput",
          "visible": "true",
          "rotateAngle": "0",
          "originDiagram": "-,-",
          "extent1Diagram": "100.0,-10.0",
          "extent2Diagram": "120.0,10.0",
          "rotation": "0",
          "output_type": "",
          "inputOutputs": [],
          "subShapes": [
            {
              "visible": "true",
              "originalPoint": "0.0,0.0",
              "rotation": "0",
              "type": "Polygon",
              "color": "0,0,127",
              "fillColor": "255,255,255",
              "linePattern": "LinePattern.Solid",
              "fillPattern": "FillPattern.Solid",
              "lineThickness": "0.25",
              "polygonPoints": [
                "-100.0,100.0",
                "100.0,0.0",
                "-100.0,-100.0"
              ],
              "smooth": "Smooth.None"
            }
          ]
        }
      ],
      "subShapes": [
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-80.0,80.0",
            "-80.0,-88.0"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "-80.0,92.0",
            "-88.0,70.0",
            "-72.0,70.0",
            "-80.0,92.0"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-90.0,-78.0",
            "82.0,-78.0"
          ],
          "color": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Polygon",
          "color": "192,192,192",
          "fillColor": "192,192,192",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "polygonPoints": [
            "90.0,-78.0",
            "68.0,-70.0",
            "68.0,-86.0",
            "90.0,-78.0"
          ],
          "smooth": "Smooth.None"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "192,192,192",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-66.0,52.0",
            "88.0,90.0"
          ],
          "originalTextString": "%order",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,0",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-138.0,-140.0",
            "162.0,-110.0"
          ],
          "originalTextString": "f_cut=%f_cut",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Rectangle",
          "color": "160,160,164",
          "fillColor": "255,255,255",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Backward",
          "lineThickness": "0.25",
          "borderPattern": "BorderPattern.None",
          "extentsPoints": [
            "-80.0,-78.0",
            "22.0,10.0"
          ],
          "radius": "0"
        },
        {
          "visible": "true",
          "originalPoint": "3.333,-6.667",
          "rotation": "0",
          "type": "Line",
          "points": [
            "-83.333,34.667",
            "24.667,34.667",
            "42.667,-71.333"
          ],
          "color": "0,0,127",
          "linePattern": "LinePattern.Solid",
          "lineThickness": "0.25",
          "arrow": "Arrow.None,Arrow.None",
          "arrowSize": "3",
          "smooth": "Smooth.Bezier"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Rectangle",
          "color": "0,0,127",
          "fillColor": "255,255,255",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.Solid",
          "lineThickness": "0.25",
          "borderPattern": "BorderPattern.None",
          "extentsPoints": [
            "-100,-100",
            "100,100"
          ],
          "radius": "0"
        },
        {
          "visible": "true",
          "originalPoint": "0.0,0.0",
          "rotation": "0",
          "type": "Text",
          "color": "0,0,255",
          "fillColor": "0,0,0",
          "linePattern": "LinePattern.Solid",
          "fillPattern": "FillPattern.None",
          "lineThickness": "0.25",
          "extentsPoints": [
            "-150,150",
            "150,110"
          ],
          "originalTextString": "%name",
          "fontSize": "0",
          "textColor": "-1,-1,-1",
          "fontName": "",
          "textStyles": [
            ""
          ],
          "horizontalAlignment": "TextAlignment.Center"
        }
      ]
    }
  ]
]
test_getmodelcode_res =[
                          "model Filter \"Demonstrates the Continuous.Filter block with various options\"\n  extends Modelica.Icons.Example;\n  parameter Integer order = 3;\n  parameter Modelica.SIunits.Frequency f_cut = 2;\n  parameter Modelica.Blocks.Types.FilterType filterType = Modelica.Blocks.Types.FilterType.LowPass \"Type of filter (LowPass/HighPass)\";\n  parameter Modelica.Blocks.Types.Init init = Modelica.Blocks.Types.Init.SteadyState \"Type of initialization (no init/steady state/initial state/initial output)\";\n  parameter Boolean normalized = true;\n  Modelica.Blocks.Sources.Step step(startTime = 0.1, offset = 0.1) annotation(\n    Placement(transformation(extent = {{-60, 40}, {-40, 60}})));\n  Modelica.Blocks.Continuous.Filter CriticalDamping(analogFilter = Modelica.Blocks.Types.AnalogFilter.CriticalDamping, normalized = normalized, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, 40}, {0, 60}})));\n  Modelica.Blocks.Continuous.Filter Bessel(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Bessel, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, 0}, {0, 20}})));\n  Modelica.Blocks.Continuous.Filter Butterworth(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.Butterworth, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, -40}, {0, -20}})));\n  Modelica.Blocks.Continuous.Filter ChebyshevI(normalized = normalized, analogFilter = Modelica.Blocks.Types.AnalogFilter.ChebyshevI, init = init, filterType = filterType, order = order, f_cut = f_cut, f_min = 0.8 * f_cut) annotation(\n    Placement(transformation(extent = {{-20, -80}, {0, -60}})));\nequation\n  connect(step.y, CriticalDamping.u) annotation(\n    Line(points = {{-39, 50}, {-22, 50}}, color = {0, 0, 127}));\n  connect(step.y, Bessel.u) annotation(\n    Line(points = {{-39, 50}, {-32, 50}, {-32, 10}, {-22, 10}}, color = {0, 0, 127}));\n  connect(Butterworth.u, step.y) annotation(\n    Line(points = {{-22, -30}, {-32, -30}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n  connect(ChebyshevI.u, step.y) annotation(\n    Line(points = {{-22, -70}, {-32, -70}, {-32, 50}, {-39, 50}}, color = {0, 0, 127}));\n  annotation(\n    experiment(StopTime = 0.9),\n    Documentation(info = \"<html>\n\n<p>\nThis example demonstrates various options of the\n<a href=\\\"modelica://Modelica.Blocks.Continuous.Filter\\\">Filter</a> block.\nA step input starts at 0.1 s with an offset of 0.1, in order to demonstrate\nthe initialization options. This step input drives 4 filter blocks that\nhave identical parameters, with the only exception of the used analog filter type\n(CriticalDamping, Bessel, Butterworth, Chebyshev of type I). All the main options\ncan be set via parameters and are then applied to all the 4 filters.\nThe default setting uses low pass filters of order 3 with a cut-off frequency of\n2 Hz resulting in the following outputs:\n</p>\n\n<img src=\\\"modelica://Modelica/Resources/Images/Blocks/Filter1.png\\\"\n     alt=\\\"Filter1.png\\\">\n</html>\"));\nend Filter;"
                        ]
test_getmodelparameters_res = [
  {
    "tab": "General",
    "type": "Normal",
    "group": "Parameters",
    "name": "height",
    "comment": "Height of step",
    "value": "",
    "defaultvalue": "1",
    "unit": ""
  },
  {
    "tab": "General",
    "type": "Normal",
    "group": "Parameters",
    "name": "offset",
    "comment": "Offset of output signal y",
    "value": "0.1",
    "defaultvalue": "0",
    "unit": ""
  },
  {
    "tab": "General",
    "type": "Normal",
    "group": "Parameters",
    "name": "startTime",
    "comment": "Output y = offset for time < startTime",
    "value": "0.1",
    "defaultvalue": "0",
    "unit": [
      "s"
    ]
  }
]
test_setmodelparameters_res = {
  "model_name": "Modelica.Blocks.Examples.Filter",
  "parameter_value": {
    "step.height": "",
    "step.offset": "0.2",
    "step.startTime": "0.1"
  }
}
