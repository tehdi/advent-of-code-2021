import time

class Pair:
    id = 0

    def __init__(self, parent):
        self.id = Pair.id
        self.left = None
        self.right = None
        self.parent = parent
        Pair.id += 1

    def pretty(self):
        left = self.left if type(self.left) == int else self.left.pretty() if type(self.left) == Pair else ''
        right = self.right if type(self.right) == int else self.right.pretty() if type(self.right) == Pair else ''
        # return f"{self.id}: [{left}, {right}]"
        return f"[{left}, {right}]"

    def magnitude(self):
        if type(self.left) == type(self.right) == int:
            return (3 * self.left) + (2 * self.right)
        if type(self.left) == int:
            return (3 * self.left) + (2 * self.right.magnitude())
        return (3 * self.left.magnitude()) + (2 * self.right.magnitude())

if __name__ == '__main__':
    with open('testmagnitude_input') as input_file:
        input_data = [line.rstrip('\n') for line in input_file]

    for line in input_data:
        root = None
        pair = None
        current_number = ''
        for char_index,char in enumerate(line):
            if char == '[':
                if pair is None:
                    # print(f"{char} starting root")
                    root = Pair(None)
                    pair = root
                # start of pair, but is it left or right? the pair knows which it's missing
                elif pair.left is None:
                    # print(f"{char} going left")
                    pair.left = Pair(pair)
                    pair = pair.left
                elif pair.right is None:
                    # print(f"{char} going right")
                    pair.right = Pair(pair)
                    pair = pair.right
            elif char == ',':
                # end of left. could be '[7,' or could be '],'
                # print(f"{char} end of left")
                if current_number != '':
                    pair.left = int(current_number)
                    # print(f"  set left number {current_number}")
                    current_number = ''
            elif char == ']':
                # end of right = end of pair
                # print(f"{char} end of right")
                if current_number != '':
                    # print(f"  set right number {current_number}")
                    pair.right = int(current_number)
                    current_number = ''
                if pair.parent is None:
                    if char_index + 1 != len(line):
                        # print(f"Parent is none which should mean EOL but char index is {char_index} vs line length {len(line)}")
                        pass
                else:
                    pair = pair.parent
            elif char.isnumeric():
                # number could be multiple digits so don't finalize it yet
                # print(f"{char} number")
                current_number += char
            # print(root.pretty())
            # time.sleep(0.5) # it's just really fun to watch it build the structure
        # print(root.magnitude()) # SUCCESS! testmagnitude_input expected: 4140

# for testsum_input, expected sum = [[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]
# that value => testmagnitude_input. expected_final_magnitude = 4140
