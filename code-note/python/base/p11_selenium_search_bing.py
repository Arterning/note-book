from selenium import webdriver
import time

def wait():
    time.sleep(3)


def scrollSlowly():
    scheight = 9.9
    while scheight > .1:
        driver.execute_script("window.scrollTo(0, document.body.scrollHeight/%s);" % scheight)
        scheight -= .002

searchEngine = "https://cn.bing.com/"

driver = webdriver.Chrome()
driver.get(searchEngine)

excludeSomething = "nodejs -CSDN"
inUrl = "java inurl:stackoverflow"
searchText = "nodejs -CSDN"


driver.find_element_by_id("sb_form_q").send_keys(searchText)
driver.find_element_by_id("sb_form_go").click()
scrollSlowly()
time.sleep(10)

driver.find_element_by_xpath("//*[@id='b_results']/li[13]/nav/ul/li[9]/a").click()
scrollSlowly()
time.sleep(10)

while True:
    driver.find_element_by_xpath("//*[@id='b_results']/li[13]/nav/ul/li[9]/a").click()
    scrollSlowly()
    time.sleep(10)

