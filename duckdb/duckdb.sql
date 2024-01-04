-- weather table
CREATE TABLE IF NOT EXISTS weather (station VARCHAR NOT NULL, temperature DOUBLE NOT NULL);


-- query to select max, min, avg of each station
SELECT station, MIN(temperature), MAX(temperature), AVG(temperature) FROM weather GROUP BY station;