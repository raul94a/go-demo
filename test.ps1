$body = @{
    id = "IDTEST_NRHERE"
    title = "YOUR_TITLE"
    artist = "YOUR_ARTIST"
    price = 55.66
} | ConvertTo-Json

$headers = @{
    "Content-Type" = "application/json"
}

$url = "http://localhost:8080/albums"

$response = Invoke-RestMethod -Uri $url -Method Post -Body $body -Headers $headers

# Print the response
$response



$response2 = Invoke-RestMethod -Uri $url -Method Get
$response2

