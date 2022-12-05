from collections import deque
from functools import reduce

newline = '\n'

hex_to_bin = {
    '0': '0000',
    '1': '0001',
    '2': '0010',
    '3': '0011',
    '4': '0100',
    '5': '0101',
    '6': '0110',
    '7': '0111',
    '8': '1000',
    '9': '1001',
    'A': '1010',
    'B': '1011',
    'C': '1100',
    'D': '1101',
    'E': '1110',
    'F': '1111'
}

def multiply(factors):
    return reduce(lambda x,y: x*y, factors)
def greater_than(values):
    return 1 if values[0] > values[1] else 0
def less_than(values):
    return 1 if values[0] < values[1] else 0
def equal_to(values):
    return 1 if values[0] == values[1] else 0

type_definitions = {
    '0': sum,
    '1': multiply,
    '2': min,
    '3': max,
    '4': 'value',
    '5': greater_than,
    '6': less_than,
    '7': equal_to
}

def generate_packet_id():
    next_packet_id = 1
    while True:
        yield next_packet_id
        next_packet_id += 1
packet_id = generate_packet_id()

class Packet:
    def __init__(self, version, packet_type):
        self.id = next(packet_id)
        self.version = version
        self.packet_type = packet_type
        self.value = ''
        self.children = []
    
    def add_children(self, children):
        self.children.extend(children)
    
    def value_decimal(self):
        if self.value != '':
            return int(self.value, 2)
        return 'n/a'

    def is_value(self):
        return self.packet_type == '4'

    def pprint(self):
        pretty = f"Packet {self.id}: type {self.packet_type}, value '{self.value}' ({self.value_decimal()}), children {len(self.children)}"
        for child in self.children:
            pretty += f"{newline}Child of {self.id}: {child.pprint()}"
        return pretty


def next_i_bits_string(i, binary):
    return ''.join(binary.popleft() for j in range(i))

def next_i_bits_deque(i, binary):
    return deque(next_i_bits_string(i, binary))

def next_i_bits_decimal(i, binary):
    return int(next_i_bits_string(i, binary), 2)

def parse_packets(binary, max=0):
    # print(f"{len(binary)} bits or {max} packets to process")
    packets = []
    while (max == 0 or len(packets) < max) and len(binary) > 6 and '1' in binary:
        version = next_i_bits_decimal(3, binary)
        packet_type = str(next_i_bits_decimal(3, binary))
        # print(f"Packet version {version}, type {packet_type}")
        packet = Packet(version, packet_type)
        packets.append(packet)
        if packet_type == '4':
            # print(f" literal value packet")
            while True:
                prefix = next_i_bits_string(1, binary)
                packet.value += next_i_bits_string(4, binary)
                if prefix == '0':
                    break
        else:
            # print(f" operator packet")
            length_type_id = next_i_bits_string(1, binary)
            if length_type_id == '0':
                # next 15 bits tells me how many bits are in this packet's subs
                subpacket_bit_count = next_i_bits_decimal(15, binary)
                subpacket_bits = next_i_bits_deque(subpacket_bit_count, binary)
                # print(f"  type 0 with {subpacket_bit_count} sub bits")
                packet.add_children(parse_packets(subpacket_bits))
            elif length_type_id == '1':
                # next 11 bits tell me how many subs this packet has
                subpacket_count = next_i_bits_decimal(11, binary)
                # print(f"  type 1 with {subpacket_count} subs")
                packet.add_children(parse_packets(binary, subpacket_count))
            else:
                print(f"Unrecognized length type ID {length_type_id}")
    return packets

def calculate_packet(packet):
    if packet.is_value(): return packet.value_decimal()
    
    children_values = []
    for child in packet.children:
        children_values.append(calculate_packet(child))
    return type_definitions[packet.packet_type](children_values)

if __name__ == '__main__':
    with open('input', 'r') as input_file:
        input_data = [line.rstrip('\n') for line in input_file]
    binary = deque(''.join(hex_to_bin[char] for line in input_data for char in line))
    packets = parse_packets(binary)

    # will only be 1 but I'm reusing parse_packets so it's easier to have a list the whole way
    for packet in packets:
        print(calculate_packet(packet))
