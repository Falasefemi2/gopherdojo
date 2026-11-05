$root = 'C:\Users\FEMI\gopherdojo'
$keep = (Get-Content (Join-Path $root 'exercises.json') -Raw | ConvertFrom-Json).name
foreach ($r in @('exercises', 'solutions', '.pristine')) {
    $missing = @()
    foreach ($n in $keep) {
        $p = Join-Path (Join-Path $root $r) $n
        if (-not (Test-Path $p)) { $missing += $n }
    }
    Write-Host "$r missing: $($missing -join ', ')"
}
