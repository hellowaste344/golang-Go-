import time

start = time.time()


def heavyComputation():
    sum = 0
    for i in range(1, int(1e9)):
        sum += i
    return sum


stop = time.time()

total = heavyComputation()
print(f"Sum: {total}")
print(f"Elapsed time: {(stop - start) * 1000} milliseconds")
