import os

import redis


def NewRedisClient():
    return redis.Redis(host=os.getenv("REDIS_HOST"), port=6379, db=0)
