from typing import Optional
from pydantic import BaseModel


class ModelSimulateModel(BaseModel):
    model_name: str
    s_type: Optional[str] = "OM"
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
