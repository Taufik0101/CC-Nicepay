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

## IMPORT POSTMAN
import KASPIN.postman_collection.json ke postman
## ATAU
import url dalam file API KASPIN CC.txt
