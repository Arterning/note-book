from django.urls import path

from . import views

# http://127.0.0.1:8000/passkeep/
# path() 参数： 
# route '' 
# view views.index
# name URL 取名 在模板中这个特性有用的
# 

urlpatterns = [
    path('', views.index, name='index'),

    path('tokens/all', views.all, name='all'),
    # ex: /tokens/5/
    path('tokens/<int:token_id>', views.detail, name='detail'),
]