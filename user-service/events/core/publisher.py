import json


class Publisher:
    def __init__(self, subject, client):
        self.subject = subject
        self.client = client

    def publish(self, data):
        self.client.publish(self.subject, data)
