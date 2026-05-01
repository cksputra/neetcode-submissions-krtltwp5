type TimeMap struct {
	storageT map[string][]int
	storageV map[string]string
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Constructor() TimeMap {
	storageT := make(map[string][]int, 0)
	storageV := make(map[string]string)
	return TimeMap{
		storageT: storageT,
		storageV: storageV,
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.storageT[key] = append(this.storageT[key], timestamp)
	
	valueKey := fmt.Sprintf("%s%d", key, timestamp)

	this.storageV[valueKey] = value
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.storageT[key]

	if len(arr) == 0 {
		return ""
	}

	l, r := 0, len(arr) - 1

	prevTimestamp := 0 

	for l <= r {
		mid := l + (r-l) / 2


		if arr[mid] == timestamp {
			prevTimestamp = arr[mid]
			break
		}

		if arr[mid] > timestamp {
			r = mid - 1
		}

		if arr[mid] < timestamp {
			prevTimestamp = arr[mid]
			l = mid + 1
		}
	}


	if timestamp < prevTimestamp {
		return ""
	}

	valueKey := fmt.Sprintf("%s%d", key, prevTimestamp)

	return this.storageV[valueKey]
}




