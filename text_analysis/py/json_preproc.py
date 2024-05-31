import json

import pandas as pd

import re

import os

root = os.getcwd()

json_path = 'json_files'
txt_path = 'txt_files'
# filename = 'liganet'
filename = input("Enter filename: ")

with open(f'{root}/{json_path}/{filename}.json', 'r') as myfile:
    obj = json.loads(myfile.read())
    
texts = ''

for i in range(0, len(obj['messages'])):
    if 'text' in obj['messages'][i].keys():
        message = obj['messages'][i]['text']
        if isinstance(message, list):
            for i in range(0, len(message)):
                if isinstance(message[i], dict):
                    texts += '\n'+ message[i]['text']
                else:
                    texts += '\n' + message[i]
        else:
            texts += '\n' + message
            
texts = texts.replace('.', ' ').replace('\n', ' ').replace(',', ' ')

texts = re.sub(r'\W+', ' ', texts)

texts = re.sub('\s+', ' ', texts)

words = texts.split(' ')

with open(f'{root}/{txt_path}/{filename}.txt','w') as file:
    file.write(texts.lower())