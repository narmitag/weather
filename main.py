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

        sql = "REPLACE INTO all_observations (stationID, obsTimeLocal, obsTimeUtc) VALUES (%s, %s, %s)"
        val = (ob['stationID'], ob['obsTimeLocal'], ob['obsTimeUtc'])
        mycursor.execute(sql, val)

        mydb.commit()

    exit()


def load_data(day, month, year):
    print(f'... processing {day}')
    process_data_file('all', day, month, year)
    #process_data_file('daily', day, month, year)
    #process_data_file('hourly', day, month, year)


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


