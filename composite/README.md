# tiktok-composite

This repository contains microservices written using Kitex, with a focus on the intersection of users and videos. The services use ETCD for load balancing and currently aim to implement the following features:

1. Provide a video stream sorted in reverse chronological order by submission time.
2. Support user like video operation.
3. Support user to query for all liked videos.
4. Support user comment video operation.
5. Support user to query a video comment list.
