GET /payment/v1/account

{
  "purchased": false,
  "notification": false,
  "plans": [
    {
      "plan_id": "2736123def7232",
      "type": "normal",
      "price": 10,
      "bill_period": "monthly",
      "description": "1 CPU Core, 512M Memory"
    }
  ],
  "status": "",
  "balance": {
    "balance": 0
  }
}

GET /payment/v1/market

{
  "plans": [
    {
      "plan_id": "1d3452ea-7f14-11e6-9fe0-2344dd5557c3",
      "type": "C",
      "price": 20,
      "bill_period": "M",
      "description": "1 CPU Core, 512M Memory",
      "creation_time": "2016-09-26T17:27:04+08:00"
    },
    {
      "plan_id": "1d3452ea-7f14-11e6-9fe0-2344dd5557c3",
      "type": "C",
      "price": 40.88,
      "bill_period": "M",
      "description": "2 CPU Cores, 1G Memory",
      "creation_time": "2016-09-26T17:27:04+08:00"
    },
    {
      "plan_id": "1d3452ea-7f14-11e6-9fe0-2344dd5557c3",
      "type": "C",
      "price": 88.88,
      "bill_period": "M",
      "description": "4 CPU Cores, 2G Memory",
      "creation_time": "2016-09-26T17:27:04+08:00"
    }
  ]
}

POST /payment/v1/checkout -d '{"plan_id":"zwwqe","namespace":"chaizs","region":"cn-north-1"}'

{
  "balance": 3000.01,
  "status": "active"
}

POST /payment/v1/recharge -d '{"amount":1234.34,namespace:"chaizs"}'

{
  "balance": 6000.89,
  "status": "active"
}

GET /payment/v1/amount
[
  {
    "trans_id": "03Fwerqe2",
    "creation_time": "2016-09-26T17:30:03+08:00",
    "amount": 12.23,
    "description": "Plan A",
    "payment_method": "balance",
    "status": "finish"
  },
  {
    "trans_id": "qwer238DKJ",
    "creation_time": "2016-09-26T17:30:03+08:00",
    "amount": 12.23,
    "description": "Plan A",
    "payment_method": "balance",
    "status": "finish"
  },
  {
    "trans_id": "03F232X238DKJ",
    "creation_time": "2016-09-26T17:30:03+08:00",
    "amount": 2.34,
    "description": "Plan A",
    "payment_method": "balance",
    "status": "refunded"
  }
]

GET /payment/v1/amount/we

{
  "trans_id": "03F232X238DKJ",
  "creation_time": "2016-09-26T17:30:04+08:00",
  "amount": 12.23,
  "description": "Plan A",
  "payment_method": "balance",
  "status": "finish"
}

GET /payment/v1/balance

{
  "balance": 50000.89,
  "status": "active"
}
