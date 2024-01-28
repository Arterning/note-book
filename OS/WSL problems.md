
listen EACCES: permission denied in windows
https://stackoverflow.com/questions/59428844/listen-eacces-permission-denied-in-windows

Hi anyone who comes here and don't find any answer. I have an answer that I think will help. it fixed the issue for me. there is cases where nothing is running on the port but some applications or even your code can not run the fix to this is in this order:

1. Open PowerShell as Admin.
    
2. stop winnat with command below:
    

```javascript
$ net stop winnat
```

3. start winnat again with command below:

```javascript
$ net start winnat
```

this fixed the problem for me in many cases. no restart needed. I hope it will help you too.

Edit: I dig in the problem and found out the problem is the port isn't closed in the right way that could be caused easily based on killing the terminal window or the app running on the port.


