from django.contrib import admin
from django.urls import path
from django.conf import settings
from django.conf.urls.static import static

# from .views import *

from .import views,hod_views,staff_views,student_views
from .views import base,login,doLogin

urlpatterns = [
    path('admin/', admin.site.urls),
    path('base/', view=base, name='base'),
    
    # login system
    path('', view=login, name='login'),
    path('dologin', view=doLogin, name="doLogin"),
    path('doLogout', views.doLogout, name="logout"),
    
    path('hod/home', hod_views.home, name='hod_home')
    
    
    
] + static(settings.MEDIA_URL,document_root = settings.MEDIA_ROOT)
