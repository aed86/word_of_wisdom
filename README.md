<h1>
"Word of wisdom" tcp server 
</h1>

<h3>How to start:</h3>
- `make build` - build docker containers
- `make start` - run client

<h3>Some notes:</h3>
I didn't write any tests just to save the time. If it is necessary - I can do.

Pow algorithm has been chosen as sha256 because:
1) It is used by bitcoin 
2) more secure 
3) more time-consuming (we have to protect our api from ddos, because of that we are interested in time-consuming process of pow)


