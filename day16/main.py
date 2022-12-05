from collections import deque

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

class Packet:
    def __init__(self, version, packet_type):
        self.version = version
        self.packet_type = packet_type


def next_i_bits_string(i, binary):
    return ''.join(binary.popleft() for j in range(i))

def next_i_bits_deque(i, binary):
    return deque(next_i_bits_string(i, binary))

def next_i_bits_decimal(i, binary):
    return int(next_i_bits_string(i, binary), 2)

def parse_packets(binary):
    packets = []
    while len(binary) > 6 and '1' in binary:
        version = next_i_bits_decimal(3, binary)
        packet_type = next_i_bits_decimal(3, binary)
        print(f"Packet version {version}, type {packet_type}")
        packets.append(Packet(version, packet_type))
        if packet_type == 4:
            print(f" literal value packet")
            while True:
                prefix = next_i_bits_string(1, binary)
                discard_for_now = next_i_bits_string(4, binary)
                if prefix == '0':
                    break
        else:
            print(f" operator packet")
            length_type_id = next_i_bits_string(1, binary)
            if length_type_id == '0':
                # next 15 bits tells me how many bits are in this packet's subs
                subpacket_bit_count = next_i_bits_decimal(15, binary)
                subpacket_bits = next_i_bits_deque(subpacket_bit_count, binary)
                print(f"  type 0 with {subpacket_bit_count} sub bits")
                packets.extend(parse_packets(subpacket_bits))
            elif length_type_id == '1':
                # next 11 bits tell me how many subs this packet has
                subpacket_count = next_i_bits_decimal(11, binary)
                print(f"  type 1 with {subpacket_count} subs")
                for i in range(subpacket_count):
                    packets.extend(parse_packets(binary))
            else:
                print(f"Unrecognized length type ID {length_type_id}")
    return packets


if __name__ == '__main__':
    with open('input') as input_file:
        input_data = [line.rstrip('\n') for line in input_file]
    binary = deque(''.join(hex_to_bin[char] for line in input_data for char in line))

    packets = parse_packets(binary)
    print(sum([packet.version for packet in packets]))
