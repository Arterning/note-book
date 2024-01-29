
How to remove %0D from end of URL when using wget?

I have a sh script with the following wget. However, when executed on the linux box, wget is attemping the second URL below (%OD attached). How do i prevent this from happening? I have multiple scripts, they're all having the same issue. Thanks!

```
wget https://example.com/info.repo


wget https://example.com/info.repo%0D
```


The `OD` character is a carriage return, part of the CRLF sequence that Windows uses for line endings just to be different as usual.

You can use [dos2unix](http://linuxcommand.org/man_pages/dos2unix1.html) to fix the line endings before executing, and in future don't use Notepad to write shell scripts.

```
dos2unix myscript.sh
./myscript.sh
```

https://stackoverflow.com/questions/22236197/how-to-remove-0d-from-end-of-url-when-using-wget
