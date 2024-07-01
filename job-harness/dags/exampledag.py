from airflow import Dataset
from airflow.decorators import dag, task
from pendulum import datetime
import requests
from airflow.providers.discord.notifications.discord import DiscordNotifier


@dag(
    start_date=datetime(2024, 1, 1),
    schedule="@daily",
    catchup=False,
    doc_md=__doc__,
    default_args={"owner": "Astro", "retries": 3},
    tags=["example"],
)
def example_astronauts():
    send_message = DiscordNotifier(
                discord_conn_id = "https://discord.com/api/webhooks/1256991119091175536/rPbkuFfN64IVt06lU9NXuHMrkIhF95Mv3TtmInCmcHn6BdMeNUO_DQawmfne1_a4IzFc",
                text = "It works!"
                ) 

    send_message


example_astronauts()
