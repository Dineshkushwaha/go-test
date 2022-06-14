# go-test
### Step 1:  Run a Redis server using :  docker run --name redis-test-instance -p 6379:6379 -d redis
### Step 2: ./mockredis-cli 127.0.0.1:6379
### Step 3: To set the values
SET key value
### Step 4: To get the value
DUMP key
### Step 5: To rename the key
RENAME key1 key2
### Step 6: To exit the cli
EXIT
