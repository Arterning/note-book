from django.shortcuts import render

# Create your views here.
from django.http import HttpResponse
from .models import TokenInfo
from django.template import loader

def index(request):
    return HttpResponse("Hello, world. You're at the passkeep index.")

def all(request):
    resp = TokenInfo.objects.all()
    template = loader.get_template('passkeep/all.html')
    context = {
        'token_list': resp,
    }
    return HttpResponse(template.render(context, request))

def detail(request, token_id):
    return HttpResponse("You're looking at question %s." % token_id)

def update(request, token_id):
    return HttpResponse("You're looking at question %s." % token_id)