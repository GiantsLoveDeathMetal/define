import os
import sys
import json

import requests


LANGUAGE = 'en'
URL = 'https://od-api.oxforddictionaries.com:443/api/v1/entries/{language}/{word}'

APP_ID = os.environ.get('DICTIONARY_APP_ID')
APP_KEY = os.environ.get('DICTIONARY_APP_KEY')

HEADERS = {
    'app_id': APP_ID,
    'app_key': APP_KEY,
}


def main():
    word = sys.argv[1]
    url = URL.format(language=LANGUAGE, word=word)

    response = requests.get(url, headers=HEADERS)
    print("status: {}".format(response.status_code))
    # print("text: {}".format(response.text))
    word_definition =  json.loads(response.text)

    print(word_definition['metadata'])
