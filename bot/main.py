import grpc
from proto import classifier_pb2, classifier_pb2_grpc


def main():
    with grpc.insecure_channel("localhost:50051") as ch:
        clientStub = classifier_pb2_grpc.ClassifierStub(ch)
        while True:
            inp = input(">>> ")
            request = classifier_pb2.BotRequest(input=inp)
            response = clientStub.GetMainContext(request)
            print("response from the target localhost:50051: ", response)


if __name__ == '__main__':
    main()