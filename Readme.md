## Installation

- `git clone https://github.com/Taufik0101/CC-Nicepay.git`
- `cd CC-Nicepay/`
- `go get ./...`
- `go run server.go

## Rule Registration CC
- Post Data yang sudah tersedia di postman

## Rule Inquiry CC (Lakukan Setelah Registrasi dan setelah Payment)
- Ganti params request timestamp, txid, referenceno, amt sesuai dengan response registrasi CC

## Rule Payment CC
- Ganti params request timestamp, txid, referenceno, amt sesuai dengan response registrasi CC
- sisanya biarkan saja

## API POSTMAN
https://api.postman.com/collections/19199067-e8e78a70-7b16-45b6-a26f-bc639fd3ec0e?access_key=PMAT-01H58A5V0949NGFQXEEHRTR3SV

## ATAU
import KASPIN.postman_collection.json ke postman
