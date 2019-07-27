// 二分查找法
function binarySearch(arr, item) {
    var low = 0;
    var high = arr.length - 1;

    if (item < arr[low] || item > arr[high]) {
        return null;
    }

    while (low <= high) {
        var mid = (low + high) / 2;
        mid = Math.round(mid);

        var guess = arr[mid];

        if (guess == item) {
            return mid;
        }
        if (guess < item) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }

}
var arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
console.log(binarySearch(arr, 3));