::TODO add debug flags

start cmd /c "cd server & node server.js 3000 1"
start cmd /c "cd server & node server.js 3001 2"

start cmd /k "cd balancer & go run ."