import os
import sys
import json

import requests

from define.local_lookup import word_exists
from define.local_lookup import load_words
from define.local_lookup import get_word
from define.local_lookup import save_file
from define.local_lookup import save_words


LANGUAGE = 'en'
URL = 'https://od-api.oxforddictionaries.com:443/api/v1/entries/{language}/{word}'

APP_ID = os.environ.get('DICTIONARY_APP_ID')
APP_KEY = os.environ.get('DICTIONARY_APP_KEY')

HEADERS = {
    'app_id': APP_ID,
    'app_key': APP_KEY,
}

def format_response(response):
    print(response)


def main():
    word = sys.argv[1]

    all_words = load_words()
    if word_exists(word, all_words):
        word_path = all_words[word]
        response = get_word(word_path)
    else:
        url = URL.format(language=LANGUAGE, word=word)
        response = requests.get(url, headers=HEADERS)
        if response.status_code == 200:
            new_file = '/.dictionary/words/{}.pkl'.format(word)
            response = json.loads(response.text)
            all_words[word] = new_file
            save_words(all_words)
            save_file(new_file, response)

    format_response(response)

