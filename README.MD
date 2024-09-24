## push-swap

## Objectives

This project is meant to make a digital version of an ant farm.

Lem-in Program reads from a file (describing the ants and the colony) given in the arguments.

Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

## How does it work?

- Make an ant farm with tunnels and rooms.
- Place the ants on one side and look at how they find the exit.
  Find the quickest way to get n ants across a colony (composed of rooms and tunnels).
- At the beginning of the game, all the ants are in the room ##start. The goal is to bring them to the room ##end with as few moves as possible.
- The shortest path is not necessarily the simplest.
- Some colonies will have many rooms and many links, but no path between ##start and ##end.
- Some will have rooms that link to themselves, sending your path-search spinning in circles. Some will have too many/too few ants, no ##start or ##end, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In those cases the program will return an error message ERROR: invalid data format. If you wish, you can elaborate a more specific error message (example: ERROR: invalid data format, invalid number of Ants or ERROR: invalid data format, no start room found).

Display the results on the standard output in the following format :

number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...

- x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.

- A room is defined by "name coord_x coord_y", and will usually look like "Room 1 2", "nameoftheroom 1 6", "4 6 7".

- The links are defined by "name1-name2" and will usually look like "1-2", "2-5".

Here is an example of this in practice :

```
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
```

Which corresponds to the following representation :

```
        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/
```

## Instructions

- Create tunnels and rooms.
- A room will never start with the letter L or with # and must have no spaces.
- Join the rooms together with as many tunnels as needed.
- A tunnel joins only two rooms together never more than that.
- A room can be linked to multiple rooms.
- Two rooms can't have more than one tunnel connecting them.
- Each room can only contain one ant at a time (except at ##start and ##end which can contain as many ants as necessary).
- Each tunnel can only be used once per turn.
- To be the first to arrive, ants will need to take the shortest path or paths. They will also need to avoid traffic jams as well as walking all over their fellow ants.
- Oonly display the ants that moved at each turn, and each ant can be moved only once and through a tunnel (the room at the receiving end must be empty).
- The rooms names will not necessarily be numbers, and in order.
- Any unknown command will be ignored.
- The program must handle errors carefully. In no way can it quit in an unexpected manner.
- The coordinates of the rooms will always be int.

## Example

https://github.com/user-attachments/assets/970a2083-8dfc-4465-a90f-5cec738b613c

## USAGE

```
go run . FileName
```

To test the program with given bad and good input files, run test.sh by
`bash test.sh`

## Logic used to implement the project

1. Parse the file (data.go):

   - Read the input file and extract the necessary information.
   - Validate the data format to ensure it follows the expected structure.

2. Find all possible paths (path.go):

   - Use a recursive function to explore all possible paths from the start room to the end room.
   - Track the visited rooms to avoid loops and backtracking.
   - Store the paths obtained for further analysis.

3. Find all possible combinations of the paths (path.go):

   - Generate all possible combinations of the paths obtained in the previous step.
   - Each combination represents a set of paths that can be used concurrently by the ants.

4. Find the best combination (steps.go):

   - Calculate the number of steps required for each combination.
   - Choose the combination with the minimum number of steps.

5. Regulate the ants movement (moves.go):

   - Simulate the movement of the ants based on the best combination of paths.
   - Distribute the ants among the paths in the combination.
   - Move the ants from room to room according to the chosen paths, ensuring no conflicts occur.
   - Keep track of the final movement sequence and the rooms occupied by the ants.

6. Handle output (moves.go):
   - Output the final movement of the ants, indicating the rooms they occupy at each step.
   - Format the output according to the desired specifications.

## AUTHORS

- Fatema Alawadhi
- Tasneem Mearaj

## LICENSES

This program developed within the scope of Reboot.
