import re


def reFindAll():
    str = 'aabbabaabbaa'
    # 一个"."就是匹配除 \n (换行符)以外的任意一个字符
    print(re.findall(r'a.b',str))#['aab', 'aab']
    # *前面的字符出现0次或以上
    print(re.findall(r'a*b',str))#['aab', 'b', 'ab', 'aab', 'b']
    # 贪婪，匹配从.*前面为开始到后面为结束的所有内容
    print(re.findall(r'a.*b',str))#['aabbabaabb']
    # 非贪婪，遇到开始和结束就进行截取，因此截取多次符合的结果，中间没有字符也会被截取
    print(re.findall(r'a.*?b',str))#['aab', 'ab', 'aab']
    # 非贪婪，与上面一样，只是与上面的相比多了一个括号，只保留括号的内容
    print(re.findall(r'a(.*?)b',str))#['a', '', 'a']
    str = '''aabbab
            aabbaa
            bb'''     #后面多加了2个b
    # 没有把最后一个换行的aab算进来
    print(re.findall(r'a.*?b',str))#['aab', 'ab', 'aab']
    # re.S不会对\n进行中断
    print(re.findall(r'a.*?b',str,re.S))#['aab', 'ab', 'aab', 'aa\n         b']

def textReplace():
    text = '123, word!'
    text1 = text.replace('123', 'Hello')
    print(text1)


def reReplace():
    content = 'abc124hello46goodbye67shit'
    list1 = re.findall(r'\d+', content)
    print(list1)
    mylist = list(map(int, list1))
    print(mylist)
    print(sum(mylist))
    print(re.sub(r'\d+[hg]', 'foo1', content))
    print()
    print(re.sub(r'\d+', '456654', content))



if __name__ == '__main__':
    reReplace()