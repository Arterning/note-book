from django import forms

class TokenForm(forms.Form):
    url = forms.CharField(label='url', max_length=500)
    name = forms.CharField(label='name', max_length=500)
    token = forms.CharField(label='token', max_length=500)