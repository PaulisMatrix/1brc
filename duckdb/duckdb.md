## solving using duckdb


1.  Import csv into duckdb table: `COPY weather FROM 'weather.csv' WITH (DELIMITER ',', HEADER true);`

2.  For checking query execution times, turn on timer: `.timer on` in seconds.

    Note: For explicit query and execution plans, `SET explain_output='all';` 

        1. Append EXPLAIN ANALYZE <query> -> To check total time taken.

        2. Append EXPLAIN <query> -> To check all unoptimized and optimized plans by duckdb.

3.  CLI options and dot commands : https://duckdb.org/docs/archive/0.9.2/api/cli

4.  db config : https://duckdb.org/docs/sql/configuration.html

5.  db pragmas : https://duckdb.org/docs/sql/pragmas.html

6.  duckdb will load the csv data into memory, `:memory:` db; as opposed to sqlite's on disk by default: 

    * source: https://duckdb.org/docs/archive/0.8.1/sql/statements/create_table.html#temporary-tables

    * It will start offloading the excess data to tmp files as needed. This is controlled by the `memory_limit`
    setting. 

    * You can specify which tmp directory to pick to store the excess data. https://duckdb.org/docs/sql/pragmas.html#temp-directory-for-spilling-data-to-disk

    * `SET temp_directory='/Users/rushiyadwade/Documents/go_dir/source/1brc';`

7.  Relevant command to load and calculate the avg temperatures/station: 

    ```
    time duckdb -c "CREATE TABLE weather(station VARCHAR NOT NULL, temperature REAL NOT NULL);" -c "COPY weather FROM 'weather1mi.txt' WITH (DELIMITER ',', HEADER true);" -c "SELECT station, ROUND(MIN(temperature),2), ROUND(MAX(temperature),2), ROUND(AVG(temperature),2) FROM weather GROUP BY station ORDER BY station;"

8.  Comparison results: https://x.com/mr_le_fox/status/1741933185175334923?s=20

9.  



        