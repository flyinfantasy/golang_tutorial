@echo off
for /f "delims=" %%i in (pigeon_induction20220319.json) do (
  	curl -d '%%i' -H "Content-Type: application/json" -X POST http://127.0.0.1:8282/api/v1/pigeons/
) pause