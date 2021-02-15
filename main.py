import urllib.request, json
import os.path
import datetime
import time
from get import get_data
import json
import mysql.connector


def process_data_file(filetype, day, month,year):
    filename = f'data/{filetype}/{year}/{month}/{day}.json'
    path = f'data/{filetype}/{year}/{month}/'

    mydb = mysql.connector.connect(
        host="192.168.0.5",
        port=3307,
        user="weather",
        password="Password11!",
        database="weather"
    )

    if not os.path.isfile(filename):
        print(f'.....{filetype}/{day} is missing???')
        exit(2)
    print(f'...... loading {filetype}/{day}')
    with open(filename, "r") as text_file:
        data = json.load(text_file)

    for ob in data['observations']:
        print(ob)
        mycursor = mydb.cursor()
        tablename = f'{filetype}_observations'
        print(tablename)
        sql = f"REPLACE INTO {tablename}   " \
              f"VALUES (%s, %s, %s,%s, %s, %s,%s, %s, %s, %s," \
              f"%s, %s, %s,%s, %s, %s,%s, %s, %s, %s," \
              f"%s, %s, %s,%s, %s, %s,%s, %s, %s, %s," \
              f"%s, %s, %s,%s, %s, %s,%s)"
        val = (ob['stationID'],
                ob['tz'],
                ob['obsTimeUtc'],
                ob['obsTimeLocal'],
                ob['epoch'],
                ob['lat'],
                ob['lon'],
                ob['solarRadiationHigh'],
                ob['uvHigh'],
                ob['winddirAvg'],
                ob['humidityHigh'],
                ob['humidityLow'],
                ob['humidityAvg'],
                ob['qcStatus'],
                ob['metric']['tempHigh'],
                ob['metric']['tempLow'],
                ob['metric']['tempAvg'],
                ob['metric']['windspeedHigh'],
                ob['metric']['windspeedLow'],
                ob['metric']['windspeedAvg'],
                ob['metric']['windgustHigh'],
                ob['metric']['windgustLow'],
                ob['metric']['windgustAvg'],
                ob['metric']['dewptHigh'],
                ob['metric']['dewptLow'],
                ob['metric']['dewptAvg'],
                ob['metric']['windchillHigh'],
                ob['metric']['windchillLow'],
                ob['metric']['windchillAvg'],
                ob['metric']['heatindexHigh'],
                ob['metric']['heatindexLow'],
                ob['metric']['heatindexAvg'],
                ob['metric']['pressureMax'],
                ob['metric']['pressureMin'],
                ob['metric']['pressureTrend'],
                ob['metric']['precipRate'],
                ob['metric']['precipTotal'])

        mycursor.execute(sql, val)
        mydb.commit()


def load_data(day, month, year):
    print(f'... processing {day}')
    process_data_file('all', day, month, year)
    process_data_file('daily', day, month, year)
    process_data_file('hourly', day, month, year)


if __name__ == '__main__':
    start_date = datetime.datetime(2019, 1, 18)
    today = datetime.datetime.today()
    delta = today - start_date
    # Don't download today
    for i in range(delta.days - 1):
        start_date += datetime.timedelta(days=1)
        download_day = start_date.strftime('%Y%m%d')
        download_year = start_date.strftime('%Y')
        download_month = start_date.strftime('%m')
        #get_data(download_day, download_month, download_year)
        load_data(download_day, download_month, download_year)
    print('Processing Complete')


