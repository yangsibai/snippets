nohup node server.js > /dev/null 2>&1 &

#1. `nohup` meanus: Do not teminate this process even when the stty is cut off
#2 `>/dev/null` means: stdout goes to /dev/null (which is a dummy device that does not record any output).
#3. `2>&1` means: stderr also goes to the stdout (which is already redirected to /dev/null).
#4. `&` at the end means: run this command as a background task.
#From: <http://stackoverflow.com/questions/4797050/how-to-run-process-as-background-and-never-die>
