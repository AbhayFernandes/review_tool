$repo = "AbhayFernandes/review_tool"
$response = Invoke-RestMethod "https://api.github.com/repos/$repo/releases/latest"
$url = $response.assets | Where-Object { $_.name -like "*windows*" } | Select-Object -ExpandProperty browser_download_url
Invoke-WebRequest -Uri $url -OutFile "$env:USERPROFILE\crev.exe"
Write-Output "Installed mybinary version $($response.tag_name)"
Write-Output "Restart your powershell instance/terminal to start using it!"
