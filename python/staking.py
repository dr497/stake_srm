from ftx.client import FtxClient
import logging

logger = logging.getLogger(__name__)

SRM_COINS = ['SRM', 'SRM_LOCKED', 'MSRM', 'MSRM_LOCKED']


def auto_staking(client: FtxClient) -> None:
    for coin in SRM_COINS:
        stake_coin(client, coin)


def stake_coin(client: FtxClient, coin: str) -> None:
    balance_coin = client.get_balance_coin(coin)
    if balance_coin is None:
        return
    available_balance = balance_coin['free']
    if available_balance > 0:
        client.stakes(coin=coin, size=available_balance)
