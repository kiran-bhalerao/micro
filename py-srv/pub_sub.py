import json


class PubSubClient:
    def __init__(self, redis_client):
        self.redis = redis_client
        self.pubsub = redis_client.pubsub()

    def subscribe(self, subject):
        self.pubsub.subscribe(subject)
        return self.pubsub

    def publish(self, subject, data):
        self.redis.publish(subject, json.dumps(data))
