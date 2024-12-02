import input

reports = input.string.split('\n')
safeReports = 0

for val in reports:
    levels = val.split(' ')
    safeLevels = True
    previous = None
    isPositiveLevel = int(levels[0]) < int(levels[1])
    for level in levels:
        numberLevel = int(level)
        if previous is not None:
            if isPositiveLevel:
                if numberLevel - previous > 3 or numberLevel - previous < 1:
                    safeLevels = False
            else:
                if previous - numberLevel > 3 or previous - numberLevel < 1:
                    safeLevels = False
        previous = int(level)
    if safeLevels:
        safeReports += 1
print(safeReports)