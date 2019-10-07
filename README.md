init
====
```bash
make init       initialize the project
    - it downloads the dependencies
    - make config file
make help       print help
```

###Запуск в режиме приема команд с терминала
Удобно использовать для отладки, после запуска программа программа выполнит переданную команду и завершится.
Вторым параметром указывается команда для выполнения
```bash
make terminal instr=list user=vasya
```
* list - показывает список машин доступных для юзера
* create - создает новую машину

###Запуск в режиме приема команд из telegram
Этот режим используется в production.
После запуска программа подключится к серверам telegram и начнет принимать команды для выполнения, одну за другой.
```bash
make telegram
    accepted actions
        - create    to create new host
        - list      to review list of hosts
```

###Подключение к бд
Для локальной отладки нужно запустить команду `make db` которая соберет образ бд postgres 11.5 и запустит сервер.
Чтобы наполнить бд таблицами нужно выполнить `make gen_tbl` она создаст таблицы в БД. Для создания дампа выполнить
`make pgdump` и сохранит его в _docker/parts/pg/dump.sql_