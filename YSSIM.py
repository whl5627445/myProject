# -- coding: utf-8 --
import logging
import os

import uvicorn
from fastapi import FastAPI, Request
# from fastapi.responses import JSONResponse, Response
from app.api.simulate.simulateview import router as simulate_view_router
from app.api.model.modelview import router as model_view_router
from app.api.modelfile.upload_file import router as upload_file_router
from app.api.modelfile.download import router as download_file_router
from app.api.user.user_space import router as user_space_router
from starlette.middleware.sessions import SessionMiddleware
from starlette.middleware.authentication import AuthenticationMiddleware
from fastapi.middleware.cors import CORSMiddleware
from starlette.authentication import (
    AuthenticationBackend, SimpleUser,
    AuthCredentials
    )
from config.InIt import InIt
from config.nacos_config import ServiceDiscovery


class BasicAuthBackend(AuthenticationBackend):
    async def authenticate (self, request):
        username = request.headers.get("username", None)
        user_space = request.headers.get("space_id", None)
        # print(request.headers)
        if not username or username == "sys":
            username = ""
        user = SimpleUser(username)
        user.user_space = user_space
        return AuthCredentials(["simtek"]), user


app = FastAPI()


app.include_router(simulate_view_router)
app.include_router(model_view_router)
app.include_router(upload_file_router)
app.include_router(download_file_router)
app.include_router(user_space_router)
app.add_middleware(AuthenticationMiddleware, backend=BasicAuthBackend())
app.add_middleware(SessionMiddleware, secret_key="simtek")


# app.add_middleware(
    # CORSMiddleware,
    # allow_origins=["*"],
    # allow_credentials=True,
    # allow_methods=["*"],
    # allow_headers=["*"],
    # expose_headers=["Content-Disposition"],
# )

if __name__ == '__main__':
    InIt()
    ServiceDiscovery()
    uvicorn.run("YSSIM:app",
                host="0.0.0.0",
                port=int(os.getenv("PORT")),
                workers=4,
                debug=False,
                )
