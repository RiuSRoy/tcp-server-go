# tcp-server-go


## Build a tcp server in go from scratch

https://www.youtube.com/watch?v=f9gUFy-9uCM&ab_channel=AsliEngineeringbyArpitBhayani


## Improvements
1. Limiting the number of threads
    - We cannot just keep on spinning a new thread for every client request
    - Because every thread eats up memeory (heap) and CPU allocation time
2. Add thread pool tp save on thread creation time and
    - There will be a thread pool of a finite no of threads, and for every request, we pick up a thread from this pool
3. Connection timeout
    - What if a client does not respond? Our thread is unncessarily occupied, so connection timeout is a must
4. TCP backlog queue configuration