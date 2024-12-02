import input

reports = input.string.split('\n')
safeReports = 0

def checkLevels(levels):
    safeLevels = True
    previous = None
    isPositiveLevel = int(levels[0]) < int(levels[1])
    for level in levels:
        numberLevel = int(level)
        if previous is not None and safeLevels:
            if isPositiveLevel:
                if numberLevel - previous > 3 or numberLevel - previous < 1:
                    safeLevels = False
            else:
                if previous - numberLevel > 3 or previous - numberLevel < 1:
                    safeLevels = False
        previous = int(level)
    return safeLevels

for val in reports:
    levels = val.split(' ')
    safeLevels = checkLevels(levels)

    if not safeLevels:
        for index, levelToDel in enumerate(levels):
            copiedLevels = levels.copy()
            del copiedLevels[index]
            newSafeLevels = checkLevels(copiedLevels)

            if newSafeLevels:
                safeLevels = True

    if safeLevels:
        safeReports += 1
print(safeReports)