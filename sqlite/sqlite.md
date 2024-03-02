## solving using sqlite


1.  Import csv into sqlite table: `.import --csv --skip 1 weather.csv weather -v`

2.  For checking query execution times, turn on timer: `.timer on` in seconds. 

    Note:

        1. user is the time that the CPU spent executing code in user space (i.e., in the database itself)
        2. sys is for code in the kernel.

3.  sqlite3 will create temp table on disk instead of in_memory which is the default behaviour.<br>
    Can change the setting to `MEMORY` to behave it as an in-memory db but since the file size is around 12GBs in this case, this wont work on RAMs < 12GBs obv.<br>
    The corresponding PRAGMA: [TEMP_STORE](https://www.sqlite.org/pragma.html#pragma_temp_store)

4.  Relevant command to load and calculate the avg temperatures/station: 

    ```
    time sqlite3 -cmd "CREATE TABLE weather(station TEXT NOT NULL, temperature REAL NOT NULL);" -cmd ".import --csv --skip 1 weather1mi.txt weather -v" -cmd "SELECT station, ROUND(MIN(temperature),2), ROUND(MAX(temperature),2), ROUND(AVG(temperature),2) FROM weather GROUP BY station ORDER BY station;" -cmd ".exit 1"   

5.  Sqlite's quirky exit terminal behaviour: 

    1.  https://sqlite.org/forum/info/4babe5d279078f95
    2.  https://stackoverflow.com/questions/31175018/how-to-run-command-line-sqlite-query-and-exit

6.  