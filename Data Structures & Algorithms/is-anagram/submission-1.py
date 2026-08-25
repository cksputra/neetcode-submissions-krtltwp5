class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        
        if len(s) != len(t):
            return False
        
        sCount = {}
        tCount = {}

        for i in range(0, len(s)):
            if s[i] not in sCount:
                sCount[s[i]]=1
            else:
                sCount[s[i]]+=1

            if t[i] not in tCount:
                tCount[t[i]]=1
            else:
                tCount[t[i]]+=1

        for key in sCount.keys():
            if key not in tCount or sCount[key] != tCount[key]:
                return False
            

        return True
        