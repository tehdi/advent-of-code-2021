# basin: { 'id': next(basin_id), 'locations': [(myLineIndex, myCharIndex)]}

BASINS = {}
def generate_basin_id():
    next_basin_id = 1
    while True:
        yield next_basin_id
        next_basin_id += 1
basin_id = generate_basin_id()

HIGH_POINTS = { 'id': next(basin_id) }

def first_valid_basin(basins):
    for basin in basins:
        if basin != None and basin != HIGH_POINTS:
            return basin
    return None

def merge_basins(old_basin, new_basin):
    for location in old_basin['locations']:
        BASINS[location] = new_basin
    new_basin['locations'].extend(old_basin['locations'])
    print(f"Merged basin {old_basin['id']} into {new_basin['id']}. Basin {old_basin['id']} should not reappear")
    return new_basin

def count_high_null(basins):
    count = 0
    for basin in basins:
        if basin == None or basin == HIGH_POINTS:
            count += 1
    return count


def what_basin_am_i_in(line_index, char_index, value='X'):
    # already tagged?
    if (line_index, char_index) in BASINS:
        return BASINS[(line_index, char_index)]
    
    left_basin = up_basin = None
    # what does my buddy to the left think?
    if char_index > 0:
        left_basin = what_basin_am_i_in(line_index, char_index - 1)
    # what does my buddy above me think?
    if line_index > 0:
        up_basin = what_basin_am_i_in(line_index - 1, char_index)
    
    high_null_neighbour_count = count_high_null([left_basin, up_basin])
    
    # SCENARIO: both neighbours are highPoints OR both are null (first char of the first line) OR it's one of each => create a new basin
    # 091239 <-- 0 and 1 are in this situation
    # 9456789
    #  ^4 is in this situation (don't worry, it'll get sorted when 5 sees that 1 and 4 disagree)
    if high_null_neighbour_count == 2:
        # return new Basin(myLineIndex, myCharIndex)
        print(f"I think I'm in a void as a {value} at {line_index}, {char_index}")
        return { 'id': next(basin_id), 'locations': []}

    # SCENARIO: one has a basin value (other is a high point or null (first line or first char in a line)) => use the non-high, non-null basin
    # 091239 <-- 2 and 3 are in this situation
    # 9456789
    #      ^8 is in this situation (left = basin, up = high point)
    if high_null_neighbour_count == 1:
        print(f"I think I have one neighbour between {left_basin} and {up_basin} as a {value} at {line_index}, {char_index}")
        return first_valid_basin([left_basin, up_basin])
    
    # SCENARIO: both have basin values and they don't match => use the (left? up? it doesn't matter as long as I cascade the change through the whole basin)
    # 091239
    # 9456789
    #   ^5 is in this situation because 4 didn't see a basin to the left or up
    if left_basin != up_basin:
        print(f"I think I'm in a mismatch between basins {left_basin['id']} and {up_basin['id']} as a {value} at {line_index}, {char_index}")
        return merge_basins(left_basin, up_basin)

    # SCENARIO: both have basin values and they match
    # 091239
    # 9456789
    #    ^6 is in this situation
    if left_basin == up_basin:
        print(f"I think everyone agrees we're in basin {left_basin['id']} as a {value} at {line_index}, {char_index}")
        return left_basin  # or upBasin. doesn't matter.

    # that SHOULD have been the only remaining scenario?
    raise Exception('you done fucked up')


if __name__ == '__main__':
    with open('day09_input') as input_file:
        input_data = [line.rstrip('\n') for line in input_file]

    for line_index,line in enumerate(input_data):
        for char_index,char in enumerate(line):
            if char == '9':
                print(f"I'm a 9! at {line_index}, {char_index}")
                BASINS[(line_index, char_index)] = HIGH_POINTS
            else:
                in_basin = what_basin_am_i_in(line_index, char_index, char)
                print(f"Adding myself ({char} at {line_index}, {char_index}) to basin {in_basin['id']}")
                in_basin['locations'].append((line_index, char_index))
                BASINS[(line_index, char_index)] = in_basin

    all_basins = {}
    for basin in BASINS.values():
        if basin != HIGH_POINTS and basin['id'] not in all_basins:
            all_basins[basin['id']] = len(basin['locations'])
    print(f'Ended up with {len(all_basins)} basins')
    for basin_id,location_count in all_basins.items():
        print(f"Basin {basin_id} has {location_count} locations")
    by_size = sorted(all_basins.values(), reverse=True)
    print(by_size)
