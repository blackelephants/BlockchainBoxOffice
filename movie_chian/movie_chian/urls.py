"""movie_chian URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/1.10/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  url(r'^$', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  url(r'^$', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.conf.urls import url, include
    2. Add a URL to urlpatterns:  url(r'^blog/', include('blog.urls'))
"""
from django.conf.urls import url
from django.contrib import admin

from movie_chian import views

urlpatterns = [
    url(r'^admin/', admin.site.urls),
    url(r'^$', views.index, name='index'),
    url(r'^cinemas/', views.cinemas, name='cinemas'),
    url(r'^cinema_reg', views.cinema_reg, name='cinema_reg'),
    url(r'^cinema_query', views.cinema_query, name='cinema_query'),
    url(r'^halls', views.halls, name='halls'),
    url(r'^hall_reg', views.hall_reg, name='halls reg'),
    url(r'^hall_query', views.hall_query, name='hall_query'),

    url(r'^seller', views.seller, name='cinema_reg'),
    url(r'^capital/', views.capital, name='capital'),
    url(r'^capital_query', views.capital_query, name='capital_query'),
    url(r'^schedule/', views.schedule, name='schedule'),
    url(r'^schedule_query', views.schedule_query, name='schedule_query'),
    url(r'^audience', views.audience, name='audience'),
    url(r'^all_plan_query', views.all_plan_query, name='all_plan_query'),
    url(r'^lock_ticket', views.lock_ticket, name='lock_ticket'),
    url(r'^check_ticket', views.check_ticket, name='check_ticket'),
]
