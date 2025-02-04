# CONTEXT: SYMULATION
Consider the following simulation: in a two-dimensional area of n x m (dimensions in meters) move in any direction and at a random speed (no more than 2.5[m/s]) i healthy individuals. The speed and direction of their movement can vary randomly over time (but cannot exceed the upper limit set above). Upon reaching any boundary of the area, each individual may:
turn back inside the area (probability 50 percent) leave the area (probability 50 percent)\
During the course of the simulation, new individuals enter the area at random points on its borders (frequency of entry and initial size and select to maintain population continuity). For each entering individual, there is a probability of virus infection of 10 percent.\
Each individual in the population is:\
-resistant to infection\
or (exclusionary alternative)\
-susceptible to infection\
if an individual is susceptible to infection it is:\
-healthy\
or (exclusionary alternative)\
-infected\
if the individual is infected then:\
-has symptoms\
or (exclusionary alternative)\
-does not have symptoms.\
A healthy and non-immune individual becomes infected from an infected individual if and only if: a) the distance between them does not exceed 2[m] and (conjunction) b) the time when this distance is maintained is not less than 3[s] of simulation. The probability of getting infected from an asymptomatic individual is 50 percent, and from an individual with symptomatic disease passage is 100 percent.\
The infected individual sustains the infection for 20 to 30 seconds of simulation after which it recovers, gaining immunity.\
# LABORATORY TASK
Design and implement a solution to simulate the development of an infection in a population. Enable recording and loading of the simulation state at any time t from the start. The solution is to allow visualization of movement and infection. Use vectors from Lab 2 to model the movement of individuals.
### VectorsLib code from lab2: https://pastebin.com/FZLF7ywZ
NOTES\
Each second of the simulation should consist of 25 steps.
