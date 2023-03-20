from django.shortcuts import render

# Create your views here.
from django.http import HttpResponse
from .models import TokenInfo
from django.template import loader
from .forms import TokenForm
from django.http import HttpResponseRedirect

from .serializer import TokenInfoSerializer
from rest_framework import viewsets


def index(request):
    return HttpResponse("Hello, world. You're at the passkeep index.")

def addToken(request):
    # if this is a POST request we need to process the form data
    if request.method == 'POST':
        # create a form instance and populate it with data from the request:
        form = TokenForm(request.POST)
        # check whether it's valid:
        if form.is_valid():
            # process the data in form.cleaned_data as required
            # ...
            # redirect to a new URL:
            # print(form)
            return HttpResponseRedirect('/thanks/')

    # if a GET (or any other method) we'll create a blank form
    else:
        form = TokenForm()

    return render(request, 'passkeep/add.html', {'form': form})

def allToken(request):
    resp = TokenInfo.objects.all()
    template = loader.get_template('passkeep/all.html')
    context = {
        'token_list': resp,
    }
    return HttpResponse(template.render(context, request))
    # if this is a POST request we need to process the form data
    if request.method == 'POST':
        # create a form instance and populate it with data from the request:
        form = TokenForm(request.POST)
        # check whether it's valid:
        if form.is_valid():
            # process the data in form.cleaned_data as required
            # ...
            # redirect to a new URL:
            return HttpResponseRedirect('/thanks/')

    # if a GET (or any other method) we'll create a blank form
    else:
        form = TokenForm()

    return render(request, 'name.html', {'form': form})


def detail(request, token_id):
    if request.method == 'PUT':
        return HttpResponse("You're posting a token")
    elif request.method == 'DELETE':
        pass
    return HttpResponse("You're looking at question %s." % token_id)

class TokenViewSet(viewsets.ModelViewSet):
    queryset = TokenInfo.objects.all().order_by('id')
    serializer_class = TokenInfoSerializer