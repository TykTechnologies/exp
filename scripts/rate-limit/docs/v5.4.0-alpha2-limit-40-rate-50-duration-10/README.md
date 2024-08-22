# Redis rate limit behaviour

## Fixed window

![incoming rate](./fixed-window-rate-in.png)
![latency](./fixed-window-latency.png)
![outgoing rate](./fixed-window-rate-out.png)
![requests](./fixed-window-requests.png)

## Sliding log (sentinel)

![incoming rate](./sentinel-rate-in.png)
![latency](./sentinel-latency.png)
![outgoing rate](./sentinel-rate-out.png)
![requests](./sentinel-requests.png)

## DRL (non-redis)

![incoming rate](./drl-rate-in.png)
![latency](./drl-latency.png)
![outgoing rate](./drl-rate-out.png)
![requests](./drl-requests.png)