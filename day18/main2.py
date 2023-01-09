import time
import logging
import argparse
from enum import Enum

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
            filename=output_file,
            filemode='w'
        )

class Side(Enum):
    LEFT = 0,
    RIGHT = 1

class Pair:
    id = 0

    def __init__(self, parent, left=None, right=None):
        self.left = left
        self.right = right
        self.parent = parent
        self.id = Pair.id
        Pair.id += 1

    def plus(self, other):
        root = Pair(None)
        self_copy = self.copy(root)
        other_copy = other.copy(root)
        root.left = self_copy
        root.right = other_copy
        root.reduce()
        return root

    def copy(self, parent):
        copy = Pair(parent)
        if type(self.left) == int:
            copy.left = self.left
        else:
            copy.left = self.left.copy(copy)
        if type(self.right) == int:
            copy.right = self.right
        else:
            copy.right = self.right.copy(copy)
        return copy

    def can_explode(self, level):
        # if nested 4+, can explode
        if level >= 4 : return True
        for side in [self.left, self.right]:
            if type(side) == Pair:
                # logging.debug(f"{side.pretty()} is at level {level + 1}")
                if level >= 3:
                    # logging.debug(f"{self.pretty()} can explode on {side.pretty()} at level {level + 1}")
                    return True
                elif side.can_explode(level + 1):
                    # logging.debug(f"{self.pretty()} can explode because it contains something explodable")
                    return True
        # logging.debug(f"{self.pretty()} can't explode")
        return False

    def explode(self, level):
        if type(self.left) == Pair and level >= 3:
            self.left.add_to_previous(self.left.left, Side.LEFT)
            self.left.add_to_next(self.left.right, Side.LEFT)
            self.left = 0
        elif type(self.right) == Pair and level >= 3:
            self.right.add_to_previous(self.right.left, Side.RIGHT)
            self.right.add_to_next(self.right.right, Side.RIGHT)
            self.right = 0
        elif type(self.left) == Pair and self.left.can_explode(level + 1):
            self.left.explode(level + 1)
        elif type(self.right) == Pair and self.right.can_explode(level + 1):
            self.right.explode(level + 1)

    def add_to_previous(self, value, origin):
        # left is the hard one. need to go up at least two, then continue until I find a left that isn't where I just came from
        # then right until I come to an int
        # if I get to the root node without finding any left paths, then there's no value to add to so I'm done
        if origin == Side.LEFT:
            node = self.parent.parent
            previous_id = self.parent.id
            if type(node.left) == int:
                node.left += value
                return
            while node.left.id == previous_id:
                previous_id = node.id
                node = node.parent
                if node is None: return
                if type(node.left) == int:
                    node.left += value
                    return
            node = node.left
            while type(node.right) == Pair:
                node = node.right
            node.right += value
            return

        # if I'm the right, then go up to my parent then left
        # then right until I come to an int
        if origin == Side.RIGHT:
            node = self.parent
            if type(node.left) == int:
                node.left += value
                return
            node = node.left
            while type(node.right) == Pair:
                node = node.right
            node.right += value

    def add_to_next(self, value, origin):
        # if I'm the left, go up, right, then left until I come to an int
        if origin == Side.LEFT:
            node = self.parent
            if type(node.right) == int:
                node.right += value
                return
            node = node.right
            while type(node.left) == Pair:
                node = node.left
            node.left += value
            return

        # if I'm the right, go up at least twice, continuing until a parent has a right that isn't the path I just came up
        # in other words, until the parent isn't a left
        # then go one right, then left until I come to an int
        if origin == Side.RIGHT:
            node = self.parent.parent
            previous_id = self.parent.id
            if type(node.right) == int:
                node.right += value
                return
            while node.right.id == previous_id:
                previous_id = node.id
                node = node.parent
                if node is None: return
                if type(node.right) == int:
                    node.right += value
                    return
            node = node.right
            while type(node.left) == Pair:
                node = node.left
            node.left += value

    def can_split(self):
        # if any contained number >=10, can split
        for side in [self.left, self.right]:
            if type(side) == int and side >= 10:
                # logging.debug(f"{self.pretty()} can split because it contains {side}")
                return True
            elif type(side) == Pair and side.can_split():
                # logging.debug(f"{self.pretty()} can split because it contains something splittable")
                return True
        # logging.debug(f"{self.pretty()} can't split")
        return False

    def split(self):
        if type(self.left) == int and self.left >= 10:
            value = self.left
            splits = split_value(value)
            # logging.debug(f"Splitting left {self.left} into {splits}")
            self.left = Pair(self, splits[0], splits[1])
        elif type(self.left) == Pair and self.left.can_split():
            # logging.debug(f"Split - going left from {self.pretty()} to {self.left.pretty()}")
            self.left.split()
        elif type(self.right) == int and self.right >= 10:
            value = self.right
            splits = split_value(value)
            # logging.debug(f"Splitting right {self.right} into {splits}")
            self.right = Pair(self, splits[0], splits[1])
        elif type(self.right) == Pair and self.right.can_split():
            # logging.debug(f"Split - going right from {self.pretty()} to {self.right.pretty()}")
            self.right.split()

    def reduce(self):
        # always look for an explosion before looking for a split
        while self.can_explode(0) or self.can_split():
            while self.can_explode(0):
                # logging.debug(f"Before explode: {self.pretty()}")
                self.explode(0)
                # logging.debug(f"After explode:  {self.pretty()}")
            if self.can_split():
                # logging.debug(f"Before split:   {self.pretty()}")
                self.split()
                # logging.debug(f"After split:  {self.pretty()}")

    def magnitude(self):
        left_value = self.left if type(self.left) == int else self.left.magnitude()
        right_value = self.right if type(self.right) == int else self.right.magnitude()
        return (3 * left_value) + (2 * right_value)

    def pretty(self):
        left = self.left if type(self.left) == int else self.left.pretty() if type(self.left) == Pair else ''
        right = self.right if type(self.right) == int else self.right.pretty() if type(self.right) == Pair else ''
        return f"[{left},{right}]"

def split_value(value):
    # eg. 11 => [5, 6]
    left = value // 2
    right = value - left
    return left, right

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

    numbers = []
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
        numbers.append(line_root)
    # END for line in input data

    largest_magnitude = 0
    for outer in numbers:
        for inner in numbers:
            if outer.id == inner.id: continue
            result = outer.plus(inner)
            magnitude = result.magnitude()
            logging.debug(f"Result: {result.pretty()}")
            logging.debug(f"Magnitude: {magnitude}")
            if magnitude > largest_magnitude: largest_magnitude = magnitude
    logging.info(f"Largest magnitude: {largest_magnitude}")

# testsum:
# largest magnitude is 3993
# from adding:
#   [[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
# + [[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
# = [[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]
