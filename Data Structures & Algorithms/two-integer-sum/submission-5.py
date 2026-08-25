class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numsDict = {}

        for i in range(0, len(nums)):
            remainder = target - nums[i]

            if remainder in numsDict:
                return [numsDict[remainder], i]

            numsDict[nums[i]] = i
        return []