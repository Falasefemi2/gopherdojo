$specPath = 'C:\Users\FEMI\gopherdojo\maintainer\SPEC.md'
$lines = Get-Content $specPath
$ex = @()
$inTable = $false
foreach ($line in $lines) {
  if ($line -match '^\|\s*id\s*\|') { $inTable = $true; continue }
  if ($line -match '^\|[-| ]+\|$') { continue }
  if ($inTable -and $line -match '^\|\s*(\d{3})\s*\|\s*([^|]+?)\s*\|\s*([^|]+?)\s*\|') {
    $id = [int]$matches[1]
    $dir = $matches[2].Trim()
    $topic = $matches[3].Trim()
    $ex += [pscustomobject]@{ id = $id; name = $dir; title = $topic; description = ''; track = 'core' }
  }
}
$json = $ex | ConvertTo-Json -Depth 3
[System.IO.File]::WriteAllText('C:\Users\FEMI\gopherdojo\exercises.json', $json, (New-Object System.Text.UTF8Encoding($false)))
Write-Host "Generated $($ex.Count) exercises"
