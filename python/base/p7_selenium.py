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


driver.find_element_by_id("kw").send_keys(searchText2)
driver.find_element_by_id("su").click()
wait()  # 强制等待3秒再执行下一步
driver.find_element_by_class_name("n").click()
wait()   # 强制等待3秒再执行下一步
while True:
    alink  = driver.find_elements_by_class_name("n")
    next = alink[1]
    next.click()
    wait() 


