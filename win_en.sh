powershell -Command "

Start-Process powershell \"-Command Start-Process -FilePath 'E:\DevEnvironment\Redis\redis-server.exe'\"

Start-Process powershell -ArgumentList \"-Command Set-Location -Path 'E:\DevEnvironment\Nacos\nacos\bin';.\startup.cmd -m standalone\"  -WindowStyle Normal

Start-Process powershell -ArgumentList \"-Command Set-Location -Path 'E:\DevEnvironment\Consul';.\consul agent -dev\" -WindowStyle Normal

Start-Process powershell \"-Command Start-Process -FilePath 'E:\DevEnvironment\ElasticSearch\elasticsearch-8.17.1\bin\elasticsearch.bat'\"
"