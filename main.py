from fastapi import FastAPI, Request, Response
import pydevd_pycharm
from app.api.simulateview import router as simulate_view_router
from app.api.modelview import router as model_view_router
from app.api.upload_file import router as upload_file_router
from starlette.middleware.sessions import SessionMiddleware
from starlette.middleware.authentication import AuthenticationMiddleware
from fastapi.middleware.cors import CORSMiddleware
from starlette.authentication import (
    AuthenticationBackend, AuthenticationError, SimpleUser,
    AuthCredentials
)

pydevd_pycharm.settrace('192.168.1.110', port=10086, stdoutToServer=True, stderrToServer=True)

class BasicAuthBackend(AuthenticationBackend):
    async def authenticate(self, request):
        # if "Authorization" not in request.headers:
        #     return
        #
        # auth = request.headers["Authorization"]
        # try:
        #     scheme, credentials = auth.split()
        #     if scheme.lower() != 'basic':
        #         return
        #     decoded = base64.b64decode(credentials).decode("ascii")
        # except (ValueError, UnicodeDecodeError, binascii.Error) as exc:
        #     raise AuthenticationError('Invalid basic auth credentials')
        #
        # username, _, password = decoded.partition(":")
        username = {
            "username": "tom"
        }
        # TODO: You'd want to verify the username and password here.
        # return AuthCredentials(["authenticated"]), SimpleUser(username)
        return True, username


app = FastAPI()
app.include_router(simulate_view_router)
app.include_router(model_view_router)
app.include_router(upload_file_router)
app.add_middleware(SessionMiddleware, secret_key="simtek")
app.add_middleware(AuthenticationMiddleware, backend=BasicAuthBackend())
origins = [
    ""
]

app.add_middleware(
    CORSMiddleware,
    # allow_origins=origins,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


# @app.middleware("http")
# async def add_process_time_header(request: Request, call_next):
#     secret = "abcdefghijklmnopqrstuvwxyz"
#     token = request.cookies.get("Admin-Token")
#     # try:
#     #     decoded = jwt.decode(token, secret, algorithms=["HS512"])
#     # except Exception as e:
#     #     response = responses.JSONResponse({"msg": "验证错误"})
#     #     return response
#     # key = "login_tokens:" + str(decoded["login_user_key"])
#     key = token
#     r_data = r.get(key)
#     if r_data:
#         data = eval(r.get(key))
#         print(data)
#     else:
#         pass
#         # response = responses.JSONResponse({"msg": "验证错误"})
#         # return response
#     response = await call_next(request)
#     return response




