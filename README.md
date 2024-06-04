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
