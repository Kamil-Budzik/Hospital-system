# scripts/run-all.sh
#!/bin/bash

# Start user service
cd services/user-service
air -c .air.toml &

# Start patient service
cd ../patient-service
air -c .air.toml &

# Start appointment service
cd ../appointment-service
air -c .air.toml &

# Start auth service
cd ../auth-service
air -c .air.toml &

# Wait for all background processes
wait
