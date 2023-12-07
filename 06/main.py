import math


s = open('input.txt').readlines()


times = [int(s) for s in ' '.join(s[0].split()).split()[1:]]
distances  =[int(s) for s in ' '.join(s[1].split()).split()[1:]]

zipped = dict(zip(times,distances))

ways_list = []

for t,d in zipped.items():
    p2_answer = 0
    for i in range(t):
        total = i*(t-i)
        if total > d:
            p2_answer+=1
    ways_list.append(p2_answer)

p1_answer = 1
for p2_answer in ways_list:
    p1_answer *= p2_answer

print(p1_answer)

# -hold_time^2 - 63789468 * hold_time - 411127420471035
duration = 63789468
record = 411127420471035

def f(hold):
    return hold*(duration-hold)

a = -1
b = -duration
c = -record

root1 = (-b+math.sqrt(b**2 - 4*a*c))/(2*a)
root2 = (-b-math.sqrt(b**2 - 4*a*c))/(2*a)

# Rounding errors?????????
print(root1, root2, max(root1, root2) - min(root1, root2))

# Brute force it, I guess
p2_answer = 0
for i in range(duration):
    if i*(duration-i) > record:
        p2_answer += 1
print(p2_answer)