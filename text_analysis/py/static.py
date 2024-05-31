import pandas as pd
import numpy as np

from collections import Counter

import os

root = os.getcwd()
filename = input("Enter filename: ")
print('-------------------------')

with open(f'{root}/txt_files/{filename}.txt','r', encoding="utf-8") as file:
    text = file.read()
        
words = text.split(' ')

vocab = pd.DataFrame.from_dict(Counter(words), orient='index').reset_index()#, columns=['word', 'freq'])\
vocab = vocab.rename(columns={"index": "word", 0: "freq"})
vocab = vocab.sort_values(by='freq', ascending=False, ignore_index=True)
vocab.index = vocab.index + 1
vocab['norm_freq'] = vocab['freq'] / vocab['freq'].sum()

N = len(vocab)
N1 = len(vocab[vocab['freq'] == 1])
N2 = len(vocab[vocab['freq'] == 2])
N_1N = N1 / N
N_2N = N2 / N
RR = 1 / len(vocab) * np.sum(np.power(vocab['norm_freq'], 2))
entropy = -np.sum(vocab['norm_freq'] * np.log(vocab['norm_freq']))

print(f'text length {len(vocab)}')
print(f'hapax legomena {N1}')
print(f'dis legomena {N2}')
print(f'hapax legomena fraction {N_1N:.4}')
print(f'dis legomena fraction {N_2N:.4}')
print(f'repeat rate {RR:.4}')
print(f'entropy {entropy:.4}')

print("-------------------------")