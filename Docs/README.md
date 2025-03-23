# Rubik_cube

- *The repository will contain code that will be used to solve the rubik's cube. The solution will be implemented in different languages*

## Milestones

### M1

- *BUild the models that define a rubik's cube.*
- *Add models that define 2/3 types of cubits in the Rubik's cube*
-

---

## Algorithm

- _The algorithm used in this program was single-handedly programmed and tested by [Anxiel Ray](https://github.com/anxielray) and now lemme show you how I did it._
- _I will beak down the content of the tutorial into 3 parts;_
  - _Botttom layer Algorithms_
  - _Middle layer Algorithms_
  - _Top layer Algorithms_

### _Bottom Layer_

- _Just like most algorithms used to solve the Rubik's cube, mine also stems from the common root, the botom layer._
- _We will try to achieve a cross, for the color? Thats all up to you!_
- _Start by any algorithm for the search of the mid-cubit faces that match the color you just picked._
- _Next we let's make sure  we have a cross, if you happend to have formed a corner as well, that doesn't really matter._
- _Once we have all 4 cardinal cubits in place, we want to confirm their alignment to their prescedent `face-rulers`{The center cubits of the faces, the mid-cubits fall in, apart from your selected faces} and match them!_
- _After that, check for the corner cubit that fits in a specific corner, and place it right under its correct position `(in the opposite face)`. After that, perform a series of the commands {U, R', U' R}. Do this until the corner cubit is aligned perfectly. Do the same for all the other corner cubuits_
- _If you are keen! The bottom layer should be already solved by now!_
- _Congratulations! you mastered the algorithm to solve any face in the cube independently!_

---

### _Middle Layer_

- _Now let us handle the middle layer_
