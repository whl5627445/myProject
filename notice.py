# -- coding: utf-8 --
from fastapi import FastAPI, WebSocket, WebSocketDisconnect
import logging
from fastapi.responses import HTMLResponse
from datetime import datetime
from config.WebsocketManager import manager
from config.redis_config import r
import json
import random
from fastapi.responses import HTMLResponse
from starlette.middleware.sessions import SessionMiddleware
from starlette.middleware.authentication import AuthenticationMiddleware
from starlette.authentication import (
    AuthenticationBackend, SimpleUser,
    AuthCredentials
    )

MANAGER = manager


class BasicAuthBackend(AuthenticationBackend):
    async def authenticate (self, request):
        username = request.headers.get("username", "wanghailong")
        # username = request.headers.get("username", "")
        if not username:
            return
        return AuthCredentials(["simtek"]), SimpleUser(username)


app = FastAPI()
app.add_middleware(AuthenticationMiddleware, backend=BasicAuthBackend())
app.add_middleware(SessionMiddleware, secret_key="simtek")
logging.basicConfig(level=logging.DEBUG,  # 控制台打印的日志级别
                    filename='/home/simtek/code/Log/notice.log',
                    filemode='a',  ##模式，有w和a，w就是写模式，每次都会重新写日志，覆盖之前的日志
                    # a是追加模式，默认如果不写的话，就是追加模式
                    format='%(asctime)s - %(pathname)s[line:%(lineno)d] - %(levelname)s: %(message)s'
                    # 日志格式
                    )


@app.websocket("/notice/")
async def NotificationView (ws: WebSocket, username: str = None):
    """
    # 消息通知
    :return: 暂时只被动接受消息， 服务器不接受数据
    """
    # logging.debug(ws.__dict__)
    try:
        ws.num = random.randint(5000, 9999)
        await MANAGER.connect(ws, username)
        while True:
            res = {"status": False, "msg": "", "type": "message"}
            req_data = await ws.receive_text()
            req_data = json.loads(req_data)
            if type(req_data) is not dict:
                res_data = str(req_data).encode()
            else:
                res_data = r.rpop(req_data.get("user", "") + "_" + "notification")
            if res_data:
                r_data = json.loads(res_data.decode())
                message = r_data.get("message", "An error occurred")
                res["msg"] = str(datetime.now().strftime('%Y-%m-%d %H:%M:%S; ')) + str(message)
                res["status"] = True
                res["type"] = r_data.get("type", "message")
            data = json.dumps(res, ensure_ascii=False)
            await manager.send_personal_message(data, ws, username)
    except WebSocketDisconnect:
        manager.disconnect(ws, username)
        # await manager.send_personal_message(res, websocket)

# r.lpush(username + "_" + "notification",
#                 str(datetime.now().strftime('%Y-%m-%d %H:%M:%S; ')) + model_name + " 编译成功，开始仿真")
html = """
<!DOCTYPE html>
<html>
    <head>
        <title>Chat</title>
    </head>
    <body>
        <h1>WebSocket Chat</h1>
        <form action="" onsubmit="sendMessage(event)">
            <input type="text" id="messageText" autocomplete="off"/>
            <button>Send</button>
        </form>
        <ul id='messages'>
        </ul>
        <script>
            var ws = new WebSocket("ws://119.3.155.11:5555/notice/?username=wanghailong");

            ws.onmessage = function(event) {
                console.log(event.data)
                var messages = document.getElementById('messages')
                var message = document.createElement('li')
                var content = document.createTextNode(event.data)
                message.appendChild(content)
                messages.appendChild(message)
            };
            function sendMessage(event) {
                var input = document.getElementById("messageText")
                ws.send(input.value)
                input.value = ''
                event.preventDefault()
            }
        </script>
    </body>
</html>
"""


@app.get("/ws")
async def get ():
    return HTMLResponse(html)
