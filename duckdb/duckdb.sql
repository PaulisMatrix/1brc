-- weather table
CREATE TABLE IF NOT EXISTS weather (station VARCHAR NOT NULL, temperature DOUBLE NOT NULL);


-- query to select max, min, avg of each station
SELECT station, ROUND(MIN(temperature),2), ROUND(MAX(temperature),2), ROUND(AVG(temperature),2) FROM weather GROUP BY station ORDER BY station;