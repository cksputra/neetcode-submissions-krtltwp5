class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        duplicateDict = {}

        for n in nums:
            if n in duplicateDict:
                return True
            else:
                duplicateDict[n] = True
        
        return False
        