import json

from events.core.listener import Listener


class PostCreatedListener(Listener):
    def __init__(self, client):
        self.subject = "POST_CREATED"
        self.client = client

    def parse(self, payload):
        post = json.loads(payload)
        print(post["id"], post["title"])

    def onMessage(self, message):
        self.parse(message)
