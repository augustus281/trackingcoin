package util

const ListingTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Subject}}</title>
</head>
<body>
    <h1><strong>{{.Subject}}</strong></h1>
    <h2>Hello, {{.RecipientName}}!</h2>
    <p>Here is the latest cryptocurrency data:</p>
    <table border="1">
        <tr>
            <th>Name</th>
            <th>Symbol</th>
            <th>Price (USD)</th>
            <th>Last Updated</th>
        </tr>
        {{range .ListingData.Data}} <!-- Correct path to the Data array -->
        <tr>
            <td>{{.Name}}</td>
            <td>{{.Symbol}}</td>
            <td>{{.Quote.USD.Price}}</td>
            <td>{{.LastUpdated}}</td>
        </tr>
        {{end}}
    </table>
    <p>Total Cryptocurrencies: {{.ListingData.Status.TotalCount}}</p>
    <p>Data fetched on: {{.ListingData.Status.Timestamp}}</p>
</body>
</html>
`
