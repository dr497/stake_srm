from ftx.client import FtxClient


COINS = ['SRM', 'SRM_LOCKED', 'MSRM', 'MSRM_LOCKED', 'SOL', 'RAY']


def auto_staking(client: FtxClient) -> None:
    for coin in COINS:
        stake_coin(client, coin)


def stake_coin(client: FtxClient, coin: str) -> None:
    balance_coin = client.get_balance_coin(coin)
    if balance_coin is None:
        return
    available_balance = balance_coin['free']
    if available_balance > 0:
        client.stakes(coin=coin, size=available_balance)
