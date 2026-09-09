import logging
from vkbottle import API
from vkbottle.bot import Bot, Message
from vkbottle.modules import logger

logging.basicConfig(level=logging.INFO)

bot = Bot(token="my_token")


@bot.on.message(text="Hello")
async def send_message(message: Message):
    await message.answer(f"Hello, {message.get_user()}")


if __name__ == "__main__":
    bot.run()