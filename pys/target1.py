import re
import requests
from bs4 import BeautifulSoup

re_numbers = re.compile('[0-9]+')
popup_url = 'http://dart.fss.or.kr/dsaf001/main.do?rcpNo={}'
report_url = 'http://dart.fss.or.kr/report/viewer.do?rcpNo={}&dcmNo={}&eleId=0&offset=0&length=0&dtd=HTML'

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

    result = [''] * 5
    level = 0
    for tr in trs:
        tds = tr.find_all('td')
        cell_title = tds[0].get_text()

        if '판매ㆍ공급계약 내용' in cell_title:
            level = 0
            result[0] = tds[1].get_text().strip()

        elif '계약 내역' in cell_title or level == 1:
            level = 1
            if '확정 계약금액' in tds[1].get_text():
                result[1] = tds[2].get_text().strip()
            elif '':
                pass

        elif '계약 상대방' in cell_title:
            level = 0
            result[3] = tds[1].get_text().strip()

        elif '판매ㆍ공급지역' in cell_title:
            level = 0
            result[4] = tds[1].get_text().strip()

        else:
            continue
    

    return table

print(get_report('20200703900292'))