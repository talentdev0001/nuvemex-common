# Details

## About

Provide item details by Elasticsearch ID

**Method:** GET<br>
**Access:** Public<br>
**Query Parameters:** <br>

- `partId`: elasticsearch key <br><br>

## Use cases

- Before item added to cart, details query made and product variants created at vendure.

## Dependencies

- Elasticsearch
- SQS
- DynamoDB

## Response

`/details?partId=MSR_68efccbbdc0876989cafc5eca1c955ed`

[Sample link](https://parts.stg.nuvemex.com/details?partId=MSR_68efccbbdc0876989cafc5eca1c955ed)

```json
{
  "id": "MSR_68efccbbdc0876989cafc5eca1c955ed",
  "partNum": "M81044/12-16-9",
  "supplierPartNum": "650-M81044/12-16-9",
  "manufacturer": "TE Connectivity / Raychem",
  "supplier": "mouser",
  "region": "none",
  "description": "Hook-up Wire 16AWG 19x29 WH PRICE PER FT",
  "dateCode": "none",
  "inStock": 7649,
  "minimumBuy": 1,
  "multiple": 1,
  "price": [
    {
      "quantity": 1,
      "price": 0.55,
      "currency": "USD"
    },
    {
      "quantity": 5,
      "price": 0.453,
      "currency": "USD"
    }
  ],
  "uploadedAt": "2021-01-24T22:31:56.747342446+01:00"
}
```
