import datetime
from get import get_data

if __name__ == '__main__':
    start_date = datetime.datetime(2024, 1, 1)
    today = datetime.datetime.today()
    delta = today - start_date
    # Don't download today
    for i in range(delta.days - 1):
        start_date += datetime.timedelta(days=1)
        download_day = start_date.strftime('%Y%m%d')
        download_year = start_date.strftime('%Y')
        download_month = start_date.strftime('%m')
        get_data(download_day, download_month, download_year)
    print('Processing Complete')


