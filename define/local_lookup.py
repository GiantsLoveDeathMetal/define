import os

import pickle

HOME = os.path.expanduser('~')
PATH_WORDS = HOME + "/.dictionary/all_words.pkl"


def load_words():
    try:
        with open(PATH_WORDS, 'rb') as _input:
            return pickle.load(_input)
    except EOFError and IOError:
        return {}


def save_words(data):
    with open(PATH_WORDS, 'wb') as output:
        pickle.dump(data, output, pickle.HIGHEST_PROTOCOL)


def word_exists(word, all_words):
    if word in all_words.keys():
        return True
    return False


def get_word(path):
    print('local_sourced')
    with open(HOME + path, 'rb') as _input:
        return pickle.load(_input)


def save_file(path, data):
    with open(HOME+ path, 'wb') as output:
        pickle.dump(data, output, pickle.HIGHEST_PROTOCOL)
