from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from datetime import datetime
from typing import Optional
import random
import string
import pymysql
import subprocess

app = FastAPI()

origins = ["*"]

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)

# Database connection settings
db_settings = {
    "host": "ec2-3-37-36-76.ap-northeast-2.compute.amazonaws.com",
    "port": 3306,
    "user": "eeta",
    "password": "eeta1@",
    "database": "eeta",
}

# Database connection
connection = pymysql.connect(**db_settings)

# Model for the data
class Deposit(BaseModel):
    id : Optional[int] = None
    address: Optional[str] = None
    bank: Optional[str] = None
    account: Optional[str] = None
    amount: Optional[int] = None
    deposit: Optional[int] = None
    state : Optional[int] = None
    event_time : Optional[datetime] = None

# CRUD operations

def generate_random_string(length):
    characters = string.digits + '-'
    return ''.join(random.choice(characters) for _ in range(length))

def create_deposit(deposit: Deposit):
    with connection.cursor() as cursor:
        random_bank_account = generate_random_string(14)
        now = datetime.now()
        sql = "INSERT INTO deposits (address, bank, account, amount, deposit, state, event_time) VALUES (%s, %s, %s, %s, %s, %s, %s)"
        cursor.execute(sql, (deposit.address, deposit.bank, random_bank_account, deposit.amount, deposit.deposit, deposit.state, now))
        connection.commit()

    with connection.cursor() as cursor:
        sql = "SELECT LAST_INSERT_ID();"
        cursor.execute(sql)
        id = cursor.fetchone()[0]
        return {"id": id, "address": deposit.address, "bank": deposit.bank, "account": random_bank_account, "amount": deposit.amount, "deposit": deposit.deposit, "state": deposit.state, "event_time": now}

def get_deposit(address: str):
    with connection.cursor() as cursor:
        sql = "SELECT * FROM deposits WHERE address = %s"
        cursor.execute(sql, (address,))
        deposits = cursor.fetchall()
        return [{"id": deposit[0], "address": deposit[1], "bank": deposit[2], "account": deposit[3], "amount": deposit[4], "deposit": deposit[5], "state": deposit[6], "event_time": deposit[7]} for deposit in deposits]

def get_all_deposits():
    with connection.cursor() as cursor:
        sql = "SELECT * FROM deposits"
        cursor.execute(sql)
        deposits = cursor.fetchall()
        return [{"id": deposit[0], "address": deposit[1], "bank": deposit[2], "account": deposit[3], "amount": deposit[4], "deposit": deposit[5], "state": deposit[6], "event_time": deposit[7]} for deposit in deposits]


def update_deposit_fn(id: int, deposit: int):
    with connection.cursor() as cursor:
        sql = "UPDATE deposits SET deposit = %s WHERE id = %s"
        cursor.execute(sql, (deposit, id))
        connection.commit()
    with connection.cursor() as cursor:
            sql = "SELECT address,deposit FROM deposits WHERE id = %s"
            cursor.execute(sql, (id,))
            result = cursor.fetchone()
            address = result[0]
            deposit = result[1]
            cmd = "/home/ubuntu/eetad tx deposit mint " + address + " " + str(deposit) + "000000krw --from u1 -y"
            result = subprocess.check_output(cmd, shell=True)
            return {"id": id, "deposit": deposit}

def update_state_fn(id: int, state: int):
    with connection.cursor() as cursor:
        sql = "UPDATE deposits SET state = %s WHERE id = %s"
        cursor.execute(sql, (state, id))
        connection.commit()
    with connection.cursor() as cursor:
        sql = "SELECT address,deposit FROM deposits WHERE id = %s"
        cursor.execute(sql, (id,))
        result = cursor.fetchone()
        address = result[0]
        deposit = result[1]
        cmd = "/home/ubuntu/eetad tx deposit mint " + address + " " + str(deposit) + "000000krw --from u1 -y"
        result = subprocess.check_output(cmd, shell=True)
        return {"id": id, "state": state}

# API Endpoints

@app.post("/Deposit/", response_model=Deposit)
def create_new_deposit(deposit: Deposit):
    return create_deposit(deposit)

@app.get("/Deposit/", response_model=list[Deposit])
def read_all_deposits():
    return get_all_deposits()

@app.get("/Deposit/{address}", response_model=list[Deposit])
def read_deposit(address: str):
    return get_deposit(address)

@app.put("/Deposit/{id}", response_model=Deposit)
def update_deposit(id: str, deposit: int):
    return update_deposit_fn(id, deposit)

@app.put("/Issue/{id}", response_model=Deposit)
def update_state(id: str, state: int):
    return update_state_fn(id, state)
