# -- coding: utf-8 --
import logging
from typing import List, Dict
from fastapi import  WebSocket, WebSocketDisconnect


class WebsocketConnectionManager(object):
    def __init__(self):
        self.active_connections: dict[str: WebSocket] = {}
    # 连接
    async def connect(self, ws, username):
        await ws.accept()
        self.active_connections[username] = ws

    # # 关闭连接
    def disconnect(self, username):
        del self.active_connections[username]

    # 发送单独消息
    async def send_personal_message(self, message: str, username):
        try:
            await self.active_connections[username].send_text(message)
        except WebSocketDisconnect:
            self.disconnect(username)

    # 所有连接发送消息
    async def broadcast(self, message: str, username):
        for name, connection in self.active_connections.items():
            try:
                await connection.send_text(message)
            except WebSocketDisconnect:
                self.disconnect(username)


manager = WebsocketConnectionManager()
