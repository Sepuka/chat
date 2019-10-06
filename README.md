init
====
```bash
make init       initialize the project
    - it downloads the dependencies
    - make config file
make help       print help
```

###Запуск в режиме приема команд с терминала
Удобно использовать для отладки. Для указания инструкции следует заполнить параметр instr
```bash
make terminal instr=[list|create]
```
* list - показывает список машин доступных для юзера
* create - создает новую машину
###Запуск в режиме приема команд из telegram
Этот режим используется в production
```bash
make telegram
    accepted actions
        - create    to create new host
        - list      to review list of hosts
```