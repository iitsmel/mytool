@echo off
go build
set /P id=Enter id: 
runas /user:%id% cmd
rem del "C:\Program Files\Go\bin\mytool.exe"
rem copy "mytool.exe" "C:\Program Files\Go\bin"