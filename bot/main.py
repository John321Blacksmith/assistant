import os
import grpc
from dotenv import load_dotenv
from proto import classifier_pb2, classifier_pb2_grpc
from vkbottle.bot import Bot, Message


load_dotenv()


token = os.getenv("BOT_TOKEN")
classifierIP = os.getenv("GRPC_IP")
classifierPort = os.getenv("GRPC_PORT")

if not token:
    raise RuntimeError("Bot token wasn't provided")


bot = Bot(token=token)


@bot.on.message()
async def handle_message(message: Message):
    response: str

    with grpc.insecure_channel(f"{classifierIP}:{classifierPort}") as ch:
        clientStub = classifier_pb2_grpc.ClassifierStub(ch)
        request = classifier_pb2.BotRequest(input=message.text)
        response = clientStub.GetMainContext(request)

    return await message.answer(message=f'you are talking about {response}')


if __name__ == '__main__':
    bot.run()