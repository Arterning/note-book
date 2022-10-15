from selenium import webdriver
import time
def wait():
    time.sleep(10)

driver = webdriver.Chrome()
driver.get("https://www.baidu.com/")
print(driver.title)

searchText = "node"
searchText1 = "(FC2)(702535)"
searchText2 = "selenium"
searchText3 = "ubuntu sougou input"
searchText4 = "fcitx找不到搜狗输入法"
searchText5 = "selenium 滚动屏幕"


driver.find_element_by_id("kw").send_keys(searchText5)
driver.find_element_by_id("su").click()
wait()
driver.find_element_by_class_name("n").click()
wait()


