## solving using sqlite


1. Import csv into sqlite table: `.import --csv --skip 1 weather.csv weather -v`

2. For checking query execution times, turn on timer: `.timer on` in seconds. 

    Note:

        1. user is the time that the CPU spent executing code in user space (i.e., in the database itself)
        2. sys is for code in the kernel.

3. sqlite3 will create temp table when you import.