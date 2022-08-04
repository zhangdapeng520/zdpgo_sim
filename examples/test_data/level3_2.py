def sub(a, b):
    """
    我还有自己自定义的方法
    """
    return a - b


# 我在自己的代码中复制了另一个代码中的某个函数，并进行了修改
def add(aaa, bbbb):
    """
    加法
    :param a: 数字1
    :param b: 数字2
    :return: 数字1 与 数字2 相加的结果
    """

    # 计算和并返回
    return aaa + bbbb


if __name__ == '__main__':
    print(add(11, 22))  # 测试1
    print(add(111, 222))  # 测试1
    print(add(1111, 2222))  # 测试3
    print(sub(11, 22))  # 测试1
    print(sub(111, 222))  # 测试1
    print(sub(1111, 2222))  # 测试3
