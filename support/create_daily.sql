drop table weather.daily_observations;

CREATE TABLE weather.daily_observations (
 stationID varchar(20),
 tz varchar(30),
 obsTimeUtc varchar(30),
 obsTimeLocal datetime,
 epoch int,
 lat float,
 lon float,
 solarRadiationHigh float,
 uvHigh float,
 winddirAvg int,
 humidityHigh float,
 humidityLow float,
 humidityAvg  float,
 qcStatus  float,
 tempHigh  float,
 tempLow float,
tempAvg  float,
windspeedHigh   float,
windspeedLow   float,
windspeedAvg float,
windgustHigh  float,
windgustLow  float,
windgustAvg  float,
dewptHigh  float,
dewptLow  float,
dewptAvg  float,
windchillHigh  float,
windchillLow  float,
windchillAvg  float,
heatindexHigh   float,
heatindexLow  float,
heatindexAvg  float,
pressureMax  float,
pressureMin   float,
pressureTrend   varchar(30),
precipRate  float,
precipTotal  float,
PRIMARY KEY (stationID, obsTimeLocal)
)