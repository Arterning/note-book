
```http
### Remote call
< {%
 const appId = 'sa7770d975ededda47';
 const key = '57797b7f19edf30b15cff4ae5c6f7667';
 const timeStamp = new Date().getTime().toString()
 const signature = crypto.hmac.sha256()
 .withTextSecret()
 .updateWithText(appId + key + timeStamp)
 .digest().toHex();
 request.variables.set("appId", appId);
 request.variables.set("signature", '029d840e953446bd54005c4b48553a01f669c7a171f16867310a595ba42b524b');
 request.variables.set("timeStamp", '1687939098614');
%}
GET https://emng-test.zoomlion.com/algorithm/auth/userRole/v1/getRolePermission?roleId=10557968154624
sa-app-id:{{appId}}
sa-timestamp:{{timeStamp}}
sa-sign:{{signature}}
```

sign = sh256(appid + appkey + timestamp)

_application programming interface_ (_API_)
