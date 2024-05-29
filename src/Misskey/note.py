# 1つ上のディレクトリの絶対パスを取得し、sys.pathに登録する
import sys
import os
from os.path import dirname
parent_dir = dirname(dirname(__file__))
if parent_dir not in sys.path:
    sys.path.append(parent_dir) 

from ngword_filter import is_ngword
import numpy as np
from misskey import Misskey
import json

misskey = Misskey(os.environ['SERVER'], i=os.environ['TOKEN'])


def note(sentence):
    if is_ngword(sentence) != True:
        misskey.notes_create(sentence)