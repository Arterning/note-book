
raw form data is 
```js
[
    { name: 'id', value: '658bd10a620f54fd0c24c74d' },
    { name: 'title', value: '' },
    { name: 'description', value: '' },
    { name: 'date', value: '' }
  ]
```

we should convert it into object 

```js
Object.fromEntries(formData);
```

the result is 
```js
{
  id: '658bd10a620f54fd0c24c74d',
  title: '',
  description: '',
  date: ''
}
```