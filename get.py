import urllib.request, json
import os.path
import datetime
import time


def get_data_file(filetype, day, month, year):

    apiKey = '15621335427e418ea21335427e118ec6'
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


def get_data(day, month, year):
    print(f'... downloading {day}')
    get_data_file('all', day, month, year)
    get_data_file('hourly', day, month, year)
    get_data_file('daily', day, month, year)