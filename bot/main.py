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
    output: str
    users_info = await bot.api.users.get(user_ids=message.from_id)
    first_name = users_info[0].first_name

    with grpc.insecure_channel(f"{classifierIP}:{classifierPort}") as ch:
        clientStub = classifier_pb2_grpc.ClassifierStub(ch)
        request = classifier_pb2.BotRequest(input=message.text)
        resp = clientStub.GetMainContext(request)

        if resp:
            output = resp.output

    if output == "greetings":
        return await message.answer(message=f"Hello, {first_name}")

    elif output == "planning":
        return await message.answer(message=f"Okay, let's plan for some day your task")
    else:
        return await message.answer(message=f"You're talking about {output}")


if __name__ == '__main__':
    bot.run()