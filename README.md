# English Documentation

## Fetching and Text Analysis

#### Environment Variables in the .env file:
API_APP_ID - Telegram API Identifier \
API_APP_HASH - Telegram API Key

For more detailed description: https://core.telegram.org/#telegram-api

### Introduction
This program, written solely in the Go programming language, is aimed at 3 general functions:


1) Text analysis through .txt files
2) Fetching and parsing messages from public Telegram groups into .json and .txt by provided username. 
   - Example: liganetchat
3) Fetching and parsing comments from Youtube videos into .json and .txt by provided url
   - Example: https://www.youtube.com/watch?v=TiB8swIXYZM

The program itself is launched by the start.bat script which creates a new terminal window to work with the program.

### Program Workflow
The generalized path for the user is as follows:

1) Choose between fetching Telegram messages, Youtube or Analysis of existing json or txt files.

2) Authenticate to Telegram or Youtube or enter filename for analysis.

3) Specify details (e.g. message or comment limit) and parse into json and txt.

4) Analyze text from file.

### Authentication
For authentication, Telegram and Youtube use two different methods, but for both APIs to work, this path is necessary.

Telegram and Youtube have a session system in the code on first authentication. Json files created after the first use of the program are stored in the .credentials folder (auth_token.json for Telegram and user_cred.json for Youtube).

In the same folder, ytclient_secret.json is stored, which allows to identify the correct client for Google API to authenticate to Youtube.

#### - Youtube

OAuth2 is used for authentication. First, the following text is shown:
```
To work with Youtube:

Follow the link and authenticate:
https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=[id]&redirect_uri=[redirect_uri]
Copy the last link here after logging into your Google account
(format should be: http://localhost/?state=state-token&code=[code]&scope=[scope_urls]):
```

Following the instructions, the session file - user_cred.json - is created so there is no need to authenticate again when reusing the Youtube Service.

If there is an attempt to repeat the process with an existing user_cred.json, the program will pause at the point with a request for the Video URL.

#### -Telegram
If the auth_token.json file is missing in the .credentials folder, the program will ask for a phone number and password (or enter na if there is no password on the account).

After that, the user is asked (on each run) to enter the code that comes in the telegram itself.

If the code is correct, the program proceeds to fetching and parsing.

### Fetching and Parsing
Fetching after authentication happens similarly in both services, and differs in the number of messages the program can fetch. For Telegram, this limit is 2700 messages, for Youtube it is 100.

For Telegram, the function responsible for fetching:
```go 
func TelegramParse(ctx context.Context, api_id int, api_hash string) ([]*tg.Message, string, error)
```
For Youtube:
```go
func YoutubeParse(ctx context.Context) ([]byte, error)
```
Parsing from json to txt is done by the function:

```go
func JsonPrepoc(filename string) error
```
This function removes punctuation and emojis, leaving text with words separated only by spaces. Works the same for Telegram and Youtube.

### Linguistic Text Analysis
The function responsible for linguistic analysis:
```go
func TextAnalysis(filename string) error
```
For linguistic analysis, the program uses 6 parameters:

1) Text Length
2) Hapax Legomena 
3) Dis Legomena 
4) Hapax Legomenon Equation 
5) Dis Legomenon Equation 
6) Entropy
7) 
Linguistic analysis, paired with parsing through the JsonPrepoc() function, can occur if you use Option [A] at the beginning, which asks for the filename and what you have ready

(in case of choosing json [j], the entire function will go through JsonPrepoc() before analysis)

Documentation and libraries used in the project
https://core.telegram.org/#telegram-api
https://developers.google.com/youtube/v3
https://pkg.go.dev/github.com/gotd/td@v0.102.0
https://pkg.go.dev/google.golang.org/api/youtube/v3

#### Complete list of libraries in the go.mod file

# Ukrainian Documentation

## Фетчинг та аналіз тексту

Середовищні змінні (Environment Variables) в файлі .env:

API_APP_ID - Ідентифікатор API Телеграму
\
API_APP_HASH - API ключ Телеграму

Для більш детального опису: https://core.telegram.org/#telegram-api

### Вступ

Ця програма, написана суто на мові програмування Go, націлена на 3 загальних функції:
1) Аналіз тексту через .txt файли
2) Фетчинг та парсинг повідомлень з публічних груп Телеграма в .json та .txt за наведеним юзернеймом
   - Приклад: liganetchat
3) Фетчинг та парсинг коментарів з відео Youtube в .json та .txt за наведеним url
   - Приклад: https://www.youtube.com/watch?v=TiB8swIXYZM

Сама програма запускається скриптом start.bat який створює нове вікно терміналу для роботи з програмою.

### Робота програми

Узагальнений шлях для користувача є таким:

1) Вибір між фетчингом повідомлень Telegram'а, Youtube'а або Аналіз вже готових json або txt файлів.
2) Авторизація до Телеграму чи Youtube або введення назви файлу для аналізу. 
3) Уточнення деталей  (для прикладу: ліміт кількості повідомлень або коментарів) та парсинг в json та в txt.
4) Аналіз тексту з файлу.

### Авторизація

Для авторизації Телеграм та Youtube використовують два різних метода, але задля того щоб обидва API працювали, цей шлях необхідний. 

Телеграм та Youtube при першій авторизації мають систему сесій в коді. Json файли які створені після першого використання програми зберігаються в папці .credentials (auth_token.json для Телеграму та user_cred.json для Ютубу).

В цій же папці зберігаються ytclient_secret.json, який дозволяє визначити правильний клієнт для Google API для авторизації до Ютубу.  

#### - Youtube
Для авторизації використовується oauth2. Спершу показується наступний текст:
```
Для роботи з Ютубом:
1) Пройдіть по посиланню та авторизуйтесь:
https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=[id]&redirect_uri=[redirect_uri]
2) Cкопіюйте сюди останнє посилання після логіну в гугл аккаунт
(формат має бути: http://localhost/?state=state-token&code=[code]&scope=[scope_urls]):
```

Слідуючи інструкціям, сесійний файл - user_cred.json - створюється для непотреби авторизуватись знову при повторному використанні Ютуб Сервісу. 

Якщо є спроба повторити процесс з існуючим user_cred.json, програма пригне на момент з питанням про Video URL.
#### - Телеграм 

Якщо відсутній файл auth_token.json в папці .credentials, тоді програма запитує телефон та пароль (або вписати na, якщо паролю на аккаунті не стоїть). 

Після цього, користувача запитують (при кожному запуску) ввести код який приходить в сам телеграм. 

Якщо код правильний, програма далі переходить до фетчингу та парсингу.

### Фетчинг та Парсинг

Фетчинг після авторизації відбувається схоже в обох сервісах, та відрізняється кількостю повідомлень які програма може фетчити. Для телеграму цей ліміт є 2700 повідомлень, для Ютубу 100.

Для Телеграму, функція яка відповідає за фетчинг: 

```go 
func TelegramParse(ctx context.Context, api_id int, api_hash string) ([]*tg.Message, string, error) 
```
Для Ютубу:
```go
func YoutubeParse(ctx context.Context) ([]byte, error)
```

Парсинг з json в txt відбувається по функції: 

```go 
func JsonPrepoc(filename string) error
```

Ця функція забирає розділові знаки та емоджі, та залишає текст словами розділеними тільки пробілами. Працює для Телеграму та Ютубу однаково. 

### Лінгвістичний Аналіз Тексту
Функція відповідаюча за лінгвістичний аналіз:

```go
func TextAnalysis(filename string) error
```

Для лінгвістичного аналізу програма використовує 6 параметрів: 

1) Довжина тексту
2) Гапакс Легомена
3) Дис Легомена
4) Рівняння Гапаксу Легомена
5) Рівняння Дису Легомена
6) Ентропія

Лінгвістичний аналіз, на парі з парсингом через функцію JsonPrepoc() може відбутись якщо на початку використати Варіант [А], при якому запитує назву файла та що у вас є готове 

(в випадку вибору json [j], вся функція пройде через JsonPrepoc() перед аналізом)

### Документація та бібліотеки використані в проекті

https://core.telegram.org/#telegram-api
https://developers.google.com/youtube/v3
https://pkg.go.dev/github.com/gotd/td@v0.102.0
https://pkg.go.dev/google.golang.org/api/youtube/v3

Повний перелік бібліотек в файлі go.mod
