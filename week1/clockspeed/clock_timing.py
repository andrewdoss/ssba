ITER = 1000000000


def test_loop():
    for i in range(ITER):
        pass


def test_loop_add():
    x = 0
    for i in range(ITER):
        x += 1
    return x


def test_loop_mult():
    x = 0
    for i in range(ITER):
        x = i * 146
    return x


def test_loop_divide():
    x = 0
    for i in range(ITER):
        x = i / 146
    return x 