### main.py
import time
from Misskey.follow_back import follow_back
from Misskey.get_timeline import get_tl_misskey
from Misskey.note import note
from yukimi_text.yukimi_text import change_yukimi
import logging

logging.basicConfig(level=logging.DEBUG)

def cron_follow_back():
    follow_back()

cron_follow_back()