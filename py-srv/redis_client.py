import redis


def NewRedisClient():
    return redis.Redis(host='localhost', port=6379, db=0)
