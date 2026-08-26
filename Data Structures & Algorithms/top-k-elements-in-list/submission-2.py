class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        
        freq = {}

        for num in nums:
            if num not in freq:
                freq[num] = 1
            else:
                freq[num] += 1


        arr = []
        for key, value in freq.items():
            arr.append([value, key])

        arr.sort()

        res = []

        for i in range (0, k):

            index = (i+1) * -1
            res.append(arr[index][1])

        return res

        
        
        