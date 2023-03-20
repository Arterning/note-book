from rest_framework import serializers

from .models import TokenInfo

class TokenInfoSerializer(serializers.HyperlinkedModelSerializer):

    class Meta:
            model = TokenInfo
            fields = ('url','name','token', 'created_date', 'modified_date')
