import threading
from ftx.client import FtxClient
from staking import auto_staking
from settings import API, SECRET, SUBACCOUNT


def main() -> None:
    threading.Timer(60 * 60.0, main).start()
    client = FtxClient(api_key=API,
                       api_secret=SECRET,
                       subaccount_name=SUBACCOUNT)
    auto_staking(client)


if __name__ == '__main__':
    main()
