-- weather table
CREATE TABLE IF NOT EXISTS weather (station TEXT NOT NULL, temperature REAL NOT NULL);

-- query to select max, min, avg of each station
SELECT station, MIN(temperature), MAX(temperature), AVG(temperature) FROM weather GROUP BY station;