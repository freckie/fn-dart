import re
import requests
from bs4 import BeautifulSoup

re_numbers = re.compile('[0-9]+')
popup_url = 'http://dart.fss.or.kr/dsaf001/main.do?rcpNo={}'
report_url = 'http://dart.fss.or.kr/report/viewer.do?rcpNo={}&dcmNo={}&eleId=0&offset=0&length=0&dtd=HTML'

pat1 = re.compile('1\. ?판매ㆍ공급 ?계약 ?(내용|구분)')
pat11 = re.compile('- ?체결 ?계약명?')
pat2 = re.compile('2\. ?계약 ?내역')
pat21 = re.compile('(확정 ?계약 ?금액)|(계약 ?금액 ?\(원\))')
pat22 = re.compile('매출액 ?대비 ?\(\%\)')
pat3 = re.compile('3\. ?계약 ?상대방?')
pat4 = re.compile('4\. ?판매ㆍ공급 ?지역')

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
    level = 1
    for tr in trs:
        tds = tr.find_all('td')
        if len(tds) <= 1:
            continue

        cell_title = tds[0].get_text().strip()

        if level == 2:
            if pat21.search(cell_title):
                result[1] = tds[1].get_text().strip()
            elif pat22.search(cell_title):
                result[2] = tds[1].get_text().strip()

        if pat1.search(cell_title):
            level = 1
            result[0] = tds[1].get_text().strip()

        elif pat11.search(cell_title):
            level = 1
            result[0] = tds[1].get_text().strip()

        elif pat3.search(cell_title):
            level = 1
            result[3] = tds[1].get_text().strip()

        elif pat4.search(cell_title):
            level = 1
            result[4] = tds[1].get_text().strip()

        elif pat2.search(cell_title) or level == 2:
            level = 2
            if pat21.search(tds[1].get_text()):
                result[1] = tds[2].get_text().strip()
            elif pat22.search(tds[1].get_text()):
                result[2] = tds[2].get_text().strip()

        else:
            continue
    
    return result

print(get_report('20200701900192'))