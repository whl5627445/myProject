# -- coding: utf-8 --
import logging
from typing import List, Dict
from fastapi import  WebSocket, WebSocketDisconnect


class WebsocketConnectionManager(object):
    def __init__(self):
        self.active_connections: dict[dict[str: WebSocket]] = {}
    # 连接
    async def connect(self, ws, username):
        await ws.accept()
        if self.active_connections.get(username, None):
            self.active_connections[username][ws.num] = ws
        else:
            self.active_connections[username] = {}
            self.active_connections[username][ws.num] = ws

    # # 关闭连接
    def disconnect(self, websocket, username):
        self.active_connections[username].pop(websocket.num, None)

    # 发送单独消息
    async def send_personal_message(self, message: str, websocket, username):
        for name, connection_dict in self.active_connections.items():
            if name == username:
                for num, connection in connection_dict.items():
                    try:
                        if connection.client_state != 2:
                            await connection.send_text(message)
                        else:
                            self.disconnect(websocket, username)
                    except WebSocketDisconnect:
                        self.disconnect(websocket, username)

    # 所有连接发送消息
    async def broadcast(self, message: str, websocket, username):
        for name, connection_dict in self.active_connections.items():
            for num, connection in connection_dict.items():
                try:
                    if connection.client_state != 2:
                        await connection.send_text(message)
                    else:
                        self.disconnect(websocket, username)
                except WebSocketDisconnect:
                    self.disconnect(websocket, username)
manager = WebsocketConnectionManager()
