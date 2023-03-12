from django.urls import include, path

from . import views
from rest_framework import routers
from . import views

# http://127.0.0.1:8000/passkeep/
# path() 参数： 
# route '' 
# view views.index
# name URL 取名 在模板中这个特性有用的


# http://127.0.0.1:8000/passkeep/rest-token/
router = routers.DefaultRouter()
router.register(r'rest-token', views.TokenViewSet)

urlpatterns = [
    path('', views.index, name='index'),
    path('addToken/', views.addToken, name='addToken'),

    # 这个url可不可以改造成适配post请求 因为我需要提交一个表单 
    # 我不希望定义成tokens/add这种形式 因为我希望符合restful风格的API
    path('tokens/', views.allToken, name='allToken'),

    # ex: /tokens/5/
    path('tokens/<int:token_id>', views.detail, name='detail'),

    path('', include(router.urls)), #使用router路由
    path('api-auth/', include('rest_framework.urls', namespace='rest_framework'))
]