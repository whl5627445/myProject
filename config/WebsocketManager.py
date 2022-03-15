# -- coding: utf-8 --
from typing import List
from fastapi import  WebSocket, WebSocketDisconnect


class WebsocketConnectionManager:
    def __init__(self):
        self.active_connections: List[WebSocket] = []

    # 连接
    async def connect(self, websocket: WebSocket):
        await websocket.accept()
        self.active_connections.append(websocket)

    # 关闭连接
    def disconnect(self, websocket: WebSocket):
        self.active_connections.remove(websocket)

    # 发送单独消息
    async def send_personal_message(self, message: str, websocket: WebSocket):
        await websocket.send_text(message)

    # 所有连接发送消息
    async def broadcast(self, message: str):
        for connection in self.active_connections:
            await connection.send_text(message)


manager = WebsocketConnectionManager()
