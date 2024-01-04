# Alem01 - Push-Swap

**Push-Swap** is a simple project centered around a Non-Comparative Sorting Algorithm. It involves two programs: `push-swap` and `checker`. 
The `push-swap` program determines and outputs the smallest sequence of instructions using the push-swap instruction language to sort 
integer arguments, while the `checker` program reads instructions from standard input and verifies if the integers are sorted.

## Instruction

Two stacks (`a` and `b`) and a set of instructions are available for sorting the stack `a` in ascending order:

- `pa`: Pushes the top first element of stack `b` to stack `a`.
- `pb`: Pushes the top first element of stack `a` to stack `b`.
- `sa`: Swaps the first 2 elements of stack `a`.
- `sb`: Swaps the first 2 elements of stack `b`.
- `ss`: Executes `sa` and `sb`.
- `ra`: Rotates stack `a` (shifts up all elements by 1).
- `rb`: Rotates stack `b`.
- `rr`: Executes `ra` and `rb`.
- `rra`: Reverse rotates `a` (shifts down all elements by 1).
- `rrb`: Reverse rotates `b`.
- `rrr`: Executes `rra` and `rrb`.

## Constraints
- No duplicate numbers.
- Valid integer-like strings separated by space.
- Sorted string numbers produce no output as well as empty input.
- Other cases are invalid and the program will output accordingly.

## Usage
Two programs are located in one program, just attached with tags while compiling.
To compile `push_swap` program, run:
```bash
make push_swap
```
To compile `checker` program, run:
```bash
make checker
```

Then, be free to use it, like:
```bash
./push_swap "0 4 3 5 6 7"
```
or like:
```bash
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
```
## Tech
Go version 1.20

Used only standard library


## References
[Jamie Dawson](https://medium.com/@jamierobertdawson/push-swap-the-least-amount-of-moves-with-two-stacks-d1e76a71789a)

[YYBer](https://medium.com/@YYBer/my-one-month-push-swap-journey-explore-an-easily-understand-and-efficient-algorithm-11449eb17752)

[Member of 42 School](https://www.youtube.com/@onaecO/community)


# Author
[@mrqwer](https://github.com/mrqwer)

