import urllib.request, json
import os.path
import datetime
import time
#import mysql.connector


def process_data_file(filetype, day, month,year):
    filename = f'data/{filetype}/{year}/{month}/{day}.json'
    path = f'data/{filetype}/{year}/{month}/'

    # mydb = mysql.connector.connect(
    #     host="localhost",
    #     port=3306,
    #     user="weather",
    #     password="Password11!",
    #     database="weather",
    #     auth_plugin='mysql_native_password'
    # )

    if not os.path.isfile(filename):
        print(f'.....{filetype}/{day} is missing???')
        exit(2)
    print(f'...... loading {filetype}/{day}')
    with open(filename, "r") as text_file:
        data = json.load(text_file)

    # for ob in data['observations']:
    #     mycursor = mydb.cursor()
    #     tablename = f'{filetype}_observations'
    #     sql = f"REPLACE INTO {tablename}   " \
    #           f"VALUES (%s, %s, %s,%s, %s, %s,%s, %s, %s, %s," \
    #           f"%s, %s, %s,%s, %s, %s,%s, %s, %s, %s," \
    #           f"%s, %s, %s,%s, %s, %s,%s, %s, %s, %s," \
    #           f"%s, %s, %s,%s, %s, %s,%s)"
    #     val = (ob['stationID'],
    #             ob['tz'],
    #             ob['obsTimeUtc'],
    #             ob['obsTimeLocal'],
    #             ob['epoch'],
    #             ob['lat'],
    #             ob['lon'],
    #             ob['solarRadiationHigh'],
    #             ob['uvHigh'],
    #             ob['winddirAvg'],
    #             ob['humidityHigh'],
    #             ob['humidityLow'],
    #             ob['humidityAvg'],
    #             ob['qcStatus'],
    #             ob['metric']['tempHigh'],
    #             ob['metric']['tempLow'],
    #             ob['metric']['tempAvg'],
    #             ob['metric']['windspeedHigh'],
    #             ob['metric']['windspeedLow'],
    #             ob['metric']['windspeedAvg'],
    #             ob['metric']['windgustHigh'],
    #             ob['metric']['windgustLow'],
    #             ob['metric']['windgustAvg'],
    #             ob['metric']['dewptHigh'],
    #             ob['metric']['dewptLow'],
    #             ob['metric']['dewptAvg'],
    #             ob['metric']['windchillHigh'],
    #             ob['metric']['windchillLow'],
    #             ob['metric']['windchillAvg'],
    #             ob['metric']['heatindexHigh'],
    #             ob['metric']['heatindexLow'],
    #             ob['metric']['heatindexAvg'],
    #             ob['metric']['pressureMax'],
    #             ob['metric']['pressureMin'],
    #             ob['metric']['pressureTrend'],
    #             ob['metric']['precipRate'],
    #             ob['metric']['precipTotal'])

    #     mycursor.execute(sql, val)
    #     mydb.commit()




def get_data_file(filetype, day, month, year):


    apiKey  = os.getenv('API_KEY')
    station = 'ISTGBUCH2'
    wgurl = f'https://api.weather.com/v2/pws/history/{filetype}?stationId={station}&format=json&units=m&apiKey={apiKey}&date={day}'
    filename = f'data/{filetype}/{year}/{month}/{day}.json'
    path = f'data/{filetype}/{year}/{month}/'

    if not os.path.exists(path):
        os.makedirs(path)

    if os.path.isfile(filename):
        print(f'.....already got file for {filetype}/{day}')
    else:
        print(f'.....downloading file for {filetype}/{day}')
        time.sleep(5)
        with urllib.request.urlopen(wgurl) as url:
            with open(filename, "w") as text_file:
                text_file.write(url.read().decode())
        process_data_file(filetype, day, month, year)


def get_data(day, month, year):
    print(f'... downloading {day}')
    get_data_file('all', day, month, year)
    get_data_file('hourly', day, month, year)
    get_data_file('daily', day, month, year)