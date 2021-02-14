import urllib.request, json
import os.path
import datetime
import time
from get import get_data

if __name__ == '__main__':
    start_date = datetime.datetime(2019, 1, 18)
    today = datetime.datetime.today()
    delta = today - start_date
    # Don't download todays
    for i in range(delta.days - 1):
        start_date += datetime.timedelta(days=1)
        download_day = start_date.strftime('%Y%m%d')
        download_year = start_date.strftime('%Y')
        download_month = start_date.strftime('%m')
        #get_data(download_day, download_month, download_year)
    print('Processing Complete')


