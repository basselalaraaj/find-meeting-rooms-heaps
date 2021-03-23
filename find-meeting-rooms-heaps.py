import heapq
from typing import List


def minMeetingRooms(meetingTimes: List[List[int]]) -> int:
    if len(meetingTimes) <= 0:
        return 0

    freeRooms: List[int] = []

    meetingTimes.sort(key=lambda x: x[0])

    heapq.heappush(freeRooms, meetingTimes[0][1])

    for meeting in meetingTimes[1:]:
        if freeRooms[0] <= meeting[0]:
            heapq.heappop(freeRooms)
        heapq.heappush(freeRooms, meeting[1])

    return len(freeRooms)


meetingTimes: List[List[int]] = [
    [3, 4], [3, 9], [2, 8], [5, 11], [8, 20], [11, 15]]
print(minMeetingRooms(meetingTimes))
