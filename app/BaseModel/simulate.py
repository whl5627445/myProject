# -- coding: utf-8 --
from typing import Optional
from pydantic import BaseModel


class ModelSimulateModel(BaseModel):
    model_name: str
    simulate_type: Optional[str] = "OM"
    start_time: Optional[float] = 0.0
    stop_time: Optional[float]
    number_of_intervals: Optional[int] = 500
    # interval: Optional[float]
    tolerance: Optional[float] = 0.000001
    # method: Optional[str]
    # options: Optional[str]
    # outputFormat: Optional[str]
    # variableFilter: Optional[str]
    # cflags: Optional[str] = None
    # simflags: Optional[str]


class SetComponentModifierValueModel(BaseModel):
    model_name: str
    parameter_value: dict


class SetComponentPropertiesModel(BaseModel):
    model_name: str
    component_name: str
    final: str
    protected: str
    replaceable: str
    variabilty: str
    inner: str
    outer: str
    causality: str


class CopyClassModel(BaseModel):
    parent_name: str
    package_name: str
    class_name: str
    copied_class_name: str


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
    new_component_name: str
    old_component_name: str
    model_name_all: str
    origin: str = "0,0"
    extent: list = ["-10,-10", "10,10"]
    rotation: str = "0"


class DeleteComponentModel(BaseModel):
    package_name: str
    package_id: str
    component_name: str
    model_name_all: str


class UpdateConnectionAnnotationModel(BaseModel):
    package_name: str
    package_id: str
    model_name_all: str
    connect_start: str
    connect_end: str
    color: str = "0,0,127"
    line_points: list = []


class DeleteConnectionModel(BaseModel):
    package_name: str
    package_id: str
    model_name_all: str
    connect_start: str
    connect_end: str


class ExperimentCreateModel(BaseModel):
    package_id: int
    model_name: str
    model_var_data: list
    simulate_var_data: list
    experiment_name: str


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
        "StartTime": 0,
        "StopTime": 4,
        "Tolerance": 1e-06,
        "Interval": 0.008
        }
