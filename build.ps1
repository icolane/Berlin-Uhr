# Berlin-Uhr Build Script
# Dieses Skript baut den Server und den CLI-Client in den 'build' Ordner.

$BuildDir = "build"

# 1. Build Ordner sicherstellen
if (!(Test-Path $BuildDir)) {
    New-Item -ItemType Directory -Path $BuildDir | Out-Null
    Write-Host "Ordner '$BuildDir' wurde erstellt." -ForegroundColor Cyan
}

Write-Host "Starte Kompilierung..." -ForegroundColor Yellow

# 2. Backend bauen
Write-Host "Baue BerlinServer.exe..." -NoNewline
cd BerlinClock
go build -o ../build/BerlinServer.exe main.go
if ($LASTEXITCODE -eq 0) {
    Write-Host " [OK]" -ForegroundColor Green
} else {
    Write-Host " [FEHLER]" -ForegroundColor Red
}
cd ..

# 3. CLI Client bauen
Write-Host "Baue BerlinCLI.exe..." -NoNewline
cd CliClient
go build -o ../build/BerlinCLI.exe main.go
if ($LASTEXITCODE -eq 0) {
    Write-Host " [OK]" -ForegroundColor Green
} else {
    Write-Host " [FEHLER]" -ForegroundColor Red
}
cd ..

Write-Host "`nBuild abgeschlossen! Deine Dateien liegen in: " -NoNewline
Write-Host (Get-Item $BuildDir).FullName -ForegroundColor Cyan
