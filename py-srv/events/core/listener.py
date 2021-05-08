import asyncio
from abc import ABC, abstractmethod


def listener(ps, cb):
    while True:
        for m in ps.listen():
            message = m['data']
            if type(message) is bytes:
                cb(m['data'].decode('utf-8'))


class Listener(ABC):
    def __init__(self, subject, client):
        self.subject = subject
        self.client = client

    def listen(self):
        ps = self.client.subscribe(self.subject)
        loop = asyncio.get_event_loop()
        loop.run_in_executor(None, lambda: listener(ps, self.onMessage))

    @abstractmethod
    def onMessage(self, message):
        pass

    @abstractmethod
    def parse(self, message):
        pass
