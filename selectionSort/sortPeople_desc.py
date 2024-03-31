class Solution:
    def findBigger(self, arr: List[int]) -> int:
        bigger = arr[0]
        bigger_index = 0

        for i in range(1, len(arr)):
            if arr[i] > bigger:
                bigger = arr[i]
                bigger_index = i
        return bigger_index

    def sortPeople(self, names: List[str], heights: List[int]) -> List[str]:
        relaonship = {}
        for i in range(len(names)):
            relaonship[heights[i]] = names[i]

        newOrder = []
        for i in range(len(heights)):
            bigger = self.findBigger(heights)
            newOrder.append(heights.pop(bigger))


        result = []
        for i in range(len(newOrder)):
            name = relaonship[newOrder[i]]
            result.append(name)    
        
        return result