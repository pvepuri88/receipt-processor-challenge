# Receipt Processor Challenge

A Go-based microservice that calculates reward points for receipts.  
Stores data in-memory and exposes two endpoints for processing and retrieving points.

---

## Prerequisites

- Go 1.20 (or higher) installed

---

# Installation

## 1. **Clone** the repository:
   ```bash
   git clone https://github.com/pvepuri88/receipt-processor-challenge.git
   cd receipt-processor-challenge
```

## 2. Download any dependencies:
```bash
go mod tidy
```  
  
## 3. Running the Service:
```bash
go run main.go
```
The service starts on localhost:8080


## 4. Curl commands create a receipt
```bash
curl -X POST http://localhost:8080/receipts/process \
  -H "Content-Type: application/json" \
  -d '{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
      { "shortDescription": "Mountain Dew 12PK", "price": "6.49" },
      { "shortDescription": "Emils Cheese Pizza", "price": "12.25" },
      { "shortDescription": "Knorr Creamy Chicken", "price": "1.26" },
      { "shortDescription": "Doritos Nacho Cheese", "price": "3.35" },
      { "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ", "price": "12.00" }
    ],
    "total": "35.35"
  }'
```
## 5. Curl command to fetch points (replace id with the actual ID returned by the POST command
```bash
curl http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points
```


