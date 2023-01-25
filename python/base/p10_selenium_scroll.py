from selenium import webdriver
import time
def wait():
    time.sleep(3)


def scrollSlowly():
    # scheight = .1
    # while scheight < 9.9:
    #     driver.execute_script("window.scrollTo(0, document.body.scrollHeight/%s);" % scheight)
    #     scheight += .01
    
    scheight = 9.9
    while scheight > .1:
        driver.execute_script("window.scrollTo(0, document.body.scrollHeight/%s);" % scheight)
        scheight -= .002
    
    # scheight = 0
    # while scheight < 2291:
    #     driver.execute_script("window.scrollTo(0, %s);" % scheight)
    #     scheight += 30
    #     time.sleep(0.5)

searchEngine = "www.biying.com"
searchEngine1 = "https://www.baidu.com/"

driver = webdriver.Chrome()
driver.get("https://www.baidu.com/")
print(driver.title + "，狗屁，百度垃圾")

# searchText = "node"
# searchText1 = "(FC2)(702535)"
# searchText2 = "selenium"
# searchText = "vs code 选择单词快捷键"
excludeSomething = "nodejs -CSDN"
inUrl = "java inurl:stackoverflow"

searchText = "搜索引擎使用技巧"


driver.find_element_by_id("kw").send_keys(searchText)
driver.find_element_by_id("su").click()
scrollSlowly()
time.sleep(10)

driver.find_element_by_class_name("n").click()
scrollSlowly()
time.sleep(10)

while True:
    alink  = driver.find_elements_by_class_name("n")
    next = alink[1]
    next.click()
    scrollSlowly()
    time.sleep(10)

