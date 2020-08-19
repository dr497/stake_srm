FROM python:latest

WORKDIR /usr/src/stake_srm_python

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

CMD [ "python", "./main.py" ]
