# Automatically stake SRM on FTX

Every hour [FTX](https://ftx.com) airdrops SRM staking rewards on people staking directly from their [FTX](https://ftx.com) account. However, there is no option to automatically compound these tokens. 

This repository allows you to automatically stake your staking rewards every hour. The code is available in Go and Python.

# Project Serum
<p align="center">
<img height="200" src="/assets/logo-serum.png">
</p>
Serum is the world's first completely decentralized derivatives exchange with trustless cross-chain trading brought to you by Project Serum, in collaboration with a consortium of crypto trading and DeFi experts.

- To find more information about Project Serum: [https://projectserum.com](https://projectserum.com)

- To find more information about SRM staking: [https://projectserum.com/staking-voting](https://projectserum.com/staking-voting)

- SRM can be staked directly on FTX: [https://ftx.com/staking](https://ftx.com/staking)



# Deploy with Docker

<p align="center">
<img height="100" src="/assets/logo-docker.png">
</p>

The app is available in Go and Python, both can be deployed with Docker

## Go

<p align="center">
<img height="100" src="/assets/logo-go.png">
</p>

```
docker build -t user/stake_srm_go .  
docker run -d --restart always user/stake_srm_go:latest
```


## Python

<p align="center">
<img height="100" src="/assets/logo-python.png">
</p>

```
docker build -t user/stake_srm_python .  
docker run -d --restart always user/stake_srm_python:latest
```

