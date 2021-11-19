# -- coding: utf-8 --
from fastapi import FastAPI, Request, Response
from fastapi.responses import JSONResponse
from app.api.simulateview import router as simulate_view_router
from app.api.modelview import router as model_view_router
from app.api.upload_file import router as upload_file_router
from app.api.download import router as download_file_router
from starlette.middleware.sessions import SessionMiddleware
from starlette.middleware.authentication import AuthenticationMiddleware
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles

from starlette.authentication import (
    AuthenticationBackend, AuthenticationError, SimpleUser,
    AuthCredentials
)

import pydevd_pycharm
pydevd_pycharm.settrace('192.168.1.110', port=10086, stdoutToServer=True, stderrToServer=True)

class BasicAuthBackend(AuthenticationBackend):
    async def authenticate(self, request):
        username = request.headers.get("username", "tom")
        if not username:
            return
        return AuthCredentials(["simtek"]), SimpleUser(username)

app = FastAPI()
app.include_router(simulate_view_router)
app.include_router(model_view_router)
app.include_router(upload_file_router)
app.include_router(download_file_router)
app.add_middleware(AuthenticationMiddleware, backend=BasicAuthBackend())
app.add_middleware(SessionMiddleware, secret_key="simtek")

app.mount("/static", StaticFiles(directory="./"), name="static")


# app.add_middleware(
#     CORSMiddleware,
#     allow_origins=["*"],
#     allow_credentials=True,
#     allow_methods=["*"],
#     allow_headers=["*"],
# )

@app.middleware("http")
async def LoginAuth(request: Request, call_next):
    response = await call_next(request)
    if not request.user.is_authenticated:
        return JSONResponse({"msg:": "please login"},status_code=400)
    return response

# @app.middleware("http")
# async def add_process_time_header(request: Request, call_next):
#     response = await call_next(request)
#     if request.headers.get("Origin"):
#         response.headers["Access-Control-Allow-Origin"] = request.headers["Origin"]
#     return response
