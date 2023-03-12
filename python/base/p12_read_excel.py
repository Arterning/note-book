
import xlrd
from xlrd import xldate_as_tuple
import datetime
 
# pip install xlrd
#导入需要读取的第一个Excel表格的路径
data1 = xlrd.open_workbook('/home/ning/Untitled 1.xls')
table = data1.sheets()[0]
#创建一个空列表，存储Excel的数据
tables = []

#将excel表格内容导入到tables列表中
def import_excel(excel):
   for rown in range(excel.nrows):
      array = {}
      array['0'] = table.cell_value(rown,0)
      array['1'] = table.cell_value(rown,1)
      array['2'] = table.cell_value(rown,2)
      array['3'] = table.cell_value(rown,3)
      array['4'] = table.cell_value(rown,4)
      array['5'] = table.cell_value(rown,5)
    #   array['6'] = table.cell_value(rown,6)
      tables.append(array)
 
if __name__ == '__main__':
   #将excel表格的内容导入到列表中
   import_excel(table)
   for i in tables:
       print(i)


