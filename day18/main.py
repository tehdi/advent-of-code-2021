import time
import logging

logging.basicConfig(format='%(message)s', level=logging.INFO)

class Pair:
    def __init__(self, parent):
        self.left = None
        self.right = None
        self.parent = parent

    def pretty(self):
        left = self.left if type(self.left) == int else self.left.pretty() if type(self.left) == Pair else ''
        right = self.right if type(self.right) == int else self.right.pretty() if type(self.right) == Pair else ''
        return f"[{left}, {right}]"

    def magnitude(self):
        left_value = self.left if type(self.left) == int else self.left.magnitude()
        right_value = self.right if type(self.right) == int else self.right.magnitude()
        return (3 * left_value) + (2 * right_value)

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
                    logging.debug(f"{char} starting root")
                    root = Pair(None)
                    pair = root
                # start of pair, but is it left or right? the pair knows which it's missing
                elif pair.left is None:
                    logging.debug(f"{char} going left")
                    pair.left = Pair(pair)
                    pair = pair.left
                elif pair.right is None:
                    logging.debug(f"{char} going right")
                    pair.right = Pair(pair)
                    pair = pair.right
            elif char == ',':
                # end of left. could be '[7,' or could be '],'
                logging.debug(f"{char} end of left")
                if current_number != '':
                    pair.left = int(current_number)
                    logging.debug(f"  set left number {current_number}")
                    current_number = ''
            elif char == ']':
                # end of right = end of pair
                logging.debug(f"{char} end of right")
                if current_number != '':
                    logging.debug(f"  set right number {current_number}")
                    pair.right = int(current_number)
                    current_number = ''
                if pair.parent is None:
                    if char_index + 1 != len(line):
                        logging.warning(f"Warning! Possible data integrity issue: Parent is none which should mean EOL but char index is {char_index} vs line length {len(line)}")
                        pass
                else:
                    pair = pair.parent
            elif char.isnumeric():
                # number could be multiple digits so don't finalize it yet
                logging.debug(f"{char} number")
                current_number += char
            logging.debug(root.pretty())
            # time.sleep(0.5) # it's just really fun to watch it build the structure
        logging.info(f"Magnitude: {root.magnitude()}") # SUCCESS! testmagnitude_input expected: 4140

# for testsum_input, expected sum = [[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]
# that value => testmagnitude_input. expected_final_magnitude = 4140
