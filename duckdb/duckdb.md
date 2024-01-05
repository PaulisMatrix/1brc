## solving using duckdb


1.  Import csv into duckdb table: `COPY weather FROM 'weather.csv' WITH (DELIMITER ',', HEADER true);`

2.  For checking query execution times, turn on timer: `.timer on` in seconds.

        Note: For explicit query and execution plans, `SET explain_output='all';` 

            1. Append EXPLAIN ANALYZE <query> -> To check total time taken.

            2. Append EXPLAIN <query> -> To check all unoptimized and optimized plans by duckdb.

3.  CLI options and dot commands : https://duckdb.org/docs/archive/0.9.2/api/cli

4.  db config : https://duckdb.org/docs/sql/configuration.html

5.  db pragmas : https://duckdb.org/docs/sql/pragmas.html

6.  duckdb will load the csv data into memory, `:memory:` db; as opposed to sqlite by default: 

    * source: https://duckdb.org/docs/archive/0.8.1/sql/statements/create_table.html#temporary-tables

    * It will start offloading the excess data to tmp files as needed. This is controlled by the `memory_limit`
    setting. 

    * You can specify which tmp directory to pick to store the excess data. https://duckdb.org/docs/sql/pragmas.html#temp-directory-for-spilling-data-to-disk

7.  


## results:

1.  




        