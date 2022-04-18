# -- coding: utf-8 --
from typing import Optional, List
from pydantic import BaseModel


class ModelSimulateModel(BaseModel):
    package_id: str
    model_name: str
    simulate_type: Optional[str] = "OM"
    start_time: Optional[str] = "0.0"
    stop_time: Optional[str] = "4.0"
    number_of_intervals: Optional[str] = "500"
    # interval: Optional[float]
    tolerance: Optional[str] = "0.000001"
    method: Optional[str] = None
    # options: Optional[str]
    # outputFormat: Optional[str]
    # variableFilter: Optional[str]
    # cflags: Optional[str] = None
    # simflags: Optional[str]


class SetComponentModifierValueModel(BaseModel):
    package_id: str
    model_name: str
    parameter_value: dict


class SetComponentPropertiesModel(BaseModel):
    package_id: str
    model_name: str
    old_component_name: str
    new_component_name: str
    final: str
    protected: str
    replaceable: str
    variabilty: str
    inner: str
    outer: str
    causality: str


class CopyClassModel(BaseModel):
    package_id: str
    parent_name: str
    package_name: str
    class_name: str
    copied_class_name: str


class GetComponentNameModel(BaseModel):
    # package_name: str
    package_id: str
    old_component_name: str
    model_name_all: str



class AddComponentModel(BaseModel):
    package_name: str
    package_id: str
    new_component_name: str
    old_component_name: str
    model_name_all: str
    origin: str = "0,0"
    extent: list = ["-10,-10", "10,10"]
    rotation: str = "0"


class UpdateComponentModel(BaseModel):
    package_name: str
    package_id: str
    component_name: str
    component_model_name: str
    model_name_all: str
    origin: str = "0,0"
    extent: list = ["10,10", "10,10"]
    rotation: str = "0"



class DeleteComponentModel(BaseModel):
    package_name: str
    package_id: str
    delete_list: Optional[list] = [
        {
            "delete_type": "",  # 可以有两种类型， 一种是组件类型，一种是连线类型， component, connector
            "component_name": "",
            "model_name_all": "",
            "connect_start": "",
            "connect_end": "",
            }
        ]



class UpdateConnectionAnnotationModel(BaseModel):
    package_name: str
    package_id: str
    model_name_all: str
    connect_start: str
    connect_end: str
    color: str = "0,0,127"
    line_points: list = []


class UpdateConnectionNamesModel(BaseModel):
    package_name: str
    package_id: str
    model_name_all: str
    from_name: str
    to_name: str
    from_name_new: str
    to_name_new: str


class DeleteConnectionModel(BaseModel):
    package_name: str
    package_id: str
    model_name_all: str
    connect_start: str
    connect_end: str


class ExperimentCreateModel(BaseModel):
    package_id: int
    model_name: str
    model_var_data: dict = {}
    simulate_var_data: dict = {}
    experiment_name: str = "test"


class DeletePackageModel(BaseModel):
    package_id: int
    package_name: str
    parent_name: str
    class_name: str


class SetSimulationOptionsModel(BaseModel):
    package_id: int
    package_name: str
    model_name: str
    experiment: Optional[dict] = {
        "startTime": 0,
        "stopTime": 4,
        "tolerance": 1e-06,
        "numberOfIntervals": 500,
        "interval": 0.008
        }


class ModelCodeSaveModel(BaseModel):
    package_id: int
    package_name: str


class FmuExportModel(BaseModel):
    package_id: int
    package_name: str
    model_name: str
    fmu_name: str
    download_local: bool = False
    storeResult: bool = False
    includeSource: bool = False
    fmiVersion: int = 2
    includeImage: int = 0
    fmiType: str = "all"


class SimulateResultExportModel(BaseModel):
    record_id: int
    export_type: str
    var_list: List = []


class SimulateALLResultModel(BaseModel):
    record_id: int
