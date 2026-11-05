$root = 'C:\Users\FEMI\gopherdojo'
$keep = (Get-Content (Join-Path $root 'exercises.json') -Raw | ConvertFrom-Json).name
foreach ($r in @('exercises', 'solutions', '.pristine')) {
    Get-ChildItem (Join-Path $root $r) -Directory | ForEach-Object {
        if ($keep -notcontains $_.Name) {
            Write-Host "Removing $r\$($_.Name)"
            Remove-Item -Recurse -Force $_.FullName
        }
    }
}
