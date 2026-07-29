import logging
import grpc
from bot.proto import classifier_pb2, classifier_pb2_grpc

logger = logging.Logger("__bot-logger__")

def main():
    text = "People go to South Korea with a long planning list. See Seoul, eat at the market, buy cosmetics, check up, get to the sea, spend the night in a Buddhist temple, check if the neighborhoods from the dramas really look the same off the screen. Every trip has its own reason. Someone starts with K-pop, with an interest in history or national cuisine."

    with grpc.insecure_channel("localhost:50051") as ch:
        clientStub = classifier_pb2_grpc.ClassifierStub(ch)
        logger.info("sending request to the target localhost:50051")
        request = classifier_pb2.BotRequest(input=text)
        response = clientStub.GetMainContext(request)
        logger.debug("response from the target localhost:50051: ", response)


if __name__ == '__main__':
    main()