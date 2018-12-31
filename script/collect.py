# coding=utf-8

import MySQLdb
import urllib2
import os
import sys
import json
import pytz
import time
import dateutil.parser
from datetime import datetime
reload(sys)
sys.setdefaultencoding('utf-8')


db = MySQLdb.connect("127.0.0.1", "root", "..w114lc..", "sniffing", charset='utf8' )
cursor = db.cursor()

def save_into_mysql(type, things):
    for t in things:
        imgs = ""
        if t.has_key("images"):
            for img in t["images"]:
                imgs = img + ","
        c_time = dateutil.parser.parse(t["createdAt"]).astimezone(pytz.timezone('Asia/Shanghai'))

        c_int = int(time.mktime(c_time.timetuple()))
        p_time =  dateutil.parser.parse(t["publishedAt"]).astimezone(pytz.timezone('Asia/Shanghai'))
        p_int = int(time.mktime(p_time.timetuple()))
        desc = t["desc"].replace('"', '\\"')
        sql = "INSERT INTO thing(title, create_time, breif, images, publish_time, source, type, origin_url, author)"
        sql = sql + " VALUES(\"" + desc + "\"," + str(c_int) + ",\"" + desc + "\",\"" + imgs
        sql = sql + "\"," + str(p_int) + ",\"" + t["source"] + "\",\"" + t["type"] + "\",\""
        sql = sql + t["url"] + "\",\"" + t["who"] + "\")"
        print sql
        cursor.execute(sql)

res = {}
try:
    f = urllib2.urlopen("http://gank.io/api/today")
    res = f.read()
except:
    print "urllib open failed"

rson = json.loads(res)

for key, value in rson["results"].iteritems():
    save_into_mysql(key, value)
    print key

db.commit()
