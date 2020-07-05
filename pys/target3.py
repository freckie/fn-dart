import re
import requests
from bs4 import BeautifulSoup

re_numbers = re.compile('[0-9]+')
popup_url = 'http://dart.fss.or.kr/dsaf001/main.do?rcpNo={}'
report_url = 'http://dart.fss.or.kr/report/viewer.do?rcpNo={}&dcmNo={}&eleId=0&offset=0&length=0&dtd=HTML'

pat1 = re.compile('1\. ?제목')

def get_report(rcp_no):
    _url = popup_url.format(rcp_no)
    req = requests.get(_url)
    bs = BeautifulSoup(req.content, 'lxml')

    atag = bs.find('div', class_='view_search').find('a')
    dcm_no = re_numbers.findall(atag['onclick'])[1]

    _url2 = report_url.format(rcp_no, dcm_no)
    req2 = requests.get(_url2)
    bs2 = BeautifulSoup(req2.content, 'lxml')

    table = bs2.find('table').find('tbody')
    trs = table.find_all('tr')

    result = ['']
    for tr in trs:
        tds = tr.find_all('td')
        if len(tds) <= 1:
            continue

        cell_title = tds[0].get_text().strip()

        if pat1.search(cell_title):
            result[0] = tds[1].get_text().strip()

        else:
            continue
    
    return result

print(get_report('20200630900836'))