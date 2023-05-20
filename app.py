def BubbleSort(arr):
    for i in range(len(arr)):
        for j in range(len(arr)-1-i):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j+ 1] = arr[j+1], arr[j]
    print(arr)


def insertionSort(arr):
    for i in range(len(arr)):
        smallestidx = i
        for j in range(i,len(arr)):
            if arr[j] < arr[smallestidx]:
                smallestidx = j
        temp = arr[i]
        arr[i] = arr[smallestidx]
        arr[smallestidx] = temp
    print(arr)

insertionSort([3,5,6,2,4, 8, 11, 1,10])

def binarySearch(arr, item):
    for i in range(len(arr)):
        if arr[i] == item:
            return item




