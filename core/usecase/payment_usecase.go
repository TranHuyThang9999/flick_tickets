package usecase

const curlCheckPayment = `curl --location 'https://api-merchant.payos.vn/v2/payment-requests/{id}' \
--header 'x-client-id: c84c857d-160c-456a-91f2-384526d7a360' \
--header 'x-api-key: f74461b1-d7d3-4fca-b918-fcb39524ce8c' \
--header 'Cookie: connect.sid=s%3Av-Fr0YxpRhIOdh934PJxPZpITtJ10JuJ.3tGHnEaRmgMIzNk6pR27NKm3VeVbG%2FCnkUOvaBWl5Ss'`
