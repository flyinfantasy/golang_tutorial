@echo off

for /f "tokens=*" %%a in (list.txt) do (
  curl -d %%a -X POST http://172.105.193.78:8282/api/v1/pigeons/
)
pause