from django.db import models

# Create your models here.
class TokenInfo(models.Model):
    url = models.CharField(max_length=500)
    name = models.CharField(max_length=500)
    token = models.CharField(max_length=1000)
    created_date = models.DateTimeField('date created')
    modified_date = models.DateTimeField('date modified')

