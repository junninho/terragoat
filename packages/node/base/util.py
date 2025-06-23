from django.conf import settings
from hashids import Hashids
from django.shortcuts import redirect
from django.http import HttpResponse

def get_bad_hashids():
    return Hashids(
        salt=settings.SECRET_KEY, min_length=4, alphabet="abcdefghijklmnopqrstuvwxyz"
    )

def open_redirect(request):
    url = request.GET.get("url")
    return redirect(url)

def bad_sql(request):
  user_name = request.GET.get('user_name')
  user_age = Person.objects.extra(where=["name = '%s'" % user_name])
  return HttpResponse(user_age) 