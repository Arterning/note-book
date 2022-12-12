from django.shortcuts import render

# Create your views here.
from django.http import HttpResponse
from .models import TokenInfo
from django.template import loader

def index(request):
    return HttpResponse("Hello, world. You're at the passkeep index.")

def addToken(request):
    context = {
    }
    template = loader.get_template('passkeep/add.html')
    return HttpResponse(template.render(context, request))

def allToken(request):
    resp = TokenInfo.objects.all()
    template = loader.get_template('passkeep/all.html')
    context = {
        'token_list': resp,
    }
    return HttpResponse(template.render(context, request))


def detail(request, token_id):
    if request.method == 'PUT':
        return HttpResponse("You're posting a token")
    elif request.method == 'DELETE':
        pass
    return HttpResponse("You're looking at question %s." % token_id)
