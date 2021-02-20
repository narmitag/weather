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
        host="localhost",
        port=3306,
        user="weather",
        password="Password11!",
        database="weather",
        auth_plugin='mysql_native_password'
    )

    if not os.path.isfile(filename):
        print(f'.....{filetype}/{day} is missing???')
        exit(2)
    print(f'...... loading {filetype}/{day}')
    with open(filename, "r") as text_file:
        data = json.load(text_file)

    for ob in data['observations']:
        mycursor = mydb.cursor()
        tablename = f'{filetype}_observations'
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


