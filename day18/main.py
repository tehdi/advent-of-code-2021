import time
import logging
import argparse

def configure_logging(verbose, output_file):
    log_level = logging.DEBUG if verbose else logging.INFO
    if output_file is None:
        logging.basicConfig(
            format='%(message)s',
            level=log_level
        )
    else:
        logging.basicConfig(
            format='%(message)s',
            level=log_level,
            filename=output_file
        )

class Pair:
    def __init__(self, parent):
        self.left = None
        self.right = None
        self.parent = parent

    def plus(self, other):
        root = Pair(None)
        root.left = self
        root.right = other
        self.parent = root
        other.parent = root
        root.reduce()
        return root

    def reduce(self):
        pass

    def magnitude(self):
        left_value = self.left if type(self.left) == int else self.left.magnitude()
        right_value = self.right if type(self.right) == int else self.right.magnitude()
        return (3 * left_value) + (2 * right_value)

    def pretty(self):
        left = self.left if type(self.left) == int else self.left.pretty() if type(self.left) == Pair else ''
        right = self.right if type(self.right) == int else self.right.pretty() if type(self.right) == Pair else ''
        return f"[{left}, {right}]"

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('-i', '-f', '--input-file', default='input.txt')
    parser.add_argument('-o', '--output-file', default=None)
    parser.add_argument('-v', '--verbose', '--debug', default=False, action='store_true')
    args = parser.parse_args()
    configure_logging(args.verbose, args.output_file)

    filename = args.input_file
    with open(filename) as input_file:
        input_data = [line.rstrip('\n') for line in input_file]

    root = None
    for line in input_data:
        line_root = None
        pair = None
        current_number = ''
        for char_index,char in enumerate(line):
            if char == '[':
                if pair is None:
                    # logging.debug(f"{char} starting line_root")
                    line_root = Pair(None)
                    pair = line_root
                # start of pair, but is it left or right? the pair knows which it's missing
                elif pair.left is None:
                    # logging.debug(f"{char} going left")
                    pair.left = Pair(pair)
                    pair = pair.left
                elif pair.right is None:
                    # logging.debug(f"{char} going right")
                    pair.right = Pair(pair)
                    pair = pair.right
            elif char == ',':
                # end of left. could be '[7,' or could be '],'
                # logging.debug(f"{char} end of left")
                if current_number != '':
                    pair.left = int(current_number)
                    # logging.debug(f"  set left number {current_number}")
                    current_number = ''
            elif char == ']':
                # end of right = end of pair
                # logging.debug(f"{char} end of right")
                if current_number != '':
                    # logging.debug(f"  set right number {current_number}")
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
                # logging.debug(f"{char} number")
                current_number += char
            # logging.debug(line_root.pretty())
            # time.sleep(0.5) # it's just really fun to watch it build the structure
        # END for char in line
        if root is None:
            root = line_root
        else:
            root = root.plus(line_root)
    # END for line in input_data
    logging.debug(f"Sum: {root.pretty()}")
    logging.info(f"Magnitude: {root.magnitude()}")

# testsum, testmagnitude:
#   sum = [[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]
#   magnitude = 4140
# testsum2, testmagnitude2:
#   sum = [[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]
#   magnitude = 3488
