import grpc
from proto import classifier_pb2, classifier_pb2_grpc


def main():
    text1 = "уч, by, проект, книг, learn, аудитори, inform, знан, class, project, библиотек, класс, librar, book, knowledge, practic"
    text2 = "Я не люблю готовить. Хоть у меня и получается вкусно, мои родители отправили меня за компьютер."
    text3 = "сестра, мам, friend, love, kind, mom, семь, mother, sister, дед, dadfatherbrother, grandmother, друг, сёстры, бабушк, дедушк, party, child, мать, grand, bro, друзья, брат, cousin, grandfather"
    with grpc.insecure_channel("localhost:50051") as ch:
        clientStub = classifier_pb2_grpc.ClassifierStub(ch)
        request = classifier_pb2.BotRequest(input=text1)
        response = clientStub.GetMainContext(request)
        print("response from the target localhost:50051: ", response)


if __name__ == '__main__':
    main()