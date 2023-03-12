# blended benchmark

here we should implement a react formik form

```http

POST http://domain/blended-benchmarks

{
    name: 'name',
    symbol: 'symbol',
    createdDate: 'date',
    modifiedDate: 'date',
    components: [
        {
            benchmarkId: 'id',
            weight: 80,
            date: 'date'
        },
        {
            benchmarkId: 'id',
            weight: 20,
            date: 'date'
        }
    ]
}

PATCH http://domain/blended-benchmarks

GET http://domain/blended-benchmarks

DELETE http://domain/blended-benchmarks
```

# investment policy

here we should implement a tree model select in react and a formik react form

```http
POST http://domain/investment-policy

{
    name: 'name',
    symbol: 'symbol',
    createdDate: 'date',
    modifiedDate: 'date',
    rules: [
        {
            asset: 'id',
            target: 50,
            min: 20,
            max: 100,
            date: 'date'
        },
        {
            asset: 'id',
            target: 50,
            min: 20,
            max: 100,
            date: 'date'
        }
    ]
}

```

# singel-transaction-model

here we will create a transaction system, implement Fund,Account,Transaction,TransactionType, Person,Orgnization Entity
and a management system

```http
POST http://domain/transaction-types
{
    name: 'name',
    type: 'inflow',
    fee:  '1%',
}
```

# question

should we use graphql or json api as api service?
