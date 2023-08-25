#! /bin/sh

# Start first process. Golang API.
cd /app/bin/
./Twitcher &

# Start second process. Next.js front-end.
cd /app/src/
npm run start &

# Wait for any process to exit
wait -n
  
# Exit with status of process that exited first
exit $?