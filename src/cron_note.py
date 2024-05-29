### main.py
from Misskey.get_timeline import get_tl_misskey
from Misskey.note import note
from yukimi_text.yukimi_text import change_yukimi
import logging

logging.basicConfig(level=logging.DEBUG)

def cron_note():
    text = get_tl_misskey()
    post_word = change_yukimi(text)
    note(post_word)

cron_note()

