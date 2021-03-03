# BinanceOrdersInfo

## Project to getting important informations about your oders from Bianace: https://www.binance.com/

- Problem to be solved: if you make many Buy/Sell orders daily in different prices, it'll be very hard for you to calculate your the average price in $ 

- Steps: all you need is getting your api key and api secret and add them to .env file 

- API Exmaple: GET /currency=sfp

```json
Response: {
    "Currency": "sfp",
    "Quantity": 183.48999999999998,
    "AveragePrice": 2.5123126491906924
}
```
