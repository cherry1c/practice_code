
from enum import EnumMeta


def enumerateDemo():
    words  = ["hello", "world", "you", "__{}__".format(0)]
    for i, word in enumerate(words):
        print("i: ", i, " word: ", word)

    s = ""
    s = s.join(words)
    print(s)
    index = s.find("__1__")
    print(index)


def printUnicode(text):
    '''
    a=\u767e\u5ea6\u5728\u7ebf\u7f51\u7edc\u6280\u672f\uff08\u5317\u4eac\uff09\u6709\u9650\u516c\u53f8
>>> print unicode(a,'utf-8')
\u767e\u5ea6\u5728\u7ebf\u7f51\u7edc\u6280\u672f\uff08\u5317\u4eac\uff09\u6709\u9650\u516c\u53f8
>>> print(a.encode("utf-8").decode('unicode_escape')) 
    '''
    print(text.encode("gbk").decode('unicode_escape'))
    print(text)

class setValue:
    def __init__(self) -> None:
        self.hello = 0
        self.world = ""

    def testDict(self):
        d = {
            "hello": lambda value: setattr(self, "hello", int(value)),
            "world": lambda value: setattr(self, "world", value),
        }
        print(self.hello)
        d["hello"](100)
        print(self.hello)


if __name__ == '__main__':
    a = setValue()
    a.testDict()