from events.core.publisher import Publisher


class UserCreatedPublisher(Publisher):
    def __init__(self, client):
        self.subject = "USER_CREATED"
        self.client = client
