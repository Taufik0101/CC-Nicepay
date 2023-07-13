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
https://api.postman.com/collections/19199067-ead104ea-bb70-44dd-a483-9a38ae0a0903?access_key=PMAT-01H560WVTT3N5FXNFSZMJX7BJ0
https://elements.getpostman.com/redirect?entityId=19199067-ead104ea-bb70-44dd-a483-9a38ae0a0903&entityType=collection
